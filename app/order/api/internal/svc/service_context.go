package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goms/app/order/api/internal/config"
	"goms/app/order/rpc/orderclient"
	"goms/app/product/rpc/productclient"
	"gopkg.in/go-playground/validator.v9"
)

type ServiceContext struct {
	Config   config.Config
	Validate *validator.Validate

	OrderRpc   orderclient.Order
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Validate: validator.New(),

		OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpcConf)),
	}
}
