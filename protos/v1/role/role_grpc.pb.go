// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: protos/v1/role/role.proto

package rolepb

import (
	context "context"
	common "github.com/xdorro/base-project-proto/protos/v1/common"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleServiceClient interface {
	// Find all Roles
	FindAllRoles(ctx context.Context, in *FindAllRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error)
	// Find Role by ID
	FindRoleByID(ctx context.Context, in *common.UUIDRequest, opts ...grpc.CallOption) (*Role, error)
	// Create new Role
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*status.Status, error)
	// Update Role by ID
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*status.Status, error)
	// Delete Role
	DeleteRole(ctx context.Context, in *common.UUIDRequest, opts ...grpc.CallOption) (*status.Status, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) FindAllRoles(ctx context.Context, in *FindAllRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error) {
	out := new(ListRolesResponse)
	err := c.cc.Invoke(ctx, "/rolepb.RoleService/FindAllRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) FindRoleByID(ctx context.Context, in *common.UUIDRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/rolepb.RoleService/FindRoleByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*status.Status, error) {
	out := new(status.Status)
	err := c.cc.Invoke(ctx, "/rolepb.RoleService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*status.Status, error) {
	out := new(status.Status)
	err := c.cc.Invoke(ctx, "/rolepb.RoleService/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) DeleteRole(ctx context.Context, in *common.UUIDRequest, opts ...grpc.CallOption) (*status.Status, error) {
	out := new(status.Status)
	err := c.cc.Invoke(ctx, "/rolepb.RoleService/DeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
// All implementations should embed UnimplementedRoleServiceServer
// for forward compatibility
type RoleServiceServer interface {
	// Find all Roles
	FindAllRoles(context.Context, *FindAllRolesRequest) (*ListRolesResponse, error)
	// Find Role by ID
	FindRoleByID(context.Context, *common.UUIDRequest) (*Role, error)
	// Create new Role
	CreateRole(context.Context, *CreateRoleRequest) (*status.Status, error)
	// Update Role by ID
	UpdateRole(context.Context, *UpdateRoleRequest) (*status.Status, error)
	// Delete Role
	DeleteRole(context.Context, *common.UUIDRequest) (*status.Status, error)
}

// UnimplementedRoleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRoleServiceServer struct {
}

func (UnimplementedRoleServiceServer) FindAllRoles(context.Context, *FindAllRolesRequest) (*ListRolesResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method FindAllRoles not implemented")
}
func (UnimplementedRoleServiceServer) FindRoleByID(context.Context, *common.UUIDRequest) (*Role, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method FindRoleByID not implemented")
}
func (UnimplementedRoleServiceServer) CreateRole(context.Context, *CreateRoleRequest) (*status.Status, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedRoleServiceServer) UpdateRole(context.Context, *UpdateRoleRequest) (*status.Status, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedRoleServiceServer) DeleteRole(context.Context, *common.UUIDRequest) (*status.Status, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}

// UnsafeRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServiceServer will
// result in compilation errors.
type UnsafeRoleServiceServer interface {
	mustEmbedUnimplementedRoleServiceServer()
}

func RegisterRoleServiceServer(s grpc.ServiceRegistrar, srv RoleServiceServer) {
	s.RegisterService(&RoleService_ServiceDesc, srv)
}

func _RoleService_FindAllRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).FindAllRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rolepb.RoleService/FindAllRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).FindAllRoles(ctx, req.(*FindAllRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_FindRoleByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.UUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).FindRoleByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rolepb.RoleService/FindRoleByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).FindRoleByID(ctx, req.(*common.UUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rolepb.RoleService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rolepb.RoleService/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.UUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rolepb.RoleService/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).DeleteRole(ctx, req.(*common.UUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleService_ServiceDesc is the grpc.ServiceDesc for RoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rolepb.RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllRoles",
			Handler:    _RoleService_FindAllRoles_Handler,
		},
		{
			MethodName: "FindRoleByID",
			Handler:    _RoleService_FindRoleByID_Handler,
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/v1/role/role.proto",
}
