// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package delivery

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

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServiceClient interface {
	AddUser(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*NoArgs, error)
	RemoveUser(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*NoArgs, error)
	ListAllUsers(ctx context.Context, in *NoArgs, opts ...grpc.CallOption) (*ListOfUsersResponse, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) AddUser(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*NoArgs, error) {
	out := new(NoArgs)
	err := c.cc.Invoke(ctx, "/service.UsersService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) RemoveUser(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*NoArgs, error) {
	out := new(NoArgs)
	err := c.cc.Invoke(ctx, "/service.UsersService/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) ListAllUsers(ctx context.Context, in *NoArgs, opts ...grpc.CallOption) (*ListOfUsersResponse, error) {
	out := new(ListOfUsersResponse)
	err := c.cc.Invoke(ctx, "/service.UsersService/ListAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations should embed UnimplementedUsersServiceServer
// for forward compatibility
type UsersServiceServer interface {
	AddUser(context.Context, *UserInfoRequest) (*NoArgs, error)
	RemoveUser(context.Context, *UserInfoRequest) (*NoArgs, error)
	ListAllUsers(context.Context, *NoArgs) (*ListOfUsersResponse, error)
}

// UnimplementedUsersServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (UnimplementedUsersServiceServer) AddUser(context.Context, *UserInfoRequest) (*NoArgs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUsersServiceServer) RemoveUser(context.Context, *UserInfoRequest) (*NoArgs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedUsersServiceServer) ListAllUsers(context.Context, *NoArgs) (*ListOfUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllUsers not implemented")
}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UsersService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).AddUser(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UsersService/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).RemoveUser(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_ListAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).ListAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.UsersService/ListAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).ListAllUsers(ctx, req.(*NoArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UsersService_AddUser_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _UsersService_RemoveUser_Handler,
		},
		{
			MethodName: "ListAllUsers",
			Handler:    _UsersService_ListAllUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/users.proto",
}
