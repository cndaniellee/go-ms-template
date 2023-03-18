package logic

import (
	"context"
	"goms/app/user/rpc/internal/svc"
	"goms/app/user/rpc/model"
	"goms/app/user/rpc/pb/user"
	"goms/common/reply"
	"google.golang.org/grpc/status"

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

func (l *UserCurrentLogic) UserCurrent(in *user.UserCurrentReq) (*user.UserCurrentReply, error) {

	// 查找用户
	u := &model.User{}
	if err := l.svcCtx.SqlDb.Where("id = ?", in.UserId).First(u).Error; err != nil {
		return nil, status.Error(reply.NoneMatching, err.Error())
	}

	return &user.UserCurrentReply{Username: u.Username}, nil
}
