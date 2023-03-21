package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goms/app/product/rpc/internal/svc"
	"goms/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *product.IdReq) (*product.Empty, error) {

	// 删除产品
	if err := l.svcCtx.ProductModel.Delete(l.ctx, in.Id); err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &product.Empty{}, nil
}
