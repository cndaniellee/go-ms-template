package order

import (
	"context"

	"goms/app/order/api/internal/svc"
	"goms/app/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/response"
	"goms/common/response/errcode/ordercode"
)

type PaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentLogic {
	return &PaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentLogic) Payment(req *types.PaymentReq) (resp *types.PaymentResp, err error) {
	// todo: add your logic here and delete this line

	err = response.ErrResp(0, ordercode.Payment, response.InternalError, "example")

	return
}
