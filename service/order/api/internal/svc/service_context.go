package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/common/validator"
	"goms/service/order/api/internal/config"
	"goms/service/order/rpc/orderclient"
	"goms/service/product/rpc/productclient"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc   orderclient.Order
	ProductRpc productclient.Product

	OrderCreatePusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 设置V9校验
	httpx.SetValidator(validator.NewV9())

	return &ServiceContext{
		Config: c,

		OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpcConf)),

		OrderCreatePusher: kq.NewPusher(c.OrderCreateConf.Brokers, c.OrderCreateConf.Topic),
	}
}
