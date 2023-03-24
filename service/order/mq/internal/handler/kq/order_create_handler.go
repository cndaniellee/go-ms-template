package kq

import (
	"goms/service/order/mq/internal/svc"
)

type OrderCreateHandler struct {
	svcCtx *svc.ServiceContext
}

func NewOrderCreateHandler(svcCtx *svc.ServiceContext) *OrderCreateHandler {
	return &OrderCreateHandler{
		svcCtx: svcCtx,
	}
}

func (l *OrderCreateHandler) Consume(key, val string) error {

	return nil
}
