package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf

	Redis redis.RedisConf

	OrderRpcConf   zrpc.RpcClientConf
	ProductRpcConf zrpc.RpcClientConf

	DtmService string

	OrderCreateConf kq.KqConf
}
