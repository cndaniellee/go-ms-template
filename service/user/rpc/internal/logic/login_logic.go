package logic

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"goms/service/user/rpc/internal/svc"
	"goms/service/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.AuthReq) (*user.AuthReply, error) {

	// 获取用户
	u, err := l.svcCtx.UserModel.FindByUsername(l.ctx, in.Username)
	switch err {
	case gorm.ErrRecordNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case nil:
		break
	default:
		l.Logger.Error(errors.Wrap(err, "query user failed"))
		return nil, status.Error(codes.Aborted, err.Error())
	}

	// 校验密码
	if fmt.Sprintf("%x", md5.Sum([]byte(in.Password))) != u.Password {
		return nil, status.Error(codes.NotFound, "wrong password")
	}

	return &user.AuthReply{UserId: u.ID}, nil
}
