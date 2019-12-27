// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/listener/quic_config.proto

package envoy_api_v3alpha_listener

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type QuicProtocolOptions struct {
	MaxConcurrentStreams   *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_concurrent_streams,json=maxConcurrentStreams,proto3" json:"max_concurrent_streams,omitempty"`
	IdleTimeout            *duration.Duration    `protobuf:"bytes,2,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	CryptoHandshakeTimeout *duration.Duration    `protobuf:"bytes,3,opt,name=crypto_handshake_timeout,json=cryptoHandshakeTimeout,proto3" json:"crypto_handshake_timeout,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *QuicProtocolOptions) Reset()         { *m = QuicProtocolOptions{} }
func (m *QuicProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*QuicProtocolOptions) ProtoMessage()    {}
func (*QuicProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_30110b89b5635b16, []int{0}
}

func (m *QuicProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuicProtocolOptions.Unmarshal(m, b)
}
func (m *QuicProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuicProtocolOptions.Marshal(b, m, deterministic)
}
func (m *QuicProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuicProtocolOptions.Merge(m, src)
}
func (m *QuicProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_QuicProtocolOptions.Size(m)
}
func (m *QuicProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_QuicProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_QuicProtocolOptions proto.InternalMessageInfo

func (m *QuicProtocolOptions) GetMaxConcurrentStreams() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConcurrentStreams
	}
	return nil
}

func (m *QuicProtocolOptions) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *QuicProtocolOptions) GetCryptoHandshakeTimeout() *duration.Duration {
	if m != nil {
		return m.CryptoHandshakeTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*QuicProtocolOptions)(nil), "envoy.api.v3alpha.listener.QuicProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/listener/quic_config.proto", fileDescriptor_30110b89b5635b16)
}

var fileDescriptor_30110b89b5635b16 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4e, 0x3a, 0x31,
	0x10, 0xc7, 0x03, 0xbf, 0xe4, 0x77, 0x28, 0x26, 0x26, 0xab, 0x21, 0x48, 0x0c, 0x51, 0x4f, 0x18,
	0x4d, 0x6b, 0xe0, 0x66, 0xf4, 0x02, 0x1e, 0xf4, 0x24, 0x82, 0x7a, 0xdd, 0x0c, 0xbb, 0x05, 0x1a,
	0xbb, 0x9d, 0xda, 0x3f, 0x08, 0x6f, 0xe0, 0x33, 0xf8, 0x10, 0xbe, 0x85, 0xef, 0x65, 0xb6, 0xbb,
	0xa0, 0x09, 0x1a, 0x8f, 0xcd, 0x7c, 0xbe, 0x9f, 0x99, 0xce, 0x90, 0x53, 0xae, 0xe6, 0xb8, 0x64,
	0xa0, 0x05, 0x9b, 0x77, 0x41, 0xea, 0x19, 0x30, 0x29, 0xac, 0xe3, 0x8a, 0x1b, 0xf6, 0xec, 0x45,
	0x12, 0x27, 0xa8, 0x26, 0x62, 0x4a, 0xb5, 0x41, 0x87, 0x51, 0x33, 0xd0, 0x14, 0xb4, 0xa0, 0x25,
	0x4d, 0x57, 0x74, 0xb3, 0x35, 0x45, 0x9c, 0x4a, 0xce, 0x02, 0x39, 0xf6, 0x13, 0x96, 0x7a, 0x03,
	0x4e, 0xa0, 0x2a, 0xb2, 0x9b, 0xf5, 0x17, 0x03, 0x5a, 0x73, 0x63, 0xcb, 0xfa, 0xa1, 0x4f, 0x35,
	0x30, 0x50, 0x0a, 0x5d, 0x88, 0x59, 0x36, 0xe7, 0xc6, 0x0a, 0x54, 0x42, 0x95, 0xed, 0x8f, 0xde,
	0xab, 0x64, 0xe7, 0xce, 0x8b, 0x64, 0x90, 0xbf, 0x12, 0x94, 0xb7, 0x3a, 0x80, 0xd1, 0x90, 0xd4,
	0x33, 0x58, 0xe4, 0xa3, 0x26, 0xde, 0x18, 0xae, 0x5c, 0x6c, 0x9d, 0xe1, 0x90, 0xd9, 0x46, 0xe5,
	0xa0, 0xd2, 0xae, 0x75, 0xf6, 0x69, 0xd1, 0x9b, 0xae, 0x7a, 0xd3, 0x87, 0x1b, 0xe5, 0xba, 0x9d,
	0x47, 0x90, 0x9e, 0x0f, 0x77, 0x33, 0x58, 0xf4, 0xd7, 0xd1, 0x51, 0x91, 0x8c, 0x2e, 0xc8, 0x96,
	0x48, 0x25, 0x8f, 0x9d, 0xc8, 0x38, 0x7a, 0xd7, 0xa8, 0x06, 0xd3, 0xde, 0x86, 0xe9, 0xaa, 0xfc,
	0xe5, 0xb0, 0x96, 0xe3, 0xf7, 0x05, 0x1d, 0x8d, 0x48, 0x23, 0x31, 0x4b, 0xed, 0x30, 0x9e, 0x81,
	0x4a, 0xed, 0x0c, 0x9e, 0xbe, 0x4c, 0xff, 0xfe, 0x32, 0xd5, 0x8b, 0xe8, 0xf5, 0x2a, 0x59, 0x4a,
	0xcf, 0xcf, 0xde, 0x3e, 0x5e, 0x5b, 0x27, 0xe4, 0xf8, 0xdb, 0x11, 0x3a, 0xeb, 0xfd, 0xd3, 0x1f,
	0x16, 0xd3, 0xbb, 0x24, 0x6d, 0x81, 0x34, 0xf0, 0xda, 0xe0, 0x62, 0x49, 0x7f, 0xbf, 0x5f, 0x6f,
	0x3b, 0x17, 0xf4, 0xc3, 0xb5, 0x83, 0x66, 0x50, 0x19, 0xff, 0x0f, 0xb3, 0x75, 0x3f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xf1, 0x53, 0xfb, 0xa2, 0x25, 0x02, 0x00, 0x00,
}