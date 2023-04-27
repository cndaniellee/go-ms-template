package lock

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

type Lock struct {
	*redis.RedisLock
	maxLock int
}

func NewRedisLock(rds *redis.Redis, key string, second int) *Lock {
	rl := &Lock{
		maxLock: second,
	}
	rl.RedisLock = redis.NewRedisLock(rds, key)
	rl.SetExpire(second)
	return rl
}

func (rl *Lock) AcquireExCtx(ctx context.Context) error {
	timeout := time.Now().Add(time.Second * time.Duration(rl.maxLock))
	for {
		result, err := rl.AcquireCtx(ctx)
		if err != nil {
			return err
		}
		if result {
			break
		}
		if timeout.Before(time.Now()) {
			return errors.New("acquire timeout")
		}
		time.Sleep(time.Millisecond * 50)
	}
	return nil
}

func (rl *Lock) ReleaseExCtx(ctx context.Context) error {
	timeout := time.Now().Add(time.Second * time.Duration(rl.maxLock))
	for {
		result, err := rl.ReleaseCtx(ctx)
		if err != nil {
			return err
		}
		if result {
			break
		}
		if timeout.Before(time.Now()) {
			return errors.New("release timeout")
		}
		time.Sleep(time.Millisecond * 50)
	}
	return nil
}
