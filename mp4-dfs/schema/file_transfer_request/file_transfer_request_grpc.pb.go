// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: schema/file_transfer_request/file_transfer_request.proto

package file_transfer_request

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

// FileTransferRequestServiceClient is the client API for FileTransferRequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileTransferRequestServiceClient interface {
	UploadRequest(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error)
}

type fileTransferRequestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileTransferRequestServiceClient(cc grpc.ClientConnInterface) FileTransferRequestServiceClient {
	return &fileTransferRequestServiceClient{cc}
}

func (c *fileTransferRequestServiceClient) UploadRequest(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error) {
	out := new(UploadFileResponse)
	err := c.cc.Invoke(ctx, "/file_transfer_request.FileTransferRequestService/UploadRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileTransferRequestServiceServer is the server API for FileTransferRequestService service.
// All implementations must embed UnimplementedFileTransferRequestServiceServer
// for forward compatibility
type FileTransferRequestServiceServer interface {
	UploadRequest(context.Context, *UploadFileRequest) (*UploadFileResponse, error)
	mustEmbedUnimplementedFileTransferRequestServiceServer()
}

// UnimplementedFileTransferRequestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileTransferRequestServiceServer struct {
}

func (UnimplementedFileTransferRequestServiceServer) UploadRequest(context.Context, *UploadFileRequest) (*UploadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadRequest not implemented")
}
func (UnimplementedFileTransferRequestServiceServer) mustEmbedUnimplementedFileTransferRequestServiceServer() {
}

// UnsafeFileTransferRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileTransferRequestServiceServer will
// result in compilation errors.
type UnsafeFileTransferRequestServiceServer interface {
	mustEmbedUnimplementedFileTransferRequestServiceServer()
}

func RegisterFileTransferRequestServiceServer(s grpc.ServiceRegistrar, srv FileTransferRequestServiceServer) {
	s.RegisterService(&FileTransferRequestService_ServiceDesc, srv)
}

func _FileTransferRequestService_UploadRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileTransferRequestServiceServer).UploadRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file_transfer_request.FileTransferRequestService/UploadRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTransferRequestServiceServer).UploadRequest(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileTransferRequestService_ServiceDesc is the grpc.ServiceDesc for FileTransferRequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileTransferRequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file_transfer_request.FileTransferRequestService",
	HandlerType: (*FileTransferRequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadRequest",
			Handler:    _FileTransferRequestService_UploadRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema/file_transfer_request/file_transfer_request.proto",
}
