package svc

import (
	"github.com/dtm-labs/client/dtmgrpc/dtmgimp"
	"github.com/dtm-labs/client/dtmgrpc/dtmgpb"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/app/order/api/internal/config"
	"goms/app/order/rpc/orderclient"
	"goms/app/product/rpc/productclient"
	"goms/common/validator"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc   orderclient.Order
	ProductRpc productclient.Product

	DtmClient dtmgpb.DtmClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 设置V9校验
	httpx.SetValidator(validator.NewV9())

	// 建立DTM连接
	conn, err := dtmgimp.GetGrpcConn(c.DtmService, false)
	logx.Must(err)

	return &ServiceContext{
		Config: c,

		OrderRpc:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpcConf)),

		DtmClient: dtmgpb.NewDtmClient(conn),
	}
}
