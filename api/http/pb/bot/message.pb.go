// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: bot/message.proto

package bot_pb

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

type MessageSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId int32  `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty" binding:"required" form:"chat_id"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty" binding:"required" form:"text"`
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*MessageChatsResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
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

// TODO DELETE
type MessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId  int32  `protobuf:"varint,1,opt,name=chatId,proto3" json:"chatId,omitempty" binding:"required" form:"chatId"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty" binding:"required" form:"content"`
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	mi := &file_bot_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_bot_message_proto_rawDescGZIP(), []int{4}
}

func (x *MessageRequest) GetChatId() int32 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *MessageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type MessageChatsResponse_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MessageChatsResponse_Item) Reset() {
	*x = MessageChatsResponse_Item{}
	mi := &file_bot_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageChatsResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageChatsResponse_Item) ProtoMessage() {}

func (x *MessageChatsResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_bot_message_proto_msgTypes[5]
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

var file_bot_message_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x6f, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x62, 0x6f, 0x74, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01,
	0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x26, 0x9a, 0x84, 0x9e, 0x03, 0x21, 0x66, 0x6f, 0x72, 0x6d,
	0x3a, 0x22, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x06, 0x63,
	0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x23, 0x9a, 0x84, 0x9e, 0x03, 0x1e, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x15,
	0x0a, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x78, 0x0a, 0x14,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x2a, 0x0a, 0x04, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x63, 0x68, 0x61,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x25, 0x9a, 0x84, 0x9e, 0x03, 0x20,
	0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x20, 0x62, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22,
	0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0x9a, 0x84, 0x9e, 0x03, 0x21,
	0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x20, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x22, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f,
	0x62, 0x6f, 0x74, 0x3b, 0x62, 0x6f, 0x74, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_bot_message_proto_rawDescOnce sync.Once
	file_bot_message_proto_rawDescData = file_bot_message_proto_rawDesc
)

func file_bot_message_proto_rawDescGZIP() []byte {
	file_bot_message_proto_rawDescOnce.Do(func() {
		file_bot_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_bot_message_proto_rawDescData)
	})
	return file_bot_message_proto_rawDescData
}

var file_bot_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bot_message_proto_goTypes = []any{
	(*MessageSendRequest)(nil),        // 0: bot.MessageSendRequest
	(*MessageSendResponse)(nil),       // 1: bot.MessageSendResponse
	(*MessageChatsRequest)(nil),       // 2: bot.MessageChatsRequest
	(*MessageChatsResponse)(nil),      // 3: bot.MessageChatsResponse
	(*MessageRequest)(nil),            // 4: bot.MessageRequest
	(*MessageChatsResponse_Item)(nil), // 5: bot.MessageChatsResponse.Item
}
var file_bot_message_proto_depIdxs = []int32{
	5, // 0: bot.MessageChatsResponse.items:type_name -> bot.MessageChatsResponse.Item
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
			RawDescriptor: file_bot_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bot_message_proto_goTypes,
		DependencyIndexes: file_bot_message_proto_depIdxs,
		MessageInfos:      file_bot_message_proto_msgTypes,
	}.Build()
	File_bot_message_proto = out.File
	file_bot_message_proto_rawDesc = nil
	file_bot_message_proto_goTypes = nil
	file_bot_message_proto_depIdxs = nil
}
