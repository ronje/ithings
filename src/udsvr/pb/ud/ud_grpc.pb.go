// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: ud.proto

package ud

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
	Ud_Ping_FullMethodName = "/ud.Ud/Ping"
)

// UdClient is the client API for Ud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UdClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type udClient struct {
	cc grpc.ClientConnInterface
}

func NewUdClient(cc grpc.ClientConnInterface) UdClient {
	return &udClient{cc}
}

func (c *udClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Ud_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UdServer is the server API for Ud service.
// All implementations must embed UnimplementedUdServer
// for forward compatibility
type UdServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedUdServer()
}

// UnimplementedUdServer must be embedded to have forward compatible implementations.
type UnimplementedUdServer struct {
}

func (UnimplementedUdServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedUdServer) mustEmbedUnimplementedUdServer() {}

// UnsafeUdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UdServer will
// result in compilation errors.
type UnsafeUdServer interface {
	mustEmbedUnimplementedUdServer()
}

func RegisterUdServer(s grpc.ServiceRegistrar, srv UdServer) {
	s.RegisterService(&Ud_ServiceDesc, srv)
}

func _Ud_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UdServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ud_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UdServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Ud_ServiceDesc is the grpc.ServiceDesc for Ud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ud.Ud",
	HandlerType: (*UdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Ud_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ud.proto",
}