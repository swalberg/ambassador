// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/access_loggers/file/v3/file.proto

package envoy_extensions_access_loggers_file_v3

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

// Custom configuration for an :ref:`AccessLog <envoy_api_msg_config.accesslog.v3.AccessLog>`
// that writes log entries directly to a file. Configures the built-in *envoy.file_access_log*
// AccessLog.
type FileAccessLog struct {
	// A path to a local file to which to write the access log entries.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Types that are valid to be assigned to AccessLogFormat:
	//	*FileAccessLog_Format
	//	*FileAccessLog_JsonFormat
	//	*FileAccessLog_TypedJsonFormat
	AccessLogFormat      isFileAccessLog_AccessLogFormat `protobuf_oneof:"access_log_format"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FileAccessLog) Reset()         { *m = FileAccessLog{} }
func (m *FileAccessLog) String() string { return proto.CompactTextString(m) }
func (*FileAccessLog) ProtoMessage()    {}
func (*FileAccessLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_b950f2ced354a07e, []int{0}
}
func (m *FileAccessLog) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileAccessLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileAccessLog.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileAccessLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileAccessLog.Merge(m, src)
}
func (m *FileAccessLog) XXX_Size() int {
	return m.Size()
}
func (m *FileAccessLog) XXX_DiscardUnknown() {
	xxx_messageInfo_FileAccessLog.DiscardUnknown(m)
}

var xxx_messageInfo_FileAccessLog proto.InternalMessageInfo

type isFileAccessLog_AccessLogFormat interface {
	isFileAccessLog_AccessLogFormat()
	MarshalTo([]byte) (int, error)
	Size() int
}

type FileAccessLog_Format struct {
	Format string `protobuf:"bytes,2,opt,name=format,proto3,oneof" json:"format,omitempty"`
}
type FileAccessLog_JsonFormat struct {
	JsonFormat *types.Struct `protobuf:"bytes,3,opt,name=json_format,json=jsonFormat,proto3,oneof" json:"json_format,omitempty"`
}
type FileAccessLog_TypedJsonFormat struct {
	TypedJsonFormat *types.Struct `protobuf:"bytes,4,opt,name=typed_json_format,json=typedJsonFormat,proto3,oneof" json:"typed_json_format,omitempty"`
}

func (*FileAccessLog_Format) isFileAccessLog_AccessLogFormat()          {}
func (*FileAccessLog_JsonFormat) isFileAccessLog_AccessLogFormat()      {}
func (*FileAccessLog_TypedJsonFormat) isFileAccessLog_AccessLogFormat() {}

func (m *FileAccessLog) GetAccessLogFormat() isFileAccessLog_AccessLogFormat {
	if m != nil {
		return m.AccessLogFormat
	}
	return nil
}

func (m *FileAccessLog) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileAccessLog) GetFormat() string {
	if x, ok := m.GetAccessLogFormat().(*FileAccessLog_Format); ok {
		return x.Format
	}
	return ""
}

func (m *FileAccessLog) GetJsonFormat() *types.Struct {
	if x, ok := m.GetAccessLogFormat().(*FileAccessLog_JsonFormat); ok {
		return x.JsonFormat
	}
	return nil
}

func (m *FileAccessLog) GetTypedJsonFormat() *types.Struct {
	if x, ok := m.GetAccessLogFormat().(*FileAccessLog_TypedJsonFormat); ok {
		return x.TypedJsonFormat
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FileAccessLog) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FileAccessLog_Format)(nil),
		(*FileAccessLog_JsonFormat)(nil),
		(*FileAccessLog_TypedJsonFormat)(nil),
	}
}

func init() {
	proto.RegisterType((*FileAccessLog)(nil), "envoy.extensions.access_loggers.file.v3.FileAccessLog")
}

func init() {
	proto.RegisterFile("envoy/extensions/access_loggers/file/v3/file.proto", fileDescriptor_b950f2ced354a07e)
}

var fileDescriptor_b950f2ced354a07e = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4e, 0x32, 0x31,
	0x14, 0xc5, 0x29, 0x1f, 0xe1, 0x0b, 0x25, 0xc6, 0x30, 0x2e, 0x20, 0x68, 0x26, 0xe8, 0x46, 0xdc,
	0xb4, 0x09, 0xc4, 0x0d, 0x3b, 0x27, 0x91, 0x10, 0xe3, 0x82, 0xe0, 0x03, 0x90, 0x32, 0x74, 0x6a,
	0xcd, 0xd8, 0x3b, 0x69, 0xcb, 0x04, 0xde, 0xc0, 0x95, 0x0f, 0xe0, 0xfb, 0x98, 0xb8, 0xf4, 0x11,
	0xcc, 0x3c, 0x86, 0x2b, 0x33, 0x9d, 0x41, 0x64, 0xa5, 0xab, 0xfe, 0xb9, 0xf7, 0xfc, 0xce, 0xe9,
	0x2d, 0x1e, 0x70, 0x95, 0xc2, 0x86, 0xf2, 0xb5, 0xe5, 0xca, 0x48, 0x50, 0x86, 0xb2, 0x30, 0xe4,
	0xc6, 0xcc, 0x63, 0x10, 0x82, 0x6b, 0x43, 0x23, 0x19, 0x73, 0x9a, 0x0e, 0xdd, 0x4a, 0x12, 0x0d,
	0x16, 0xbc, 0x73, 0xa7, 0x21, 0x3b, 0x0d, 0xd9, 0xd7, 0x10, 0xd7, 0x9b, 0x0e, 0xbb, 0x27, 0x02,
	0x40, 0xc4, 0x9c, 0x3a, 0xd9, 0x62, 0x15, 0x51, 0x63, 0xf5, 0x2a, 0xb4, 0x05, 0xa6, 0x7b, 0xba,
	0x5a, 0x26, 0x8c, 0x32, 0xa5, 0xc0, 0x32, 0xeb, 0xac, 0x53, 0xae, 0x73, 0x9e, 0x54, 0xa2, 0x6c,
	0x69, 0xa7, 0x2c, 0x96, 0x4b, 0x66, 0x39, 0xdd, 0x6e, 0x8a, 0xc2, 0xd9, 0x73, 0x15, 0x1f, 0x8c,
	0x65, 0xcc, 0xaf, 0x9c, 0xf1, 0x2d, 0x08, 0xef, 0x18, 0xd7, 0x12, 0x66, 0xef, 0x3b, 0xa8, 0x87,
	0xfa, 0x8d, 0xe0, 0xff, 0x67, 0x50, 0xd3, 0xd5, 0x1e, 0x9a, 0xb9, 0x4b, 0xaf, 0x83, 0xeb, 0x11,
	0xe8, 0x47, 0x66, 0x3b, 0xd5, 0xbc, 0x3c, 0xa9, 0xcc, 0xca, 0xb3, 0x37, 0xc2, 0xcd, 0x07, 0x03,
	0x6a, 0x5e, 0x96, 0xff, 0xf5, 0x50, 0xbf, 0x39, 0x68, 0x93, 0x22, 0x38, 0xd9, 0x06, 0x27, 0x77,
	0x2e, 0xf8, 0xa4, 0x32, 0xc3, 0x79, 0xf7, 0xb8, 0xd0, 0x5e, 0xe3, 0x96, 0xdd, 0x24, 0x7c, 0x39,
	0xff, 0x49, 0xa8, 0xfd, 0x46, 0x38, 0x74, 0x9a, 0x9b, 0x6f, 0xcc, 0x88, 0xbc, 0xbc, 0x3e, 0xf9,
	0x17, 0xb8, 0x9c, 0x6a, 0x08, 0x2a, 0x92, 0xa2, 0x9c, 0x68, 0x0c, 0x82, 0xa4, 0x03, 0xb2, 0xf7,
	0xd2, 0xe0, 0x08, 0xb7, 0x76, 0xf3, 0x2e, 0x6d, 0x83, 0xe9, 0x5b, 0xe6, 0xa3, 0xf7, 0xcc, 0x47,
	0x1f, 0x99, 0x8f, 0xf0, 0xa5, 0x04, 0xe2, 0x70, 0x89, 0x86, 0xf5, 0x86, 0xfc, 0xf1, 0xbf, 0x82,
	0x46, 0x6e, 0x34, 0xcd, 0x13, 0x4f, 0xd1, 0xa2, 0xee, 0xa2, 0x0f, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xcd, 0x24, 0xe0, 0x73, 0x22, 0x02, 0x00, 0x00,
}

func (m *FileAccessLog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileAccessLog) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileAccessLog) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.AccessLogFormat != nil {
		{
			size := m.AccessLogFormat.Size()
			i -= size
			if _, err := m.AccessLogFormat.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintFile(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FileAccessLog_Format) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileAccessLog_Format) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Format)
	copy(dAtA[i:], m.Format)
	i = encodeVarintFile(dAtA, i, uint64(len(m.Format)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func (m *FileAccessLog_JsonFormat) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileAccessLog_JsonFormat) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.JsonFormat != nil {
		{
			size, err := m.JsonFormat.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFile(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *FileAccessLog_TypedJsonFormat) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileAccessLog_TypedJsonFormat) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.TypedJsonFormat != nil {
		{
			size, err := m.TypedJsonFormat.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFile(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func encodeVarintFile(dAtA []byte, offset int, v uint64) int {
	offset -= sovFile(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FileAccessLog) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovFile(uint64(l))
	}
	if m.AccessLogFormat != nil {
		n += m.AccessLogFormat.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FileAccessLog_Format) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Format)
	n += 1 + l + sovFile(uint64(l))
	return n
}
func (m *FileAccessLog_JsonFormat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.JsonFormat != nil {
		l = m.JsonFormat.Size()
		n += 1 + l + sovFile(uint64(l))
	}
	return n
}
func (m *FileAccessLog_TypedJsonFormat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TypedJsonFormat != nil {
		l = m.TypedJsonFormat.Size()
		n += 1 + l + sovFile(uint64(l))
	}
	return n
}

func sovFile(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFile(x uint64) (n int) {
	return sovFile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FileAccessLog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFile
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
			return fmt.Errorf("proto: FileAccessLog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileAccessLog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
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
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
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
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessLogFormat = &FileAccessLog_Format{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JsonFormat", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
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
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.Struct{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.AccessLogFormat = &FileAccessLog_JsonFormat{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypedJsonFormat", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
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
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.Struct{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.AccessLogFormat = &FileAccessLog_TypedJsonFormat{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFile
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFile
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
func skipFile(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFile
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
					return 0, ErrIntOverflowFile
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
					return 0, ErrIntOverflowFile
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
				return 0, ErrInvalidLengthFile
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFile
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFile
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFile        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFile          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFile = fmt.Errorf("proto: unexpected end of group")
)