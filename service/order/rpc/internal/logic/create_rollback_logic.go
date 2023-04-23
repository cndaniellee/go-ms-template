package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/pkg/errors"
	"goms/service/order/rpc/internal/svc"
	"goms/service/order/rpc/pb/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRollbackLogic {
	return &CreateRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DTM
func (l *CreateRollbackLogic) CreateRollback(in *order.CreateReq) (*order.Empty, error) {

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

	// 在Barrier中执行事务
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		_, err = tx.Exec("DELETE FROM `product` WHERE `ref_no` = ?", in.RefNo)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		return nil
	}); err != nil {
		l.Error(errors.Wrap(err, "dtm tx execute failed"))
		return nil, err
	}

	return &order.Empty{}, nil
}
