package logic

import (
	"context"
	"github.com/CNDanielLee/go-utils"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/app/user/rpc/internal/svc"
	"goms/app/user/rpc/model"
	"goms/app/user/rpc/pb/user"
	"goms/common/reply"
	"google.golang.org/grpc/status"
)

type AuthLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLoginLogic {
	return &AuthLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthLoginLogic) AuthLogin(in *user.AuthReq) (*user.AuthReply, error) {

	// 校验用户
	u := &model.User{}
	if err := l.svcCtx.SqlDb.Where("username = ?", in.Username).First(u).Error; err != nil {
		return nil, status.Error(reply.NoneMatching, err.Error())
	}
	if utils.EncryptStrToMd5(in.Password) != u.Password {
		return nil, status.Error(reply.NoneMatching, "wrong password")
	}

	return &user.AuthReply{UserId: u.ID}, nil
}
