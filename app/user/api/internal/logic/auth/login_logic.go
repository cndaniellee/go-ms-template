package auth

import (
	"context"
	"goms/app/user/rpc/userclient"

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

	rpcResp, rpcErr := l.svcCtx.UserRpc.AuthLogin(l.ctx, &userclient.AuthReq{
		Username: req.Username,
		Password: req.Password,
	})
	if rpcErr != nil {
		l.Logger.Errorf("rpc call error: %v", rpcErr.Error())
		err = response.ErrResp(0, errcode.Login, response.RpcCallError)
		return
	}

	resp = &types.AuthResp{Token: rpcResp.Token}

	return
}
