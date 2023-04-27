package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"goms/common/lock"
	"goms/common/model"
	"goms/service/order/rpc/model/enum"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"goms/service/order/rpc/internal/svc"
	"goms/service/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckPaymentTimeoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckPaymentTimeoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckPaymentTimeoutLogic {
	return &CheckPaymentTimeoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Internal
func (l *CheckPaymentTimeoutLogic) CheckPaymentTimeout(in *order.IdReq) (*order.Empty, error) {

	// 使用Redis分布式锁，不锁数据库
	rl := lock.NewRedisLock(l.svcCtx.Redis, fmt.Sprintf(model.IdLockKey, l.svcCtx.OrderModel.Name(), in.Id), 5)
	if err := rl.AcquireExCtx(l.ctx); err != nil {
		l.Error(errors.Wrap(err, "acquire redis lock failed"))
		return nil, status.Error(codes.Aborted, err.Error())
	}
	// 解锁
	defer func(rl *lock.Lock, ctx context.Context) {
		if err := rl.ReleaseExCtx(ctx); err != nil {
			l.Error(errors.Wrap(err, "release redis lock failed"))
		}
	}(rl, l.ctx)

	// 查询订单数据
	o, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		case nil:
			break
		default:
			return nil, status.Error(codes.Aborted, err.Error())
		}
	}

	// 如果订单未支付，则关闭订单
	if o.Status == enum.OrderStatusUnpaid {
		if err = l.svcCtx.OrderModel.UpdateStatus(l.ctx, o.ID, enum.OrderStatusClose); err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
	}

	return &order.Empty{}, nil
}
