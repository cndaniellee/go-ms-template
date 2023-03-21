// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"goms/app/user/rpc/internal/logic"
	"goms/app/user/rpc/internal/svc"
	"goms/app/user/rpc/pb/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *user.AuthReq) (*user.AuthReply, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Register(ctx context.Context, in *user.AuthReq) (*user.AuthReply, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Current(ctx context.Context, in *user.CurrentReq) (*user.CurrentReply, error) {
	l := logic.NewCurrentLogic(ctx, s.svcCtx)
	return l.Current(in)
}
