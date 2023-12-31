// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: shard.proto

package grpc

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

const (
	Shard_Get_FullMethodName = "/Shard/Proxy"
)

// ShardClient is the client API for Shard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShardClient interface {
	Get(ctx context.Context, in *GetShardRequest, opts ...grpc.CallOption) (*GetShardResponse, error)
}

type shardClient struct {
	cc grpc.ClientConnInterface
}

func NewShardClient(cc grpc.ClientConnInterface) ShardClient {
	return &shardClient{cc}
}

func (c *shardClient) Get(ctx context.Context, in *GetShardRequest, opts ...grpc.CallOption) (*GetShardResponse, error) {
	out := new(GetShardResponse)
	err := c.cc.Invoke(ctx, Shard_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShardServer is the server API for Shard service.
// All implementations must embed UnimplementedShardServer
// for forward compatibility
type ShardServer interface {
	Get(context.Context, *GetShardRequest) (*GetShardResponse, error)
	mustEmbedUnimplementedShardServer()
}

// UnimplementedShardServer must be embedded to have forward compatible implementations.
type UnimplementedShardServer struct {
}

func (UnimplementedShardServer) Get(context.Context, *GetShardRequest) (*GetShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proxy not implemented")
}
func (UnimplementedShardServer) mustEmbedUnimplementedShardServer() {}

// UnsafeShardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShardServer will
// result in compilation errors.
type UnsafeShardServer interface {
	mustEmbedUnimplementedShardServer()
}

func RegisterShardServer(s grpc.ServiceRegistrar, srv ShardServer) {
	s.RegisterService(&Shard_ServiceDesc, srv)
}

func _Shard_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShardServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shard_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShardServer).Get(ctx, req.(*GetShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shard_ServiceDesc is the grpc.ServiceDesc for Shard service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shard_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Shard",
	HandlerType: (*ShardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Proxy",
			Handler:    _Shard_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shard.proto",
}
