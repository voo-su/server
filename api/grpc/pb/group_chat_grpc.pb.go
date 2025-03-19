// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.3
// source: group_chat.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GroupChatService_CreateGroupChat_FullMethodName         = "/group_chat.GroupChatService/CreateGroupChat"
	GroupChatService_GetGroupChat_FullMethodName            = "/group_chat.GroupChatService/GetGroupChat"
	GroupChatService_GetMembers_FullMethodName              = "/group_chat.GroupChatService/GetMembers"
	GroupChatService_AddUserToGroupChat_FullMethodName      = "/group_chat.GroupChatService/AddUserToGroupChat"
	GroupChatService_RemoveUserFromGroupChat_FullMethodName = "/group_chat.GroupChatService/RemoveUserFromGroupChat"
	GroupChatService_LeaveGroupChat_FullMethodName          = "/group_chat.GroupChatService/LeaveGroupChat"
	GroupChatService_DeleteGroupChat_FullMethodName         = "/group_chat.GroupChatService/DeleteGroupChat"
	GroupChatService_EditNameGroupChat_FullMethodName       = "/group_chat.GroupChatService/EditNameGroupChat"
	GroupChatService_EditAboutGroupChat_FullMethodName      = "/group_chat.GroupChatService/EditAboutGroupChat"
)

// GroupChatServiceClient is the client API for GroupChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupChatServiceClient interface {
	// Creating a group chat
	CreateGroupChat(ctx context.Context, in *CreateGroupChatRequest, opts ...grpc.CallOption) (*CreateGroupChatResponse, error)
	// Retrieving information about the group chat
	GetGroupChat(ctx context.Context, in *GetGroupChatRequest, opts ...grpc.CallOption) (*GetGroupChatResponse, error)
	// Retrieving the list of participants in the group chat
	GetMembers(ctx context.Context, in *GetMembersRequest, opts ...grpc.CallOption) (*GetMembersResponse, error)
	// Adding participants to the group chat
	AddUserToGroupChat(ctx context.Context, in *AddUserToGroupChatRequest, opts ...grpc.CallOption) (*AddUserToGroupChatResponse, error)
	// Removing a participant from the group chat
	RemoveUserFromGroupChat(ctx context.Context, in *RemoveUserFromGroupChatRequest, opts ...grpc.CallOption) (*RemoveUserFromGroupChatResponse, error)
	// A user leaving the group chat
	LeaveGroupChat(ctx context.Context, in *LeaveGroupChatRequest, opts ...grpc.CallOption) (*LeaveGroupChatResponse, error)
	// Deleting a group chat
	DeleteGroupChat(ctx context.Context, in *DeleteGroupChatRequest, opts ...grpc.CallOption) (*DeleteGroupChatResponse, error)
	// Editing the group chat name
	EditNameGroupChat(ctx context.Context, in *EditNameGroupChatRequest, opts ...grpc.CallOption) (*EditNameGroupChatResponse, error)
	// Editing the group chat description
	EditAboutGroupChat(ctx context.Context, in *EditAboutGroupChatRequest, opts ...grpc.CallOption) (*EditAboutGroupChatResponse, error)
}

type groupChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupChatServiceClient(cc grpc.ClientConnInterface) GroupChatServiceClient {
	return &groupChatServiceClient{cc}
}

func (c *groupChatServiceClient) CreateGroupChat(ctx context.Context, in *CreateGroupChatRequest, opts ...grpc.CallOption) (*CreateGroupChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateGroupChatResponse)
	err := c.cc.Invoke(ctx, GroupChatService_CreateGroupChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupChatServiceClient) GetGroupChat(ctx context.Context, in *GetGroupChatRequest, opts ...grpc.CallOption) (*GetGroupChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGroupChatResponse)
	err := c.cc.Invoke(ctx, GroupChatService_GetGroupChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupChatServiceClient) GetMembers(ctx context.Context, in *GetMembersRequest, opts ...grpc.CallOption) (*GetMembersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMembersResponse)
	err := c.cc.Invoke(ctx, GroupChatService_GetMembers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupChatServiceClient) AddUserToGroupChat(ctx context.Context, in *AddUserToGroupChatRequest, opts ...grpc.CallOption) (*AddUserToGroupChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserToGroupChatResponse)
	err := c.cc.Invoke(ctx, GroupChatService_AddUserToGroupChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupChatServiceClient) RemoveUserFromGroupChat(ctx context.Context, in *RemoveUserFromGroupChatRequest, opts ...grpc.CallOption) (*RemoveUserFromGroupChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveUserFromGroupChatResponse)
	err := c.cc.Invoke(ctx, GroupChatService_RemoveUserFromGroupChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupChatServiceClient) LeaveGroupChat(ctx context.Context, in *LeaveGroupChatRequest, opts ...grpc.CallOption) (*LeaveGroupChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeaveGroupChatResponse)
	err := c.cc.Invoke(ctx, GroupChatService_LeaveGroupChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupChatServiceClient) DeleteGroupChat(ctx context.Context, in *DeleteGroupChatRequest, opts ...grpc.CallOption) (*DeleteGroupChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteGroupChatResponse)
	err := c.cc.Invoke(ctx, GroupChatService_DeleteGroupChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupChatServiceClient) EditNameGroupChat(ctx context.Context, in *EditNameGroupChatRequest, opts ...grpc.CallOption) (*EditNameGroupChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditNameGroupChatResponse)
	err := c.cc.Invoke(ctx, GroupChatService_EditNameGroupChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupChatServiceClient) EditAboutGroupChat(ctx context.Context, in *EditAboutGroupChatRequest, opts ...grpc.CallOption) (*EditAboutGroupChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditAboutGroupChatResponse)
	err := c.cc.Invoke(ctx, GroupChatService_EditAboutGroupChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupChatServiceServer is the server API for GroupChatService service.
// All implementations must embed UnimplementedGroupChatServiceServer
// for forward compatibility.
type GroupChatServiceServer interface {
	// Creating a group chat
	CreateGroupChat(context.Context, *CreateGroupChatRequest) (*CreateGroupChatResponse, error)
	// Retrieving information about the group chat
	GetGroupChat(context.Context, *GetGroupChatRequest) (*GetGroupChatResponse, error)
	// Retrieving the list of participants in the group chat
	GetMembers(context.Context, *GetMembersRequest) (*GetMembersResponse, error)
	// Adding participants to the group chat
	AddUserToGroupChat(context.Context, *AddUserToGroupChatRequest) (*AddUserToGroupChatResponse, error)
	// Removing a participant from the group chat
	RemoveUserFromGroupChat(context.Context, *RemoveUserFromGroupChatRequest) (*RemoveUserFromGroupChatResponse, error)
	// A user leaving the group chat
	LeaveGroupChat(context.Context, *LeaveGroupChatRequest) (*LeaveGroupChatResponse, error)
	// Deleting a group chat
	DeleteGroupChat(context.Context, *DeleteGroupChatRequest) (*DeleteGroupChatResponse, error)
	// Editing the group chat name
	EditNameGroupChat(context.Context, *EditNameGroupChatRequest) (*EditNameGroupChatResponse, error)
	// Editing the group chat description
	EditAboutGroupChat(context.Context, *EditAboutGroupChatRequest) (*EditAboutGroupChatResponse, error)
	mustEmbedUnimplementedGroupChatServiceServer()
}

// UnimplementedGroupChatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGroupChatServiceServer struct{}

func (UnimplementedGroupChatServiceServer) CreateGroupChat(context.Context, *CreateGroupChatRequest) (*CreateGroupChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupChat not implemented")
}
func (UnimplementedGroupChatServiceServer) GetGroupChat(context.Context, *GetGroupChatRequest) (*GetGroupChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupChat not implemented")
}
func (UnimplementedGroupChatServiceServer) GetMembers(context.Context, *GetMembersRequest) (*GetMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMembers not implemented")
}
func (UnimplementedGroupChatServiceServer) AddUserToGroupChat(context.Context, *AddUserToGroupChatRequest) (*AddUserToGroupChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToGroupChat not implemented")
}
func (UnimplementedGroupChatServiceServer) RemoveUserFromGroupChat(context.Context, *RemoveUserFromGroupChatRequest) (*RemoveUserFromGroupChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserFromGroupChat not implemented")
}
func (UnimplementedGroupChatServiceServer) LeaveGroupChat(context.Context, *LeaveGroupChatRequest) (*LeaveGroupChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveGroupChat not implemented")
}
func (UnimplementedGroupChatServiceServer) DeleteGroupChat(context.Context, *DeleteGroupChatRequest) (*DeleteGroupChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupChat not implemented")
}
func (UnimplementedGroupChatServiceServer) EditNameGroupChat(context.Context, *EditNameGroupChatRequest) (*EditNameGroupChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditNameGroupChat not implemented")
}
func (UnimplementedGroupChatServiceServer) EditAboutGroupChat(context.Context, *EditAboutGroupChatRequest) (*EditAboutGroupChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAboutGroupChat not implemented")
}
func (UnimplementedGroupChatServiceServer) mustEmbedUnimplementedGroupChatServiceServer() {}
func (UnimplementedGroupChatServiceServer) testEmbeddedByValue()                          {}

// UnsafeGroupChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupChatServiceServer will
// result in compilation errors.
type UnsafeGroupChatServiceServer interface {
	mustEmbedUnimplementedGroupChatServiceServer()
}

func RegisterGroupChatServiceServer(s grpc.ServiceRegistrar, srv GroupChatServiceServer) {
	// If the following call pancis, it indicates UnimplementedGroupChatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GroupChatService_ServiceDesc, srv)
}

func _GroupChatService_CreateGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupChatServiceServer).CreateGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupChatService_CreateGroupChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupChatServiceServer).CreateGroupChat(ctx, req.(*CreateGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupChatService_GetGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupChatServiceServer).GetGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupChatService_GetGroupChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupChatServiceServer).GetGroupChat(ctx, req.(*GetGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupChatService_GetMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupChatServiceServer).GetMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupChatService_GetMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupChatServiceServer).GetMembers(ctx, req.(*GetMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupChatService_AddUserToGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupChatServiceServer).AddUserToGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupChatService_AddUserToGroupChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupChatServiceServer).AddUserToGroupChat(ctx, req.(*AddUserToGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupChatService_RemoveUserFromGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserFromGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupChatServiceServer).RemoveUserFromGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupChatService_RemoveUserFromGroupChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupChatServiceServer).RemoveUserFromGroupChat(ctx, req.(*RemoveUserFromGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupChatService_LeaveGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupChatServiceServer).LeaveGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupChatService_LeaveGroupChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupChatServiceServer).LeaveGroupChat(ctx, req.(*LeaveGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupChatService_DeleteGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupChatServiceServer).DeleteGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupChatService_DeleteGroupChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupChatServiceServer).DeleteGroupChat(ctx, req.(*DeleteGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupChatService_EditNameGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditNameGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupChatServiceServer).EditNameGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupChatService_EditNameGroupChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupChatServiceServer).EditNameGroupChat(ctx, req.(*EditNameGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupChatService_EditAboutGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditAboutGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupChatServiceServer).EditAboutGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupChatService_EditAboutGroupChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupChatServiceServer).EditAboutGroupChat(ctx, req.(*EditAboutGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupChatService_ServiceDesc is the grpc.ServiceDesc for GroupChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "group_chat.GroupChatService",
	HandlerType: (*GroupChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroupChat",
			Handler:    _GroupChatService_CreateGroupChat_Handler,
		},
		{
			MethodName: "GetGroupChat",
			Handler:    _GroupChatService_GetGroupChat_Handler,
		},
		{
			MethodName: "GetMembers",
			Handler:    _GroupChatService_GetMembers_Handler,
		},
		{
			MethodName: "AddUserToGroupChat",
			Handler:    _GroupChatService_AddUserToGroupChat_Handler,
		},
		{
			MethodName: "RemoveUserFromGroupChat",
			Handler:    _GroupChatService_RemoveUserFromGroupChat_Handler,
		},
		{
			MethodName: "LeaveGroupChat",
			Handler:    _GroupChatService_LeaveGroupChat_Handler,
		},
		{
			MethodName: "DeleteGroupChat",
			Handler:    _GroupChatService_DeleteGroupChat_Handler,
		},
		{
			MethodName: "EditNameGroupChat",
			Handler:    _GroupChatService_EditNameGroupChat_Handler,
		},
		{
			MethodName: "EditAboutGroupChat",
			Handler:    _GroupChatService_EditAboutGroupChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group_chat.proto",
}
