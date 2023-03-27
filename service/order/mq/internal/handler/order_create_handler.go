package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/CNDanielLee/go-utils"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/message"
	"goms/service/order/mq/internal/svc"
	"goms/service/order/rpc/orderclient"
	"goms/service/order/rpc/pb/order"
	"goms/service/product/rpc/pb/product"
	"goms/service/product/rpc/productclient"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type OrderCreateHandler struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderCreateHandler(svcCtx *svc.ServiceContext) *OrderCreateHandler {
	ctx := context.Background()
	return &OrderCreateHandler{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderCreateHandler) Consume(_, val string) error {

	// 解析消息
	var msg message.KqOrderCreateMsg
	if err := json.Unmarshal([]byte(val), &msg); err != nil {
		l.Logger.Error(errors.Wrap(err, "unmarshal json failed"))
		return err
	}

	// 构建RPC连接信息
	orderRpc, err := l.svcCtx.Config.OrderRpcConf.BuildTarget()
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "order rpc build failed"))
		return err
	}
	productRpc, err := l.svcCtx.Config.ProductRpcConf.BuildTarget()
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "product rpc build failed"))
		return err
	}

	// 构建Saga事务
	reply, err := l.svcCtx.DtmClient.NewGid(l.ctx, &emptypb.Empty{})
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "dtm gid fetch failed"))
		return err
	}
	saga := dtmgrpc.NewSagaGrpc(l.svcCtx.Config.DtmService, reply.Gid)

	// 依次添加每个产品的扣减事务
	products := make([]*orderclient.Product, len(msg.Products))
	for i, item := range msg.Products {
		saga = saga.Add(
			fmt.Sprint(productRpc, product.Product_Deduct_FullMethodName),
			fmt.Sprint(productRpc, product.Product_DeductRollback_FullMethodName),
			&productclient.DeductReq{
				Id:     item.ID,
				Amount: item.Amount,
			})
		// 转换数据
		products[i] = &orderclient.Product{
			Id:     item.ID,
			Amount: item.Amount,
			Price:  item.Price,
		}
	}

	// 添加订单创建事务
	saga = saga.Add(
		fmt.Sprint(orderRpc, order.Order_Create_FullMethodName),
		fmt.Sprint(orderRpc, order.Order_CreateRollback_FullMethodName),
		&orderclient.CreateReq{
			UserId:    msg.UserID,
			RefNo:     fmt.Sprint(time.Now().Format("20060102150405"), utils.RandStr(6, utils.Number)),
			Products:  products,
			Consignee: msg.Consignee,
			Phone:     msg.Phone,
			Address:   msg.Address,
		})

	// 提交Saga事务
	if err = saga.Submit(); err != nil {
		l.Logger.Error(errors.Wrap(err, "dtm saga submit failed"))
		return err
	}

	return nil
}
