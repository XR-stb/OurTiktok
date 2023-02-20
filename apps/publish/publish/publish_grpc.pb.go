// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: publish.proto

package publish

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

// PublishClient is the client API for Publish service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublishClient interface {
	Action(ctx context.Context, in *ActionReq, opts ...grpc.CallOption) (*ActionRes, error)
	List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListRes, error)
	GetVideos(ctx context.Context, in *GetVideosReq, opts ...grpc.CallOption) (*GetVideosRes, error)
	GetVideoIds(ctx context.Context, in *GetVideoIdsReq, opts ...grpc.CallOption) (*GetVideoIdsRes, error)
	GetWorkCount(ctx context.Context, in *GetWorkCountReq, opts ...grpc.CallOption) (*GetWorkCountRes, error)
}

type publishClient struct {
	cc grpc.ClientConnInterface
}

func NewPublishClient(cc grpc.ClientConnInterface) PublishClient {
	return &publishClient{cc}
}

func (c *publishClient) Action(ctx context.Context, in *ActionReq, opts ...grpc.CallOption) (*ActionRes, error) {
	out := new(ActionRes)
	err := c.cc.Invoke(ctx, "/publish.Publish/Action", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishClient) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListRes, error) {
	out := new(ListRes)
	err := c.cc.Invoke(ctx, "/publish.Publish/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishClient) GetVideos(ctx context.Context, in *GetVideosReq, opts ...grpc.CallOption) (*GetVideosRes, error) {
	out := new(GetVideosRes)
	err := c.cc.Invoke(ctx, "/publish.Publish/GetVideos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishClient) GetVideoIds(ctx context.Context, in *GetVideoIdsReq, opts ...grpc.CallOption) (*GetVideoIdsRes, error) {
	out := new(GetVideoIdsRes)
	err := c.cc.Invoke(ctx, "/publish.Publish/GetVideoIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishClient) GetWorkCount(ctx context.Context, in *GetWorkCountReq, opts ...grpc.CallOption) (*GetWorkCountRes, error) {
	out := new(GetWorkCountRes)
	err := c.cc.Invoke(ctx, "/publish.Publish/GetWorkCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublishServer is the server API for Publish service.
// All implementations must embed UnimplementedPublishServer
// for forward compatibility
type PublishServer interface {
	Action(context.Context, *ActionReq) (*ActionRes, error)
	List(context.Context, *ListReq) (*ListRes, error)
	GetVideos(context.Context, *GetVideosReq) (*GetVideosRes, error)
	GetVideoIds(context.Context, *GetVideoIdsReq) (*GetVideoIdsRes, error)
	GetWorkCount(context.Context, *GetWorkCountReq) (*GetWorkCountRes, error)
	mustEmbedUnimplementedPublishServer()
}

// UnimplementedPublishServer must be embedded to have forward compatible implementations.
type UnimplementedPublishServer struct {
}

func (UnimplementedPublishServer) Action(context.Context, *ActionReq) (*ActionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Action not implemented")
}
func (UnimplementedPublishServer) List(context.Context, *ListReq) (*ListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPublishServer) GetVideos(context.Context, *GetVideosReq) (*GetVideosRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideos not implemented")
}
func (UnimplementedPublishServer) GetVideoIds(context.Context, *GetVideoIdsReq) (*GetVideoIdsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoIds not implemented")
}
func (UnimplementedPublishServer) GetWorkCount(context.Context, *GetWorkCountReq) (*GetWorkCountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkCount not implemented")
}
func (UnimplementedPublishServer) mustEmbedUnimplementedPublishServer() {}

// UnsafePublishServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublishServer will
// result in compilation errors.
type UnsafePublishServer interface {
	mustEmbedUnimplementedPublishServer()
}

func RegisterPublishServer(s grpc.ServiceRegistrar, srv PublishServer) {
	s.RegisterService(&Publish_ServiceDesc, srv)
}

func _Publish_Action_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServer).Action(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publish.Publish/Action",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServer).Action(ctx, req.(*ActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publish_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publish.Publish/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServer).List(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publish_GetVideos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideosReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServer).GetVideos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publish.Publish/GetVideos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServer).GetVideos(ctx, req.(*GetVideosReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publish_GetVideoIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServer).GetVideoIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publish.Publish/GetVideoIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServer).GetVideoIds(ctx, req.(*GetVideoIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publish_GetWorkCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServer).GetWorkCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publish.Publish/GetWorkCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServer).GetWorkCount(ctx, req.(*GetWorkCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Publish_ServiceDesc is the grpc.ServiceDesc for Publish service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Publish_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "publish.Publish",
	HandlerType: (*PublishServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Action",
			Handler:    _Publish_Action_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Publish_List_Handler,
		},
		{
			MethodName: "GetVideos",
			Handler:    _Publish_GetVideos_Handler,
		},
		{
			MethodName: "GetVideoIds",
			Handler:    _Publish_GetVideoIds_Handler,
		},
		{
			MethodName: "GetWorkCount",
			Handler:    _Publish_GetWorkCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "publish.proto",
}
