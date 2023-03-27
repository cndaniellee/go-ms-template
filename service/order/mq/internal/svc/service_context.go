package svc

import (
	"github.com/dtm-labs/client/dtmgrpc/dtmgimp"
	"github.com/dtm-labs/client/dtmgrpc/dtmgpb"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"goms/service/order/mq/internal/config"
	"goms/service/order/rpc/orderclient"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc orderclient.Order

	DtmClient dtmgpb.DtmClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 建立DTM连接
	conn, err := dtmgimp.GetGrpcConn(c.DtmService, false)
	logx.Must(err)

	return &ServiceContext{
		Config: c,

		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),

		DtmClient: dtmgpb.NewDtmClient(conn),
	}
}
