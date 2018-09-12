// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue444.proto

package issue444

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SizeMe struct {
	Foo                  string   `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SizeMe) Reset()         { *m = SizeMe{} }
func (m *SizeMe) String() string { return proto.CompactTextString(m) }
func (*SizeMe) ProtoMessage()    {}
func (*SizeMe) Descriptor() ([]byte, []int) {
	return fileDescriptor_issue444_1285670608cee74f, []int{0}
}
func (m *SizeMe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SizeMe.Unmarshal(m, b)
}
func (m *SizeMe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SizeMe.Marshal(b, m, deterministic)
}
func (dst *SizeMe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SizeMe.Merge(dst, src)
}
func (m *SizeMe) XXX_Size() int {
	return xxx_messageInfo_SizeMe.Size(m)
}
func (m *SizeMe) XXX_DiscardUnknown() {
	xxx_messageInfo_SizeMe.DiscardUnknown(m)
}

var xxx_messageInfo_SizeMe proto.InternalMessageInfo

func (m *SizeMe) GetFoo() string {
	if m != nil {
		return m.Foo
	}
	return ""
}

func init() {
	proto.RegisterType((*SizeMe)(nil), "issue444.SizeMe")
}
func (m *SizeMe) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Foo)
	if l > 0 {
		n += 1 + l + sovIssue444(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIssue444(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIssue444(x uint64) (n int) {
	return sovIssue444(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func init() { proto.RegisterFile("issue444.proto", fileDescriptor_issue444_1285670608cee74f) }

var fileDescriptor_issue444_1285670608cee74f = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x31, 0x31, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x74,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5,
	0xc1, 0x0a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x54, 0x92, 0xe2, 0x62,
	0x0b, 0xce, 0xac, 0x4a, 0xf5, 0x4d, 0x15, 0x12, 0xe0, 0x62, 0x4e, 0xcb, 0xcf, 0x97, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x9d, 0x58, 0x1e, 0x3c, 0x92, 0x63, 0x4c, 0x62, 0x03, 0x2b,
	0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xe0, 0xfa, 0x18, 0x73, 0x00, 0x00, 0x00,
}
