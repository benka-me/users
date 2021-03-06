// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc-users.proto

package users

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("rpc-users.proto", fileDescriptor_896b3a902ed8a400) }

var fileDescriptor_896b3a902ed8a400 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2a, 0x48, 0xd6,
	0x2d, 0x2d, 0x4e, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4,
	0xb8, 0x91, 0xc4, 0x8c, 0x2e, 0x31, 0x72, 0xb1, 0x86, 0x82, 0xf8, 0x42, 0x46, 0x5c, 0x1c, 0x41,
	0xa9, 0xe9, 0x99, 0xc5, 0x25, 0xa9, 0x45, 0x42, 0x42, 0x7a, 0x10, 0x35, 0x30, 0x81, 0xa0, 0xd4,
	0x42, 0x29, 0x4c, 0xb1, 0x62, 0x21, 0x4d, 0x2e, 0x56, 0x9f, 0xfc, 0xf4, 0xcc, 0x3c, 0x21, 0x7e,
	0xa8, 0x24, 0x98, 0x07, 0x52, 0x8d, 0x26, 0x50, 0x2c, 0xa4, 0xc8, 0xc5, 0xec, 0x9e, 0x5a, 0x22,
	0xc4, 0x0b, 0x15, 0x77, 0x4f, 0x2d, 0x01, 0x29, 0xe3, 0x86, 0x72, 0x5d, 0x12, 0x4b, 0x12, 0x85,
	0x94, 0xb8, 0xd8, 0xdc, 0x53, 0x4b, 0x1c, 0x73, 0x72, 0x84, 0x78, 0xa0, 0xc2, 0xae, 0xb9, 0x05,
	0x25, 0x95, 0x52, 0x5c, 0x50, 0x1e, 0x48, 0x46, 0x99, 0x8b, 0xc5, 0xb1, 0xb4, 0x24, 0x03, 0xae,
	0x22, 0x24, 0x3f, 0x3b, 0x35, 0x4f, 0x0a, 0x66, 0xaa, 0x67, 0x31, 0x48, 0xd2, 0x29, 0xe6, 0xc2,
	0x43, 0x39, 0x86, 0x1b, 0x0f, 0xe5, 0x18, 0x3e, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8,
	0xe2, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7,
	0xf8, 0xe2, 0x91, 0x1c, 0xc3, 0x87, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c,
	0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x5a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x7e, 0x52, 0x6a, 0x5e, 0x76, 0xa2, 0x6e, 0x6e, 0xaa, 0x3e, 0xd8, 0x5c, 0xfd, 0xf4,
	0x7c, 0xdd, 0x82, 0xec, 0x74, 0x08, 0x27, 0x89, 0x0d, 0x1c, 0x72, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x32, 0xca, 0xe2, 0xed, 0x60, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*Data, error)
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*All, error)
	Auth(ctx context.Context, in *Token, opts ...grpc.CallOption) (*IsAuth, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error) {
	out := new(RegisterRes)
	err := c.cc.Invoke(ctx, "/users.Users/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, "/users.Users/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/users.Users/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*All, error) {
	out := new(All)
	err := c.cc.Invoke(ctx, "/users.Users/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Auth(ctx context.Context, in *Token, opts ...grpc.CallOption) (*IsAuth, error) {
	out := new(IsAuth)
	err := c.cc.Invoke(ctx, "/users.Users/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	Register(context.Context, *RegisterReq) (*RegisterRes, error)
	Login(context.Context, *LoginReq) (*LoginRes, error)
	Get(context.Context, *GetReq) (*Data, error)
	GetAll(context.Context, *Empty) (*All, error)
	Auth(context.Context, *Token) (*IsAuth, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) Register(ctx context.Context, req *RegisterReq) (*RegisterRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedUsersServer) Login(ctx context.Context, req *LoginReq) (*LoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUsersServer) Get(ctx context.Context, req *GetReq) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedUsersServer) GetAll(ctx context.Context, req *Empty) (*All, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedUsersServer) Auth(ctx context.Context, req *Token) (*IsAuth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Auth(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Users_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Users_Login_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Users_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Users_GetAll_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _Users_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc-users.proto",
}
