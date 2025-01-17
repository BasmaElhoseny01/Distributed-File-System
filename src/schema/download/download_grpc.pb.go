// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: schema/download/download.proto

package download

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
	DownloadService_Download_FullMethodName  = "/download.DownloadService/Download"
	DownloadService_GetServer_FullMethodName = "/download.DownloadService/GetServer"
)

// DownloadServiceClient is the client API for DownloadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DownloadServiceClient interface {
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (DownloadService_DownloadClient, error)
	GetServer(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadServerResponse, error)
}

type downloadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDownloadServiceClient(cc grpc.ClientConnInterface) DownloadServiceClient {
	return &downloadServiceClient{cc}
}

func (c *downloadServiceClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (DownloadService_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &DownloadService_ServiceDesc.Streams[0], DownloadService_Download_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &downloadServiceDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DownloadService_DownloadClient interface {
	Recv() (*DownloadResponse, error)
	grpc.ClientStream
}

type downloadServiceDownloadClient struct {
	grpc.ClientStream
}

func (x *downloadServiceDownloadClient) Recv() (*DownloadResponse, error) {
	m := new(DownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *downloadServiceClient) GetServer(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadServerResponse, error) {
	out := new(DownloadServerResponse)
	err := c.cc.Invoke(ctx, DownloadService_GetServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DownloadServiceServer is the server API for DownloadService service.
// All implementations must embed UnimplementedDownloadServiceServer
// for forward compatibility
type DownloadServiceServer interface {
	Download(*DownloadRequest, DownloadService_DownloadServer) error
	GetServer(context.Context, *DownloadRequest) (*DownloadServerResponse, error)
	mustEmbedUnimplementedDownloadServiceServer()
}

// UnimplementedDownloadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDownloadServiceServer struct {
}

func (UnimplementedDownloadServiceServer) Download(*DownloadRequest, DownloadService_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedDownloadServiceServer) GetServer(context.Context, *DownloadRequest) (*DownloadServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServer not implemented")
}
func (UnimplementedDownloadServiceServer) mustEmbedUnimplementedDownloadServiceServer() {}

// UnsafeDownloadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DownloadServiceServer will
// result in compilation errors.
type UnsafeDownloadServiceServer interface {
	mustEmbedUnimplementedDownloadServiceServer()
}

func RegisterDownloadServiceServer(s grpc.ServiceRegistrar, srv DownloadServiceServer) {
	s.RegisterService(&DownloadService_ServiceDesc, srv)
}

func _DownloadService_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DownloadServiceServer).Download(m, &downloadServiceDownloadServer{stream})
}

type DownloadService_DownloadServer interface {
	Send(*DownloadResponse) error
	grpc.ServerStream
}

type downloadServiceDownloadServer struct {
	grpc.ServerStream
}

func (x *downloadServiceDownloadServer) Send(m *DownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DownloadService_GetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadServiceServer).GetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DownloadService_GetServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadServiceServer).GetServer(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DownloadService_ServiceDesc is the grpc.ServiceDesc for DownloadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DownloadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "download.DownloadService",
	HandlerType: (*DownloadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServer",
			Handler:    _DownloadService_GetServer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Download",
			Handler:       _DownloadService_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "schema/download/download.proto",
}
