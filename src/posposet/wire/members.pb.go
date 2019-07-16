// Code generated by protoc-gen-go. DO NOT EDIT.
// source: members.proto

package wire

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

type Members struct {
	All                  map[string]uint64 `protobuf:"bytes,1,rep,name=All,json=all,proto3" json:"All,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Members) Reset()         { *m = Members{} }
func (m *Members) String() string { return proto.CompactTextString(m) }
func (*Members) ProtoMessage()    {}
func (*Members) Descriptor() ([]byte, []int) {
	return fileDescriptor_a130ba3d4c49b931, []int{0}
}

func (m *Members) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Members.Unmarshal(m, b)
}
func (m *Members) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Members.Marshal(b, m, deterministic)
}
func (m *Members) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Members.Merge(m, src)
}
func (m *Members) XXX_Size() int {
	return xxx_messageInfo_Members.Size(m)
}
func (m *Members) XXX_DiscardUnknown() {
	xxx_messageInfo_Members.DiscardUnknown(m)
}

var xxx_messageInfo_Members proto.InternalMessageInfo

func (m *Members) GetAll() map[string]uint64 {
	if m != nil {
		return m.All
	}
	return nil
}

func init() {
	proto.RegisterType((*Members)(nil), "wire.Members")
	proto.RegisterMapType((map[string]uint64)(nil), "wire.Members.AllEntry")
}

func init() { proto.RegisterFile("members.proto", fileDescriptor_a130ba3d4c49b931) }

var fileDescriptor_a130ba3d4c49b931 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0xcd, 0x4d,
	0x4a, 0x2d, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0xcf, 0x2c, 0x4a, 0x55,
	0xca, 0xe6, 0x62, 0xf7, 0x85, 0x08, 0x0b, 0x69, 0x70, 0x31, 0x3b, 0xe6, 0xe4, 0x48, 0x30, 0x2a,
	0x30, 0x6b, 0x70, 0x1b, 0x89, 0xe9, 0x81, 0xa4, 0xf5, 0xa0, 0x72, 0x7a, 0x8e, 0x39, 0x39, 0xae,
	0x79, 0x25, 0x45, 0x95, 0x41, 0xcc, 0x89, 0x39, 0x39, 0x52, 0x66, 0x5c, 0x1c, 0x30, 0x01, 0x21,
	0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x48,
	0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x25, 0x08, 0xc2,
	0xb1, 0x62, 0xb2, 0x60, 0x4c, 0x62, 0x03, 0xdb, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd5,
	0xd8, 0xe4, 0x59, 0x8a, 0x00, 0x00, 0x00,
}
