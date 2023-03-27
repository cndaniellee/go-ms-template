package handler

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/message"
	"goms/service/order/mq/internal/svc"
	"goms/service/order/rpc/orderclient"
)

type PaymentTimeoutHandler struct {
	svcCtx *svc.ServiceContext
}

func NewPaymentTimeoutHandler(svcCtx *svc.ServiceContext) *PaymentTimeoutHandler {
	return &PaymentTimeoutHandler{
		svcCtx: svcCtx,
	}
}

func (l *PaymentTimeoutHandler) ProcessTask(ctx context.Context, task *asynq.Task) error {

	// 解析消息
	var msg message.DqOrderIdMsg
	if err := json.Unmarshal(task.Payload(), &msg); err != nil {
		logx.WithContext(ctx).Error(errors.Wrap(err, "unmarshal json failed"))
		return err
	}

	// 调用RPC服务
	if _, err := l.svcCtx.OrderRpc.CheckPaymentTimeout(ctx, &orderclient.IdReq{
		Id: msg.OrderID,
	}); err != nil {
		logx.WithContext(ctx).Error(errors.Wrap(err, "order rpc call failed"))
		return err
	}

	return nil
}
