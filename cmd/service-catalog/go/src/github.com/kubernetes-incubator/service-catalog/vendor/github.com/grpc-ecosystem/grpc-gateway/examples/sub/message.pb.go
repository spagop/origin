// Code generated by protoc-gen-go.
// source: examples/sub/message.proto
// DO NOT EDIT!

/*
Package sub is a generated protocol buffer package.

It is generated from these files:
	examples/sub/message.proto

It has these top-level messages:
	StringMessage
*/
package sub

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StringMessage struct {
	Value            *string `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StringMessage) Reset()                    { *m = StringMessage{} }
func (m *StringMessage) String() string            { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()               {}
func (*StringMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StringMessage) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StringMessage)(nil), "grpc.gateway.examples.sub.StringMessage")
}

func init() { proto.RegisterFile("examples/sub/message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x2e, 0x4d, 0xd2, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4c, 0x2f, 0x2a, 0x48, 0xd6, 0x4b, 0x4f, 0x2c,
	0x49, 0x2d, 0x4f, 0xac, 0xd4, 0x83, 0x29, 0xd4, 0x2b, 0x2e, 0x4d, 0x52, 0x52, 0xe5, 0xe2, 0x0d,
	0x2e, 0x29, 0xca, 0xcc, 0x4b, 0xf7, 0x85, 0xe8, 0x10, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29,
	0x4d, 0x95, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x0c, 0x82, 0x70, 0x9c, 0x58, 0xa3, 0x98, 0x8b, 0x4b,
	0x93, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x10, 0x60, 0xa9, 0x65, 0x00, 0x00, 0x00,
}
