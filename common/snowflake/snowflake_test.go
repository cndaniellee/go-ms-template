package snowflake

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"testing"
	"time"
)

func TestGetWorkerId(t *testing.T) {
	hKey := fmt.Sprintf("snowflake:%d", 31)
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: "127.0.0.1:6379",
		Type: "node",
		Pass: "",
		Tls:  false,
	})
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 60; i++ {
		workerId, err := getWorkerId(rds, hKey)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("WorkerId: ", workerId)
		if err = rds.Hset(hKey, strconv.FormatInt(workerId, 10), strconv.FormatInt(time.Now().Add(KeepAlivePeriod+GracePeriod).UnixMilli(), 10)); err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second * 3)
	}
}
