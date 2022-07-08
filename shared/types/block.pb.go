// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

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

type Block struct {
	BlockHash            []byte   `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	BlockNumber          uint64   `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	MinerHash            string   `protobuf:"bytes,3,opt,name=miner_hash,json=minerHash,proto3" json:"miner_hash,omitempty"`
	Timestamp            uint64   `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	GasLimit             uint64   `protobuf:"varint,5,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasUsed              uint64   `protobuf:"varint,6,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	SizeBytes            uint32   `protobuf:"varint,7,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	TxsCount             uint32   `protobuf:"varint,8,opt,name=txs_count,json=txsCount,proto3" json:"txs_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{0}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Block.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return m.Size()
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *Block) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *Block) GetMinerHash() string {
	if m != nil {
		return m.MinerHash
	}
	return ""
}

func (m *Block) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Block) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *Block) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *Block) GetSizeBytes() uint32 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

func (m *Block) GetTxsCount() uint32 {
	if m != nil {
		return m.TxsCount
	}
	return 0
}

type BlockDetails struct {
	BlockHash            []byte   `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	BlockNumber          uint64   `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Coinbase             []byte   `protobuf:"bytes,3,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	Difficulty           uint64   `protobuf:"varint,4,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Nonce                []byte   `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ParentHash           []byte   `protobuf:"bytes,6,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Timestamp            uint64   `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	GasLimit             uint64   `protobuf:"varint,8,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasUsed              uint64   `protobuf:"varint,9,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	SizeBytes            uint32   `protobuf:"varint,10,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	TxsCount             uint32   `protobuf:"varint,11,opt,name=txs_count,json=txsCount,proto3" json:"txs_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockDetails) Reset()         { *m = BlockDetails{} }
func (m *BlockDetails) String() string { return proto.CompactTextString(m) }
func (*BlockDetails) ProtoMessage()    {}
func (*BlockDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{1}
}
func (m *BlockDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockDetails.Merge(m, src)
}
func (m *BlockDetails) XXX_Size() int {
	return m.Size()
}
func (m *BlockDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockDetails.DiscardUnknown(m)
}

var xxx_messageInfo_BlockDetails proto.InternalMessageInfo

func (m *BlockDetails) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *BlockDetails) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *BlockDetails) GetCoinbase() []byte {
	if m != nil {
		return m.Coinbase
	}
	return nil
}

func (m *BlockDetails) GetDifficulty() uint64 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *BlockDetails) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *BlockDetails) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *BlockDetails) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockDetails) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *BlockDetails) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *BlockDetails) GetSizeBytes() uint32 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

func (m *BlockDetails) GetTxsCount() uint32 {
	if m != nil {
		return m.TxsCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Block)(nil), "com.ankr.Block")
	proto.RegisterType((*BlockDetails)(nil), "com.ankr.BlockDetails")
}

func init() { proto.RegisterFile("block.proto", fileDescriptor_8e550b1f5926e92d) }

var fileDescriptor_8e550b1f5926e92d = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x3d, 0xaf, 0x9b, 0x30,
	0x14, 0x86, 0xeb, 0x34, 0x1f, 0x70, 0x42, 0x3b, 0x58, 0x19, 0xe8, 0x17, 0xa5, 0x99, 0x98, 0x92,
	0xa1, 0x63, 0x37, 0xda, 0xa1, 0x43, 0x55, 0x55, 0x48, 0x5d, 0xba, 0x20, 0x03, 0x4e, 0xb0, 0x02,
	0x06, 0x71, 0x8c, 0x94, 0xf4, 0x97, 0xf4, 0x27, 0x75, 0xec, 0xd2, 0xbd, 0x4a, 0xff, 0xc5, 0x9d,
	0xae, 0x7c, 0xb8, 0x9f, 0xba, 0x12, 0xd3, 0x5d, 0x90, 0xcf, 0xf3, 0x9a, 0x17, 0xfc, 0xc8, 0xb0,
	0xcc, 0xaa, 0x26, 0x3f, 0x6c, 0xda, 0xae, 0x31, 0x0d, 0x77, 0xf2, 0xa6, 0xde, 0x08, 0x7d, 0xe8,
	0xd6, 0x17, 0x0c, 0x66, 0xb1, 0x4d, 0xf8, 0x1b, 0x00, 0xda, 0x92, 0x96, 0x02, 0x4b, 0x9f, 0x85,
	0x2c, 0xf2, 0x12, 0x97, 0xc8, 0x67, 0x81, 0x25, 0x7f, 0x07, 0xde, 0x10, 0xeb, 0xbe, 0xce, 0x64,
	0xe7, 0x4f, 0x42, 0x16, 0x4d, 0x93, 0xa1, 0xf5, 0x2b, 0x21, 0xdb, 0x50, 0x2b, 0x2d, 0xbb, 0xa1,
	0xe1, 0x69, 0xc8, 0x22, 0x37, 0x71, 0x89, 0x50, 0xc3, 0x6b, 0x70, 0x8d, 0xaa, 0x25, 0x1a, 0x51,
	0xb7, 0xfe, 0x94, 0x5e, 0xbf, 0x05, 0xfc, 0x15, 0xb8, 0x7b, 0x81, 0x69, 0xa5, 0x6a, 0x65, 0xfc,
	0x19, 0xa5, 0xce, 0x5e, 0xe0, 0x17, 0x3b, 0xf3, 0x17, 0x60, 0xd7, 0x69, 0x8f, 0xb2, 0xf0, 0xe7,
	0x94, 0x2d, 0xf6, 0x02, 0xbf, 0xa3, 0x2c, 0xec, 0x47, 0x51, 0xfd, 0x94, 0x69, 0x76, 0x32, 0x12,
	0xfd, 0x45, 0xc8, 0xa2, 0x67, 0x89, 0x6b, 0x49, 0x6c, 0x81, 0xad, 0x35, 0x47, 0x4c, 0xf3, 0xa6,
	0xd7, 0xc6, 0x77, 0x28, 0x75, 0xcc, 0x11, 0x3f, 0xda, 0x79, 0xfd, 0x77, 0x02, 0x1e, 0x1d, 0xfe,
	0x93, 0x34, 0x42, 0x55, 0xf8, 0x08, 0x0e, 0x5e, 0x82, 0x93, 0x37, 0x4a, 0x67, 0x02, 0x25, 0x19,
	0xf0, 0x92, 0x9b, 0x99, 0x07, 0x00, 0x85, 0xda, 0xed, 0x54, 0xde, 0x57, 0xe6, 0x74, 0x65, 0xe0,
	0x0e, 0xe1, 0x2b, 0x98, 0xe9, 0x46, 0xe7, 0x92, 0x8e, 0xef, 0x25, 0xc3, 0xc0, 0xdf, 0xc2, 0xb2,
	0x15, 0x9d, 0xd4, 0x66, 0xf8, 0xa9, 0x39, 0x65, 0x30, 0xa0, 0x87, 0x5e, 0x17, 0xa3, 0x5e, 0x9d,
	0x11, 0xaf, 0xee, 0x98, 0x57, 0x18, 0xf5, 0xba, 0xbc, 0xef, 0x35, 0x8e, 0x7f, 0x9f, 0x03, 0xf6,
	0xe7, 0x1c, 0xb0, 0x7f, 0xe7, 0x80, 0xfd, 0xfa, 0x1f, 0x3c, 0x81, 0xe7, 0xd7, 0x17, 0x6e, 0xb8,
	0x80, 0x31, 0x90, 0xf6, 0x6f, 0x76, 0xfd, 0x63, 0xb5, 0xd9, 0x62, 0x29, 0x3a, 0x59, 0x6c, 0xcd,
	0xa9, 0x95, 0xf8, 0x81, 0x9e, 0xd9, 0x9c, 0x36, 0xbe, 0xbf, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x00,
	0xe1, 0x47, 0x05, 0xb8, 0x02, 0x00, 0x00,
}

func (m *Block) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Block) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Block) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TxsCount != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.TxsCount))
		i--
		dAtA[i] = 0x40
	}
	if m.SizeBytes != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.SizeBytes))
		i--
		dAtA[i] = 0x38
	}
	if m.GasUsed != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x30
	}
	if m.GasLimit != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x28
	}
	if m.Timestamp != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.MinerHash) > 0 {
		i -= len(m.MinerHash)
		copy(dAtA[i:], m.MinerHash)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.MinerHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockNumber != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x10
	}
	if len(m.BlockHash) > 0 {
		i -= len(m.BlockHash)
		copy(dAtA[i:], m.BlockHash)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.BlockHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TxsCount != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.TxsCount))
		i--
		dAtA[i] = 0x58
	}
	if m.SizeBytes != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.SizeBytes))
		i--
		dAtA[i] = 0x50
	}
	if m.GasUsed != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x48
	}
	if m.GasLimit != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.GasLimit))
		i--
		dAtA[i] = 0x40
	}
	if m.Timestamp != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ParentHash) > 0 {
		i -= len(m.ParentHash)
		copy(dAtA[i:], m.ParentHash)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.ParentHash)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Difficulty != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Difficulty))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Coinbase) > 0 {
		i -= len(m.Coinbase)
		copy(dAtA[i:], m.Coinbase)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Coinbase)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockNumber != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x10
	}
	if len(m.BlockHash) > 0 {
		i -= len(m.BlockHash)
		copy(dAtA[i:], m.BlockHash)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.BlockHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Block) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovBlock(uint64(m.BlockNumber))
	}
	l = len(m.MinerHash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovBlock(uint64(m.Timestamp))
	}
	if m.GasLimit != 0 {
		n += 1 + sovBlock(uint64(m.GasLimit))
	}
	if m.GasUsed != 0 {
		n += 1 + sovBlock(uint64(m.GasUsed))
	}
	if m.SizeBytes != 0 {
		n += 1 + sovBlock(uint64(m.SizeBytes))
	}
	if m.TxsCount != 0 {
		n += 1 + sovBlock(uint64(m.TxsCount))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlockDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovBlock(uint64(m.BlockNumber))
	}
	l = len(m.Coinbase)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Difficulty != 0 {
		n += 1 + sovBlock(uint64(m.Difficulty))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.ParentHash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovBlock(uint64(m.Timestamp))
	}
	if m.GasLimit != 0 {
		n += 1 + sovBlock(uint64(m.GasLimit))
	}
	if m.GasUsed != 0 {
		n += 1 + sovBlock(uint64(m.GasUsed))
	}
	if m.SizeBytes != 0 {
		n += 1 + sovBlock(uint64(m.SizeBytes))
	}
	if m.TxsCount != 0 {
		n += 1 + sovBlock(uint64(m.TxsCount))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Block) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = append(m.BlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockHash == nil {
				m.BlockHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinerHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinerHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SizeBytes", wireType)
			}
			m.SizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SizeBytes |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxsCount", wireType)
			}
			m.TxsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxsCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *BlockDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: BlockDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = append(m.BlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockHash == nil {
				m.BlockHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coinbase", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coinbase = append(m.Coinbase[:0], dAtA[iNdEx:postIndex]...)
			if m.Coinbase == nil {
				m.Coinbase = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Difficulty", wireType)
			}
			m.Difficulty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Difficulty |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = append(m.Nonce[:0], dAtA[iNdEx:postIndex]...)
			if m.Nonce == nil {
				m.Nonce = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentHash = append(m.ParentHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ParentHash == nil {
				m.ParentHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SizeBytes", wireType)
			}
			m.SizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SizeBytes |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxsCount", wireType)
			}
			m.TxsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxsCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)
