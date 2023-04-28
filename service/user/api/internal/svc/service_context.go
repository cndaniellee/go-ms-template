package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/common/middleware"
	"goms/service/user/api/internal/config"
	"goms/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	AuthConvertor rest.Middleware

	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,

		AuthConvertor: middleware.NewAuthConvertor().Handle,

		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
