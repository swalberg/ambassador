// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/network/thrift_proxy/v3/thrift_proxy.proto

package envoy_extensions_filters_network_thrift_proxy_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
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

// Thrift transport types supported by Envoy.
type TransportType int32

const (
	// For downstream connections, the Thrift proxy will attempt to determine which transport to use.
	// For upstream connections, the Thrift proxy will use same transport as the downstream
	// connection.
	TransportType_AUTO_TRANSPORT TransportType = 0
	// The Thrift proxy will use the Thrift framed transport.
	TransportType_FRAMED TransportType = 1
	// The Thrift proxy will use the Thrift unframed transport.
	TransportType_UNFRAMED TransportType = 2
	// The Thrift proxy will assume the client is using the Thrift header transport.
	TransportType_HEADER TransportType = 3
)

var TransportType_name = map[int32]string{
	0: "AUTO_TRANSPORT",
	1: "FRAMED",
	2: "UNFRAMED",
	3: "HEADER",
}

var TransportType_value = map[string]int32{
	"AUTO_TRANSPORT": 0,
	"FRAMED":         1,
	"UNFRAMED":       2,
	"HEADER":         3,
}

func (x TransportType) String() string {
	return proto.EnumName(TransportType_name, int32(x))
}

func (TransportType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a5befaf4ce0c1223, []int{0}
}

// Thrift Protocol types supported by Envoy.
type ProtocolType int32

const (
	// For downstream connections, the Thrift proxy will attempt to determine which protocol to use.
	// Note that the older, non-strict (or lax) binary protocol is not included in automatic protocol
	// detection. For upstream connections, the Thrift proxy will use the same protocol as the
	// downstream connection.
	ProtocolType_AUTO_PROTOCOL ProtocolType = 0
	// The Thrift proxy will use the Thrift binary protocol.
	ProtocolType_BINARY ProtocolType = 1
	// The Thrift proxy will use Thrift non-strict binary protocol.
	ProtocolType_LAX_BINARY ProtocolType = 2
	// The Thrift proxy will use the Thrift compact protocol.
	ProtocolType_COMPACT ProtocolType = 3
	// The Thrift proxy will use the Thrift "Twitter" protocol implemented by the finagle library.
	ProtocolType_TWITTER ProtocolType = 4
)

var ProtocolType_name = map[int32]string{
	0: "AUTO_PROTOCOL",
	1: "BINARY",
	2: "LAX_BINARY",
	3: "COMPACT",
	4: "TWITTER",
}

var ProtocolType_value = map[string]int32{
	"AUTO_PROTOCOL": 0,
	"BINARY":        1,
	"LAX_BINARY":    2,
	"COMPACT":       3,
	"TWITTER":       4,
}

func (x ProtocolType) String() string {
	return proto.EnumName(ProtocolType_name, int32(x))
}

func (ProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a5befaf4ce0c1223, []int{1}
}

// [#next-free-field: 6]
type ThriftProxy struct {
	// Supplies the type of transport that the Thrift proxy should use. Defaults to
	// :ref:`AUTO_TRANSPORT<envoy_api_enum_value_extensions.filters.network.thrift_proxy.v3.TransportType.AUTO_TRANSPORT>`.
	Transport TransportType `protobuf:"varint,2,opt,name=transport,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.TransportType" json:"transport,omitempty"`
	// Supplies the type of protocol that the Thrift proxy should use. Defaults to
	// :ref:`AUTO_PROTOCOL<envoy_api_enum_value_extensions.filters.network.thrift_proxy.v3.ProtocolType.AUTO_PROTOCOL>`.
	Protocol ProtocolType `protobuf:"varint,3,opt,name=protocol,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType" json:"protocol,omitempty"`
	// The human readable prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The route table for the connection manager is static and is specified in this property.
	RouteConfig *RouteConfiguration `protobuf:"bytes,4,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	// A list of individual Thrift filters that make up the filter chain for requests made to the
	// Thrift proxy. Order matters as the filters are processed sequentially. For backwards
	// compatibility, if no thrift_filters are specified, a default Thrift router filter
	// (`envoy.filters.thrift.router`) is used.
	ThriftFilters        []*ThriftFilter `protobuf:"bytes,5,rep,name=thrift_filters,json=thriftFilters,proto3" json:"thrift_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ThriftProxy) Reset()         { *m = ThriftProxy{} }
func (m *ThriftProxy) String() string { return proto.CompactTextString(m) }
func (*ThriftProxy) ProtoMessage()    {}
func (*ThriftProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5befaf4ce0c1223, []int{0}
}
func (m *ThriftProxy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ThriftProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ThriftProxy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ThriftProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProxy.Merge(m, src)
}
func (m *ThriftProxy) XXX_Size() int {
	return m.Size()
}
func (m *ThriftProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProxy proto.InternalMessageInfo

func (m *ThriftProxy) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProxy) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func (m *ThriftProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ThriftProxy) GetRouteConfig() *RouteConfiguration {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

func (m *ThriftProxy) GetThriftFilters() []*ThriftFilter {
	if m != nil {
		return m.ThriftFilters
	}
	return nil
}

// ThriftFilter configures a Thrift filter.
type ThriftFilter struct {
	// The name of the filter to instantiate. The name must match a supported
	// filter. The built-in filters are:
	//
	// [#comment:TODO(zuercher): Auto generate the following list]
	// * :ref:`envoy.filters.thrift.router <config_thrift_filters_router>`
	// * :ref:`envoy.filters.thrift.rate_limit <config_thrift_filters_rate_limit>`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being instantiated. See the supported
	// filters for further documentation.
	//
	// Types that are valid to be assigned to ConfigType:
	//	*ThriftFilter_TypedConfig
	ConfigType           isThriftFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ThriftFilter) Reset()         { *m = ThriftFilter{} }
func (m *ThriftFilter) String() string { return proto.CompactTextString(m) }
func (*ThriftFilter) ProtoMessage()    {}
func (*ThriftFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5befaf4ce0c1223, []int{1}
}
func (m *ThriftFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ThriftFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ThriftFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ThriftFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftFilter.Merge(m, src)
}
func (m *ThriftFilter) XXX_Size() int {
	return m.Size()
}
func (m *ThriftFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftFilter proto.InternalMessageInfo

type isThriftFilter_ConfigType interface {
	isThriftFilter_ConfigType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ThriftFilter_TypedConfig struct {
	TypedConfig *types.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof" json:"typed_config,omitempty"`
}

func (*ThriftFilter_TypedConfig) isThriftFilter_ConfigType() {}

func (m *ThriftFilter) GetConfigType() isThriftFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ThriftFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ThriftFilter) GetTypedConfig() *types.Any {
	if x, ok := m.GetConfigType().(*ThriftFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ThriftFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ThriftFilter_TypedConfig)(nil),
	}
}

// ThriftProtocolOptions specifies Thrift upstream protocol options. This object is used in
// in
// :ref:`typed_extension_protocol_options<envoy_api_field_config.cluster.v3.Cluster.typed_extension_protocol_options>`,
// keyed by the name `envoy.filters.network.thrift_proxy`.
type ThriftProtocolOptions struct {
	// Supplies the type of transport that the Thrift proxy should use for upstream connections.
	// Selecting
	// :ref:`AUTO_TRANSPORT<envoy_api_enum_value_extensions.filters.network.thrift_proxy.v3.TransportType.AUTO_TRANSPORT>`,
	// which is the default, causes the proxy to use the same transport as the downstream connection.
	Transport TransportType `protobuf:"varint,1,opt,name=transport,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.TransportType" json:"transport,omitempty"`
	// Supplies the type of protocol that the Thrift proxy should use for upstream connections.
	// Selecting
	// :ref:`AUTO_PROTOCOL<envoy_api_enum_value_extensions.filters.network.thrift_proxy.v3.ProtocolType.AUTO_PROTOCOL>`,
	// which is the default, causes the proxy to use the same protocol as the downstream connection.
	Protocol             ProtocolType `protobuf:"varint,2,opt,name=protocol,proto3,enum=envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ThriftProtocolOptions) Reset()         { *m = ThriftProtocolOptions{} }
func (m *ThriftProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*ThriftProtocolOptions) ProtoMessage()    {}
func (*ThriftProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5befaf4ce0c1223, []int{2}
}
func (m *ThriftProtocolOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ThriftProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ThriftProtocolOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ThriftProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProtocolOptions.Merge(m, src)
}
func (m *ThriftProtocolOptions) XXX_Size() int {
	return m.Size()
}
func (m *ThriftProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProtocolOptions proto.InternalMessageInfo

func (m *ThriftProtocolOptions) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProtocolOptions) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func init() {
	proto.RegisterEnum("envoy.extensions.filters.network.thrift_proxy.v3.TransportType", TransportType_name, TransportType_value)
	proto.RegisterEnum("envoy.extensions.filters.network.thrift_proxy.v3.ProtocolType", ProtocolType_name, ProtocolType_value)
	proto.RegisterType((*ThriftProxy)(nil), "envoy.extensions.filters.network.thrift_proxy.v3.ThriftProxy")
	proto.RegisterType((*ThriftFilter)(nil), "envoy.extensions.filters.network.thrift_proxy.v3.ThriftFilter")
	proto.RegisterType((*ThriftProtocolOptions)(nil), "envoy.extensions.filters.network.thrift_proxy.v3.ThriftProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/thrift_proxy/v3/thrift_proxy.proto", fileDescriptor_a5befaf4ce0c1223)
}

var fileDescriptor_a5befaf4ce0c1223 = []byte{
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xbb, 0x4e, 0xda, 0xa6, 0xe3, 0x24, 0x32, 0x2b, 0x10, 0xa1, 0xa0, 0x28, 0xf4, 0x14,
	0xf5, 0x60, 0x43, 0x7b, 0xa2, 0x82, 0x16, 0x3b, 0x49, 0x69, 0x50, 0x9b, 0x58, 0x8b, 0xab, 0xc2,
	0x29, 0xb8, 0xcd, 0x26, 0xb5, 0x08, 0x5e, 0x6b, 0xbd, 0x09, 0xcd, 0x95, 0x13, 0xcf, 0x80, 0x78,
	0x0e, 0xde, 0x00, 0x09, 0x71, 0xe2, 0x11, 0x50, 0x1f, 0xa3, 0x27, 0xe4, 0x5d, 0xa7, 0x4d, 0x2a,
	0x71, 0x48, 0x11, 0xdc, 0x3c, 0x3b, 0xbb, 0xdf, 0xfc, 0xf3, 0xcf, 0x18, 0x6a, 0x34, 0x1c, 0xb1,
	0xb1, 0x45, 0xcf, 0x04, 0x0d, 0xe3, 0x80, 0x85, 0xb1, 0xd5, 0x0b, 0x06, 0x82, 0xf2, 0xd8, 0x0a,
	0xa9, 0xf8, 0xc0, 0xf8, 0x3b, 0x4b, 0x9c, 0xf2, 0xa0, 0x27, 0x3a, 0x11, 0x67, 0x67, 0x63, 0x6b,
	0xb4, 0x39, 0x13, 0x9b, 0x11, 0x67, 0x82, 0xe1, 0x47, 0x12, 0x62, 0x5e, 0x41, 0xcc, 0x14, 0x62,
	0xa6, 0x10, 0x73, 0xe6, 0xd1, 0x68, 0x73, 0xf5, 0xe9, 0xdc, 0x65, 0x39, 0x1b, 0x0a, 0xaa, 0xea,
	0xad, 0xde, 0xeb, 0x33, 0xd6, 0x1f, 0x50, 0x4b, 0x46, 0xc7, 0xc3, 0x9e, 0xe5, 0x87, 0xa9, 0x94,
	0xd5, 0x07, 0xd7, 0x53, 0xb1, 0xe0, 0xc3, 0x13, 0x91, 0x66, 0x1f, 0x0e, 0xbb, 0x91, 0x6f, 0xf9,
	0x61, 0xc8, 0x84, 0x2f, 0x64, 0xd9, 0x11, 0xe5, 0x49, 0xfd, 0x20, 0xec, 0xa7, 0x57, 0xee, 0x8e,
	0xfc, 0x41, 0xd0, 0xf5, 0x05, 0xb5, 0x26, 0x1f, 0x2a, 0xb1, 0xf6, 0x25, 0x0b, 0xba, 0x27, 0x45,
	0xb9, 0x89, 0x26, 0xdc, 0x87, 0x15, 0xc1, 0xfd, 0x30, 0x8e, 0x18, 0x17, 0x25, 0xad, 0x82, 0xaa,
	0xc5, 0x8d, 0x1d, 0x73, 0x5e, 0x23, 0x4c, 0x6f, 0x82, 0xf0, 0xc6, 0x11, 0x75, 0x72, 0x17, 0xce,
	0xe2, 0x47, 0xa4, 0x19, 0x88, 0x5c, 0xb1, 0x71, 0x17, 0x72, 0x52, 0xc1, 0x09, 0x1b, 0x94, 0x32,
	0xb2, 0xce, 0xf6, 0xfc, 0x75, 0xdc, 0x94, 0x70, 0xad, 0xcc, 0x25, 0x19, 0x57, 0x41, 0x8f, 0x85,
	0x9f, 0xbc, 0xa0, 0xbd, 0xe0, 0xac, 0x84, 0x2a, 0xa8, 0xba, 0xe2, 0x2c, 0x5f, 0x38, 0x59, 0xae,
	0x55, 0x10, 0x81, 0x24, 0xe7, 0xca, 0x14, 0xee, 0x43, 0x5e, 0x0e, 0xa3, 0x73, 0xc2, 0xc2, 0x5e,
	0xd0, 0x2f, 0x65, 0x2b, 0xa8, 0xaa, 0x6f, 0xd4, 0xe7, 0xd7, 0x44, 0x12, 0x4a, 0x4d, 0x42, 0x86,
	0x5c, 0xce, 0x83, 0xe8, 0xfc, 0xea, 0x0c, 0x53, 0x28, 0xa6, 0x4f, 0x52, 0x52, 0x69, 0xb1, 0x92,
	0xa9, 0xea, 0x37, 0x69, 0x5f, 0x0d, 0x6e, 0x57, 0xde, 0x22, 0x05, 0x31, 0x15, 0xc5, 0x5b, 0xf5,
	0xcf, 0xdf, 0x3e, 0x95, 0x77, 0xe0, 0x99, 0x82, 0xaa, 0xa6, 0x52, 0xe0, 0x1f, 0x78, 0x1b, 0xfe,
	0x20, 0x3a, 0xf5, 0x1f, 0x9b, 0x53, 0xeb, 0xb0, 0xf6, 0x03, 0x41, 0x7e, 0xba, 0x0a, 0xbe, 0x0f,
	0xd9, 0xd0, 0x7f, 0x4f, 0xaf, 0x3b, 0x29, 0x0f, 0xf1, 0x13, 0xc8, 0x8b, 0x71, 0x44, 0xbb, 0x13,
	0x0f, 0x33, 0xd2, 0xc3, 0xdb, 0xa6, 0xda, 0x5e, 0x73, 0xb2, 0xbd, 0xa6, 0x1d, 0x8e, 0xf7, 0x16,
	0x88, 0x2e, 0xef, 0x2a, 0x57, 0xb6, 0x1a, 0x89, 0xdc, 0xe7, 0xb0, 0x7d, 0x53, 0xb9, 0x4a, 0x9e,
	0x53, 0x00, 0x5d, 0xbd, 0xed, 0x24, 0xf0, 0x97, 0xd9, 0x9c, 0x66, 0x64, 0xc8, 0x92, 0x3a, 0x5a,
	0xfb, 0xaa, 0xc1, 0x9d, 0xcb, 0xe6, 0xe4, 0x7e, 0xb4, 0x23, 0xf9, 0xbf, 0xcc, 0x6e, 0x3d, 0xfa,
	0x4f, 0x5b, 0xaf, 0xfd, 0xab, 0xad, 0xdf, 0x6a, 0x25, 0x66, 0x36, 0xe1, 0xc5, 0x5f, 0xcc, 0x7e,
	0xda, 0x9e, 0xf5, 0x26, 0x14, 0x66, 0x7a, 0xc3, 0x18, 0x8a, 0xf6, 0xa1, 0xd7, 0xee, 0x78, 0xc4,
	0x6e, 0xbd, 0x72, 0xdb, 0xc4, 0x33, 0x16, 0x30, 0xc0, 0xd2, 0x2e, 0xb1, 0x0f, 0x1a, 0x75, 0x03,
	0xe1, 0x3c, 0xe4, 0x0e, 0x5b, 0x69, 0xa4, 0x25, 0x99, 0xbd, 0x86, 0x5d, 0x6f, 0x10, 0x23, 0xb3,
	0x7e, 0x04, 0xf9, 0x69, 0xf9, 0xf8, 0x16, 0x14, 0x24, 0xc9, 0x25, 0x6d, 0xaf, 0x5d, 0x6b, 0xef,
	0x2b, 0x90, 0xd3, 0x6c, 0xd9, 0xe4, 0x8d, 0x81, 0x70, 0x11, 0x60, 0xdf, 0x7e, 0xdd, 0x49, 0x63,
	0x0d, 0xeb, 0xb0, 0x5c, 0x6b, 0x1f, 0xb8, 0x76, 0xcd, 0x33, 0x32, 0x49, 0xe0, 0x1d, 0x35, 0x3d,
	0xaf, 0x41, 0x8c, 0xac, 0xf3, 0xf6, 0xfb, 0x79, 0x19, 0xfd, 0x3c, 0x2f, 0xa3, 0x5f, 0xe7, 0x65,
	0x04, 0xdb, 0x01, 0x53, 0xbe, 0xaa, 0xfe, 0xe6, 0xb5, 0xd8, 0x31, 0xa6, 0x7e, 0x02, 0xa9, 0xd7,
	0x45, 0xc7, 0x4b, 0xd2, 0xdf, 0xcd, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x58, 0x84, 0x11, 0x98,
	0x5d, 0x06, 0x00, 0x00,
}

func (m *ThriftProxy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ThriftProxy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ThriftProxy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ThriftFilters) > 0 {
		for iNdEx := len(m.ThriftFilters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ThriftFilters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintThriftProxy(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.RouteConfig != nil {
		{
			size, err := m.RouteConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintThriftProxy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Protocol != 0 {
		i = encodeVarintThriftProxy(dAtA, i, uint64(m.Protocol))
		i--
		dAtA[i] = 0x18
	}
	if m.Transport != 0 {
		i = encodeVarintThriftProxy(dAtA, i, uint64(m.Transport))
		i--
		dAtA[i] = 0x10
	}
	if len(m.StatPrefix) > 0 {
		i -= len(m.StatPrefix)
		copy(dAtA[i:], m.StatPrefix)
		i = encodeVarintThriftProxy(dAtA, i, uint64(len(m.StatPrefix)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ThriftFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ThriftFilter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ThriftFilter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ConfigType != nil {
		{
			size := m.ConfigType.Size()
			i -= size
			if _, err := m.ConfigType.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintThriftProxy(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ThriftFilter_TypedConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ThriftFilter_TypedConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.TypedConfig != nil {
		{
			size, err := m.TypedConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintThriftProxy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *ThriftProtocolOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ThriftProtocolOptions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ThriftProtocolOptions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Protocol != 0 {
		i = encodeVarintThriftProxy(dAtA, i, uint64(m.Protocol))
		i--
		dAtA[i] = 0x10
	}
	if m.Transport != 0 {
		i = encodeVarintThriftProxy(dAtA, i, uint64(m.Transport))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintThriftProxy(dAtA []byte, offset int, v uint64) int {
	offset -= sovThriftProxy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ThriftProxy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovThriftProxy(uint64(l))
	}
	if m.Transport != 0 {
		n += 1 + sovThriftProxy(uint64(m.Transport))
	}
	if m.Protocol != 0 {
		n += 1 + sovThriftProxy(uint64(m.Protocol))
	}
	if m.RouteConfig != nil {
		l = m.RouteConfig.Size()
		n += 1 + l + sovThriftProxy(uint64(l))
	}
	if len(m.ThriftFilters) > 0 {
		for _, e := range m.ThriftFilters {
			l = e.Size()
			n += 1 + l + sovThriftProxy(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ThriftFilter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovThriftProxy(uint64(l))
	}
	if m.ConfigType != nil {
		n += m.ConfigType.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ThriftFilter_TypedConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TypedConfig != nil {
		l = m.TypedConfig.Size()
		n += 1 + l + sovThriftProxy(uint64(l))
	}
	return n
}
func (m *ThriftProtocolOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Transport != 0 {
		n += 1 + sovThriftProxy(uint64(m.Transport))
	}
	if m.Protocol != 0 {
		n += 1 + sovThriftProxy(uint64(m.Protocol))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovThriftProxy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozThriftProxy(x uint64) (n int) {
	return sovThriftProxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ThriftProxy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowThriftProxy
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
			return fmt.Errorf("proto: ThriftProxy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ThriftProxy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowThriftProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthThriftProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthThriftProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transport", wireType)
			}
			m.Transport = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowThriftProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Transport |= TransportType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			m.Protocol = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowThriftProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Protocol |= ProtocolType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RouteConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowThriftProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthThriftProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthThriftProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RouteConfig == nil {
				m.RouteConfig = &RouteConfiguration{}
			}
			if err := m.RouteConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThriftFilters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowThriftProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthThriftProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthThriftProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ThriftFilters = append(m.ThriftFilters, &ThriftFilter{})
			if err := m.ThriftFilters[len(m.ThriftFilters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipThriftProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthThriftProxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthThriftProxy
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
func (m *ThriftFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowThriftProxy
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
			return fmt.Errorf("proto: ThriftFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ThriftFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowThriftProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthThriftProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthThriftProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypedConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowThriftProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthThriftProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthThriftProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.Any{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ConfigType = &ThriftFilter_TypedConfig{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipThriftProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthThriftProxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthThriftProxy
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
func (m *ThriftProtocolOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowThriftProxy
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
			return fmt.Errorf("proto: ThriftProtocolOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ThriftProtocolOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transport", wireType)
			}
			m.Transport = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowThriftProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Transport |= TransportType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			m.Protocol = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowThriftProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Protocol |= ProtocolType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipThriftProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthThriftProxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthThriftProxy
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
func skipThriftProxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowThriftProxy
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
					return 0, ErrIntOverflowThriftProxy
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
					return 0, ErrIntOverflowThriftProxy
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
				return 0, ErrInvalidLengthThriftProxy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupThriftProxy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthThriftProxy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthThriftProxy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowThriftProxy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupThriftProxy = fmt.Errorf("proto: unexpected end of group")
)
