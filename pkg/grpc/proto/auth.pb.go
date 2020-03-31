// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package proto_auth

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

type Credentials struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credentials) Reset()         { *m = Credentials{} }
func (m *Credentials) String() string { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()    {}
func (*Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Credentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credentials.Unmarshal(m, b)
}
func (m *Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credentials.Marshal(b, m, deterministic)
}
func (m *Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credentials.Merge(m, src)
}
func (m *Credentials) XXX_Size() int {
	return xxx_messageInfo_Credentials.Size(m)
}
func (m *Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Credentials proto.InternalMessageInfo

func (m *Credentials) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Credentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Session struct {
	// expiration time can be worked out client side since our jwt holds
	// the expiration time.
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Jwt                  string   `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
	RefreshToken         string   `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Session) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

func (m *Session) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type LogoutStatus struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutStatus) Reset()         { *m = LogoutStatus{} }
func (m *LogoutStatus) String() string { return proto.CompactTextString(m) }
func (*LogoutStatus) ProtoMessage()    {}
func (*LogoutStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *LogoutStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutStatus.Unmarshal(m, b)
}
func (m *LogoutStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutStatus.Marshal(b, m, deterministic)
}
func (m *LogoutStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutStatus.Merge(m, src)
}
func (m *LogoutStatus) XXX_Size() int {
	return xxx_messageInfo_LogoutStatus.Size(m)
}
func (m *LogoutStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutStatus.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutStatus proto.InternalMessageInfo

func (m *LogoutStatus) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *LogoutStatus) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LogoutStatus) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type JWT struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWT) Reset()         { *m = JWT{} }
func (m *JWT) String() string { return proto.CompactTextString(m) }
func (*JWT) ProtoMessage()    {}
func (*JWT) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *JWT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWT.Unmarshal(m, b)
}
func (m *JWT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWT.Marshal(b, m, deterministic)
}
func (m *JWT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWT.Merge(m, src)
}
func (m *JWT) XXX_Size() int {
	return xxx_messageInfo_JWT.Size(m)
}
func (m *JWT) XXX_DiscardUnknown() {
	xxx_messageInfo_JWT.DiscardUnknown(m)
}

var xxx_messageInfo_JWT proto.InternalMessageInfo

func (m *JWT) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidityStatus struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidityStatus) Reset()         { *m = ValidityStatus{} }
func (m *ValidityStatus) String() string { return proto.CompactTextString(m) }
func (*ValidityStatus) ProtoMessage()    {}
func (*ValidityStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *ValidityStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidityStatus.Unmarshal(m, b)
}
func (m *ValidityStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidityStatus.Marshal(b, m, deterministic)
}
func (m *ValidityStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidityStatus.Merge(m, src)
}
func (m *ValidityStatus) XXX_Size() int {
	return xxx_messageInfo_ValidityStatus.Size(m)
}
func (m *ValidityStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidityStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ValidityStatus proto.InternalMessageInfo

func (m *ValidityStatus) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type Permission struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PermLvl              int32    `protobuf:"varint,2,opt,name=perm_lvl,json=permLvl,proto3" json:"perm_lvl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
}

func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Permission) GetPermLvl() int32 {
	if m != nil {
		return m.PermLvl
	}
	return 0
}

type CheckStatus struct {
	HasPermission        bool     `protobuf:"varint,1,opt,name=HasPermission,proto3" json:"HasPermission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckStatus) Reset()         { *m = CheckStatus{} }
func (m *CheckStatus) String() string { return proto.CompactTextString(m) }
func (*CheckStatus) ProtoMessage()    {}
func (*CheckStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6}
}

func (m *CheckStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckStatus.Unmarshal(m, b)
}
func (m *CheckStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckStatus.Marshal(b, m, deterministic)
}
func (m *CheckStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckStatus.Merge(m, src)
}
func (m *CheckStatus) XXX_Size() int {
	return xxx_messageInfo_CheckStatus.Size(m)
}
func (m *CheckStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckStatus.DiscardUnknown(m)
}

var xxx_messageInfo_CheckStatus proto.InternalMessageInfo

func (m *CheckStatus) GetHasPermission() bool {
	if m != nil {
		return m.HasPermission
	}
	return false
}

func init() {
	proto.RegisterType((*Credentials)(nil), "proto.auth.Credentials")
	proto.RegisterType((*Session)(nil), "proto.auth.Session")
	proto.RegisterType((*LogoutStatus)(nil), "proto.auth.LogoutStatus")
	proto.RegisterType((*JWT)(nil), "proto.auth.JWT")
	proto.RegisterType((*ValidityStatus)(nil), "proto.auth.ValidityStatus")
	proto.RegisterType((*Permission)(nil), "proto.auth.Permission")
	proto.RegisterType((*CheckStatus)(nil), "proto.auth.CheckStatus")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xcf, 0x6f, 0xda, 0x30,
	0x18, 0x15, 0x43, 0x21, 0xd9, 0xc7, 0x8f, 0x4d, 0x1e, 0x1a, 0x59, 0x76, 0x99, 0xb2, 0x69, 0xda,
	0x29, 0x07, 0xd0, 0xb4, 0xa9, 0xa7, 0x56, 0xa8, 0x52, 0x41, 0x1c, 0xaa, 0x80, 0x8a, 0x7a, 0x42,
	0x6e, 0xe2, 0x12, 0x97, 0x24, 0x46, 0xb6, 0x03, 0xea, 0x9f, 0xdd, 0xff, 0xa0, 0xb2, 0x1d, 0x9a,
	0x50, 0x51, 0xf5, 0x14, 0xbf, 0xf7, 0xf9, 0x7b, 0x2f, 0x7e, 0x0f, 0x00, 0x17, 0x32, 0x09, 0xb6,
	0x9c, 0x49, 0x86, 0x40, 0x7f, 0x02, 0xc5, 0xf8, 0x97, 0xd0, 0x1e, 0x73, 0x12, 0x93, 0x5c, 0x52,
	0x9c, 0x0a, 0xe4, 0x81, 0x53, 0x08, 0xc2, 0x73, 0x9c, 0x11, 0xb7, 0xf1, 0xa3, 0xf1, 0xe7, 0x63,
	0xf8, 0x82, 0xd5, 0x6c, 0x8b, 0x85, 0xd8, 0x33, 0x1e, 0xbb, 0x1f, 0xcc, 0xec, 0x80, 0xfd, 0x5b,
	0xb0, 0xe7, 0x44, 0x08, 0xca, 0x72, 0x34, 0x00, 0x5b, 0xad, 0xac, 0x68, 0x5c, 0x2a, 0xb4, 0x14,
	0x9c, 0xc4, 0xe8, 0x33, 0x34, 0x1f, 0xf6, 0xb2, 0x5c, 0x55, 0x47, 0xf4, 0x13, 0xba, 0x9c, 0xdc,
	0x73, 0x22, 0x92, 0x95, 0x64, 0x1b, 0x92, 0xbb, 0x4d, 0x3d, 0xeb, 0x94, 0xe4, 0x42, 0x71, 0xfe,
	0x1c, 0x3a, 0x33, 0xb6, 0x66, 0x85, 0x9c, 0x4b, 0x2c, 0x0b, 0xf1, 0xb6, 0xbe, 0x0b, 0xb6, 0x28,
	0xa2, 0x88, 0x08, 0xa1, 0x3d, 0x9c, 0xf0, 0x00, 0x95, 0x73, 0x26, 0xd6, 0xa5, 0xba, 0x3a, 0xfa,
	0xdf, 0xa1, 0x39, 0x5d, 0x2e, 0x50, 0x1f, 0x2c, 0x63, 0x6c, 0x94, 0x0c, 0xf0, 0x7f, 0x43, 0xef,
	0x06, 0xa7, 0x34, 0xa6, 0xf2, 0xb1, 0xf4, 0xec, 0x83, 0xb5, 0x53, 0x8c, 0xbe, 0xe7, 0x84, 0x06,
	0xf8, 0xe7, 0x00, 0xd7, 0x84, 0x67, 0xf4, 0x9d, 0x77, 0x7f, 0x03, 0x67, 0x4b, 0x78, 0xb6, 0x4a,
	0x77, 0xa9, 0xfe, 0x31, 0x2b, 0xb4, 0x15, 0x9e, 0xed, 0x52, 0x7f, 0x04, 0xed, 0x71, 0x42, 0xa2,
	0x4d, 0x69, 0xf3, 0x0b, 0xba, 0x57, 0x58, 0x54, 0x9a, 0xa5, 0xdd, 0x31, 0x39, 0x7c, 0x6a, 0x40,
	0xef, 0xa2, 0x90, 0x89, 0xea, 0x2c, 0xc2, 0x52, 0x79, 0xff, 0x05, 0x6b, 0xc6, 0xd6, 0x34, 0x47,
	0x83, 0xa0, 0xea, 0x36, 0xa8, 0x15, 0xeb, 0x7d, 0xa9, 0x0f, 0x0e, 0x55, 0x8d, 0xc0, 0x0e, 0x4d,
	0xd4, 0xe8, 0xd4, 0xfc, 0xf4, 0xd2, 0x19, 0xb4, 0x75, 0x3a, 0x58, 0x12, 0x15, 0xe1, 0xa7, 0xfa,
	0x9d, 0xe9, 0x72, 0xe1, 0x79, 0x75, 0xe2, 0x55, 0x8e, 0xff, 0xa0, 0x65, 0xba, 0x3c, 0xed, 0xe7,
	0xd6, 0xc9, 0x7a, 0xe9, 0xc3, 0x09, 0x74, 0xd5, 0x93, 0x19, 0xa7, 0xc2, 0xbc, 0xf8, 0x3f, 0x58,
	0x3a, 0x39, 0xf4, 0xb5, 0xbe, 0x53, 0xa5, 0xe4, 0x1d, 0x27, 0x51, 0x85, 0x7c, 0xd7, 0xd2, 0xfc,
	0xe8, 0x39, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x5e, 0x05, 0xfe, 0x12, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationClient interface {
	Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Session, error)
	Refresh(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Session, error)
	ValidateJWT(ctx context.Context, in *JWT, opts ...grpc.CallOption) (*ValidityStatus, error)
	Logout(ctx context.Context, in *Session, opts ...grpc.CallOption) (*LogoutStatus, error)
}

type authenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationClient(cc grpc.ClientConnInterface) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/proto.auth.Authentication/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Refresh(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/proto.auth.Authentication/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) ValidateJWT(ctx context.Context, in *JWT, opts ...grpc.CallOption) (*ValidityStatus, error) {
	out := new(ValidityStatus)
	err := c.cc.Invoke(ctx, "/proto.auth.Authentication/ValidateJWT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Logout(ctx context.Context, in *Session, opts ...grpc.CallOption) (*LogoutStatus, error) {
	out := new(LogoutStatus)
	err := c.cc.Invoke(ctx, "/proto.auth.Authentication/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
type AuthenticationServer interface {
	Login(context.Context, *Credentials) (*Session, error)
	Refresh(context.Context, *Session) (*Session, error)
	ValidateJWT(context.Context, *JWT) (*ValidityStatus, error)
	Logout(context.Context, *Session) (*LogoutStatus, error)
}

// UnimplementedAuthenticationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (*UnimplementedAuthenticationServer) Login(ctx context.Context, req *Credentials) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthenticationServer) Refresh(ctx context.Context, req *Session) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (*UnimplementedAuthenticationServer) ValidateJWT(ctx context.Context, req *JWT) (*ValidityStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateJWT not implemented")
}
func (*UnimplementedAuthenticationServer) Logout(ctx context.Context, req *Session) (*LogoutStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.auth.Authentication/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Login(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.auth.Authentication/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Refresh(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_ValidateJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JWT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).ValidateJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.auth.Authentication/ValidateJWT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).ValidateJWT(ctx, req.(*JWT))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.auth.Authentication/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Logout(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.auth.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Authentication_Login_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Authentication_Refresh_Handler,
		},
		{
			MethodName: "ValidateJWT",
			Handler:    _Authentication_ValidateJWT_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Authentication_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

// AuthorisationClient is the client API for Authorisation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorisationClient interface {
	Check(ctx context.Context, in *Permission, opts ...grpc.CallOption) (*CheckStatus, error)
}

type authorisationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorisationClient(cc grpc.ClientConnInterface) AuthorisationClient {
	return &authorisationClient{cc}
}

func (c *authorisationClient) Check(ctx context.Context, in *Permission, opts ...grpc.CallOption) (*CheckStatus, error) {
	out := new(CheckStatus)
	err := c.cc.Invoke(ctx, "/proto.auth.Authorisation/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorisationServer is the server API for Authorisation service.
type AuthorisationServer interface {
	Check(context.Context, *Permission) (*CheckStatus, error)
}

// UnimplementedAuthorisationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorisationServer struct {
}

func (*UnimplementedAuthorisationServer) Check(ctx context.Context, req *Permission) (*CheckStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterAuthorisationServer(s *grpc.Server, srv AuthorisationServer) {
	s.RegisterService(&_Authorisation_serviceDesc, srv)
}

func _Authorisation_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Permission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorisationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.auth.Authorisation/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorisationServer).Check(ctx, req.(*Permission))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorisation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.auth.Authorisation",
	HandlerType: (*AuthorisationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Authorisation_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
