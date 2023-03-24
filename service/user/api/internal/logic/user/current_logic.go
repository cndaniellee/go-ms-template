package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/request"
	"goms/common/response"
	"goms/common/response/errcode/usercode"
	"goms/service/user/api/internal/svc"
	"goms/service/user/api/internal/types"
	"goms/service/user/rpc/userclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		l.Logger.Error(errors.Wrap(err, "user id parse failed"))
		err = response.ErrResp(0, usercode.Register, response.InternalError, err.Error())
		return
	}

	// 调用RPC服务
	reply, err := l.svcCtx.UserRpc.Current(l.ctx, &userclient.CurrentReq{
		UserId: userId,
	})
	if err != nil {
		switch s, _ := status.FromError(err); s.Code() {
		case codes.NotFound:
			err = response.ErrResp(1, usercode.Current, response.NoneMatching, s.Message())
		case codes.Aborted:
			err = response.ErrResp(2, usercode.Current, response.InternalError, s.Message())
		default:
			l.Logger.Error(errors.Wrap(err, "user rpc call failed"))
			err = response.ErrResp(3, usercode.Current, response.ServiceError, s.Message())
		}
		return
	}

	resp = &types.CurrentResp{Username: reply.Username}

	return
}
