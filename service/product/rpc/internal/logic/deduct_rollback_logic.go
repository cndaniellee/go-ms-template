package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/pkg/errors"
	"goms/common/lock"
	"goms/common/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goms/service/product/rpc/internal/svc"
	"goms/service/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductRollbackLogic {
	return &DeductRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DTM
func (l *DeductRollbackLogic) DeductRollback(in *product.DeductReq) (*product.Empty, error) {

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		l.Error(errors.Wrap(err, "get dtm barrier failed"))
		return nil, status.Error(codes.Internal, err.Error())
	}
	db, err := l.svcCtx.SqlDB.DB()
	if err != nil {
		l.Error(errors.Wrap(err, "get db instance failed"))
		return nil, status.Error(codes.Internal, err.Error())
	}

	// 使用Redis分布式锁，不锁数据库
	rl := lock.NewRedisLock(l.svcCtx.Redis, fmt.Sprintf(model.IdLockKey, l.svcCtx.ProductModel.Name(), in.Id), 5)
	if err = rl.AcquireExCtx(l.ctx); err != nil {
		l.Error(errors.Wrap(err, "acquire redis lock failed"))
		return nil, status.Error(codes.Aborted, err.Error())
	}
	// 解锁
	defer func(rl *lock.Lock, ctx context.Context) {
		if err = rl.ReleaseExCtx(ctx); err != nil {
			l.Error(errors.Wrap(err, "release redis lock failed"))
		}
	}(rl, l.ctx)

	// 在Barrier中执行事务
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 恢复库存
		result, err := tx.Exec("UPDATE `product` SET `stock` = `stock` + ? WHERE `id` = ? AND `deleted_at` IS NULL", in.Amount, in.Id)
		if err != nil {
			l.Error(errors.Wrap(err, "stock rollback failed"))
			return status.Error(codes.Internal, err.Error())
		}
		affected, err := result.RowsAffected()
		if err != nil {
			l.Error(errors.Wrap(err, "get rows affected failed"))
			return status.Error(codes.Internal, err.Error())
		}
		// 影响行数0代表回滚失败，可能是产品无效，返回事务失败
		if affected <= 0 {
			l.Error("wrong stock number, terminate dtm tx")
			return status.Error(codes.Aborted, dtmcli.ResultFailure)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &product.Empty{}, nil
}
