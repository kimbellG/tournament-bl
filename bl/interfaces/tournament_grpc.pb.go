// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package interfaces

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TournamentServiceClient is the client API for TournamentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TournamentServiceClient interface {
	SaveUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*SaveResponse, error)
	GetUserById(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error)
	DeleteUserByID(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SumToBalance(ctx context.Context, in *RequestToUpdateBalance, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type tournamentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTournamentServiceClient(cc grpc.ClientConnInterface) TournamentServiceClient {
	return &tournamentServiceClient{cc}
}

func (c *tournamentServiceClient) SaveUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/interfaces.TournamentService/SaveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tournamentServiceClient) GetUserById(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/interfaces.TournamentService/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tournamentServiceClient) DeleteUserByID(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/interfaces.TournamentService/DeleteUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tournamentServiceClient) SumToBalance(ctx context.Context, in *RequestToUpdateBalance, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/interfaces.TournamentService/SumToBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TournamentServiceServer is the server API for TournamentService service.
// All implementations must embed UnimplementedTournamentServiceServer
// for forward compatibility
type TournamentServiceServer interface {
	SaveUser(context.Context, *User) (*SaveResponse, error)
	GetUserById(context.Context, *UserRequest) (*User, error)
	DeleteUserByID(context.Context, *UserRequest) (*emptypb.Empty, error)
	SumToBalance(context.Context, *RequestToUpdateBalance) (*emptypb.Empty, error)
	mustEmbedUnimplementedTournamentServiceServer()
}

// UnimplementedTournamentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTournamentServiceServer struct {
}

func (UnimplementedTournamentServiceServer) SaveUser(context.Context, *User) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUser not implemented")
}
func (UnimplementedTournamentServiceServer) GetUserById(context.Context, *UserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedTournamentServiceServer) DeleteUserByID(context.Context, *UserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserByID not implemented")
}
func (UnimplementedTournamentServiceServer) SumToBalance(context.Context, *RequestToUpdateBalance) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumToBalance not implemented")
}
func (UnimplementedTournamentServiceServer) mustEmbedUnimplementedTournamentServiceServer() {}

// UnsafeTournamentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TournamentServiceServer will
// result in compilation errors.
type UnsafeTournamentServiceServer interface {
	mustEmbedUnimplementedTournamentServiceServer()
}

func RegisterTournamentServiceServer(s grpc.ServiceRegistrar, srv TournamentServiceServer) {
	s.RegisterService(&TournamentService_ServiceDesc, srv)
}

func _TournamentService_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TournamentServiceServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interfaces.TournamentService/SaveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TournamentServiceServer).SaveUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _TournamentService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TournamentServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interfaces.TournamentService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TournamentServiceServer).GetUserById(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TournamentService_DeleteUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TournamentServiceServer).DeleteUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interfaces.TournamentService/DeleteUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TournamentServiceServer).DeleteUserByID(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TournamentService_SumToBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestToUpdateBalance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TournamentServiceServer).SumToBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interfaces.TournamentService/SumToBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TournamentServiceServer).SumToBalance(ctx, req.(*RequestToUpdateBalance))
	}
	return interceptor(ctx, in, info, handler)
}

// TournamentService_ServiceDesc is the grpc.ServiceDesc for TournamentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TournamentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interfaces.TournamentService",
	HandlerType: (*TournamentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUser",
			Handler:    _TournamentService_SaveUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _TournamentService_GetUserById_Handler,
		},
		{
			MethodName: "DeleteUserByID",
			Handler:    _TournamentService_DeleteUserByID_Handler,
		},
		{
			MethodName: "SumToBalance",
			Handler:    _TournamentService_SumToBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interfaces/tournament.proto",
}
