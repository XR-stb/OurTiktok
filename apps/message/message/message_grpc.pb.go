// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: message.proto

package message

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

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageClient interface {
	Action(ctx context.Context, in *MessageActionReq, opts ...grpc.CallOption) (*MessageActionRes, error)
	Chat(ctx context.Context, in *MessageChatReq, opts ...grpc.CallOption) (*MessageChatRes, error)
	GetLastMsg(ctx context.Context, in *GetLastMsgReq, opts ...grpc.CallOption) (*GetLastMsgRes, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) Action(ctx context.Context, in *MessageActionReq, opts ...grpc.CallOption) (*MessageActionRes, error) {
	out := new(MessageActionRes)
	err := c.cc.Invoke(ctx, "/message.Message/Action", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) Chat(ctx context.Context, in *MessageChatReq, opts ...grpc.CallOption) (*MessageChatRes, error) {
	out := new(MessageChatRes)
	err := c.cc.Invoke(ctx, "/message.Message/Chat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) GetLastMsg(ctx context.Context, in *GetLastMsgReq, opts ...grpc.CallOption) (*GetLastMsgRes, error) {
	out := new(GetLastMsgRes)
	err := c.cc.Invoke(ctx, "/message.Message/GetLastMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
// All implementations must embed UnimplementedMessageServer
// for forward compatibility
type MessageServer interface {
	Action(context.Context, *MessageActionReq) (*MessageActionRes, error)
	Chat(context.Context, *MessageChatReq) (*MessageChatRes, error)
	GetLastMsg(context.Context, *GetLastMsgReq) (*GetLastMsgRes, error)
	mustEmbedUnimplementedMessageServer()
}

// UnimplementedMessageServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (UnimplementedMessageServer) Action(context.Context, *MessageActionReq) (*MessageActionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Action not implemented")
}
func (UnimplementedMessageServer) Chat(context.Context, *MessageChatReq) (*MessageChatRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedMessageServer) GetLastMsg(context.Context, *GetLastMsgReq) (*GetLastMsgRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastMsg not implemented")
}
func (UnimplementedMessageServer) mustEmbedUnimplementedMessageServer() {}

// UnsafeMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServer will
// result in compilation errors.
type UnsafeMessageServer interface {
	mustEmbedUnimplementedMessageServer()
}

func RegisterMessageServer(s grpc.ServiceRegistrar, srv MessageServer) {
	s.RegisterService(&Message_ServiceDesc, srv)
}

func _Message_Action_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).Action(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Message/Action",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).Action(ctx, req.(*MessageActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_Chat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).Chat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Message/Chat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).Chat(ctx, req.(*MessageChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_GetLastMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).GetLastMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Message/GetLastMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).GetLastMsg(ctx, req.(*GetLastMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Message_ServiceDesc is the grpc.ServiceDesc for Message service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Message_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Action",
			Handler:    _Message_Action_Handler,
		},
		{
			MethodName: "Chat",
			Handler:    _Message_Chat_Handler,
		},
		{
			MethodName: "GetLastMsg",
			Handler:    _Message_GetLastMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
