// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: rio.proto

package rio_grpc

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

// RioClient is the client API for Rio service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RioClient interface {
	SetGPIObyOffset(ctx context.Context, in *GPIOselected, opts ...grpc.CallOption) (*ServerResponse, error)
	SetGPIObyAlias(ctx context.Context, in *GPIOselected, opts ...grpc.CallOption) (*ServerResponse, error)
}

type rioClient struct {
	cc grpc.ClientConnInterface
}

func NewRioClient(cc grpc.ClientConnInterface) RioClient {
	return &rioClient{cc}
}

func (c *rioClient) SetGPIObyOffset(ctx context.Context, in *GPIOselected, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/riogrpc.Rio/SetGPIObyOffset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rioClient) SetGPIObyAlias(ctx context.Context, in *GPIOselected, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/riogrpc.Rio/SetGPIObyAlias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RioServer is the server API for Rio service.
// All implementations must embed UnimplementedRioServer
// for forward compatibility
type RioServer interface {
	SetGPIObyOffset(context.Context, *GPIOselected) (*ServerResponse, error)
	SetGPIObyAlias(context.Context, *GPIOselected) (*ServerResponse, error)
	mustEmbedUnimplementedRioServer()
}

// UnimplementedRioServer must be embedded to have forward compatible implementations.
type UnimplementedRioServer struct {
}

func (UnimplementedRioServer) SetGPIObyOffset(context.Context, *GPIOselected) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGPIObyOffset not implemented")
}
func (UnimplementedRioServer) SetGPIObyAlias(context.Context, *GPIOselected) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGPIObyAlias not implemented")
}
func (UnimplementedRioServer) mustEmbedUnimplementedRioServer() {}

// UnsafeRioServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RioServer will
// result in compilation errors.
type UnsafeRioServer interface {
	mustEmbedUnimplementedRioServer()
}

func RegisterRioServer(s grpc.ServiceRegistrar, srv RioServer) {
	s.RegisterService(&Rio_ServiceDesc, srv)
}

func _Rio_SetGPIObyOffset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GPIOselected)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RioServer).SetGPIObyOffset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/riogrpc.Rio/SetGPIObyOffset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RioServer).SetGPIObyOffset(ctx, req.(*GPIOselected))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rio_SetGPIObyAlias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GPIOselected)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RioServer).SetGPIObyAlias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/riogrpc.Rio/SetGPIObyAlias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RioServer).SetGPIObyAlias(ctx, req.(*GPIOselected))
	}
	return interceptor(ctx, in, info, handler)
}

// Rio_ServiceDesc is the grpc.ServiceDesc for Rio service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rio_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "riogrpc.Rio",
	HandlerType: (*RioServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetGPIObyOffset",
			Handler:    _Rio_SetGPIObyOffset_Handler,
		},
		{
			MethodName: "SetGPIObyAlias",
			Handler:    _Rio_SetGPIObyAlias_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rio.proto",
}
