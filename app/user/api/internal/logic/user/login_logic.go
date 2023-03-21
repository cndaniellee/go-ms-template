package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/app/user/api/internal/svc"
	"goms/app/user/api/internal/types"
	"goms/app/user/rpc/userclient"
	"goms/common/auth"
	"goms/common/response"
	"goms/common/response/errcode/usercode"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
			err = response.ErrResp(0, usercode.Login, response.NoneMatching, s.Message())
		case codes.Aborted:
			err = response.ErrResp(1, usercode.Login, response.InternalError, s.Message())
		default:
			l.Logger.Error(errors.Wrap(err, "user rpc call failed"))
			err = response.ErrResp(2, usercode.Login, response.ServiceError, s.Message())
		}
		return
	}
	// 生成用户Token
	token, err := auth.GenerateUserToken(l.svcCtx.Config.JwtAuth, reply.UserId)
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "generate token failed"))
		err = response.ErrResp(3, usercode.Login, response.InternalError, err.Error())
		return nil, err
	}

	resp = &types.AuthResp{Token: token}

	return
}
