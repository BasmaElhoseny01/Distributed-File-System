// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: schema/upload/upload.proto

package upload

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

// RequestUpload
type RequestUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName     string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	ClientSocket string `protobuf:"bytes,2,opt,name=client_socket,json=clientSocket,proto3" json:"client_socket,omitempty"` // to be used in Notifcation
}

func (x *RequestUploadRequest) Reset() {
	*x = RequestUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_upload_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestUploadRequest) ProtoMessage() {}

func (x *RequestUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_upload_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestUploadRequest.ProtoReflect.Descriptor instead.
func (*RequestUploadRequest) Descriptor() ([]byte, []int) {
	return file_schema_upload_upload_proto_rawDescGZIP(), []int{0}
}

func (x *RequestUploadRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *RequestUploadRequest) GetClientSocket() string {
	if x != nil {
		return x.ClientSocket
	}
	return ""
}

type RequestUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeSocket string `protobuf:"bytes,1,opt,name=node_socket,json=nodeSocket,proto3" json:"node_socket,omitempty"`
}

func (x *RequestUploadResponse) Reset() {
	*x = RequestUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_upload_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestUploadResponse) ProtoMessage() {}

func (x *RequestUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_upload_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestUploadResponse.ProtoReflect.Descriptor instead.
func (*RequestUploadResponse) Descriptor() ([]byte, []int) {
	return file_schema_upload_upload_proto_rawDescGZIP(), []int{1}
}

func (x *RequestUploadResponse) GetNodeSocket() string {
	if x != nil {
		return x.NodeSocket
	}
	return ""
}

// UploadFile
type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*UploadFileRequest_FileInfo
	//	*UploadFileRequest_ChuckData
	Data isUploadFileRequest_Data `protobuf_oneof:"data"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_upload_upload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_upload_upload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_schema_upload_upload_proto_rawDescGZIP(), []int{2}
}

func (m *UploadFileRequest) GetData() isUploadFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadFileRequest) GetFileInfo() *FileInfo {
	if x, ok := x.GetData().(*UploadFileRequest_FileInfo); ok {
		return x.FileInfo
	}
	return nil
}

func (x *UploadFileRequest) GetChuckData() []byte {
	if x, ok := x.GetData().(*UploadFileRequest_ChuckData); ok {
		return x.ChuckData
	}
	return nil
}

type isUploadFileRequest_Data interface {
	isUploadFileRequest_Data()
}

type UploadFileRequest_FileInfo struct {
	FileInfo *FileInfo `protobuf:"bytes,1,opt,name=file_info,json=fileInfo,proto3,oneof"`
}

type UploadFileRequest_ChuckData struct {
	ChuckData []byte `protobuf:"bytes,2,opt,name=chuck_data,json=chuckData,proto3,oneof"`
}

func (*UploadFileRequest_FileInfo) isUploadFileRequest_Data() {}

func (*UploadFileRequest_ChuckData) isUploadFileRequest_Data() {}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_upload_upload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_schema_upload_upload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_schema_upload_upload_proto_rawDescGZIP(), []int{3}
}

func (x *FileInfo) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type UploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_upload_upload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_upload_upload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_schema_upload_upload_proto_rawDescGZIP(), []int{4}
}

// NotifyMaster
type NotifyMasterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId   string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"` //id for datanode known by the master
	FileName string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FilePath string `protobuf:"bytes,3,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"` // file path on the datanode
}

func (x *NotifyMasterRequest) Reset() {
	*x = NotifyMasterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_upload_upload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyMasterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyMasterRequest) ProtoMessage() {}

func (x *NotifyMasterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_upload_upload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyMasterRequest.ProtoReflect.Descriptor instead.
func (*NotifyMasterRequest) Descriptor() ([]byte, []int) {
	return file_schema_upload_upload_proto_rawDescGZIP(), []int{5}
}

func (x *NotifyMasterRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *NotifyMasterRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *NotifyMasterRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type NotifyMasterReponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyMasterReponse) Reset() {
	*x = NotifyMasterReponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_upload_upload_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyMasterReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyMasterReponse) ProtoMessage() {}

func (x *NotifyMasterReponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_upload_upload_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyMasterReponse.ProtoReflect.Descriptor instead.
func (*NotifyMasterReponse) Descriptor() ([]byte, []int) {
	return file_schema_upload_upload_proto_rawDescGZIP(), []int{6}
}

// NotifyClient
type NotifyClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *NotifyClientRequest) Reset() {
	*x = NotifyClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_upload_upload_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyClientRequest) ProtoMessage() {}

func (x *NotifyClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_upload_upload_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyClientRequest.ProtoReflect.Descriptor instead.
func (*NotifyClientRequest) Descriptor() ([]byte, []int) {
	return file_schema_upload_upload_proto_rawDescGZIP(), []int{7}
}

func (x *NotifyClientRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type NotifyClientReponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyClientReponse) Reset() {
	*x = NotifyClientReponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_upload_upload_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyClientReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyClientReponse) ProtoMessage() {}

func (x *NotifyClientReponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_upload_upload_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyClientReponse.ProtoReflect.Descriptor instead.
func (*NotifyClientReponse) Descriptor() ([]byte, []int) {
	return file_schema_upload_upload_proto_rawDescGZIP(), []int{8}
}

var File_schema_upload_upload_proto protoreflect.FileDescriptor

var file_schema_upload_upload_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x58, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x38,
	0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x6f,
	0x64, 0x65, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x6d, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f,
	0x0a, 0x0a, 0x63, 0x68, 0x75, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x75, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x27, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x68, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x15, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xb8, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x19, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x48, 0x0a, 0x0c, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0f, 0x5a,
	0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_upload_upload_proto_rawDescOnce sync.Once
	file_schema_upload_upload_proto_rawDescData = file_schema_upload_upload_proto_rawDesc
)

func file_schema_upload_upload_proto_rawDescGZIP() []byte {
	file_schema_upload_upload_proto_rawDescOnce.Do(func() {
		file_schema_upload_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_upload_upload_proto_rawDescData)
	})
	return file_schema_upload_upload_proto_rawDescData
}

var file_schema_upload_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_schema_upload_upload_proto_goTypes = []interface{}{
	(*RequestUploadRequest)(nil),  // 0: upload.RequestUploadRequest
	(*RequestUploadResponse)(nil), // 1: upload.RequestUploadResponse
	(*UploadFileRequest)(nil),     // 2: upload.UploadFileRequest
	(*FileInfo)(nil),              // 3: upload.FileInfo
	(*UploadFileResponse)(nil),    // 4: upload.UploadFileResponse
	(*NotifyMasterRequest)(nil),   // 5: upload.NotifyMasterRequest
	(*NotifyMasterReponse)(nil),   // 6: upload.NotifyMasterReponse
	(*NotifyClientRequest)(nil),   // 7: upload.NotifyClientRequest
	(*NotifyClientReponse)(nil),   // 8: upload.NotifyClientReponse
}
var file_schema_upload_upload_proto_depIdxs = []int32{
	3, // 0: upload.UploadFileRequest.file_info:type_name -> upload.FileInfo
	0, // 1: upload.UploadService.RequestUpload:input_type -> upload.RequestUploadRequest
	2, // 2: upload.UploadService.UploadFile:input_type -> upload.UploadFileRequest
	5, // 3: upload.UploadService.NotifyMaster:input_type -> upload.NotifyMasterRequest
	7, // 4: upload.UploadService.NotifyClient:input_type -> upload.NotifyClientRequest
	1, // 5: upload.UploadService.RequestUpload:output_type -> upload.RequestUploadResponse
	4, // 6: upload.UploadService.UploadFile:output_type -> upload.UploadFileResponse
	6, // 7: upload.UploadService.NotifyMaster:output_type -> upload.NotifyMasterReponse
	8, // 8: upload.UploadService.NotifyClient:output_type -> upload.NotifyClientReponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_schema_upload_upload_proto_init() }
func file_schema_upload_upload_proto_init() {
	if File_schema_upload_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_upload_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestUploadRequest); i {
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
		file_schema_upload_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestUploadResponse); i {
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
		file_schema_upload_upload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_schema_upload_upload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_schema_upload_upload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResponse); i {
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
		file_schema_upload_upload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyMasterRequest); i {
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
		file_schema_upload_upload_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyMasterReponse); i {
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
		file_schema_upload_upload_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyClientRequest); i {
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
		file_schema_upload_upload_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyClientReponse); i {
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
	file_schema_upload_upload_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*UploadFileRequest_FileInfo)(nil),
		(*UploadFileRequest_ChuckData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_upload_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schema_upload_upload_proto_goTypes,
		DependencyIndexes: file_schema_upload_upload_proto_depIdxs,
		MessageInfos:      file_schema_upload_upload_proto_msgTypes,
	}.Build()
	File_schema_upload_upload_proto = out.File
	file_schema_upload_upload_proto_rawDesc = nil
	file_schema_upload_upload_proto_goTypes = nil
	file_schema_upload_upload_proto_depIdxs = nil
}
