// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transaction.proto

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

type TransactionStatus int32

const (
	TransactionStatus_REVERTED TransactionStatus = 0
	TransactionStatus_SUCCESS  TransactionStatus = 1
)

var TransactionStatus_name = map[int32]string{
	0: "REVERTED",
	1: "SUCCESS",
}

var TransactionStatus_value = map[string]int32{
	"REVERTED": 0,
	"SUCCESS":  1,
}

func (x TransactionStatus) String() string {
	return proto.EnumName(TransactionStatus_name, int32(x))
}

func (TransactionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}

type Transaction struct {
	TxHash               []byte            `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Status               TransactionStatus `protobuf:"varint,2,opt,name=status,proto3,enum=com.ankr.TransactionStatus" json:"status,omitempty"`
	Gas                  uint64            `protobuf:"varint,3,opt,name=gas,proto3" json:"gas,omitempty"`
	CumulativeGasUsed    uint64            `protobuf:"varint,4,opt,name=cumulative_gas_used,json=cumulativeGasUsed,proto3" json:"cumulative_gas_used,omitempty"`
	GasUsed              uint64            `protobuf:"varint,5,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	GasPrice             uint64            `protobuf:"varint,6,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	TxIndex              uint64            `protobuf:"varint,7,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	Input                []byte            `protobuf:"bytes,8,opt,name=input,proto3" json:"input,omitempty"`
	Nonce                uint64            `protobuf:"varint,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Value                string            `protobuf:"bytes,10,opt,name=value,proto3" json:"value,omitempty"`
	Error                string            `protobuf:"bytes,11,opt,name=error,proto3" json:"error,omitempty"`
	BlockHash            []byte            `protobuf:"bytes,12,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	BlockNumber          uint64            `protobuf:"varint,13,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Sender               []byte            `protobuf:"bytes,14,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            []byte            `protobuf:"bytes,15,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Contract             []byte            `protobuf:"bytes,16,opt,name=contract,proto3" json:"contract,omitempty"`
	RevertReason         string            `protobuf:"bytes,17,opt,name=revert_reason,json=revertReason,proto3" json:"revert_reason,omitempty"`
	Type                 uint32            `protobuf:"varint,18,opt,name=type,proto3" json:"type,omitempty"`
	InternalFailed       bool              `protobuf:"varint,19,opt,name=internal_failed,json=internalFailed,proto3" json:"internal_failed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return m.Size()
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *Transaction) GetStatus() TransactionStatus {
	if m != nil {
		return m.Status
	}
	return TransactionStatus_REVERTED
}

func (m *Transaction) GetGas() uint64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *Transaction) GetCumulativeGasUsed() uint64 {
	if m != nil {
		return m.CumulativeGasUsed
	}
	return 0
}

func (m *Transaction) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *Transaction) GetGasPrice() uint64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *Transaction) GetTxIndex() uint64 {
	if m != nil {
		return m.TxIndex
	}
	return 0
}

func (m *Transaction) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *Transaction) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Transaction) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Transaction) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *Transaction) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *Transaction) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Transaction) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *Transaction) GetContract() []byte {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *Transaction) GetRevertReason() string {
	if m != nil {
		return m.RevertReason
	}
	return ""
}

func (m *Transaction) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Transaction) GetInternalFailed() bool {
	if m != nil {
		return m.InternalFailed
	}
	return false
}

func init() {
	proto.RegisterEnum("com.ankr.TransactionStatus", TransactionStatus_name, TransactionStatus_value)
	proto.RegisterType((*Transaction)(nil), "com.ankr.Transaction")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490) }

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xbb, 0x34, 0x4d, 0x9c, 0xc9, 0x9f, 0x26, 0xdb, 0x0a, 0x16, 0x0a, 0x91, 0x81, 0x03,
	0x16, 0x07, 0x57, 0xa2, 0x47, 0x6e, 0x29, 0xe1, 0xcf, 0x05, 0x55, 0x4e, 0xcb, 0x81, 0x8b, 0xb5,
	0xb1, 0x87, 0xc4, 0x6a, 0xb2, 0x8e, 0x76, 0xd7, 0x91, 0x79, 0x13, 0x9e, 0x86, 0x33, 0x47, 0x1e,
	0x01, 0x85, 0x17, 0x41, 0x3b, 0xdb, 0xa6, 0x48, 0xbd, 0x58, 0xf3, 0xfd, 0xbe, 0x6f, 0xc6, 0xeb,
	0xf1, 0xc2, 0xd0, 0x6a, 0xa9, 0x8c, 0xcc, 0x6c, 0x51, 0xaa, 0x78, 0xad, 0x4b, 0x5b, 0xf2, 0x20,
	0x2b, 0x57, 0xb1, 0x54, 0xd7, 0xfa, 0xc5, 0xcf, 0x06, 0x74, 0x2e, 0xef, 0x7c, 0xfe, 0x08, 0x5a,
	0xb6, 0x4e, 0x17, 0xd2, 0x2c, 0x04, 0x0b, 0x59, 0xd4, 0x4d, 0x9a, 0xb6, 0xfe, 0x28, 0xcd, 0x82,
	0x9f, 0x41, 0xd3, 0x58, 0x69, 0x2b, 0x23, 0x1e, 0x84, 0x2c, 0xea, 0xbf, 0x39, 0x89, 0x6f, 0x67,
	0xc4, 0xff, 0xf5, 0x4f, 0x29, 0x92, 0xdc, 0x44, 0xf9, 0x00, 0xf6, 0xe7, 0xd2, 0x88, 0xfd, 0x90,
	0x45, 0x8d, 0xc4, 0x95, 0x3c, 0x86, 0xa3, 0xac, 0x5a, 0x55, 0x4b, 0x69, 0x8b, 0x0d, 0xa6, 0x73,
	0x69, 0xd2, 0xca, 0x60, 0x2e, 0x1a, 0x94, 0x18, 0xde, 0x59, 0x1f, 0xa4, 0xb9, 0x32, 0x98, 0xf3,
	0xc7, 0x10, 0xec, 0x42, 0x07, 0x14, 0x6a, 0xcd, 0x6f, 0xac, 0x13, 0x68, 0x3b, 0x6b, 0xad, 0x8b,
	0x0c, 0x45, 0x93, 0x3c, 0x97, 0xbd, 0x70, 0xda, 0xf5, 0xd9, 0x3a, 0x2d, 0x54, 0x8e, 0xb5, 0x68,
	0xf9, 0x3e, 0x5b, 0x7f, 0x72, 0x92, 0x1f, 0xc3, 0x41, 0xa1, 0xd6, 0x95, 0x15, 0x01, 0x7d, 0xa0,
	0x17, 0x8e, 0xaa, 0x52, 0x65, 0x28, 0xda, 0x94, 0xf6, 0xc2, 0xd1, 0x8d, 0x5c, 0x56, 0x28, 0x20,
	0x64, 0x51, 0x3b, 0xf1, 0xc2, 0x51, 0xd4, 0xba, 0xd4, 0xa2, 0xe3, 0x29, 0x09, 0xfe, 0x0c, 0x60,
	0xb6, 0x2c, 0xb3, 0x6b, 0xbf, 0xbd, 0x2e, 0x0d, 0x6f, 0x13, 0xa1, 0x05, 0x3e, 0x87, 0xae, 0xb7,
	0x55, 0xb5, 0x9a, 0xa1, 0x16, 0x3d, 0x7a, 0x4f, 0x87, 0xd8, 0x67, 0x42, 0xfc, 0x21, 0x34, 0x0d,
	0xaa, 0x1c, 0xb5, 0xe8, 0xfb, 0xdd, 0x7b, 0xc5, 0x9f, 0x42, 0x5b, 0x63, 0x56, 0xac, 0x0b, 0x54,
	0x56, 0x1c, 0xfa, 0xc1, 0x3b, 0xc0, 0x9f, 0x40, 0x90, 0x95, 0xca, 0x6a, 0x99, 0x59, 0x31, 0x20,
	0x73, 0xa7, 0xf9, 0x4b, 0xe8, 0x69, 0xdc, 0xa0, 0xb6, 0xa9, 0x46, 0x69, 0x4a, 0x25, 0x86, 0x74,
	0xe2, 0xae, 0x87, 0x09, 0x31, 0xce, 0xa1, 0x61, 0xbf, 0xaf, 0x51, 0xf0, 0x90, 0x45, 0xbd, 0x84,
	0x6a, 0xfe, 0x0a, 0x0e, 0x0b, 0x65, 0x51, 0x2b, 0xb9, 0x4c, 0xbf, 0xc9, 0x62, 0x89, 0xb9, 0x38,
	0x0a, 0x59, 0x14, 0x24, 0xfd, 0x5b, 0xfc, 0x9e, 0xe8, 0xeb, 0x18, 0x86, 0xf7, 0xfe, 0x3f, 0xef,
	0x42, 0x90, 0x4c, 0xbe, 0x4c, 0x92, 0xcb, 0xc9, 0xbb, 0xc1, 0x1e, 0xef, 0x40, 0x6b, 0x7a, 0x75,
	0x7e, 0x3e, 0x99, 0x4e, 0x07, 0x6c, 0x3c, 0xfe, 0xb5, 0x1d, 0xb1, 0xdf, 0xdb, 0x11, 0xfb, 0xb3,
	0x1d, 0xb1, 0x1f, 0x7f, 0x47, 0x7b, 0xd0, 0xdf, 0x5d, 0x24, 0xba, 0x9c, 0x63, 0x18, 0xbb, 0x95,
	0x5c, 0xb8, 0xfa, 0xeb, 0x71, 0x7c, 0x6a, 0x16, 0x52, 0x63, 0x7e, 0xea, 0x4e, 0x65, 0xde, 0xd2,
	0x73, 0xd6, 0xa4, 0xe0, 0xd9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x2b, 0xa0, 0x63, 0xda,
	0x02, 0x00, 0x00,
}

func (m *Transaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Transaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.InternalFailed {
		i--
		if m.InternalFailed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x98
	}
	if m.Type != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	if len(m.RevertReason) > 0 {
		i -= len(m.RevertReason)
		copy(dAtA[i:], m.RevertReason)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.RevertReason)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x72
	}
	if m.BlockNumber != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x68
	}
	if len(m.BlockHash) > 0 {
		i -= len(m.BlockHash)
		copy(dAtA[i:], m.BlockHash)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.BlockHash)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x52
	}
	if m.Nonce != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Input) > 0 {
		i -= len(m.Input)
		copy(dAtA[i:], m.Input)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.Input)))
		i--
		dAtA[i] = 0x42
	}
	if m.TxIndex != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.TxIndex))
		i--
		dAtA[i] = 0x38
	}
	if m.GasPrice != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.GasPrice))
		i--
		dAtA[i] = 0x30
	}
	if m.GasUsed != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x28
	}
	if m.CumulativeGasUsed != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.CumulativeGasUsed))
		i--
		dAtA[i] = 0x20
	}
	if m.Gas != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.Gas))
		i--
		dAtA[i] = 0x18
	}
	if m.Status != 0 {
		i = encodeVarintTransaction(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintTransaction(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransaction(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransaction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Transaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovTransaction(uint64(m.Status))
	}
	if m.Gas != 0 {
		n += 1 + sovTransaction(uint64(m.Gas))
	}
	if m.CumulativeGasUsed != 0 {
		n += 1 + sovTransaction(uint64(m.CumulativeGasUsed))
	}
	if m.GasUsed != 0 {
		n += 1 + sovTransaction(uint64(m.GasUsed))
	}
	if m.GasPrice != 0 {
		n += 1 + sovTransaction(uint64(m.GasPrice))
	}
	if m.TxIndex != 0 {
		n += 1 + sovTransaction(uint64(m.TxIndex))
	}
	l = len(m.Input)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovTransaction(uint64(m.Nonce))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovTransaction(uint64(m.BlockNumber))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTransaction(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 2 + l + sovTransaction(uint64(l))
	}
	l = len(m.RevertReason)
	if l > 0 {
		n += 2 + l + sovTransaction(uint64(l))
	}
	if m.Type != 0 {
		n += 2 + sovTransaction(uint64(m.Type))
	}
	if m.InternalFailed {
		n += 3
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTransaction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransaction(x uint64) (n int) {
	return sovTransaction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Transaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransaction
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
			return fmt.Errorf("proto: Transaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= TransactionStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gas", wireType)
			}
			m.Gas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeGasUsed", wireType)
			}
			m.CumulativeGasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CumulativeGasUsed |= uint64(b&0x7F) << shift
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
					return ErrIntOverflowTransaction
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
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			m.GasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIndex", wireType)
			}
			m.TxIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Input", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Input = append(m.Input[:0], dAtA[iNdEx:postIndex]...)
			if m.Input == nil {
				m.Input = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = append(m.BlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockHash == nil {
				m.BlockHash = []byte{}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = append(m.Recipient[:0], dAtA[iNdEx:postIndex]...)
			if m.Recipient == nil {
				m.Recipient = []byte{}
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = append(m.Contract[:0], dAtA[iNdEx:postIndex]...)
			if m.Contract == nil {
				m.Contract = []byte{}
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevertReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
				return ErrInvalidLengthTransaction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTransaction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RevertReason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 19:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalFailed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransaction
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
			m.InternalFailed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTransaction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransaction
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
func skipTransaction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
					return 0, ErrIntOverflowTransaction
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
				return 0, ErrInvalidLengthTransaction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransaction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransaction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransaction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransaction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransaction = fmt.Errorf("proto: unexpected end of group")
)
