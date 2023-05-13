// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elysium/blob/v1/event.proto

package types

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

// EventPayForBlobs defines an event that is emitted after a pay for blob has
// been processed.
type EventPayForBlobs struct {
	Signer    string   `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	BlobSizes []uint32 `protobuf:"varint,2,rep,packed,name=blob_sizes,json=blobSizes,proto3" json:"blob_sizes,omitempty"`
	// namespaces is a list of namespaces that the blobs in blob_sizes belong to.
	// A namespace has length of 33 bytes where the first byte is the
	// namespaceVersion and the subsequent 32 bytes are the namespaceID.
	Namespaces [][]byte `protobuf:"bytes,3,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (m *EventPayForBlobs) Reset()         { *m = EventPayForBlobs{} }
func (m *EventPayForBlobs) String() string { return proto.CompactTextString(m) }
func (*EventPayForBlobs) ProtoMessage()    {}
func (*EventPayForBlobs) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d90f0a63835a06e, []int{0}
}
func (m *EventPayForBlobs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventPayForBlobs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventPayForBlobs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventPayForBlobs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPayForBlobs.Merge(m, src)
}
func (m *EventPayForBlobs) XXX_Size() int {
	return m.Size()
}
func (m *EventPayForBlobs) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPayForBlobs.DiscardUnknown(m)
}

var xxx_messageInfo_EventPayForBlobs proto.InternalMessageInfo

func (m *EventPayForBlobs) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *EventPayForBlobs) GetBlobSizes() []uint32 {
	if m != nil {
		return m.BlobSizes
	}
	return nil
}

func (m *EventPayForBlobs) GetNamespaces() [][]byte {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func init() {
	proto.RegisterType((*EventPayForBlobs)(nil), "elysium.blob.v1.EventPayForBlobs")
}

func init() { proto.RegisterFile("elysium/blob/v1/event.proto", fileDescriptor_9d90f0a63835a06e) }

var fileDescriptor_9d90f0a63835a06e = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x4e, 0xcd, 0x49,
	0x2d, 0x2e, 0xc9, 0x4c, 0xd4, 0x4f, 0xca, 0xc9, 0x4f, 0xd2, 0x2f, 0x33, 0xd4, 0x4f, 0x2d, 0x4b,
	0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0xc9, 0xea, 0x81, 0x64, 0xf5,
	0xca, 0x0c, 0x95, 0x32, 0xb9, 0x04, 0x5c, 0x41, 0x0a, 0x02, 0x12, 0x2b, 0xdd, 0xf2, 0x8b, 0x9c,
	0x72, 0xf2, 0x93, 0x8a, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x33, 0xd3, 0xf3, 0x52, 0x8b, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21, 0x59, 0x2e, 0x2e, 0x90, 0xb6, 0xf8, 0xe2, 0xcc,
	0xaa, 0xd4, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0xde, 0x20, 0x4e, 0x90, 0x48, 0x30, 0x48, 0x40,
	0x48, 0x8e, 0x8b, 0x2b, 0x2f, 0x31, 0x37, 0xb5, 0xb8, 0x20, 0x31, 0x39, 0xb5, 0x58, 0x82, 0x59,
	0x81, 0x59, 0x83, 0x27, 0x08, 0x49, 0xc4, 0xc9, 0xeb, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f,
	0xe5, 0x18, 0xa2, 0x0c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x61,
	0x2e, 0xcc, 0x2f, 0x4a, 0x87, 0xb3, 0x75, 0x13, 0x0b, 0x0a, 0xf4, 0x2b, 0x20, 0x3e, 0x2a, 0xa9,
	0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xfb, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x7b,
	0x9e, 0x60, 0xef, 0x00, 0x00, 0x00,
}

func (m *EventPayForBlobs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPayForBlobs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventPayForBlobs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Namespaces) > 0 {
		for iNdEx := len(m.Namespaces) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Namespaces[iNdEx])
			copy(dAtA[i:], m.Namespaces[iNdEx])
			i = encodeVarintEvent(dAtA, i, uint64(len(m.Namespaces[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.BlobSizes) > 0 {
		dAtA2 := make([]byte, len(m.BlobSizes)*10)
		var j1 int
		for _, num := range m.BlobSizes {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintEvent(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventPayForBlobs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if len(m.BlobSizes) > 0 {
		l = 0
		for _, e := range m.BlobSizes {
			l += sovEvent(uint64(e))
		}
		n += 1 + sovEvent(uint64(l)) + l
	}
	if len(m.Namespaces) > 0 {
		for _, b := range m.Namespaces {
			l = len(b)
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventPayForBlobs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventPayForBlobs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPayForBlobs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.BlobSizes = append(m.BlobSizes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvent
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthEvent
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthEvent
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.BlobSizes) == 0 {
					m.BlobSizes = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEvent
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.BlobSizes = append(m.BlobSizes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BlobSizes", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespaces", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespaces = append(m.Namespaces, make([]byte, postIndex-iNdEx))
			copy(m.Namespaces[len(m.Namespaces)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
