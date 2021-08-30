// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Forum

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

// ForumClient is the client API for Forum service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ForumClient interface {
	Speak(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	GetMessage(ctx context.Context, in *Uid, opts ...grpc.CallOption) (*Messages, error)
	GetHistory(ctx context.Context, in *Classid, opts ...grpc.CallOption) (*Messages, error)
}

type forumClient struct {
	cc grpc.ClientConnInterface
}

func NewForumClient(cc grpc.ClientConnInterface) ForumClient {
	return &forumClient{cc}
}

func (c *forumClient) Speak(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/Forum.Forum/speak", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumClient) GetMessage(ctx context.Context, in *Uid, opts ...grpc.CallOption) (*Messages, error) {
	out := new(Messages)
	err := c.cc.Invoke(ctx, "/Forum.Forum/get_message", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forumClient) GetHistory(ctx context.Context, in *Classid, opts ...grpc.CallOption) (*Messages, error) {
	out := new(Messages)
	err := c.cc.Invoke(ctx, "/Forum.Forum/get_history", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForumServer is the server API for Forum service.
// All implementations must embed UnimplementedForumServer
// for forward compatibility
type ForumServer interface {
	Speak(context.Context, *Message) (*Message, error)
	GetMessage(context.Context, *Uid) (*Messages, error)
	GetHistory(context.Context, *Classid) (*Messages, error)
	mustEmbedUnimplementedForumServer()
}

// UnimplementedForumServer must be embedded to have forward compatible implementations.
type UnimplementedForumServer struct {
}

func (UnimplementedForumServer) Speak(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Speak not implemented")
}
func (UnimplementedForumServer) GetMessage(context.Context, *Uid) (*Messages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedForumServer) GetHistory(context.Context, *Classid) (*Messages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistory not implemented")
}
func (UnimplementedForumServer) mustEmbedUnimplementedForumServer() {}

// UnsafeForumServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ForumServer will
// result in compilation errors.
type UnsafeForumServer interface {
	mustEmbedUnimplementedForumServer()
}

func RegisterForumServer(s grpc.ServiceRegistrar, srv ForumServer) {
	s.RegisterService(&Forum_ServiceDesc, srv)
}

func _Forum_Speak_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServer).Speak(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Forum.Forum/speak",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServer).Speak(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forum_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Forum.Forum/get_message",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServer).GetMessage(ctx, req.(*Uid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forum_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Classid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForumServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Forum.Forum/get_history",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForumServer).GetHistory(ctx, req.(*Classid))
	}
	return interceptor(ctx, in, info, handler)
}

// Forum_ServiceDesc is the grpc.ServiceDesc for Forum service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Forum_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Forum.Forum",
	HandlerType: (*ForumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "speak",
			Handler:    _Forum_Speak_Handler,
		},
		{
			MethodName: "get_message",
			Handler:    _Forum_GetMessage_Handler,
		},
		{
			MethodName: "get_history",
			Handler:    _Forum_GetHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Forum.proto",
}
