// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auth

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AuthorityServiceClient is the client API for AuthorityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorityServiceClient interface {
	// AuthorizeAction checks if the user is authorized to perform the action
	AuthorizeAction(ctx context.Context, in *Action, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authorityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorityServiceClient(cc grpc.ClientConnInterface) AuthorityServiceClient {
	return &authorityServiceClient{cc}
}

func (c *authorityServiceClient) AuthorizeAction(ctx context.Context, in *Action, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/auth.AuthorityService/AuthorizeAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorityServiceServer is the server API for AuthorityService service.
// All implementations should embed UnimplementedAuthorityServiceServer
// for forward compatibility
type AuthorityServiceServer interface {
	// AuthorizeAction checks if the user is authorized to perform the action
	AuthorizeAction(context.Context, *Action) (*empty.Empty, error)
}

// UnimplementedAuthorityServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAuthorityServiceServer struct {
}

func (UnimplementedAuthorityServiceServer) AuthorizeAction(context.Context, *Action) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeAction not implemented")
}

// UnsafeAuthorityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorityServiceServer will
// result in compilation errors.
type UnsafeAuthorityServiceServer interface {
	mustEmbedUnimplementedAuthorityServiceServer()
}

func RegisterAuthorityServiceServer(s grpc.ServiceRegistrar, srv AuthorityServiceServer) {
	s.RegisterService(&_AuthorityService_serviceDesc, srv)
}

func _AuthorityService_AuthorizeAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Action)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorityServiceServer).AuthorizeAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthorityService/AuthorizeAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorityServiceServer).AuthorizeAction(ctx, req.(*Action))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthorityService",
	HandlerType: (*AuthorityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthorizeAction",
			Handler:    _AuthorityService_AuthorizeAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authority_service.proto",
}
