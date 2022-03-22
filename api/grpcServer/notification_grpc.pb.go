// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/grpcServer/protos/notification.proto

package grpcServer

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

// NotificationClient is the client API for Notification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationClient interface {
	ListNoti(ctx context.Context, in *ListNotiRequest, opts ...grpc.CallOption) (*ListNotiResponse, error)
	CreateNoti(ctx context.Context, in *CreateNotiRequest, opts ...grpc.CallOption) (*CreateNotiResponse, error)
	SendNoti(ctx context.Context, in *SendNotiRequest, opts ...grpc.CallOption) (*SendNotiResponse, error)
}

type notificationClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationClient(cc grpc.ClientConnInterface) NotificationClient {
	return &notificationClient{cc}
}

func (c *notificationClient) ListNoti(ctx context.Context, in *ListNotiRequest, opts ...grpc.CallOption) (*ListNotiResponse, error) {
	out := new(ListNotiResponse)
	err := c.cc.Invoke(ctx, "/grpcServer.Notification/ListNoti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) CreateNoti(ctx context.Context, in *CreateNotiRequest, opts ...grpc.CallOption) (*CreateNotiResponse, error) {
	out := new(CreateNotiResponse)
	err := c.cc.Invoke(ctx, "/grpcServer.Notification/CreateNoti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationClient) SendNoti(ctx context.Context, in *SendNotiRequest, opts ...grpc.CallOption) (*SendNotiResponse, error) {
	out := new(SendNotiResponse)
	err := c.cc.Invoke(ctx, "/grpcServer.Notification/SendNoti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServer is the server API for Notification service.
// All implementations must embed UnimplementedNotificationServer
// for forward compatibility
type NotificationServer interface {
	ListNoti(context.Context, *ListNotiRequest) (*ListNotiResponse, error)
	CreateNoti(context.Context, *CreateNotiRequest) (*CreateNotiResponse, error)
	SendNoti(context.Context, *SendNotiRequest) (*SendNotiResponse, error)
	mustEmbedUnimplementedNotificationServer()
}

// UnimplementedNotificationServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServer struct {
}

func (UnimplementedNotificationServer) ListNoti(context.Context, *ListNotiRequest) (*ListNotiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNoti not implemented")
}
func (UnimplementedNotificationServer) CreateNoti(context.Context, *CreateNotiRequest) (*CreateNotiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNoti not implemented")
}
func (UnimplementedNotificationServer) SendNoti(context.Context, *SendNotiRequest) (*SendNotiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNoti not implemented")
}
func (UnimplementedNotificationServer) mustEmbedUnimplementedNotificationServer() {}

// UnsafeNotificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServer will
// result in compilation errors.
type UnsafeNotificationServer interface {
	mustEmbedUnimplementedNotificationServer()
}

func RegisterNotificationServer(s grpc.ServiceRegistrar, srv NotificationServer) {
	s.RegisterService(&Notification_ServiceDesc, srv)
}

func _Notification_ListNoti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).ListNoti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcServer.Notification/ListNoti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).ListNoti(ctx, req.(*ListNotiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_CreateNoti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).CreateNoti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcServer.Notification/CreateNoti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).CreateNoti(ctx, req.(*CreateNotiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notification_SendNoti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServer).SendNoti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcServer.Notification/SendNoti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServer).SendNoti(ctx, req.(*SendNotiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notification_ServiceDesc is the grpc.ServiceDesc for Notification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcServer.Notification",
	HandlerType: (*NotificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNoti",
			Handler:    _Notification_ListNoti_Handler,
		},
		{
			MethodName: "CreateNoti",
			Handler:    _Notification_CreateNoti_Handler,
		},
		{
			MethodName: "SendNoti",
			Handler:    _Notification_SendNoti_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpcServer/protos/notification.proto",
}
