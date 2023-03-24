package dq

import (
	"context"
	"github.com/hibiken/asynq"
	"goms/service/order/mq/internal/svc"
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

	return nil
}
