// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: lancr/v1/team.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// the categories are not exhaustive, but are meant to be a starting point
type RoleCategory int32

const (
	RoleCategory_ROLE_CATEGORY_DESIGNER_UNSPECIFIED RoleCategory = 0
	RoleCategory_ROLE_CATEGORY_PRODUCT_DESIGNER     RoleCategory = 1
	RoleCategory_ROLE_CATEGORY_SOFTWARE_DEVELOPMENT RoleCategory = 2
	RoleCategory_ROLE_CATEGORY_PRESENTOR            RoleCategory = 3
	RoleCategory_ROLE_CATEGORY_MANAGER              RoleCategory = 4
	RoleCategory_ROLE_CATEGORY_ENGINEER             RoleCategory = 5
	RoleCategory_ROLE_CATEGORY_UNSPECIFIED          RoleCategory = 6
)

// Enum value maps for RoleCategory.
var (
	RoleCategory_name = map[int32]string{
		0: "ROLE_CATEGORY_DESIGNER_UNSPECIFIED",
		1: "ROLE_CATEGORY_PRODUCT_DESIGNER",
		2: "ROLE_CATEGORY_SOFTWARE_DEVELOPMENT",
		3: "ROLE_CATEGORY_PRESENTOR",
		4: "ROLE_CATEGORY_MANAGER",
		5: "ROLE_CATEGORY_ENGINEER",
		6: "ROLE_CATEGORY_UNSPECIFIED",
	}
	RoleCategory_value = map[string]int32{
		"ROLE_CATEGORY_DESIGNER_UNSPECIFIED": 0,
		"ROLE_CATEGORY_PRODUCT_DESIGNER":     1,
		"ROLE_CATEGORY_SOFTWARE_DEVELOPMENT": 2,
		"ROLE_CATEGORY_PRESENTOR":            3,
		"ROLE_CATEGORY_MANAGER":              4,
		"ROLE_CATEGORY_ENGINEER":             5,
		"ROLE_CATEGORY_UNSPECIFIED":          6,
	}
)

func (x RoleCategory) Enum() *RoleCategory {
	p := new(RoleCategory)
	*p = x
	return p
}

func (x RoleCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoleCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_lancr_v1_team_proto_enumTypes[0].Descriptor()
}

func (RoleCategory) Type() protoreflect.EnumType {
	return &file_lancr_v1_team_proto_enumTypes[0]
}

func (x RoleCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoleCategory.Descriptor instead.
func (RoleCategory) EnumDescriptor() ([]byte, []int) {
	return file_lancr_v1_team_proto_rawDescGZIP(), []int{0}
}

// same in case of RoleCategory
type Industry int32

const (
	Industry_INDUSTRY_FINTECH_UNSPECIFIED Industry = 0
	Industry_INDUSTRY_TECHNOLOGY          Industry = 1
	Industry_INDUSTRY_ACCOUNTING          Industry = 2
	Industry_INDUSTRY_CONSTRUCTION        Industry = 3
)

// Enum value maps for Industry.
var (
	Industry_name = map[int32]string{
		0: "INDUSTRY_FINTECH_UNSPECIFIED",
		1: "INDUSTRY_TECHNOLOGY",
		2: "INDUSTRY_ACCOUNTING",
		3: "INDUSTRY_CONSTRUCTION",
	}
	Industry_value = map[string]int32{
		"INDUSTRY_FINTECH_UNSPECIFIED": 0,
		"INDUSTRY_TECHNOLOGY":          1,
		"INDUSTRY_ACCOUNTING":          2,
		"INDUSTRY_CONSTRUCTION":        3,
	}
)

func (x Industry) Enum() *Industry {
	p := new(Industry)
	*p = x
	return p
}

func (x Industry) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Industry) Descriptor() protoreflect.EnumDescriptor {
	return file_lancr_v1_team_proto_enumTypes[1].Descriptor()
}

func (Industry) Type() protoreflect.EnumType {
	return &file_lancr_v1_team_proto_enumTypes[1]
}

func (x Industry) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Industry.Descriptor instead.
func (Industry) EnumDescriptor() ([]byte, []int) {
	return file_lancr_v1_team_proto_rawDescGZIP(), []int{1}
}

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorId      string                 `protobuf:"bytes,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Title          *string                `protobuf:"bytes,3,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Description    *string                `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Highlight      *string                `protobuf:"bytes,5,opt,name=highlight,proto3,oneof" json:"highlight,omitempty"`
	Industry       *Industry              `protobuf:"varint,6,opt,name=industry,proto3,enum=lancr.v1.Industry,oneof" json:"industry,omitempty"`
	Image          *Image                 `protobuf:"bytes,7,opt,name=image,proto3,oneof" json:"image,omitempty"`
	TimeEstimate   *string                `protobuf:"bytes,8,opt,name=time_estimate,json=timeEstimate,proto3,oneof" json:"time_estimate,omitempty"`
	Bidders        []*Bid                 `protobuf:"bytes,9,rep,name=bidders,proto3" json:"bidders,omitempty"`
	Members        []*Hero                `protobuf:"bytes,10,rep,name=members,proto3" json:"members,omitempty"`
	RoleCategories []RoleCategory         `protobuf:"varint,11,rep,packed,name=role_categories,json=roleCategories,proto3,enum=lancr.v1.RoleCategory" json:"role_categories,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_team_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_team_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_lancr_v1_team_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Team) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Team) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *Team) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Team) GetHighlight() string {
	if x != nil && x.Highlight != nil {
		return *x.Highlight
	}
	return ""
}

func (x *Team) GetIndustry() Industry {
	if x != nil && x.Industry != nil {
		return *x.Industry
	}
	return Industry_INDUSTRY_FINTECH_UNSPECIFIED
}

func (x *Team) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Team) GetTimeEstimate() string {
	if x != nil && x.TimeEstimate != nil {
		return *x.TimeEstimate
	}
	return ""
}

func (x *Team) GetBidders() []*Bid {
	if x != nil {
		return x.Bidders
	}
	return nil
}

func (x *Team) GetMembers() []*Hero {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Team) GetRoleCategories() []RoleCategory {
	if x != nil {
		return x.RoleCategories
	}
	return nil
}

func (x *Team) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       *string                `protobuf:"bytes,2,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Description *string                `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Status      *string                `protobuf:"bytes,4,opt,name=status,proto3,oneof" json:"status,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Image       *Image                 `protobuf:"bytes,6,opt,name=image,proto3,oneof" json:"image,omitempty"`
	MemberId    string                 `protobuf:"bytes,7,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_team_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_team_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_lancr_v1_team_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *Role) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Role) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *Role) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Role) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Role) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

var File_lancr_v1_team_proto protoreflect.FileDescriptor

var file_lancr_v1_team_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6c, 0x61, 0x6e,
	0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc5, 0x04, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x68, 0x69,
	0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x09, 0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a,
	0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x75, 0x73,
	0x74, 0x72, 0x79, 0x48, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x88,
	0x01, 0x01, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x48, 0x04, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28,
	0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x73, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x07, 0x62, 0x69, 0x64, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x61, 0x6e, 0x63,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x07, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x72, 0x6f, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x3f, 0x0a, 0x0f, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0e, 0x72, 0x6f,
	0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x22, 0xa8, 0x02, 0x0a, 0x04, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x61, 0x6e,
	0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x03, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2a, 0xf5, 0x01, 0x0a, 0x0c, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x22, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x54,
	0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x44, 0x45, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x52, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e,
	0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x44, 0x45, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x52, 0x10, 0x01,
	0x12, 0x26, 0x0a, 0x22, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52,
	0x59, 0x5f, 0x53, 0x4f, 0x46, 0x54, 0x57, 0x41, 0x52, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x45, 0x4c,
	0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x4f, 0x4c, 0x45,
	0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x45, 0x4e,
	0x54, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x04,
	0x12, 0x1a, 0x0a, 0x16, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52,
	0x59, 0x5f, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x45, 0x52, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19,
	0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x06, 0x2a, 0x79, 0x0a, 0x08, 0x49,
	0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x1c, 0x49, 0x4e, 0x44, 0x55, 0x53,
	0x54, 0x52, 0x59, 0x5f, 0x46, 0x49, 0x4e, 0x54, 0x45, 0x43, 0x48, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x44,
	0x55, 0x53, 0x54, 0x52, 0x59, 0x5f, 0x54, 0x45, 0x43, 0x48, 0x4e, 0x4f, 0x4c, 0x4f, 0x47, 0x59,
	0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x4e, 0x44, 0x55, 0x53, 0x54, 0x52, 0x59, 0x5f, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x49,
	0x4e, 0x44, 0x55, 0x53, 0x54, 0x52, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x6f, 0x6c, 0x74, 0x64, 0x2f, 0x6c, 0x61, 0x6e, 0x63,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lancr_v1_team_proto_rawDescOnce sync.Once
	file_lancr_v1_team_proto_rawDescData = file_lancr_v1_team_proto_rawDesc
)

func file_lancr_v1_team_proto_rawDescGZIP() []byte {
	file_lancr_v1_team_proto_rawDescOnce.Do(func() {
		file_lancr_v1_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_lancr_v1_team_proto_rawDescData)
	})
	return file_lancr_v1_team_proto_rawDescData
}

var file_lancr_v1_team_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_lancr_v1_team_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lancr_v1_team_proto_goTypes = []interface{}{
	(RoleCategory)(0),             // 0: lancr.v1.RoleCategory
	(Industry)(0),                 // 1: lancr.v1.Industry
	(*Team)(nil),                  // 2: lancr.v1.Team
	(*Role)(nil),                  // 3: lancr.v1.Role
	(*Image)(nil),                 // 4: lancr.v1.Image
	(*Bid)(nil),                   // 5: lancr.v1.Bid
	(*Hero)(nil),                  // 6: lancr.v1.Hero
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_lancr_v1_team_proto_depIdxs = []int32{
	1, // 0: lancr.v1.Team.industry:type_name -> lancr.v1.Industry
	4, // 1: lancr.v1.Team.image:type_name -> lancr.v1.Image
	5, // 2: lancr.v1.Team.bidders:type_name -> lancr.v1.Bid
	6, // 3: lancr.v1.Team.members:type_name -> lancr.v1.Hero
	0, // 4: lancr.v1.Team.role_categories:type_name -> lancr.v1.RoleCategory
	7, // 5: lancr.v1.Team.created_at:type_name -> google.protobuf.Timestamp
	7, // 6: lancr.v1.Role.created_at:type_name -> google.protobuf.Timestamp
	4, // 7: lancr.v1.Role.image:type_name -> lancr.v1.Image
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_lancr_v1_team_proto_init() }
func file_lancr_v1_team_proto_init() {
	if File_lancr_v1_team_proto != nil {
		return
	}
	file_lancr_v1_common_proto_init()
	file_lancr_v1_bid_proto_init()
	file_lancr_v1_hero_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lancr_v1_team_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Team); i {
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
		file_lancr_v1_team_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
	file_lancr_v1_team_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_lancr_v1_team_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lancr_v1_team_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lancr_v1_team_proto_goTypes,
		DependencyIndexes: file_lancr_v1_team_proto_depIdxs,
		EnumInfos:         file_lancr_v1_team_proto_enumTypes,
		MessageInfos:      file_lancr_v1_team_proto_msgTypes,
	}.Build()
	File_lancr_v1_team_proto = out.File
	file_lancr_v1_team_proto_rawDesc = nil
	file_lancr_v1_team_proto_goTypes = nil
	file_lancr_v1_team_proto_depIdxs = nil
}
