package svc

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/common/validator"
	"goms/service/product/api/internal/config"
	"goms/service/product/rpc/productclient"
)

type ServiceContext struct {
	Config config.Config

	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 设置V9校验
	httpx.SetValidator(validator.NewV9())

	return &ServiceContext{
		Config: c,

		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpcConf)),
	}
}
