// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: lancr/v1/course.proto

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

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Xp            int32                  `protobuf:"varint,2,opt,name=xp,proto3" json:"xp,omitempty"`
	Verified      bool                   `protobuf:"varint,3,opt,name=verified,proto3" json:"verified,omitempty"`
	Title         string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Description   *string                `protobuf:"bytes,5,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Image         *Image                 `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	Provider      string                 `protobuf:"bytes,7,opt,name=provider,proto3" json:"provider,omitempty"`
	CreatorId     string                 `protobuf:"bytes,8,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Rating        int32                  `protobuf:"varint,9,opt,name=rating,proto3" json:"rating,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	EnrolledUsers []*Hero                `protobuf:"bytes,11,rep,name=enrolled_users,json=enrolledUsers,proto3" json:"enrolled_users,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_lancr_v1_course_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Course) GetXp() int32 {
	if x != nil {
		return x.Xp
	}
	return 0
}

func (x *Course) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *Course) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Course) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Course) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Course) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Course) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Course) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Course) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Course) GetEnrolledUsers() []*Hero {
	if x != nil {
		return x.EnrolledUsers
	}
	return nil
}

var File_lancr_v1_course_proto protoreflect.FileDescriptor

var file_lancr_v1_course_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6c, 0x61, 0x6e, 0x63, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd,
	0x02, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x78, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x78, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x0e, 0x65, 0x6e, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x52,
	0x0d, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x6f,
	0x6c, 0x74, 0x64, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_lancr_v1_course_proto_rawDescOnce sync.Once
	file_lancr_v1_course_proto_rawDescData = file_lancr_v1_course_proto_rawDesc
)

func file_lancr_v1_course_proto_rawDescGZIP() []byte {
	file_lancr_v1_course_proto_rawDescOnce.Do(func() {
		file_lancr_v1_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_lancr_v1_course_proto_rawDescData)
	})
	return file_lancr_v1_course_proto_rawDescData
}

var file_lancr_v1_course_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lancr_v1_course_proto_goTypes = []interface{}{
	(*Course)(nil),                // 0: lancr.v1.Course
	(*Image)(nil),                 // 1: lancr.v1.Image
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*Hero)(nil),                  // 3: lancr.v1.Hero
}
var file_lancr_v1_course_proto_depIdxs = []int32{
	1, // 0: lancr.v1.Course.image:type_name -> lancr.v1.Image
	2, // 1: lancr.v1.Course.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: lancr.v1.Course.enrolled_users:type_name -> lancr.v1.Hero
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_lancr_v1_course_proto_init() }
func file_lancr_v1_course_proto_init() {
	if File_lancr_v1_course_proto != nil {
		return
	}
	file_lancr_v1_common_proto_init()
	file_lancr_v1_hero_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lancr_v1_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
	file_lancr_v1_course_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lancr_v1_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lancr_v1_course_proto_goTypes,
		DependencyIndexes: file_lancr_v1_course_proto_depIdxs,
		MessageInfos:      file_lancr_v1_course_proto_msgTypes,
	}.Build()
	File_lancr_v1_course_proto = out.File
	file_lancr_v1_course_proto_rawDesc = nil
	file_lancr_v1_course_proto_goTypes = nil
	file_lancr_v1_course_proto_depIdxs = nil
}
