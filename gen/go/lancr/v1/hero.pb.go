// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: lancr/v1/hero.proto

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

// Experience might be another service possibly since it requires separate
// CRUD and might be nice to have it separate for D&A workloads later on
type Experience struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Position  string `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Company   string `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	StartDate string `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"` // ISO date string
	EndDate   string `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *Experience) Reset() {
	*x = Experience{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_hero_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Experience) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Experience) ProtoMessage() {}

func (x *Experience) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_hero_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Experience.ProtoReflect.Descriptor instead.
func (*Experience) Descriptor() ([]byte, []int) {
	return file_lancr_v1_hero_proto_rawDescGZIP(), []int{0}
}

func (x *Experience) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Experience) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Experience) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *Experience) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Experience) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type Hero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// internal
	FirebaseId     string        `protobuf:"bytes,2,opt,name=firebase_id,json=firebaseId,proto3" json:"firebase_id,omitempty"`
	Name           *Name         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Visibility     string        `protobuf:"bytes,4,opt,name=visibility,proto3" json:"visibility,omitempty"`
	ProfilePicture string        `protobuf:"bytes,5,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	Email          string        `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber    string        `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Location       *Location     `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
	UserName       string        `protobuf:"bytes,10,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Bio            string        `protobuf:"bytes,11,opt,name=bio,proto3" json:"bio,omitempty"`
	WorkType       string        `protobuf:"bytes,12,opt,name=work_type,json=workType,proto3" json:"work_type,omitempty"`
	SubWorkType    string        `protobuf:"bytes,13,opt,name=sub_work_type,json=subWorkType,proto3" json:"sub_work_type,omitempty"`
	Twitter        string        `protobuf:"bytes,14,opt,name=twitter,proto3" json:"twitter,omitempty"`
	Linkedin       string        `protobuf:"bytes,15,opt,name=linkedin,proto3" json:"linkedin,omitempty"`
	Website        string        `protobuf:"bytes,16,opt,name=website,proto3" json:"website,omitempty"`
	Experience     []*Experience `protobuf:"bytes,17,rep,name=experience,proto3" json:"experience,omitempty"`
	Rating         float32       `protobuf:"fixed32,18,opt,name=rating,proto3" json:"rating,omitempty"`
	Xp             string        `protobuf:"bytes,19,opt,name=xp,proto3" json:"xp,omitempty"`
	Region         string        `protobuf:"bytes,20,opt,name=region,proto3" json:"region,omitempty"`
	Language       string        `protobuf:"bytes,21,opt,name=language,proto3" json:"language,omitempty"`
	Level          uint32        `protobuf:"varint,22,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *Hero) Reset() {
	*x = Hero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_hero_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hero) ProtoMessage() {}

func (x *Hero) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_hero_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hero.ProtoReflect.Descriptor instead.
func (*Hero) Descriptor() ([]byte, []int) {
	return file_lancr_v1_hero_proto_rawDescGZIP(), []int{1}
}

func (x *Hero) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Hero) GetFirebaseId() string {
	if x != nil {
		return x.FirebaseId
	}
	return ""
}

func (x *Hero) GetName() *Name {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Hero) GetVisibility() string {
	if x != nil {
		return x.Visibility
	}
	return ""
}

func (x *Hero) GetProfilePicture() string {
	if x != nil {
		return x.ProfilePicture
	}
	return ""
}

func (x *Hero) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Hero) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Hero) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Hero) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Hero) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *Hero) GetWorkType() string {
	if x != nil {
		return x.WorkType
	}
	return ""
}

func (x *Hero) GetSubWorkType() string {
	if x != nil {
		return x.SubWorkType
	}
	return ""
}

func (x *Hero) GetTwitter() string {
	if x != nil {
		return x.Twitter
	}
	return ""
}

func (x *Hero) GetLinkedin() string {
	if x != nil {
		return x.Linkedin
	}
	return ""
}

func (x *Hero) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *Hero) GetExperience() []*Experience {
	if x != nil {
		return x.Experience
	}
	return nil
}

func (x *Hero) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Hero) GetXp() string {
	if x != nil {
		return x.Xp
	}
	return ""
}

func (x *Hero) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Hero) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Hero) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

var File_lancr_v1_hero_proto protoreflect.FileDescriptor

var file_lancr_v1_hero_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x94, 0x01, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x22, 0x82, 0x05, 0x0a, 0x04, 0x48, 0x65, 0x72,
	0x6f, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x72,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69,
	0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x1b, 0x0a, 0x09,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x75, 0x62,
	0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x75, 0x62, 0x57, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x65,
	0x64, 0x69, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x65,
	0x64, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x34, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x11, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x78,
	0x70, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x78, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x6f, 0x6c,
	0x74, 0x64, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lancr_v1_hero_proto_rawDescOnce sync.Once
	file_lancr_v1_hero_proto_rawDescData = file_lancr_v1_hero_proto_rawDesc
)

func file_lancr_v1_hero_proto_rawDescGZIP() []byte {
	file_lancr_v1_hero_proto_rawDescOnce.Do(func() {
		file_lancr_v1_hero_proto_rawDescData = protoimpl.X.CompressGZIP(file_lancr_v1_hero_proto_rawDescData)
	})
	return file_lancr_v1_hero_proto_rawDescData
}

var file_lancr_v1_hero_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_lancr_v1_hero_proto_goTypes = []interface{}{
	(*Experience)(nil), // 0: lancr.v1.Experience
	(*Hero)(nil),       // 1: lancr.v1.Hero
	(*Name)(nil),       // 2: lancr.v1.Name
	(*Location)(nil),   // 3: lancr.v1.Location
}
var file_lancr_v1_hero_proto_depIdxs = []int32{
	2, // 0: lancr.v1.Hero.name:type_name -> lancr.v1.Name
	3, // 1: lancr.v1.Hero.location:type_name -> lancr.v1.Location
	0, // 2: lancr.v1.Hero.experience:type_name -> lancr.v1.Experience
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_lancr_v1_hero_proto_init() }
func file_lancr_v1_hero_proto_init() {
	if File_lancr_v1_hero_proto != nil {
		return
	}
	file_lancr_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lancr_v1_hero_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Experience); i {
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
		file_lancr_v1_hero_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hero); i {
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
			RawDescriptor: file_lancr_v1_hero_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lancr_v1_hero_proto_goTypes,
		DependencyIndexes: file_lancr_v1_hero_proto_depIdxs,
		MessageInfos:      file_lancr_v1_hero_proto_msgTypes,
	}.Build()
	File_lancr_v1_hero_proto = out.File
	file_lancr_v1_hero_proto_rawDesc = nil
	file_lancr_v1_hero_proto_goTypes = nil
	file_lancr_v1_hero_proto_depIdxs = nil
}
