package logic

import (
	"context"
	"github.com/pkg/errors"
	"goms/app/order/rpc/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goms/app/order/rpc/internal/svc"
	"goms/app/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitLogic {
	return &SubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubmitLogic) Submit(in *order.SubmitReq) (*order.Empty, error) {

	// DTM

	// 转换数据
	var totalAmount, totalPrice int64
	ops := make([]*model.OrderProduct, len(in.Products))
	for i, item := range in.Products {
		ops[i] = &model.OrderProduct{
			ProductID: item.Id,
			Amount:    item.Amount,
			Price:     item.Price,
		}
		totalAmount += item.Amount
		totalPrice += item.Amount * item.Price
	}

	o := &model.Order{
		Status:        model.OrderStatusUnpaid,
		Consignee:     in.Consignee,
		Phone:         in.Phone,
		Address:       in.Address,
		TotalAmount:   totalAmount,
		TotalPrice:    totalPrice,
		OrderProducts: ops,
	}

	// 创建订单
	if err := l.svcCtx.OrderModel.Create(l.ctx, o); err != nil {
		l.Logger.Error(errors.Wrap(err, "create order failed"))
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &order.Empty{}, nil
}
