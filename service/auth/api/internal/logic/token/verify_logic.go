package token

import (
	"context"
	"goms/common/middleware"
	"goms/common/request"

	"github.com/zeromicro/go-zero/core/logx"
	"goms/service/auth/api/internal/svc"
)

type VerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyLogic) Verify() (resp *middleware.AuthRequestUserInfo, err error) {

	// 解析用户ID
	userId, err := request.ParseUserId(l.ctx)
	if err != nil {
		return nil, err
	}

	resp = &middleware.AuthRequestUserInfo{
		UserId: userId,
	}

	return
}
