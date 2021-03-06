// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/network/rbac/v3/rbac.proto

package envoy_extensions_filters_network_rbac_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/config/rbac/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
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

type RBAC_EnforcementType int32

const (
	// Apply RBAC policies when the first byte of data arrives on the connection.
	RBAC_ONE_TIME_ON_FIRST_BYTE RBAC_EnforcementType = 0
	// Continuously apply RBAC policies as data arrives. Use this mode when
	// using RBAC with message oriented protocols such as Mongo, MySQL, Kafka,
	// etc. when the protocol decoders emit dynamic metadata such as the
	// resources being accessed and the operations on the resources.
	RBAC_CONTINUOUS RBAC_EnforcementType = 1
)

var RBAC_EnforcementType_name = map[int32]string{
	0: "ONE_TIME_ON_FIRST_BYTE",
	1: "CONTINUOUS",
}

var RBAC_EnforcementType_value = map[string]int32{
	"ONE_TIME_ON_FIRST_BYTE": 0,
	"CONTINUOUS":             1,
}

func (x RBAC_EnforcementType) String() string {
	return proto.EnumName(RBAC_EnforcementType_name, int32(x))
}

func (RBAC_EnforcementType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_287a9975e9b41589, []int{0, 0}
}

// RBAC network filter config.
//
// Header should not be used in rules/shadow_rules in RBAC network filter as
// this information is only available in :ref:`RBAC http filter <config_http_filters_rbac>`.
type RBAC struct {
	// Specify the RBAC rules to be applied globally.
	// If absent, no enforcing RBAC policy will be applied.
	Rules *v3.RBAC `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	// Shadow rules are not enforced by the filter but will emit stats and logs
	// and can be used for rule testing.
	// If absent, no shadow RBAC policy will be applied.
	ShadowRules *v3.RBAC `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	// The prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,3,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// RBAC enforcement strategy. By default RBAC will be enforced only once
	// when the first byte of data arrives from the downstream. When used in
	// conjunction with filters that emit dynamic metadata after decoding
	// every payload (e.g., Mongo, MySQL, Kafka) set the enforcement type to
	// CONTINUOUS to enforce RBAC policies on every message boundary.
	EnforcementType      RBAC_EnforcementType `protobuf:"varint,4,opt,name=enforcement_type,json=enforcementType,proto3,enum=envoy.extensions.filters.network.rbac.v3.RBAC_EnforcementType" json:"enforcement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_287a9975e9b41589, []int{0}
}
func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return m.Size()
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetRules() *v3.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RBAC) GetShadowRules() *v3.RBAC {
	if m != nil {
		return m.ShadowRules
	}
	return nil
}

func (m *RBAC) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RBAC) GetEnforcementType() RBAC_EnforcementType {
	if m != nil {
		return m.EnforcementType
	}
	return RBAC_ONE_TIME_ON_FIRST_BYTE
}

func init() {
	proto.RegisterEnum("envoy.extensions.filters.network.rbac.v3.RBAC_EnforcementType", RBAC_EnforcementType_name, RBAC_EnforcementType_value)
	proto.RegisterType((*RBAC)(nil), "envoy.extensions.filters.network.rbac.v3.RBAC")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/rbac/v3/rbac.proto", fileDescriptor_287a9975e9b41589)
}

var fileDescriptor_287a9975e9b41589 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x9d, 0x5a, 0x95, 0x3b, 0x95, 0x7b, 0x4b, 0x16, 0x5a, 0xb2, 0x88, 0xf1, 0xae, 0x82,
	0x8b, 0x19, 0x69, 0xc0, 0x85, 0x70, 0x05, 0x53, 0x22, 0x74, 0x61, 0x52, 0xd3, 0x74, 0xe1, 0x2a,
	0x4c, 0xd3, 0x49, 0x1d, 0xac, 0x33, 0x61, 0x32, 0x4d, 0x9b, 0x37, 0xf0, 0x19, 0x7c, 0x1f, 0xc1,
	0xa5, 0x8f, 0x20, 0x7d, 0x0c, 0x71, 0x21, 0x99, 0x69, 0x29, 0x2d, 0x82, 0x5d, 0x25, 0xcc, 0xf9,
	0xbf, 0xff, 0xfc, 0xe7, 0x1c, 0xe8, 0x53, 0x5e, 0x8b, 0x06, 0xd3, 0xad, 0xa2, 0xbc, 0x62, 0x82,
	0x57, 0xb8, 0x60, 0x2b, 0x45, 0x65, 0x85, 0x39, 0x55, 0x1b, 0x21, 0x3f, 0x63, 0x39, 0x27, 0x39,
	0xae, 0x7d, 0xfd, 0x45, 0xa5, 0x14, 0x4a, 0x58, 0x9e, 0x86, 0xd0, 0x11, 0x42, 0x7b, 0x08, 0xed,
	0x21, 0xa4, 0xc5, 0xb5, 0x6f, 0x3f, 0x33, 0xf6, 0xb9, 0xe0, 0x05, 0x5b, 0xfe, 0xc3, 0xca, 0x7e,
	0xbe, 0x5e, 0x94, 0x04, 0x13, 0xce, 0x85, 0x22, 0x4a, 0xf7, 0xaf, 0xa9, 0x6c, 0x3d, 0x19, 0x5f,
	0xee, 0x25, 0x4f, 0x6b, 0xb2, 0x62, 0x0b, 0xa2, 0x28, 0x3e, 0xfc, 0x98, 0xc2, 0xed, 0x9f, 0x0e,
	0xec, 0x26, 0xc1, 0xdb, 0x91, 0xf5, 0x12, 0x3e, 0x90, 0xeb, 0x15, 0xad, 0x06, 0xc0, 0x05, 0x5e,
	0x6f, 0x68, 0x23, 0x93, 0xcf, 0x74, 0x3d, 0x64, 0x41, 0xad, 0x34, 0x31, 0x42, 0xeb, 0x0e, 0x3e,
	0xae, 0x3e, 0x91, 0x85, 0xd8, 0x64, 0x06, 0xec, 0xfc, 0x17, 0xec, 0x19, 0x7d, 0xa2, 0x71, 0x0f,
	0xf6, 0x2a, 0x45, 0x54, 0x56, 0x4a, 0x5a, 0xb0, 0xed, 0xe0, 0xbe, 0x0b, 0xbc, 0xab, 0xe0, 0xd1,
	0xef, 0xa0, 0x2b, 0x3b, 0x2e, 0x48, 0x60, 0x5b, 0x9b, 0xe8, 0x92, 0xc5, 0x60, 0x9f, 0xf2, 0x42,
	0xc8, 0x9c, 0x7e, 0xa1, 0x5c, 0x65, 0xaa, 0x29, 0xe9, 0xa0, 0xeb, 0x02, 0xef, 0x7a, 0xf8, 0x06,
	0x5d, 0xba, 0x45, 0x1d, 0x00, 0x85, 0x47, 0x9b, 0xb4, 0x29, 0x69, 0x72, 0x43, 0x4f, 0x1f, 0x6e,
	0xef, 0xe0, 0xcd, 0x99, 0xc6, 0xb2, 0xe1, 0x93, 0x38, 0x0a, 0xb3, 0x74, 0xfc, 0x3e, 0xcc, 0xe2,
	0x28, 0x7b, 0x37, 0x4e, 0xa6, 0x69, 0x16, 0x7c, 0x4c, 0xc3, 0xfe, 0x3d, 0xeb, 0x1a, 0xc2, 0x51,
	0x1c, 0xa5, 0xe3, 0x68, 0x16, 0xcf, 0xa6, 0x7d, 0xf0, 0x1a, 0x7f, 0xfb, 0xfe, 0xd5, 0x79, 0x01,
	0xbd, 0x93, 0x15, 0x98, 0x44, 0x67, 0x81, 0x86, 0x3a, 0x50, 0xf0, 0xe1, 0xc7, 0xce, 0x01, 0x3f,
	0x77, 0x0e, 0xf8, 0xb5, 0x73, 0x00, 0x7c, 0xc5, 0x84, 0x19, 0xa8, 0x94, 0x62, 0xdb, 0x5c, 0x3c,
	0x5b, 0x70, 0x95, 0xcc, 0x49, 0x3e, 0x69, 0xef, 0x39, 0x01, 0xf3, 0x87, 0xfa, 0xb0, 0xfe, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xbe, 0x06, 0x92, 0x96, 0x02, 0x00, 0x00,
}

func (m *RBAC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RBAC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RBAC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.EnforcementType != 0 {
		i = encodeVarintRbac(dAtA, i, uint64(m.EnforcementType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.StatPrefix) > 0 {
		i -= len(m.StatPrefix)
		copy(dAtA[i:], m.StatPrefix)
		i = encodeVarintRbac(dAtA, i, uint64(len(m.StatPrefix)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ShadowRules != nil {
		{
			size, err := m.ShadowRules.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRbac(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Rules != nil {
		{
			size, err := m.Rules.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRbac(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRbac(dAtA []byte, offset int, v uint64) int {
	offset -= sovRbac(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RBAC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Rules != nil {
		l = m.Rules.Size()
		n += 1 + l + sovRbac(uint64(l))
	}
	if m.ShadowRules != nil {
		l = m.ShadowRules.Size()
		n += 1 + l + sovRbac(uint64(l))
	}
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovRbac(uint64(l))
	}
	if m.EnforcementType != 0 {
		n += 1 + sovRbac(uint64(m.EnforcementType))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRbac(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRbac(x uint64) (n int) {
	return sovRbac(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RBAC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRbac
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
			return fmt.Errorf("proto: RBAC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RBAC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
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
				return ErrInvalidLengthRbac
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRbac
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Rules == nil {
				m.Rules = &v3.RBAC{}
			}
			if err := m.Rules.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShadowRules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
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
				return ErrInvalidLengthRbac
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRbac
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ShadowRules == nil {
				m.ShadowRules = &v3.RBAC{}
			}
			if err := m.ShadowRules.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
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
				return ErrInvalidLengthRbac
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRbac
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnforcementType", wireType)
			}
			m.EnforcementType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRbac
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EnforcementType |= RBAC_EnforcementType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRbac(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRbac
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRbac
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
func skipRbac(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRbac
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
					return 0, ErrIntOverflowRbac
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
					return 0, ErrIntOverflowRbac
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
				return 0, ErrInvalidLengthRbac
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRbac
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRbac
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRbac        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRbac          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRbac = fmt.Errorf("proto: unexpected end of group")
)
