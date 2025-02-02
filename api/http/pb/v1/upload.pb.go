// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v4.25.1
// source: v1/upload.proto

package v1_pb

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UploadAvatarResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Avatar        string                 `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadAvatarResponse) Reset() {
	*x = UploadAvatarResponse{}
	mi := &file_v1_upload_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadAvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAvatarResponse) ProtoMessage() {}

func (x *UploadAvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_upload_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAvatarResponse.ProtoReflect.Descriptor instead.
func (*UploadAvatarResponse) Descriptor() ([]byte, []int) {
	return file_v1_upload_proto_rawDescGZIP(), []int{0}
}

func (x *UploadAvatarResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type UploadImagerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadImagerRequest) Reset() {
	*x = UploadImagerRequest{}
	mi := &file_v1_upload_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadImagerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImagerRequest) ProtoMessage() {}

func (x *UploadImagerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_upload_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImagerRequest.ProtoReflect.Descriptor instead.
func (*UploadImagerRequest) Descriptor() ([]byte, []int) {
	return file_v1_upload_proto_rawDescGZIP(), []int{1}
}

type UploadImageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Src           string                 `protobuf:"bytes,1,opt,name=src,proto3" json:"src"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadImageResponse) Reset() {
	*x = UploadImageResponse{}
	mi := &file_v1_upload_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageResponse) ProtoMessage() {}

func (x *UploadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_upload_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageResponse.ProtoReflect.Descriptor instead.
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return file_v1_upload_proto_rawDescGZIP(), []int{2}
}

func (x *UploadImageResponse) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

type UploadInitiateMultipartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty" binding:"required" label:"file_name"`
	FileSize      int64                  `protobuf:"varint,2,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty" binding:"required" label:"file_size"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadInitiateMultipartRequest) Reset() {
	*x = UploadInitiateMultipartRequest{}
	mi := &file_v1_upload_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadInitiateMultipartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadInitiateMultipartRequest) ProtoMessage() {}

func (x *UploadInitiateMultipartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_upload_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadInitiateMultipartRequest.ProtoReflect.Descriptor instead.
func (*UploadInitiateMultipartRequest) Descriptor() ([]byte, []int) {
	return file_v1_upload_proto_rawDescGZIP(), []int{3}
}

func (x *UploadInitiateMultipartRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadInitiateMultipartRequest) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

type UploadInitiateMultipartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UploadId      string                 `protobuf:"bytes,1,opt,name=upload_id,json=uploadId,proto3" json:"upload_id"`
	ShardSize     int32                  `protobuf:"varint,2,opt,name=shard_size,json=shardSize,proto3" json:"shard_size"`
	ShardNum      int32                  `protobuf:"varint,3,opt,name=shard_num,json=shardNum,proto3" json:"shard_num"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadInitiateMultipartResponse) Reset() {
	*x = UploadInitiateMultipartResponse{}
	mi := &file_v1_upload_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadInitiateMultipartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadInitiateMultipartResponse) ProtoMessage() {}

func (x *UploadInitiateMultipartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_upload_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadInitiateMultipartResponse.ProtoReflect.Descriptor instead.
func (*UploadInitiateMultipartResponse) Descriptor() ([]byte, []int) {
	return file_v1_upload_proto_rawDescGZIP(), []int{4}
}

func (x *UploadInitiateMultipartResponse) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

func (x *UploadInitiateMultipartResponse) GetShardSize() int32 {
	if x != nil {
		return x.ShardSize
	}
	return 0
}

func (x *UploadInitiateMultipartResponse) GetShardNum() int32 {
	if x != nil {
		return x.ShardNum
	}
	return 0
}

type UploadMultipartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UploadId      string                 `protobuf:"bytes,1,opt,name=upload_id,json=uploadId,proto3" json:"upload_id,omitempty" binding:"required" form:"upload_id" label:"upload_id"`
	SplitIndex    int32                  `protobuf:"varint,2,opt,name=split_index,json=splitIndex,proto3" json:"split_index,omitempty" binding:"min=0" form:"split_index" label:"split_index"`
	SplitNum      int32                  `protobuf:"varint,3,opt,name=split_num,json=splitNum,proto3" json:"split_num,omitempty" binding:"required,min=1" form:"split_num" label:"split_num"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadMultipartRequest) Reset() {
	*x = UploadMultipartRequest{}
	mi := &file_v1_upload_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadMultipartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadMultipartRequest) ProtoMessage() {}

func (x *UploadMultipartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_upload_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadMultipartRequest.ProtoReflect.Descriptor instead.
func (*UploadMultipartRequest) Descriptor() ([]byte, []int) {
	return file_v1_upload_proto_rawDescGZIP(), []int{5}
}

func (x *UploadMultipartRequest) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

func (x *UploadMultipartRequest) GetSplitIndex() int32 {
	if x != nil {
		return x.SplitIndex
	}
	return 0
}

func (x *UploadMultipartRequest) GetSplitNum() int32 {
	if x != nil {
		return x.SplitNum
	}
	return 0
}

type UploadMultipartResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UploadId      string                 `protobuf:"bytes,1,opt,name=upload_id,json=uploadId,proto3" json:"upload_id"`
	IsMerge       bool                   `protobuf:"varint,2,opt,name=is_merge,json=isMerge,proto3" json:"is_merge"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadMultipartResponse) Reset() {
	*x = UploadMultipartResponse{}
	mi := &file_v1_upload_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadMultipartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadMultipartResponse) ProtoMessage() {}

func (x *UploadMultipartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_upload_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadMultipartResponse.ProtoReflect.Descriptor instead.
func (*UploadMultipartResponse) Descriptor() ([]byte, []int) {
	return file_v1_upload_proto_rawDescGZIP(), []int{6}
}

func (x *UploadMultipartResponse) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

func (x *UploadMultipartResponse) GetIsMerge() bool {
	if x != nil {
		return x.IsMerge
	}
	return false
}

var File_v1_upload_proto protoreflect.FileDescriptor

var file_v1_upload_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x14, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x12, 0x9a, 0x84, 0x9e, 0x03, 0x0d, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x15,
	0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x03,
	0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x9a, 0x84, 0x9e, 0x03, 0x0a,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x72, 0x63, 0x22, 0x52, 0x03, 0x73, 0x72, 0x63, 0x22,
	0xb0, 0x01, 0x0a, 0x1e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x46, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0x9a, 0x84, 0x9e, 0x03, 0x24, 0x62, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x20, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x3a, 0x22, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x29, 0x9a,
	0x84, 0x9e, 0x03, 0x24, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x20, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3a, 0x22, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x1f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0x9a, 0x84, 0x9e, 0x03, 0x10,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x22,
	0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0a, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x16,
	0x9a, 0x84, 0x9e, 0x03, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x52, 0x09, 0x73, 0x68, 0x61, 0x72, 0x64, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x32, 0x0a, 0x09, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x15, 0x9a, 0x84, 0x9e, 0x03, 0x10, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x22, 0x52, 0x08, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x4e, 0x75, 0x6d, 0x22, 0xae, 0x02, 0x0a, 0x16, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x57, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x3a, 0x9a, 0x84, 0x9e, 0x03, 0x35, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x20, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x3a, 0x22, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x22, 0x52,
	0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x5c, 0x0a, 0x0b, 0x73, 0x70, 0x6c,
	0x69, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x3b,
	0x9a, 0x84, 0x9e, 0x03, 0x36, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x73, 0x70, 0x6c, 0x69, 0x74,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a,
	0x22, 0x6d, 0x69, 0x6e, 0x3d, 0x30, 0x22, 0x20, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3a, 0x22, 0x73,
	0x70, 0x6c, 0x69, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x52, 0x0a, 0x73, 0x70, 0x6c,
	0x69, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x5d, 0x0a, 0x09, 0x73, 0x70, 0x6c, 0x69, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x40, 0x9a, 0x84, 0x9e, 0x03,
	0x3b, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x2c, 0x6d, 0x69, 0x6e, 0x3d, 0x31, 0x22, 0x20, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x3a, 0x22, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x22, 0x52, 0x08, 0x73, 0x70,
	0x6c, 0x69, 0x74, 0x4e, 0x75, 0x6d, 0x22, 0x7e, 0x0a, 0x17, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0x9a, 0x84, 0x9e, 0x03, 0x10, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x08, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x72, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x14, 0x9a, 0x84, 0x9e, 0x03, 0x0f, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x22, 0x52, 0x07, 0x69,
	0x73, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_v1_upload_proto_rawDescOnce sync.Once
	file_v1_upload_proto_rawDescData []byte
)

func file_v1_upload_proto_rawDescGZIP() []byte {
	file_v1_upload_proto_rawDescOnce.Do(func() {
		file_v1_upload_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_upload_proto_rawDesc), len(file_v1_upload_proto_rawDesc)))
	})
	return file_v1_upload_proto_rawDescData
}

var file_v1_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_upload_proto_goTypes = []any{
	(*UploadAvatarResponse)(nil),            // 0: v1.UploadAvatarResponse
	(*UploadImagerRequest)(nil),             // 1: v1.UploadImagerRequest
	(*UploadImageResponse)(nil),             // 2: v1.UploadImageResponse
	(*UploadInitiateMultipartRequest)(nil),  // 3: v1.UploadInitiateMultipartRequest
	(*UploadInitiateMultipartResponse)(nil), // 4: v1.UploadInitiateMultipartResponse
	(*UploadMultipartRequest)(nil),          // 5: v1.UploadMultipartRequest
	(*UploadMultipartResponse)(nil),         // 6: v1.UploadMultipartResponse
}
var file_v1_upload_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_upload_proto_init() }
func file_v1_upload_proto_init() {
	if File_v1_upload_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_upload_proto_rawDesc), len(file_v1_upload_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_upload_proto_goTypes,
		DependencyIndexes: file_v1_upload_proto_depIdxs,
		MessageInfos:      file_v1_upload_proto_msgTypes,
	}.Build()
	File_v1_upload_proto = out.File
	file_v1_upload_proto_goTypes = nil
	file_v1_upload_proto_depIdxs = nil
}
