// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/lancr/v1/hero_service.proto

package v1

import (
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

type GetHeroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHeroRequest) Reset() {
	*x = GetHeroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHeroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHeroRequest) ProtoMessage() {}

func (x *GetHeroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHeroRequest.ProtoReflect.Descriptor instead.
func (*GetHeroRequest) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_hero_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetHeroRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetHeroResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hero *Hero `protobuf:"bytes,1,opt,name=hero,proto3" json:"hero,omitempty"`
}

func (x *GetHeroResponse) Reset() {
	*x = GetHeroResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHeroResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHeroResponse) ProtoMessage() {}

func (x *GetHeroResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHeroResponse.ProtoReflect.Descriptor instead.
func (*GetHeroResponse) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_hero_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetHeroResponse) GetHero() *Hero {
	if x != nil {
		return x.Hero
	}
	return nil
}

type ListHeroesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken int32 `protobuf:"varint,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListHeroesRequest) Reset() {
	*x = ListHeroesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHeroesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHeroesRequest) ProtoMessage() {}

func (x *ListHeroesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHeroesRequest.ProtoReflect.Descriptor instead.
func (*ListHeroesRequest) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_hero_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListHeroesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListHeroesRequest) GetPageToken() int32 {
	if x != nil {
		return x.PageToken
	}
	return 0
}

type ListHeroesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Heroes        []*Hero `protobuf:"bytes,1,rep,name=heroes,proto3" json:"heroes,omitempty"`
	NextPageToken int32   `protobuf:"varint,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	TotalSize     int32   `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	PageSize      int32   `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListHeroesResponse) Reset() {
	*x = ListHeroesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHeroesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHeroesResponse) ProtoMessage() {}

func (x *ListHeroesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHeroesResponse.ProtoReflect.Descriptor instead.
func (*ListHeroesResponse) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_hero_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListHeroesResponse) GetHeroes() []*Hero {
	if x != nil {
		return x.Heroes
	}
	return nil
}

func (x *ListHeroesResponse) GetNextPageToken() int32 {
	if x != nil {
		return x.NextPageToken
	}
	return 0
}

func (x *ListHeroesResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *ListHeroesResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type CreateHeroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hero *Hero `protobuf:"bytes,1,opt,name=hero,proto3" json:"hero,omitempty"`
}

func (x *CreateHeroRequest) Reset() {
	*x = CreateHeroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHeroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHeroRequest) ProtoMessage() {}

func (x *CreateHeroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHeroRequest.ProtoReflect.Descriptor instead.
func (*CreateHeroRequest) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_hero_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateHeroRequest) GetHero() *Hero {
	if x != nil {
		return x.Hero
	}
	return nil
}

type CreateHeroResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hero *Hero `protobuf:"bytes,1,opt,name=hero,proto3" json:"hero,omitempty"`
}

func (x *CreateHeroResponse) Reset() {
	*x = CreateHeroResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHeroResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHeroResponse) ProtoMessage() {}

func (x *CreateHeroResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHeroResponse.ProtoReflect.Descriptor instead.
func (*CreateHeroResponse) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_hero_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateHeroResponse) GetHero() *Hero {
	if x != nil {
		return x.Hero
	}
	return nil
}

type UpdateHeroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hero *Hero  `protobuf:"bytes,2,opt,name=hero,proto3" json:"hero,omitempty"`
	// If set to true, and the hero is not found, a new hero will be created.
	AllowMissing bool `protobuf:"varint,3,opt,name=allow_missing,json=allowMissing,proto3" json:"allow_missing,omitempty"`
}

func (x *UpdateHeroRequest) Reset() {
	*x = UpdateHeroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHeroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHeroRequest) ProtoMessage() {}

func (x *UpdateHeroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHeroRequest.ProtoReflect.Descriptor instead.
func (*UpdateHeroRequest) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_hero_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateHeroRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateHeroRequest) GetHero() *Hero {
	if x != nil {
		return x.Hero
	}
	return nil
}

func (x *UpdateHeroRequest) GetAllowMissing() bool {
	if x != nil {
		return x.AllowMissing
	}
	return false
}

type UpdateHeroResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hero *Hero `protobuf:"bytes,1,opt,name=hero,proto3" json:"hero,omitempty"`
}

func (x *UpdateHeroResponse) Reset() {
	*x = UpdateHeroResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHeroResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHeroResponse) ProtoMessage() {}

func (x *UpdateHeroResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHeroResponse.ProtoReflect.Descriptor instead.
func (*UpdateHeroResponse) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_hero_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateHeroResponse) GetHero() *Hero {
	if x != nil {
		return x.Hero
	}
	return nil
}

type DeleteHeroRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteHeroRequest) Reset() {
	*x = DeleteHeroRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHeroRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHeroRequest) ProtoMessage() {}

func (x *DeleteHeroRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHeroRequest.ProtoReflect.Descriptor instead.
func (*DeleteHeroRequest) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_hero_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteHeroRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteHeroResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteHeroResponse) Reset() {
	*x = DeleteHeroResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHeroResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHeroResponse) ProtoMessage() {}

func (x *DeleteHeroResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_hero_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHeroResponse.ProtoReflect.Descriptor instead.
func (*DeleteHeroResponse) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_hero_service_proto_rawDescGZIP(), []int{9}
}

var File_proto_lancr_v1_hero_service_proto protoreflect.FileDescriptor

var file_proto_lancr_v1_hero_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x68, 0x65, 0x72, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x72, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x72,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48,
	0x65, 0x72, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x68,
	0x65, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x6e, 0x63,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x22,
	0x4f, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0xa0, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0x37, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x72,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x68, 0x65, 0x72, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x22, 0x38, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x72, 0x6f,
	0x52, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x22, 0x6c, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x68,
	0x65, 0x72, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x6e, 0x63,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x12,
	0x23, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4d, 0x69, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x22, 0x38, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65,
	0x72, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x68, 0x65,
	0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x22, 0x23,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x72,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xea, 0x03, 0x0a, 0x0b, 0x48, 0x65,
	0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x48, 0x65, 0x72, 0x6f, 0x12, 0x18, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x72,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x5b, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x65, 0x73,
	0x12, 0x1b, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x72,
	0x6f, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x12,
	0x5e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x1b, 0x2e,
	0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48,
	0x65, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x61, 0x6e,
	0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x72, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x12,
	0x63, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x1b, 0x2e,
	0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x65, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x61, 0x6e,
	0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x72, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x3a, 0x01, 0x2a, 0x1a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x60, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65,
	0x72, 0x6f, 0x12, 0x1b, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x2a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x65,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x6f, 0x6c, 0x74, 0x64, 0x2f, 0x6c, 0x61, 0x6e, 0x63,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_lancr_v1_hero_service_proto_rawDescOnce sync.Once
	file_proto_lancr_v1_hero_service_proto_rawDescData = file_proto_lancr_v1_hero_service_proto_rawDesc
)

func file_proto_lancr_v1_hero_service_proto_rawDescGZIP() []byte {
	file_proto_lancr_v1_hero_service_proto_rawDescOnce.Do(func() {
		file_proto_lancr_v1_hero_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_lancr_v1_hero_service_proto_rawDescData)
	})
	return file_proto_lancr_v1_hero_service_proto_rawDescData
}

var file_proto_lancr_v1_hero_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_lancr_v1_hero_service_proto_goTypes = []interface{}{
	(*GetHeroRequest)(nil),     // 0: lancr.v1.GetHeroRequest
	(*GetHeroResponse)(nil),    // 1: lancr.v1.GetHeroResponse
	(*ListHeroesRequest)(nil),  // 2: lancr.v1.ListHeroesRequest
	(*ListHeroesResponse)(nil), // 3: lancr.v1.ListHeroesResponse
	(*CreateHeroRequest)(nil),  // 4: lancr.v1.CreateHeroRequest
	(*CreateHeroResponse)(nil), // 5: lancr.v1.CreateHeroResponse
	(*UpdateHeroRequest)(nil),  // 6: lancr.v1.UpdateHeroRequest
	(*UpdateHeroResponse)(nil), // 7: lancr.v1.UpdateHeroResponse
	(*DeleteHeroRequest)(nil),  // 8: lancr.v1.DeleteHeroRequest
	(*DeleteHeroResponse)(nil), // 9: lancr.v1.DeleteHeroResponse
	(*Hero)(nil),               // 10: lancr.v1.Hero
}
var file_proto_lancr_v1_hero_service_proto_depIdxs = []int32{
	10, // 0: lancr.v1.GetHeroResponse.hero:type_name -> lancr.v1.Hero
	10, // 1: lancr.v1.ListHeroesResponse.heroes:type_name -> lancr.v1.Hero
	10, // 2: lancr.v1.CreateHeroRequest.hero:type_name -> lancr.v1.Hero
	10, // 3: lancr.v1.CreateHeroResponse.hero:type_name -> lancr.v1.Hero
	10, // 4: lancr.v1.UpdateHeroRequest.hero:type_name -> lancr.v1.Hero
	10, // 5: lancr.v1.UpdateHeroResponse.hero:type_name -> lancr.v1.Hero
	0,  // 6: lancr.v1.HeroService.GetHero:input_type -> lancr.v1.GetHeroRequest
	2,  // 7: lancr.v1.HeroService.ListHeroes:input_type -> lancr.v1.ListHeroesRequest
	4,  // 8: lancr.v1.HeroService.CreateHero:input_type -> lancr.v1.CreateHeroRequest
	6,  // 9: lancr.v1.HeroService.UpdateHero:input_type -> lancr.v1.UpdateHeroRequest
	8,  // 10: lancr.v1.HeroService.DeleteHero:input_type -> lancr.v1.DeleteHeroRequest
	1,  // 11: lancr.v1.HeroService.GetHero:output_type -> lancr.v1.GetHeroResponse
	3,  // 12: lancr.v1.HeroService.ListHeroes:output_type -> lancr.v1.ListHeroesResponse
	5,  // 13: lancr.v1.HeroService.CreateHero:output_type -> lancr.v1.CreateHeroResponse
	7,  // 14: lancr.v1.HeroService.UpdateHero:output_type -> lancr.v1.UpdateHeroResponse
	9,  // 15: lancr.v1.HeroService.DeleteHero:output_type -> lancr.v1.DeleteHeroResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_lancr_v1_hero_service_proto_init() }
func file_proto_lancr_v1_hero_service_proto_init() {
	if File_proto_lancr_v1_hero_service_proto != nil {
		return
	}
	file_proto_lancr_v1_hero_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_lancr_v1_hero_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHeroRequest); i {
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
		file_proto_lancr_v1_hero_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHeroResponse); i {
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
		file_proto_lancr_v1_hero_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHeroesRequest); i {
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
		file_proto_lancr_v1_hero_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHeroesResponse); i {
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
		file_proto_lancr_v1_hero_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHeroRequest); i {
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
		file_proto_lancr_v1_hero_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHeroResponse); i {
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
		file_proto_lancr_v1_hero_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHeroRequest); i {
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
		file_proto_lancr_v1_hero_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHeroResponse); i {
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
		file_proto_lancr_v1_hero_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHeroRequest); i {
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
		file_proto_lancr_v1_hero_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHeroResponse); i {
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
			RawDescriptor: file_proto_lancr_v1_hero_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_lancr_v1_hero_service_proto_goTypes,
		DependencyIndexes: file_proto_lancr_v1_hero_service_proto_depIdxs,
		MessageInfos:      file_proto_lancr_v1_hero_service_proto_msgTypes,
	}.Build()
	File_proto_lancr_v1_hero_service_proto = out.File
	file_proto_lancr_v1_hero_service_proto_rawDesc = nil
	file_proto_lancr_v1_hero_service_proto_goTypes = nil
	file_proto_lancr_v1_hero_service_proto_depIdxs = nil
}
