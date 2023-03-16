package auth

import (
	"context"

	"goms/app/user/api/internal/svc"
	"goms/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/response"
	"goms/common/response/errcode"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.AuthReq) (resp *types.AuthResp, err error) {
	// todo: add your logic here and delete this line

	err = response.ErrResp(0, errcode.Login, "example")

	return
}
