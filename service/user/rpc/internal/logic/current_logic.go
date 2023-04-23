package logic

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"goms/service/user/rpc/internal/svc"
	"goms/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCurrentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentLogic {
	return &CurrentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CurrentLogic) Current(in *user.CurrentReq) (*user.CurrentReply, error) {

	// 获取用户
	u, err := l.svcCtx.UserModel.FindById(l.ctx, in.UserId)
	switch err {
	case gorm.ErrRecordNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case nil:
		break
	default:
		l.Error(errors.Wrap(err, "query user failed"))
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &user.CurrentReply{Username: u.Username}, nil
}
