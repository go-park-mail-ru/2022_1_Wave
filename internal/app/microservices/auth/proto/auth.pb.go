// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package proto

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

type Session struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
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

func (m *Session) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type LoginData struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Session              *Session `protobuf:"bytes,3,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginData) Reset()         { *m = LoginData{} }
func (m *LoginData) String() string { return proto.CompactTextString(m) }
func (*LoginData) ProtoMessage()    {}
func (*LoginData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *LoginData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginData.Unmarshal(m, b)
}
func (m *LoginData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginData.Marshal(b, m, deterministic)
}
func (m *LoginData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginData.Merge(m, src)
}
func (m *LoginData) XXX_Size() int {
	return xxx_messageInfo_LoginData.Size(m)
}
func (m *LoginData) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginData.DiscardUnknown(m)
}

var xxx_messageInfo_LoginData proto.InternalMessageInfo

func (m *LoginData) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *LoginData) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginData) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type LoginResult struct {
	NewSession           *Session `protobuf:"bytes,2,opt,name=new_session,json=newSession,proto3" json:"new_session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResult) Reset()         { *m = LoginResult{} }
func (m *LoginResult) String() string { return proto.CompactTextString(m) }
func (*LoginResult) ProtoMessage()    {}
func (*LoginResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *LoginResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResult.Unmarshal(m, b)
}
func (m *LoginResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResult.Marshal(b, m, deterministic)
}
func (m *LoginResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResult.Merge(m, src)
}
func (m *LoginResult) XXX_Size() int {
	return xxx_messageInfo_LoginResult.Size(m)
}
func (m *LoginResult) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResult.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResult proto.InternalMessageInfo

func (m *LoginResult) GetNewSession() *Session {
	if m != nil {
		return m.NewSession
	}
	return nil
}

type LogoutResult struct {
	NewSession           *Session `protobuf:"bytes,2,opt,name=new_session,json=newSession,proto3" json:"new_session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResult) Reset()         { *m = LogoutResult{} }
func (m *LogoutResult) String() string { return proto.CompactTextString(m) }
func (*LogoutResult) ProtoMessage()    {}
func (*LogoutResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *LogoutResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResult.Unmarshal(m, b)
}
func (m *LogoutResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResult.Marshal(b, m, deterministic)
}
func (m *LogoutResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResult.Merge(m, src)
}
func (m *LogoutResult) XXX_Size() int {
	return xxx_messageInfo_LogoutResult.Size(m)
}
func (m *LogoutResult) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResult.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResult proto.InternalMessageInfo

func (m *LogoutResult) GetNewSession() *Session {
	if m != nil {
		return m.NewSession
	}
	return nil
}

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignUpData struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Session              *Session `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpData) Reset()         { *m = SignUpData{} }
func (m *SignUpData) String() string { return proto.CompactTextString(m) }
func (*SignUpData) ProtoMessage()    {}
func (*SignUpData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6}
}

func (m *SignUpData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpData.Unmarshal(m, b)
}
func (m *SignUpData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpData.Marshal(b, m, deterministic)
}
func (m *SignUpData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpData.Merge(m, src)
}
func (m *SignUpData) XXX_Size() int {
	return xxx_messageInfo_SignUpData.Size(m)
}
func (m *SignUpData) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpData.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpData proto.InternalMessageInfo

func (m *SignUpData) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *SignUpData) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type SignUpResult struct {
	NewSession           *Session `protobuf:"bytes,2,opt,name=new_session,json=newSession,proto3" json:"new_session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpResult) Reset()         { *m = SignUpResult{} }
func (m *SignUpResult) String() string { return proto.CompactTextString(m) }
func (*SignUpResult) ProtoMessage()    {}
func (*SignUpResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{7}
}

func (m *SignUpResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpResult.Unmarshal(m, b)
}
func (m *SignUpResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpResult.Marshal(b, m, deterministic)
}
func (m *SignUpResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpResult.Merge(m, src)
}
func (m *SignUpResult) XXX_Size() int {
	return xxx_messageInfo_SignUpResult.Size(m)
}
func (m *SignUpResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpResult.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpResult proto.InternalMessageInfo

func (m *SignUpResult) GetNewSession() *Session {
	if m != nil {
		return m.NewSession
	}
	return nil
}

type GetSessionResult struct {
	Session              *Session `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionResult) Reset()         { *m = GetSessionResult{} }
func (m *GetSessionResult) String() string { return proto.CompactTextString(m) }
func (*GetSessionResult) ProtoMessage()    {}
func (*GetSessionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{8}
}

func (m *GetSessionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionResult.Unmarshal(m, b)
}
func (m *GetSessionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionResult.Marshal(b, m, deterministic)
}
func (m *GetSessionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionResult.Merge(m, src)
}
func (m *GetSessionResult) XXX_Size() int {
	return xxx_messageInfo_GetSessionResult.Size(m)
}
func (m *GetSessionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionResult proto.InternalMessageInfo

func (m *GetSessionResult) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func init() {
	proto.RegisterType((*Session)(nil), "auth.Session")
	proto.RegisterType((*Empty)(nil), "auth.Empty")
	proto.RegisterType((*LoginData)(nil), "auth.LoginData")
	proto.RegisterType((*LoginResult)(nil), "auth.LoginResult")
	proto.RegisterType((*LogoutResult)(nil), "auth.LogoutResult")
	proto.RegisterType((*User)(nil), "auth.User")
	proto.RegisterType((*SignUpData)(nil), "auth.SignUpData")
	proto.RegisterType((*SignUpResult)(nil), "auth.SignUpResult")
	proto.RegisterType((*GetSessionResult)(nil), "auth.GetSessionResult")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x0b, 0xd3, 0x40,
	0x10, 0x4d, 0xd3, 0x2f, 0x33, 0x69, 0xb1, 0x2e, 0x52, 0x42, 0x40, 0x29, 0x7b, 0xb1, 0x1e, 0xda,
	0x4a, 0x3d, 0x8a, 0x15, 0x45, 0x29, 0x85, 0x9e, 0x52, 0x73, 0xf1, 0x52, 0xd6, 0x66, 0x6d, 0x17,
	0x9a, 0x6c, 0xc8, 0x6e, 0x2c, 0xfa, 0xd3, 0xfc, 0x75, 0xb2, 0x1f, 0x09, 0x69, 0x11, 0x29, 0x3d,
	0x65, 0x67, 0x67, 0xde, 0x9b, 0x79, 0x6f, 0x27, 0x00, 0xa4, 0x94, 0xa7, 0x79, 0x5e, 0x70, 0xc9,
	0x51, 0x47, 0x9d, 0xf1, 0x14, 0xfa, 0x3b, 0x2a, 0x04, 0xe3, 0x19, 0x7a, 0x01, 0x20, 0xcc, 0x71,
	0xcf, 0x92, 0xa0, 0x35, 0x69, 0x4d, 0xbd, 0xc8, 0xb3, 0x37, 0x9b, 0x04, 0xf7, 0xa1, 0xfb, 0x25,
	0xcd, 0xe5, 0x2f, 0xfc, 0x03, 0xbc, 0x2d, 0x3f, 0xb2, 0xec, 0x33, 0x91, 0x04, 0x3d, 0x87, 0xee,
	0x59, 0x05, 0xb6, 0xde, 0x04, 0x28, 0x84, 0x27, 0x39, 0x11, 0xe2, 0xc2, 0x8b, 0x24, 0x70, 0x75,
	0xa2, 0x8e, 0xd1, 0x2b, 0xe8, 0x5b, 0xd2, 0xa0, 0x3d, 0x69, 0x4d, 0xfd, 0xe5, 0x70, 0xae, 0xa7,
	0xb2, 0x63, 0x44, 0x55, 0x16, 0xbf, 0x07, 0x5f, 0xf7, 0x89, 0xa8, 0x28, 0xcf, 0x12, 0xcd, 0xc1,
	0xcf, 0xe8, 0x65, 0x5f, 0x61, 0xdd, 0x7f, 0x61, 0x21, 0xa3, 0x17, 0x7b, 0xc6, 0x2b, 0x18, 0x6c,
	0xf9, 0x91, 0x97, 0xf2, 0x41, 0xfc, 0x57, 0xe8, 0xc4, 0x82, 0x16, 0x4a, 0x4b, 0x29, 0x68, 0x91,
	0x91, 0x94, 0x5a, 0x91, 0x75, 0xac, 0xd4, 0xd3, 0x94, 0xb0, 0xb3, 0x15, 0x69, 0x82, 0x2b, 0xf5,
	0xed, 0x6b, 0xf5, 0x38, 0x06, 0xd8, 0xb1, 0x63, 0x16, 0xe7, 0xda, 0xbd, 0x97, 0xd0, 0x51, 0x5c,
	0x9a, 0xd7, 0x5f, 0x82, 0x19, 0x46, 0x75, 0x8d, 0xf4, 0x7d, 0xd3, 0x2b, 0xf7, 0xbf, 0x5e, 0xad,
	0x60, 0x60, 0x68, 0x1f, 0x14, 0xfb, 0x0e, 0x46, 0x6b, 0x2a, 0xab, 0x8c, 0xe1, 0xb8, 0xb7, 0xf9,
	0xf2, 0x8f, 0x0b, 0xc3, 0x8f, 0xa5, 0x3c, 0xf1, 0x82, 0xfd, 0x26, 0x52, 0xad, 0xd2, 0x0c, 0xba,
	0xfa, 0xe9, 0xd0, 0x53, 0x03, 0xa9, 0xf7, 0x25, 0x7c, 0xd6, 0xb8, 0x30, 0x7d, 0xb0, 0x83, 0x66,
	0xd0, 0x33, 0x4f, 0x85, 0xae, 0x5b, 0x84, 0xa8, 0xae, 0xae, 0xdf, 0x11, 0x3b, 0xe8, 0x0d, 0xf4,
	0x8c, 0x58, 0x34, 0xb2, 0xe5, 0xb5, 0xa3, 0x15, 0xa2, 0x69, 0x06, 0x76, 0xd0, 0x07, 0x18, 0xaf,
	0xa9, 0x8c, 0x33, 0x62, 0xa7, 0xa4, 0x49, 0xb5, 0xf4, 0xbe, 0xa9, 0xd7, 0x9b, 0x1d, 0x8e, 0x4d,
	0x70, 0xeb, 0x04, 0x76, 0xd0, 0x6b, 0xf0, 0x36, 0xa2, 0xc2, 0xdc, 0x0c, 0xd9, 0xa4, 0xd0, 0x62,
	0x86, 0x1b, 0xa1, 0xec, 0xb8, 0xab, 0xfc, 0x53, 0xf8, 0x2d, 0x48, 0xd9, 0xa1, 0xe0, 0x82, 0x16,
	0x3f, 0xd9, 0x81, 0x8a, 0x85, 0xca, 0x2e, 0xf4, 0x2f, 0xfa, 0xbd, 0xa7, 0x3f, 0x6f, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x12, 0x61, 0x31, 0x48, 0xb7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	Login(ctx context.Context, in *LoginData, opts ...grpc.CallOption) (*LoginResult, error)
	Logout(ctx context.Context, in *Session, opts ...grpc.CallOption) (*LogoutResult, error)
	SignUp(ctx context.Context, in *SignUpData, opts ...grpc.CallOption) (*SignUpResult, error)
	GetUnauthorizedSession(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetSessionResult, error)
	IsSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Empty, error)
	IsAuthSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Empty, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Login(ctx context.Context, in *LoginData, opts ...grpc.CallOption) (*LoginResult, error) {
	out := new(LoginResult)
	err := c.cc.Invoke(ctx, "/auth.Authorization/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) Logout(ctx context.Context, in *Session, opts ...grpc.CallOption) (*LogoutResult, error) {
	out := new(LogoutResult)
	err := c.cc.Invoke(ctx, "/auth.Authorization/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) SignUp(ctx context.Context, in *SignUpData, opts ...grpc.CallOption) (*SignUpResult, error) {
	out := new(SignUpResult)
	err := c.cc.Invoke(ctx, "/auth.Authorization/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) GetUnauthorizedSession(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetSessionResult, error) {
	out := new(GetSessionResult)
	err := c.cc.Invoke(ctx, "/auth.Authorization/GetUnauthorizedSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) IsSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/auth.Authorization/IsSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) IsAuthSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/auth.Authorization/IsAuthSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	Login(context.Context, *LoginData) (*LoginResult, error)
	Logout(context.Context, *Session) (*LogoutResult, error)
	SignUp(context.Context, *SignUpData) (*SignUpResult, error)
	GetUnauthorizedSession(context.Context, *Empty) (*GetSessionResult, error)
	IsSession(context.Context, *Session) (*Empty, error)
	IsAuthSession(context.Context, *Session) (*Empty, error)
}

// UnimplementedAuthorizationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServer struct {
}

func (*UnimplementedAuthorizationServer) Login(ctx context.Context, req *LoginData) (*LoginResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthorizationServer) Logout(ctx context.Context, req *Session) (*LogoutResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (*UnimplementedAuthorizationServer) SignUp(ctx context.Context, req *SignUpData) (*SignUpResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedAuthorizationServer) GetUnauthorizedSession(ctx context.Context, req *Empty) (*GetSessionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnauthorizedSession not implemented")
}
func (*UnimplementedAuthorizationServer) IsSession(ctx context.Context, req *Session) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsSession not implemented")
}
func (*UnimplementedAuthorizationServer) IsAuthSession(ctx context.Context, req *Session) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuthSession not implemented")
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Authorization/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Login(ctx, req.(*LoginData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Authorization/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Logout(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Authorization/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).SignUp(ctx, req.(*SignUpData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_GetUnauthorizedSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).GetUnauthorizedSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Authorization/GetUnauthorizedSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).GetUnauthorizedSession(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_IsSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).IsSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Authorization/IsSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).IsSession(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_IsAuthSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).IsAuthSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Authorization/IsAuthSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).IsAuthSession(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Authorization_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Authorization_Logout_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _Authorization_SignUp_Handler,
		},
		{
			MethodName: "GetUnauthorizedSession",
			Handler:    _Authorization_GetUnauthorizedSession_Handler,
		},
		{
			MethodName: "IsSession",
			Handler:    _Authorization_IsSession_Handler,
		},
		{
			MethodName: "IsAuthSession",
			Handler:    _Authorization_IsAuthSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
