// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: schema/register/register.proto

package register

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

type DataKeeperRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *DataKeeperRegisterRequest) Reset() {
	*x = DataKeeperRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_register_register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataKeeperRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataKeeperRegisterRequest) ProtoMessage() {}

func (x *DataKeeperRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_register_register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataKeeperRegisterRequest.ProtoReflect.Descriptor instead.
func (*DataKeeperRegisterRequest) Descriptor() ([]byte, []int) {
	return file_schema_register_register_proto_rawDescGZIP(), []int{0}
}

func (x *DataKeeperRegisterRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *DataKeeperRegisterRequest) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type DataKeeperRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataKeeperId string `protobuf:"bytes,1,opt,name=data_keeper_id,json=dataKeeperId,proto3" json:"data_keeper_id,omitempty"`
	Success      string `protobuf:"bytes,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DataKeeperRegisterResponse) Reset() {
	*x = DataKeeperRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_register_register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataKeeperRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataKeeperRegisterResponse) ProtoMessage() {}

func (x *DataKeeperRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_register_register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataKeeperRegisterResponse.ProtoReflect.Descriptor instead.
func (*DataKeeperRegisterResponse) Descriptor() ([]byte, []int) {
	return file_schema_register_register_proto_rawDescGZIP(), []int{1}
}

func (x *DataKeeperRegisterResponse) GetDataKeeperId() string {
	if x != nil {
		return x.DataKeeperId
	}
	return ""
}

func (x *DataKeeperRegisterResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

type AlivePingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataKeeperId string `protobuf:"bytes,1,opt,name=data_keeper_id,json=dataKeeperId,proto3" json:"data_keeper_id,omitempty"`
}

func (x *AlivePingRequest) Reset() {
	*x = AlivePingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_register_register_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlivePingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlivePingRequest) ProtoMessage() {}

func (x *AlivePingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_register_register_proto_msgTypes[2]
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
	return file_schema_register_register_proto_rawDescGZIP(), []int{2}
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
		mi := &file_schema_register_register_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlivePingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlivePingResponse) ProtoMessage() {}

func (x *AlivePingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_register_register_proto_msgTypes[3]
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
	return file_schema_register_register_proto_rawDescGZIP(), []int{3}
}

func (x *AlivePingResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

var File_schema_register_register_proto protoreflect.FileDescriptor

var file_schema_register_register_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x19, 0x44, 0x61,
	0x74, 0x61, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x5c, 0x0a, 0x1a, 0x44,
	0x61, 0x74, 0x61, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x38, 0x0a, 0x10, 0x41, 0x6c, 0x69,
	0x76, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x32, 0x72, 0x0a, 0x19, 0x44, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x55, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x58, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42,
	0x65, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x41, 0x6c,
	0x69, 0x76, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x41,
	0x6c, 0x69, 0x76, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x11, 0x5a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_register_register_proto_rawDescOnce sync.Once
	file_schema_register_register_proto_rawDescData = file_schema_register_register_proto_rawDesc
)

func file_schema_register_register_proto_rawDescGZIP() []byte {
	file_schema_register_register_proto_rawDescOnce.Do(func() {
		file_schema_register_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_register_register_proto_rawDescData)
	})
	return file_schema_register_register_proto_rawDescData
}

var file_schema_register_register_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_schema_register_register_proto_goTypes = []interface{}{
	(*DataKeeperRegisterRequest)(nil),  // 0: register.DataKeeperRegisterRequest
	(*DataKeeperRegisterResponse)(nil), // 1: register.DataKeeperRegisterResponse
	(*AlivePingRequest)(nil),           // 2: register.AlivePingRequest
	(*AlivePingResponse)(nil),          // 3: register.AlivePingResponse
}
var file_schema_register_register_proto_depIdxs = []int32{
	0, // 0: register.DataKeeperRegisterService.Register:input_type -> register.DataKeeperRegisterRequest
	2, // 1: register.HeartBeatService.AlivePing:input_type -> register.AlivePingRequest
	1, // 2: register.DataKeeperRegisterService.Register:output_type -> register.DataKeeperRegisterResponse
	3, // 3: register.HeartBeatService.AlivePing:output_type -> register.AlivePingResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schema_register_register_proto_init() }
func file_schema_register_register_proto_init() {
	if File_schema_register_register_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_register_register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataKeeperRegisterRequest); i {
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
		file_schema_register_register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataKeeperRegisterResponse); i {
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
		file_schema_register_register_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_schema_register_register_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_schema_register_register_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_schema_register_register_proto_goTypes,
		DependencyIndexes: file_schema_register_register_proto_depIdxs,
		MessageInfos:      file_schema_register_register_proto_msgTypes,
	}.Build()
	File_schema_register_register_proto = out.File
	file_schema_register_register_proto_rawDesc = nil
	file_schema_register_register_proto_goTypes = nil
	file_schema_register_register_proto_depIdxs = nil
}
