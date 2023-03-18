package logic

import (
	"context"
	"github.com/pkg/errors"
	"goms/app/user/rpc/internal/svc"
	"goms/app/user/rpc/model"
	"goms/app/user/rpc/pb/user"
	"goms/common/reply"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthRegisterLogic {
	return &AuthRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthRegisterLogic) AuthRegister(in *user.AuthReq) (*user.AuthReply, error) {

	// 检查用户是否已存在
	if l.svcCtx.SqlDb.Where("username = ?", in.Username).First(&model.User{}).Error != gorm.ErrRecordNotFound {
		return nil, status.Error(reply.DataConflict, "username exists")
	}

	// 创建用户
	u := &model.User{Username: in.Username, Password: in.Password}
	if err := l.svcCtx.SqlDb.Create(u).Error; err != nil {
		l.Logger.Error(errors.Wrapf(err, "create user failed"))
		return nil, status.Error(reply.ServiceError, err.Error())
	}

	return &user.AuthReply{UserId: u.ID}, nil
}
