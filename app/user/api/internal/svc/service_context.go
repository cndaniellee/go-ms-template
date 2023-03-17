package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goms/app/user/api/internal/config"
	"goms/app/user/rpc/userclient"
	"gopkg.in/go-playground/validator.v9"
)

type ServiceContext struct {
	Config   config.Config
	Validate *validator.Validate

	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Validate: validator.New(),

		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
