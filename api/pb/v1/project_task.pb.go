// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: v1/project_task.proto

package api_v1

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

type ProjectTaskCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId   int64  `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" form:"project_id" binding:"required"`
	TypeId      int64  `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ProjectTaskCreateRequest) Reset() {
	*x = ProjectTaskCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_project_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectTaskCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectTaskCreateRequest) ProtoMessage() {}

func (x *ProjectTaskCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectTaskCreateRequest.ProtoReflect.Descriptor instead.
func (*ProjectTaskCreateRequest) Descriptor() ([]byte, []int) {
	return file_v1_project_task_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectTaskCreateRequest) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *ProjectTaskCreateRequest) GetTypeId() int64 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *ProjectTaskCreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProjectTaskCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ProjectTaskCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProjectTaskCreateResponse) Reset() {
	*x = ProjectTaskCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_project_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectTaskCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectTaskCreateResponse) ProtoMessage() {}

func (x *ProjectTaskCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectTaskCreateResponse.ProtoReflect.Descriptor instead.
func (*ProjectTaskCreateResponse) Descriptor() ([]byte, []int) {
	return file_v1_project_task_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectTaskCreateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProjectTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId int64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" form:"project_id" binding:"required"`
}

func (x *ProjectTaskRequest) Reset() {
	*x = ProjectTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_project_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectTaskRequest) ProtoMessage() {}

func (x *ProjectTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectTaskRequest.ProtoReflect.Descriptor instead.
func (*ProjectTaskRequest) Descriptor() ([]byte, []int) {
	return file_v1_project_task_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectTaskRequest) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type ProjectTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*ProjectTaskResponse_Categories `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *ProjectTaskResponse) Reset() {
	*x = ProjectTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_project_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectTaskResponse) ProtoMessage() {}

func (x *ProjectTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectTaskResponse.ProtoReflect.Descriptor instead.
func (*ProjectTaskResponse) Descriptor() ([]byte, []int) {
	return file_v1_project_task_proto_rawDescGZIP(), []int{3}
}

func (x *ProjectTaskResponse) GetCategories() []*ProjectTaskResponse_Categories {
	if x != nil {
		return x.Categories
	}
	return nil
}

type ProjectTaskMoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId int64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" form:"project_id" binding:"required"`
	TaskId    int64 `protobuf:"varint,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty" form:"task_id" binding:"required"`
	FromId    int64 `protobuf:"varint,3,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty" form:"from_id" binding:"required"`
	ToId      int64 `protobuf:"varint,4,opt,name=to_id,json=toId,proto3" json:"to_id,omitempty" form:"to_id" binding:"required"`
}

func (x *ProjectTaskMoveRequest) Reset() {
	*x = ProjectTaskMoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_project_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectTaskMoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectTaskMoveRequest) ProtoMessage() {}

func (x *ProjectTaskMoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectTaskMoveRequest.ProtoReflect.Descriptor instead.
func (*ProjectTaskMoveRequest) Descriptor() ([]byte, []int) {
	return file_v1_project_task_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectTaskMoveRequest) GetProjectId() int64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *ProjectTaskMoveRequest) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *ProjectTaskMoveRequest) GetFromId() int64 {
	if x != nil {
		return x.FromId
	}
	return 0
}

func (x *ProjectTaskMoveRequest) GetToId() int64 {
	if x != nil {
		return x.ToId
	}
	return 0
}

type ProjectTaskMoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProjectTaskMoveResponse) Reset() {
	*x = ProjectTaskMoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_project_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectTaskMoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectTaskMoveResponse) ProtoMessage() {}

func (x *ProjectTaskMoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectTaskMoveResponse.ProtoReflect.Descriptor instead.
func (*ProjectTaskMoveResponse) Descriptor() ([]byte, []int) {
	return file_v1_project_task_proto_rawDescGZIP(), []int{5}
}

type ProjectTaskDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId int64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty" form:"task_id" binding:"required"`
}

func (x *ProjectTaskDetailRequest) Reset() {
	*x = ProjectTaskDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_project_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectTaskDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectTaskDetailRequest) ProtoMessage() {}

func (x *ProjectTaskDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectTaskDetailRequest.ProtoReflect.Descriptor instead.
func (*ProjectTaskDetailRequest) Descriptor() ([]byte, []int) {
	return file_v1_project_task_proto_rawDescGZIP(), []int{6}
}

func (x *ProjectTaskDetailRequest) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type ProjectTaskDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ProjectTaskDetailResponse) Reset() {
	*x = ProjectTaskDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_project_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectTaskDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectTaskDetailResponse) ProtoMessage() {}

func (x *ProjectTaskDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectTaskDetailResponse.ProtoReflect.Descriptor instead.
func (*ProjectTaskDetailResponse) Descriptor() ([]byte, []int) {
	return file_v1_project_task_proto_rawDescGZIP(), []int{7}
}

func (x *ProjectTaskDetailResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProjectTaskDetailResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProjectTaskDetailResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ProjectTaskTypeNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId int64  `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty" form:"task_id" binding:"required"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" form:"name" binding:"required"`
}

func (x *ProjectTaskTypeNameRequest) Reset() {
	*x = ProjectTaskTypeNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_project_task_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectTaskTypeNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectTaskTypeNameRequest) ProtoMessage() {}

func (x *ProjectTaskTypeNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_task_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectTaskTypeNameRequest.ProtoReflect.Descriptor instead.
func (*ProjectTaskTypeNameRequest) Descriptor() ([]byte, []int) {
	return file_v1_project_task_proto_rawDescGZIP(), []int{8}
}

func (x *ProjectTaskTypeNameRequest) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *ProjectTaskTypeNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProjectTaskTypeNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProjectTaskTypeNameResponse) Reset() {
	*x = ProjectTaskTypeNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_project_task_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectTaskTypeNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectTaskTypeNameResponse) ProtoMessage() {}

func (x *ProjectTaskTypeNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_task_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectTaskTypeNameResponse.ProtoReflect.Descriptor instead.
func (*ProjectTaskTypeNameResponse) Descriptor() ([]byte, []int) {
	return file_v1_project_task_proto_rawDescGZIP(), []int{9}
}

type ProjectTaskResponse_Tasks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ProjectTaskResponse_Tasks) Reset() {
	*x = ProjectTaskResponse_Tasks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_project_task_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectTaskResponse_Tasks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectTaskResponse_Tasks) ProtoMessage() {}

func (x *ProjectTaskResponse_Tasks) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_task_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectTaskResponse_Tasks.ProtoReflect.Descriptor instead.
func (*ProjectTaskResponse_Tasks) Descriptor() ([]byte, []int) {
	return file_v1_project_task_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ProjectTaskResponse_Tasks) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectTaskResponse_Tasks) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ProjectTaskResponse_Categories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64                        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string                       `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Tasks []*ProjectTaskResponse_Tasks `protobuf:"bytes,3,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *ProjectTaskResponse_Categories) Reset() {
	*x = ProjectTaskResponse_Categories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_project_task_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectTaskResponse_Categories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectTaskResponse_Categories) ProtoMessage() {}

func (x *ProjectTaskResponse_Categories) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_task_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectTaskResponse_Categories.ProtoReflect.Descriptor instead.
func (*ProjectTaskResponse_Categories) Descriptor() ([]byte, []int) {
	return file_v1_project_task_proto_rawDescGZIP(), []int{3, 1}
}

func (x *ProjectTaskResponse_Categories) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectTaskResponse_Categories) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProjectTaskResponse_Categories) GetTasks() []*ProjectTaskResponse_Tasks {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_v1_project_task_proto protoreflect.FileDescriptor

var file_v1_project_task_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f,
	0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a,
	0x18, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x29, 0x9a,
	0x84, 0x9e, 0x03, 0x24, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x5e, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x29, 0x9a, 0x84, 0x9e,
	0x03, 0x24, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x22, 0xf1, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x2d, 0x0a,
	0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x1a, 0x67, 0x0a, 0x0a,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x33, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x05,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x9f, 0x02, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x48, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x29, 0x9a, 0x84, 0x9e, 0x03, 0x24, 0x66, 0x6f, 0x72, 0x6d, 0x3a,
	0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x62, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x07, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x26, 0x9a, 0x84, 0x9e,
	0x03, 0x21, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x22,
	0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x22, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x07, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x26, 0x9a, 0x84,
	0x9e, 0x03, 0x21, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x64,
	0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x22, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x05,
	0x74, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x24, 0x9a, 0x84, 0x9e,
	0x03, 0x1f, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x74, 0x6f, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x22, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x5b, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f,
	0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x26, 0x9a, 0x84, 0x9e, 0x03, 0x21, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x69, 0x64, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22,
	0x72, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x26, 0x9a, 0x84, 0x9e, 0x03, 0x21, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x23, 0x9a, 0x84, 0x9e, 0x03, 0x1e, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x1b,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_v1_project_task_proto_rawDescOnce sync.Once
	file_v1_project_task_proto_rawDescData = file_v1_project_task_proto_rawDesc
)

func file_v1_project_task_proto_rawDescGZIP() []byte {
	file_v1_project_task_proto_rawDescOnce.Do(func() {
		file_v1_project_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_project_task_proto_rawDescData)
	})
	return file_v1_project_task_proto_rawDescData
}

var file_v1_project_task_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_v1_project_task_proto_goTypes = []interface{}{
	(*ProjectTaskCreateRequest)(nil),       // 0: v1.ProjectTaskCreateRequest
	(*ProjectTaskCreateResponse)(nil),      // 1: v1.ProjectTaskCreateResponse
	(*ProjectTaskRequest)(nil),             // 2: v1.ProjectTaskRequest
	(*ProjectTaskResponse)(nil),            // 3: v1.ProjectTaskResponse
	(*ProjectTaskMoveRequest)(nil),         // 4: v1.ProjectTaskMoveRequest
	(*ProjectTaskMoveResponse)(nil),        // 5: v1.ProjectTaskMoveResponse
	(*ProjectTaskDetailRequest)(nil),       // 6: v1.ProjectTaskDetailRequest
	(*ProjectTaskDetailResponse)(nil),      // 7: v1.ProjectTaskDetailResponse
	(*ProjectTaskTypeNameRequest)(nil),     // 8: v1.ProjectTaskTypeNameRequest
	(*ProjectTaskTypeNameResponse)(nil),    // 9: v1.ProjectTaskTypeNameResponse
	(*ProjectTaskResponse_Tasks)(nil),      // 10: v1.ProjectTaskResponse.Tasks
	(*ProjectTaskResponse_Categories)(nil), // 11: v1.ProjectTaskResponse.Categories
}
var file_v1_project_task_proto_depIdxs = []int32{
	11, // 0: v1.ProjectTaskResponse.categories:type_name -> v1.ProjectTaskResponse.Categories
	10, // 1: v1.ProjectTaskResponse.Categories.tasks:type_name -> v1.ProjectTaskResponse.Tasks
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_v1_project_task_proto_init() }
func file_v1_project_task_proto_init() {
	if File_v1_project_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_project_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectTaskCreateRequest); i {
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
		file_v1_project_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectTaskCreateResponse); i {
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
		file_v1_project_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectTaskRequest); i {
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
		file_v1_project_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectTaskResponse); i {
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
		file_v1_project_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectTaskMoveRequest); i {
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
		file_v1_project_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectTaskMoveResponse); i {
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
		file_v1_project_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectTaskDetailRequest); i {
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
		file_v1_project_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectTaskDetailResponse); i {
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
		file_v1_project_task_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectTaskTypeNameRequest); i {
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
		file_v1_project_task_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectTaskTypeNameResponse); i {
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
		file_v1_project_task_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectTaskResponse_Tasks); i {
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
		file_v1_project_task_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectTaskResponse_Categories); i {
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
			RawDescriptor: file_v1_project_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_project_task_proto_goTypes,
		DependencyIndexes: file_v1_project_task_proto_depIdxs,
		MessageInfos:      file_v1_project_task_proto_msgTypes,
	}.Build()
	File_v1_project_task_proto = out.File
	file_v1_project_task_proto_rawDesc = nil
	file_v1_project_task_proto_goTypes = nil
	file_v1_project_task_proto_depIdxs = nil
}
