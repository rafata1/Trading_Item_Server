// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protopack

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

// AccountServicesClient is the client API for AccountServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServicesClient interface {
	SignupService(ctx context.Context, in *SignupReq, opts ...grpc.CallOption) (*SignupRes, error)
	LoginService(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
}

type accountServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServicesClient(cc grpc.ClientConnInterface) AccountServicesClient {
	return &accountServicesClient{cc}
}

func (c *accountServicesClient) SignupService(ctx context.Context, in *SignupReq, opts ...grpc.CallOption) (*SignupRes, error) {
	out := new(SignupRes)
	err := c.cc.Invoke(ctx, "/protopack.AccountServices/SignupService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServicesClient) LoginService(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, "/protopack.AccountServices/LoginService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServicesServer is the server API for AccountServices service.
// All implementations must embed UnimplementedAccountServicesServer
// for forward compatibility
type AccountServicesServer interface {
	SignupService(context.Context, *SignupReq) (*SignupRes, error)
	LoginService(context.Context, *LoginReq) (*LoginRes, error)
	mustEmbedUnimplementedAccountServicesServer()
}

// UnimplementedAccountServicesServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServicesServer struct {
}

func (UnimplementedAccountServicesServer) SignupService(context.Context, *SignupReq) (*SignupRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignupService not implemented")
}
func (UnimplementedAccountServicesServer) LoginService(context.Context, *LoginReq) (*LoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginService not implemented")
}
func (UnimplementedAccountServicesServer) mustEmbedUnimplementedAccountServicesServer() {}

// UnsafeAccountServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServicesServer will
// result in compilation errors.
type UnsafeAccountServicesServer interface {
	mustEmbedUnimplementedAccountServicesServer()
}

func RegisterAccountServicesServer(s grpc.ServiceRegistrar, srv AccountServicesServer) {
	s.RegisterService(&AccountServices_ServiceDesc, srv)
}

func _AccountServices_SignupService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServicesServer).SignupService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protopack.AccountServices/SignupService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServicesServer).SignupService(ctx, req.(*SignupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountServices_LoginService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServicesServer).LoginService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protopack.AccountServices/LoginService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServicesServer).LoginService(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountServices_ServiceDesc is the grpc.ServiceDesc for AccountServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protopack.AccountServices",
	HandlerType: (*AccountServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignupService",
			Handler:    _AccountServices_SignupService_Handler,
		},
		{
			MethodName: "LoginService",
			Handler:    _AccountServices_LoginService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
