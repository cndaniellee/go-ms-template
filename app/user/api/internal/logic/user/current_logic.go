package user

import (
	"context"
	"goms/app/user/rpc/userclient"

	"goms/app/user/api/internal/svc"
	"goms/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/response"
	"goms/common/response/errcode"
)

type CurrentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentLogic {
	return &CurrentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurrentLogic) Current() (resp *types.CurrentResp, err error) {

	rpcResp, rpcErr := l.svcCtx.UserRpc.UserCurrent(l.ctx, &userclient.UserCurrentReq{
		UserId: 1,
	})
	if rpcErr != nil {
		l.Logger.Errorf("rpc call error: %v", rpcErr.Error())
		err = response.ErrResp(0, errcode.Current, response.RpcCallError)
		return
	}
	resp.Username = rpcResp.Username

	return
}
