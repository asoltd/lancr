// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: lancr/v1/quest_service.proto

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

type ReadQuestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadQuestRequest) Reset() {
	*x = ReadQuestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_quest_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadQuestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadQuestRequest) ProtoMessage() {}

func (x *ReadQuestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_quest_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadQuestRequest.ProtoReflect.Descriptor instead.
func (*ReadQuestRequest) Descriptor() ([]byte, []int) {
	return file_lancr_v1_quest_service_proto_rawDescGZIP(), []int{0}
}

func (x *ReadQuestRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadQuestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Quest `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ReadQuestResponse) Reset() {
	*x = ReadQuestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_quest_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadQuestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadQuestResponse) ProtoMessage() {}

func (x *ReadQuestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_quest_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadQuestResponse.ProtoReflect.Descriptor instead.
func (*ReadQuestResponse) Descriptor() ([]byte, []int) {
	return file_lancr_v1_quest_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReadQuestResponse) GetResult() *Quest {
	if x != nil {
		return x.Result
	}
	return nil
}

type ListQuestsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken int32 `protobuf:"varint,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListQuestsRequest) Reset() {
	*x = ListQuestsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_quest_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuestsRequest) ProtoMessage() {}

func (x *ListQuestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_quest_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuestsRequest.ProtoReflect.Descriptor instead.
func (*ListQuestsRequest) Descriptor() ([]byte, []int) {
	return file_lancr_v1_quest_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListQuestsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListQuestsRequest) GetPageToken() int32 {
	if x != nil {
		return x.PageToken
	}
	return 0
}

type ListQuestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results       []*Quest `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	NextPageToken int32    `protobuf:"varint,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	TotalSize     int32    `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
}

func (x *ListQuestsResponse) Reset() {
	*x = ListQuestsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_quest_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuestsResponse) ProtoMessage() {}

func (x *ListQuestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_quest_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuestsResponse.ProtoReflect.Descriptor instead.
func (*ListQuestsResponse) Descriptor() ([]byte, []int) {
	return file_lancr_v1_quest_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListQuestsResponse) GetResults() []*Quest {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListQuestsResponse) GetNextPageToken() int32 {
	if x != nil {
		return x.NextPageToken
	}
	return 0
}

func (x *ListQuestsResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

type CreateQuestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *Quest `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *CreateQuestRequest) Reset() {
	*x = CreateQuestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_quest_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuestRequest) ProtoMessage() {}

func (x *CreateQuestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_quest_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuestRequest.ProtoReflect.Descriptor instead.
func (*CreateQuestRequest) Descriptor() ([]byte, []int) {
	return file_lancr_v1_quest_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateQuestRequest) GetPayload() *Quest {
	if x != nil {
		return x.Payload
	}
	return nil
}

type CreateQuestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Quest `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateQuestResponse) Reset() {
	*x = CreateQuestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_quest_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuestResponse) ProtoMessage() {}

func (x *CreateQuestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_quest_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuestResponse.ProtoReflect.Descriptor instead.
func (*CreateQuestResponse) Descriptor() ([]byte, []int) {
	return file_lancr_v1_quest_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateQuestResponse) GetResult() *Quest {
	if x != nil {
		return x.Result
	}
	return nil
}

type UpdateQuestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload      *Quest `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	AllowMissing bool   `protobuf:"varint,3,opt,name=allow_missing,json=allowMissing,proto3" json:"allow_missing,omitempty"`
}

func (x *UpdateQuestRequest) Reset() {
	*x = UpdateQuestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_quest_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuestRequest) ProtoMessage() {}

func (x *UpdateQuestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_quest_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuestRequest.ProtoReflect.Descriptor instead.
func (*UpdateQuestRequest) Descriptor() ([]byte, []int) {
	return file_lancr_v1_quest_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateQuestRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateQuestRequest) GetPayload() *Quest {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *UpdateQuestRequest) GetAllowMissing() bool {
	if x != nil {
		return x.AllowMissing
	}
	return false
}

type UpdateQuestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Quest `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UpdateQuestResponse) Reset() {
	*x = UpdateQuestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_quest_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuestResponse) ProtoMessage() {}

func (x *UpdateQuestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_quest_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuestResponse.ProtoReflect.Descriptor instead.
func (*UpdateQuestResponse) Descriptor() ([]byte, []int) {
	return file_lancr_v1_quest_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateQuestResponse) GetResult() *Quest {
	if x != nil {
		return x.Result
	}
	return nil
}

type DeleteQuestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteQuestRequest) Reset() {
	*x = DeleteQuestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_quest_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuestRequest) ProtoMessage() {}

func (x *DeleteQuestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_quest_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuestRequest.ProtoReflect.Descriptor instead.
func (*DeleteQuestRequest) Descriptor() ([]byte, []int) {
	return file_lancr_v1_quest_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteQuestRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteQuestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteQuestResponse) Reset() {
	*x = DeleteQuestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_quest_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuestResponse) ProtoMessage() {}

func (x *DeleteQuestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_quest_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuestResponse.ProtoReflect.Descriptor instead.
func (*DeleteQuestResponse) Descriptor() ([]byte, []int) {
	return file_lancr_v1_quest_service_proto_rawDescGZIP(), []int{9}
}

type GetTopQuestersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTopQuestersRequest) Reset() {
	*x = GetTopQuestersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_quest_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopQuestersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopQuestersRequest) ProtoMessage() {}

func (x *GetTopQuestersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_quest_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopQuestersRequest.ProtoReflect.Descriptor instead.
func (*GetTopQuestersRequest) Descriptor() ([]byte, []int) {
	return file_lancr_v1_quest_service_proto_rawDescGZIP(), []int{10}
}

type GetTopQuestersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopQuesters []*Hero `protobuf:"bytes,1,rep,name=top_questers,json=topQuesters,proto3" json:"top_questers,omitempty"`
}

func (x *GetTopQuestersResponse) Reset() {
	*x = GetTopQuestersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_quest_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopQuestersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopQuestersResponse) ProtoMessage() {}

func (x *GetTopQuestersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_quest_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopQuestersResponse.ProtoReflect.Descriptor instead.
func (*GetTopQuestersResponse) Descriptor() ([]byte, []int) {
	return file_lancr_v1_quest_service_proto_rawDescGZIP(), []int{11}
}

func (x *GetTopQuestersResponse) GetTopQuesters() []*Hero {
	if x != nil {
		return x.TopQuesters
	}
	return nil
}

var File_lancr_v1_quest_service_proto protoreflect.FileDescriptor

var file_lancr_v1_quest_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x10, 0x52,
	0x65, 0x61, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x3c, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x4f, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x86,
	0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x3f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x3e, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x74, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x22, 0x3e,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x24,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x70, 0x51, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x4b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x0c, 0x74, 0x6f, 0x70, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x72, 0x6f, 0x52, 0x0b, 0x74, 0x6f, 0x70, 0x51, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x32, 0x85, 0x05, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5d, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x61,
	0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x5b, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12,
	0x1b, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c,
	0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x61,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x2e,
	0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x61,
	0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x12, 0x66, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x1a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6e, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0xba, 0xb9, 0x19, 0x07, 0x0a, 0x05, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x2a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x76, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x70, 0x51, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x6c, 0x61,
	0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6c,
	0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x2a, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x3a, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x51, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x1a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x6f, 0x6c, 0x74, 0x64, 0x2f, 0x6c,
	0x61, 0x6e, 0x63, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x63,
	0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lancr_v1_quest_service_proto_rawDescOnce sync.Once
	file_lancr_v1_quest_service_proto_rawDescData = file_lancr_v1_quest_service_proto_rawDesc
)

func file_lancr_v1_quest_service_proto_rawDescGZIP() []byte {
	file_lancr_v1_quest_service_proto_rawDescOnce.Do(func() {
		file_lancr_v1_quest_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_lancr_v1_quest_service_proto_rawDescData)
	})
	return file_lancr_v1_quest_service_proto_rawDescData
}

var file_lancr_v1_quest_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_lancr_v1_quest_service_proto_goTypes = []interface{}{
	(*ReadQuestRequest)(nil),       // 0: lancr.v1.ReadQuestRequest
	(*ReadQuestResponse)(nil),      // 1: lancr.v1.ReadQuestResponse
	(*ListQuestsRequest)(nil),      // 2: lancr.v1.ListQuestsRequest
	(*ListQuestsResponse)(nil),     // 3: lancr.v1.ListQuestsResponse
	(*CreateQuestRequest)(nil),     // 4: lancr.v1.CreateQuestRequest
	(*CreateQuestResponse)(nil),    // 5: lancr.v1.CreateQuestResponse
	(*UpdateQuestRequest)(nil),     // 6: lancr.v1.UpdateQuestRequest
	(*UpdateQuestResponse)(nil),    // 7: lancr.v1.UpdateQuestResponse
	(*DeleteQuestRequest)(nil),     // 8: lancr.v1.DeleteQuestRequest
	(*DeleteQuestResponse)(nil),    // 9: lancr.v1.DeleteQuestResponse
	(*GetTopQuestersRequest)(nil),  // 10: lancr.v1.GetTopQuestersRequest
	(*GetTopQuestersResponse)(nil), // 11: lancr.v1.GetTopQuestersResponse
	(*Quest)(nil),                  // 12: lancr.v1.Quest
	(*Hero)(nil),                   // 13: lancr.v1.Hero
}
var file_lancr_v1_quest_service_proto_depIdxs = []int32{
	12, // 0: lancr.v1.ReadQuestResponse.result:type_name -> lancr.v1.Quest
	12, // 1: lancr.v1.ListQuestsResponse.results:type_name -> lancr.v1.Quest
	12, // 2: lancr.v1.CreateQuestRequest.payload:type_name -> lancr.v1.Quest
	12, // 3: lancr.v1.CreateQuestResponse.result:type_name -> lancr.v1.Quest
	12, // 4: lancr.v1.UpdateQuestRequest.payload:type_name -> lancr.v1.Quest
	12, // 5: lancr.v1.UpdateQuestResponse.result:type_name -> lancr.v1.Quest
	13, // 6: lancr.v1.GetTopQuestersResponse.top_questers:type_name -> lancr.v1.Hero
	0,  // 7: lancr.v1.QuestService.ReadQuest:input_type -> lancr.v1.ReadQuestRequest
	2,  // 8: lancr.v1.QuestService.ListQuests:input_type -> lancr.v1.ListQuestsRequest
	4,  // 9: lancr.v1.QuestService.CreateQuest:input_type -> lancr.v1.CreateQuestRequest
	6,  // 10: lancr.v1.QuestService.UpdateQuest:input_type -> lancr.v1.UpdateQuestRequest
	8,  // 11: lancr.v1.QuestService.DeleteQuest:input_type -> lancr.v1.DeleteQuestRequest
	10, // 12: lancr.v1.QuestService.GetTopQuesters:input_type -> lancr.v1.GetTopQuestersRequest
	1,  // 13: lancr.v1.QuestService.ReadQuest:output_type -> lancr.v1.ReadQuestResponse
	3,  // 14: lancr.v1.QuestService.ListQuests:output_type -> lancr.v1.ListQuestsResponse
	5,  // 15: lancr.v1.QuestService.CreateQuest:output_type -> lancr.v1.CreateQuestResponse
	7,  // 16: lancr.v1.QuestService.UpdateQuest:output_type -> lancr.v1.UpdateQuestResponse
	9,  // 17: lancr.v1.QuestService.DeleteQuest:output_type -> lancr.v1.DeleteQuestResponse
	11, // 18: lancr.v1.QuestService.GetTopQuesters:output_type -> lancr.v1.GetTopQuestersResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_lancr_v1_quest_service_proto_init() }
func file_lancr_v1_quest_service_proto_init() {
	if File_lancr_v1_quest_service_proto != nil {
		return
	}
	file_lancr_v1_quest_proto_init()
	file_lancr_v1_hero_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lancr_v1_quest_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadQuestRequest); i {
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
		file_lancr_v1_quest_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadQuestResponse); i {
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
		file_lancr_v1_quest_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuestsRequest); i {
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
		file_lancr_v1_quest_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuestsResponse); i {
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
		file_lancr_v1_quest_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuestRequest); i {
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
		file_lancr_v1_quest_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuestResponse); i {
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
		file_lancr_v1_quest_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuestRequest); i {
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
		file_lancr_v1_quest_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuestResponse); i {
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
		file_lancr_v1_quest_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQuestRequest); i {
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
		file_lancr_v1_quest_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQuestResponse); i {
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
		file_lancr_v1_quest_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopQuestersRequest); i {
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
		file_lancr_v1_quest_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopQuestersResponse); i {
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
			RawDescriptor: file_lancr_v1_quest_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lancr_v1_quest_service_proto_goTypes,
		DependencyIndexes: file_lancr_v1_quest_service_proto_depIdxs,
		MessageInfos:      file_lancr_v1_quest_service_proto_msgTypes,
	}.Build()
	File_lancr_v1_quest_service_proto = out.File
	file_lancr_v1_quest_service_proto_rawDesc = nil
	file_lancr_v1_quest_service_proto_goTypes = nil
	file_lancr_v1_quest_service_proto_depIdxs = nil
}