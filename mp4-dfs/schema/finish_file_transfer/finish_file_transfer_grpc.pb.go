// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: schema/finish_file_transfer/finish_file_transfer.proto

package finish_file_transfer

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

// FinishFileTransferServiceClient is the client API for FinishFileTransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FinishFileTransferServiceClient interface {
	FinishFileUpload(ctx context.Context, in *FinishFileUploadRequest, opts ...grpc.CallOption) (*FinishFileUploadResponse, error)
}

type finishFileTransferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFinishFileTransferServiceClient(cc grpc.ClientConnInterface) FinishFileTransferServiceClient {
	return &finishFileTransferServiceClient{cc}
}

func (c *finishFileTransferServiceClient) FinishFileUpload(ctx context.Context, in *FinishFileUploadRequest, opts ...grpc.CallOption) (*FinishFileUploadResponse, error) {
	out := new(FinishFileUploadResponse)
	err := c.cc.Invoke(ctx, "/finish_file_transfer.FinishFileTransferService/FinishFileUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinishFileTransferServiceServer is the server API for FinishFileTransferService service.
// All implementations must embed UnimplementedFinishFileTransferServiceServer
// for forward compatibility
type FinishFileTransferServiceServer interface {
	FinishFileUpload(context.Context, *FinishFileUploadRequest) (*FinishFileUploadResponse, error)
	mustEmbedUnimplementedFinishFileTransferServiceServer()
}

// UnimplementedFinishFileTransferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFinishFileTransferServiceServer struct {
}

func (UnimplementedFinishFileTransferServiceServer) FinishFileUpload(context.Context, *FinishFileUploadRequest) (*FinishFileUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishFileUpload not implemented")
}
func (UnimplementedFinishFileTransferServiceServer) mustEmbedUnimplementedFinishFileTransferServiceServer() {
}

// UnsafeFinishFileTransferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinishFileTransferServiceServer will
// result in compilation errors.
type UnsafeFinishFileTransferServiceServer interface {
	mustEmbedUnimplementedFinishFileTransferServiceServer()
}

func RegisterFinishFileTransferServiceServer(s grpc.ServiceRegistrar, srv FinishFileTransferServiceServer) {
	s.RegisterService(&FinishFileTransferService_ServiceDesc, srv)
}

func _FinishFileTransferService_FinishFileUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishFileUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinishFileTransferServiceServer).FinishFileUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/finish_file_transfer.FinishFileTransferService/FinishFileUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinishFileTransferServiceServer).FinishFileUpload(ctx, req.(*FinishFileUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FinishFileTransferService_ServiceDesc is the grpc.ServiceDesc for FinishFileTransferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FinishFileTransferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "finish_file_transfer.FinishFileTransferService",
	HandlerType: (*FinishFileTransferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FinishFileUpload",
			Handler:    _FinishFileTransferService_FinishFileUpload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema/finish_file_transfer/finish_file_transfer.proto",
}
