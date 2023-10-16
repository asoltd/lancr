// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: lancr/v1/teams_service.proto

package v1

import (
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ReadTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadTeamRequest) Reset() {
	*x = ReadTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_teams_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadTeamRequest) ProtoMessage() {}

func (x *ReadTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_teams_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadTeamRequest.ProtoReflect.Descriptor instead.
func (*ReadTeamRequest) Descriptor() ([]byte, []int) {
	return file_lancr_v1_teams_service_proto_rawDescGZIP(), []int{0}
}

func (x *ReadTeamRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Team `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ReadTeamResponse) Reset() {
	*x = ReadTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_teams_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadTeamResponse) ProtoMessage() {}

func (x *ReadTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_teams_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadTeamResponse.ProtoReflect.Descriptor instead.
func (*ReadTeamResponse) Descriptor() ([]byte, []int) {
	return file_lancr_v1_teams_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReadTeamResponse) GetResult() *Team {
	if x != nil {
		return x.Result
	}
	return nil
}

type ListTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken int32 `protobuf:"varint,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListTeamsRequest) Reset() {
	*x = ListTeamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_teams_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeamsRequest) ProtoMessage() {}

func (x *ListTeamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_teams_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeamsRequest.ProtoReflect.Descriptor instead.
func (*ListTeamsRequest) Descriptor() ([]byte, []int) {
	return file_lancr_v1_teams_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListTeamsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTeamsRequest) GetPageToken() int32 {
	if x != nil {
		return x.PageToken
	}
	return 0
}

type ListTeamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results       []*Team `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	NextPageToken int32   `protobuf:"varint,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	TotalSize     int32   `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
}

func (x *ListTeamsResponse) Reset() {
	*x = ListTeamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_teams_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeamsResponse) ProtoMessage() {}

func (x *ListTeamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_teams_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeamsResponse.ProtoReflect.Descriptor instead.
func (*ListTeamsResponse) Descriptor() ([]byte, []int) {
	return file_lancr_v1_teams_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListTeamsResponse) GetResults() []*Team {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListTeamsResponse) GetNextPageToken() int32 {
	if x != nil {
		return x.NextPageToken
	}
	return 0
}

func (x *ListTeamsResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

type CreateTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *Team `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *CreateTeamRequest) Reset() {
	*x = CreateTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_teams_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamRequest) ProtoMessage() {}

func (x *CreateTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_teams_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamRequest.ProtoReflect.Descriptor instead.
func (*CreateTeamRequest) Descriptor() ([]byte, []int) {
	return file_lancr_v1_teams_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTeamRequest) GetPayload() *Team {
	if x != nil {
		return x.Payload
	}
	return nil
}

type CreateTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Team `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateTeamResponse) Reset() {
	*x = CreateTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_teams_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamResponse) ProtoMessage() {}

func (x *CreateTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_teams_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamResponse.ProtoReflect.Descriptor instead.
func (*CreateTeamResponse) Descriptor() ([]byte, []int) {
	return file_lancr_v1_teams_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTeamResponse) GetResult() *Team {
	if x != nil {
		return x.Result
	}
	return nil
}

type UpdateTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *Team `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *UpdateTeamRequest) Reset() {
	*x = UpdateTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_teams_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeamRequest) ProtoMessage() {}

func (x *UpdateTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_teams_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeamRequest.ProtoReflect.Descriptor instead.
func (*UpdateTeamRequest) Descriptor() ([]byte, []int) {
	return file_lancr_v1_teams_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTeamRequest) GetPayload() *Team {
	if x != nil {
		return x.Payload
	}
	return nil
}

type UpdateTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Team `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UpdateTeamResponse) Reset() {
	*x = UpdateTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_teams_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeamResponse) ProtoMessage() {}

func (x *UpdateTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_teams_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeamResponse.ProtoReflect.Descriptor instead.
func (*UpdateTeamResponse) Descriptor() ([]byte, []int) {
	return file_lancr_v1_teams_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTeamResponse) GetResult() *Team {
	if x != nil {
		return x.Result
	}
	return nil
}

type DeleteTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTeamRequest) Reset() {
	*x = DeleteTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_teams_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeamRequest) ProtoMessage() {}

func (x *DeleteTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_teams_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeamRequest.ProtoReflect.Descriptor instead.
func (*DeleteTeamRequest) Descriptor() ([]byte, []int) {
	return file_lancr_v1_teams_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteTeamRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTeamResponse) Reset() {
	*x = DeleteTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_teams_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeamResponse) ProtoMessage() {}

func (x *DeleteTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_teams_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeamResponse.ProtoReflect.Descriptor instead.
func (*DeleteTeamResponse) Descriptor() ([]byte, []int) {
	return file_lancr_v1_teams_service_proto_rawDescGZIP(), []int{9}
}

var File_lancr_v1_teams_service_proto protoreflect.FileDescriptor

var file_lancr_v1_teams_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c,
	0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x4e, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61,
	0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x3d, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x3c, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x3c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xf3, 0x03, 0x0a, 0x0c, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x59, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e,
	0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x54, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x57, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x2e, 0x6c, 0x61, 0x6e, 0x63,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x5d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x6d, 0x12, 0x1b, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x65, 0x61, 0x6d, 0x73, 0x12, 0x5d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x12, 0x1b, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x1a, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x61, 0x6d, 0x73, 0x12, 0x69, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61,
	0x6d, 0x12, 0x1b, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0xba, 0xb9,
	0x19, 0x06, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x1a, 0x06,
	0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x6f, 0x6c, 0x74, 0x64, 0x2f, 0x6c, 0x61, 0x6e, 0x63,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lancr_v1_teams_service_proto_rawDescOnce sync.Once
	file_lancr_v1_teams_service_proto_rawDescData = file_lancr_v1_teams_service_proto_rawDesc
)

func file_lancr_v1_teams_service_proto_rawDescGZIP() []byte {
	file_lancr_v1_teams_service_proto_rawDescOnce.Do(func() {
		file_lancr_v1_teams_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_lancr_v1_teams_service_proto_rawDescData)
	})
	return file_lancr_v1_teams_service_proto_rawDescData
}

var file_lancr_v1_teams_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_lancr_v1_teams_service_proto_goTypes = []interface{}{
	(*ReadTeamRequest)(nil),    // 0: lancr.v1.ReadTeamRequest
	(*ReadTeamResponse)(nil),   // 1: lancr.v1.ReadTeamResponse
	(*ListTeamsRequest)(nil),   // 2: lancr.v1.ListTeamsRequest
	(*ListTeamsResponse)(nil),  // 3: lancr.v1.ListTeamsResponse
	(*CreateTeamRequest)(nil),  // 4: lancr.v1.CreateTeamRequest
	(*CreateTeamResponse)(nil), // 5: lancr.v1.CreateTeamResponse
	(*UpdateTeamRequest)(nil),  // 6: lancr.v1.UpdateTeamRequest
	(*UpdateTeamResponse)(nil), // 7: lancr.v1.UpdateTeamResponse
	(*DeleteTeamRequest)(nil),  // 8: lancr.v1.DeleteTeamRequest
	(*DeleteTeamResponse)(nil), // 9: lancr.v1.DeleteTeamResponse
	(*Team)(nil),               // 10: lancr.v1.Team
}
var file_lancr_v1_teams_service_proto_depIdxs = []int32{
	10, // 0: lancr.v1.ReadTeamResponse.result:type_name -> lancr.v1.Team
	10, // 1: lancr.v1.ListTeamsResponse.results:type_name -> lancr.v1.Team
	10, // 2: lancr.v1.CreateTeamRequest.payload:type_name -> lancr.v1.Team
	10, // 3: lancr.v1.CreateTeamResponse.result:type_name -> lancr.v1.Team
	10, // 4: lancr.v1.UpdateTeamRequest.payload:type_name -> lancr.v1.Team
	10, // 5: lancr.v1.UpdateTeamResponse.result:type_name -> lancr.v1.Team
	0,  // 6: lancr.v1.TeamsService.ReadTeam:input_type -> lancr.v1.ReadTeamRequest
	2,  // 7: lancr.v1.TeamsService.ListTeams:input_type -> lancr.v1.ListTeamsRequest
	4,  // 8: lancr.v1.TeamsService.CreateTeam:input_type -> lancr.v1.CreateTeamRequest
	6,  // 9: lancr.v1.TeamsService.UpdateTeam:input_type -> lancr.v1.UpdateTeamRequest
	8,  // 10: lancr.v1.TeamsService.DeleteTeam:input_type -> lancr.v1.DeleteTeamRequest
	1,  // 11: lancr.v1.TeamsService.ReadTeam:output_type -> lancr.v1.ReadTeamResponse
	3,  // 12: lancr.v1.TeamsService.ListTeams:output_type -> lancr.v1.ListTeamsResponse
	5,  // 13: lancr.v1.TeamsService.CreateTeam:output_type -> lancr.v1.CreateTeamResponse
	7,  // 14: lancr.v1.TeamsService.UpdateTeam:output_type -> lancr.v1.UpdateTeamResponse
	9,  // 15: lancr.v1.TeamsService.DeleteTeam:output_type -> lancr.v1.DeleteTeamResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_lancr_v1_teams_service_proto_init() }
func file_lancr_v1_teams_service_proto_init() {
	if File_lancr_v1_teams_service_proto != nil {
		return
	}
	file_lancr_v1_team_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lancr_v1_teams_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadTeamRequest); i {
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
		file_lancr_v1_teams_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadTeamResponse); i {
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
		file_lancr_v1_teams_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTeamsRequest); i {
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
		file_lancr_v1_teams_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTeamsResponse); i {
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
		file_lancr_v1_teams_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamRequest); i {
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
		file_lancr_v1_teams_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamResponse); i {
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
		file_lancr_v1_teams_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTeamRequest); i {
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
		file_lancr_v1_teams_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTeamResponse); i {
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
		file_lancr_v1_teams_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTeamRequest); i {
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
		file_lancr_v1_teams_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTeamResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lancr_v1_teams_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lancr_v1_teams_service_proto_goTypes,
		DependencyIndexes: file_lancr_v1_teams_service_proto_depIdxs,
		MessageInfos:      file_lancr_v1_teams_service_proto_msgTypes,
	}.Build()
	File_lancr_v1_teams_service_proto = out.File
	file_lancr_v1_teams_service_proto_rawDesc = nil
	file_lancr_v1_teams_service_proto_goTypes = nil
	file_lancr_v1_teams_service_proto_depIdxs = nil
}
