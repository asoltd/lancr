// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: lancr/v1/bid.proto

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

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BidderId       string                 `protobuf:"bytes,2,opt,name=bidder_id,json=bidderId,proto3" json:"bidder_id,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Rate           float32                `protobuf:"fixed32,5,opt,name=rate,proto3" json:"rate,omitempty"`
	Amount         float32                `protobuf:"fixed32,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency       string                 `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
	TimeRequired   string                 `protobuf:"bytes,8,opt,name=time_required,json=timeRequired,proto3" json:"time_required,omitempty"`
	WorkingTime    string                 `protobuf:"bytes,9,opt,name=working_time,json=workingTime,proto3" json:"working_time,omitempty"`
	QuestId        string                 `protobuf:"bytes,10,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
	Apprentice     *Apprentice            `protobuf:"bytes,11,opt,name=apprentice,proto3" json:"apprentice,omitempty"`
	ApprenticeRate float32                `protobuf:"fixed32,12,opt,name=apprentice_rate,json=apprenticeRate,proto3" json:"apprentice_rate,omitempty"`
	ApprenticeCut  float32                `protobuf:"fixed32,13,opt,name=apprentice_cut,json=apprenticeCut,proto3" json:"apprentice_cut,omitempty"`
	TotalEarnings  float32                `protobuf:"fixed32,14,opt,name=total_earnings,json=totalEarnings,proto3" json:"total_earnings,omitempty"`
	Status         string                 `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	TotalBidValue  float32                `protobuf:"fixed32,16,opt,name=total_bid_value,json=totalBidValue,proto3" json:"total_bid_value,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lancr_v1_bid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_lancr_v1_bid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_lancr_v1_bid_proto_rawDescGZIP(), []int{0}
}

func (x *Bid) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bid) GetBidderId() string {
	if x != nil {
		return x.BidderId
	}
	return ""
}

func (x *Bid) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Bid) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Bid) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *Bid) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Bid) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Bid) GetTimeRequired() string {
	if x != nil {
		return x.TimeRequired
	}
	return ""
}

func (x *Bid) GetWorkingTime() string {
	if x != nil {
		return x.WorkingTime
	}
	return ""
}

func (x *Bid) GetQuestId() string {
	if x != nil {
		return x.QuestId
	}
	return ""
}

func (x *Bid) GetApprentice() *Apprentice {
	if x != nil {
		return x.Apprentice
	}
	return nil
}

func (x *Bid) GetApprenticeRate() float32 {
	if x != nil {
		return x.ApprenticeRate
	}
	return 0
}

func (x *Bid) GetApprenticeCut() float32 {
	if x != nil {
		return x.ApprenticeCut
	}
	return 0
}

func (x *Bid) GetTotalEarnings() float32 {
	if x != nil {
		return x.TotalEarnings
	}
	return 0
}

func (x *Bid) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Bid) GetTotalBidValue() float32 {
	if x != nil {
		return x.TotalBidValue
	}
	return 0
}

var File_lancr_v1_bid_proto protoreflect.FileDescriptor

var file_lancr_v1_bid_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x04, 0x0a, 0x03, 0x42,
	0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x34, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x72, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x70, 0x70, 0x72, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x65, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x72,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x72, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0e, 0x61, 0x70, 0x70, 0x72, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x75,
	0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x72, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x65, 0x43, 0x75, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62,
	0x69, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x69, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x6f, 0x6c,
	0x74, 0x64, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lancr_v1_bid_proto_rawDescOnce sync.Once
	file_lancr_v1_bid_proto_rawDescData = file_lancr_v1_bid_proto_rawDesc
)

func file_lancr_v1_bid_proto_rawDescGZIP() []byte {
	file_lancr_v1_bid_proto_rawDescOnce.Do(func() {
		file_lancr_v1_bid_proto_rawDescData = protoimpl.X.CompressGZIP(file_lancr_v1_bid_proto_rawDescData)
	})
	return file_lancr_v1_bid_proto_rawDescData
}

var file_lancr_v1_bid_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lancr_v1_bid_proto_goTypes = []interface{}{
	(*Bid)(nil),                   // 0: lancr.v1.Bid
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*Apprentice)(nil),            // 2: lancr.v1.Apprentice
}
var file_lancr_v1_bid_proto_depIdxs = []int32{
	1, // 0: lancr.v1.Bid.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: lancr.v1.Bid.updated_at:type_name -> google.protobuf.Timestamp
	2, // 2: lancr.v1.Bid.apprentice:type_name -> lancr.v1.Apprentice
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_lancr_v1_bid_proto_init() }
func file_lancr_v1_bid_proto_init() {
	if File_lancr_v1_bid_proto != nil {
		return
	}
	file_lancr_v1_apprentice_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lancr_v1_bid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bid); i {
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
			RawDescriptor: file_lancr_v1_bid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lancr_v1_bid_proto_goTypes,
		DependencyIndexes: file_lancr_v1_bid_proto_depIdxs,
		MessageInfos:      file_lancr_v1_bid_proto_msgTypes,
	}.Build()
	File_lancr_v1_bid_proto = out.File
	file_lancr_v1_bid_proto_rawDesc = nil
	file_lancr_v1_bid_proto_goTypes = nil
	file_lancr_v1_bid_proto_depIdxs = nil
}
