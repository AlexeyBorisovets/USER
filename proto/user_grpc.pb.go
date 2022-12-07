// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/user.proto

package __

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

// USERClient is the client API for USER service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type USERClient interface {
	GetUserByUserType(ctx context.Context, in *GetUserByUserTypeRequest, opts ...grpc.CallOption) (*GetUserByUserTypeResponse, error)
	GetBalanceByID(ctx context.Context, in *GetBalanceByIDRequest, opts ...grpc.CallOption) (*GetBalanceByIDResponse, error)
	UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*Response, error)
	Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Response, error)
	Authentication(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
	RefreshMyTokens(ctx context.Context, in *RefreshTokensRequest, opts ...grpc.CallOption) (*RefreshTokensResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*Response, error)
}

type uSERClient struct {
	cc grpc.ClientConnInterface
}

func NewUSERClient(cc grpc.ClientConnInterface) USERClient {
	return &uSERClient{cc}
}

func (c *uSERClient) GetUserByUserType(ctx context.Context, in *GetUserByUserTypeRequest, opts ...grpc.CallOption) (*GetUserByUserTypeResponse, error) {
	out := new(GetUserByUserTypeResponse)
	err := c.cc.Invoke(ctx, "/USER/GetUserByUserType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERClient) GetBalanceByID(ctx context.Context, in *GetBalanceByIDRequest, opts ...grpc.CallOption) (*GetBalanceByIDResponse, error) {
	out := new(GetBalanceByIDResponse)
	err := c.cc.Invoke(ctx, "/USER/GetBalanceByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERClient) UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/USER/UpdateBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERClient) Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, "/USER/Registration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/USER/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERClient) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error) {
	out := new(GetAllUsersResponse)
	err := c.cc.Invoke(ctx, "/USER/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/USER/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/USER/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERClient) Authentication(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, "/USER/Authentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERClient) RefreshMyTokens(ctx context.Context, in *RefreshTokensRequest, opts ...grpc.CallOption) (*RefreshTokensResponse, error) {
	out := new(RefreshTokensResponse)
	err := c.cc.Invoke(ctx, "/USER/RefreshMyTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uSERClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/USER/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// USERServer is the server API for USER service.
// All implementations must embed UnimplementedUSERServer
// for forward compatibility
type USERServer interface {
	GetUserByUserType(context.Context, *GetUserByUserTypeRequest) (*GetUserByUserTypeResponse, error)
	GetBalanceByID(context.Context, *GetBalanceByIDRequest) (*GetBalanceByIDResponse, error)
	UpdateBalance(context.Context, *UpdateBalanceRequest) (*Response, error)
	Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*Response, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*Response, error)
	Authentication(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error)
	RefreshMyTokens(context.Context, *RefreshTokensRequest) (*RefreshTokensResponse, error)
	Logout(context.Context, *LogoutRequest) (*Response, error)
	mustEmbedUnimplementedUSERServer()
}

// UnimplementedUSERServer must be embedded to have forward compatible implementations.
type UnimplementedUSERServer struct {
}

func (UnimplementedUSERServer) GetUserByUserType(context.Context, *GetUserByUserTypeRequest) (*GetUserByUserTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUserType not implemented")
}
func (UnimplementedUSERServer) GetBalanceByID(context.Context, *GetBalanceByIDRequest) (*GetBalanceByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalanceByID not implemented")
}
func (UnimplementedUSERServer) UpdateBalance(context.Context, *UpdateBalanceRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBalance not implemented")
}
func (UnimplementedUSERServer) Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}
func (UnimplementedUSERServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUSERServer) GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUSERServer) DeleteUser(context.Context, *DeleteUserRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUSERServer) UpdateUser(context.Context, *UpdateUserRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUSERServer) Authentication(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authentication not implemented")
}
func (UnimplementedUSERServer) RefreshMyTokens(context.Context, *RefreshTokensRequest) (*RefreshTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshMyTokens not implemented")
}
func (UnimplementedUSERServer) Logout(context.Context, *LogoutRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUSERServer) mustEmbedUnimplementedUSERServer() {}

// UnsafeUSERServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to USERServer will
// result in compilation errors.
type UnsafeUSERServer interface {
	mustEmbedUnimplementedUSERServer()
}

func RegisterUSERServer(s grpc.ServiceRegistrar, srv USERServer) {
	s.RegisterService(&USER_ServiceDesc, srv)
}

func _USER_GetUserByUserType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUserTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServer).GetUserByUserType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/USER/GetUserByUserType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServer).GetUserByUserType(ctx, req.(*GetUserByUserTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _USER_GetBalanceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServer).GetBalanceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/USER/GetBalanceByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServer).GetBalanceByID(ctx, req.(*GetBalanceByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _USER_UpdateBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServer).UpdateBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/USER/UpdateBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServer).UpdateBalance(ctx, req.(*UpdateBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _USER_Registration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServer).Registration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/USER/Registration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServer).Registration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _USER_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/USER/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _USER_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/USER/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServer).GetAllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _USER_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/USER/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _USER_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/USER/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _USER_Authentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServer).Authentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/USER/Authentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServer).Authentication(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _USER_RefreshMyTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServer).RefreshMyTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/USER/RefreshMyTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServer).RefreshMyTokens(ctx, req.(*RefreshTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _USER_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(USERServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/USER/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(USERServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// USER_ServiceDesc is the grpc.ServiceDesc for USER service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var USER_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "USER",
	HandlerType: (*USERServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByUserType",
			Handler:    _USER_GetUserByUserType_Handler,
		},
		{
			MethodName: "GetBalanceByID",
			Handler:    _USER_GetBalanceByID_Handler,
		},
		{
			MethodName: "UpdateBalance",
			Handler:    _USER_UpdateBalance_Handler,
		},
		{
			MethodName: "Registration",
			Handler:    _USER_Registration_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _USER_GetUser_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _USER_GetAllUsers_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _USER_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _USER_UpdateUser_Handler,
		},
		{
			MethodName: "Authentication",
			Handler:    _USER_Authentication_Handler,
		},
		{
			MethodName: "RefreshMyTokens",
			Handler:    _USER_RefreshMyTokens_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _USER_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}