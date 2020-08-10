// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sessions.proto

package session

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SessionData struct {
	Cookies              string   `protobuf:"bytes,1,opt,name=Cookies,proto3" json:"Cookies,omitempty"`
	Csrf                 string   `protobuf:"bytes,2,opt,name=Csrf,proto3" json:"Csrf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionData) Reset()         { *m = SessionData{} }
func (m *SessionData) String() string { return proto.CompactTextString(m) }
func (*SessionData) ProtoMessage()    {}
func (*SessionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0475a55364240b58, []int{0}
}

func (m *SessionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionData.Unmarshal(m, b)
}
func (m *SessionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionData.Marshal(b, m, deterministic)
}
func (m *SessionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionData.Merge(m, src)
}
func (m *SessionData) XXX_Size() int {
	return xxx_messageInfo_SessionData.Size(m)
}
func (m *SessionData) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionData.DiscardUnknown(m)
}

var xxx_messageInfo_SessionData proto.InternalMessageInfo

func (m *SessionData) GetCookies() string {
	if m != nil {
		return m.Cookies
	}
	return ""
}

func (m *SessionData) GetCsrf() string {
	if m != nil {
		return m.Csrf
	}
	return ""
}

type UserId struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserId) Reset()         { *m = UserId{} }
func (m *UserId) String() string { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()    {}
func (*UserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_0475a55364240b58, []int{1}
}

func (m *UserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserId.Unmarshal(m, b)
}
func (m *UserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserId.Marshal(b, m, deterministic)
}
func (m *UserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserId.Merge(m, src)
}
func (m *UserId) XXX_Size() int {
	return xxx_messageInfo_UserId.Size(m)
}
func (m *UserId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserId.DiscardUnknown(m)
}

var xxx_messageInfo_UserId proto.InternalMessageInfo

func (m *UserId) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type Auth struct {
	Login                string   `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_0475a55364240b58, []int{2}
}

func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (m *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(m, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *Auth) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Register struct {
	Email                string   `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Surname              string   `protobuf:"bytes,4,opt,name=Surname,proto3" json:"Surname,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Date                 string   `protobuf:"bytes,6,opt,name=Date,proto3" json:"Date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_0475a55364240b58, []int{3}
}

func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (m *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(m, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

func (m *Register) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Register) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Register) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Register) GetSurname() string {
	if m != nil {
		return m.Surname
	}
	return ""
}

func (m *Register) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Register) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func init() {
	proto.RegisterType((*SessionData)(nil), "session.SessionData")
	proto.RegisterType((*UserId)(nil), "session.UserId")
	proto.RegisterType((*Auth)(nil), "session.Auth")
	proto.RegisterType((*Register)(nil), "session.Register")
}

func init() { proto.RegisterFile("sessions.proto", fileDescriptor_0475a55364240b58) }

var fileDescriptor_0475a55364240b58 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0xe9, 0xf7, 0x41, 0x81, 0xab, 0x60, 0xbc, 0x21, 0x66, 0xc2, 0x8a, 0xcc, 0xca, 0x55,
	0x17, 0xb8, 0x90, 0xe8, 0xca, 0xb4, 0x2e, 0x4c, 0x0c, 0x21, 0x25, 0x3e, 0xc0, 0x28, 0x57, 0xda,
	0x00, 0x1d, 0x33, 0x33, 0x0d, 0x2f, 0xe2, 0x9b, 0xfa, 0x02, 0x66, 0xfe, 0xb4, 0x71, 0xa1, 0x26,
	0xee, 0xce, 0xef, 0x66, 0xce, 0xe9, 0x9c, 0xdb, 0x81, 0xb1, 0x26, 0xad, 0x4b, 0x59, 0xe9, 0xe4,
	0x4d, 0x49, 0x23, 0xb1, 0x1f, 0x98, 0xdf, 0xc2, 0xc9, 0xda, 0xcb, 0x4c, 0x18, 0x81, 0x0c, 0xfa,
	0xa9, 0x94, 0xbb, 0x92, 0x34, 0x8b, 0x66, 0xd1, 0xe5, 0x30, 0x6f, 0x10, 0x11, 0xba, 0xa9, 0x56,
	0xaf, 0xec, 0x9f, 0x1b, 0x3b, 0xcd, 0x67, 0x10, 0x3f, 0x69, 0x52, 0x0f, 0x1b, 0xbc, 0x80, 0xb8,
	0x76, 0xca, 0xd9, 0x7a, 0x79, 0x20, 0xbe, 0x80, 0xee, 0x5d, 0x6d, 0x0a, 0x9c, 0x40, 0xef, 0x51,
	0x6e, 0xcb, 0x2a, 0xa4, 0x7a, 0xc0, 0x29, 0x0c, 0x56, 0x42, 0xeb, 0xa3, 0x54, 0x9b, 0x90, 0xdb,
	0x32, 0x7f, 0x8f, 0x60, 0x90, 0xd3, 0xb6, 0xd4, 0x86, 0x94, 0xb5, 0xdf, 0x1f, 0x44, 0xb9, 0x6f,
	0xec, 0x0e, 0x7e, 0xb3, 0xdb, 0xeb, 0x2e, 0xc5, 0x81, 0xd8, 0x7f, 0x7f, 0x5d, 0xab, 0x6d, 0xb9,
	0x75, 0xad, 0x2a, 0x3b, 0xee, 0xfa, 0x72, 0x01, 0x6d, 0xfe, 0xaa, 0x90, 0x15, 0xb1, 0x9e, 0xcf,
	0x77, 0x60, 0x33, 0x32, 0x61, 0x88, 0xc5, 0x3e, 0xc3, 0xea, 0xf9, 0x47, 0x04, 0xe3, 0xb0, 0xb0,
	0xb4, 0xa0, 0x97, 0x1d, 0x29, 0x9c, 0xc3, 0xd0, 0xd5, 0xb1, 0xab, 0xc0, 0x51, 0x12, 0x36, 0x9b,
	0xd8, 0xde, 0xd3, 0x49, 0x8b, 0x5f, 0xb6, 0xcc, 0x3b, 0x78, 0x03, 0xa3, 0x54, 0x91, 0x30, 0xb4,
	0xa4, 0xa3, 0xf3, 0x9d, 0xb7, 0x07, 0x9b, 0xd2, 0x3f, 0x7a, 0xaf, 0xe1, 0xd4, 0x7d, 0x3a, 0x4c,
	0xf1, 0xdb, 0x73, 0xd3, 0xb3, 0x76, 0xea, 0x7f, 0x11, 0xef, 0xe0, 0x02, 0x46, 0x19, 0xed, 0xc9,
	0xd0, 0x5f, 0x9d, 0xcf, 0xb1, 0x7b, 0x35, 0x57, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x42,
	0xc6, 0xc1, 0x47, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionCheckerClient is the client API for SessionChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionCheckerClient interface {
	LoginUser(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*SessionData, error)
	CreateNewUser(ctx context.Context, in *Register, opts ...grpc.CallOption) (*SessionData, error)
	CheckSession(ctx context.Context, in *SessionData, opts ...grpc.CallOption) (*UserId, error)
	DeleteSession(ctx context.Context, in *SessionData, opts ...grpc.CallOption) (*UserId, error)
}

type sessionCheckerClient struct {
	cc *grpc.ClientConn
}

func NewSessionCheckerClient(cc *grpc.ClientConn) SessionCheckerClient {
	return &sessionCheckerClient{cc}
}

func (c *sessionCheckerClient) LoginUser(ctx context.Context, in *Auth, opts ...grpc.CallOption) (*SessionData, error) {
	out := new(SessionData)
	err := c.cc.Invoke(ctx, "/session.SessionChecker/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionCheckerClient) CreateNewUser(ctx context.Context, in *Register, opts ...grpc.CallOption) (*SessionData, error) {
	out := new(SessionData)
	err := c.cc.Invoke(ctx, "/session.SessionChecker/CreateNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionCheckerClient) CheckSession(ctx context.Context, in *SessionData, opts ...grpc.CallOption) (*UserId, error) {
	out := new(UserId)
	err := c.cc.Invoke(ctx, "/session.SessionChecker/CheckSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionCheckerClient) DeleteSession(ctx context.Context, in *SessionData, opts ...grpc.CallOption) (*UserId, error) {
	out := new(UserId)
	err := c.cc.Invoke(ctx, "/session.SessionChecker/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionCheckerServer is the server API for SessionChecker service.
type SessionCheckerServer interface {
	LoginUser(context.Context, *Auth) (*SessionData, error)
	CreateNewUser(context.Context, *Register) (*SessionData, error)
	CheckSession(context.Context, *SessionData) (*UserId, error)
	DeleteSession(context.Context, *SessionData) (*UserId, error)
}

// UnimplementedSessionCheckerServer can be embedded to have forward compatible implementations.
type UnimplementedSessionCheckerServer struct {
}

func (*UnimplementedSessionCheckerServer) LoginUser(ctx context.Context, req *Auth) (*SessionData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (*UnimplementedSessionCheckerServer) CreateNewUser(ctx context.Context, req *Register) (*SessionData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewUser not implemented")
}
func (*UnimplementedSessionCheckerServer) CheckSession(ctx context.Context, req *SessionData) (*UserId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSession not implemented")
}
func (*UnimplementedSessionCheckerServer) DeleteSession(ctx context.Context, req *SessionData) (*UserId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}

func RegisterSessionCheckerServer(s *grpc.Server, srv SessionCheckerServer) {
	s.RegisterService(&_SessionChecker_serviceDesc, srv)
}

func _SessionChecker_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionCheckerServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.SessionChecker/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionCheckerServer).LoginUser(ctx, req.(*Auth))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionChecker_CreateNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Register)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionCheckerServer).CreateNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.SessionChecker/CreateNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionCheckerServer).CreateNewUser(ctx, req.(*Register))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionChecker_CheckSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionCheckerServer).CheckSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.SessionChecker/CheckSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionCheckerServer).CheckSession(ctx, req.(*SessionData))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionChecker_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionCheckerServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.SessionChecker/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionCheckerServer).DeleteSession(ctx, req.(*SessionData))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "session.SessionChecker",
	HandlerType: (*SessionCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginUser",
			Handler:    _SessionChecker_LoginUser_Handler,
		},
		{
			MethodName: "CreateNewUser",
			Handler:    _SessionChecker_CreateNewUser_Handler,
		},
		{
			MethodName: "CheckSession",
			Handler:    _SessionChecker_CheckSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _SessionChecker_DeleteSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sessions.proto",
}