package auth

import (
	"context"
	"github.com/pkg/errors"
	"goms/service/user/rpc/userclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goms/service/auth/api/internal/svc"
	"goms/service/auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/response"
	"goms/common/response/errcode/authcode"
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
	reply, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.AuthReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		switch s, _ := status.FromError(err); s.Code() {
		case codes.NotFound:
			err = response.ErrResp(1, authcode.Login, response.NoneMatching, s.Message())
		case codes.Aborted:
			err = response.ErrResp(2, authcode.Login, response.InternalError, s.Message())
		default:
			l.Error(errors.Wrap(err, "user rpc call failed"))
			err = response.ErrResp(3, authcode.Login, response.ServiceError, s.Message())
		}
		return
	}
	// 生成用户Token
	userToken, refreshToken, err := l.svcCtx.TokenGenerator.ExecuteUser(l.ctx, reply.UserId)
	if err != nil {
		l.Error(errors.Wrap(err, "generate token failed"))
		err = response.ErrResp(4, authcode.Login, response.InternalError, err.Error())
		return nil, err
	}

	resp = &types.AuthResp{
		Token:        userToken,
		RefreshToken: refreshToken,
	}

	return
}
