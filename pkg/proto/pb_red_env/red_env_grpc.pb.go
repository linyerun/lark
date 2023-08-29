// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: pb_red_env/red_env.proto

package pb_red_env

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
	RedEnv_GiveRedEnvelope_FullMethodName = "/pb_red_env.RedEnv/GiveRedEnvelope"
)

// RedEnvClient is the client API for RedEnv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RedEnvClient interface {
	GiveRedEnvelope(ctx context.Context, in *GiveRedEnvelopeReq, opts ...grpc.CallOption) (*GiveRedEnvelopeResp, error)
}

type redEnvClient struct {
	cc grpc.ClientConnInterface
}

func NewRedEnvClient(cc grpc.ClientConnInterface) RedEnvClient {
	return &redEnvClient{cc}
}

func (c *redEnvClient) GiveRedEnvelope(ctx context.Context, in *GiveRedEnvelopeReq, opts ...grpc.CallOption) (*GiveRedEnvelopeResp, error) {
	out := new(GiveRedEnvelopeResp)
	err := c.cc.Invoke(ctx, RedEnv_GiveRedEnvelope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedEnvServer is the server API for RedEnv service.
// All implementations must embed UnimplementedRedEnvServer
// for forward compatibility
type RedEnvServer interface {
	GiveRedEnvelope(context.Context, *GiveRedEnvelopeReq) (*GiveRedEnvelopeResp, error)
	mustEmbedUnimplementedRedEnvServer()
}

// UnimplementedRedEnvServer must be embedded to have forward compatible implementations.
type UnimplementedRedEnvServer struct {
}

func (UnimplementedRedEnvServer) GiveRedEnvelope(context.Context, *GiveRedEnvelopeReq) (*GiveRedEnvelopeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GiveRedEnvelope not implemented")
}
func (UnimplementedRedEnvServer) mustEmbedUnimplementedRedEnvServer() {}

// UnsafeRedEnvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RedEnvServer will
// result in compilation errors.
type UnsafeRedEnvServer interface {
	mustEmbedUnimplementedRedEnvServer()
}

func RegisterRedEnvServer(s grpc.ServiceRegistrar, srv RedEnvServer) {
	s.RegisterService(&RedEnv_ServiceDesc, srv)
}

func _RedEnv_GiveRedEnvelope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GiveRedEnvelopeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedEnvServer).GiveRedEnvelope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedEnv_GiveRedEnvelope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedEnvServer).GiveRedEnvelope(ctx, req.(*GiveRedEnvelopeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RedEnv_ServiceDesc is the grpc.ServiceDesc for RedEnv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedEnv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_red_env.RedEnv",
	HandlerType: (*RedEnvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GiveRedEnvelope",
			Handler:    _RedEnv_GiveRedEnvelope_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_red_env/red_env.proto",
}
