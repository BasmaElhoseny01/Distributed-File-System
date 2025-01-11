// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: schema/heart_beat/heart_beat.proto

package heart_beat

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AlivePingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataKeeperId string `protobuf:"bytes,1,opt,name=data_keeper_id,json=dataKeeperId,proto3" json:"data_keeper_id,omitempty"`
}

func (x *AlivePingRequest) Reset() {
	*x = AlivePingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_heart_beat_heart_beat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlivePingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlivePingRequest) ProtoMessage() {}

func (x *AlivePingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_heart_beat_heart_beat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlivePingRequest.ProtoReflect.Descriptor instead.
func (*AlivePingRequest) Descriptor() ([]byte, []int) {
	return file_schema_heart_beat_heart_beat_proto_rawDescGZIP(), []int{0}
}

func (x *AlivePingRequest) GetDataKeeperId() string {
	if x != nil {
		return x.DataKeeperId
	}
	return ""
}

type AlivePingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success string `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AlivePingResponse) Reset() {
	*x = AlivePingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_heart_beat_heart_beat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlivePingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlivePingResponse) ProtoMessage() {}

func (x *AlivePingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_heart_beat_heart_beat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlivePingResponse.ProtoReflect.Descriptor instead.
func (*AlivePingResponse) Descriptor() ([]byte, []int) {
	return file_schema_heart_beat_heart_beat_proto_rawDescGZIP(), []int{1}
}

func (x *AlivePingResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

var File_schema_heart_beat_heart_beat_proto protoreflect.FileDescriptor

var file_schema_heart_beat_heart_beat_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x5f, 0x62,
	0x65, 0x61, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68, 0x65, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x65, 0x61, 0x74,
	0x22, 0x38, 0x0a, 0x10, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61,
	0x74, 0x61, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x41, 0x6c,
	0x69, 0x76, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x5c, 0x0a, 0x10, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a,
	0x09, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x2e, 0x68, 0x65, 0x61,
	0x72, 0x74, 0x5f, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x68, 0x65, 0x61, 0x72, 0x74,
	0x5f, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x65, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_heart_beat_heart_beat_proto_rawDescOnce sync.Once
	file_schema_heart_beat_heart_beat_proto_rawDescData = file_schema_heart_beat_heart_beat_proto_rawDesc
)

func file_schema_heart_beat_heart_beat_proto_rawDescGZIP() []byte {
	file_schema_heart_beat_heart_beat_proto_rawDescOnce.Do(func() {
		file_schema_heart_beat_heart_beat_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_heart_beat_heart_beat_proto_rawDescData)
	})
	return file_schema_heart_beat_heart_beat_proto_rawDescData
}

var file_schema_heart_beat_heart_beat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_schema_heart_beat_heart_beat_proto_goTypes = []interface{}{
	(*AlivePingRequest)(nil),  // 0: heart_beat.AlivePingRequest
	(*AlivePingResponse)(nil), // 1: heart_beat.AlivePingResponse
}
var file_schema_heart_beat_heart_beat_proto_depIdxs = []int32{
	0, // 0: heart_beat.HeartBeatService.AlivePing:input_type -> heart_beat.AlivePingRequest
	1, // 1: heart_beat.HeartBeatService.AlivePing:output_type -> heart_beat.AlivePingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schema_heart_beat_heart_beat_proto_init() }
func file_schema_heart_beat_heart_beat_proto_init() {
	if File_schema_heart_beat_heart_beat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_heart_beat_heart_beat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlivePingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_heart_beat_heart_beat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlivePingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_heart_beat_heart_beat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schema_heart_beat_heart_beat_proto_goTypes,
		DependencyIndexes: file_schema_heart_beat_heart_beat_proto_depIdxs,
		MessageInfos:      file_schema_heart_beat_heart_beat_proto_msgTypes,
	}.Build()
	File_schema_heart_beat_heart_beat_proto = out.File
	file_schema_heart_beat_heart_beat_proto_rawDesc = nil
	file_schema_heart_beat_heart_beat_proto_goTypes = nil
	file_schema_heart_beat_heart_beat_proto_depIdxs = nil
}
