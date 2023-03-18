package auth

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/app/user/api/internal/svc"
	"goms/app/user/api/internal/types"
	"goms/app/user/rpc/userclient"
	"goms/common/auth"
	"goms/common/logtool"
	"goms/common/reply"
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

	// 调用RPC服务
	rpcResp, rpcErr := l.svcCtx.UserRpc.AuthLogin(l.ctx, &userclient.AuthReq{
		Username: req.Username,
		Password: req.Password,
	})
	if rpcErr != nil {
		if code, msg := logtool.CheckRpcConnErr(l.Logger, rpcErr); code == reply.NoneMatching {
			err = response.ErrResp(0, errcode.Login, response.NoneMatching, msg)
		} else {
			err = response.ErrResp(1, errcode.Login, response.ServiceError)
		}
		return
	}

	// 生成用户Token
	token, err := auth.GenerateUserToken(l.svcCtx.Config.JwtAuth, rpcResp.UserId)
	if err != nil {
		l.Logger.Error(errors.Wrapf(err, "generate token failed"))
		err = response.ErrResp(2, errcode.Login, response.InternalError)
		return nil, err
	}

	resp = &types.AuthResp{Token: token}

	return
}
