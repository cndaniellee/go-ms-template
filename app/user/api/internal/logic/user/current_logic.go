package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/app/user/api/internal/svc"
	"goms/app/user/api/internal/types"
	"goms/app/user/rpc/userclient"
	"goms/common/logtool"
	"goms/common/request"
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

	// 解析用户ID
	userId, err := request.ParseUserId(l.ctx)
	if err != nil {
		l.Logger.Error(errors.Wrapf(err, "user id parse failed"))
		err = response.ErrResp(0, errcode.Register, response.InternalError)
		return
	}

	// 调用RPC服务
	rpcResp, rpcErr := l.svcCtx.UserRpc.UserCurrent(l.ctx, &userclient.UserCurrentReq{
		UserId: userId,
	})
	if rpcErr != nil {
		logtool.CheckRpcConnErr(l.Logger, rpcErr)
		err = response.ErrResp(1, errcode.Current, response.ServiceError)
		return
	}

	resp = &types.CurrentResp{Username: rpcResp.Username}

	return
}
