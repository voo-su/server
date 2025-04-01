// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.1
// source: bot/message.proto

package bot_pb

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

type MessageSendRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChatId        int32                  `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty" binding:"required" form:"chat_id" label:"chat_id"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty" binding:"required" form:"text" label:"text"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageSendRequest) Reset() {
	*x = MessageSendRequest{}
	mi := &file_bot_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSendRequest) ProtoMessage() {}

func (x *MessageSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bot_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSendRequest.ProtoReflect.Descriptor instead.
func (*MessageSendRequest) Descriptor() ([]byte, []int) {
	return file_bot_message_proto_rawDescGZIP(), []int{0}
}

func (x *MessageSendRequest) GetChatId() int32 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *MessageSendRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type MessageSendResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageSendResponse) Reset() {
	*x = MessageSendResponse{}
	mi := &file_bot_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSendResponse) ProtoMessage() {}

func (x *MessageSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bot_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSendResponse.ProtoReflect.Descriptor instead.
func (*MessageSendResponse) Descriptor() ([]byte, []int) {
	return file_bot_message_proto_rawDescGZIP(), []int{1}
}

type MessageChatsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageChatsRequest) Reset() {
	*x = MessageChatsRequest{}
	mi := &file_bot_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageChatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageChatsRequest) ProtoMessage() {}

func (x *MessageChatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bot_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageChatsRequest.ProtoReflect.Descriptor instead.
func (*MessageChatsRequest) Descriptor() ([]byte, []int) {
	return file_bot_message_proto_rawDescGZIP(), []int{2}
}

type MessageChatsResponse struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Items         []*MessageChatsResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageChatsResponse) Reset() {
	*x = MessageChatsResponse{}
	mi := &file_bot_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageChatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageChatsResponse) ProtoMessage() {}

func (x *MessageChatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bot_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageChatsResponse.ProtoReflect.Descriptor instead.
func (*MessageChatsResponse) Descriptor() ([]byte, []int) {
	return file_bot_message_proto_rawDescGZIP(), []int{3}
}

func (x *MessageChatsResponse) GetItems() []*MessageChatsResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type MessageChatsResponse_Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageChatsResponse_Item) Reset() {
	*x = MessageChatsResponse_Item{}
	mi := &file_bot_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageChatsResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageChatsResponse_Item) ProtoMessage() {}

func (x *MessageChatsResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_bot_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageChatsResponse_Item.ProtoReflect.Descriptor instead.
func (*MessageChatsResponse_Item) Descriptor() ([]byte, []int) {
	return file_bot_message_proto_rawDescGZIP(), []int{3, 0}
}

func (x *MessageChatsResponse_Item) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MessageChatsResponse_Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_bot_message_proto protoreflect.FileDescriptor

const file_bot_message_proto_rawDesc = "" +
	"\n" +
	"\x11bot/message.proto\x12\x03bot\x1a\x13tagger/tagger.proto\"\xab\x01\n" +
	"\x12MessageSendRequest\x12O\n" +
	"\achat_id\x18\x01 \x01(\x05B6\x9a\x84\x9e\x031form:\"chat_id\" binding:\"required\" label:\"chat_id\"R\x06chatId\x12D\n" +
	"\x04text\x18\x02 \x01(\tB0\x9a\x84\x9e\x03+form:\"text\" binding:\"required\" label:\"text\"R\x04text\"\x15\n" +
	"\x13MessageSendResponse\"\x15\n" +
	"\x13MessageChatsRequest\"\xad\x01\n" +
	"\x14MessageChatsResponse\x12G\n" +
	"\x05items\x18\x01 \x03(\v2\x1e.bot.MessageChatsResponse.ItemB\x11\x9a\x84\x9e\x03\fjson:\"items\"R\x05items\x1aL\n" +
	"\x04Item\x12\x1e\n" +
	"\x02id\x18\x01 \x01(\x05B\x0e\x9a\x84\x9e\x03\tjson:\"id\"R\x02id\x12$\n" +
	"\x04name\x18\x02 \x01(\tB\x10\x9a\x84\x9e\x03\vjson:\"name\"R\x04nameB\x0eZ\f./bot;bot_pbb\x06proto3"

var (
	file_bot_message_proto_rawDescOnce sync.Once
	file_bot_message_proto_rawDescData []byte
)

func file_bot_message_proto_rawDescGZIP() []byte {
	file_bot_message_proto_rawDescOnce.Do(func() {
		file_bot_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_bot_message_proto_rawDesc), len(file_bot_message_proto_rawDesc)))
	})
	return file_bot_message_proto_rawDescData
}

var file_bot_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bot_message_proto_goTypes = []any{
	(*MessageSendRequest)(nil),        // 0: bot.MessageSendRequest
	(*MessageSendResponse)(nil),       // 1: bot.MessageSendResponse
	(*MessageChatsRequest)(nil),       // 2: bot.MessageChatsRequest
	(*MessageChatsResponse)(nil),      // 3: bot.MessageChatsResponse
	(*MessageChatsResponse_Item)(nil), // 4: bot.MessageChatsResponse.Item
}
var file_bot_message_proto_depIdxs = []int32{
	4, // 0: bot.MessageChatsResponse.items:type_name -> bot.MessageChatsResponse.Item
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bot_message_proto_init() }
func file_bot_message_proto_init() {
	if File_bot_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_bot_message_proto_rawDesc), len(file_bot_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bot_message_proto_goTypes,
		DependencyIndexes: file_bot_message_proto_depIdxs,
		MessageInfos:      file_bot_message_proto_msgTypes,
	}.Build()
	File_bot_message_proto = out.File
	file_bot_message_proto_goTypes = nil
	file_bot_message_proto_depIdxs = nil
}
