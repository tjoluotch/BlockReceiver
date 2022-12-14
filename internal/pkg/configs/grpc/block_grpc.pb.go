// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package blockresolver

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

// BlockResolverClient is the client API for BlockResolver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockResolverClient interface {
	SendBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*BlockId, error)
}

type blockResolverClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockResolverClient(cc grpc.ClientConnInterface) BlockResolverClient {
	return &blockResolverClient{cc}
}

func (c *blockResolverClient) SendBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*BlockId, error) {
	out := new(BlockId)
	err := c.cc.Invoke(ctx, "/BlockResolver/SendBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockResolverServer is the server API for BlockResolver service.
// All implementations must embed UnimplementedBlockResolverServer
// for forward compatibility
type BlockResolverServer interface {
	SendBlock(context.Context, *Block) (*BlockId, error)
	mustEmbedUnimplementedBlockResolverServer()
}

// UnimplementedBlockResolverServer must be embedded to have forward compatible implementations.
type UnimplementedBlockResolverServer struct {
}

func (UnimplementedBlockResolverServer) SendBlock(context.Context, *Block) (*BlockId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBlock not implemented")
}
func (UnimplementedBlockResolverServer) mustEmbedUnimplementedBlockResolverServer() {}

// UnsafeBlockResolverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockResolverServer will
// result in compilation errors.
type UnsafeBlockResolverServer interface {
	mustEmbedUnimplementedBlockResolverServer()
}

func RegisterBlockResolverServer(s grpc.ServiceRegistrar, srv BlockResolverServer) {
	s.RegisterService(&BlockResolver_ServiceDesc, srv)
}

func _BlockResolver_SendBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockResolverServer).SendBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockResolver/SendBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockResolverServer).SendBlock(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockResolver_ServiceDesc is the grpc.ServiceDesc for BlockResolver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockResolver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BlockResolver",
	HandlerType: (*BlockResolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendBlock",
			Handler:    _BlockResolver_SendBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "configs/grpc/block.proto",
}
