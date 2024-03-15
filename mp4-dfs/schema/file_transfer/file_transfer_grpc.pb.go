// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: schema/file_transfer/file_transfer.proto

package file_transfer

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

// FileTransferServiceClient is the client API for FileTransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileTransferServiceClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileTransferService_UploadFileClient, error)
}

type fileTransferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileTransferServiceClient(cc grpc.ClientConnInterface) FileTransferServiceClient {
	return &fileTransferServiceClient{cc}
}

func (c *fileTransferServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileTransferService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileTransferService_ServiceDesc.Streams[0], "/file_transfer.FileTransferService/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileTransferServiceUploadFileClient{stream}
	return x, nil
}

type FileTransferService_UploadFileClient interface {
	Send(*UploadVideoRequest) error
	CloseAndRecv() (*UploadVideoResponse, error)
	grpc.ClientStream
}

type fileTransferServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *fileTransferServiceUploadFileClient) Send(m *UploadVideoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileTransferServiceUploadFileClient) CloseAndRecv() (*UploadVideoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadVideoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileTransferServiceServer is the server API for FileTransferService service.
// All implementations must embed UnimplementedFileTransferServiceServer
// for forward compatibility
type FileTransferServiceServer interface {
	UploadFile(FileTransferService_UploadFileServer) error
	mustEmbedUnimplementedFileTransferServiceServer()
}

// UnimplementedFileTransferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileTransferServiceServer struct {
}

func (UnimplementedFileTransferServiceServer) UploadFile(FileTransferService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFileTransferServiceServer) mustEmbedUnimplementedFileTransferServiceServer() {}

// UnsafeFileTransferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileTransferServiceServer will
// result in compilation errors.
type UnsafeFileTransferServiceServer interface {
	mustEmbedUnimplementedFileTransferServiceServer()
}

func RegisterFileTransferServiceServer(s grpc.ServiceRegistrar, srv FileTransferServiceServer) {
	s.RegisterService(&FileTransferService_ServiceDesc, srv)
}

func _FileTransferService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileTransferServiceServer).UploadFile(&fileTransferServiceUploadFileServer{stream})
}

type FileTransferService_UploadFileServer interface {
	SendAndClose(*UploadVideoResponse) error
	Recv() (*UploadVideoRequest, error)
	grpc.ServerStream
}

type fileTransferServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *fileTransferServiceUploadFileServer) SendAndClose(m *UploadVideoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileTransferServiceUploadFileServer) Recv() (*UploadVideoRequest, error) {
	m := new(UploadVideoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileTransferService_ServiceDesc is the grpc.ServiceDesc for FileTransferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileTransferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file_transfer.FileTransferService",
	HandlerType: (*FileTransferServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FileTransferService_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "schema/file_transfer/file_transfer.proto",
}