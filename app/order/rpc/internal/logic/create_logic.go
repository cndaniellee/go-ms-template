package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/pkg/errors"
	"goms/app/order/rpc/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"goms/app/order/rpc/internal/svc"
	"goms/app/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DTM
func (l *CreateLogic) Create(in *order.CreateReq) (*order.Empty, error) {

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "get dtm barrier failed"))
		return nil, status.Error(codes.Internal, err.Error())
	}
	db, err := l.svcCtx.SqlDB.DB()
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "get db instance failed"))
		return nil, status.Error(codes.Internal, err.Error())
	}

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
		UserID:        in.UserId,
		RefNo:         in.RefNo,
		Status:        model.OrderStatusUnpaid,
		Consignee:     in.Consignee,
		Phone:         in.Phone,
		Address:       in.Address,
		TotalAmount:   totalAmount,
		TotalPrice:    totalPrice,
		OrderProducts: ops,
	}

	// 在Barrier中执行事务
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		query := l.svcCtx.SqlDB.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Create(o)
		})
		result, err := tx.Exec(query)
		if err != nil {
			return err
		}
		// 清除缓存
		id, _ := result.LastInsertId()
		l.svcCtx.OrderModel.RemoveCache(l.ctx, id)
		return nil
	}); err != nil {
		l.Logger.Error(errors.Wrap(err, "dtm tx execute failed"))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &order.Empty{}, nil
}
