// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package NodePack

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

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClient interface {
	Permission(ctx context.Context, in *RequestPermission, opts ...grpc.CallOption) (*GivePermission, error)
	AccesCrit(ctx context.Context, in *GoIntoCrit, opts ...grpc.CallOption) (*ServerDoneInCrit, error)
	ExitCrit(ctx context.Context, in *ReleaseToken, opts ...grpc.CallOption) (*Empty, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Permission(ctx context.Context, in *RequestPermission, opts ...grpc.CallOption) (*GivePermission, error) {
	out := new(GivePermission)
	err := c.cc.Invoke(ctx, "/main.Node/Permission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) AccesCrit(ctx context.Context, in *GoIntoCrit, opts ...grpc.CallOption) (*ServerDoneInCrit, error) {
	out := new(ServerDoneInCrit)
	err := c.cc.Invoke(ctx, "/main.Node/AccesCrit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ExitCrit(ctx context.Context, in *ReleaseToken, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/main.Node/ExitCrit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
// All implementations must embed UnimplementedNodeServer
// for forward compatibility
type NodeServer interface {
	Permission(context.Context, *RequestPermission) (*GivePermission, error)
	AccesCrit(context.Context, *GoIntoCrit) (*ServerDoneInCrit, error)
	ExitCrit(context.Context, *ReleaseToken) (*Empty, error)
	
}

// UnimplementedNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (UnimplementedNodeServer) Permission(context.Context, *RequestPermission) (*GivePermission, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Permission not implemented")
}
func (UnimplementedNodeServer) AccesCrit(context.Context, *GoIntoCrit) (*ServerDoneInCrit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccesCrit not implemented")
}
func (UnimplementedNodeServer) ExitCrit(context.Context, *ReleaseToken) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExitCrit not implemented")
}
func (UnimplementedNodeServer) mustEmbedUnimplementedNodeServer() {}

// UnsafeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServer will
// result in compilation errors.
type UnsafeNodeServer interface {
	mustEmbedUnimplementedNodeServer()
}

func RegisterNodeServer(s grpc.ServiceRegistrar, srv NodeServer) {
	s.RegisterService(&Node_ServiceDesc, srv)
}

func _Node_Permission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPermission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Permission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Node/Permission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Permission(ctx, req.(*RequestPermission))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_AccesCrit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoIntoCrit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).AccesCrit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Node/AccesCrit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).AccesCrit(ctx, req.(*GoIntoCrit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ExitCrit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ExitCrit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Node/ExitCrit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ExitCrit(ctx, req.(*ReleaseToken))
	}
	return interceptor(ctx, in, info, handler)
}

// Node_ServiceDesc is the grpc.ServiceDesc for Node service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Node_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Permission",
			Handler:    _Node_Permission_Handler,
		},
		{
			MethodName: "AccesCrit",
			Handler:    _Node_AccesCrit_Handler,
		},
		{
			MethodName: "ExitCrit",
			Handler:    _Node_ExitCrit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "NodePack/node.proto",
}
