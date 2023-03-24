package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"goms/service/product/rpc/internal/svc"
	"goms/service/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *product.IdReq) (*product.DetailReply, error) {

	// 获取产品
	p, err := l.svcCtx.ProductModel.FindById(l.ctx, in.Id)
	switch err {
	case gorm.ErrRecordNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case nil:
		break
	default:
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &product.DetailReply{
		Id:          p.ID,
		Title:       p.Title,
		Category:    int32(p.Category),
		Stock:       p.Stock,
		Price:       p.Price,
		Description: p.Description,
		CreatedAt:   p.CreatedAt.UnixMilli(),
	}, nil
}
