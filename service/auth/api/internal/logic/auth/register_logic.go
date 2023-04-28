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
	reply, err := l.svcCtx.UserRpc.Register(l.ctx, &userclient.AuthReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.AlreadyExists:
				err = response.ErrResp(1, authcode.Register, response.AlreadyExists, s.Message())
			case codes.Aborted:
				err = response.ErrResp(2, authcode.Register, response.InternalError, s.Message())
			default:
				l.Error(errors.Wrap(err, "user rpc call failed"))
				err = response.ErrResp(3, authcode.Register, response.ServiceError, s.Message())
			}
			return
		}
		return
	}

	// 生成用户Token
	userToken, refreshToken, err := l.svcCtx.TokenGenerator.ExecuteUser(l.ctx, reply.UserId)
	if err != nil {
		l.Error(errors.Wrap(err, "generate token failed"))
		err = response.ErrResp(4, authcode.Register, response.InternalError, err.Error())
		return
	}

	resp = &types.AuthResp{
		Token:        userToken,
		RefreshToken: refreshToken,
	}

	return
}
