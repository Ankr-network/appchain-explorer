// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: address.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Address struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Balance              uint64   `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Tokens               uint32   `protobuf:"varint,3,opt,name=tokens,proto3" json:"tokens,omitempty"`
	Transactions         uint32   `protobuf:"varint,4,opt,name=transactions,proto3" json:"transactions,omitempty"`
	GasUsed              uint64   `protobuf:"varint,5,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	BlocksValidated      uint32   `protobuf:"varint,6,opt,name=blocks_validated,json=blocksValidated,proto3" json:"blocks_validated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{0}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Address.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return m.Size()
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Address) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Address) GetTokens() uint32 {
	if m != nil {
		return m.Tokens
	}
	return 0
}

func (m *Address) GetTransactions() uint32 {
	if m != nil {
		return m.Transactions
	}
	return 0
}

func (m *Address) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *Address) GetBlocksValidated() uint32 {
	if m != nil {
		return m.BlocksValidated
	}
	return 0
}

func init() {
	proto.RegisterType((*Address)(nil), "com.ankr.Address")
}

func init() { proto.RegisterFile("address.proto", fileDescriptor_982c640dad8fe78e) }

var fileDescriptor_982c640dad8fe78e = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4c, 0x49, 0x29,
	0x4a, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0xce, 0xcf, 0xd5, 0x4b,
	0xcc, 0xcb, 0x2e, 0x52, 0xda, 0xc9, 0xc8, 0xc5, 0xee, 0x08, 0x91, 0x13, 0x12, 0xe2, 0x62, 0xc9,
	0x48, 0x2c, 0xce, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x24, 0xb8, 0xd8,
	0x93, 0x12, 0x73, 0x12, 0xf3, 0x92, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x58, 0x82, 0x60, 0x5c,
	0x21, 0x31, 0x2e, 0xb6, 0x92, 0xfc, 0xec, 0xd4, 0xbc, 0x62, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xde,
	0x20, 0x28, 0x4f, 0x48, 0x89, 0x8b, 0xa7, 0xa4, 0x28, 0x31, 0xaf, 0x38, 0x31, 0xb9, 0x24, 0x33,
	0x3f, 0xaf, 0x58, 0x82, 0x05, 0x2c, 0x8b, 0x22, 0x26, 0x24, 0xc9, 0xc5, 0x91, 0x9e, 0x58, 0x1c,
	0x5f, 0x5a, 0x9c, 0x9a, 0x22, 0xc1, 0x0a, 0x31, 0x36, 0x3d, 0xb1, 0x38, 0xb4, 0x38, 0x35, 0x45,
	0x48, 0x93, 0x4b, 0x20, 0x29, 0x27, 0x3f, 0x39, 0xbb, 0x38, 0xbe, 0x2c, 0x31, 0x27, 0x33, 0x25,
	0xb1, 0x24, 0x35, 0x45, 0x82, 0x0d, 0x6c, 0x04, 0x3f, 0x44, 0x3c, 0x0c, 0x26, 0xec, 0xe4, 0x72,
	0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7,
	0xc0, 0xc5, 0x07, 0xf3, 0x17, 0xc4, 0x9f, 0x4e, 0x3c, 0x50, 0xaf, 0x05, 0x80, 0x78, 0x51, 0x22,
	0x7a, 0xfa, 0xc5, 0x19, 0x89, 0x45, 0xa9, 0x29, 0xfa, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0xd6, 0x60,
	0x32, 0x89, 0x0d, 0xac, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xc7, 0x4c, 0x40, 0x23,
	0x01, 0x00, 0x00,
}

func (m *Address) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Address) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Address) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.BlocksValidated != 0 {
		i = encodeVarintAddress(dAtA, i, uint64(m.BlocksValidated))
		i--
		dAtA[i] = 0x30
	}
	if m.GasUsed != 0 {
		i = encodeVarintAddress(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x28
	}
	if m.Transactions != 0 {
		i = encodeVarintAddress(dAtA, i, uint64(m.Transactions))
		i--
		dAtA[i] = 0x20
	}
	if m.Tokens != 0 {
		i = encodeVarintAddress(dAtA, i, uint64(m.Tokens))
		i--
		dAtA[i] = 0x18
	}
	if m.Balance != 0 {
		i = encodeVarintAddress(dAtA, i, uint64(m.Balance))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintAddress(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAddress(dAtA []byte, offset int, v uint64) int {
	offset -= sovAddress(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Address) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovAddress(uint64(l))
	}
	if m.Balance != 0 {
		n += 1 + sovAddress(uint64(m.Balance))
	}
	if m.Tokens != 0 {
		n += 1 + sovAddress(uint64(m.Tokens))
	}
	if m.Transactions != 0 {
		n += 1 + sovAddress(uint64(m.Transactions))
	}
	if m.GasUsed != 0 {
		n += 1 + sovAddress(uint64(m.GasUsed))
	}
	if m.BlocksValidated != 0 {
		n += 1 + sovAddress(uint64(m.BlocksValidated))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAddress(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAddress(x uint64) (n int) {
	return sovAddress(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Address) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAddress
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
			return fmt.Errorf("proto: Address: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Address: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
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
				return ErrInvalidLengthAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			m.Balance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Balance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			m.Tokens = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tokens |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
			}
			m.Transactions = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Transactions |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksValidated", wireType)
			}
			m.BlocksValidated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksValidated |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAddress
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
func skipAddress(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAddress
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
					return 0, ErrIntOverflowAddress
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
					return 0, ErrIntOverflowAddress
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
				return 0, ErrInvalidLengthAddress
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAddress
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAddress
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAddress        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAddress          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAddress = fmt.Errorf("proto: unexpected end of group")
)
