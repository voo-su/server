// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.1
// source: v1/project.proto

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

type ProjectCreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title" form:"title" label:"title"`
	Ids           string                 `protobuf:"bytes,2,opt,name=ids,proto3" json:"ids" binding:"ids" form:"ids" label:"ids"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *ProjectCreateRequest) GetIds() string {
	if x != nil {
		return x.Ids
	}
	return ""
}

type ProjectCreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *ProjectCreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ProjectListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Items         []*ProjectListResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

type ProjectDetailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" binding:"required" form:"id" label:"id"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectDetailRequest) Reset() {
	*x = ProjectDetailRequest{}
	mi := &file_v1_project_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectDetailRequest) ProtoMessage() {}

func (x *ProjectDetailRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProjectDetailRequest.ProtoReflect.Descriptor instead.
func (*ProjectDetailRequest) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectDetailRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ProjectDetailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectDetailResponse) Reset() {
	*x = ProjectDetailResponse{}
	mi := &file_v1_project_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectDetailResponse) ProtoMessage() {}

func (x *ProjectDetailResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProjectDetailResponse.ProtoReflect.Descriptor instead.
func (*ProjectDetailResponse) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{5}
}

func (x *ProjectDetailResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProjectDetailResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProjectMembersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProjectId     string                 `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" binding:"required" form:"project_id" label:"project_id"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectMembersRequest) Reset() {
	*x = ProjectMembersRequest{}
	mi := &file_v1_project_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectMembersRequest) ProtoMessage() {}

func (x *ProjectMembersRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProjectMembersRequest.ProtoReflect.Descriptor instead.
func (*ProjectMembersRequest) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{6}
}

func (x *ProjectMembersRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type ProjectMembersResponse struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Items         []*ProjectMembersResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectMembersResponse) Reset() {
	*x = ProjectMembersResponse{}
	mi := &file_v1_project_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectMembersResponse) ProtoMessage() {}

func (x *ProjectMembersResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProjectMembersResponse.ProtoReflect.Descriptor instead.
func (*ProjectMembersResponse) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{7}
}

func (x *ProjectMembersResponse) GetItems() []*ProjectMembersResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type ProjectInviteFriendsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProjectId     string                 `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" form:"project_id" label:"project_id"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectInviteFriendsRequest) Reset() {
	*x = ProjectInviteFriendsRequest{}
	mi := &file_v1_project_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectInviteFriendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInviteFriendsRequest) ProtoMessage() {}

func (x *ProjectInviteFriendsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProjectInviteFriendsRequest.ProtoReflect.Descriptor instead.
func (*ProjectInviteFriendsRequest) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{8}
}

func (x *ProjectInviteFriendsRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type ProjectInviteFriendsResponse struct {
	state         protoimpl.MessageState               `protogen:"open.v1"`
	Items         []*ProjectInviteFriendsResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectInviteFriendsResponse) Reset() {
	*x = ProjectInviteFriendsResponse{}
	mi := &file_v1_project_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectInviteFriendsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInviteFriendsResponse) ProtoMessage() {}

func (x *ProjectInviteFriendsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ProjectInviteFriendsResponse.ProtoReflect.Descriptor instead.
func (*ProjectInviteFriendsResponse) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{9}
}

func (x *ProjectInviteFriendsResponse) GetItems() []*ProjectInviteFriendsResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type ProjectInviteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProjectId     string                 `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" binding:"required" label:"project_id"`
	Ids           string                 `protobuf:"bytes,2,opt,name=ids,proto3" json:"ids,omitempty" binding:"required,ids" label:"ids"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectInviteRequest) Reset() {
	*x = ProjectInviteRequest{}
	mi := &file_v1_project_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectInviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInviteRequest) ProtoMessage() {}

func (x *ProjectInviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[10]
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
	return file_v1_project_proto_rawDescGZIP(), []int{10}
}

func (x *ProjectInviteRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ProjectInviteRequest) GetIds() string {
	if x != nil {
		return x.Ids
	}
	return ""
}

type ProjectInviteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectInviteResponse) Reset() {
	*x = ProjectInviteResponse{}
	mi := &file_v1_project_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectInviteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInviteResponse) ProtoMessage() {}

func (x *ProjectInviteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[11]
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
	return file_v1_project_proto_rawDescGZIP(), []int{11}
}

type ProjectListResponse_Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectListResponse_Item) Reset() {
	*x = ProjectListResponse_Item{}
	mi := &file_v1_project_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectListResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectListResponse_Item) ProtoMessage() {}

func (x *ProjectListResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[12]
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

func (x *ProjectListResponse_Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProjectListResponse_Item) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ProjectMembersResponse_Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectMembersResponse_Item) Reset() {
	*x = ProjectMembersResponse_Item{}
	mi := &file_v1_project_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectMembersResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectMembersResponse_Item) ProtoMessage() {}

func (x *ProjectMembersResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[13]
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
	return file_v1_project_proto_rawDescGZIP(), []int{7, 0}
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

type ProjectInviteFriendsResponse_Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProjectInviteFriendsResponse_Item) Reset() {
	*x = ProjectInviteFriendsResponse_Item{}
	mi := &file_v1_project_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProjectInviteFriendsResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInviteFriendsResponse_Item) ProtoMessage() {}

func (x *ProjectInviteFriendsResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_v1_project_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectInviteFriendsResponse_Item.ProtoReflect.Descriptor instead.
func (*ProjectInviteFriendsResponse_Item) Descriptor() ([]byte, []int) {
	return file_v1_project_proto_rawDescGZIP(), []int{9, 0}
}

func (x *ProjectInviteFriendsResponse_Item) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProjectInviteFriendsResponse_Item) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_v1_project_proto protoreflect.FileDescriptor

const file_v1_project_proto_rawDesc = "" +
	"\n" +
	"\x10v1/project.proto\x12\x02v1\x1a\x13tagger/tagger.proto\"\xa2\x01\n" +
	"\x14ProjectCreateRequest\x12B\n" +
	"\x05title\x18\x01 \x01(\tB,\x9a\x84\x9e\x03'form:\"title\" json:\"title\" label:\"title\"R\x05title\x12F\n" +
	"\x03ids\x18\x02 \x01(\tB4\x9a\x84\x9e\x03/form:\"ids\" json:\"ids\" binding:\"ids\" label:\"ids\"R\x03ids\"7\n" +
	"\x15ProjectCreateResponse\x12\x1e\n" +
	"\x02id\x18\x01 \x01(\tB\x0e\x9a\x84\x9e\x03\tjson:\"id\"R\x02id\"\x14\n" +
	"\x12ProjectListRequest\"\xad\x01\n" +
	"\x13ProjectListResponse\x12E\n" +
	"\x05items\x18\x01 \x03(\v2\x1c.v1.ProjectListResponse.ItemB\x11\x9a\x84\x9e\x03\fjson:\"items\"R\x05items\x1aO\n" +
	"\x04Item\x12\x1e\n" +
	"\x02id\x18\x01 \x01(\tB\x0e\x9a\x84\x9e\x03\tjson:\"id\"R\x02id\x12'\n" +
	"\x05title\x18\x02 \x01(\tB\x11\x9a\x84\x9e\x03\fjson:\"title\"R\x05title\"T\n" +
	"\x14ProjectDetailRequest\x12<\n" +
	"\x02id\x18\x01 \x01(\tB,\x9a\x84\x9e\x03'form:\"id\" binding:\"required\" label:\"id\"R\x02id\"]\n" +
	"\x15ProjectDetailResponse\x12\x1e\n" +
	"\x02id\x18\x01 \x01(\tB\x0e\x9a\x84\x9e\x03\tjson:\"id\"R\x02id\x12$\n" +
	"\x04name\x18\x02 \x01(\tB\x10\x9a\x84\x9e\x03\vjson:\"name\"R\x04name\"t\n" +
	"\x15ProjectMembersRequest\x12[\n" +
	"\n" +
	"project_id\x18\x01 \x01(\tB<\x9a\x84\x9e\x037form:\"project_id\" binding:\"required\" label:\"project_id\"R\tprojectId\"\xbc\x01\n" +
	"\x16ProjectMembersResponse\x12H\n" +
	"\x05items\x18\x01 \x03(\v2\x1f.v1.ProjectMembersResponse.ItemB\x11\x9a\x84\x9e\x03\fjson:\"items\"R\x05items\x1aX\n" +
	"\x04Item\x12\x1e\n" +
	"\x02id\x18\x01 \x01(\x03B\x0e\x9a\x84\x9e\x03\tjson:\"id\"R\x02id\x120\n" +
	"\busername\x18\x02 \x01(\tB\x14\x9a\x84\x9e\x03\x0fjson:\"username\"R\busername\"g\n" +
	"\x1bProjectInviteFriendsRequest\x12H\n" +
	"\n" +
	"project_id\x18\x01 \x01(\tB)\x9a\x84\x9e\x03$form:\"project_id\" label:\"project_id\"R\tprojectId\"\xc8\x01\n" +
	"\x1cProjectInviteFriendsResponse\x12N\n" +
	"\x05items\x18\x01 \x03(\v2%.v1.ProjectInviteFriendsResponse.ItemB\x11\x9a\x84\x9e\x03\fjson:\"items\"R\x05items\x1aX\n" +
	"\x04Item\x12\x1e\n" +
	"\x02id\x18\x01 \x01(\x03B\x0e\x9a\x84\x9e\x03\tjson:\"id\"R\x02id\x120\n" +
	"\busername\x18\x02 \x01(\tB\x14\x9a\x84\x9e\x03\x0fjson:\"username\"R\busername\"\x9c\x01\n" +
	"\x14ProjectInviteRequest\x12I\n" +
	"\n" +
	"project_id\x18\x01 \x01(\tB*\x9a\x84\x9e\x03%binding:\"required\" label:\"project_id\"R\tprojectId\x129\n" +
	"\x03ids\x18\x02 \x01(\tB'\x9a\x84\x9e\x03\"binding:\"required,ids\" label:\"ids\"R\x03ids\"\x17\n" +
	"\x15ProjectInviteResponseB\fZ\n" +
	"./v1;v1_pbb\x06proto3"

var (
	file_v1_project_proto_rawDescOnce sync.Once
	file_v1_project_proto_rawDescData []byte
)

func file_v1_project_proto_rawDescGZIP() []byte {
	file_v1_project_proto_rawDescOnce.Do(func() {
		file_v1_project_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_project_proto_rawDesc), len(file_v1_project_proto_rawDesc)))
	})
	return file_v1_project_proto_rawDescData
}

var file_v1_project_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_v1_project_proto_goTypes = []any{
	(*ProjectCreateRequest)(nil),              // 0: v1.ProjectCreateRequest
	(*ProjectCreateResponse)(nil),             // 1: v1.ProjectCreateResponse
	(*ProjectListRequest)(nil),                // 2: v1.ProjectListRequest
	(*ProjectListResponse)(nil),               // 3: v1.ProjectListResponse
	(*ProjectDetailRequest)(nil),              // 4: v1.ProjectDetailRequest
	(*ProjectDetailResponse)(nil),             // 5: v1.ProjectDetailResponse
	(*ProjectMembersRequest)(nil),             // 6: v1.ProjectMembersRequest
	(*ProjectMembersResponse)(nil),            // 7: v1.ProjectMembersResponse
	(*ProjectInviteFriendsRequest)(nil),       // 8: v1.ProjectInviteFriendsRequest
	(*ProjectInviteFriendsResponse)(nil),      // 9: v1.ProjectInviteFriendsResponse
	(*ProjectInviteRequest)(nil),              // 10: v1.ProjectInviteRequest
	(*ProjectInviteResponse)(nil),             // 11: v1.ProjectInviteResponse
	(*ProjectListResponse_Item)(nil),          // 12: v1.ProjectListResponse.Item
	(*ProjectMembersResponse_Item)(nil),       // 13: v1.ProjectMembersResponse.Item
	(*ProjectInviteFriendsResponse_Item)(nil), // 14: v1.ProjectInviteFriendsResponse.Item
}
var file_v1_project_proto_depIdxs = []int32{
	12, // 0: v1.ProjectListResponse.items:type_name -> v1.ProjectListResponse.Item
	13, // 1: v1.ProjectMembersResponse.items:type_name -> v1.ProjectMembersResponse.Item
	14, // 2: v1.ProjectInviteFriendsResponse.items:type_name -> v1.ProjectInviteFriendsResponse.Item
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_project_proto_rawDesc), len(file_v1_project_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_project_proto_goTypes,
		DependencyIndexes: file_v1_project_proto_depIdxs,
		MessageInfos:      file_v1_project_proto_msgTypes,
	}.Build()
	File_v1_project_proto = out.File
	file_v1_project_proto_goTypes = nil
	file_v1_project_proto_depIdxs = nil
}
