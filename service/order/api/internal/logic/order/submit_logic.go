package order

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/message"
	"goms/common/request"
	"goms/common/response"
	"goms/common/response/errcode/ordercode"
	"goms/service/order/api/internal/svc"
	"goms/service/order/api/internal/types"
	"goms/service/product/rpc/productclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitLogic {
	return &SubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitLogic) Submit(req *types.SubmitReq) (err error) {

	// 解析用户ID
	userId, err := request.ParseUserId(l.ctx)
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "user id parse failed"))
		err = response.ErrResp(1, ordercode.Submit, response.InternalError, err.Error())
		return
	}

	// 转换数据
	ids := make([]int64, len(req.Products))
	products := make([]message.OrderProduct, len(req.Products))
	for i, item := range products {
		ids[i] = item.ID
		products[i] = message.OrderProduct{
			ID:     item.ID,
			Amount: item.Amount,
		}
	}

	// 调用RPC服务
	reply, err := l.svcCtx.ProductRpc.ListByIds(l.ctx, &productclient.ListByIdsReq{
		Ids: ids,
	})
	if err != nil {
		switch s, _ := status.FromError(err); s.Code() {
		case codes.Aborted:
			err = response.ErrResp(2, ordercode.Submit, response.InternalError, s.Message())
		default:
			l.Logger.Error(errors.Wrap(err, "product rpc call failed"))
			err = response.ErrResp(3, ordercode.Submit, response.ServiceError, s.Message())
		}
		return
	}

	// 回写产品价格
	for i, item := range reply.List {
		products[i].Price = item.Price
	}

	// 构建Kafka消息
	msg, err := json.Marshal(message.KqOrderCreateMsg{
		UserID:    userId,
		Products:  products,
		Consignee: req.Consignee,
		Phone:     req.Phone,
		Address:   req.Address,
	})
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "marshal json failed"))
		err = response.ErrResp(2, ordercode.Submit, response.InternalError, err.Error())
		return
	}

	// 推送订单创建消息到Kafka
	if err = l.svcCtx.OrderCreatePusher.Push(string(msg)); err != nil {
		l.Logger.Error(errors.Wrap(err, "push kafka failed"))
		err = response.ErrResp(3, ordercode.Submit, response.InternalError, err.Error())
		return
	}

	// 推送订单支付超时到Asynq

	return
}
