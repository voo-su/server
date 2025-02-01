// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.1
// source: manager/dashboard.proto

package manager_pb

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

type ManagerDashboardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ManagerDashboardRequest) Reset() {
	*x = ManagerDashboardRequest{}
	mi := &file_manager_dashboard_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ManagerDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerDashboardRequest) ProtoMessage() {}

func (x *ManagerDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dashboard_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerDashboardRequest.ProtoReflect.Descriptor instead.
func (*ManagerDashboardRequest) Descriptor() ([]byte, []int) {
	return file_manager_dashboard_proto_rawDescGZIP(), []int{0}
}

type ManagerDashboardResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Users           int64                  `protobuf:"varint,1,opt,name=users,proto3" json:"users,omitempty"`
	Bots            int64                  `protobuf:"varint,2,opt,name=bots,proto3" json:"bots,omitempty"`
	TotalMessages   int64                  `protobuf:"varint,3,opt,name=total_messages,json=totalMessages,proto3" json:"total_messages,omitempty"`
	GroupChats      int64                  `protobuf:"varint,4,opt,name=group_chats,json=groupChats,proto3" json:"group_chats,omitempty"`
	GroupMessages   int64                  `protobuf:"varint,5,opt,name=group_messages,json=groupMessages,proto3" json:"group_messages,omitempty"`
	PrivateMessages int64                  `protobuf:"varint,6,opt,name=private_messages,json=privateMessages,proto3" json:"private_messages,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ManagerDashboardResponse) Reset() {
	*x = ManagerDashboardResponse{}
	mi := &file_manager_dashboard_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ManagerDashboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerDashboardResponse) ProtoMessage() {}

func (x *ManagerDashboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_dashboard_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerDashboardResponse.ProtoReflect.Descriptor instead.
func (*ManagerDashboardResponse) Descriptor() ([]byte, []int) {
	return file_manager_dashboard_proto_rawDescGZIP(), []int{1}
}

func (x *ManagerDashboardResponse) GetUsers() int64 {
	if x != nil {
		return x.Users
	}
	return 0
}

func (x *ManagerDashboardResponse) GetBots() int64 {
	if x != nil {
		return x.Bots
	}
	return 0
}

func (x *ManagerDashboardResponse) GetTotalMessages() int64 {
	if x != nil {
		return x.TotalMessages
	}
	return 0
}

func (x *ManagerDashboardResponse) GetGroupChats() int64 {
	if x != nil {
		return x.GroupChats
	}
	return 0
}

func (x *ManagerDashboardResponse) GetGroupMessages() int64 {
	if x != nil {
		return x.GroupMessages
	}
	return 0
}

func (x *ManagerDashboardResponse) GetPrivateMessages() int64 {
	if x != nil {
		return x.PrivateMessages
	}
	return 0
}

var File_manager_dashboard_proto protoreflect.FileDescriptor

var file_manager_dashboard_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x17, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xf2, 0x02, 0x0a, 0x18, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x11,
	0x9a, 0x84, 0x9e, 0x03, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x22, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x6f, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x62, 0x6f, 0x74, 0x73, 0x22, 0x52, 0x04, 0x62, 0x6f, 0x74, 0x73, 0x12, 0x41,
	0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1a, 0x9a, 0x84, 0x9e, 0x03, 0x15, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x22, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x38, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x73, 0x22, 0x52,
	0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x61, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x1a, 0x9a, 0x84, 0x9e, 0x03, 0x15, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x52,
	0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x47,
	0x0a, 0x10, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x1c, 0x9a, 0x84, 0x9e, 0x03, 0x17, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x3b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manager_dashboard_proto_rawDescOnce sync.Once
	file_manager_dashboard_proto_rawDescData = file_manager_dashboard_proto_rawDesc
)

func file_manager_dashboard_proto_rawDescGZIP() []byte {
	file_manager_dashboard_proto_rawDescOnce.Do(func() {
		file_manager_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_manager_dashboard_proto_rawDescData)
	})
	return file_manager_dashboard_proto_rawDescData
}

var file_manager_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_manager_dashboard_proto_goTypes = []any{
	(*ManagerDashboardRequest)(nil),  // 0: manager.ManagerDashboardRequest
	(*ManagerDashboardResponse)(nil), // 1: manager.ManagerDashboardResponse
}
var file_manager_dashboard_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_manager_dashboard_proto_init() }
func file_manager_dashboard_proto_init() {
	if File_manager_dashboard_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_manager_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_manager_dashboard_proto_goTypes,
		DependencyIndexes: file_manager_dashboard_proto_depIdxs,
		MessageInfos:      file_manager_dashboard_proto_msgTypes,
	}.Build()
	File_manager_dashboard_proto = out.File
	file_manager_dashboard_proto_rawDesc = nil
	file_manager_dashboard_proto_goTypes = nil
	file_manager_dashboard_proto_depIdxs = nil
}
