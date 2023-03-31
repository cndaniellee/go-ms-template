package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"goms/common/model"
	"goms/service/product/rpc/internal/svc"
	"goms/service/product/rpc/pb/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductLogic {
	return &DeductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DTM
func (l *DeductLogic) Deduct(in *product.DeductReq) (*product.Empty, error) {

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

	// 使用Redis分布式锁，不锁数据库
	lock := redis.NewRedisLock(l.svcCtx.Redis, fmt.Sprintf(model.IdLockKey, l.svcCtx.ProductModel.Name(), in.Id))
	if _, err = lock.AcquireCtx(l.ctx); err != nil {
		l.Logger.Error(errors.Wrap(err, "acquire redis lock failed"))
		return nil, status.Error(codes.Internal, err.Error())
	}
	// 解锁
	defer func(lock *redis.RedisLock, ctx context.Context) {
		if _, err = lock.ReleaseCtx(ctx); err != nil {
			l.Logger.Error(errors.Wrap(err, "release redis lock failed"))
		}
	}(lock, l.ctx)

	// 在Barrier中执行事务
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 扣减库存
		result, err := tx.Exec("UPDATE `product` SET `stock` = `stock` - ? WHERE `id` = ? AND `stock` >= ? AND `deleted_at` IS NULL", in.Amount, in.Id, in.Amount)
		if err != nil {
			l.Logger.Error(errors.Wrap(err, "stock deduct failed"))
			return status.Error(codes.Internal, err.Error())
		}
		affected, err := result.RowsAffected()
		if err != nil {
			l.Logger.Error(errors.Wrap(err, "get rows affected failed"))
			return status.Error(codes.Internal, err.Error())
		}
		// 影响行数0代表扣减失败，可能是库存不足或产品无效，返回事务失败
		if affected <= 0 {
			l.Logger.Error("stock not enough, terminate dtm tx")
			return status.Error(codes.Aborted, dtmcli.ResultFailure)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &product.Empty{}, nil
}
