package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/message"
	"goms/service/order/rpc/internal/svc"
	"goms/service/order/rpc/model"
	"goms/service/order/rpc/model/enum"
	"goms/service/order/rpc/pb/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
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
		Status:        enum.OrderStatusUnpaid,
		Consignee:     in.Consignee,
		Phone:         in.Phone,
		Address:       in.Address,
		TotalAmount:   totalAmount,
		TotalPrice:    totalPrice,
		OrderProducts: ops,
	}

	// 在Barrier中执行事务
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 这里用Gorm生成SQL
		mainSql := l.svcCtx.SqlDB.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Create(o)
		})
		result, err := tx.Exec(mainSql)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		orderId, err := result.LastInsertId()
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		for _, product := range o.OrderProducts {
			product.OrderID = orderId
		}
		subSql := l.svcCtx.SqlDB.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Create(o.OrderProducts)
		})
		if _, err = tx.Exec(subSql); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		// 添加订单超时延迟队列
		msg, err := json.Marshal(message.DqOrderIdMsg{
			OrderID: orderId,
		})
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		if _, err = l.svcCtx.Asynq.EnqueueContext(l.ctx, asynq.NewTask(message.DqOrderPaymentTimeout, msg), asynq.ProcessAt(time.Now().Add(time.Minute*15))); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		return nil
	}); err != nil {
		l.Logger.Error(errors.Wrap(err, "dtm tx execute failed"))
		return nil, err
	}

	return &order.Empty{}, nil
}
