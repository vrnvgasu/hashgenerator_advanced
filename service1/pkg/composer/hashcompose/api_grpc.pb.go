// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.13.0
// source: api/api.proto

package hashcompose

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
	HashComposeService_Gen_FullMethodName = "/hashcompose.HashComposeService/Gen"
)

// HashComposeServiceClient is the client API for HashComposeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HashComposeServiceClient interface {
	Gen(ctx context.Context, in *StringList, opts ...grpc.CallOption) (HashComposeService_GenClient, error)
}

type hashComposeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHashComposeServiceClient(cc grpc.ClientConnInterface) HashComposeServiceClient {
	return &hashComposeServiceClient{cc}
}

func (c *hashComposeServiceClient) Gen(ctx context.Context, in *StringList, opts ...grpc.CallOption) (HashComposeService_GenClient, error) {
	stream, err := c.cc.NewStream(ctx, &HashComposeService_ServiceDesc.Streams[0], HashComposeService_Gen_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &hashComposeServiceGenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HashComposeService_GenClient interface {
	Recv() (*HashObj, error)
	grpc.ClientStream
}

type hashComposeServiceGenClient struct {
	grpc.ClientStream
}

func (x *hashComposeServiceGenClient) Recv() (*HashObj, error) {
	m := new(HashObj)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HashComposeServiceServer is the server API for HashComposeService service.
// All implementations must embed UnimplementedHashComposeServiceServer
// for forward compatibility
type HashComposeServiceServer interface {
	Gen(*StringList, HashComposeService_GenServer) error
	mustEmbedUnimplementedHashComposeServiceServer()
}

// UnimplementedHashComposeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHashComposeServiceServer struct {
}

func (UnimplementedHashComposeServiceServer) Gen(*StringList, HashComposeService_GenServer) error {
	return status.Errorf(codes.Unimplemented, "method Gen not implemented")
}
func (UnimplementedHashComposeServiceServer) mustEmbedUnimplementedHashComposeServiceServer() {}

// UnsafeHashComposeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HashComposeServiceServer will
// result in compilation errors.
type UnsafeHashComposeServiceServer interface {
	mustEmbedUnimplementedHashComposeServiceServer()
}

func RegisterHashComposeServiceServer(s grpc.ServiceRegistrar, srv HashComposeServiceServer) {
	s.RegisterService(&HashComposeService_ServiceDesc, srv)
}

func _HashComposeService_Gen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StringList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HashComposeServiceServer).Gen(m, &hashComposeServiceGenServer{stream})
}

type HashComposeService_GenServer interface {
	Send(*HashObj) error
	grpc.ServerStream
}

type hashComposeServiceGenServer struct {
	grpc.ServerStream
}

func (x *hashComposeServiceGenServer) Send(m *HashObj) error {
	return x.ServerStream.SendMsg(m)
}

// HashComposeService_ServiceDesc is the grpc.ServiceDesc for HashComposeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HashComposeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hashcompose.HashComposeService",
	HandlerType: (*HashComposeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Gen",
			Handler:       _HashComposeService_Gen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/api.proto",
}