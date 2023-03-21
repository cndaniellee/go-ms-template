package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goms/app/product/api/internal/config"
	"goms/app/product/rpc/productclient"
	"gopkg.in/go-playground/validator.v9"
)

type ServiceContext struct {
	Config   config.Config
	Validate *validator.Validate

	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Validate: validator.New(),

		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpcConf)),
	}
}
