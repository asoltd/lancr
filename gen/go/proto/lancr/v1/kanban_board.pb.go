// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/lancr/v1/kanban_board.proto

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

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Index       uint32                 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description *string                `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// TODO(oliwierost): the assignees should be migrated to IDs
	Assignees []string `protobuf:"bytes,7,rep,name=assignees,proto3" json:"assignees,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_kanban_board_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_kanban_board_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_kanban_board_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Task) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Task) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Task) GetAssignees() []string {
	if x != nil {
		return x.Assignees
	}
	return nil
}

type KanbanCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Index       uint32                 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description *string                `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Tasks       []*Task                `protobuf:"bytes,7,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *KanbanCategory) Reset() {
	*x = KanbanCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_kanban_board_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KanbanCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KanbanCategory) ProtoMessage() {}

func (x *KanbanCategory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_kanban_board_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KanbanCategory.ProtoReflect.Descriptor instead.
func (*KanbanCategory) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_kanban_board_proto_rawDescGZIP(), []int{1}
}

func (x *KanbanCategory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *KanbanCategory) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *KanbanCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KanbanCategory) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *KanbanCategory) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *KanbanCategory) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *KanbanCategory) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatorId     string                 `protobuf:"bytes,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Text          string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	EditHistories []string               `protobuf:"bytes,4,rep,name=edit_histories,json=editHistories,proto3" json:"edit_histories,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lancr_v1_kanban_board_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lancr_v1_kanban_board_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_proto_lancr_v1_kanban_board_proto_rawDescGZIP(), []int{2}
}

func (x *Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Comment) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Comment) GetEditHistories() []string {
	if x != nil {
		return x.EditHistories
	}
	return nil
}

func (x *Comment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Comment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_proto_lancr_v1_kanban_board_proto protoreflect.FileDescriptor

var file_proto_lancr_v1_kanban_board_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x6b, 0x61, 0x6e, 0x62, 0x61, 0x6e, 0x5f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b,
	0x02, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x73, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9d, 0x02, 0x0a,
	0x0e, 0x4b, 0x61, 0x6e, 0x62, 0x61, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x61, 0x6e, 0x63, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe9, 0x01, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x65,
	0x64, 0x69, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x64, 0x69, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x6f, 0x6c, 0x74, 0x64, 0x2f, 0x6c, 0x61,
	0x6e, 0x63, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x61, 0x6e, 0x63, 0x72,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_lancr_v1_kanban_board_proto_rawDescOnce sync.Once
	file_proto_lancr_v1_kanban_board_proto_rawDescData = file_proto_lancr_v1_kanban_board_proto_rawDesc
)

func file_proto_lancr_v1_kanban_board_proto_rawDescGZIP() []byte {
	file_proto_lancr_v1_kanban_board_proto_rawDescOnce.Do(func() {
		file_proto_lancr_v1_kanban_board_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_lancr_v1_kanban_board_proto_rawDescData)
	})
	return file_proto_lancr_v1_kanban_board_proto_rawDescData
}

var file_proto_lancr_v1_kanban_board_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_lancr_v1_kanban_board_proto_goTypes = []interface{}{
	(*Task)(nil),                  // 0: lancr.v1.Task
	(*KanbanCategory)(nil),        // 1: lancr.v1.KanbanCategory
	(*Comment)(nil),               // 2: lancr.v1.Comment
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_proto_lancr_v1_kanban_board_proto_depIdxs = []int32{
	3, // 0: lancr.v1.Task.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: lancr.v1.Task.updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: lancr.v1.KanbanCategory.created_at:type_name -> google.protobuf.Timestamp
	3, // 3: lancr.v1.KanbanCategory.updated_at:type_name -> google.protobuf.Timestamp
	0, // 4: lancr.v1.KanbanCategory.tasks:type_name -> lancr.v1.Task
	3, // 5: lancr.v1.Comment.created_at:type_name -> google.protobuf.Timestamp
	3, // 6: lancr.v1.Comment.updated_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_lancr_v1_kanban_board_proto_init() }
func file_proto_lancr_v1_kanban_board_proto_init() {
	if File_proto_lancr_v1_kanban_board_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_lancr_v1_kanban_board_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_proto_lancr_v1_kanban_board_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KanbanCategory); i {
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
		file_proto_lancr_v1_kanban_board_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
	file_proto_lancr_v1_kanban_board_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_lancr_v1_kanban_board_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_lancr_v1_kanban_board_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_lancr_v1_kanban_board_proto_goTypes,
		DependencyIndexes: file_proto_lancr_v1_kanban_board_proto_depIdxs,
		MessageInfos:      file_proto_lancr_v1_kanban_board_proto_msgTypes,
	}.Build()
	File_proto_lancr_v1_kanban_board_proto = out.File
	file_proto_lancr_v1_kanban_board_proto_rawDesc = nil
	file_proto_lancr_v1_kanban_board_proto_goTypes = nil
	file_proto_lancr_v1_kanban_board_proto_depIdxs = nil
}
