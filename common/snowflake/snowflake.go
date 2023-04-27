package snowflake

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go.uber.org/atomic"
	"goms/common/lock"
	"sort"
	"strconv"
	"time"
)

const (
	// CacheServiceKey 容器列表Key
	CacheServiceKey = "snowflake:service:%d"
	// CacheLockKey 服务加锁Key
	CacheLockKey = "snowflake:lock:%d"
	// KeepAlivePeriod 存活更新时间
	KeepAlivePeriod = time.Second * 10
	// GracePeriod 检查宽限时间
	GracePeriod = time.Second * 5
)

type SnowFlake struct {
	rds  *redis.Redis
	hKey string

	epoch     int64 // 起始时间戳
	timestamp int64 // 当前时间戳，毫秒
	serviceId int64 // 数据中心机房ID
	workerId  int64 // 机器ID

	timestampBits  int64 // 时间戳占用位数
	serviceIdBits  int64 // 数据中心id所占位数
	workerIdBits   int64 // 机器id所占位数
	sequenceBits   int64 // 序列所占的位数
	sequenceMask   int64 // 生成序列的掩码最大值
	workerIdShift  int64 // 机器id左移偏移量
	centerIdShift  int64 // 数据中心机房id左移偏移量
	timestampShift int64 // 时间戳左移偏移量
	maxTimeStamp   int64 // 最大支持的时间

	sequence      *atomic.Int64 // 毫秒内序列号
	lastTimestamp *atomic.Int64 // 上一次生成ID的时间戳
}

func New(rds *redis.Redis, serviceId int64) (sf *SnowFlake, err error) {

	// 加锁
	ctx := context.Background()
	rl := lock.NewRedisLock(rds, fmt.Sprintf(CacheLockKey, serviceId), 5)
	if err = rl.AcquireExCtx(ctx); err != nil {
		return
	}
	// 解锁
	defer func(rl *lock.Lock, ctx context.Context) {
		err = rl.ReleaseExCtx(ctx)
	}(rl, ctx)

	// Redis中以Map存储当前服务下的已使用ID+过期时间
	hKey := fmt.Sprintf(CacheServiceKey, serviceId)

	workerId, err := getWorkerId(rds, hKey)
	if err != nil {
		return nil, err
	}

	sf = &SnowFlake{rds: rds, hKey: hKey}
	sf.epoch = int64(1672502400000) //设置起始时间戳：2023-01-01 00:00:00
	sf.serviceId = serviceId
	sf.workerId = workerId
	sf.serviceIdBits = 5  // 支持的最大服务ID占位数，最大是31
	sf.workerIdBits = 6   // 支持的最大容器ID占位数，最大是63
	sf.timestampBits = 41 // 时间戳占用位数
	sf.maxTimeStamp = -1 ^ (-1 << sf.timestampBits)

	maxServiceId := -1 ^ (-1 << sf.serviceIdBits)
	maxWorkerId := -1 ^ (-1 << sf.workerIdBits)

	// 参数校验
	if int(serviceId) > maxServiceId || serviceId < 0 {
		return nil, errors.New(fmt.Sprintf("serviceId can't be greater than %d or less than 0", maxServiceId))
	}
	if int(workerId) > maxWorkerId || workerId < 0 {
		return nil, errors.New(fmt.Sprintf("workerId can't be greater than %d or less than 0", maxWorkerId))
	}

	sf.sequenceBits = 11                                                     // 序列在ID中占的位数（1毫秒中生成的），最大为2047
	sf.sequenceMask = -1 ^ (-1 << sf.sequenceBits)                           // 计算毫秒内，最大的序列号
	sf.workerIdShift = sf.sequenceBits                                       // 机器ID向左移11位
	sf.centerIdShift = sf.sequenceBits + sf.workerIdBits                     // 机房ID向左移17位
	sf.timestampShift = sf.sequenceBits + sf.workerIdBits + sf.serviceIdBits // 时间截向左移22位

	sf.sequence = atomic.NewInt64(-1)
	sf.lastTimestamp = atomic.NewInt64(-1) // 上次生成 ID 的时间戳

	// 开携程更新存活时间
	go sf.keepAlive()

	return
}

func getWorkerId(rds *redis.Redis, hKey string) (int64, error) {
	all, err := rds.Hgetall(hKey)
	if err != nil {
		return 0, err
	}
	var workerId int64
	if len(all) > 0 {
		// 获取所有有效的ID
		occupiedIds := make([]int64, 0, len(all))
		now := time.Now()
		for k, v := range all {
			var parse int64
			// 解析租约
			parse, err = strconv.ParseInt(v, 10, 64)
			if err != nil {
				return 0, err
			}
			// 此ID的租约已到期
			if time.UnixMilli(parse).Before(now) {
				if _, err = rds.Hdel(hKey, k); err != nil {
					return 0, err
				}
			} else {
				// 解析ID
				parse, err = strconv.ParseInt(k, 10, 64)
				if err != nil {
					return 0, err
				}
				occupiedIds = append(occupiedIds, parse)
			}
		}

		if len(occupiedIds) > 0 {
			// 将ID排序
			sort.Sort(Int64Slice(occupiedIds))
			// 找到可用的ID
			var occupiedIndex int
			for {
				if workerId < occupiedIds[occupiedIndex] {
					break
				} else if workerId == occupiedIds[occupiedIndex] {
					if occupiedIndex < len(occupiedIds)-1 {
						occupiedIndex++
					}
				} else {
					break
				}
				workerId++
			}
		}
	}
	return workerId, nil
}

func (s *SnowFlake) NextId() int64 {
	now := time.Now().UnixMilli()     // 获取当前时间戳
	if now < s.lastTimestamp.Load() { // 如果当前时间小于上一次 ID 生成的时间戳，说明发生时钟回拨
		panic(fmt.Sprintf("clock moved backwards. Refusing to generate id for %d milliseconds", s.lastTimestamp.Load()-now))
	}

	t := now - s.epoch
	if t > s.maxTimeStamp {
		panic(fmt.Sprintf("epoch must be between 0 and %d", s.maxTimeStamp-1))
	}

	// 同一时间生成的，则序号+1
	if s.lastTimestamp.Load() == now {
		s.sequence.Store(s.sequence.Inc() & s.sequenceMask)
		// 毫秒内序列溢出：超过最大值; 阻塞到下一个毫秒，获得新的时间戳
		if s.sequence.Load() == 0 {
			for now <= s.lastTimestamp.Load() {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence.Store(0) // 时间戳改变，序列重置
	}
	// 保存本次的时间戳
	s.lastTimestamp.Store(now)

	// 根据偏移量，向左位移达到
	return (t << s.timestampShift) | (s.serviceId << s.centerIdShift) | (s.workerId << s.workerIdShift) | s.sequence.Load()
}

func (s *SnowFlake) keepAlive() {
	for {
		if err := s.rds.Hset(s.hKey, strconv.FormatInt(s.workerId, 10), strconv.FormatInt(time.Now().Add(KeepAlivePeriod+GracePeriod).UnixMilli(), 10)); err != nil {
			logx.Error("snowflake keep alive failed", err)
			return
		}
		time.Sleep(KeepAlivePeriod)
	}
}

// Int64Slice ID排序
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
