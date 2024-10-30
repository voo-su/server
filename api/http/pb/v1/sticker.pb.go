// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: v1/sticker.proto

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

type StickerListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaId int32  `protobuf:"varint,1,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	Src     string `protobuf:"bytes,2,opt,name=src,proto3" json:"src,omitempty"`
}

func (x *StickerListItem) Reset() {
	*x = StickerListItem{}
	mi := &file_v1_sticker_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StickerListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StickerListItem) ProtoMessage() {}

func (x *StickerListItem) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sticker_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StickerListItem.ProtoReflect.Descriptor instead.
func (*StickerListItem) Descriptor() ([]byte, []int) {
	return file_v1_sticker_proto_rawDescGZIP(), []int{0}
}

func (x *StickerListItem) GetMediaId() int32 {
	if x != nil {
		return x.MediaId
	}
	return 0
}

func (x *StickerListItem) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

type StickerSetSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StickerId int32 `protobuf:"varint,1,opt,name=sticker_id,json=stickerId,proto3" json:"sticker_id,omitempty" binding:"required"`
	Type      int32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty" binding:"required,oneof=1 2"`
}

func (x *StickerSetSystemRequest) Reset() {
	*x = StickerSetSystemRequest{}
	mi := &file_v1_sticker_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StickerSetSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StickerSetSystemRequest) ProtoMessage() {}

func (x *StickerSetSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sticker_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StickerSetSystemRequest.ProtoReflect.Descriptor instead.
func (*StickerSetSystemRequest) Descriptor() ([]byte, []int) {
	return file_v1_sticker_proto_rawDescGZIP(), []int{1}
}

func (x *StickerSetSystemRequest) GetStickerId() int32 {
	if x != nil {
		return x.StickerId
	}
	return 0
}

func (x *StickerSetSystemRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type StickerSetSystemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StickerId int32              `protobuf:"varint,1,opt,name=sticker_id,json=stickerId,proto3" json:"sticker_id,omitempty"`
	Url       string             `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Name      string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	List      []*StickerListItem `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *StickerSetSystemResponse) Reset() {
	*x = StickerSetSystemResponse{}
	mi := &file_v1_sticker_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StickerSetSystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StickerSetSystemResponse) ProtoMessage() {}

func (x *StickerSetSystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sticker_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StickerSetSystemResponse.ProtoReflect.Descriptor instead.
func (*StickerSetSystemResponse) Descriptor() ([]byte, []int) {
	return file_v1_sticker_proto_rawDescGZIP(), []int{2}
}

func (x *StickerSetSystemResponse) GetStickerId() int32 {
	if x != nil {
		return x.StickerId
	}
	return 0
}

func (x *StickerSetSystemResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *StickerSetSystemResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StickerSetSystemResponse) GetList() []*StickerListItem {
	if x != nil {
		return x.List
	}
	return nil
}

type StickerDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids string `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty" binding:"required,ids" form:"ids"`
}

func (x *StickerDeleteRequest) Reset() {
	*x = StickerDeleteRequest{}
	mi := &file_v1_sticker_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StickerDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StickerDeleteRequest) ProtoMessage() {}

func (x *StickerDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sticker_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StickerDeleteRequest.ProtoReflect.Descriptor instead.
func (*StickerDeleteRequest) Descriptor() ([]byte, []int) {
	return file_v1_sticker_proto_rawDescGZIP(), []int{3}
}

func (x *StickerDeleteRequest) GetIds() string {
	if x != nil {
		return x.Ids
	}
	return ""
}

type StickerSysListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*StickerSysListResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *StickerSysListResponse) Reset() {
	*x = StickerSysListResponse{}
	mi := &file_v1_sticker_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StickerSysListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StickerSysListResponse) ProtoMessage() {}

func (x *StickerSysListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sticker_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StickerSysListResponse.ProtoReflect.Descriptor instead.
func (*StickerSysListResponse) Descriptor() ([]byte, []int) {
	return file_v1_sticker_proto_rawDescGZIP(), []int{4}
}

func (x *StickerSysListResponse) GetItems() []*StickerSysListResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type StickerListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SysSticker     []*StickerListResponse_SysSticker `protobuf:"bytes,1,rep,name=sys_sticker,json=sysSticker,proto3" json:"sys_sticker,omitempty"`
	CollectSticker []*StickerListItem                `protobuf:"bytes,2,rep,name=collect_sticker,json=collectSticker,proto3" json:"collect_sticker,omitempty"`
}

func (x *StickerListResponse) Reset() {
	*x = StickerListResponse{}
	mi := &file_v1_sticker_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StickerListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StickerListResponse) ProtoMessage() {}

func (x *StickerListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sticker_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StickerListResponse.ProtoReflect.Descriptor instead.
func (*StickerListResponse) Descriptor() ([]byte, []int) {
	return file_v1_sticker_proto_rawDescGZIP(), []int{5}
}

func (x *StickerListResponse) GetSysSticker() []*StickerListResponse_SysSticker {
	if x != nil {
		return x.SysSticker
	}
	return nil
}

func (x *StickerListResponse) GetCollectSticker() []*StickerListItem {
	if x != nil {
		return x.CollectSticker
	}
	return nil
}

type StickerUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaId int32  `protobuf:"varint,1,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	Src     string `protobuf:"bytes,2,opt,name=src,proto3" json:"src,omitempty"`
}

func (x *StickerUploadResponse) Reset() {
	*x = StickerUploadResponse{}
	mi := &file_v1_sticker_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StickerUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StickerUploadResponse) ProtoMessage() {}

func (x *StickerUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sticker_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StickerUploadResponse.ProtoReflect.Descriptor instead.
func (*StickerUploadResponse) Descriptor() ([]byte, []int) {
	return file_v1_sticker_proto_rawDescGZIP(), []int{6}
}

func (x *StickerUploadResponse) GetMediaId() int32 {
	if x != nil {
		return x.MediaId
	}
	return 0
}

func (x *StickerUploadResponse) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

type StickerSysListResponse_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon   string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Status int32  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StickerSysListResponse_Item) Reset() {
	*x = StickerSysListResponse_Item{}
	mi := &file_v1_sticker_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StickerSysListResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StickerSysListResponse_Item) ProtoMessage() {}

func (x *StickerSysListResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sticker_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StickerSysListResponse_Item.ProtoReflect.Descriptor instead.
func (*StickerSysListResponse_Item) Descriptor() ([]byte, []int) {
	return file_v1_sticker_proto_rawDescGZIP(), []int{4, 0}
}

func (x *StickerSysListResponse_Item) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StickerSysListResponse_Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StickerSysListResponse_Item) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *StickerSysListResponse_Item) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type StickerListResponse_SysSticker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StickerId int32              `protobuf:"varint,1,opt,name=sticker_id,json=stickerId,proto3" json:"sticker_id,omitempty"`
	Url       string             `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Name      string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	List      []*StickerListItem `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *StickerListResponse_SysSticker) Reset() {
	*x = StickerListResponse_SysSticker{}
	mi := &file_v1_sticker_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StickerListResponse_SysSticker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StickerListResponse_SysSticker) ProtoMessage() {}

func (x *StickerListResponse_SysSticker) ProtoReflect() protoreflect.Message {
	mi := &file_v1_sticker_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StickerListResponse_SysSticker.ProtoReflect.Descriptor instead.
func (*StickerListResponse_SysSticker) Descriptor() ([]byte, []int) {
	return file_v1_sticker_proto_rawDescGZIP(), []int{5, 0}
}

func (x *StickerListResponse_SysSticker) GetStickerId() int32 {
	if x != nil {
		return x.StickerId
	}
	return 0
}

func (x *StickerListResponse_SysSticker) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *StickerListResponse_SysSticker) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StickerListResponse_SysSticker) GetList() []*StickerListItem {
	if x != nil {
		return x.List
	}
	return nil
}

var File_v1_sticker_proto protoreflect.FileDescriptor

var file_v1_sticker_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0f, 0x53,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19,
	0x0a, 0x08, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x22, 0x88, 0x01, 0x0a, 0x17,
	0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x73, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a, 0x84, 0x9e,
	0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x22, 0x52, 0x09, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x21, 0x9a,
	0x84, 0x9e, 0x03, 0x1c, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x3d, 0x31, 0x20, 0x32, 0x22,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x18, 0x53, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x72, 0x53, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x50, 0x0a, 0x14, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0x9a, 0x84, 0x9e, 0x03, 0x21, 0x66, 0x6f, 0x72,
	0x6d, 0x3a, 0x22, 0x69, 0x64, 0x73, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a,
	0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x69, 0x64, 0x73, 0x22, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x53,
	0x79, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x79, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x56, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x94, 0x02,
	0x0a, 0x13, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x5f, 0x73, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x0a,
	0x73, 0x79, 0x73, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x1a, 0x7a, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x53,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x15, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_sticker_proto_rawDescOnce sync.Once
	file_v1_sticker_proto_rawDescData = file_v1_sticker_proto_rawDesc
)

func file_v1_sticker_proto_rawDescGZIP() []byte {
	file_v1_sticker_proto_rawDescOnce.Do(func() {
		file_v1_sticker_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_sticker_proto_rawDescData)
	})
	return file_v1_sticker_proto_rawDescData
}

var file_v1_sticker_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v1_sticker_proto_goTypes = []any{
	(*StickerListItem)(nil),                // 0: v1.StickerListItem
	(*StickerSetSystemRequest)(nil),        // 1: v1.StickerSetSystemRequest
	(*StickerSetSystemResponse)(nil),       // 2: v1.StickerSetSystemResponse
	(*StickerDeleteRequest)(nil),           // 3: v1.StickerDeleteRequest
	(*StickerSysListResponse)(nil),         // 4: v1.StickerSysListResponse
	(*StickerListResponse)(nil),            // 5: v1.StickerListResponse
	(*StickerUploadResponse)(nil),          // 6: v1.StickerUploadResponse
	(*StickerSysListResponse_Item)(nil),    // 7: v1.StickerSysListResponse.Item
	(*StickerListResponse_SysSticker)(nil), // 8: v1.StickerListResponse.SysSticker
}
var file_v1_sticker_proto_depIdxs = []int32{
	0, // 0: v1.StickerSetSystemResponse.list:type_name -> v1.StickerListItem
	7, // 1: v1.StickerSysListResponse.items:type_name -> v1.StickerSysListResponse.Item
	8, // 2: v1.StickerListResponse.sys_sticker:type_name -> v1.StickerListResponse.SysSticker
	0, // 3: v1.StickerListResponse.collect_sticker:type_name -> v1.StickerListItem
	0, // 4: v1.StickerListResponse.SysSticker.list:type_name -> v1.StickerListItem
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_sticker_proto_init() }
func file_v1_sticker_proto_init() {
	if File_v1_sticker_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_sticker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_sticker_proto_goTypes,
		DependencyIndexes: file_v1_sticker_proto_depIdxs,
		MessageInfos:      file_v1_sticker_proto_msgTypes,
	}.Build()
	File_v1_sticker_proto = out.File
	file_v1_sticker_proto_rawDesc = nil
	file_v1_sticker_proto_goTypes = nil
	file_v1_sticker_proto_depIdxs = nil
}
