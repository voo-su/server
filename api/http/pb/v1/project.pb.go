// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: v1/project.proto

package v1_pb

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
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

type ProjectCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ProjectCreateRequest) Reset() {
	*x = ProjectCreateRequest{}
	mi := &file_v1_project_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectCreateRequest) ProtoMessage() {}

func (x *ProjectCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectCreateRequest.ProtoReflect.Descriptor instead.
func (*ProjectCreateRequest) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectCreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ProjectCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProjectCreateResponse) Reset() {
	*x = ProjectCreateResponse{}
	mi := &file_v1_project_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectCreateResponse) ProtoMessage() {}

func (x *ProjectCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectCreateResponse.ProtoReflect.Descriptor instead.
func (*ProjectCreateResponse) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectCreateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProjectListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProjectListRequest) Reset() {
	*x = ProjectListRequest{}
	mi := &file_v1_project_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectListRequest) ProtoMessage() {}

func (x *ProjectListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectListRequest.ProtoReflect.Descriptor instead.
func (*ProjectListRequest) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{2}
}

type ProjectListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ProjectListResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ProjectListResponse) Reset() {
	*x = ProjectListResponse{}
	mi := &file_v1_project_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectListResponse) ProtoMessage() {}

func (x *ProjectListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectListResponse.ProtoReflect.Descriptor instead.
func (*ProjectListResponse) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{3}
}

func (x *ProjectListResponse) GetItems() []*ProjectListResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type ProjectMembersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId int64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" binding:"required" form:"project_id"`
}

func (x *ProjectMembersRequest) Reset() {
	*x = ProjectMembersRequest{}
	mi := &file_v1_project_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectMembersRequest) ProtoMessage() {}

func (x *ProjectMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectMembersRequest.ProtoReflect.Descriptor instead.
func (*ProjectMembersRequest) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectMembersRequest) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type ProjectMembersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ProjectMembersResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ProjectMembersResponse) Reset() {
	*x = ProjectMembersResponse{}
	mi := &file_v1_project_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectMembersResponse) ProtoMessage() {}

func (x *ProjectMembersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectMembersResponse.ProtoReflect.Descriptor instead.
func (*ProjectMembersResponse) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{5}
}

func (x *ProjectMembersResponse) GetItems() []*ProjectMembersResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type ProjectInviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId int32  `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" binding:"required"`
	Ids       string `protobuf:"bytes,2,opt,name=ids,proto3" json:"ids,omitempty" binding:"required,ids"`
}

func (x *ProjectInviteRequest) Reset() {
	*x = ProjectInviteRequest{}
	mi := &file_v1_project_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectInviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInviteRequest) ProtoMessage() {}

func (x *ProjectInviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectInviteRequest.ProtoReflect.Descriptor instead.
func (*ProjectInviteRequest) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{6}
}

func (x *ProjectInviteRequest) GetProjectId() int32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *ProjectInviteRequest) GetIds() string {
	if x != nil {
		return x.Ids
	}
	return ""
}

type ProjectInviteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProjectInviteResponse) Reset() {
	*x = ProjectInviteResponse{}
	mi := &file_v1_project_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectInviteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInviteResponse) ProtoMessage() {}

func (x *ProjectInviteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectInviteResponse.ProtoReflect.Descriptor instead.
func (*ProjectInviteResponse) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{7}
}

type ProjectListResponse_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ProjectListResponse_Item) Reset() {
	*x = ProjectListResponse_Item{}
	mi := &file_v1_project_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectListResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectListResponse_Item) ProtoMessage() {}

func (x *ProjectListResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectListResponse_Item.ProtoReflect.Descriptor instead.
func (*ProjectListResponse_Item) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ProjectListResponse_Item) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectListResponse_Item) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ProjectMembersResponse_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ProjectMembersResponse_Item) Reset() {
	*x = ProjectMembersResponse_Item{}
	mi := &file_v1_project_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectMembersResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectMembersResponse_Item) ProtoMessage() {}

func (x *ProjectMembersResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectMembersResponse_Item.ProtoReflect.Descriptor instead.
func (*ProjectMembersResponse_Item) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ProjectMembersResponse_Item) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectMembersResponse_Item) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_v1_project_proto protoreflect.FileDescriptor

var file_v1_project_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x14, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x77, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x1a, 0x2c, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0x61, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x29,
	0x9a, 0x84, 0x9e, 0x03, 0x24, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x32, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7d, 0x0a, 0x14, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0x9a, 0x84, 0x9e, 0x03, 0x16, 0x62, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c,
	0x69, 0x64, 0x73, 0x22, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x5f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_project_proto_rawDescOnce sync.Once
	file_v1_project_proto_rawDescData = file_v1_project_proto_rawDesc
)

func file_v1_project_proto_rawDescGZIP() []byte {
	file_v1_project_proto_rawDescOnce.Do(func() {
		file_v1_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_project_proto_rawDescData)
	})
	return file_v1_project_proto_rawDescData
}

var file_v1_project_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_v1_project_proto_goTypes = []any{
	(*ProjectCreateRequest)(nil),        // 0: v1.ProjectCreateRequest
	(*ProjectCreateResponse)(nil),       // 1: v1.ProjectCreateResponse
	(*ProjectListRequest)(nil),          // 2: v1.ProjectListRequest
	(*ProjectListResponse)(nil),         // 3: v1.ProjectListResponse
	(*ProjectMembersRequest)(nil),       // 4: v1.ProjectMembersRequest
	(*ProjectMembersResponse)(nil),      // 5: v1.ProjectMembersResponse
	(*ProjectInviteRequest)(nil),        // 6: v1.ProjectInviteRequest
	(*ProjectInviteResponse)(nil),       // 7: v1.ProjectInviteResponse
	(*ProjectListResponse_Item)(nil),    // 8: v1.ProjectListResponse.Item
	(*ProjectMembersResponse_Item)(nil), // 9: v1.ProjectMembersResponse.Item
}
var file_v1_project_proto_depIdxs = []int32{
	8, // 0: v1.ProjectListResponse.items:type_name -> v1.ProjectListResponse.Item
	9, // 1: v1.ProjectMembersResponse.items:type_name -> v1.ProjectMembersResponse.Item
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_project_proto_init() }
func file_v1_project_proto_init() {
	if File_v1_project_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_project_proto_goTypes,
		DependencyIndexes: file_v1_project_proto_depIdxs,
		MessageInfos:      file_v1_project_proto_msgTypes,
	}.Build()
	File_v1_project_proto = out.File
	file_v1_project_proto_rawDesc = nil
	file_v1_project_proto_goTypes = nil
	file_v1_project_proto_depIdxs = nil
}
