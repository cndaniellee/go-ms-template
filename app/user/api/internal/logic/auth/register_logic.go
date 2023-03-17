package auth

import (
	"context"
	"github.com/pkg/errors"
	"goms/app/user/rpc/userclient"
	"goms/common/auth"
	"goms/common/logtool"
	"goms/common/reply"

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

	// 调用RPC服务
	rpcResp, rpcErr := l.svcCtx.UserRpc.AuthRegister(l.ctx, &userclient.AuthReq{
		Username: req.Username,
		Password: req.Password,
	})
	if rpcErr != nil {
		if msg, note := logtool.CheckRpcConnErr(l.Logger, rpcErr); msg == reply.NoneMatching {
			err = response.ErrResp(0, errcode.Register, response.NoneMatching, note)
		} else {
			err = response.ErrResp(1, errcode.Register, response.ServiceError)
		}
		return
	}

	// 生成用户Token
	token, err := auth.GenerateUserToken(l.svcCtx.Config.JwtAuth, rpcResp.UserId)
	if err != nil {
		l.Logger.Error(errors.Wrapf(err, "generate token error"))
		err = response.ErrResp(2, errcode.Register, response.InternalError)
		return
	}

	resp = &types.AuthResp{Token: token}

	return
}
