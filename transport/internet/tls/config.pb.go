// Code generated by protoc-gen-go.
// source: v2ray.com/core/transport/internet/tls/config.proto
// DO NOT EDIT!

/*
Package tls is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/transport/internet/tls/config.proto

It has these top-level messages:
	Certificate
	Config
*/
package tls

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

type Certificate struct {
	Certificate []byte `protobuf:"bytes,1,opt,name=Certificate,json=certificate,proto3" json:"Certificate,omitempty"`
	Key         []byte `protobuf:"bytes,2,opt,name=Key,json=key,proto3" json:"Key,omitempty"`
}

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Config struct {
	AllowInsecure bool           `protobuf:"varint,1,opt,name=allow_insecure,json=allowInsecure" json:"allow_insecure,omitempty"`
	Certificate   []*Certificate `protobuf:"bytes,2,rep,name=certificate" json:"certificate,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetCertificate() []*Certificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func init() {
	proto.RegisterType((*Certificate)(nil), "v2ray.core.transport.internet.tls.Certificate")
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.tls.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/transport/internet/tls/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4b, 0x43, 0x31,
	0x10, 0xc6, 0x79, 0x0d, 0x14, 0x49, 0x54, 0x24, 0xd3, 0x1b, 0x9f, 0x85, 0x42, 0xa7, 0x0b, 0x3c,
	0x27, 0x47, 0xdb, 0x49, 0x5c, 0x4a, 0x47, 0x17, 0x89, 0xe1, 0x2a, 0xc1, 0x34, 0x57, 0x2e, 0xa7,
	0xf2, 0x46, 0xff, 0x73, 0x31, 0xb5, 0xf2, 0x3a, 0x75, 0xbc, 0xef, 0xbe, 0xfb, 0xee, 0xc7, 0xa7,
	0xfb, 0xcf, 0x9e, 0xfd, 0x00, 0x81, 0x76, 0x2e, 0x10, 0xa3, 0x13, 0xf6, 0xb9, 0xec, 0x89, 0xc5,
	0xc5, 0x2c, 0xc8, 0x19, 0xc5, 0x49, 0x2a, 0x2e, 0x50, 0xde, 0xc6, 0x37, 0xd8, 0x33, 0x09, 0xd9,
	0xdb, 0xe3, 0x0d, 0x23, 0xfc, 0xfb, 0xe1, 0xe8, 0x07, 0x49, 0x65, 0xf6, 0xa0, 0xcd, 0x0a, 0x59,
	0xe2, 0x36, 0x06, 0x2f, 0x68, 0xbb, 0x93, 0xb1, 0x6d, 0xba, 0x66, 0x71, 0xb9, 0x31, 0x61, 0xe4,
	0xb8, 0xd1, 0xea, 0x09, 0x87, 0x76, 0x52, 0x37, 0xea, 0x1d, 0x87, 0xd9, 0x77, 0xa3, 0xa7, 0xab,
	0xfa, 0xd6, 0xce, 0xf5, 0xb5, 0x4f, 0x89, 0xbe, 0x5e, 0x62, 0x2e, 0x18, 0x3e, 0xf8, 0x90, 0x70,
	0xb1, 0xb9, 0xaa, 0xea, 0xe3, 0x9f, 0x68, 0xd7, 0x7a, 0x1c, 0xd9, 0x4e, 0x3a, 0xb5, 0x30, 0x3d,
	0xc0, 0x59, 0x5a, 0x18, 0xb1, 0x9d, 0x50, 0x2d, 0xef, 0xf5, 0x3c, 0xd0, 0xee, 0x7c, 0xc2, 0xd2,
	0x1c, 0x48, 0xd7, 0xbf, 0xfd, 0x3c, 0x2b, 0x49, 0xe5, 0x75, 0x5a, 0xbb, 0xba, 0xfb, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xbb, 0x82, 0x9d, 0x49, 0x61, 0x01, 0x00, 0x00,
}
