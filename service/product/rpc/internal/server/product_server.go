// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package server

import (
	"context"

	"goms/service/product/rpc/internal/logic"
	"goms/service/product/rpc/internal/svc"
	"goms/service/product/rpc/pb/product"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) List(ctx context.Context, in *product.ListReq) (*product.ListReply, error) {
	l := logic.NewListLogic(ctx, s.svcCtx)
	return l.List(in)
}

func (s *ProductServer) Detail(ctx context.Context, in *product.IdReq) (*product.DetailReply, error) {
	l := logic.NewDetailLogic(ctx, s.svcCtx)
	return l.Detail(in)
}

func (s *ProductServer) Edit(ctx context.Context, in *product.EditReq) (*product.IdReply, error) {
	l := logic.NewEditLogic(ctx, s.svcCtx)
	return l.Edit(in)
}

func (s *ProductServer) Remove(ctx context.Context, in *product.IdReq) (*product.Empty, error) {
	l := logic.NewRemoveLogic(ctx, s.svcCtx)
	return l.Remove(in)
}

// Internal
func (s *ProductServer) ListByIds(ctx context.Context, in *product.ListByIdsReq) (*product.ListByIdsReply, error) {
	l := logic.NewListByIdsLogic(ctx, s.svcCtx)
	return l.ListByIds(in)
}

// DTM
func (s *ProductServer) Deduct(ctx context.Context, in *product.DeductReq) (*product.Empty, error) {
	l := logic.NewDeductLogic(ctx, s.svcCtx)
	return l.Deduct(in)
}

func (s *ProductServer) DeductRollback(ctx context.Context, in *product.DeductReq) (*product.Empty, error) {
	l := logic.NewDeductRollbackLogic(ctx, s.svcCtx)
	return l.DeductRollback(in)
}
