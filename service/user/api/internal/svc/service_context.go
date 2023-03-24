package svc

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/common/validator"
	"goms/service/user/api/internal/config"
	"goms/service/user/rpc/userclient"
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
