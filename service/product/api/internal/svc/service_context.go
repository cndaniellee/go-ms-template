package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/common/middleware"
	"goms/service/product/api/internal/config"
	"goms/service/product/rpc/productclient"
)

type ServiceContext struct {
	Config config.Config

	AuthConvertor rest.Middleware

	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,

		AuthConvertor: middleware.NewAuthConvertor().Handle,

		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpcConf)),
	}
}
