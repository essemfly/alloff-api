// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/grpcServer/protos/hometab.proto

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

// HomeTabItemClient is the client API for HomeTabItem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HomeTabItemClient interface {
	GetHomeTabItem(ctx context.Context, in *GetHomeTabItemRequest, opts ...grpc.CallOption) (*GetHomeTabItemResponse, error)
	ListHomeTabItems(ctx context.Context, in *ListHomeTabItemsRequest, opts ...grpc.CallOption) (*ListHomeTabItemsResponse, error)
	EditHomeTabItem(ctx context.Context, in *EditHomeTabItemRequest, opts ...grpc.CallOption) (*EditHomeTabItemResponse, error)
	CreateHomeTabItem(ctx context.Context, in *CreateHomeTabItemRequest, opts ...grpc.CallOption) (*CreateHomeTabItemResponse, error)
}

type homeTabItemClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeTabItemClient(cc grpc.ClientConnInterface) HomeTabItemClient {
	return &homeTabItemClient{cc}
}

func (c *homeTabItemClient) GetHomeTabItem(ctx context.Context, in *GetHomeTabItemRequest, opts ...grpc.CallOption) (*GetHomeTabItemResponse, error) {
	out := new(GetHomeTabItemResponse)
	err := c.cc.Invoke(ctx, "/grpcServer.HomeTabItem/GetHomeTabItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeTabItemClient) ListHomeTabItems(ctx context.Context, in *ListHomeTabItemsRequest, opts ...grpc.CallOption) (*ListHomeTabItemsResponse, error) {
	out := new(ListHomeTabItemsResponse)
	err := c.cc.Invoke(ctx, "/grpcServer.HomeTabItem/ListHomeTabItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeTabItemClient) EditHomeTabItem(ctx context.Context, in *EditHomeTabItemRequest, opts ...grpc.CallOption) (*EditHomeTabItemResponse, error) {
	out := new(EditHomeTabItemResponse)
	err := c.cc.Invoke(ctx, "/grpcServer.HomeTabItem/EditHomeTabItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeTabItemClient) CreateHomeTabItem(ctx context.Context, in *CreateHomeTabItemRequest, opts ...grpc.CallOption) (*CreateHomeTabItemResponse, error) {
	out := new(CreateHomeTabItemResponse)
	err := c.cc.Invoke(ctx, "/grpcServer.HomeTabItem/CreateHomeTabItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeTabItemServer is the server API for HomeTabItem service.
// All implementations must embed UnimplementedHomeTabItemServer
// for forward compatibility
type HomeTabItemServer interface {
	GetHomeTabItem(context.Context, *GetHomeTabItemRequest) (*GetHomeTabItemResponse, error)
	ListHomeTabItems(context.Context, *ListHomeTabItemsRequest) (*ListHomeTabItemsResponse, error)
	EditHomeTabItem(context.Context, *EditHomeTabItemRequest) (*EditHomeTabItemResponse, error)
	CreateHomeTabItem(context.Context, *CreateHomeTabItemRequest) (*CreateHomeTabItemResponse, error)
	mustEmbedUnimplementedHomeTabItemServer()
}

// UnimplementedHomeTabItemServer must be embedded to have forward compatible implementations.
type UnimplementedHomeTabItemServer struct {
}

func (UnimplementedHomeTabItemServer) GetHomeTabItem(context.Context, *GetHomeTabItemRequest) (*GetHomeTabItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomeTabItem not implemented")
}
func (UnimplementedHomeTabItemServer) ListHomeTabItems(context.Context, *ListHomeTabItemsRequest) (*ListHomeTabItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHomeTabItems not implemented")
}
func (UnimplementedHomeTabItemServer) EditHomeTabItem(context.Context, *EditHomeTabItemRequest) (*EditHomeTabItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditHomeTabItem not implemented")
}
func (UnimplementedHomeTabItemServer) CreateHomeTabItem(context.Context, *CreateHomeTabItemRequest) (*CreateHomeTabItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHomeTabItem not implemented")
}
func (UnimplementedHomeTabItemServer) mustEmbedUnimplementedHomeTabItemServer() {}

// UnsafeHomeTabItemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeTabItemServer will
// result in compilation errors.
type UnsafeHomeTabItemServer interface {
	mustEmbedUnimplementedHomeTabItemServer()
}

func RegisterHomeTabItemServer(s grpc.ServiceRegistrar, srv HomeTabItemServer) {
	s.RegisterService(&HomeTabItem_ServiceDesc, srv)
}

func _HomeTabItem_GetHomeTabItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomeTabItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeTabItemServer).GetHomeTabItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcServer.HomeTabItem/GetHomeTabItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeTabItemServer).GetHomeTabItem(ctx, req.(*GetHomeTabItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeTabItem_ListHomeTabItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHomeTabItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeTabItemServer).ListHomeTabItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcServer.HomeTabItem/ListHomeTabItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeTabItemServer).ListHomeTabItems(ctx, req.(*ListHomeTabItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeTabItem_EditHomeTabItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditHomeTabItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeTabItemServer).EditHomeTabItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcServer.HomeTabItem/EditHomeTabItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeTabItemServer).EditHomeTabItem(ctx, req.(*EditHomeTabItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeTabItem_CreateHomeTabItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHomeTabItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeTabItemServer).CreateHomeTabItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcServer.HomeTabItem/CreateHomeTabItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeTabItemServer).CreateHomeTabItem(ctx, req.(*CreateHomeTabItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HomeTabItem_ServiceDesc is the grpc.ServiceDesc for HomeTabItem service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HomeTabItem_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcServer.HomeTabItem",
	HandlerType: (*HomeTabItemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHomeTabItem",
			Handler:    _HomeTabItem_GetHomeTabItem_Handler,
		},
		{
			MethodName: "ListHomeTabItems",
			Handler:    _HomeTabItem_ListHomeTabItems_Handler,
		},
		{
			MethodName: "EditHomeTabItem",
			Handler:    _HomeTabItem_EditHomeTabItem_Handler,
		},
		{
			MethodName: "CreateHomeTabItem",
			Handler:    _HomeTabItem_CreateHomeTabItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpcServer/protos/hometab.proto",
}
