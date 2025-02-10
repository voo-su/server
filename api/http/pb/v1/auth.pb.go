// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.1
// source: v1/auth.proto

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

type AuthLoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" binding:"required" label:"email"`
	Platform      string                 `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty" binding:"required,oneof=web android ios linux macos windows" label:"platform"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthLoginRequest) Reset() {
	*x = AuthLoginRequest{}
	mi := &file_v1_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLoginRequest) ProtoMessage() {}

func (x *AuthLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLoginRequest.ProtoReflect.Descriptor instead.
func (*AuthLoginRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthLoginRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

type AuthLoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	ExpiresIn     int32                  `protobuf:"varint,2,opt,name=expiresIn,proto3" json:"expiresIn"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthLoginResponse) Reset() {
	*x = AuthLoginResponse{}
	mi := &file_v1_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLoginResponse) ProtoMessage() {}

func (x *AuthLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLoginResponse.ProtoReflect.Descriptor instead.
func (*AuthLoginResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthLoginResponse) GetExpiresIn() int32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

type AuthVerifyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" binding:"required" label:"token"`
	Code          string                 `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty" binding:"required,len=6" label:"code"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthVerifyRequest) Reset() {
	*x = AuthVerifyRequest{}
	mi := &file_v1_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthVerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthVerifyRequest) ProtoMessage() {}

func (x *AuthVerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthVerifyRequest.ProtoReflect.Descriptor instead.
func (*AuthVerifyRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AuthVerifyRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthVerifyRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type AuthVerifyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	AccessToken   string                 `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken"`
	ExpiresIn     int32                  `protobuf:"varint,3,opt,name=expiresIn,proto3" json:"expiresIn"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthVerifyResponse) Reset() {
	*x = AuthVerifyResponse{}
	mi := &file_v1_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthVerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthVerifyResponse) ProtoMessage() {}

func (x *AuthVerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthVerifyResponse.ProtoReflect.Descriptor instead.
func (*AuthVerifyResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *AuthVerifyResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AuthVerifyResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthVerifyResponse) GetExpiresIn() int32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

type AuthRefreshRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthRefreshRequest) Reset() {
	*x = AuthRefreshRequest{}
	mi := &file_v1_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthRefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRefreshRequest) ProtoMessage() {}

func (x *AuthRefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRefreshRequest.ProtoReflect.Descriptor instead.
func (*AuthRefreshRequest) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{4}
}

type AuthRefreshResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type" label:"type"`
	AccessToken   string                 `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken" label:"accessToken"`
	ExpiresIn     int32                  `protobuf:"varint,3,opt,name=expiresIn,proto3" json:"expiresIn" label:"expiresIn"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthRefreshResponse) Reset() {
	*x = AuthRefreshResponse{}
	mi := &file_v1_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthRefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRefreshResponse) ProtoMessage() {}

func (x *AuthRefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRefreshResponse.ProtoReflect.Descriptor instead.
func (*AuthRefreshResponse) Descriptor() ([]byte, []int) {
	return file_v1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *AuthRefreshResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AuthRefreshResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthRefreshResponse) GetExpiresIn() int32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

var File_v1_auth_proto protoreflect.FileDescriptor

var file_v1_auth_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x10, 0x41, 0x75, 0x74,
	0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0x9a, 0x84,
	0x9e, 0x03, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x22, 0x20, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3a, 0x22, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x6e, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x52, 0x9a, 0x84,
	0x9e, 0x03, 0x4d, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x2c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x3d, 0x77, 0x65, 0x62, 0x20, 0x61,
	0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x20, 0x69, 0x6f, 0x73, 0x20, 0x6c, 0x69, 0x6e, 0x75, 0x78,
	0x20, 0x6d, 0x61, 0x63, 0x6f, 0x73, 0x20, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x22, 0x20,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3a, 0x22, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x71, 0x0a, 0x11, 0x41, 0x75,
	0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11,
	0x9a, 0x84, 0x9e, 0x03, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x33, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x49, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x15, 0x9a, 0x84, 0x9e,
	0x03, 0x10, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49,
	0x6e, 0x22, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x22, 0x90, 0x01,
	0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x25, 0x9a, 0x84, 0x9e, 0x03, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x20, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x3a, 0x22, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x3e, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a,
	0x9a, 0x84, 0x9e, 0x03, 0x25, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6c, 0x65, 0x6e, 0x3d, 0x36, 0x22, 0x20, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x3a, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0xaa, 0x01, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x33, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x49, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x15, 0x9a, 0x84, 0x9e,
	0x03, 0x10, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49,
	0x6e, 0x22, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x22, 0x14, 0x0a,
	0x12, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0xde, 0x01, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0x9a, 0x84, 0x9e, 0x03, 0x18,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x20, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x4d,
	0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2b, 0x9a, 0x84, 0x9e, 0x03, 0x26, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x20, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x3a, 0x22, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x45, 0x0a,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x27, 0x9a, 0x84, 0x9e, 0x03, 0x22, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x22, 0x20, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3a, 0x22, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x22, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x49, 0x6e, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x5f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_v1_auth_proto_rawDescOnce sync.Once
	file_v1_auth_proto_rawDescData []byte
)

func file_v1_auth_proto_rawDescGZIP() []byte {
	file_v1_auth_proto_rawDescOnce.Do(func() {
		file_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_auth_proto_rawDesc), len(file_v1_auth_proto_rawDesc)))
	})
	return file_v1_auth_proto_rawDescData
}

var file_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_auth_proto_goTypes = []any{
	(*AuthLoginRequest)(nil),    // 0: v1.AuthLoginRequest
	(*AuthLoginResponse)(nil),   // 1: v1.AuthLoginResponse
	(*AuthVerifyRequest)(nil),   // 2: v1.AuthVerifyRequest
	(*AuthVerifyResponse)(nil),  // 3: v1.AuthVerifyResponse
	(*AuthRefreshRequest)(nil),  // 4: v1.AuthRefreshRequest
	(*AuthRefreshResponse)(nil), // 5: v1.AuthRefreshResponse
}
var file_v1_auth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_auth_proto_init() }
func file_v1_auth_proto_init() {
	if File_v1_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_auth_proto_rawDesc), len(file_v1_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_auth_proto_goTypes,
		DependencyIndexes: file_v1_auth_proto_depIdxs,
		MessageInfos:      file_v1_auth_proto_msgTypes,
	}.Build()
	File_v1_auth_proto = out.File
	file_v1_auth_proto_goTypes = nil
	file_v1_auth_proto_depIdxs = nil
}
