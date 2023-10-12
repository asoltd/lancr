// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: lancr/v1/teams_service.proto

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
	TeamsService_ReadTeam_FullMethodName   = "/lancr.v1.TeamsService/ReadTeam"
	TeamsService_ListTeams_FullMethodName  = "/lancr.v1.TeamsService/ListTeams"
	TeamsService_CreateTeam_FullMethodName = "/lancr.v1.TeamsService/CreateTeam"
	TeamsService_UpdateTeam_FullMethodName = "/lancr.v1.TeamsService/UpdateTeam"
	TeamsService_DeleteTeam_FullMethodName = "/lancr.v1.TeamsService/DeleteTeam"
)

// TeamsServiceClient is the client API for TeamsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamsServiceClient interface {
	ReadTeam(ctx context.Context, in *ReadTeamRequest, opts ...grpc.CallOption) (*ReadTeamResponse, error)
	ListTeams(ctx context.Context, in *ListTeamsRequest, opts ...grpc.CallOption) (*ListTeamsResponse, error)
	CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamResponse, error)
	UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamResponse, error)
	DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamResponse, error)
}

type teamsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamsServiceClient(cc grpc.ClientConnInterface) TeamsServiceClient {
	return &teamsServiceClient{cc}
}

func (c *teamsServiceClient) ReadTeam(ctx context.Context, in *ReadTeamRequest, opts ...grpc.CallOption) (*ReadTeamResponse, error) {
	out := new(ReadTeamResponse)
	err := c.cc.Invoke(ctx, TeamsService_ReadTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsServiceClient) ListTeams(ctx context.Context, in *ListTeamsRequest, opts ...grpc.CallOption) (*ListTeamsResponse, error) {
	out := new(ListTeamsResponse)
	err := c.cc.Invoke(ctx, TeamsService_ListTeams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsServiceClient) CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamResponse, error) {
	out := new(CreateTeamResponse)
	err := c.cc.Invoke(ctx, TeamsService_CreateTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsServiceClient) UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamResponse, error) {
	out := new(UpdateTeamResponse)
	err := c.cc.Invoke(ctx, TeamsService_UpdateTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsServiceClient) DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamResponse, error) {
	out := new(DeleteTeamResponse)
	err := c.cc.Invoke(ctx, TeamsService_DeleteTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamsServiceServer is the server API for TeamsService service.
// All implementations must embed UnimplementedTeamsServiceServer
// for forward compatibility
type TeamsServiceServer interface {
	ReadTeam(context.Context, *ReadTeamRequest) (*ReadTeamResponse, error)
	ListTeams(context.Context, *ListTeamsRequest) (*ListTeamsResponse, error)
	CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamResponse, error)
	UpdateTeam(context.Context, *UpdateTeamRequest) (*UpdateTeamResponse, error)
	DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamResponse, error)
	mustEmbedUnimplementedTeamsServiceServer()
}

// UnimplementedTeamsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeamsServiceServer struct {
}

func (UnimplementedTeamsServiceServer) ReadTeam(context.Context, *ReadTeamRequest) (*ReadTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadTeam not implemented")
}
func (UnimplementedTeamsServiceServer) ListTeams(context.Context, *ListTeamsRequest) (*ListTeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTeams not implemented")
}
func (UnimplementedTeamsServiceServer) CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeam not implemented")
}
func (UnimplementedTeamsServiceServer) UpdateTeam(context.Context, *UpdateTeamRequest) (*UpdateTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeam not implemented")
}
func (UnimplementedTeamsServiceServer) DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeam not implemented")
}
func (UnimplementedTeamsServiceServer) mustEmbedUnimplementedTeamsServiceServer() {}

// UnsafeTeamsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamsServiceServer will
// result in compilation errors.
type UnsafeTeamsServiceServer interface {
	mustEmbedUnimplementedTeamsServiceServer()
}

func RegisterTeamsServiceServer(s grpc.ServiceRegistrar, srv TeamsServiceServer) {
	s.RegisterService(&TeamsService_ServiceDesc, srv)
}

func _TeamsService_ReadTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).ReadTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamsService_ReadTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).ReadTeam(ctx, req.(*ReadTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamsService_ListTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).ListTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamsService_ListTeams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).ListTeams(ctx, req.(*ListTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamsService_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamsService_CreateTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).CreateTeam(ctx, req.(*CreateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamsService_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamsService_UpdateTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).UpdateTeam(ctx, req.(*UpdateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamsService_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamsService_DeleteTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).DeleteTeam(ctx, req.(*DeleteTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeamsService_ServiceDesc is the grpc.ServiceDesc for TeamsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeamsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lancr.v1.TeamsService",
	HandlerType: (*TeamsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadTeam",
			Handler:    _TeamsService_ReadTeam_Handler,
		},
		{
			MethodName: "ListTeams",
			Handler:    _TeamsService_ListTeams_Handler,
		},
		{
			MethodName: "CreateTeam",
			Handler:    _TeamsService_CreateTeam_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _TeamsService_UpdateTeam_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _TeamsService_DeleteTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lancr/v1/teams_service.proto",
}