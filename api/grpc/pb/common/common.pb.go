// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.3
// source: common/common.proto

package common

import (
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

type EntityNotifySettings struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Time until notifications are disabled (Unix timestamp)
	// If mute_until = 0, notifications are enabled
	// If mute_until > current time, notifications are disabled until the specified time
	// If mute_until = 2147483647, notifications are disabled permanently
	MuteUntil int32 `protobuf:"varint,1,opt,name=mute_until,json=muteUntil,proto3" json:"mute_until,omitempty"`
	// Mute notifications sound flag
	// If silent = true, notifications arrive without sound
	// If silent = false, notifications arrive with sound
	Silent bool `protobuf:"varint,2,opt,name=silent,proto3" json:"silent,omitempty"`
	// Message preview display flag
	// If show_previews = true, the message text is displayed in the notification
	// If show_previews = false, only the sender's name is shown in the notification
	ShowPreviews  bool `protobuf:"varint,3,opt,name=show_previews,json=showPreviews,proto3" json:"show_previews,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntityNotifySettings) Reset() {
	*x = EntityNotifySettings{}
	mi := &file_common_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityNotifySettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityNotifySettings) ProtoMessage() {}

func (x *EntityNotifySettings) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityNotifySettings.ProtoReflect.Descriptor instead.
func (*EntityNotifySettings) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *EntityNotifySettings) GetMuteUntil() int32 {
	if x != nil {
		return x.MuteUntil
	}
	return 0
}

func (x *EntityNotifySettings) GetSilent() bool {
	if x != nil {
		return x.Silent
	}
	return false
}

func (x *EntityNotifySettings) GetShowPreviews() bool {
	if x != nil {
		return x.ShowPreviews
	}
	return false
}

type InputFile struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Parts         int32                  `protobuf:"varint,2,opt,name=parts,proto3" json:"parts,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InputFile) Reset() {
	*x = InputFile{}
	mi := &file_common_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InputFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputFile) ProtoMessage() {}

func (x *InputFile) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputFile.ProtoReflect.Descriptor instead.
func (*InputFile) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *InputFile) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InputFile) GetParts() int32 {
	if x != nil {
		return x.Parts
	}
	return 0
}

func (x *InputFile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type InputDocumentFileLocation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InputDocumentFileLocation) Reset() {
	*x = InputDocumentFileLocation{}
	mi := &file_common_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InputDocumentFileLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputDocumentFileLocation) ProtoMessage() {}

func (x *InputDocumentFileLocation) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputDocumentFileLocation.ProtoReflect.Descriptor instead.
func (*InputDocumentFileLocation) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *InputDocumentFileLocation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type InputPhoto struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Photo:
	//
	//	*InputPhoto_Empty
	//	*InputPhoto_File
	Photo         isInputPhoto_Photo `protobuf_oneof:"photo"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InputPhoto) Reset() {
	*x = InputPhoto{}
	mi := &file_common_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InputPhoto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputPhoto) ProtoMessage() {}

func (x *InputPhoto) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputPhoto.ProtoReflect.Descriptor instead.
func (*InputPhoto) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{3}
}

func (x *InputPhoto) GetPhoto() isInputPhoto_Photo {
	if x != nil {
		return x.Photo
	}
	return nil
}

func (x *InputPhoto) GetEmpty() *InputPhotoEmpty {
	if x != nil {
		if x, ok := x.Photo.(*InputPhoto_Empty); ok {
			return x.Empty
		}
	}
	return nil
}

func (x *InputPhoto) GetFile() *InputFile {
	if x != nil {
		if x, ok := x.Photo.(*InputPhoto_File); ok {
			return x.File
		}
	}
	return nil
}

type isInputPhoto_Photo interface {
	isInputPhoto_Photo()
}

type InputPhoto_Empty struct {
	Empty *InputPhotoEmpty `protobuf:"bytes,1,opt,name=empty,proto3,oneof"`
}

type InputPhoto_File struct {
	File *InputFile `protobuf:"bytes,2,opt,name=file,proto3,oneof"`
}

func (*InputPhoto_Empty) isInputPhoto_Photo() {}

func (*InputPhoto_File) isInputPhoto_Photo() {}

type InputPhotoEmpty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InputPhotoEmpty) Reset() {
	*x = InputPhotoEmpty{}
	mi := &file_common_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InputPhotoEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputPhotoEmpty) ProtoMessage() {}

func (x *InputPhotoEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputPhotoEmpty.ProtoReflect.Descriptor instead.
func (*InputPhotoEmpty) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{4}
}

type DocumentAttributeVideo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Duration      int32                  `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"` // Video duration in seconds
	Width         int32                  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`       // Video width in pixels
	Height        int32                  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`     // Video height in pixels
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DocumentAttributeVideo) Reset() {
	*x = DocumentAttributeVideo{}
	mi := &file_common_common_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DocumentAttributeVideo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentAttributeVideo) ProtoMessage() {}

func (x *DocumentAttributeVideo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentAttributeVideo.ProtoReflect.Descriptor instead.
func (*DocumentAttributeVideo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{5}
}

func (x *DocumentAttributeVideo) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *DocumentAttributeVideo) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *DocumentAttributeVideo) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type DocumentAttributeAudio struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Duration      int32                  `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"` // Audio file duration in seconds
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Performer     string                 `protobuf:"bytes,3,opt,name=performer,proto3" json:"performer,omitempty"`
	Voice         bool                   `protobuf:"varint,4,opt,name=voice,proto3" json:"voice,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DocumentAttributeAudio) Reset() {
	*x = DocumentAttributeAudio{}
	mi := &file_common_common_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DocumentAttributeAudio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentAttributeAudio) ProtoMessage() {}

func (x *DocumentAttributeAudio) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentAttributeAudio.ProtoReflect.Descriptor instead.
func (*DocumentAttributeAudio) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{6}
}

func (x *DocumentAttributeAudio) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *DocumentAttributeAudio) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DocumentAttributeAudio) GetPerformer() string {
	if x != nil {
		return x.Performer
	}
	return ""
}

func (x *DocumentAttributeAudio) GetVoice() bool {
	if x != nil {
		return x.Voice
	}
	return false
}

type DocumentAttributeFilename struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DocumentAttributeFilename) Reset() {
	*x = DocumentAttributeFilename{}
	mi := &file_common_common_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DocumentAttributeFilename) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentAttributeFilename) ProtoMessage() {}

func (x *DocumentAttributeFilename) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentAttributeFilename.ProtoReflect.Descriptor instead.
func (*DocumentAttributeFilename) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{7}
}

func (x *DocumentAttributeFilename) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x72, 0x0a,
	0x14, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x75, 0x74, 0x65, 0x5f, 0x75, 0x6e,
	0x74, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x75, 0x74, 0x65, 0x55,
	0x6e, 0x74, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x68, 0x6f, 0x77, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x77, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x22, 0x45, 0x0a, 0x09, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x61, 0x72, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x19, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6f, 0x0a, 0x0a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x05, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x07, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x62, 0x0a, 0x16, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x7e, 0x0a,
	0x16, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x72,
	0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x38, 0x0a,
	0x19, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x76, 0x6f, 0x6f, 0x2e, 0x73,
	0x75, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData []byte
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_common_proto_rawDesc), len(file_common_common_proto_rawDesc)))
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_common_proto_goTypes = []any{
	(*EntityNotifySettings)(nil),      // 0: common.EntityNotifySettings
	(*InputFile)(nil),                 // 1: common.InputFile
	(*InputDocumentFileLocation)(nil), // 2: common.InputDocumentFileLocation
	(*InputPhoto)(nil),                // 3: common.InputPhoto
	(*InputPhotoEmpty)(nil),           // 4: common.InputPhotoEmpty
	(*DocumentAttributeVideo)(nil),    // 5: common.DocumentAttributeVideo
	(*DocumentAttributeAudio)(nil),    // 6: common.DocumentAttributeAudio
	(*DocumentAttributeFilename)(nil), // 7: common.DocumentAttributeFilename
}
var file_common_common_proto_depIdxs = []int32{
	4, // 0: common.InputPhoto.empty:type_name -> common.InputPhotoEmpty
	1, // 1: common.InputPhoto.file:type_name -> common.InputFile
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	file_common_common_proto_msgTypes[3].OneofWrappers = []any{
		(*InputPhoto_Empty)(nil),
		(*InputPhoto_File)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_common_proto_rawDesc), len(file_common_common_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}
