package logic

import (
	"context"
	"goms/service/product/rpc/model"
	"goms/service/product/rpc/model/enum"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goms/service/product/rpc/internal/svc"
	"goms/service/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditLogic) Edit(in *product.EditReq) (*product.IdReply, error) {

	// 新增或编辑产品
	p := &model.Product{
		Title:       in.Title,
		Category:    enum.ProductCategory(in.Category),
		Stock:       in.Stock,
		Price:       in.Price,
		Description: in.Description,
	}
	// 0：新增，Other：编辑
	p.ID = in.Id
	if err := l.svcCtx.ProductModel.Upsert(l.ctx, p); err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	// ES：更新
	if err := l.svcCtx.ProductES.Upsert(l.ctx, p); err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &product.IdReply{Id: p.ID}, nil
}
