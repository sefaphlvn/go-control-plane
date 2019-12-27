// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/common/dynamic_forward_proxy/v3alpha/dns_cache.proto

package envoy_config_common_dynamic_forward_proxy_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type DnsCacheConfig struct {
	Name                 string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DnsLookupFamily      v3alpha.Cluster_DnsLookupFamily `protobuf:"varint,2,opt,name=dns_lookup_family,json=dnsLookupFamily,proto3,enum=envoy.api.v3alpha.Cluster_DnsLookupFamily" json:"dns_lookup_family,omitempty"`
	DnsRefreshRate       *duration.Duration              `protobuf:"bytes,3,opt,name=dns_refresh_rate,json=dnsRefreshRate,proto3" json:"dns_refresh_rate,omitempty"`
	HostTtl              *duration.Duration              `protobuf:"bytes,4,opt,name=host_ttl,json=hostTtl,proto3" json:"host_ttl,omitempty"`
	MaxHosts             *wrappers.UInt32Value           `protobuf:"bytes,5,opt,name=max_hosts,json=maxHosts,proto3" json:"max_hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *DnsCacheConfig) Reset()         { *m = DnsCacheConfig{} }
func (m *DnsCacheConfig) String() string { return proto.CompactTextString(m) }
func (*DnsCacheConfig) ProtoMessage()    {}
func (*DnsCacheConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f21fdb25ee49be2, []int{0}
}

func (m *DnsCacheConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsCacheConfig.Unmarshal(m, b)
}
func (m *DnsCacheConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsCacheConfig.Marshal(b, m, deterministic)
}
func (m *DnsCacheConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsCacheConfig.Merge(m, src)
}
func (m *DnsCacheConfig) XXX_Size() int {
	return xxx_messageInfo_DnsCacheConfig.Size(m)
}
func (m *DnsCacheConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsCacheConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DnsCacheConfig proto.InternalMessageInfo

func (m *DnsCacheConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DnsCacheConfig) GetDnsLookupFamily() v3alpha.Cluster_DnsLookupFamily {
	if m != nil {
		return m.DnsLookupFamily
	}
	return v3alpha.Cluster_AUTO
}

func (m *DnsCacheConfig) GetDnsRefreshRate() *duration.Duration {
	if m != nil {
		return m.DnsRefreshRate
	}
	return nil
}

func (m *DnsCacheConfig) GetHostTtl() *duration.Duration {
	if m != nil {
		return m.HostTtl
	}
	return nil
}

func (m *DnsCacheConfig) GetMaxHosts() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxHosts
	}
	return nil
}

func init() {
	proto.RegisterType((*DnsCacheConfig)(nil), "envoy.config.common.dynamic_forward_proxy.v3alpha.DnsCacheConfig")
}

func init() {
	proto.RegisterFile("envoy/config/common/dynamic_forward_proxy/v3alpha/dns_cache.proto", fileDescriptor_4f21fdb25ee49be2)
}

var fileDescriptor_4f21fdb25ee49be2 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4d, 0x6f, 0xd3, 0x3e,
	0x1c, 0xc7, 0x97, 0xac, 0xfb, 0xaf, 0xf3, 0x5f, 0x94, 0x91, 0x0b, 0x61, 0x43, 0x53, 0xe0, 0x54,
	0xf5, 0x60, 0x8b, 0xf6, 0xc6, 0x81, 0x87, 0xb4, 0xe2, 0x41, 0x02, 0x69, 0x8a, 0x80, 0x6b, 0xf8,
	0x2d, 0x76, 0x5a, 0x0b, 0xc7, 0xb6, 0x6c, 0xa7, 0x6b, 0xaf, 0x9c, 0x78, 0x0d, 0xbc, 0x04, 0x78,
	0x1d, 0xbc, 0xa9, 0x9d, 0x50, 0xec, 0x06, 0x69, 0x02, 0x24, 0x76, 0x8b, 0xf3, 0xf3, 0xf7, 0xe3,
	0x7c, 0x3f, 0x31, 0x7a, 0xce, 0xe4, 0x5a, 0x6d, 0x49, 0xa5, 0x64, 0xcd, 0x97, 0xa4, 0x52, 0x4d,
	0xa3, 0x24, 0xa1, 0x5b, 0x09, 0x0d, 0xaf, 0xca, 0x5a, 0x99, 0x4b, 0x30, 0xb4, 0xd4, 0x46, 0x6d,
	0xb6, 0x64, 0x3d, 0x03, 0xa1, 0x57, 0x40, 0xa8, 0xb4, 0x65, 0x05, 0xd5, 0x8a, 0x61, 0x6d, 0x94,
	0x53, 0xc9, 0x23, 0x8f, 0xc0, 0x01, 0x81, 0x03, 0x02, 0xff, 0x11, 0x81, 0x77, 0x88, 0x93, 0xd3,
	0x70, 0x2a, 0x68, 0xfe, 0x8b, 0x5a, 0x51, 0x1b, 0x78, 0x27, 0x67, 0x4b, 0xa5, 0x96, 0x82, 0x11,
	0xbf, 0xba, 0x68, 0x6b, 0x42, 0x5b, 0x03, 0x8e, 0x2b, 0xf9, 0xb7, 0xf9, 0xa5, 0x01, 0xad, 0x99,
	0xe9, 0xf3, 0x0f, 0x5a, 0xaa, 0x81, 0x80, 0x94, 0xca, 0xf9, 0x98, 0x25, 0x6b, 0x66, 0x2c, 0x57,
	0x92, 0xcb, 0xe5, 0x6e, 0xcb, 0xdd, 0x35, 0x08, 0x4e, 0xc1, 0x31, 0xd2, 0x3f, 0x84, 0xc1, 0xc3,
	0xef, 0xfb, 0x68, 0xb4, 0x90, 0x76, 0xde, 0xd5, 0x9b, 0xfb, 0x42, 0xc9, 0x29, 0x1a, 0x48, 0x68,
	0x58, 0x1a, 0x65, 0xd1, 0xf8, 0x28, 0x3f, 0xbc, 0xca, 0x07, 0x26, 0xce, 0xa2, 0xc2, 0xbf, 0x4c,
	0x3e, 0xa2, 0x3b, 0x9d, 0x0e, 0xa1, 0xd4, 0xa7, 0x56, 0x97, 0x35, 0x34, 0x5c, 0x6c, 0xd3, 0x38,
	0x8b, 0xc6, 0xa3, 0xe9, 0x04, 0x07, 0x2f, 0xa0, 0x79, 0xdf, 0x1b, 0xcf, 0x45, 0x6b, 0x1d, 0x33,
	0x78, 0x21, 0xed, 0x1b, 0x1f, 0x79, 0xe1, 0x13, 0xf9, 0xf0, 0x2a, 0x3f, 0xf8, 0x1c, 0xc5, 0xc7,
	0x51, 0x71, 0x9b, 0x5e, 0x1f, 0x25, 0x6f, 0xd1, 0x71, 0x77, 0x82, 0x61, 0xb5, 0x61, 0x76, 0x55,
	0x1a, 0x70, 0x2c, 0xdd, 0xcf, 0xa2, 0xf1, 0xff, 0xd3, 0x7b, 0x38, 0x88, 0xc0, 0xbd, 0x08, 0xbc,
	0xd8, 0x89, 0xf2, 0xbc, 0x6f, 0x51, 0x3c, 0xd9, 0x2b, 0x46, 0x54, 0xda, 0x22, 0x64, 0x0b, 0x70,
	0x2c, 0x79, 0x82, 0x86, 0x2b, 0x65, 0x5d, 0xe9, 0x9c, 0x48, 0x07, 0xff, 0x8e, 0x39, 0xec, 0x42,
	0xef, 0x9c, 0x48, 0x72, 0x74, 0xd4, 0xc0, 0xa6, 0xec, 0x96, 0x36, 0x3d, 0xf0, 0x80, 0xfb, 0xbf,
	0x01, 0xde, 0xbf, 0x96, 0x6e, 0x36, 0xfd, 0x00, 0xa2, 0x65, 0x5e, 0xd8, 0x24, 0xce, 0xf6, 0x8a,
	0x61, 0x03, 0x9b, 0x57, 0x5d, 0xec, 0xf1, 0xcb, 0xaf, 0x3f, 0xbe, 0x9c, 0xe5, 0xe8, 0xd9, 0x0d,
	0xee, 0xcd, 0x34, 0xf8, 0xbb, 0xfe, 0x6b, 0xf2, 0x02, 0x3d, 0xe5, 0x2a, 0x68, 0x0e, 0xfb, 0x6e,
	0x7c, 0x13, 0xf3, 0x5b, 0x3d, 0xf2, 0xbc, 0xfb, 0xf8, 0xf3, 0xe8, 0xe2, 0x3f, 0xdf, 0x62, 0xf6,
	0x33, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xa5, 0x44, 0x8f, 0x19, 0x03, 0x00, 0x00,
}