// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sizeunderscore.proto

package sizeunderscore

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type SizeMessage struct {
	Size_                *int64   `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Equal_               *bool    `protobuf:"varint,2,opt,name=Equal" json:"Equal,omitempty"`
	String_              *string  `protobuf:"bytes,3,opt,name=String" json:"String,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SizeMessage) Reset()         { *m = SizeMessage{} }
func (m *SizeMessage) String() string { return proto.CompactTextString(m) }
func (*SizeMessage) ProtoMessage()    {}
func (*SizeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ba72c70f4572a8b, []int{0}
}
func (m *SizeMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SizeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SizeMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SizeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SizeMessage.Merge(m, src)
}
func (m *SizeMessage) XXX_Size() int {
	return m.Size()
}
func (m *SizeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SizeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SizeMessage proto.InternalMessageInfo

func (m *SizeMessage) GetSize_() int64 {
	if m != nil && m.Size_ != nil {
		return *m.Size_
	}
	return 0
}

func (m *SizeMessage) GetEqual_() bool {
	if m != nil && m.Equal_ != nil {
		return *m.Equal_
	}
	return false
}

func (m *SizeMessage) GetString_() string {
	if m != nil && m.String_ != nil {
		return *m.String_
	}
	return ""
}

func init() {
	proto.RegisterType((*SizeMessage)(nil), "sizeunderscore.SizeMessage")
}

func init() { proto.RegisterFile("sizeunderscore.proto", fileDescriptor_4ba72c70f4572a8b) }

var fileDescriptor_4ba72c70f4572a8b = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xce, 0xac, 0x4a,
	0x2d, 0xcd, 0x4b, 0x49, 0x2d, 0x2a, 0x4e, 0xce, 0x2f, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x43, 0x15, 0x95, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x2b, 0x4b, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30,
	0x0b, 0xa2, 0x5d, 0xc9, 0x9f, 0x8b, 0x3b, 0x38, 0xb3, 0x2a, 0xd5, 0x37, 0xb5, 0xb8, 0x38, 0x31,
	0x3d, 0x55, 0x48, 0x88, 0x8b, 0x05, 0x64, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x98,
	0x2d, 0x24, 0xc2, 0xc5, 0xea, 0x5a, 0x58, 0x9a, 0x98, 0x23, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x11,
	0x04, 0xe1, 0x08, 0x89, 0x71, 0xb1, 0x05, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x4b, 0x30, 0x2b, 0x30,
	0x6a, 0x70, 0x06, 0x41, 0x79, 0x4e, 0x12, 0x3f, 0x1e, 0xca, 0x31, 0xae, 0x78, 0x24, 0xc7, 0xb8,
	0xe3, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7,
	0x08, 0x08, 0x00, 0x00, 0xff, 0xff, 0x37, 0x1c, 0x48, 0xa4, 0xc0, 0x00, 0x00, 0x00,
}

func (this *SizeMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SizeMessage)
	if !ok {
		that2, ok := that.(SizeMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Size_ != nil && that1.Size_ != nil {
		if *this.Size_ != *that1.Size_ {
			return false
		}
	} else if this.Size_ != nil {
		return false
	} else if that1.Size_ != nil {
		return false
	}
	if this.Equal_ != nil && that1.Equal_ != nil {
		if *this.Equal_ != *that1.Equal_ {
			return false
		}
	} else if this.Equal_ != nil {
		return false
	} else if that1.Equal_ != nil {
		return false
	}
	if this.String_ != nil && that1.String_ != nil {
		if *this.String_ != *that1.String_ {
			return false
		}
	} else if this.String_ != nil {
		return false
	} else if that1.String_ != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *SizeMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SizeMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SizeMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.String_ != nil {
		i -= len(*m.String_)
		copy(dAtA[i:], *m.String_)
		i = encodeVarintSizeunderscore(dAtA, i, uint64(len(*m.String_)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Equal_ != nil {
		i--
		if *m.Equal_ {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Size_ != nil {
		i = encodeVarintSizeunderscore(dAtA, i, uint64(*m.Size_))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSizeunderscore(dAtA []byte, offset int, v uint64) int {
	offset -= sovSizeunderscore(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedSizeMessage(r randySizeunderscore, easy bool) *SizeMessage {
	this := &SizeMessage{}
	if r.Intn(5) != 0 {
		v1 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v1 *= -1
		}
		this.Size_ = &v1
	}
	if r.Intn(5) != 0 {
		v2 := bool(bool(r.Intn(2) == 0))
		this.Equal_ = &v2
	}
	if r.Intn(5) != 0 {
		v3 := string(randStringSizeunderscore(r))
		this.String_ = &v3
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedSizeunderscore(r, 4)
	}
	return this
}

type randySizeunderscore interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSizeunderscore(r randySizeunderscore) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringSizeunderscore(r randySizeunderscore) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneSizeunderscore(r)
	}
	return string(tmps)
}
func randUnrecognizedSizeunderscore(r randySizeunderscore, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldSizeunderscore(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldSizeunderscore(dAtA []byte, r randySizeunderscore, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateSizeunderscore(dAtA, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		dAtA = encodeVarintPopulateSizeunderscore(dAtA, uint64(v5))
	case 1:
		dAtA = encodeVarintPopulateSizeunderscore(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateSizeunderscore(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateSizeunderscore(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateSizeunderscore(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateSizeunderscore(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *SizeMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Size_ != nil {
		n += 1 + sovSizeunderscore(uint64(*m.Size_))
	}
	if m.Equal_ != nil {
		n += 2
	}
	if m.String_ != nil {
		l = len(*m.String_)
		n += 1 + l + sovSizeunderscore(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSizeunderscore(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSizeunderscore(x uint64) (n int) {
	return sovSizeunderscore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SizeMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSizeunderscore
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
			return fmt.Errorf("proto: SizeMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SizeMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSizeunderscore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Size_ = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Equal_", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSizeunderscore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Equal_ = &b
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field String_", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSizeunderscore
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
				return ErrInvalidLengthSizeunderscore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSizeunderscore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.String_ = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSizeunderscore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSizeunderscore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSizeunderscore
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
func skipSizeunderscore(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSizeunderscore
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
					return 0, ErrIntOverflowSizeunderscore
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSizeunderscore
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
				return 0, ErrInvalidLengthSizeunderscore
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSizeunderscore
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSizeunderscore
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSizeunderscore(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSizeunderscore
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSizeunderscore = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSizeunderscore   = fmt.Errorf("proto: integer overflow")
)
