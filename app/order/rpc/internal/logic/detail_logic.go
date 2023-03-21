package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"goms/app/order/rpc/internal/svc"
	"goms/app/order/rpc/pb/order"

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

func (l *DetailLogic) Detail(in *order.IdReq) (*order.DetailReply, error) {

	// 获取订单
	o, err := l.svcCtx.OrderModel.FindById(l.ctx, in.Id)
	switch err {
	case gorm.ErrRecordNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case nil:
		break
	default:
		return nil, status.Error(codes.Aborted, err.Error())
	}

	// 转换数据
	products := make([]*order.Product, len(o.OrderProducts))
	for i, product := range o.OrderProducts {
		products[i] = &order.Product{
			Id:     product.ProductID,
			Amount: product.Amount,
			Price:  product.Price,
		}·
	}

	return &order.DetailReply{
		Id:          o.ID,
		Status:      int32(o.Status),
		Products:    products,
		Consignee:   o.Consignee,
		Phone:       o.Phone,
		Address:     o.Address,
		TotalAmount: o.TotalAmount,
		TotalPrice:  o.TotalPrice,
		CreatedAt:   o.CreatedAt.UnixMilli(),
	}, nil
}
