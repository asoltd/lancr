// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: lancr/v1/hero_service.proto

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
	HeroService_ReadHero_FullMethodName        = "/lancr.v1.HeroService/ReadHero"
	HeroService_ListHeroes_FullMethodName      = "/lancr.v1.HeroService/ListHeroes"
	HeroService_CreateHero_FullMethodName      = "/lancr.v1.HeroService/CreateHero"
	HeroService_UpdateHero_FullMethodName      = "/lancr.v1.HeroService/UpdateHero"
	HeroService_DeleteHero_FullMethodName      = "/lancr.v1.HeroService/DeleteHero"
	HeroService_GetTotalHeroes_FullMethodName  = "/lancr.v1.HeroService/GetTotalHeroes"
	HeroService_GetQuestCreator_FullMethodName = "/lancr.v1.HeroService/GetQuestCreator"
)

// HeroServiceClient is the client API for HeroService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeroServiceClient interface {
	ReadHero(ctx context.Context, in *ReadHeroRequest, opts ...grpc.CallOption) (*ReadHeroResponse, error)
	ListHeroes(ctx context.Context, in *ListHeroesRequest, opts ...grpc.CallOption) (*ListHeroesResponse, error)
	CreateHero(ctx context.Context, in *CreateHeroRequest, opts ...grpc.CallOption) (*CreateHeroResponse, error)
	UpdateHero(ctx context.Context, in *UpdateHeroRequest, opts ...grpc.CallOption) (*UpdateHeroResponse, error)
	DeleteHero(ctx context.Context, in *DeleteHeroRequest, opts ...grpc.CallOption) (*DeleteHeroResponse, error)
	GetTotalHeroes(ctx context.Context, in *GetTotalHeroesRequest, opts ...grpc.CallOption) (*GetTotalHeroesResponse, error)
	GetQuestCreator(ctx context.Context, in *GetQuestCreatorRequest, opts ...grpc.CallOption) (*GetQuestCreatorResponse, error)
}

type heroServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeroServiceClient(cc grpc.ClientConnInterface) HeroServiceClient {
	return &heroServiceClient{cc}
}

func (c *heroServiceClient) ReadHero(ctx context.Context, in *ReadHeroRequest, opts ...grpc.CallOption) (*ReadHeroResponse, error) {
	out := new(ReadHeroResponse)
	err := c.cc.Invoke(ctx, HeroService_ReadHero_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heroServiceClient) ListHeroes(ctx context.Context, in *ListHeroesRequest, opts ...grpc.CallOption) (*ListHeroesResponse, error) {
	out := new(ListHeroesResponse)
	err := c.cc.Invoke(ctx, HeroService_ListHeroes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heroServiceClient) CreateHero(ctx context.Context, in *CreateHeroRequest, opts ...grpc.CallOption) (*CreateHeroResponse, error) {
	out := new(CreateHeroResponse)
	err := c.cc.Invoke(ctx, HeroService_CreateHero_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heroServiceClient) UpdateHero(ctx context.Context, in *UpdateHeroRequest, opts ...grpc.CallOption) (*UpdateHeroResponse, error) {
	out := new(UpdateHeroResponse)
	err := c.cc.Invoke(ctx, HeroService_UpdateHero_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heroServiceClient) DeleteHero(ctx context.Context, in *DeleteHeroRequest, opts ...grpc.CallOption) (*DeleteHeroResponse, error) {
	out := new(DeleteHeroResponse)
	err := c.cc.Invoke(ctx, HeroService_DeleteHero_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heroServiceClient) GetTotalHeroes(ctx context.Context, in *GetTotalHeroesRequest, opts ...grpc.CallOption) (*GetTotalHeroesResponse, error) {
	out := new(GetTotalHeroesResponse)
	err := c.cc.Invoke(ctx, HeroService_GetTotalHeroes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heroServiceClient) GetQuestCreator(ctx context.Context, in *GetQuestCreatorRequest, opts ...grpc.CallOption) (*GetQuestCreatorResponse, error) {
	out := new(GetQuestCreatorResponse)
	err := c.cc.Invoke(ctx, HeroService_GetQuestCreator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeroServiceServer is the server API for HeroService service.
// All implementations must embed UnimplementedHeroServiceServer
// for forward compatibility
type HeroServiceServer interface {
	ReadHero(context.Context, *ReadHeroRequest) (*ReadHeroResponse, error)
	ListHeroes(context.Context, *ListHeroesRequest) (*ListHeroesResponse, error)
	CreateHero(context.Context, *CreateHeroRequest) (*CreateHeroResponse, error)
	UpdateHero(context.Context, *UpdateHeroRequest) (*UpdateHeroResponse, error)
	DeleteHero(context.Context, *DeleteHeroRequest) (*DeleteHeroResponse, error)
	GetTotalHeroes(context.Context, *GetTotalHeroesRequest) (*GetTotalHeroesResponse, error)
	GetQuestCreator(context.Context, *GetQuestCreatorRequest) (*GetQuestCreatorResponse, error)
	mustEmbedUnimplementedHeroServiceServer()
}

// UnimplementedHeroServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHeroServiceServer struct {
}

func (UnimplementedHeroServiceServer) ReadHero(context.Context, *ReadHeroRequest) (*ReadHeroResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadHero not implemented")
}
func (UnimplementedHeroServiceServer) ListHeroes(context.Context, *ListHeroesRequest) (*ListHeroesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHeroes not implemented")
}
func (UnimplementedHeroServiceServer) CreateHero(context.Context, *CreateHeroRequest) (*CreateHeroResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHero not implemented")
}
func (UnimplementedHeroServiceServer) UpdateHero(context.Context, *UpdateHeroRequest) (*UpdateHeroResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHero not implemented")
}
func (UnimplementedHeroServiceServer) DeleteHero(context.Context, *DeleteHeroRequest) (*DeleteHeroResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHero not implemented")
}
func (UnimplementedHeroServiceServer) GetTotalHeroes(context.Context, *GetTotalHeroesRequest) (*GetTotalHeroesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalHeroes not implemented")
}
func (UnimplementedHeroServiceServer) GetQuestCreator(context.Context, *GetQuestCreatorRequest) (*GetQuestCreatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestCreator not implemented")
}
func (UnimplementedHeroServiceServer) mustEmbedUnimplementedHeroServiceServer() {}

// UnsafeHeroServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeroServiceServer will
// result in compilation errors.
type UnsafeHeroServiceServer interface {
	mustEmbedUnimplementedHeroServiceServer()
}

func RegisterHeroServiceServer(s grpc.ServiceRegistrar, srv HeroServiceServer) {
	s.RegisterService(&HeroService_ServiceDesc, srv)
}

func _HeroService_ReadHero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadHeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeroServiceServer).ReadHero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HeroService_ReadHero_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeroServiceServer).ReadHero(ctx, req.(*ReadHeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeroService_ListHeroes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHeroesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeroServiceServer).ListHeroes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HeroService_ListHeroes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeroServiceServer).ListHeroes(ctx, req.(*ListHeroesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeroService_CreateHero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeroServiceServer).CreateHero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HeroService_CreateHero_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeroServiceServer).CreateHero(ctx, req.(*CreateHeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeroService_UpdateHero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeroServiceServer).UpdateHero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HeroService_UpdateHero_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeroServiceServer).UpdateHero(ctx, req.(*UpdateHeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeroService_DeleteHero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeroServiceServer).DeleteHero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HeroService_DeleteHero_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeroServiceServer).DeleteHero(ctx, req.(*DeleteHeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeroService_GetTotalHeroes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTotalHeroesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeroServiceServer).GetTotalHeroes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HeroService_GetTotalHeroes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeroServiceServer).GetTotalHeroes(ctx, req.(*GetTotalHeroesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeroService_GetQuestCreator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestCreatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeroServiceServer).GetQuestCreator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HeroService_GetQuestCreator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeroServiceServer).GetQuestCreator(ctx, req.(*GetQuestCreatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HeroService_ServiceDesc is the grpc.ServiceDesc for HeroService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeroService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lancr.v1.HeroService",
	HandlerType: (*HeroServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadHero",
			Handler:    _HeroService_ReadHero_Handler,
		},
		{
			MethodName: "ListHeroes",
			Handler:    _HeroService_ListHeroes_Handler,
		},
		{
			MethodName: "CreateHero",
			Handler:    _HeroService_CreateHero_Handler,
		},
		{
			MethodName: "UpdateHero",
			Handler:    _HeroService_UpdateHero_Handler,
		},
		{
			MethodName: "DeleteHero",
			Handler:    _HeroService_DeleteHero_Handler,
		},
		{
			MethodName: "GetTotalHeroes",
			Handler:    _HeroService_GetTotalHeroes_Handler,
		},
		{
			MethodName: "GetQuestCreator",
			Handler:    _HeroService_GetQuestCreator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lancr/v1/hero_service.proto",
}
