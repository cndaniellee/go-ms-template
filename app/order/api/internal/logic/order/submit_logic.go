package order

import (
	"context"

	"goms/app/order/api/internal/svc"
	"goms/app/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/response"
	"goms/common/response/errcode/ordercode"
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

func (l *SubmitLogic) Submit(req *types.SubmitReq) error {
	// todo: add your logic here and delete this line

	err = response.ErrResp(0, ordercode.Submit, response.InternalError, "example")

	return nil
}
