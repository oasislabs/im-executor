// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: gateway_min.proto

package types

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

// WebClient is the client API for Web service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebClient interface {
	InitWithdraw(ctx context.Context, in *InitWithdrawRequest, opts ...grpc.CallOption) (*InitWithdrawResponse, error)
	InitPegRefund(ctx context.Context, in *InitPegRefundRequest, opts ...grpc.CallOption) (*InitPegRefundResponse, error)
}

type webClient struct {
	cc grpc.ClientConnInterface
}

func NewWebClient(cc grpc.ClientConnInterface) WebClient {
	return &webClient{cc}
}

func (c *webClient) InitWithdraw(ctx context.Context, in *InitWithdrawRequest, opts ...grpc.CallOption) (*InitWithdrawResponse, error) {
	out := new(InitWithdrawResponse)
	err := c.cc.Invoke(ctx, "/sgn.gateway.v1.Web/InitWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webClient) InitPegRefund(ctx context.Context, in *InitPegRefundRequest, opts ...grpc.CallOption) (*InitPegRefundResponse, error) {
	out := new(InitPegRefundResponse)
	err := c.cc.Invoke(ctx, "/sgn.gateway.v1.Web/InitPegRefund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebServer is the server API for Web service.
// All implementations must embed UnimplementedWebServer
// for forward compatibility
type WebServer interface {
	InitWithdraw(context.Context, *InitWithdrawRequest) (*InitWithdrawResponse, error)
	InitPegRefund(context.Context, *InitPegRefundRequest) (*InitPegRefundResponse, error)
	mustEmbedUnimplementedWebServer()
}

// UnimplementedWebServer must be embedded to have forward compatible implementations.
type UnimplementedWebServer struct {
}

func (UnimplementedWebServer) InitWithdraw(context.Context, *InitWithdrawRequest) (*InitWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitWithdraw not implemented")
}
func (UnimplementedWebServer) InitPegRefund(context.Context, *InitPegRefundRequest) (*InitPegRefundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitPegRefund not implemented")
}
func (UnimplementedWebServer) mustEmbedUnimplementedWebServer() {}

// UnsafeWebServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebServer will
// result in compilation errors.
type UnsafeWebServer interface {
	mustEmbedUnimplementedWebServer()
}

func RegisterWebServer(s grpc.ServiceRegistrar, srv WebServer) {
	s.RegisterService(&Web_ServiceDesc, srv)
}

func _Web_InitWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).InitWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgn.gateway.v1.Web/InitWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).InitWithdraw(ctx, req.(*InitWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Web_InitPegRefund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitPegRefundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServer).InitPegRefund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgn.gateway.v1.Web/InitPegRefund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServer).InitPegRefund(ctx, req.(*InitPegRefundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Web_ServiceDesc is the grpc.ServiceDesc for Web service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Web_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sgn.gateway.v1.Web",
	HandlerType: (*WebServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitWithdraw",
			Handler:    _Web_InitWithdraw_Handler,
		},
		{
			MethodName: "InitPegRefund",
			Handler:    _Web_InitPegRefund_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway_min.proto",
}
