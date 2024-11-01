// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: lancr/v1/quest_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	QuestService_ReadQuest_FullMethodName      = "/lancr.v1.QuestService/ReadQuest"
	QuestService_ListQuests_FullMethodName     = "/lancr.v1.QuestService/ListQuests"
	QuestService_CreateQuest_FullMethodName    = "/lancr.v1.QuestService/CreateQuest"
	QuestService_UpdateQuest_FullMethodName    = "/lancr.v1.QuestService/UpdateQuest"
	QuestService_DeleteQuest_FullMethodName    = "/lancr.v1.QuestService/DeleteQuest"
	QuestService_GetTopQuesters_FullMethodName = "/lancr.v1.QuestService/GetTopQuesters"
)

// QuestServiceClient is the client API for QuestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestServiceClient interface {
	ReadQuest(ctx context.Context, in *ReadQuestRequest, opts ...grpc.CallOption) (*ReadQuestResponse, error)
	ListQuests(ctx context.Context, in *ListQuestsRequest, opts ...grpc.CallOption) (*ListQuestsResponse, error)
	CreateQuest(ctx context.Context, in *CreateQuestRequest, opts ...grpc.CallOption) (*CreateQuestResponse, error)
	UpdateQuest(ctx context.Context, in *UpdateQuestRequest, opts ...grpc.CallOption) (*UpdateQuestResponse, error)
	DeleteQuest(ctx context.Context, in *DeleteQuestRequest, opts ...grpc.CallOption) (*DeleteQuestResponse, error)
	GetTopQuesters(ctx context.Context, in *GetTopQuestersRequest, opts ...grpc.CallOption) (*GetTopQuestersResponse, error)
}

type questServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestServiceClient(cc grpc.ClientConnInterface) QuestServiceClient {
	return &questServiceClient{cc}
}

func (c *questServiceClient) ReadQuest(ctx context.Context, in *ReadQuestRequest, opts ...grpc.CallOption) (*ReadQuestResponse, error) {
	out := new(ReadQuestResponse)
	err := c.cc.Invoke(ctx, QuestService_ReadQuest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questServiceClient) ListQuests(ctx context.Context, in *ListQuestsRequest, opts ...grpc.CallOption) (*ListQuestsResponse, error) {
	out := new(ListQuestsResponse)
	err := c.cc.Invoke(ctx, QuestService_ListQuests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questServiceClient) CreateQuest(ctx context.Context, in *CreateQuestRequest, opts ...grpc.CallOption) (*CreateQuestResponse, error) {
	out := new(CreateQuestResponse)
	err := c.cc.Invoke(ctx, QuestService_CreateQuest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questServiceClient) UpdateQuest(ctx context.Context, in *UpdateQuestRequest, opts ...grpc.CallOption) (*UpdateQuestResponse, error) {
	out := new(UpdateQuestResponse)
	err := c.cc.Invoke(ctx, QuestService_UpdateQuest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questServiceClient) DeleteQuest(ctx context.Context, in *DeleteQuestRequest, opts ...grpc.CallOption) (*DeleteQuestResponse, error) {
	out := new(DeleteQuestResponse)
	err := c.cc.Invoke(ctx, QuestService_DeleteQuest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questServiceClient) GetTopQuesters(ctx context.Context, in *GetTopQuestersRequest, opts ...grpc.CallOption) (*GetTopQuestersResponse, error) {
	out := new(GetTopQuestersResponse)
	err := c.cc.Invoke(ctx, QuestService_GetTopQuesters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestServiceServer is the server API for QuestService service.
// All implementations must embed UnimplementedQuestServiceServer
// for forward compatibility
type QuestServiceServer interface {
	ReadQuest(context.Context, *ReadQuestRequest) (*ReadQuestResponse, error)
	ListQuests(context.Context, *ListQuestsRequest) (*ListQuestsResponse, error)
	CreateQuest(context.Context, *CreateQuestRequest) (*CreateQuestResponse, error)
	UpdateQuest(context.Context, *UpdateQuestRequest) (*UpdateQuestResponse, error)
	DeleteQuest(context.Context, *DeleteQuestRequest) (*DeleteQuestResponse, error)
	GetTopQuesters(context.Context, *GetTopQuestersRequest) (*GetTopQuestersResponse, error)
	mustEmbedUnimplementedQuestServiceServer()
}

// UnimplementedQuestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuestServiceServer struct {
}

func (UnimplementedQuestServiceServer) ReadQuest(context.Context, *ReadQuestRequest) (*ReadQuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadQuest not implemented")
}
func (UnimplementedQuestServiceServer) ListQuests(context.Context, *ListQuestsRequest) (*ListQuestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuests not implemented")
}
func (UnimplementedQuestServiceServer) CreateQuest(context.Context, *CreateQuestRequest) (*CreateQuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuest not implemented")
}
func (UnimplementedQuestServiceServer) UpdateQuest(context.Context, *UpdateQuestRequest) (*UpdateQuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuest not implemented")
}
func (UnimplementedQuestServiceServer) DeleteQuest(context.Context, *DeleteQuestRequest) (*DeleteQuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuest not implemented")
}
func (UnimplementedQuestServiceServer) GetTopQuesters(context.Context, *GetTopQuestersRequest) (*GetTopQuestersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopQuesters not implemented")
}
func (UnimplementedQuestServiceServer) mustEmbedUnimplementedQuestServiceServer() {}

// UnsafeQuestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestServiceServer will
// result in compilation errors.
type UnsafeQuestServiceServer interface {
	mustEmbedUnimplementedQuestServiceServer()
}

func RegisterQuestServiceServer(s grpc.ServiceRegistrar, srv QuestServiceServer) {
	s.RegisterService(&QuestService_ServiceDesc, srv)
}

func _QuestService_ReadQuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadQuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestServiceServer).ReadQuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestService_ReadQuest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestServiceServer).ReadQuest(ctx, req.(*ReadQuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestService_ListQuests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQuestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestServiceServer).ListQuests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestService_ListQuests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestServiceServer).ListQuests(ctx, req.(*ListQuestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestService_CreateQuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestServiceServer).CreateQuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestService_CreateQuest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestServiceServer).CreateQuest(ctx, req.(*CreateQuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestService_UpdateQuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestServiceServer).UpdateQuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestService_UpdateQuest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestServiceServer).UpdateQuest(ctx, req.(*UpdateQuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestService_DeleteQuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestServiceServer).DeleteQuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestService_DeleteQuest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestServiceServer).DeleteQuest(ctx, req.(*DeleteQuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestService_GetTopQuesters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopQuestersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestServiceServer).GetTopQuesters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestService_GetTopQuesters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestServiceServer).GetTopQuesters(ctx, req.(*GetTopQuestersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuestService_ServiceDesc is the grpc.ServiceDesc for QuestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lancr.v1.QuestService",
	HandlerType: (*QuestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadQuest",
			Handler:    _QuestService_ReadQuest_Handler,
		},
		{
			MethodName: "ListQuests",
			Handler:    _QuestService_ListQuests_Handler,
		},
		{
			MethodName: "CreateQuest",
			Handler:    _QuestService_CreateQuest_Handler,
		},
		{
			MethodName: "UpdateQuest",
			Handler:    _QuestService_UpdateQuest_Handler,
		},
		{
			MethodName: "DeleteQuest",
			Handler:    _QuestService_DeleteQuest_Handler,
		},
		{
			MethodName: "GetTopQuesters",
			Handler:    _QuestService_GetTopQuesters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lancr/v1/quest_service.proto",
}
