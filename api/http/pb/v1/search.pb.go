// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: v1/search.proto

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

type SearchUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Q string `protobuf:"bytes,1,opt,name=q,proto3" json:"q,omitempty" binding:"required" form:"q"`
}

func (x *SearchUsersRequest) Reset() {
	*x = SearchUsersRequest{}
	mi := &file_v1_search_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUsersRequest) ProtoMessage() {}

func (x *SearchUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_search_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUsersRequest.ProtoReflect.Descriptor instead.
func (*SearchUsersRequest) Descriptor() ([]byte, []int) {
	return file_v1_search_proto_rawDescGZIP(), []int{0}
}

func (x *SearchUsersRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

type SearchUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*SearchUserResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *SearchUserResponse) Reset() {
	*x = SearchUserResponse{}
	mi := &file_v1_search_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserResponse) ProtoMessage() {}

func (x *SearchUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_search_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserResponse.ProtoReflect.Descriptor instead.
func (*SearchUserResponse) Descriptor() ([]byte, []int) {
	return file_v1_search_proto_rawDescGZIP(), []int{1}
}

func (x *SearchUserResponse) GetItems() []*SearchUserResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type SearchGroupChatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty" binding:"required" form:"page"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" binding:"max=50" form:"name"`
}

func (x *SearchGroupChatsRequest) Reset() {
	*x = SearchGroupChatsRequest{}
	mi := &file_v1_search_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchGroupChatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchGroupChatsRequest) ProtoMessage() {}

func (x *SearchGroupChatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_search_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchGroupChatsRequest.ProtoReflect.Descriptor instead.
func (*SearchGroupChatsRequest) Descriptor() ([]byte, []int) {
	return file_v1_search_proto_rawDescGZIP(), []int{2}
}

func (x *SearchGroupChatsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchGroupChatsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SearchGroupChatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*SearchGroupChatsResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Next  bool                             `protobuf:"varint,2,opt,name=next,proto3" json:"next,omitempty"`
}

func (x *SearchGroupChatsResponse) Reset() {
	*x = SearchGroupChatsResponse{}
	mi := &file_v1_search_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchGroupChatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchGroupChatsResponse) ProtoMessage() {}

func (x *SearchGroupChatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_search_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchGroupChatsResponse.ProtoReflect.Descriptor instead.
func (*SearchGroupChatsResponse) Descriptor() ([]byte, []int) {
	return file_v1_search_proto_rawDescGZIP(), []int{3}
}

func (x *SearchGroupChatsResponse) GetItems() []*SearchGroupChatsResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *SearchGroupChatsResponse) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

type SearchUserResponse_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Avatar   string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Surname  string `protobuf:"bytes,5,opt,name=surname,proto3" json:"surname,omitempty"`
}

func (x *SearchUserResponse_Item) Reset() {
	*x = SearchUserResponse_Item{}
	mi := &file_v1_search_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchUserResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserResponse_Item) ProtoMessage() {}

func (x *SearchUserResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_v1_search_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserResponse_Item.ProtoReflect.Descriptor instead.
func (*SearchUserResponse_Item) Descriptor() ([]byte, []int) {
	return file_v1_search_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SearchUserResponse_Item) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SearchUserResponse_Item) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SearchUserResponse_Item) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *SearchUserResponse_Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchUserResponse_Item) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

type SearchGroupChatsResponse_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type        int32  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Avatar      string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Count       int32  `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
	MaxNum      int32  `protobuf:"varint,7,opt,name=max_num,json=maxNum,proto3" json:"max_num,omitempty"`
	IsMember    bool   `protobuf:"varint,8,opt,name=is_member,json=isMember,proto3" json:"is_member,omitempty"`
	CreatedAt   string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *SearchGroupChatsResponse_Item) Reset() {
	*x = SearchGroupChatsResponse_Item{}
	mi := &file_v1_search_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchGroupChatsResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchGroupChatsResponse_Item) ProtoMessage() {}

func (x *SearchGroupChatsResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_v1_search_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchGroupChatsResponse_Item.ProtoReflect.Descriptor instead.
func (*SearchGroupChatsResponse_Item) Descriptor() ([]byte, []int) {
	return file_v1_search_proto_rawDescGZIP(), []int{3, 0}
}

func (x *SearchGroupChatsResponse_Item) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SearchGroupChatsResponse_Item) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *SearchGroupChatsResponse_Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchGroupChatsResponse_Item) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *SearchGroupChatsResponse_Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SearchGroupChatsResponse_Item) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SearchGroupChatsResponse_Item) GetMaxNum() int32 {
	if x != nil {
		return x.MaxNum
	}
	return 0
}

func (x *SearchGroupChatsResponse_Item) GetIsMember() bool {
	if x != nil {
		return x.IsMember
	}
	return false
}

func (x *SearchGroupChatsResponse_Item) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_v1_search_proto protoreflect.FileDescriptor

var file_v1_search_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x12, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2e, 0x0a, 0x01, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0x9a, 0x84, 0x9e,
	0x03, 0x1b, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x71, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x01, 0x71,
	0x22, 0xc1, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x78, 0x0a, 0x04, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x37, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x23,
	0x9a, 0x84, 0x9e, 0x03, 0x1e, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x70, 0x61, 0x67, 0x65, 0x22,
	0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x22, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0x9a, 0x84, 0x9e, 0x03, 0x1c, 0x66, 0x6f,
	0x72, 0x6d, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x3a, 0x22, 0x6d, 0x61, 0x78, 0x3d, 0x35, 0x30, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xcd, 0x02, 0x0a, 0x18, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x1a, 0xe3, 0x01, 0x0a, 0x04, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6d,
	0x61, 0x78, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61,
	0x78, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x5f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_search_proto_rawDescOnce sync.Once
	file_v1_search_proto_rawDescData = file_v1_search_proto_rawDesc
)

func file_v1_search_proto_rawDescGZIP() []byte {
	file_v1_search_proto_rawDescOnce.Do(func() {
		file_v1_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_search_proto_rawDescData)
	})
	return file_v1_search_proto_rawDescData
}

var file_v1_search_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_search_proto_goTypes = []any{
	(*SearchUsersRequest)(nil),            // 0: v1.SearchUsersRequest
	(*SearchUserResponse)(nil),            // 1: v1.SearchUserResponse
	(*SearchGroupChatsRequest)(nil),       // 2: v1.SearchGroupChatsRequest
	(*SearchGroupChatsResponse)(nil),      // 3: v1.SearchGroupChatsResponse
	(*SearchUserResponse_Item)(nil),       // 4: v1.SearchUserResponse.Item
	(*SearchGroupChatsResponse_Item)(nil), // 5: v1.SearchGroupChatsResponse.Item
}
var file_v1_search_proto_depIdxs = []int32{
	4, // 0: v1.SearchUserResponse.items:type_name -> v1.SearchUserResponse.Item
	5, // 1: v1.SearchGroupChatsResponse.items:type_name -> v1.SearchGroupChatsResponse.Item
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_search_proto_init() }
func file_v1_search_proto_init() {
	if File_v1_search_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_search_proto_goTypes,
		DependencyIndexes: file_v1_search_proto_depIdxs,
		MessageInfos:      file_v1_search_proto_msgTypes,
	}.Build()
	File_v1_search_proto = out.File
	file_v1_search_proto_rawDesc = nil
	file_v1_search_proto_goTypes = nil
	file_v1_search_proto_depIdxs = nil
}
