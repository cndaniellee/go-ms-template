// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package orderclient

import (
	"context"

	"goms/app/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DetailReply = order.DetailReply
	Empty       = order.Empty
	IdReq       = order.IdReq
	ListItem    = order.ListItem
	ListReply   = order.ListReply
	ListReq     = order.ListReq
	Page        = order.Page
	Product     = order.Product
	SubmitReq   = order.SubmitReq

	Order interface {
		Submit(ctx context.Context, in *SubmitReq, opts ...grpc.CallOption) (*Empty, error)
		List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListReply, error)
		Detail(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*DetailReply, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) Submit(ctx context.Context, in *SubmitReq, opts ...grpc.CallOption) (*Empty, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.Submit(ctx, in, opts...)
}

func (m *defaultOrder) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListReply, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.List(ctx, in, opts...)
}

func (m *defaultOrder) Detail(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*DetailReply, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.Detail(ctx, in, opts...)
}
