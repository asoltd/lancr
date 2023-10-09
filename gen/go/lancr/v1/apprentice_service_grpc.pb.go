// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: lancr/v1/apprentice_service.proto

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
	ApprenticeService_ReadApprentice_FullMethodName   = "/lancr.v1.ApprenticeService/ReadApprentice"
	ApprenticeService_ListApprentices_FullMethodName  = "/lancr.v1.ApprenticeService/ListApprentices"
	ApprenticeService_CreateApprentice_FullMethodName = "/lancr.v1.ApprenticeService/CreateApprentice"
	ApprenticeService_UpdateApprentice_FullMethodName = "/lancr.v1.ApprenticeService/UpdateApprentice"
	ApprenticeService_DeleteApprentice_FullMethodName = "/lancr.v1.ApprenticeService/DeleteApprentice"
)

// ApprenticeServiceClient is the client API for ApprenticeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApprenticeServiceClient interface {
	ReadApprentice(ctx context.Context, in *ReadApprenticeRequest, opts ...grpc.CallOption) (*ReadApprenticeResponse, error)
	ListApprentices(ctx context.Context, in *ListApprenticesRequest, opts ...grpc.CallOption) (*ListApprenticesResponse, error)
	CreateApprentice(ctx context.Context, in *CreateApprenticeRequest, opts ...grpc.CallOption) (*CreateApprenticeResponse, error)
	UpdateApprentice(ctx context.Context, in *UpdateApprenticeRequest, opts ...grpc.CallOption) (*UpdateApprenticeResponse, error)
	DeleteApprentice(ctx context.Context, in *DeleteApprenticeRequest, opts ...grpc.CallOption) (*DeleteApprenticeResponse, error)
}

type apprenticeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApprenticeServiceClient(cc grpc.ClientConnInterface) ApprenticeServiceClient {
	return &apprenticeServiceClient{cc}
}

func (c *apprenticeServiceClient) ReadApprentice(ctx context.Context, in *ReadApprenticeRequest, opts ...grpc.CallOption) (*ReadApprenticeResponse, error) {
	out := new(ReadApprenticeResponse)
	err := c.cc.Invoke(ctx, ApprenticeService_ReadApprentice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apprenticeServiceClient) ListApprentices(ctx context.Context, in *ListApprenticesRequest, opts ...grpc.CallOption) (*ListApprenticesResponse, error) {
	out := new(ListApprenticesResponse)
	err := c.cc.Invoke(ctx, ApprenticeService_ListApprentices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apprenticeServiceClient) CreateApprentice(ctx context.Context, in *CreateApprenticeRequest, opts ...grpc.CallOption) (*CreateApprenticeResponse, error) {
	out := new(CreateApprenticeResponse)
	err := c.cc.Invoke(ctx, ApprenticeService_CreateApprentice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apprenticeServiceClient) UpdateApprentice(ctx context.Context, in *UpdateApprenticeRequest, opts ...grpc.CallOption) (*UpdateApprenticeResponse, error) {
	out := new(UpdateApprenticeResponse)
	err := c.cc.Invoke(ctx, ApprenticeService_UpdateApprentice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apprenticeServiceClient) DeleteApprentice(ctx context.Context, in *DeleteApprenticeRequest, opts ...grpc.CallOption) (*DeleteApprenticeResponse, error) {
	out := new(DeleteApprenticeResponse)
	err := c.cc.Invoke(ctx, ApprenticeService_DeleteApprentice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApprenticeServiceServer is the server API for ApprenticeService service.
// All implementations must embed UnimplementedApprenticeServiceServer
// for forward compatibility
type ApprenticeServiceServer interface {
	ReadApprentice(context.Context, *ReadApprenticeRequest) (*ReadApprenticeResponse, error)
	ListApprentices(context.Context, *ListApprenticesRequest) (*ListApprenticesResponse, error)
	CreateApprentice(context.Context, *CreateApprenticeRequest) (*CreateApprenticeResponse, error)
	UpdateApprentice(context.Context, *UpdateApprenticeRequest) (*UpdateApprenticeResponse, error)
	DeleteApprentice(context.Context, *DeleteApprenticeRequest) (*DeleteApprenticeResponse, error)
	mustEmbedUnimplementedApprenticeServiceServer()
}

// UnimplementedApprenticeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApprenticeServiceServer struct {
}

func (UnimplementedApprenticeServiceServer) ReadApprentice(context.Context, *ReadApprenticeRequest) (*ReadApprenticeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadApprentice not implemented")
}
func (UnimplementedApprenticeServiceServer) ListApprentices(context.Context, *ListApprenticesRequest) (*ListApprenticesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApprentices not implemented")
}
func (UnimplementedApprenticeServiceServer) CreateApprentice(context.Context, *CreateApprenticeRequest) (*CreateApprenticeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApprentice not implemented")
}
func (UnimplementedApprenticeServiceServer) UpdateApprentice(context.Context, *UpdateApprenticeRequest) (*UpdateApprenticeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApprentice not implemented")
}
func (UnimplementedApprenticeServiceServer) DeleteApprentice(context.Context, *DeleteApprenticeRequest) (*DeleteApprenticeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApprentice not implemented")
}
func (UnimplementedApprenticeServiceServer) mustEmbedUnimplementedApprenticeServiceServer() {}

// UnsafeApprenticeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApprenticeServiceServer will
// result in compilation errors.
type UnsafeApprenticeServiceServer interface {
	mustEmbedUnimplementedApprenticeServiceServer()
}

func RegisterApprenticeServiceServer(s grpc.ServiceRegistrar, srv ApprenticeServiceServer) {
	s.RegisterService(&ApprenticeService_ServiceDesc, srv)
}

func _ApprenticeService_ReadApprentice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadApprenticeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprenticeServiceServer).ReadApprentice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprenticeService_ReadApprentice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprenticeServiceServer).ReadApprentice(ctx, req.(*ReadApprenticeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprenticeService_ListApprentices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApprenticesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprenticeServiceServer).ListApprentices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprenticeService_ListApprentices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprenticeServiceServer).ListApprentices(ctx, req.(*ListApprenticesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprenticeService_CreateApprentice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApprenticeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprenticeServiceServer).CreateApprentice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprenticeService_CreateApprentice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprenticeServiceServer).CreateApprentice(ctx, req.(*CreateApprenticeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprenticeService_UpdateApprentice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApprenticeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprenticeServiceServer).UpdateApprentice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprenticeService_UpdateApprentice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprenticeServiceServer).UpdateApprentice(ctx, req.(*UpdateApprenticeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApprenticeService_DeleteApprentice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApprenticeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApprenticeServiceServer).DeleteApprentice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApprenticeService_DeleteApprentice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApprenticeServiceServer).DeleteApprentice(ctx, req.(*DeleteApprenticeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApprenticeService_ServiceDesc is the grpc.ServiceDesc for ApprenticeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApprenticeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lancr.v1.ApprenticeService",
	HandlerType: (*ApprenticeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadApprentice",
			Handler:    _ApprenticeService_ReadApprentice_Handler,
		},
		{
			MethodName: "ListApprentices",
			Handler:    _ApprenticeService_ListApprentices_Handler,
		},
		{
			MethodName: "CreateApprentice",
			Handler:    _ApprenticeService_CreateApprentice_Handler,
		},
		{
			MethodName: "UpdateApprentice",
			Handler:    _ApprenticeService_UpdateApprentice_Handler,
		},
		{
			MethodName: "DeleteApprentice",
			Handler:    _ApprenticeService_DeleteApprentice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lancr/v1/apprentice_service.proto",
}