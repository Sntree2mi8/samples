// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: omiyage/v1/service.proto

package omiyagev1

import (
	v11 "github.com/Sntree2mi8/samples/omiyage_management/proto/gen/go/types/recipient/v1"
	v1 "github.com/Sntree2mi8/samples/omiyage_management/proto/gen/go/types/user/v1"
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

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ユーザ名
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omiyage_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omiyage_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_omiyage_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type SignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *SignUpResponse_DomainError `protobuf:"bytes,1,opt,name=error,proto3,oneof" json:"error,omitempty"`
	User  *v1.User                    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omiyage_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omiyage_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_omiyage_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpResponse) GetError() *SignUpResponse_DomainError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SignUpResponse) GetUser() *v1.User {
	if x != nil {
		return x.User
	}
	return nil
}

type AddRecipientGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 送り手ユーザID
	//
	// サンプルのプロジェクトでガッツリ作り込みたくないからどのユーザのリクエストかはここで判断しちゃう. contextとかからみるみたいなことはしない.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 受け取り手集団
	RecipientGroups []*v11.RecipientGroup `protobuf:"bytes,2,rep,name=recipient_groups,json=recipientGroups,proto3" json:"recipient_groups,omitempty"`
}

func (x *AddRecipientGroupsRequest) Reset() {
	*x = AddRecipientGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omiyage_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecipientGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecipientGroupsRequest) ProtoMessage() {}

func (x *AddRecipientGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omiyage_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecipientGroupsRequest.ProtoReflect.Descriptor instead.
func (*AddRecipientGroupsRequest) Descriptor() ([]byte, []int) {
	return file_omiyage_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *AddRecipientGroupsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddRecipientGroupsRequest) GetRecipientGroups() []*v11.RecipientGroup {
	if x != nil {
		return x.RecipientGroups
	}
	return nil
}

type AddRecipientGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *AddRecipientGroupsResponse_DomainError `protobuf:"bytes,1,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *AddRecipientGroupsResponse) Reset() {
	*x = AddRecipientGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omiyage_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecipientGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecipientGroupsResponse) ProtoMessage() {}

func (x *AddRecipientGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omiyage_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecipientGroupsResponse.ProtoReflect.Descriptor instead.
func (*AddRecipientGroupsResponse) Descriptor() ([]byte, []int) {
	return file_omiyage_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddRecipientGroupsResponse) GetError() *AddRecipientGroupsResponse_DomainError {
	if x != nil {
		return x.Error
	}
	return nil
}

type SignUpResponse_DomainError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignUpResponse_DomainError) Reset() {
	*x = SignUpResponse_DomainError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omiyage_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse_DomainError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse_DomainError) ProtoMessage() {}

func (x *SignUpResponse_DomainError) ProtoReflect() protoreflect.Message {
	mi := &file_omiyage_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse_DomainError.ProtoReflect.Descriptor instead.
func (*SignUpResponse_DomainError) Descriptor() ([]byte, []int) {
	return file_omiyage_v1_service_proto_rawDescGZIP(), []int{1, 0}
}

type AddRecipientGroupsResponse_DomainError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddRecipientGroupsResponse_DomainError) Reset() {
	*x = AddRecipientGroupsResponse_DomainError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omiyage_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecipientGroupsResponse_DomainError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecipientGroupsResponse_DomainError) ProtoMessage() {}

func (x *AddRecipientGroupsResponse_DomainError) ProtoReflect() protoreflect.Message {
	mi := &file_omiyage_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecipientGroupsResponse_DomainError.ProtoReflect.Descriptor instead.
func (*AddRecipientGroupsResponse_DomainError) Descriptor() ([]byte, []int) {
	return file_omiyage_v1_service_proto_rawDescGZIP(), []int{3, 0}
}

var File_omiyage_v1_service_proto protoreflect.FileDescriptor

var file_omiyage_v1_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6f, 0x6d, 0x69, 0x79, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x6d, 0x69, 0x79,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x0d, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x0e, 0x53, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x6d, 0x69,
	0x79, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x27,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0d, 0x0a, 0x0b, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x83, 0x01, 0x0a, 0x19, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4d, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x1a, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6f, 0x6d, 0x69, 0x79, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x88, 0x01, 0x01, 0x1a, 0x0d, 0x0a, 0x0b, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xb6, 0x01,
	0x0a, 0x0e, 0x4f, 0x6d, 0x69, 0x79, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3f, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x19, 0x2e, 0x6f, 0x6d, 0x69,
	0x79, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6f, 0x6d, 0x69, 0x79, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x63, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x25, 0x2e, 0x6f, 0x6d, 0x69, 0x79, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x6f, 0x6d, 0x69, 0x79, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xbb, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x6f,
	0x6d, 0x69, 0x79, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x74, 0x72, 0x65, 0x65, 0x32, 0x6d, 0x69, 0x38,
	0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x6f, 0x6d, 0x69, 0x79, 0x61, 0x67, 0x65,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x6d, 0x69, 0x79, 0x61, 0x67, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x6d, 0x69, 0x79, 0x61, 0x67, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x4f, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x4f, 0x6d, 0x69, 0x79, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0a, 0x4f, 0x6d, 0x69, 0x79, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16,
	0x4f, 0x6d, 0x69, 0x79, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x4f, 0x6d, 0x69, 0x79, 0x61, 0x67, 0x65,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_omiyage_v1_service_proto_rawDescOnce sync.Once
	file_omiyage_v1_service_proto_rawDescData = file_omiyage_v1_service_proto_rawDesc
)

func file_omiyage_v1_service_proto_rawDescGZIP() []byte {
	file_omiyage_v1_service_proto_rawDescOnce.Do(func() {
		file_omiyage_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_omiyage_v1_service_proto_rawDescData)
	})
	return file_omiyage_v1_service_proto_rawDescData
}

var file_omiyage_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_omiyage_v1_service_proto_goTypes = []interface{}{
	(*SignUpRequest)(nil),                          // 0: omiyage.v1.SignUpRequest
	(*SignUpResponse)(nil),                         // 1: omiyage.v1.SignUpResponse
	(*AddRecipientGroupsRequest)(nil),              // 2: omiyage.v1.AddRecipientGroupsRequest
	(*AddRecipientGroupsResponse)(nil),             // 3: omiyage.v1.AddRecipientGroupsResponse
	(*SignUpResponse_DomainError)(nil),             // 4: omiyage.v1.SignUpResponse.DomainError
	(*AddRecipientGroupsResponse_DomainError)(nil), // 5: omiyage.v1.AddRecipientGroupsResponse.DomainError
	(*v1.User)(nil),                                // 6: types.user.v1.User
	(*v11.RecipientGroup)(nil),                     // 7: types.recipient.v1.RecipientGroup
}
var file_omiyage_v1_service_proto_depIdxs = []int32{
	4, // 0: omiyage.v1.SignUpResponse.error:type_name -> omiyage.v1.SignUpResponse.DomainError
	6, // 1: omiyage.v1.SignUpResponse.user:type_name -> types.user.v1.User
	7, // 2: omiyage.v1.AddRecipientGroupsRequest.recipient_groups:type_name -> types.recipient.v1.RecipientGroup
	5, // 3: omiyage.v1.AddRecipientGroupsResponse.error:type_name -> omiyage.v1.AddRecipientGroupsResponse.DomainError
	0, // 4: omiyage.v1.OmiyageService.SignUp:input_type -> omiyage.v1.SignUpRequest
	2, // 5: omiyage.v1.OmiyageService.AddRecipientGroups:input_type -> omiyage.v1.AddRecipientGroupsRequest
	1, // 6: omiyage.v1.OmiyageService.SignUp:output_type -> omiyage.v1.SignUpResponse
	3, // 7: omiyage.v1.OmiyageService.AddRecipientGroups:output_type -> omiyage.v1.AddRecipientGroupsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_omiyage_v1_service_proto_init() }
func file_omiyage_v1_service_proto_init() {
	if File_omiyage_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_omiyage_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omiyage_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omiyage_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecipientGroupsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omiyage_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecipientGroupsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omiyage_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse_DomainError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omiyage_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecipientGroupsResponse_DomainError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_omiyage_v1_service_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_omiyage_v1_service_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_omiyage_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_omiyage_v1_service_proto_goTypes,
		DependencyIndexes: file_omiyage_v1_service_proto_depIdxs,
		MessageInfos:      file_omiyage_v1_service_proto_msgTypes,
	}.Build()
	File_omiyage_v1_service_proto = out.File
	file_omiyage_v1_service_proto_rawDesc = nil
	file_omiyage_v1_service_proto_goTypes = nil
	file_omiyage_v1_service_proto_depIdxs = nil
}
