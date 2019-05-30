// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gataway.proto

package gataway

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Login struct {
	Login                int32    `protobuf:"varint,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Login) Reset()         { *m = Login{} }
func (m *Login) String() string { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()    {}
func (*Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_a588296e513e1f02, []int{0}
}

func (m *Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login.Unmarshal(m, b)
}
func (m *Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login.Marshal(b, m, deterministic)
}
func (m *Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login.Merge(m, src)
}
func (m *Login) XXX_Size() int {
	return xxx_messageInfo_Login.Size(m)
}
func (m *Login) XXX_DiscardUnknown() {
	xxx_messageInfo_Login.DiscardUnknown(m)
}

var xxx_messageInfo_Login proto.InternalMessageInfo

func (m *Login) GetLogin() int32 {
	if m != nil {
		return m.Login
	}
	return 0
}

func (m *Login) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReply struct {
	// Types that are valid to be assigned to LoginReplyMessage:
	//	*LoginReply_LoginType
	//	*LoginReply_LoginMessage
	LoginReplyMessage    isLoginReply_LoginReplyMessage `protobuf_oneof:"LoginReplyMessage"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_a588296e513e1f02, []int{1}
}

func (m *LoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReply.Unmarshal(m, b)
}
func (m *LoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReply.Marshal(b, m, deterministic)
}
func (m *LoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReply.Merge(m, src)
}
func (m *LoginReply) XXX_Size() int {
	return xxx_messageInfo_LoginReply.Size(m)
}
func (m *LoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReply proto.InternalMessageInfo

type isLoginReply_LoginReplyMessage interface {
	isLoginReply_LoginReplyMessage()
}

type LoginReply_LoginType struct {
	LoginType bool `protobuf:"varint,1,opt,name=LoginType,proto3,oneof"`
}

type LoginReply_LoginMessage struct {
	LoginMessage string `protobuf:"bytes,2,opt,name=LoginMessage,proto3,oneof"`
}

func (*LoginReply_LoginType) isLoginReply_LoginReplyMessage() {}

func (*LoginReply_LoginMessage) isLoginReply_LoginReplyMessage() {}

func (m *LoginReply) GetLoginReplyMessage() isLoginReply_LoginReplyMessage {
	if m != nil {
		return m.LoginReplyMessage
	}
	return nil
}

func (m *LoginReply) GetLoginType() bool {
	if x, ok := m.GetLoginReplyMessage().(*LoginReply_LoginType); ok {
		return x.LoginType
	}
	return false
}

func (m *LoginReply) GetLoginMessage() string {
	if x, ok := m.GetLoginReplyMessage().(*LoginReply_LoginMessage); ok {
		return x.LoginMessage
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LoginReply) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LoginReply_LoginType)(nil),
		(*LoginReply_LoginMessage)(nil),
	}
}

type Move struct {
	Type                 int32    `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Direction            []int32  `protobuf:"varint,2,rep,packed,name=direction,proto3" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Move) Reset()         { *m = Move{} }
func (m *Move) String() string { return proto.CompactTextString(m) }
func (*Move) ProtoMessage()    {}
func (*Move) Descriptor() ([]byte, []int) {
	return fileDescriptor_a588296e513e1f02, []int{2}
}

func (m *Move) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Move.Unmarshal(m, b)
}
func (m *Move) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Move.Marshal(b, m, deterministic)
}
func (m *Move) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Move.Merge(m, src)
}
func (m *Move) XXX_Size() int {
	return xxx_messageInfo_Move.Size(m)
}
func (m *Move) XXX_DiscardUnknown() {
	xxx_messageInfo_Move.DiscardUnknown(m)
}

var xxx_messageInfo_Move proto.InternalMessageInfo

func (m *Move) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Move) GetDirection() []int32 {
	if m != nil {
		return m.Direction
	}
	return nil
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
	return fileDescriptor_a588296e513e1f02, []int{3}
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

func init() {
	proto.RegisterType((*Login)(nil), "Login")
	proto.RegisterType((*LoginReply)(nil), "LoginReply")
	proto.RegisterType((*Move)(nil), "Move")
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() { proto.RegisterFile("gataway.proto", fileDescriptor_a588296e513e1f02) }

var fileDescriptor_a588296e513e1f02 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x9b, 0xda, 0x8d, 0xcd, 0xab, 0x1e, 0x1c, 0x3d, 0x84, 0x28, 0x12, 0x16, 0x0f, 0x39,
	0xed, 0x41, 0x2f, 0x7a, 0x54, 0x10, 0x7b, 0x30, 0x20, 0xab, 0x7f, 0x60, 0xb5, 0x4b, 0x08, 0x68,
	0x37, 0x24, 0x4b, 0xcb, 0xfe, 0x7b, 0xe9, 0x74, 0x69, 0xf4, 0x36, 0xef, 0xe3, 0xf1, 0x0d, 0x33,
	0x38, 0x6d, 0x8c, 0x37, 0x5b, 0x13, 0x54, 0xd7, 0x3b, 0xef, 0xe4, 0x03, 0xc4, 0xab, 0x6b, 0xda,
	0x35, 0x5d, 0xc4, 0x21, 0x4f, 0xca, 0xa4, 0x12, 0x3a, 0xd2, 0x02, 0xf3, 0x37, 0x33, 0x0c, 0x5b,
	0xd7, 0xaf, 0xf2, 0x69, 0x99, 0x54, 0x99, 0x3e, 0x64, 0xd9, 0x00, 0x5c, 0xd2, 0xb6, 0xfb, 0x0e,
	0x74, 0x8d, 0x8c, 0xd3, 0x47, 0xe8, 0x2c, 0x3b, 0xe6, 0xcb, 0x89, 0x1e, 0x11, 0xdd, 0xe0, 0x84,
	0x43, 0x6d, 0x87, 0xc1, 0x34, 0x76, 0x6f, 0x5b, 0x4e, 0xf4, 0x3f, 0xfa, 0x74, 0x8e, 0xb3, 0xd1,
	0x19, 0xa1, 0xbc, 0xc7, 0xac, 0x76, 0x1b, 0x4b, 0x84, 0xd9, 0xc1, 0x2e, 0x34, 0xcf, 0x74, 0x85,
	0x6c, 0xd5, 0xf6, 0xf6, 0xcb, 0xb7, 0x6e, 0x9d, 0x4f, 0xcb, 0xa3, 0x4a, 0xe8, 0x11, 0xc8, 0x63,
	0x88, 0xe7, 0x9f, 0xce, 0x87, 0xdb, 0x1a, 0x78, 0xd9, 0xdf, 0xfd, 0xde, 0x6f, 0x48, 0x62, 0xc1,
	0x5b, 0x1e, 0xb9, 0x45, 0xa9, 0xe2, 0x54, 0x2c, 0xd4, 0x9f, 0x7b, 0x2e, 0x81, 0xdd, 0xd2, 0x58,
	0x11, 0x6a, 0x17, 0x8a, 0x54, 0xb1, 0xee, 0x33, 0xe5, 0xe7, 0xdd, 0xfd, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xe8, 0xf0, 0xb1, 0x2f, 0x4d, 0x01, 0x00, 0x00,
}
