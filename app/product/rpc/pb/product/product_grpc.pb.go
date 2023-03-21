// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: doc/goctl/rpc/product.proto

package product

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Product_List_FullMethodName   = "/product.product/List"
	Product_Detail_FullMethodName = "/product.product/Detail"
	Product_Edit_FullMethodName   = "/product.product/Edit"
	Product_Remove_FullMethodName = "/product.product/Remove"
)

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListReply, error)
	Detail(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*DetailReply, error)
	Edit(ctx context.Context, in *EditReq, opts ...grpc.CallOption) (*IdReply, error)
	Remove(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Empty, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, Product_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) Detail(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*DetailReply, error) {
	out := new(DetailReply)
	err := c.cc.Invoke(ctx, Product_Detail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) Edit(ctx context.Context, in *EditReq, opts ...grpc.CallOption) (*IdReply, error) {
	out := new(IdReply)
	err := c.cc.Invoke(ctx, Product_Edit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) Remove(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Product_Remove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	List(context.Context, *ListReq) (*ListReply, error)
	Detail(context.Context, *IdReq) (*DetailReply, error)
	Edit(context.Context, *EditReq) (*IdReply, error)
	Remove(context.Context, *IdReq) (*Empty, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) List(context.Context, *ListReq) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProductServer) Detail(context.Context, *IdReq) (*DetailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (UnimplementedProductServer) Edit(context.Context, *EditReq) (*IdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedProductServer) Remove(context.Context, *IdReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).List(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_Detail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Detail(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_Edit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Edit(ctx, req.(*EditReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_Remove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Remove(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Product_List_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _Product_Detail_Handler,
		},
		{
			MethodName: "Edit",
			Handler:    _Product_Edit_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Product_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doc/goctl/rpc/product.proto",
}
