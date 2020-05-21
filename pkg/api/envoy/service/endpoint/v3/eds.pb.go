// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/endpoint/v3/eds.proto

package envoy_service_endpoint_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/datawire/ambassador/pkg/api/envoy/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/service/discovery/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221 and protoxform to upgrade the file.
type EdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdsDummy) Reset()         { *m = EdsDummy{} }
func (m *EdsDummy) String() string { return proto.CompactTextString(m) }
func (*EdsDummy) ProtoMessage()    {}
func (*EdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_583a91558308c82d, []int{0}
}
func (m *EdsDummy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EdsDummy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdsDummy.Merge(m, src)
}
func (m *EdsDummy) XXX_Size() int {
	return m.Size()
}
func (m *EdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_EdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_EdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EdsDummy)(nil), "envoy.service.endpoint.v3.EdsDummy")
}

func init() {
	proto.RegisterFile("envoy/service/endpoint/v3/eds.proto", fileDescriptor_583a91558308c82d)
}

var fileDescriptor_583a91558308c82d = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcd, 0x6a, 0x14, 0x41,
	0x14, 0x85, 0xad, 0x04, 0xa2, 0xd4, 0x22, 0xca, 0x80, 0x8c, 0x69, 0x62, 0xa3, 0xe3, 0xc2, 0x21,
	0x68, 0x55, 0x9c, 0x01, 0x17, 0xb3, 0x33, 0x4e, 0x5c, 0x65, 0x11, 0xcc, 0x13, 0x54, 0xba, 0x6e,
	0xda, 0x82, 0x99, 0xba, 0x65, 0xfd, 0xe9, 0x2c, 0x04, 0x71, 0x15, 0xdc, 0x0a, 0x0a, 0xae, 0x7d,
	0x0a, 0xf7, 0x42, 0x96, 0x8a, 0x2f, 0xa0, 0x83, 0x6f, 0xe0, 0x0b, 0xc8, 0x74, 0x77, 0x25, 0x69,
	0x42, 0x44, 0x37, 0xd9, 0xf5, 0xe5, 0x7c, 0xf7, 0xdc, 0xee, 0xc3, 0x69, 0x7a, 0x07, 0x74, 0xc4,
	0x19, 0x77, 0x60, 0xa3, 0x2a, 0x80, 0x83, 0x96, 0x06, 0x95, 0xf6, 0x3c, 0x0e, 0x39, 0x48, 0xc7,
	0x8c, 0x45, 0x8f, 0x9d, 0xb5, 0x0a, 0x62, 0x0d, 0xc4, 0x12, 0xc4, 0xe2, 0x30, 0xdb, 0x68, 0xef,
	0x4b, 0xe5, 0x0a, 0x8c, 0x60, 0x67, 0x0b, 0x83, 0xe3, 0xa1, 0xb6, 0xc9, 0xd6, 0x4b, 0xc4, 0x72,
	0x02, 0x5c, 0x18, 0xc5, 0x85, 0xd6, 0xe8, 0x85, 0x57, 0xa8, 0x9b, 0x23, 0x59, 0xde, 0xa8, 0xd5,
	0xb4, 0x1f, 0x0e, 0xb8, 0x0c, 0xb6, 0x02, 0xce, 0xd3, 0x5f, 0x58, 0x61, 0x0c, 0xd8, 0xb4, 0x7f,
	0xab, 0x7e, 0x93, 0x53, 0xc6, 0xdc, 0x82, 0xc3, 0x60, 0x0b, 0x68, 0x88, 0x9b, 0x41, 0x1a, 0xd1,
	0x02, 0x9c, 0x17, 0x3e, 0x24, 0x83, 0xdb, 0x67, 0xe4, 0x08, 0xd6, 0x29, 0xd4, 0x4a, 0x97, 0x0d,
	0xd2, 0x8d, 0x62, 0xa2, 0xa4, 0xf0, 0xc0, 0xd3, 0x43, 0x2d, 0xf4, 0xfa, 0xf4, 0xca, 0xb6, 0x74,
	0xe3, 0x30, 0x9d, 0xce, 0x46, 0xeb, 0x1f, 0xbf, 0x1c, 0xe6, 0x5d, 0x7a, 0xbd, 0x0e, 0x4d, 0x18,
	0xc5, 0xe2, 0x80, 0x25, 0x75, 0xf0, 0x73, 0x99, 0xde, 0xd8, 0x6e, 0x02, 0x1c, 0xa7, 0x80, 0xf6,
	0xea, 0xf8, 0x3a, 0x9e, 0x5e, 0xdd, 0xf3, 0x16, 0xc4, 0x34, 0x11, 0xae, 0x73, 0x8f, 0xb5, 0xc3,
	0x3f, 0x09, 0x35, 0x0e, 0xd9, 0xb1, 0xc1, 0x53, 0x78, 0x1e, 0xc0, 0xf9, 0xec, 0xfe, 0x3f, 0xd2,
	0xce, 0xa0, 0x76, 0xd0, 0xbb, 0xd4, 0x27, 0x9b, 0xa4, 0xf3, 0x8a, 0xae, 0x8e, 0x61, 0xe2, 0xc5,
	0xc9, 0xd1, 0x07, 0x7f, 0xb5, 0x59, 0xb0, 0x67, 0x2e, 0x0f, 0xfe, 0x67, 0xa5, 0x75, 0xfe, 0x03,
	0xa1, 0xab, 0x4f, 0xc0, 0x17, 0xcf, 0x2e, 0xe8, 0xa3, 0xfb, 0x6f, 0xbe, 0xff, 0x7a, 0xb7, 0xb4,
	0xd6, 0xeb, 0xb6, 0x3a, 0x3a, 0x4a, 0xad, 0x76, 0x95, 0xbc, 0x3c, 0x22, 0x1b, 0xd9, 0xc3, 0xb7,
	0x9f, 0xde, 0xff, 0xbe, 0xbc, 0x49, 0x1b, 0xff, 0x02, 0xf5, 0x81, 0x2a, 0x4f, 0xd7, 0x9f, 0x3d,
	0x9e, 0x04, 0xe7, 0xc1, 0xee, 0xa0, 0x90, 0x8f, 0x9c, 0x53, 0xa5, 0x9e, 0x82, 0xf6, 0x5b, 0x3b,
	0x47, 0xf3, 0x9c, 0x7c, 0x9d, 0xe7, 0xe4, 0xc7, 0x3c, 0x27, 0x9f, 0x5f, 0x1f, 0x7d, 0x5b, 0x59,
	0xba, 0x46, 0xe8, 0x5d, 0x85, 0xb5, 0x91, 0xb1, 0xf8, 0x72, 0xc6, 0xce, 0xfd, 0xa7, 0xb6, 0x16,
	0x55, 0xda, 0x5d, 0xd4, 0x6a, 0x97, 0x1c, 0x12, 0xb2, 0xbf, 0x52, 0x55, 0x6c, 0xf8, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x5e, 0x9b, 0x54, 0xc2, 0xab, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndpointDiscoveryServiceClient is the client API for EndpointDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointDiscoveryServiceClient interface {
	// The resource_names field in DiscoveryRequest specifies a list of clusters
	// to subscribe to updates for.
	StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error)
	DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error)
	FetchEndpoints(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type endpointDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointDiscoveryServiceClient(cc *grpc.ClientConn) EndpointDiscoveryServiceClient {
	return &endpointDiscoveryServiceClient{cc}
}

func (c *endpointDiscoveryServiceClient) StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[0], "/envoy.service.endpoint.v3.EndpointDiscoveryService/StreamEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamEndpointsClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[1], "/envoy.service.endpoint.v3.EndpointDiscoveryService/DeltaEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceDeltaEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_DeltaEndpointsClient interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceDeltaEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) FetchEndpoints(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.endpoint.v3.EndpointDiscoveryService/FetchEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointDiscoveryServiceServer is the server API for EndpointDiscoveryService service.
type EndpointDiscoveryServiceServer interface {
	// The resource_names field in DiscoveryRequest specifies a list of clusters
	// to subscribe to updates for.
	StreamEndpoints(EndpointDiscoveryService_StreamEndpointsServer) error
	DeltaEndpoints(EndpointDiscoveryService_DeltaEndpointsServer) error
	FetchEndpoints(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
}

// UnimplementedEndpointDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEndpointDiscoveryServiceServer struct {
}

func (*UnimplementedEndpointDiscoveryServiceServer) StreamEndpoints(srv EndpointDiscoveryService_StreamEndpointsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEndpoints not implemented")
}
func (*UnimplementedEndpointDiscoveryServiceServer) DeltaEndpoints(srv EndpointDiscoveryService_DeltaEndpointsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaEndpoints not implemented")
}
func (*UnimplementedEndpointDiscoveryServiceServer) FetchEndpoints(ctx context.Context, req *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchEndpoints not implemented")
}

func RegisterEndpointDiscoveryServiceServer(s *grpc.Server, srv EndpointDiscoveryServiceServer) {
	s.RegisterService(&_EndpointDiscoveryService_serviceDesc, srv)
}

func _EndpointDiscoveryService_StreamEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamEndpoints(&endpointDiscoveryServiceStreamEndpointsServer{stream})
}

type EndpointDiscoveryService_StreamEndpointsServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_DeltaEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).DeltaEndpoints(&endpointDiscoveryServiceDeltaEndpointsServer{stream})
}

type EndpointDiscoveryService_DeltaEndpointsServer interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceDeltaEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_FetchEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.endpoint.v3.EndpointDiscoveryService/FetchEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.endpoint.v3.EndpointDiscoveryService",
	HandlerType: (*EndpointDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchEndpoints",
			Handler:    _EndpointDiscoveryService_FetchEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEndpoints",
			Handler:       _EndpointDiscoveryService_StreamEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaEndpoints",
			Handler:       _EndpointDiscoveryService_DeltaEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/endpoint/v3/eds.proto",
}

func (m *EdsDummy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EdsDummy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EdsDummy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintEds(dAtA []byte, offset int, v uint64) int {
	offset -= sovEds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EdsDummy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEds(x uint64) (n int) {
	return sovEds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EdsDummy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEds
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EdsDummy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EdsDummy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipEds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEds
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEds
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEds
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEds
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEds = fmt.Errorf("proto: unexpected end of group")
)
