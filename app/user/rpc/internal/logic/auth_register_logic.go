package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/hash"

	"goms/app/user/rpc/internal/svc"
	"goms/app/user/rpc/pb/user"

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

func (l *AuthRegisterLogic) AuthRegister(in *user.AuthReq) (*user.AuthResp, error) {

	return &user.AuthResp{Token: string(hash.Md5([]byte(in.Username)))}, nil
}
