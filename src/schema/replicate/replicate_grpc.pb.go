// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: schema/replicate/replicate.proto

package replicate

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

// ReplicateServiceClient is the client API for ReplicateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplicateServiceClient interface {
	NotifyToCopy(ctx context.Context, in *NotifyToCopyRequest, opts ...grpc.CallOption) (*NotifyToCopyResponse, error)
	Copying(ctx context.Context, in *CopyingRequest, opts ...grpc.CallOption) (ReplicateService_CopyingClient, error)
	ConfirmCopy(ctx context.Context, in *ConfirmCopyRequest, opts ...grpc.CallOption) (*ConfirmCopyResponse, error)
}

type replicateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicateServiceClient(cc grpc.ClientConnInterface) ReplicateServiceClient {
	return &replicateServiceClient{cc}
}

func (c *replicateServiceClient) NotifyToCopy(ctx context.Context, in *NotifyToCopyRequest, opts ...grpc.CallOption) (*NotifyToCopyResponse, error) {
	out := new(NotifyToCopyResponse)
	err := c.cc.Invoke(ctx, "/replicate.ReplicateService/NotifyToCopy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicateServiceClient) Copying(ctx context.Context, in *CopyingRequest, opts ...grpc.CallOption) (ReplicateService_CopyingClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReplicateService_ServiceDesc.Streams[0], "/replicate.ReplicateService/Copying", opts...)
	if err != nil {
		return nil, err
	}
	x := &replicateServiceCopyingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReplicateService_CopyingClient interface {
	Recv() (*CopyingResponse, error)
	grpc.ClientStream
}

type replicateServiceCopyingClient struct {
	grpc.ClientStream
}

func (x *replicateServiceCopyingClient) Recv() (*CopyingResponse, error) {
	m := new(CopyingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *replicateServiceClient) ConfirmCopy(ctx context.Context, in *ConfirmCopyRequest, opts ...grpc.CallOption) (*ConfirmCopyResponse, error) {
	out := new(ConfirmCopyResponse)
	err := c.cc.Invoke(ctx, "/replicate.ReplicateService/ConfirmCopy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicateServiceServer is the server API for ReplicateService service.
// All implementations must embed UnimplementedReplicateServiceServer
// for forward compatibility
type ReplicateServiceServer interface {
	NotifyToCopy(context.Context, *NotifyToCopyRequest) (*NotifyToCopyResponse, error)
	Copying(*CopyingRequest, ReplicateService_CopyingServer) error
	ConfirmCopy(context.Context, *ConfirmCopyRequest) (*ConfirmCopyResponse, error)
	mustEmbedUnimplementedReplicateServiceServer()
}

// UnimplementedReplicateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReplicateServiceServer struct {
}

func (UnimplementedReplicateServiceServer) NotifyToCopy(context.Context, *NotifyToCopyRequest) (*NotifyToCopyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyToCopy not implemented")
}
func (UnimplementedReplicateServiceServer) Copying(*CopyingRequest, ReplicateService_CopyingServer) error {
	return status.Errorf(codes.Unimplemented, "method Copying not implemented")
}
func (UnimplementedReplicateServiceServer) ConfirmCopy(context.Context, *ConfirmCopyRequest) (*ConfirmCopyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmCopy not implemented")
}
func (UnimplementedReplicateServiceServer) mustEmbedUnimplementedReplicateServiceServer() {}

// UnsafeReplicateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicateServiceServer will
// result in compilation errors.
type UnsafeReplicateServiceServer interface {
	mustEmbedUnimplementedReplicateServiceServer()
}

func RegisterReplicateServiceServer(s grpc.ServiceRegistrar, srv ReplicateServiceServer) {
	s.RegisterService(&ReplicateService_ServiceDesc, srv)
}

func _ReplicateService_NotifyToCopy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyToCopyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicateServiceServer).NotifyToCopy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/replicate.ReplicateService/NotifyToCopy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicateServiceServer).NotifyToCopy(ctx, req.(*NotifyToCopyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplicateService_Copying_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CopyingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReplicateServiceServer).Copying(m, &replicateServiceCopyingServer{stream})
}

type ReplicateService_CopyingServer interface {
	Send(*CopyingResponse) error
	grpc.ServerStream
}

type replicateServiceCopyingServer struct {
	grpc.ServerStream
}

func (x *replicateServiceCopyingServer) Send(m *CopyingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ReplicateService_ConfirmCopy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmCopyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicateServiceServer).ConfirmCopy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/replicate.ReplicateService/ConfirmCopy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicateServiceServer).ConfirmCopy(ctx, req.(*ConfirmCopyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReplicateService_ServiceDesc is the grpc.ServiceDesc for ReplicateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReplicateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "replicate.ReplicateService",
	HandlerType: (*ReplicateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyToCopy",
			Handler:    _ReplicateService_NotifyToCopy_Handler,
		},
		{
			MethodName: "ConfirmCopy",
			Handler:    _ReplicateService_ConfirmCopy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Copying",
			Handler:       _ReplicateService_Copying_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "schema/replicate/replicate.proto",
}
