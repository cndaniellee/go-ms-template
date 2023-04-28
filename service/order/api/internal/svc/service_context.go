package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/common/middleware"
	"goms/service/order/api/internal/config"
	"goms/service/order/rpc/orderclient"
	"goms/service/product/rpc/productclient"
)

type ServiceContext struct {
	Config config.Config

	AuthConvertor rest.Middleware

	OrderRpc   orderclient.Order
	ProductRpc productclient.Product

	OrderCreatePusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,

		AuthConvertor: middleware.NewAuthConvertor().Handle,

		OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpcConf)),

		OrderCreatePusher: kq.NewPusher(c.OrderCreateConf.Brokers, c.OrderCreateConf.Topic),
	}
}
