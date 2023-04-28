package token

import (
	"context"
	"github.com/pkg/errors"
	"goms/common/request"
	"goms/common/response/errcode/authcode"

	"goms/service/auth/api/internal/svc"
	"goms/service/auth/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/response"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh() (resp *types.AuthResp, err error) {

	// 解析用户ID
	userId, err := request.ParseUserId(l.ctx)
	if err != nil {
		l.Error(errors.Wrap(err, "user id parse failed"))
		err = response.ErrResp(1, authcode.Refresh, response.InternalError, err.Error())
		return
	}

	// 生成用户Token
	userToken, refreshToken, err := l.svcCtx.TokenGenerator.ExecuteUser(l.ctx, userId)
	if err != nil {
		l.Error(errors.Wrap(err, "generate token failed"))
		err = response.ErrResp(2, authcode.Refresh, response.InternalError, err.Error())
		return
	}

	resp = &types.AuthResp{
		Token:        userToken,
		RefreshToken: refreshToken,
	}

	return
}
