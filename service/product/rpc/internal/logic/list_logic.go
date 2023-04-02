package logic

import (
	"context"
	"github.com/pkg/errors"
	"goms/service/product/rpc/model/enum"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goms/service/product/rpc/internal/svc"
	"goms/service/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *product.ListReq) (*product.ListReply, error) {

	// ES：获取产品列表
	products, total, err := l.svcCtx.ProductES.Search(l.ctx, in.Search, enum.ProductCategory(in.Category), int(in.Page), int(in.PageSize))
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "query products failed"))
		return nil, status.Error(codes.Aborted, err.Error())
	}

	// 获取产品列表
	//products, total, err := l.svcCtx.ProductModel.List(l.ctx, in.Search, enum.ProductCategory(in.Category), int(in.Page), int(in.PageSize))
	//if err != nil {
	//	l.Logger.Error(errors.Wrap(err, "query products failed"))
	//	return nil, status.Error(codes.Aborted, err.Error())
	//}

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

	return &product.ListReply{List: list, Page: &product.Page{Page: in.Page, PageSize: in.PageSize, Total: total}}, nil
}
