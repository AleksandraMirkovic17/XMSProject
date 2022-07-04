// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: connection_service.proto

package connection_service

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

// ConnectionServiceClient is the client API for ConnectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionServiceClient interface {
	GetFriends(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error)
	GetBlockeds(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error)
	GetFriendRequests(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*ActionResult, error)
	AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*ActionResult, error)
	RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*ActionResult, error)
	AddBlockUser(ctx context.Context, in *AddBlockUserRequest, opts ...grpc.CallOption) (*ActionResult, error)
	UnblockUser(ctx context.Context, in *UnblockUserRequest, opts ...grpc.CallOption) (*ActionResult, error)
	SendFriendRequest(ctx context.Context, in *SendFriendRequestRequest, opts ...grpc.CallOption) (*ActionResult, error)
	UnsendFriendRequest(ctx context.Context, in *UnsendFriendRequestRequest, opts ...grpc.CallOption) (*ActionResult, error)
	GetRecommendation(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error)
	GetConnectionDetail(ctx context.Context, in *GetConnectionDetailRequest, opts ...grpc.CallOption) (*ConnectionDetail, error)
	ChangePrivacy(ctx context.Context, in *ChangePrivacyRequest, opts ...grpc.CallOption) (*ActionResult, error)
	GetMyContacts(ctx context.Context, in *GetMyContactsRequest, opts ...grpc.CallOption) (*ContactsResponse, error)
}

type connectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionServiceClient(cc grpc.ClientConnInterface) ConnectionServiceClient {
	return &connectionServiceClient{cc}
}

func (c *connectionServiceClient) GetFriends(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/GetFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetBlockeds(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/GetBlockeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetFriendRequests(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/GetFriendRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*ActionResult, error) {
	out := new(ActionResult)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*ActionResult, error) {
	out := new(ActionResult)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*ActionResult, error) {
	out := new(ActionResult)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/RemoveFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) AddBlockUser(ctx context.Context, in *AddBlockUserRequest, opts ...grpc.CallOption) (*ActionResult, error) {
	out := new(ActionResult)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/AddBlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) UnblockUser(ctx context.Context, in *UnblockUserRequest, opts ...grpc.CallOption) (*ActionResult, error) {
	out := new(ActionResult)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/UnblockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) SendFriendRequest(ctx context.Context, in *SendFriendRequestRequest, opts ...grpc.CallOption) (*ActionResult, error) {
	out := new(ActionResult)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/SendFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) UnsendFriendRequest(ctx context.Context, in *UnsendFriendRequestRequest, opts ...grpc.CallOption) (*ActionResult, error) {
	out := new(ActionResult)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/UnsendFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetRecommendation(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/GetRecommendation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetConnectionDetail(ctx context.Context, in *GetConnectionDetailRequest, opts ...grpc.CallOption) (*ConnectionDetail, error) {
	out := new(ConnectionDetail)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/GetConnectionDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) ChangePrivacy(ctx context.Context, in *ChangePrivacyRequest, opts ...grpc.CallOption) (*ActionResult, error) {
	out := new(ActionResult)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/ChangePrivacy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetMyContacts(ctx context.Context, in *GetMyContactsRequest, opts ...grpc.CallOption) (*ContactsResponse, error) {
	out := new(ContactsResponse)
	err := c.cc.Invoke(ctx, "/connection_service.ConnectionService/GetMyContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectionServiceServer is the server API for ConnectionService service.
// All implementations must embed UnimplementedConnectionServiceServer
// for forward compatibility
type ConnectionServiceServer interface {
	GetFriends(context.Context, *GetRequest) (*Users, error)
	GetBlockeds(context.Context, *GetRequest) (*Users, error)
	GetFriendRequests(context.Context, *GetRequest) (*Users, error)
	Register(context.Context, *RegisterRequest) (*ActionResult, error)
	AddFriend(context.Context, *AddFriendRequest) (*ActionResult, error)
	RemoveFriend(context.Context, *RemoveFriendRequest) (*ActionResult, error)
	AddBlockUser(context.Context, *AddBlockUserRequest) (*ActionResult, error)
	UnblockUser(context.Context, *UnblockUserRequest) (*ActionResult, error)
	SendFriendRequest(context.Context, *SendFriendRequestRequest) (*ActionResult, error)
	UnsendFriendRequest(context.Context, *UnsendFriendRequestRequest) (*ActionResult, error)
	GetRecommendation(context.Context, *GetRequest) (*Users, error)
	GetConnectionDetail(context.Context, *GetConnectionDetailRequest) (*ConnectionDetail, error)
	ChangePrivacy(context.Context, *ChangePrivacyRequest) (*ActionResult, error)
	GetMyContacts(context.Context, *GetMyContactsRequest) (*ContactsResponse, error)
	mustEmbedUnimplementedConnectionServiceServer()
}

// UnimplementedConnectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionServiceServer struct {
}

func (UnimplementedConnectionServiceServer) GetFriends(context.Context, *GetRequest) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriends not implemented")
}
func (UnimplementedConnectionServiceServer) GetBlockeds(context.Context, *GetRequest) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockeds not implemented")
}
func (UnimplementedConnectionServiceServer) GetFriendRequests(context.Context, *GetRequest) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendRequests not implemented")
}
func (UnimplementedConnectionServiceServer) Register(context.Context, *RegisterRequest) (*ActionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedConnectionServiceServer) AddFriend(context.Context, *AddFriendRequest) (*ActionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedConnectionServiceServer) RemoveFriend(context.Context, *RemoveFriendRequest) (*ActionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFriend not implemented")
}
func (UnimplementedConnectionServiceServer) AddBlockUser(context.Context, *AddBlockUserRequest) (*ActionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlockUser not implemented")
}
func (UnimplementedConnectionServiceServer) UnblockUser(context.Context, *UnblockUserRequest) (*ActionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnblockUser not implemented")
}
func (UnimplementedConnectionServiceServer) SendFriendRequest(context.Context, *SendFriendRequestRequest) (*ActionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendFriendRequest not implemented")
}
func (UnimplementedConnectionServiceServer) UnsendFriendRequest(context.Context, *UnsendFriendRequestRequest) (*ActionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsendFriendRequest not implemented")
}
func (UnimplementedConnectionServiceServer) GetRecommendation(context.Context, *GetRequest) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendation not implemented")
}
func (UnimplementedConnectionServiceServer) GetConnectionDetail(context.Context, *GetConnectionDetailRequest) (*ConnectionDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectionDetail not implemented")
}
func (UnimplementedConnectionServiceServer) ChangePrivacy(context.Context, *ChangePrivacyRequest) (*ActionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePrivacy not implemented")
}
func (UnimplementedConnectionServiceServer) GetMyContacts(context.Context, *GetMyContactsRequest) (*ContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyContacts not implemented")
}
func (UnimplementedConnectionServiceServer) mustEmbedUnimplementedConnectionServiceServer() {}

// UnsafeConnectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionServiceServer will
// result in compilation errors.
type UnsafeConnectionServiceServer interface {
	mustEmbedUnimplementedConnectionServiceServer()
}

func RegisterConnectionServiceServer(s grpc.ServiceRegistrar, srv ConnectionServiceServer) {
	s.RegisterService(&ConnectionService_ServiceDesc, srv)
}

func _ConnectionService_GetFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/GetFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetFriends(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetBlockeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetBlockeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/GetBlockeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetBlockeds(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetFriendRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetFriendRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/GetFriendRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetFriendRequests(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).AddFriend(ctx, req.(*AddFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_RemoveFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).RemoveFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/RemoveFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).RemoveFriend(ctx, req.(*RemoveFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_AddBlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).AddBlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/AddBlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).AddBlockUser(ctx, req.(*AddBlockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_UnblockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnblockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).UnblockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/UnblockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).UnblockUser(ctx, req.(*UnblockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_SendFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).SendFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/SendFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).SendFriendRequest(ctx, req.(*SendFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_UnsendFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsendFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).UnsendFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/UnsendFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).UnsendFriendRequest(ctx, req.(*UnsendFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetRecommendation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetRecommendation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/GetRecommendation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetRecommendation(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetConnectionDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectionDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetConnectionDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/GetConnectionDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetConnectionDetail(ctx, req.(*GetConnectionDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_ChangePrivacy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePrivacyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).ChangePrivacy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/ChangePrivacy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).ChangePrivacy(ctx, req.(*ChangePrivacyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetMyContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyContactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetMyContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection_service.ConnectionService/GetMyContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetMyContacts(ctx, req.(*GetMyContactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectionService_ServiceDesc is the grpc.ServiceDesc for ConnectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connection_service.ConnectionService",
	HandlerType: (*ConnectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFriends",
			Handler:    _ConnectionService_GetFriends_Handler,
		},
		{
			MethodName: "GetBlockeds",
			Handler:    _ConnectionService_GetBlockeds_Handler,
		},
		{
			MethodName: "GetFriendRequests",
			Handler:    _ConnectionService_GetFriendRequests_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _ConnectionService_Register_Handler,
		},
		{
			MethodName: "AddFriend",
			Handler:    _ConnectionService_AddFriend_Handler,
		},
		{
			MethodName: "RemoveFriend",
			Handler:    _ConnectionService_RemoveFriend_Handler,
		},
		{
			MethodName: "AddBlockUser",
			Handler:    _ConnectionService_AddBlockUser_Handler,
		},
		{
			MethodName: "UnblockUser",
			Handler:    _ConnectionService_UnblockUser_Handler,
		},
		{
			MethodName: "SendFriendRequest",
			Handler:    _ConnectionService_SendFriendRequest_Handler,
		},
		{
			MethodName: "UnsendFriendRequest",
			Handler:    _ConnectionService_UnsendFriendRequest_Handler,
		},
		{
			MethodName: "GetRecommendation",
			Handler:    _ConnectionService_GetRecommendation_Handler,
		},
		{
			MethodName: "GetConnectionDetail",
			Handler:    _ConnectionService_GetConnectionDetail_Handler,
		},
		{
			MethodName: "ChangePrivacy",
			Handler:    _ConnectionService_ChangePrivacy_Handler,
		},
		{
			MethodName: "GetMyContacts",
			Handler:    _ConnectionService_GetMyContacts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connection_service.proto",
}
