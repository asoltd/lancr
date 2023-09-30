// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/lancr/v1/quest.proto

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

type Tag int32

const (
	Tag_TAG_DESIGN_UNSPECIFIED  Tag = 0
	Tag_TAG_MANAGEMENT          Tag = 1
	Tag_TAG_RESEARCH            Tag = 2
	Tag_TAG_PRESENTATION        Tag = 3
	Tag_TAG_SOFTWAREDEVELOPMENT Tag = 4
	Tag_TAG_CUSTOMERSUCCESS     Tag = 5
	Tag_TAG_LEADERSHIP          Tag = 6
)

// Enum value maps for Tag.
var (
	Tag_name = map[int32]string{
		0: "TAG_DESIGN_UNSPECIFIED",
		1: "TAG_MANAGEMENT",
		2: "TAG_RESEARCH",
		3: "TAG_PRESENTATION",
		4: "TAG_SOFTWAREDEVELOPMENT",
		5: "TAG_CUSTOMERSUCCESS",
		6: "TAG_LEADERSHIP",
	}
	Tag_value = map[string]int32{
		"TAG_DESIGN_UNSPECIFIED":  0,
		"TAG_MANAGEMENT":          1,
		"TAG_RESEARCH":            2,
		"TAG_PRESENTATION":        3,
		"TAG_SOFTWAREDEVELOPMENT": 4,
		"TAG_CUSTOMERSUCCESS":     5,
		"TAG_LEADERSHIP":          6,
	}
)

func (x Tag) Enum() *Tag {
	p := new(Tag)
	*p = x
	return p
}

func (x Tag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Tag) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_lancr_v1_quest_proto_enumTypes[0].Descriptor()
}

func (Tag) Type() protoreflect.EnumType {
	return &file_proto_lancr_v1_quest_proto_enumTypes[0]
}

func (x Tag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Tag.Descriptor instead.
func (Tag) EnumDescriptor() ([]byte, []int) {
	return file_proto_lancr_v1_quest_proto_rawDescGZIP(), []int{0}
}

type Quest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorId   string                 `protobuf:"bytes,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Reward      float32                `protobuf:"fixed32,3,opt,name=reward,proto3" json:"reward,omitempty"`
	Title       string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Description string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Tags        []Tag                  `protobuf:"varint,6,rep,packed,name=tags,proto3,enum=lancr.v1.Tag" json:"tags,omitempty"`
	Images      []*Image               `protobuf:"bytes,7,rep,name=images,proto3" json:"images,omitempty"`
	Bidders     []string               `protobuf:"bytes,8,rep,name=bidders,proto3" json:"bidders,omitempty"`
	Status      string                 `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Summary     string                 `protobuf:"bytes,11,opt,name=summary,proto3" json:"summary,omitempty"`
	Level       uint32                 `protobuf:"varint,12,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *Quest) Reset() {
	*x = Quest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_quest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quest) ProtoMessage() {}

func (x *Quest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_quest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quest.ProtoReflect.Descriptor instead.
func (*Quest) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_quest_proto_rawDescGZIP(), []int{0}
}

func (x *Quest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Quest) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Quest) GetReward() float32 {
	if x != nil {
		return x.Reward
	}
	return 0
}

func (x *Quest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Quest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Quest) GetTags() []Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Quest) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Quest) GetBidders() []string {
	if x != nil {
		return x.Bidders
	}
	return nil
}

func (x *Quest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Quest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Quest) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Quest) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

var File_proto_lancr_v1_quest_proto protoreflect.FileDescriptor

var file_proto_lancr_v1_quest_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x61,
	0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c,
	0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x02, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6c, 0x61,
	0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x27, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x69, 0x64,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x62, 0x69, 0x64, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2a, 0xa7, 0x01, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x1a,
	0x0a, 0x16, 0x54, 0x41, 0x47, 0x5f, 0x44, 0x45, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x41,
	0x47, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x10,
	0x0a, 0x0c, 0x54, 0x41, 0x47, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x10, 0x02,
	0x12, 0x14, 0x0a, 0x10, 0x54, 0x41, 0x47, 0x5f, 0x50, 0x52, 0x45, 0x53, 0x45, 0x4e, 0x54, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x41, 0x47, 0x5f, 0x53, 0x4f,
	0x46, 0x54, 0x57, 0x41, 0x52, 0x45, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x41, 0x47, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f,
	0x4d, 0x45, 0x52, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e,
	0x54, 0x41, 0x47, 0x5f, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x53, 0x48, 0x49, 0x50, 0x10, 0x06,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x73, 0x6f, 0x6c, 0x74, 0x64, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_lancr_v1_quest_proto_rawDescOnce sync.Once
	file_proto_lancr_v1_quest_proto_rawDescData = file_proto_lancr_v1_quest_proto_rawDesc
)

func file_proto_lancr_v1_quest_proto_rawDescGZIP() []byte {
	file_proto_lancr_v1_quest_proto_rawDescOnce.Do(func() {
		file_proto_lancr_v1_quest_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_lancr_v1_quest_proto_rawDescData)
	})
	return file_proto_lancr_v1_quest_proto_rawDescData
}

var file_proto_lancr_v1_quest_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_lancr_v1_quest_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_lancr_v1_quest_proto_goTypes = []interface{}{
	(Tag)(0),                      // 0: lancr.v1.Tag
	(*Quest)(nil),                 // 1: lancr.v1.Quest
	(*Image)(nil),                 // 2: lancr.v1.Image
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_proto_lancr_v1_quest_proto_depIdxs = []int32{
	0, // 0: lancr.v1.Quest.tags:type_name -> lancr.v1.Tag
	2, // 1: lancr.v1.Quest.images:type_name -> lancr.v1.Image
	3, // 2: lancr.v1.Quest.created_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_lancr_v1_quest_proto_init() }
func file_proto_lancr_v1_quest_proto_init() {
	if File_proto_lancr_v1_quest_proto != nil {
		return
	}
	file_proto_lancr_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_lancr_v1_quest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quest); i {
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
			RawDescriptor: file_proto_lancr_v1_quest_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_lancr_v1_quest_proto_goTypes,
		DependencyIndexes: file_proto_lancr_v1_quest_proto_depIdxs,
		EnumInfos:         file_proto_lancr_v1_quest_proto_enumTypes,
		MessageInfos:      file_proto_lancr_v1_quest_proto_msgTypes,
	}.Build()
	File_proto_lancr_v1_quest_proto = out.File
	file_proto_lancr_v1_quest_proto_rawDesc = nil
	file_proto_lancr_v1_quest_proto_goTypes = nil
	file_proto_lancr_v1_quest_proto_depIdxs = nil
}
