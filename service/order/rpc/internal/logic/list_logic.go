package logic

import (
	"context"
	"github.com/pkg/errors"
	"goms/service/order/rpc/model/enum"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goms/service/order/rpc/internal/svc"
	"goms/service/order/rpc/pb/order"

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

func (l *ListLogic) List(in *order.ListReq) (*order.ListReply, error) {

	// 获取订单列表
	orders, total, err := l.svcCtx.OrderModel.List(l.ctx, in.UserId, enum.OrderStatus(in.Status), int(in.Page), int(in.PageSize))
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "query orders failed"))
		return nil, status.Error(codes.Aborted, err.Error())
	}

	// 转换数据
	list := make([]*order.ListItem, len(orders))
	for i, item := range orders {
		products := make([]*order.Product, len(item.OrderProducts))
		for j, product := range item.OrderProducts {
			products[j] = &order.Product{
				Id:     product.ProductID,
				Amount: product.Amount,
				Price:  product.Price,
			}
		}
		list[i] = &order.ListItem{
			Id:          item.ID,
			Status:      int32(item.Status),
			Products:    products,
			TotalAmount: item.TotalAmount,
			TotalPrice:  item.TotalPrice,
		}
	}

	return &order.ListReply{List: list, Page: &order.Page{Page: in.Page, PageSize: in.PageSize, Total: total}}, nil
}
