// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.1
// source: v1/sticker.proto

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

type StickerListItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MediaId       int32                  `protobuf:"varint,1,opt,name=media_id,json=mediaId,proto3" json:"media_id"`
	Src           string                 `protobuf:"bytes,2,opt,name=src,proto3" json:"src"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	StickerId     int32                  `protobuf:"varint,1,opt,name=sticker_id,json=stickerId,proto3" json:"sticker_id,omitempty" binding:"required" label:"sticker_id"`
	Type          int32                  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty" binding:"required,oneof=1 2" label:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	StickerId     int32                  `protobuf:"varint,1,opt,name=sticker_id,json=stickerId,proto3" json:"sticker_id"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	List          []*StickerListItem     `protobuf:"bytes,4,rep,name=list,proto3" json:"list"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ids           string                 `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty" binding:"required,ids" form:"ids" label:"ids"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Items         []*StickerSysListResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state          protoimpl.MessageState            `protogen:"open.v1"`
	SysSticker     []*StickerListResponse_SysSticker `protobuf:"bytes,1,rep,name=sys_sticker,json=sysSticker,proto3" json:"sys_sticker"`
	CollectSticker []*StickerListItem                `protobuf:"bytes,2,rep,name=collect_sticker,json=collectSticker,proto3" json:"collect_sticker"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	MediaId       int32                  `protobuf:"varint,1,opt,name=media_id,json=mediaId,proto3" json:"media_id"`
	Src           string                 `protobuf:"bytes,2,opt,name=src,proto3" json:"src"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Icon          string                 `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon"`
	Status        int32                  `protobuf:"varint,4,opt,name=status,proto3" json:"status"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	StickerId     int32                  `protobuf:"varint,1,opt,name=sticker_id,json=stickerId,proto3" json:"sticker_id"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	List          []*StickerListItem     `protobuf:"bytes,4,rep,name=list,proto3" json:"list"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

const file_v1_sticker_proto_rawDesc = "" +
	"\n" +
	"\x10v1/sticker.proto\x12\x02v1\x1a\x13tagger/tagger.proto\"e\n" +
	"\x0fStickerListItem\x12/\n" +
	"\bmedia_id\x18\x01 \x01(\x05B\x14\x9a\x84\x9e\x03\x0fjson:\"media_id\"R\amediaId\x12!\n" +
	"\x03src\x18\x02 \x01(\tB\x0f\x9a\x84\x9e\x03\n" +
	"json:\"src\"R\x03src\"\xa8\x01\n" +
	"\x17StickerSetSystemRequest\x12I\n" +
	"\n" +
	"sticker_id\x18\x01 \x01(\x05B*\x9a\x84\x9e\x03%binding:\"required\" label:\"sticker_id\"R\tstickerId\x12B\n" +
	"\x04type\x18\x02 \x01(\x05B.\x9a\x84\x9e\x03)binding:\"required,oneof=1 2\" label:\"type\"R\x04type\"\xd5\x01\n" +
	"\x18StickerSetSystemResponse\x125\n" +
	"\n" +
	"sticker_id\x18\x01 \x01(\x05B\x16\x9a\x84\x9e\x03\x11json:\"sticker_id\"R\tstickerId\x12!\n" +
	"\x03url\x18\x02 \x01(\tB\x0f\x9a\x84\x9e\x03\n" +
	"json:\"url\"R\x03url\x12$\n" +
	"\x04name\x18\x03 \x01(\tB\x10\x9a\x84\x9e\x03\vjson:\"name\"R\x04name\x129\n" +
	"\x04list\x18\x04 \x03(\v2\x13.v1.StickerListItemB\x10\x9a\x84\x9e\x03\vjson:\"list\"R\x04list\"\\\n" +
	"\x14StickerDeleteRequest\x12D\n" +
	"\x03ids\x18\x01 \x01(\tB2\x9a\x84\x9e\x03-form:\"ids\" binding:\"required,ids\" label:\"ids\"R\x03ids\"\x83\x02\n" +
	"\x16StickerSysListResponse\x12H\n" +
	"\x05items\x18\x01 \x03(\v2\x1f.v1.StickerSysListResponse.ItemB\x11\x9a\x84\x9e\x03\fjson:\"items\"R\x05items\x1a\x9e\x01\n" +
	"\x04Item\x12\x1e\n" +
	"\x02id\x18\x01 \x01(\x05B\x0e\x9a\x84\x9e\x03\tjson:\"id\"R\x02id\x12$\n" +
	"\x04name\x18\x02 \x01(\tB\x10\x9a\x84\x9e\x03\vjson:\"name\"R\x04name\x12$\n" +
	"\x04icon\x18\x03 \x01(\tB\x10\x9a\x84\x9e\x03\vjson:\"icon\"R\x04icon\x12*\n" +
	"\x06status\x18\x04 \x01(\x05B\x12\x9a\x84\x9e\x03\rjson:\"status\"R\x06status\"\x98\x03\n" +
	"\x13StickerListResponse\x12\\\n" +
	"\vsys_sticker\x18\x01 \x03(\v2\".v1.StickerListResponse.SysStickerB\x17\x9a\x84\x9e\x03\x12json:\"sys_sticker\"R\n" +
	"sysSticker\x12Y\n" +
	"\x0fcollect_sticker\x18\x02 \x03(\v2\x13.v1.StickerListItemB\x1b\x9a\x84\x9e\x03\x16json:\"collect_sticker\"R\x0ecollectSticker\x1a\xc7\x01\n" +
	"\n" +
	"SysSticker\x125\n" +
	"\n" +
	"sticker_id\x18\x01 \x01(\x05B\x16\x9a\x84\x9e\x03\x11json:\"sticker_id\"R\tstickerId\x12!\n" +
	"\x03url\x18\x02 \x01(\tB\x0f\x9a\x84\x9e\x03\n" +
	"json:\"url\"R\x03url\x12$\n" +
	"\x04name\x18\x03 \x01(\tB\x10\x9a\x84\x9e\x03\vjson:\"name\"R\x04name\x129\n" +
	"\x04list\x18\x04 \x03(\v2\x13.v1.StickerListItemB\x10\x9a\x84\x9e\x03\vjson:\"list\"R\x04list\"k\n" +
	"\x15StickerUploadResponse\x12/\n" +
	"\bmedia_id\x18\x01 \x01(\x05B\x14\x9a\x84\x9e\x03\x0fjson:\"media_id\"R\amediaId\x12!\n" +
	"\x03src\x18\x02 \x01(\tB\x0f\x9a\x84\x9e\x03\n" +
	"json:\"src\"R\x03srcB\fZ\n" +
	"./v1;v1_pbb\x06proto3"

var (
	file_v1_sticker_proto_rawDescOnce sync.Once
	file_v1_sticker_proto_rawDescData []byte
)

func file_v1_sticker_proto_rawDescGZIP() []byte {
	file_v1_sticker_proto_rawDescOnce.Do(func() {
		file_v1_sticker_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_sticker_proto_rawDesc), len(file_v1_sticker_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_sticker_proto_rawDesc), len(file_v1_sticker_proto_rawDesc)),
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
	file_v1_sticker_proto_goTypes = nil
	file_v1_sticker_proto_depIdxs = nil
}
