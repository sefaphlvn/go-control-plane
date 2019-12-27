// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/auth/v3alpha/attribute_context.proto

package envoy_service_auth_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type AttributeContext struct {
	Source               *AttributeContext_Peer    `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Destination          *AttributeContext_Peer    `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Request              *AttributeContext_Request `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
	ContextExtensions    map[string]string         `protobuf:"bytes,10,rep,name=context_extensions,json=contextExtensions,proto3" json:"context_extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MetadataContext      *core.Metadata            `protobuf:"bytes,11,opt,name=metadata_context,json=metadataContext,proto3" json:"metadata_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AttributeContext) Reset()         { *m = AttributeContext{} }
func (m *AttributeContext) String() string { return proto.CompactTextString(m) }
func (*AttributeContext) ProtoMessage()    {}
func (*AttributeContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_000310fa99e78275, []int{0}
}

func (m *AttributeContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext.Unmarshal(m, b)
}
func (m *AttributeContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext.Marshal(b, m, deterministic)
}
func (m *AttributeContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext.Merge(m, src)
}
func (m *AttributeContext) XXX_Size() int {
	return xxx_messageInfo_AttributeContext.Size(m)
}
func (m *AttributeContext) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext proto.InternalMessageInfo

func (m *AttributeContext) GetSource() *AttributeContext_Peer {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *AttributeContext) GetDestination() *AttributeContext_Peer {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *AttributeContext) GetRequest() *AttributeContext_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *AttributeContext) GetContextExtensions() map[string]string {
	if m != nil {
		return m.ContextExtensions
	}
	return nil
}

func (m *AttributeContext) GetMetadataContext() *core.Metadata {
	if m != nil {
		return m.MetadataContext
	}
	return nil
}

type AttributeContext_Peer struct {
	Address              *core.Address     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Service              string            `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Principal            string            `protobuf:"bytes,4,opt,name=principal,proto3" json:"principal,omitempty"`
	Certificate          string            `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AttributeContext_Peer) Reset()         { *m = AttributeContext_Peer{} }
func (m *AttributeContext_Peer) String() string { return proto.CompactTextString(m) }
func (*AttributeContext_Peer) ProtoMessage()    {}
func (*AttributeContext_Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_000310fa99e78275, []int{0, 0}
}

func (m *AttributeContext_Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext_Peer.Unmarshal(m, b)
}
func (m *AttributeContext_Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext_Peer.Marshal(b, m, deterministic)
}
func (m *AttributeContext_Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext_Peer.Merge(m, src)
}
func (m *AttributeContext_Peer) XXX_Size() int {
	return xxx_messageInfo_AttributeContext_Peer.Size(m)
}
func (m *AttributeContext_Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext_Peer proto.InternalMessageInfo

func (m *AttributeContext_Peer) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *AttributeContext_Peer) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *AttributeContext_Peer) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *AttributeContext_Peer) GetPrincipal() string {
	if m != nil {
		return m.Principal
	}
	return ""
}

func (m *AttributeContext_Peer) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

type AttributeContext_Request struct {
	Time                 *timestamp.Timestamp          `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Http                 *AttributeContext_HttpRequest `protobuf:"bytes,2,opt,name=http,proto3" json:"http,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AttributeContext_Request) Reset()         { *m = AttributeContext_Request{} }
func (m *AttributeContext_Request) String() string { return proto.CompactTextString(m) }
func (*AttributeContext_Request) ProtoMessage()    {}
func (*AttributeContext_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_000310fa99e78275, []int{0, 1}
}

func (m *AttributeContext_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext_Request.Unmarshal(m, b)
}
func (m *AttributeContext_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext_Request.Marshal(b, m, deterministic)
}
func (m *AttributeContext_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext_Request.Merge(m, src)
}
func (m *AttributeContext_Request) XXX_Size() int {
	return xxx_messageInfo_AttributeContext_Request.Size(m)
}
func (m *AttributeContext_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext_Request.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext_Request proto.InternalMessageInfo

func (m *AttributeContext_Request) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *AttributeContext_Request) GetHttp() *AttributeContext_HttpRequest {
	if m != nil {
		return m.Http
	}
	return nil
}

type AttributeContext_HttpRequest struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Method               string            `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Headers              map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Path                 string            `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Host                 string            `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	Scheme               string            `protobuf:"bytes,6,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Query                string            `protobuf:"bytes,7,opt,name=query,proto3" json:"query,omitempty"`
	Fragment             string            `protobuf:"bytes,8,opt,name=fragment,proto3" json:"fragment,omitempty"`
	Size                 int64             `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	Protocol             string            `protobuf:"bytes,10,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Body                 string            `protobuf:"bytes,11,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AttributeContext_HttpRequest) Reset()         { *m = AttributeContext_HttpRequest{} }
func (m *AttributeContext_HttpRequest) String() string { return proto.CompactTextString(m) }
func (*AttributeContext_HttpRequest) ProtoMessage()    {}
func (*AttributeContext_HttpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_000310fa99e78275, []int{0, 2}
}

func (m *AttributeContext_HttpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeContext_HttpRequest.Unmarshal(m, b)
}
func (m *AttributeContext_HttpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeContext_HttpRequest.Marshal(b, m, deterministic)
}
func (m *AttributeContext_HttpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeContext_HttpRequest.Merge(m, src)
}
func (m *AttributeContext_HttpRequest) XXX_Size() int {
	return xxx_messageInfo_AttributeContext_HttpRequest.Size(m)
}
func (m *AttributeContext_HttpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeContext_HttpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeContext_HttpRequest proto.InternalMessageInfo

func (m *AttributeContext_HttpRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *AttributeContext_HttpRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetFragment() string {
	if m != nil {
		return m.Fragment
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *AttributeContext_HttpRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *AttributeContext_HttpRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*AttributeContext)(nil), "envoy.service.auth.v3alpha.AttributeContext")
	proto.RegisterMapType((map[string]string)(nil), "envoy.service.auth.v3alpha.AttributeContext.ContextExtensionsEntry")
	proto.RegisterType((*AttributeContext_Peer)(nil), "envoy.service.auth.v3alpha.AttributeContext.Peer")
	proto.RegisterMapType((map[string]string)(nil), "envoy.service.auth.v3alpha.AttributeContext.Peer.LabelsEntry")
	proto.RegisterType((*AttributeContext_Request)(nil), "envoy.service.auth.v3alpha.AttributeContext.Request")
	proto.RegisterType((*AttributeContext_HttpRequest)(nil), "envoy.service.auth.v3alpha.AttributeContext.HttpRequest")
	proto.RegisterMapType((map[string]string)(nil), "envoy.service.auth.v3alpha.AttributeContext.HttpRequest.HeadersEntry")
}

func init() {
	proto.RegisterFile("envoy/service/auth/v3alpha/attribute_context.proto", fileDescriptor_000310fa99e78275)
}

var fileDescriptor_000310fa99e78275 = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x4e, 0x1b, 0x47,
	0x14, 0x96, 0x7f, 0xb0, 0xf1, 0x31, 0x6a, 0xe9, 0xa8, 0x45, 0xab, 0x55, 0x55, 0xdc, 0xaa, 0xaa,
	0x2c, 0x55, 0xcc, 0xb6, 0x86, 0x48, 0x60, 0x29, 0x17, 0x40, 0x90, 0x88, 0x42, 0x22, 0xb4, 0x49,
	0xae, 0xd1, 0x78, 0xf7, 0xe0, 0x1d, 0xc5, 0xde, 0x59, 0x66, 0x66, 0x2d, 0x9c, 0x27, 0xc8, 0x33,
	0xe4, 0x36, 0xef, 0x91, 0x9b, 0x28, 0x0f, 0x92, 0x37, 0x89, 0xe6, 0x67, 0x1d, 0x0b, 0x41, 0x82,
	0xb9, 0x62, 0xce, 0xec, 0x39, 0xdf, 0x9c, 0xef, 0xfb, 0xce, 0xc1, 0x30, 0xc0, 0x7c, 0x26, 0xe6,
	0x91, 0x42, 0x39, 0xe3, 0x09, 0x46, 0xac, 0xd4, 0x59, 0x34, 0xdb, 0x65, 0x93, 0x22, 0x63, 0x11,
	0xd3, 0x5a, 0xf2, 0x51, 0xa9, 0xf1, 0x22, 0x11, 0xb9, 0xc6, 0x6b, 0x4d, 0x0b, 0x29, 0xb4, 0x20,
	0xa1, 0xad, 0xa1, 0xbe, 0x86, 0x9a, 0x1a, 0xea, 0x6b, 0xc2, 0xbf, 0x1d, 0x1e, 0x2b, 0xf8, 0x02,
	0x26, 0x11, 0x12, 0x23, 0x96, 0xa6, 0x12, 0x95, 0x72, 0x08, 0xe1, 0x9f, 0x77, 0x64, 0x8d, 0x98,
	0x42, 0x9f, 0xb2, 0x3d, 0x16, 0x62, 0x3c, 0xc1, 0xc8, 0x46, 0xa3, 0xf2, 0x32, 0xd2, 0x7c, 0x8a,
	0x4a, 0xb3, 0x69, 0x51, 0x61, 0x94, 0x69, 0xc1, 0x22, 0x96, 0xe7, 0x42, 0x33, 0xcd, 0x45, 0xae,
	0xa2, 0x19, 0x4a, 0xc5, 0x45, 0xce, 0xf3, 0xb1, 0x4b, 0xf9, 0xeb, 0xc3, 0x06, 0x6c, 0x1e, 0x56,
	0x24, 0x8e, 0x1d, 0x07, 0xf2, 0x14, 0x5a, 0x4a, 0x94, 0x32, 0xc1, 0xa0, 0xd6, 0xab, 0xf5, 0xbb,
	0x83, 0xff, 0xe9, 0xdd, 0x74, 0xe8, 0xcd, 0x6a, 0x7a, 0x8e, 0x28, 0x63, 0x0f, 0x40, 0x5e, 0x42,
	0x37, 0x45, 0xa5, 0x79, 0x6e, 0x1b, 0x08, 0xea, 0x0f, 0xc5, 0x5b, 0x46, 0x21, 0x2f, 0xa0, 0x2d,
	0xf1, 0xaa, 0x44, 0xa5, 0x83, 0xa6, 0x05, 0xdc, 0x5b, 0x09, 0x30, 0x76, 0xb5, 0x71, 0x05, 0x42,
	0x24, 0x10, 0x6f, 0xdf, 0x05, 0x5e, 0x6b, 0xcc, 0x8d, 0x44, 0x2a, 0x80, 0x5e, 0xa3, 0xdf, 0x1d,
	0x1c, 0xaf, 0x04, 0xed, 0xff, 0x9e, 0x2c, 0x50, 0x4e, 0x72, 0x2d, 0xe7, 0xf1, 0x2f, 0xc9, 0xcd,
	0x7b, 0xf2, 0x0c, 0x36, 0xa7, 0xa8, 0x59, 0xca, 0x34, 0xab, 0x66, 0x27, 0xe8, 0x5a, 0x32, 0x3d,
	0xff, 0x22, 0x2b, 0xf8, 0xe2, 0x21, 0x63, 0x3d, 0x7d, 0xee, 0xf3, 0xe3, 0x9f, 0xab, 0x4a, 0xff,
	0x5c, 0xf8, 0xa5, 0x0e, 0x4d, 0x23, 0x13, 0x39, 0x80, 0xb6, 0x1f, 0x23, 0x6f, 0xdd, 0xf6, 0x5d,
	0x60, 0x87, 0x2e, 0x2d, 0xae, 0xf2, 0x49, 0x00, 0x6d, 0xcf, 0xd1, 0xba, 0xd4, 0x89, 0xab, 0x90,
	0xbc, 0x86, 0xd6, 0x84, 0x8d, 0x70, 0xa2, 0x82, 0x86, 0x95, 0xe4, 0xf1, 0xca, 0xf6, 0xd1, 0x33,
	0x5b, 0xef, 0xc4, 0xf0, 0x60, 0xe4, 0x77, 0xe8, 0x14, 0x92, 0xe7, 0x09, 0x2f, 0xd8, 0xc4, 0xfa,
	0xd8, 0x89, 0xbf, 0x5d, 0x90, 0x1e, 0x74, 0x13, 0x94, 0x9a, 0x5f, 0xf2, 0x84, 0x69, 0x0c, 0xd6,
	0xec, 0xf7, 0xe5, 0xab, 0xf0, 0x00, 0xba, 0x4b, 0xb0, 0x64, 0x13, 0x1a, 0x6f, 0x70, 0x6e, 0x69,
	0x77, 0x62, 0x73, 0x24, 0xbf, 0xc2, 0xda, 0x8c, 0x4d, 0xca, 0x8a, 0x8f, 0x0b, 0x86, 0xf5, 0xfd,
	0xda, 0x70, 0xf0, 0xfe, 0xf3, 0xbb, 0x3f, 0x76, 0xe0, 0xdf, 0xdb, 0x78, 0x0c, 0x6e, 0xa7, 0x10,
	0x7e, 0xac, 0x41, 0xdb, 0x4f, 0x0e, 0xa1, 0xd0, 0x34, 0xbb, 0xe6, 0x35, 0x0e, 0xa9, 0x5b, 0x44,
	0x5a, 0x2d, 0x22, 0x7d, 0x55, 0x2d, 0x62, 0x6c, 0xf3, 0xc8, 0x19, 0x34, 0x33, 0xad, 0x0b, 0x3f,
	0xfe, 0xfb, 0x2b, 0xe9, 0x77, 0xaa, 0x75, 0x51, 0x4d, 0xac, 0x45, 0x19, 0x3e, 0x32, 0xdd, 0xff,
	0x07, 0xf4, 0x9e, 0xdd, 0xfb, 0xe2, 0xf0, 0x53, 0x03, 0xba, 0x4b, 0x60, 0xe4, 0x27, 0xa8, 0xf3,
	0xd4, 0xeb, 0x55, 0xe7, 0x29, 0xd9, 0x82, 0xd6, 0x14, 0x75, 0x26, 0x52, 0xaf, 0x97, 0x8f, 0xc8,
	0x05, 0xb4, 0x33, 0x64, 0x29, 0xca, 0xca, 0xff, 0x93, 0x87, 0xf6, 0x4f, 0x4f, 0x1d, 0x8e, 0x9b,
	0x83, 0x0a, 0x95, 0x10, 0x68, 0x16, 0x4c, 0x67, 0x7e, 0x06, 0xec, 0xd9, 0xdc, 0x65, 0x42, 0x69,
	0xef, 0xbb, 0x3d, 0x9b, 0x06, 0x55, 0x92, 0xe1, 0x14, 0x83, 0x96, 0x6b, 0xd0, 0x45, 0xc6, 0xe7,
	0xab, 0x12, 0xe5, 0x3c, 0x68, 0x3b, 0x9f, 0x6d, 0x40, 0x42, 0x58, 0xbf, 0x94, 0x6c, 0x3c, 0xc5,
	0x5c, 0x07, 0xeb, 0xf6, 0xc3, 0x22, 0x36, 0xe8, 0x8a, 0xbf, 0xc5, 0xa0, 0xd3, 0xab, 0xf5, 0x1b,
	0xb1, 0x3d, 0x9b, 0x7c, 0xeb, 0x5f, 0x22, 0x26, 0x01, 0xb8, 0xfc, 0x2a, 0x36, 0xf9, 0x23, 0x91,
	0xce, 0xed, 0x82, 0x76, 0x62, 0x7b, 0x0e, 0x87, 0xb0, 0xb1, 0x4c, 0x67, 0xa5, 0xf9, 0x3b, 0x30,
	0x0e, 0xee, 0xf9, 0x5f, 0x96, 0x1f, 0x3b, 0xb8, 0x24, 0x61, 0xf8, 0x04, 0xb6, 0x6e, 0xff, 0x27,
	0xb3, 0x52, 0x03, 0x3b, 0xa6, 0x81, 0x3e, 0xfc, 0x73, 0xbf, 0x06, 0x8e, 0x8e, 0xa1, 0xcf, 0x85,
	0x73, 0xbd, 0x90, 0xe2, 0x7a, 0xfe, 0x9d, 0x01, 0x38, 0xfa, 0xed, 0x66, 0xf5, 0xb9, 0x51, 0xf1,
	0xbc, 0x36, 0x6a, 0x59, 0x39, 0x77, 0xbf, 0x06, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x65, 0x2f, 0x8d,
	0x50, 0x07, 0x00, 0x00,
}