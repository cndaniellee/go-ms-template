package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/common/jwtauth"
)

type Config struct {
	rest.RestConf

	JwtAuth    jwtauth.JwtAuthConf
	JwtRefresh jwtauth.JwtAuthConf

	Redis redis.RedisConf

	UserRpcConf zrpc.RpcClientConf
}
