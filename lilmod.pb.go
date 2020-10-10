// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lilmod.proto

package lilmod

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

type ProxyRequest struct {
	ProxyToken           string   `protobuf:"bytes,1,opt,name=proxyToken,proto3" json:"proxyToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyRequest) Reset()         { *m = ProxyRequest{} }
func (m *ProxyRequest) String() string { return proto.CompactTextString(m) }
func (*ProxyRequest) ProtoMessage()    {}
func (*ProxyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38431b6be50411ac, []int{0}
}

func (m *ProxyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyRequest.Unmarshal(m, b)
}
func (m *ProxyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyRequest.Marshal(b, m, deterministic)
}
func (m *ProxyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyRequest.Merge(m, src)
}
func (m *ProxyRequest) XXX_Size() int {
	return xxx_messageInfo_ProxyRequest.Size(m)
}
func (m *ProxyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyRequest proto.InternalMessageInfo

func (m *ProxyRequest) GetProxyToken() string {
	if m != nil {
		return m.ProxyToken
	}
	return ""
}

type ProxyResponse struct {
	URL                  string   `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
	PrivatePaths         []string `protobuf:"bytes,2,rep,name=privatePaths,proto3" json:"privatePaths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyResponse) Reset()         { *m = ProxyResponse{} }
func (m *ProxyResponse) String() string { return proto.CompactTextString(m) }
func (*ProxyResponse) ProtoMessage()    {}
func (*ProxyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38431b6be50411ac, []int{1}
}

func (m *ProxyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyResponse.Unmarshal(m, b)
}
func (m *ProxyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyResponse.Marshal(b, m, deterministic)
}
func (m *ProxyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyResponse.Merge(m, src)
}
func (m *ProxyResponse) XXX_Size() int {
	return xxx_messageInfo_ProxyResponse.Size(m)
}
func (m *ProxyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyResponse proto.InternalMessageInfo

func (m *ProxyResponse) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *ProxyResponse) GetPrivatePaths() []string {
	if m != nil {
		return m.PrivatePaths
	}
	return nil
}

func init() {
	proto.RegisterType((*ProxyRequest)(nil), "lilmod.ProxyRequest")
	proto.RegisterType((*ProxyResponse)(nil), "lilmod.ProxyResponse")
}

func init() {
	proto.RegisterFile("lilmod.proto", fileDescriptor_38431b6be50411ac)
}

var fileDescriptor_38431b6be50411ac = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xc9, 0xcc, 0xc9,
	0xcd, 0x4f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xf4, 0xb8, 0x78,
	0x02, 0x8a, 0xf2, 0x2b, 0x2a, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xe4, 0xb8, 0xb8,
	0x0a, 0x40, 0xfc, 0x90, 0xfc, 0xec, 0xd4, 0x3c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x24,
	0x11, 0x25, 0x57, 0x2e, 0x5e, 0xa8, 0xfa, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x01, 0x2e,
	0xe6, 0xd0, 0x20, 0x1f, 0xa8, 0x4a, 0x10, 0x53, 0x48, 0x89, 0x8b, 0xa7, 0xa0, 0x28, 0xb3, 0x2c,
	0xb1, 0x24, 0x35, 0x20, 0xb1, 0x24, 0xa3, 0x58, 0x82, 0x49, 0x81, 0x59, 0x83, 0x33, 0x08, 0x45,
	0xcc, 0xc8, 0x9a, 0x8b, 0xd9, 0x31, 0xc0, 0x53, 0xc8, 0x84, 0x8b, 0x15, 0x6c, 0x9a, 0x90, 0x88,
	0x1e, 0xd4, 0x75, 0xc8, 0x8e, 0x91, 0x12, 0x45, 0x13, 0x85, 0x58, 0xe9, 0xc4, 0x11, 0x05, 0x75,
	0x7d, 0x12, 0x1b, 0xd8, 0x33, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0x7a, 0x92, 0x9e,
	0xdc, 0x00, 0x00, 0x00,
}
