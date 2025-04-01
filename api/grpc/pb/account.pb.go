// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.1
// source: account.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
	common "voo.su/api/grpc/pb/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	mi := &file_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

type GetProfileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Avatar        string                 `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Name          string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Surname       string                 `protobuf:"bytes,6,opt,name=surname,proto3" json:"surname,omitempty"`
	Gender        int32                  `protobuf:"varint,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthday      string                 `protobuf:"bytes,8,opt,name=birthday,proto3" json:"birthday,omitempty"`
	About         string                 `protobuf:"bytes,9,opt,name=about,proto3" json:"about,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProfileResponse) Reset() {
	*x = GetProfileResponse{}
	mi := &file_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponse) ProtoMessage() {}

func (x *GetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponse.ProtoReflect.Descriptor instead.
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

func (x *GetProfileResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetProfileResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetProfileResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetProfileResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *GetProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProfileResponse) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *GetProfileResponse) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *GetProfileResponse) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *GetProfileResponse) GetAbout() string {
	if x != nil {
		return x.About
	}
	return ""
}

type UpdateProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname       string                 `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	Gender        int32                  `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Birthday      string                 `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	About         string                 `protobuf:"bytes,5,opt,name=about,proto3" json:"about,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProfileRequest) Reset() {
	*x = UpdateProfileRequest{}
	mi := &file_account_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileRequest) ProtoMessage() {}

func (x *UpdateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProfileRequest) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *UpdateProfileRequest) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UpdateProfileRequest) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *UpdateProfileRequest) GetAbout() string {
	if x != nil {
		return x.About
	}
	return ""
}

type UpdateProfileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProfileResponse) Reset() {
	*x = UpdateProfileResponse{}
	mi := &file_account_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileResponse) ProtoMessage() {}

func (x *UpdateProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileResponse.ProtoReflect.Descriptor instead.
func (*UpdateProfileResponse) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProfileResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UpdateProfilePhotoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	File          *common.InputFile      `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProfilePhotoRequest) Reset() {
	*x = UpdateProfilePhotoRequest{}
	mi := &file_account_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProfilePhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfilePhotoRequest) ProtoMessage() {}

func (x *UpdateProfilePhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfilePhotoRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfilePhotoRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProfilePhotoRequest) GetFile() *common.InputFile {
	if x != nil {
		return x.File
	}
	return nil
}

type UpdateProfilePhotoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProfilePhotoResponse) Reset() {
	*x = UpdateProfilePhotoResponse{}
	mi := &file_account_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProfilePhotoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfilePhotoResponse) ProtoMessage() {}

func (x *UpdateProfilePhotoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfilePhotoResponse.ProtoReflect.Descriptor instead.
func (*UpdateProfilePhotoResponse) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateProfilePhotoResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type NotifyEntity struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Entity:
	//
	//	*NotifyEntity_Chats
	//	*NotifyEntity_Groups
	//	*NotifyEntity_Chat
	//	*NotifyEntity_Group
	Entity        isNotifyEntity_Entity `protobuf_oneof:"entity"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotifyEntity) Reset() {
	*x = NotifyEntity{}
	mi := &file_account_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyEntity) ProtoMessage() {}

func (x *NotifyEntity) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyEntity.ProtoReflect.Descriptor instead.
func (*NotifyEntity) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{6}
}

func (x *NotifyEntity) GetEntity() isNotifyEntity_Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *NotifyEntity) GetChats() *EntityChats {
	if x != nil {
		if x, ok := x.Entity.(*NotifyEntity_Chats); ok {
			return x.Chats
		}
	}
	return nil
}

func (x *NotifyEntity) GetGroups() *EntityGroups {
	if x != nil {
		if x, ok := x.Entity.(*NotifyEntity_Groups); ok {
			return x.Groups
		}
	}
	return nil
}

func (x *NotifyEntity) GetChat() *EntityChat {
	if x != nil {
		if x, ok := x.Entity.(*NotifyEntity_Chat); ok {
			return x.Chat
		}
	}
	return nil
}

func (x *NotifyEntity) GetGroup() *EntityGroup {
	if x != nil {
		if x, ok := x.Entity.(*NotifyEntity_Group); ok {
			return x.Group
		}
	}
	return nil
}

type isNotifyEntity_Entity interface {
	isNotifyEntity_Entity()
}

type NotifyEntity_Chats struct {
	Chats *EntityChats `protobuf:"bytes,2,opt,name=chats,proto3,oneof"`
}

type NotifyEntity_Groups struct {
	Groups *EntityGroups `protobuf:"bytes,1,opt,name=groups,proto3,oneof"`
}

type NotifyEntity_Chat struct {
	Chat *EntityChat `protobuf:"bytes,4,opt,name=chat,proto3,oneof"`
}

type NotifyEntity_Group struct {
	Group *EntityGroup `protobuf:"bytes,5,opt,name=group,proto3,oneof"`
}

func (*NotifyEntity_Chats) isNotifyEntity_Entity() {}

func (*NotifyEntity_Groups) isNotifyEntity_Entity() {}

func (*NotifyEntity_Chat) isNotifyEntity_Entity() {}

func (*NotifyEntity_Group) isNotifyEntity_Entity() {}

type EntityChats struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntityChats) Reset() {
	*x = EntityChats{}
	mi := &file_account_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityChats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityChats) ProtoMessage() {}

func (x *EntityChats) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityChats.ProtoReflect.Descriptor instead.
func (*EntityChats) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{7}
}

type EntityGroups struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntityGroups) Reset() {
	*x = EntityGroups{}
	mi := &file_account_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityGroups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityGroups) ProtoMessage() {}

func (x *EntityGroups) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityGroups.ProtoReflect.Descriptor instead.
func (*EntityGroups) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{8}
}

type EntityChat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChatId        int64                  `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntityChat) Reset() {
	*x = EntityChat{}
	mi := &file_account_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityChat) ProtoMessage() {}

func (x *EntityChat) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityChat.ProtoReflect.Descriptor instead.
func (*EntityChat) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{9}
}

func (x *EntityChat) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type EntityGroup struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GroupId       int64                  `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntityGroup) Reset() {
	*x = EntityGroup{}
	mi := &file_account_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityGroup) ProtoMessage() {}

func (x *EntityGroup) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityGroup.ProtoReflect.Descriptor instead.
func (*EntityGroup) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{10}
}

func (x *EntityGroup) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type GetNotifySettingsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Entity        *NotifyEntity          `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNotifySettingsRequest) Reset() {
	*x = GetNotifySettingsRequest{}
	mi := &file_account_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNotifySettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotifySettingsRequest) ProtoMessage() {}

func (x *GetNotifySettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotifySettingsRequest.ProtoReflect.Descriptor instead.
func (*GetNotifySettingsRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{11}
}

func (x *GetNotifySettingsRequest) GetEntity() *NotifyEntity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type GetNotifySettingsResponse struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Settings      *common.EntityNotifySettings `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNotifySettingsResponse) Reset() {
	*x = GetNotifySettingsResponse{}
	mi := &file_account_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNotifySettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotifySettingsResponse) ProtoMessage() {}

func (x *GetNotifySettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotifySettingsResponse.ProtoReflect.Descriptor instead.
func (*GetNotifySettingsResponse) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{12}
}

func (x *GetNotifySettingsResponse) GetSettings() *common.EntityNotifySettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type UpdateNotifySettingsRequest struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Entity        *NotifyEntity                `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Settings      *common.EntityNotifySettings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateNotifySettingsRequest) Reset() {
	*x = UpdateNotifySettingsRequest{}
	mi := &file_account_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateNotifySettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNotifySettingsRequest) ProtoMessage() {}

func (x *UpdateNotifySettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNotifySettingsRequest.ProtoReflect.Descriptor instead.
func (*UpdateNotifySettingsRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{13}
}

func (x *UpdateNotifySettingsRequest) GetEntity() *NotifyEntity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *UpdateNotifySettingsRequest) GetSettings() *common.EntityNotifySettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type UpdateNotifySettingsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateNotifySettingsResponse) Reset() {
	*x = UpdateNotifySettingsResponse{}
	mi := &file_account_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateNotifySettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNotifySettingsResponse) ProtoMessage() {}

func (x *UpdateNotifySettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNotifySettingsResponse.ProtoReflect.Descriptor instead.
func (*UpdateNotifySettingsResponse) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{14}
}

func (x *UpdateNotifySettingsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RegisterDeviceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TokenType     int32                  `protobuf:"varint,1,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterDeviceRequest) Reset() {
	*x = RegisterDeviceRequest{}
	mi := &file_account_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDeviceRequest) ProtoMessage() {}

func (x *RegisterDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDeviceRequest.ProtoReflect.Descriptor instead.
func (*RegisterDeviceRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{15}
}

func (x *RegisterDeviceRequest) GetTokenType() int32 {
	if x != nil {
		return x.TokenType
	}
	return 0
}

func (x *RegisterDeviceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RegisterDeviceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterDeviceResponse) Reset() {
	*x = RegisterDeviceResponse{}
	mi := &file_account_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDeviceResponse) ProtoMessage() {}

func (x *RegisterDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDeviceResponse.ProtoReflect.Descriptor instead.
func (*RegisterDeviceResponse) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{16}
}

func (x *RegisterDeviceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_account_proto protoreflect.FileDescriptor

const file_account_proto_rawDesc = "" +
	"\n" +
	"\raccount.proto\x12\aaccount\x1a\x13common/common.proto\"\x13\n" +
	"\x11GetProfileRequest\"\xe6\x01\n" +
	"\x12GetProfileResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1a\n" +
	"\busername\x18\x03 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x04 \x01(\tR\x05email\x12\x16\n" +
	"\x06avatar\x18\x02 \x01(\tR\x06avatar\x12\x12\n" +
	"\x04name\x18\x05 \x01(\tR\x04name\x12\x18\n" +
	"\asurname\x18\x06 \x01(\tR\asurname\x12\x16\n" +
	"\x06gender\x18\a \x01(\x05R\x06gender\x12\x1a\n" +
	"\bbirthday\x18\b \x01(\tR\bbirthday\x12\x14\n" +
	"\x05about\x18\t \x01(\tR\x05about\"\x8e\x01\n" +
	"\x14UpdateProfileRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x18\n" +
	"\asurname\x18\x02 \x01(\tR\asurname\x12\x16\n" +
	"\x06gender\x18\x03 \x01(\x05R\x06gender\x12\x1a\n" +
	"\bbirthday\x18\x04 \x01(\tR\bbirthday\x12\x14\n" +
	"\x05about\x18\x05 \x01(\tR\x05about\"1\n" +
	"\x15UpdateProfileResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"B\n" +
	"\x19UpdateProfilePhotoRequest\x12%\n" +
	"\x04file\x18\x01 \x01(\v2\x11.common.InputFileR\x04file\"6\n" +
	"\x1aUpdateProfilePhotoResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"\xd0\x01\n" +
	"\fNotifyEntity\x12,\n" +
	"\x05chats\x18\x02 \x01(\v2\x14.account.EntityChatsH\x00R\x05chats\x12/\n" +
	"\x06groups\x18\x01 \x01(\v2\x15.account.EntityGroupsH\x00R\x06groups\x12)\n" +
	"\x04chat\x18\x04 \x01(\v2\x13.account.EntityChatH\x00R\x04chat\x12,\n" +
	"\x05group\x18\x05 \x01(\v2\x14.account.EntityGroupH\x00R\x05groupB\b\n" +
	"\x06entity\"\r\n" +
	"\vEntityChats\"\x0e\n" +
	"\fEntityGroups\"%\n" +
	"\n" +
	"EntityChat\x12\x17\n" +
	"\achat_id\x18\x01 \x01(\x03R\x06chatId\"(\n" +
	"\vEntityGroup\x12\x19\n" +
	"\bgroup_id\x18\x01 \x01(\x03R\agroupId\"I\n" +
	"\x18GetNotifySettingsRequest\x12-\n" +
	"\x06entity\x18\x01 \x01(\v2\x15.account.NotifyEntityR\x06entity\"U\n" +
	"\x19GetNotifySettingsResponse\x128\n" +
	"\bsettings\x18\x01 \x01(\v2\x1c.common.EntityNotifySettingsR\bsettings\"\x86\x01\n" +
	"\x1bUpdateNotifySettingsRequest\x12-\n" +
	"\x06entity\x18\x01 \x01(\v2\x15.account.NotifyEntityR\x06entity\x128\n" +
	"\bsettings\x18\x02 \x01(\v2\x1c.common.EntityNotifySettingsR\bsettings\"8\n" +
	"\x1cUpdateNotifySettingsResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"L\n" +
	"\x15RegisterDeviceRequest\x12\x1d\n" +
	"\n" +
	"token_type\x18\x01 \x01(\x05R\ttokenType\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token\"2\n" +
	"\x16RegisterDeviceResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\x9a\x04\n" +
	"\x0eAccountService\x12E\n" +
	"\n" +
	"GetProfile\x12\x1a.account.GetProfileRequest\x1a\x1b.account.GetProfileResponse\x12N\n" +
	"\rUpdateProfile\x12\x1d.account.UpdateProfileRequest\x1a\x1e.account.UpdateProfileResponse\x12]\n" +
	"\x12UpdateProfilePhoto\x12\".account.UpdateProfilePhotoRequest\x1a#.account.UpdateProfilePhotoResponse\x12Z\n" +
	"\x11GetNotifySettings\x12!.account.GetNotifySettingsRequest\x1a\".account.GetNotifySettingsResponse\x12c\n" +
	"\x14UpdateNotifySettings\x12$.account.UpdateNotifySettingsRequest\x1a%.account.UpdateNotifySettingsResponse\x12Q\n" +
	"\x0eRegisterDevice\x12\x1e.account.RegisterDeviceRequest\x1a\x1f.account.RegisterDeviceResponseB\x1fZ\x1dvoo.su/api/grpc/pb/account;pbb\x06proto3"

var (
	file_account_proto_rawDescOnce sync.Once
	file_account_proto_rawDescData []byte
)

func file_account_proto_rawDescGZIP() []byte {
	file_account_proto_rawDescOnce.Do(func() {
		file_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_account_proto_rawDesc), len(file_account_proto_rawDesc)))
	})
	return file_account_proto_rawDescData
}

var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_account_proto_goTypes = []any{
	(*GetProfileRequest)(nil),            // 0: account.GetProfileRequest
	(*GetProfileResponse)(nil),           // 1: account.GetProfileResponse
	(*UpdateProfileRequest)(nil),         // 2: account.UpdateProfileRequest
	(*UpdateProfileResponse)(nil),        // 3: account.UpdateProfileResponse
	(*UpdateProfilePhotoRequest)(nil),    // 4: account.UpdateProfilePhotoRequest
	(*UpdateProfilePhotoResponse)(nil),   // 5: account.UpdateProfilePhotoResponse
	(*NotifyEntity)(nil),                 // 6: account.NotifyEntity
	(*EntityChats)(nil),                  // 7: account.EntityChats
	(*EntityGroups)(nil),                 // 8: account.EntityGroups
	(*EntityChat)(nil),                   // 9: account.EntityChat
	(*EntityGroup)(nil),                  // 10: account.EntityGroup
	(*GetNotifySettingsRequest)(nil),     // 11: account.GetNotifySettingsRequest
	(*GetNotifySettingsResponse)(nil),    // 12: account.GetNotifySettingsResponse
	(*UpdateNotifySettingsRequest)(nil),  // 13: account.UpdateNotifySettingsRequest
	(*UpdateNotifySettingsResponse)(nil), // 14: account.UpdateNotifySettingsResponse
	(*RegisterDeviceRequest)(nil),        // 15: account.RegisterDeviceRequest
	(*RegisterDeviceResponse)(nil),       // 16: account.RegisterDeviceResponse
	(*common.InputFile)(nil),             // 17: common.InputFile
	(*common.EntityNotifySettings)(nil),  // 18: common.EntityNotifySettings
}
var file_account_proto_depIdxs = []int32{
	17, // 0: account.UpdateProfilePhotoRequest.file:type_name -> common.InputFile
	7,  // 1: account.NotifyEntity.chats:type_name -> account.EntityChats
	8,  // 2: account.NotifyEntity.groups:type_name -> account.EntityGroups
	9,  // 3: account.NotifyEntity.chat:type_name -> account.EntityChat
	10, // 4: account.NotifyEntity.group:type_name -> account.EntityGroup
	6,  // 5: account.GetNotifySettingsRequest.entity:type_name -> account.NotifyEntity
	18, // 6: account.GetNotifySettingsResponse.settings:type_name -> common.EntityNotifySettings
	6,  // 7: account.UpdateNotifySettingsRequest.entity:type_name -> account.NotifyEntity
	18, // 8: account.UpdateNotifySettingsRequest.settings:type_name -> common.EntityNotifySettings
	0,  // 9: account.AccountService.GetProfile:input_type -> account.GetProfileRequest
	2,  // 10: account.AccountService.UpdateProfile:input_type -> account.UpdateProfileRequest
	4,  // 11: account.AccountService.UpdateProfilePhoto:input_type -> account.UpdateProfilePhotoRequest
	11, // 12: account.AccountService.GetNotifySettings:input_type -> account.GetNotifySettingsRequest
	13, // 13: account.AccountService.UpdateNotifySettings:input_type -> account.UpdateNotifySettingsRequest
	15, // 14: account.AccountService.RegisterDevice:input_type -> account.RegisterDeviceRequest
	1,  // 15: account.AccountService.GetProfile:output_type -> account.GetProfileResponse
	3,  // 16: account.AccountService.UpdateProfile:output_type -> account.UpdateProfileResponse
	5,  // 17: account.AccountService.UpdateProfilePhoto:output_type -> account.UpdateProfilePhotoResponse
	12, // 18: account.AccountService.GetNotifySettings:output_type -> account.GetNotifySettingsResponse
	14, // 19: account.AccountService.UpdateNotifySettings:output_type -> account.UpdateNotifySettingsResponse
	16, // 20: account.AccountService.RegisterDevice:output_type -> account.RegisterDeviceResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	file_account_proto_msgTypes[6].OneofWrappers = []any{
		(*NotifyEntity_Chats)(nil),
		(*NotifyEntity_Groups)(nil),
		(*NotifyEntity_Chat)(nil),
		(*NotifyEntity_Group)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_account_proto_rawDesc), len(file_account_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_proto_goTypes,
		DependencyIndexes: file_account_proto_depIdxs,
		MessageInfos:      file_account_proto_msgTypes,
	}.Build()
	File_account_proto = out.File
	file_account_proto_goTypes = nil
	file_account_proto_depIdxs = nil
}
