package logic

import (
	"context"
	"strconv"

	"goms/app/user/rpc/internal/svc"
	"goms/app/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCurrentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCurrentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCurrentLogic {
	return &UserCurrentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCurrentLogic) UserCurrent(in *user.UserCurrentReq) (*user.UserCurrentResp, error) {

	return &user.UserCurrentResp{Username: "User_" + strconv.Itoa(int(in.UserId))}, nil
}
