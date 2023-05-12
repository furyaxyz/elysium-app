// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elysium/da/data_availability_header.proto

package da

import (
	fmt "fmt"
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

// DataAvailabilityHeader contains the row and column roots of the erasure
// coded version of the data in Block.Data.
// Therefor the original Block.Data is arranged in a
// k × k matrix, which is then "extended" to a
// 2k × 2k matrix applying multiple times Reed-Solomon encoding.
// For details see Section 5.2: https://arxiv.org/abs/1809.09044
// or the Elysium specification:
// https://github.com/elysiumorg/elysium-specs/blob/master/src/specs/data_structures.md#availabledataheader
// Note that currently we list row and column roots in separate fields
// (different from the spec).
type DataAvailabilityHeader struct {
	// RowRoot_j 	= root((M_{j,1} || M_{j,2} || ... || M_{j,2k} ))
	RowRoots [][]byte `protobuf:"bytes,1,rep,name=row_roots,json=rowRoots,proto3" json:"row_roots,omitempty"`
	// ColumnRoot_j = root((M_{1,j} || M_{2,j} || ... || M_{2k,j} ))
	ColumnRoots [][]byte `protobuf:"bytes,2,rep,name=column_roots,json=columnRoots,proto3" json:"column_roots,omitempty"`
}

func (m *DataAvailabilityHeader) Reset()         { *m = DataAvailabilityHeader{} }
func (m *DataAvailabilityHeader) String() string { return proto.CompactTextString(m) }
func (*DataAvailabilityHeader) ProtoMessage()    {}
func (*DataAvailabilityHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_93b487fd8444a5fd, []int{0}
}
func (m *DataAvailabilityHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataAvailabilityHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataAvailabilityHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataAvailabilityHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataAvailabilityHeader.Merge(m, src)
}
func (m *DataAvailabilityHeader) XXX_Size() int {
	return m.Size()
}
func (m *DataAvailabilityHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_DataAvailabilityHeader.DiscardUnknown(m)
}

var xxx_messageInfo_DataAvailabilityHeader proto.InternalMessageInfo

func (m *DataAvailabilityHeader) GetRowRoots() [][]byte {
	if m != nil {
		return m.RowRoots
	}
	return nil
}

func (m *DataAvailabilityHeader) GetColumnRoots() [][]byte {
	if m != nil {
		return m.ColumnRoots
	}
	return nil
}

func init() {
	proto.RegisterType((*DataAvailabilityHeader)(nil), "elysium.da.DataAvailabilityHeader")
}

func init() {
	proto.RegisterFile("elysium/da/data_availability_header.proto", fileDescriptor_93b487fd8444a5fd)
}

var fileDescriptor_93b487fd8444a5fd = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4a, 0x4e, 0xcd, 0x49,
	0x2d, 0x2e, 0xc9, 0x4c, 0xd4, 0x4f, 0x01, 0xa1, 0x92, 0xc4, 0xf8, 0xc4, 0xb2, 0xc4, 0xcc, 0x9c,
	0xc4, 0xa4, 0xcc, 0x9c, 0xcc, 0x92, 0xca, 0xf8, 0x8c, 0xd4, 0xc4, 0x94, 0xd4, 0x22, 0xbd, 0x82,
	0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x6e, 0x98, 0x5a, 0xbd, 0x94, 0x44, 0xa5, 0x08, 0x2e, 0x31, 0x97,
	0xc4, 0x92, 0x44, 0x47, 0x24, 0xd5, 0x1e, 0x60, 0xc5, 0x42, 0xd2, 0x5c, 0x9c, 0x45, 0xf9, 0xe5,
	0xf1, 0x45, 0xf9, 0xf9, 0x25, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x3c, 0x41, 0x1c, 0x45, 0xf9,
	0xe5, 0x41, 0x20, 0xbe, 0x90, 0x22, 0x17, 0x4f, 0x72, 0x7e, 0x4e, 0x69, 0x6e, 0x1e, 0x54, 0x9e,
	0x09, 0x2c, 0xcf, 0x0d, 0x11, 0x03, 0x2b, 0x71, 0xf2, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23,
	0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6,
	0x63, 0x39, 0x86, 0x28, 0xd3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d,
	0x98, 0x5b, 0xf2, 0x8b, 0xd2, 0xe1, 0x6c, 0xdd, 0xc4, 0x82, 0x02, 0x7d, 0xb0, 0x5b, 0xf5, 0x91,
	0xbc, 0x95, 0xc4, 0x06, 0x16, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xa3, 0x97, 0x36,
	0xec, 0x00, 0x00, 0x00,
}

func (m *DataAvailabilityHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataAvailabilityHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataAvailabilityHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ColumnRoots) > 0 {
		for iNdEx := len(m.ColumnRoots) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ColumnRoots[iNdEx])
			copy(dAtA[i:], m.ColumnRoots[iNdEx])
			i = encodeVarintDataAvailabilityHeader(dAtA, i, uint64(len(m.ColumnRoots[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RowRoots) > 0 {
		for iNdEx := len(m.RowRoots) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RowRoots[iNdEx])
			copy(dAtA[i:], m.RowRoots[iNdEx])
			i = encodeVarintDataAvailabilityHeader(dAtA, i, uint64(len(m.RowRoots[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDataAvailabilityHeader(dAtA []byte, offset int, v uint64) int {
	offset -= sovDataAvailabilityHeader(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataAvailabilityHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RowRoots) > 0 {
		for _, b := range m.RowRoots {
			l = len(b)
			n += 1 + l + sovDataAvailabilityHeader(uint64(l))
		}
	}
	if len(m.ColumnRoots) > 0 {
		for _, b := range m.ColumnRoots {
			l = len(b)
			n += 1 + l + sovDataAvailabilityHeader(uint64(l))
		}
	}
	return n
}

func sovDataAvailabilityHeader(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDataAvailabilityHeader(x uint64) (n int) {
	return sovDataAvailabilityHeader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataAvailabilityHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataAvailabilityHeader
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
			return fmt.Errorf("proto: DataAvailabilityHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataAvailabilityHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RowRoots", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataAvailabilityHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDataAvailabilityHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDataAvailabilityHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RowRoots = append(m.RowRoots, make([]byte, postIndex-iNdEx))
			copy(m.RowRoots[len(m.RowRoots)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnRoots", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataAvailabilityHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDataAvailabilityHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDataAvailabilityHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ColumnRoots = append(m.ColumnRoots, make([]byte, postIndex-iNdEx))
			copy(m.ColumnRoots[len(m.ColumnRoots)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataAvailabilityHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDataAvailabilityHeader
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDataAvailabilityHeader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDataAvailabilityHeader
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
					return 0, ErrIntOverflowDataAvailabilityHeader
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
					return 0, ErrIntOverflowDataAvailabilityHeader
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
				return 0, ErrInvalidLengthDataAvailabilityHeader
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDataAvailabilityHeader
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDataAvailabilityHeader
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDataAvailabilityHeader        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDataAvailabilityHeader          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDataAvailabilityHeader = fmt.Errorf("proto: unexpected end of group")
)
