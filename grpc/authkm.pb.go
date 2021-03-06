// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authkm.proto

package grpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The authentication request message containing the user's credentials
type UserCredentialsReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCredentialsReq) Reset()         { *m = UserCredentialsReq{} }
func (m *UserCredentialsReq) String() string { return proto.CompactTextString(m) }
func (*UserCredentialsReq) ProtoMessage()    {}
func (*UserCredentialsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9ccb2c8c67153f7, []int{0}
}

func (m *UserCredentialsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCredentialsReq.Unmarshal(m, b)
}
func (m *UserCredentialsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCredentialsReq.Marshal(b, m, deterministic)
}
func (m *UserCredentialsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCredentialsReq.Merge(m, src)
}
func (m *UserCredentialsReq) XXX_Size() int {
	return xxx_messageInfo_UserCredentialsReq.Size(m)
}
func (m *UserCredentialsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCredentialsReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserCredentialsReq proto.InternalMessageInfo

func (m *UserCredentialsReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserCredentialsReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// The response message containing the user's authentication claims
type AuthClaimsResp struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Role                 int32    `protobuf:"varint,4,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthClaimsResp) Reset()         { *m = AuthClaimsResp{} }
func (m *AuthClaimsResp) String() string { return proto.CompactTextString(m) }
func (*AuthClaimsResp) ProtoMessage()    {}
func (*AuthClaimsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9ccb2c8c67153f7, []int{1}
}

func (m *AuthClaimsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthClaimsResp.Unmarshal(m, b)
}
func (m *AuthClaimsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthClaimsResp.Marshal(b, m, deterministic)
}
func (m *AuthClaimsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthClaimsResp.Merge(m, src)
}
func (m *AuthClaimsResp) XXX_Size() int {
	return xxx_messageInfo_AuthClaimsResp.Size(m)
}
func (m *AuthClaimsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthClaimsResp.DiscardUnknown(m)
}

var xxx_messageInfo_AuthClaimsResp proto.InternalMessageInfo

func (m *AuthClaimsResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthClaimsResp) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AuthClaimsResp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthClaimsResp) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

// The enrollment request message containing the user's details
type UserEnrollReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserEnrollReq) Reset()         { *m = UserEnrollReq{} }
func (m *UserEnrollReq) String() string { return proto.CompactTextString(m) }
func (*UserEnrollReq) ProtoMessage()    {}
func (*UserEnrollReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9ccb2c8c67153f7, []int{2}
}

func (m *UserEnrollReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEnrollReq.Unmarshal(m, b)
}
func (m *UserEnrollReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEnrollReq.Marshal(b, m, deterministic)
}
func (m *UserEnrollReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEnrollReq.Merge(m, src)
}
func (m *UserEnrollReq) XXX_Size() int {
	return xxx_messageInfo_UserEnrollReq.Size(m)
}
func (m *UserEnrollReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEnrollReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserEnrollReq proto.InternalMessageInfo

func (m *UserEnrollReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserEnrollReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserEnrollReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// The response containing the user's enrollment status
type EnrollStatusResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnrollStatusResp) Reset()         { *m = EnrollStatusResp{} }
func (m *EnrollStatusResp) String() string { return proto.CompactTextString(m) }
func (*EnrollStatusResp) ProtoMessage()    {}
func (*EnrollStatusResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9ccb2c8c67153f7, []int{3}
}

func (m *EnrollStatusResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnrollStatusResp.Unmarshal(m, b)
}
func (m *EnrollStatusResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnrollStatusResp.Marshal(b, m, deterministic)
}
func (m *EnrollStatusResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnrollStatusResp.Merge(m, src)
}
func (m *EnrollStatusResp) XXX_Size() int {
	return xxx_messageInfo_EnrollStatusResp.Size(m)
}
func (m *EnrollStatusResp) XXX_DiscardUnknown() {
	xxx_messageInfo_EnrollStatusResp.DiscardUnknown(m)
}

var xxx_messageInfo_EnrollStatusResp proto.InternalMessageInfo

func (m *EnrollStatusResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *EnrollStatusResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*UserCredentialsReq)(nil), "grpc.UserCredentialsReq")
	proto.RegisterType((*AuthClaimsResp)(nil), "grpc.AuthClaimsResp")
	proto.RegisterType((*UserEnrollReq)(nil), "grpc.UserEnrollReq")
	proto.RegisterType((*EnrollStatusResp)(nil), "grpc.EnrollStatusResp")
}

func init() { proto.RegisterFile("authkm.proto", fileDescriptor_d9ccb2c8c67153f7) }

var fileDescriptor_d9ccb2c8c67153f7 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0x4b, 0xff, 0x00, 0x4e, 0xaa, 0x31, 0x6b, 0xa3, 0x84, 0x53, 0xc3, 0xa9, 0x27, 0x0e,
	0x7a, 0xf0, 0xaa, 0x69, 0x6c, 0x62, 0xbc, 0x61, 0x3c, 0x1a, 0xb3, 0xc2, 0xa4, 0x25, 0xdd, 0x65,
	0x71, 0x67, 0x89, 0x5f, 0xc0, 0x0f, 0x6e, 0x96, 0x2d, 0xa0, 0x35, 0xde, 0xe6, 0xc7, 0x84, 0xf7,
	0xde, 0xbc, 0x85, 0x39, 0x6f, 0xcc, 0x6e, 0x2f, 0xd3, 0x5a, 0x2b, 0xa3, 0xd8, 0x74, 0xab, 0xeb,
	0x3c, 0xd9, 0x00, 0x7b, 0x21, 0xd4, 0x6b, 0x8d, 0x05, 0x56, 0xa6, 0xe4, 0x82, 0x32, 0xfc, 0x60,
	0x0b, 0x98, 0xa1, 0xe4, 0xa5, 0x88, 0xbc, 0xa5, 0xb7, 0x3a, 0xc9, 0x1c, 0xb0, 0x18, 0xc2, 0x9a,
	0x13, 0x7d, 0x2a, 0x5d, 0x44, 0xe3, 0x76, 0xd1, 0x73, 0x52, 0xc2, 0xd9, 0x7d, 0x63, 0x76, 0x6b,
	0xc1, 0x4b, 0x49, 0x19, 0x52, 0x6d, 0x35, 0x8c, 0xda, 0x63, 0xd5, 0x69, 0xb4, 0xc0, 0xae, 0x20,
	0x68, 0x08, 0xf5, 0x5b, 0xd9, 0x49, 0xf8, 0x16, 0x1f, 0x8b, 0xc1, 0x72, 0xf2, 0xd3, 0x92, 0xc1,
	0x54, 0x2b, 0x81, 0xd1, 0x74, 0xe9, 0xad, 0x66, 0x59, 0x3b, 0x27, 0xaf, 0x70, 0x6a, 0x23, 0x3f,
	0x54, 0x5a, 0x09, 0x61, 0xd3, 0xc6, 0x10, 0x5a, 0x91, 0x8a, 0x4b, 0x3c, 0x98, 0xf5, 0x3c, 0xc8,
	0x8e, 0xff, 0xbb, 0x64, 0x72, 0x74, 0xc9, 0x06, 0xce, 0x9d, 0xf4, 0xb3, 0xe1, 0xa6, 0x71, 0xb7,
	0x44, 0x10, 0x50, 0x93, 0xe7, 0x48, 0xd4, 0x1a, 0x84, 0x59, 0x87, 0x76, 0x23, 0x91, 0x88, 0x6f,
	0xf1, 0xe0, 0xd0, 0xe1, 0xf5, 0x97, 0x07, 0x81, 0xad, 0xe4, 0x89, 0x4b, 0x76, 0x07, 0x73, 0x3b,
	0xda, 0x86, 0x73, 0x6e, 0x90, 0x45, 0xa9, 0x2d, 0x3f, 0xfd, 0xdb, 0x7c, 0xbc, 0x70, 0x9b, 0xdf,
	0x5d, 0x26, 0x23, 0x76, 0x0b, 0xbe, 0x4b, 0xc5, 0x2e, 0x86, 0x7f, 0xfb, 0x0a, 0xe2, 0x4b, 0xf7,
	0xf1, 0x38, 0x78, 0x32, 0x7a, 0xf7, 0xdb, 0xd7, 0xbe, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xa7,
	0x4e, 0x0c, 0xa3, 0xfd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthKamClient is the client API for AuthKam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthKamClient interface {
	// Authenticate
	Authenticate(ctx context.Context, in *UserCredentialsReq, opts ...grpc.CallOption) (*AuthClaimsResp, error)
	// Enroll
	Enroll(ctx context.Context, in *UserEnrollReq, opts ...grpc.CallOption) (*EnrollStatusResp, error)
}

type authKamClient struct {
	cc *grpc.ClientConn
}

func NewAuthKamClient(cc *grpc.ClientConn) AuthKamClient {
	return &authKamClient{cc}
}

func (c *authKamClient) Authenticate(ctx context.Context, in *UserCredentialsReq, opts ...grpc.CallOption) (*AuthClaimsResp, error) {
	out := new(AuthClaimsResp)
	err := c.cc.Invoke(ctx, "/grpc.AuthKam/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authKamClient) Enroll(ctx context.Context, in *UserEnrollReq, opts ...grpc.CallOption) (*EnrollStatusResp, error) {
	out := new(EnrollStatusResp)
	err := c.cc.Invoke(ctx, "/grpc.AuthKam/Enroll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthKamServer is the server API for AuthKam service.
type AuthKamServer interface {
	// Authenticate
	Authenticate(context.Context, *UserCredentialsReq) (*AuthClaimsResp, error)
	// Enroll
	Enroll(context.Context, *UserEnrollReq) (*EnrollStatusResp, error)
}

func RegisterAuthKamServer(s *grpc.Server, srv AuthKamServer) {
	s.RegisterService(&_AuthKam_serviceDesc, srv)
}

func _AuthKam_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCredentialsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthKamServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthKam/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthKamServer).Authenticate(ctx, req.(*UserCredentialsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthKam_Enroll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEnrollReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthKamServer).Enroll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.AuthKam/Enroll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthKamServer).Enroll(ctx, req.(*UserEnrollReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthKam_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.AuthKam",
	HandlerType: (*AuthKamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthKam_Authenticate_Handler,
		},
		{
			MethodName: "Enroll",
			Handler:    _AuthKam_Enroll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authkm.proto",
}
