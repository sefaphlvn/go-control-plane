// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/discovery/v3alpha/sds.proto

package envoy_service_discovery_v3alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SdsDummy) Reset()         { *m = SdsDummy{} }
func (m *SdsDummy) String() string { return proto.CompactTextString(m) }
func (*SdsDummy) ProtoMessage()    {}
func (*SdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_1969cf11de54a1fa, []int{0}
}

func (m *SdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SdsDummy.Unmarshal(m, b)
}
func (m *SdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SdsDummy.Marshal(b, m, deterministic)
}
func (m *SdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SdsDummy.Merge(m, src)
}
func (m *SdsDummy) XXX_Size() int {
	return xxx_messageInfo_SdsDummy.Size(m)
}
func (m *SdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_SdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_SdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SdsDummy)(nil), "envoy.service.discovery.v3alpha.SdsDummy")
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v3alpha/sds.proto", fileDescriptor_1969cf11de54a1fa)
}

var fileDescriptor_1969cf11de54a1fa = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0x3f, 0xf7, 0x93, 0x10, 0xb2, 0xca, 0x92, 0x81, 0x21, 0xaa, 0x40, 0x6d, 0x41, 0x6a,
	0x8b, 0x70, 0x50, 0x2b, 0x31, 0x74, 0x60, 0xa8, 0x2a, 0xe6, 0x8a, 0xdc, 0x00, 0x26, 0x39, 0x6a,
	0x2d, 0xb5, 0x3e, 0xc6, 0x76, 0x23, 0xb2, 0x32, 0xa0, 0xee, 0x6c, 0xdc, 0x0f, 0x57, 0xc0, 0x2d,
	0x30, 0x73, 0x0d, 0xa8, 0x76, 0x12, 0xfe, 0x7f, 0x16, 0x56, 0x9f, 0xe7, 0xbc, 0x8f, 0xad, 0xd7,
	0xb4, 0x0b, 0x32, 0xc3, 0x3c, 0x32, 0xa0, 0x33, 0x91, 0x40, 0x94, 0x0a, 0x93, 0x60, 0x06, 0x3a,
	0x8f, 0xb2, 0x01, 0x9f, 0xab, 0x19, 0x8f, 0x4c, 0x6a, 0x98, 0xd2, 0x68, 0x31, 0xd8, 0x75, 0x28,
	0x2b, 0x50, 0x56, 0xa1, 0xac, 0x40, 0xc3, 0xa6, 0xcf, 0xe2, 0x4a, 0x54, 0xdb, 0x2f, 0x90, 0xcb,
	0x08, 0x1b, 0x53, 0xc4, 0xe9, 0x1c, 0x1c, 0xc3, 0xa5, 0x44, 0xcb, 0xad, 0x40, 0x59, 0x18, 0xc2,
	0xe6, 0x32, 0x55, 0xfc, 0xf5, 0x79, 0x94, 0x81, 0x36, 0x02, 0xa5, 0x90, 0x53, 0x8f, 0xb4, 0x8e,
	0xe9, 0x66, 0x9c, 0x9a, 0xf1, 0x72, 0xb1, 0xc8, 0x87, 0xbd, 0xbb, 0xfb, 0xd5, 0xce, 0x3e, 0x6d,
	0x7f, 0x79, 0xaf, 0x3e, 0x2b, 0xd9, 0xfe, 0x53, 0x8d, 0x6e, 0xc7, 0x90, 0x68, 0xb0, 0xe3, 0x72,
	0x1e, 0xfb, 0x85, 0x40, 0xd0, 0xfa, 0x18, 0xe6, 0x96, 0xfb, 0xb1, 0x09, 0x3a, 0xcc, 0x07, 0x72,
	0x25, 0xca, 0xa7, 0x31, 0x07, 0x54, 0x9b, 0x67, 0x70, 0xb9, 0x04, 0x63, 0xc3, 0xee, 0x2f, 0x48,
	0xa3, 0x50, 0x1a, 0x68, 0xfd, 0xeb, 0x90, 0x23, 0x12, 0x9c, 0xd3, 0xad, 0xd8, 0x6a, 0xe0, 0x8b,
	0xd2, 0xd5, 0xfe, 0x2c, 0xe1, 0xbd, 0x66, 0xef, 0x7b, 0xe8, 0x8d, 0xe1, 0x86, 0xd0, 0xfa, 0x29,
	0xd8, 0x64, 0xf6, 0x07, 0x86, 0x83, 0xeb, 0x87, 0xc7, 0xdb, 0x5a, 0xa3, 0x15, 0x7e, 0xac, 0x77,
	0x68, 0xbc, 0xce, 0x11, 0xff, 0x87, 0xa4, 0x37, 0x3a, 0xa1, 0x87, 0x02, 0x7d, 0xac, 0xd2, 0x78,
	0x95, 0xb3, 0x1f, 0x7e, 0xcf, 0x68, 0xdd, 0xeb, 0x64, 0xdd, 0xf1, 0x84, 0xac, 0x08, 0xb9, 0xd8,
	0x70, 0x7d, 0x0f, 0x9e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x1c, 0x4d, 0x01, 0xa1, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecretDiscoveryServiceClient is the client API for SecretDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretDiscoveryServiceClient interface {
	DeltaSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_DeltaSecretsClient, error)
	StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error)
	FetchSecrets(ctx context.Context, in *v3alpha.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha.DiscoveryResponse, error)
}

type secretDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecretDiscoveryServiceClient(cc *grpc.ClientConn) SecretDiscoveryServiceClient {
	return &secretDiscoveryServiceClient{cc}
}

func (c *secretDiscoveryServiceClient) DeltaSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_DeltaSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SecretDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v3alpha.SecretDiscoveryService/DeltaSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretDiscoveryServiceDeltaSecretsClient{stream}
	return x, nil
}

type SecretDiscoveryService_DeltaSecretsClient interface {
	Send(*v3alpha.DeltaDiscoveryRequest) error
	Recv() (*v3alpha.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type secretDiscoveryServiceDeltaSecretsClient struct {
	grpc.ClientStream
}

func (x *secretDiscoveryServiceDeltaSecretsClient) Send(m *v3alpha.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretDiscoveryServiceDeltaSecretsClient) Recv() (*v3alpha.DeltaDiscoveryResponse, error) {
	m := new(v3alpha.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretDiscoveryServiceClient) StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SecretDiscoveryService_serviceDesc.Streams[1], "/envoy.service.discovery.v3alpha.SecretDiscoveryService/StreamSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretDiscoveryServiceStreamSecretsClient{stream}
	return x, nil
}

type SecretDiscoveryService_StreamSecretsClient interface {
	Send(*v3alpha.DiscoveryRequest) error
	Recv() (*v3alpha.DiscoveryResponse, error)
	grpc.ClientStream
}

type secretDiscoveryServiceStreamSecretsClient struct {
	grpc.ClientStream
}

func (x *secretDiscoveryServiceStreamSecretsClient) Send(m *v3alpha.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsClient) Recv() (*v3alpha.DiscoveryResponse, error) {
	m := new(v3alpha.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretDiscoveryServiceClient) FetchSecrets(ctx context.Context, in *v3alpha.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha.DiscoveryResponse, error) {
	out := new(v3alpha.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.discovery.v3alpha.SecretDiscoveryService/FetchSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretDiscoveryServiceServer is the server API for SecretDiscoveryService service.
type SecretDiscoveryServiceServer interface {
	DeltaSecrets(SecretDiscoveryService_DeltaSecretsServer) error
	StreamSecrets(SecretDiscoveryService_StreamSecretsServer) error
	FetchSecrets(context.Context, *v3alpha.DiscoveryRequest) (*v3alpha.DiscoveryResponse, error)
}

// UnimplementedSecretDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSecretDiscoveryServiceServer struct {
}

func (*UnimplementedSecretDiscoveryServiceServer) DeltaSecrets(srv SecretDiscoveryService_DeltaSecretsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaSecrets not implemented")
}
func (*UnimplementedSecretDiscoveryServiceServer) StreamSecrets(srv SecretDiscoveryService_StreamSecretsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSecrets not implemented")
}
func (*UnimplementedSecretDiscoveryServiceServer) FetchSecrets(ctx context.Context, req *v3alpha.DiscoveryRequest) (*v3alpha.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchSecrets not implemented")
}

func RegisterSecretDiscoveryServiceServer(s *grpc.Server, srv SecretDiscoveryServiceServer) {
	s.RegisterService(&_SecretDiscoveryService_serviceDesc, srv)
}

func _SecretDiscoveryService_DeltaSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretDiscoveryServiceServer).DeltaSecrets(&secretDiscoveryServiceDeltaSecretsServer{stream})
}

type SecretDiscoveryService_DeltaSecretsServer interface {
	Send(*v3alpha.DeltaDiscoveryResponse) error
	Recv() (*v3alpha.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type secretDiscoveryServiceDeltaSecretsServer struct {
	grpc.ServerStream
}

func (x *secretDiscoveryServiceDeltaSecretsServer) Send(m *v3alpha.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretDiscoveryServiceDeltaSecretsServer) Recv() (*v3alpha.DeltaDiscoveryRequest, error) {
	m := new(v3alpha.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretDiscoveryService_StreamSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretDiscoveryServiceServer).StreamSecrets(&secretDiscoveryServiceStreamSecretsServer{stream})
}

type SecretDiscoveryService_StreamSecretsServer interface {
	Send(*v3alpha.DiscoveryResponse) error
	Recv() (*v3alpha.DiscoveryRequest, error)
	grpc.ServerStream
}

type secretDiscoveryServiceStreamSecretsServer struct {
	grpc.ServerStream
}

func (x *secretDiscoveryServiceStreamSecretsServer) Send(m *v3alpha.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsServer) Recv() (*v3alpha.DiscoveryRequest, error) {
	m := new(v3alpha.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretDiscoveryService_FetchSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3alpha.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.discovery.v3alpha.SecretDiscoveryService/FetchSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, req.(*v3alpha.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v3alpha.SecretDiscoveryService",
	HandlerType: (*SecretDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchSecrets",
			Handler:    _SecretDiscoveryService_FetchSecrets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaSecrets",
			Handler:       _SecretDiscoveryService_DeltaSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamSecrets",
			Handler:       _SecretDiscoveryService_StreamSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v3alpha/sds.proto",
}