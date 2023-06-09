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

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleServiceClient interface {
	// returns existing role
	GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*Role, error)
	// creates a new role
	CreateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
	// updates existing role
	UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
	// deletes existing role
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// binds role to user in a scope of organization
	CreateRoleBinding(ctx context.Context, in *CreateRoleBindingRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// unbinds role from user in all scopes
	DeleteRoleBinding(ctx context.Context, in *DeleteRoleBindingRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/auth.RoleService/GetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) CreateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/auth.RoleService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/auth.RoleService/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/auth.RoleService/DeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) CreateRoleBinding(ctx context.Context, in *CreateRoleBindingRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/auth.RoleService/CreateRoleBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) DeleteRoleBinding(ctx context.Context, in *DeleteRoleBindingRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/auth.RoleService/DeleteRoleBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
// All implementations should embed UnimplementedRoleServiceServer
// for forward compatibility
type RoleServiceServer interface {
	// returns existing role
	GetRole(context.Context, *GetRoleRequest) (*Role, error)
	// creates a new role
	CreateRole(context.Context, *Role) (*Role, error)
	// updates existing role
	UpdateRole(context.Context, *Role) (*Role, error)
	// deletes existing role
	DeleteRole(context.Context, *DeleteRoleRequest) (*empty.Empty, error)
	// binds role to user in a scope of organization
	CreateRoleBinding(context.Context, *CreateRoleBindingRequest) (*empty.Empty, error)
	// unbinds role from user in all scopes
	DeleteRoleBinding(context.Context, *DeleteRoleBindingRequest) (*empty.Empty, error)
}

// UnimplementedRoleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRoleServiceServer struct {
}

func (UnimplementedRoleServiceServer) GetRole(context.Context, *GetRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedRoleServiceServer) CreateRole(context.Context, *Role) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedRoleServiceServer) UpdateRole(context.Context, *Role) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedRoleServiceServer) DeleteRole(context.Context, *DeleteRoleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedRoleServiceServer) CreateRoleBinding(context.Context, *CreateRoleBindingRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoleBinding not implemented")
}
func (UnimplementedRoleServiceServer) DeleteRoleBinding(context.Context, *DeleteRoleBindingRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoleBinding not implemented")
}

// UnsafeRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServiceServer will
// result in compilation errors.
type UnsafeRoleServiceServer interface {
	mustEmbedUnimplementedRoleServiceServer()
}

func RegisterRoleServiceServer(s grpc.ServiceRegistrar, srv RoleServiceServer) {
	s.RegisterService(&_RoleService_serviceDesc, srv)
}

func _RoleService_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.RoleService/GetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).GetRole(ctx, req.(*GetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.RoleService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).CreateRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.RoleService/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).UpdateRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.RoleService/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_CreateRoleBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).CreateRoleBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.RoleService/CreateRoleBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).CreateRoleBinding(ctx, req.(*CreateRoleBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_DeleteRoleBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleBindingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).DeleteRoleBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.RoleService/DeleteRoleBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).DeleteRoleBinding(ctx, req.(*DeleteRoleBindingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRole",
			Handler:    _RoleService_GetRole_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _RoleService_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _RoleService_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _RoleService_DeleteRole_Handler,
		},
		{
			MethodName: "CreateRoleBinding",
			Handler:    _RoleService_CreateRoleBinding_Handler,
		},
		{
			MethodName: "DeleteRoleBinding",
			Handler:    _RoleService_DeleteRoleBinding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role_service.proto",
}
