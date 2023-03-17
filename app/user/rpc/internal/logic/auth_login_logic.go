package logic

import (
	"context"
	"goms/app/user/rpc/internal/svc"
	"goms/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
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

	return &user.AuthReply{UserId: 1}, nil
}
