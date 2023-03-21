package logic

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goms/app/product/rpc/internal/svc"
	"goms/app/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByIdsLogic {
	return &ListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Internal
func (l *ListByIdsLogic) ListByIds(in *product.ListByIdsReq) (*product.ListByIdsReply, error) {

	// 获取产品列表
	products, err := l.svcCtx.ProductModel.ListByIds(l.ctx, in.Ids)
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "query products failed"))
		return nil, status.Error(codes.Aborted, err.Error())
	}

	// 转换数据
	list := make([]*product.ListItem, len(products))
	for i, item := range products {
		list[i] = &product.ListItem{
			Id:       item.ID,
			Title:    item.Title,
			Category: int32(item.Category),
			Stock:    item.Stock,
			Price:    item.Price,
		}
	}

	return &product.ListByIdsReply{List: list}, nil
}
