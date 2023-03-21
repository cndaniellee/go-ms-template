package logic

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/pkg/errors"
	"goms/app/user/rpc/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"goms/app/user/rpc/internal/svc"
	"goms/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.AuthReq) (*user.AuthReply, error) {

	// 检查用户是否已存在
	_, err := l.svcCtx.UserModel.FindByUsername(l.ctx, in.Username)
	switch err {
	case nil:
		return nil, status.Error(codes.AlreadyExists, "username exists")
	case gorm.ErrRecordNotFound:
		break
	default:
		l.Logger.Error(errors.Wrap(err, "query user failed"))
		return nil, status.Error(codes.Aborted, err.Error())
	}

	// 创建用户
	u := &model.User{Username: in.Username, Password: fmt.Sprintf("%x", md5.Sum([]byte(in.Password)))}
	if err = l.svcCtx.UserModel.Create(l.ctx, u); err != nil {
		l.Logger.Error(errors.Wrap(err, "create user failed"))
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &user.AuthReply{UserId: u.ID}, nil
}
