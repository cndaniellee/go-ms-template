package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/common/jwtauth"
	"goms/common/middleware"
	"goms/service/auth/api/internal/config"
	"goms/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	Redis *redis.Redis

	TokenValidator rest.Middleware
	TokenGenerator *jwtauth.TokenGenerator

	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 初始化缓存
	rds, err := redis.NewRedis(c.Redis)
	logx.Must(err)

	return &ServiceContext{
		Config: c,

		TokenValidator: middleware.NewTokenValidator(rds).Handle,
		TokenGenerator: jwtauth.NewTokenGenerator(rds, c.JwtAuth, c.JwtRefresh),

		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
