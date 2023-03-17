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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.AuthReq) (resp *types.AuthResp, err error) {

	rpcResp, rpcErr := l.svcCtx.UserRpc.AuthRegister(l.ctx, &userclient.AuthReq{
		Username: req.Username,
		Password: req.Password,
	})
	if rpcErr != nil {
		l.Logger.Errorf("rpc call error: %v", rpcErr.Error())
		err = response.ErrResp(0, errcode.Register, response.RpcCallError)
		return
	}
	resp.Token = rpcResp.Token

	return
}
