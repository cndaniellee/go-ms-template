package svc

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/app/user/api/internal/config"
	"goms/app/user/rpc/userclient"
	"goms/common/validator"
)

type ServiceContext struct {
	Config config.Config

	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	httpx.SetValidator(validator.NewV9())

	return &ServiceContext{
		Config: c,

		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
