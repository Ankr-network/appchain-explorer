// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gateway.proto

package types

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type GetRecentBlocksRequest struct {
	FromBlock            uint64   `protobuf:"varint,1,opt,name=from_block,json=fromBlock,proto3" json:"from_block,omitempty"`
	Limit                uint64   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRecentBlocksRequest) Reset()         { *m = GetRecentBlocksRequest{} }
func (m *GetRecentBlocksRequest) String() string { return proto.CompactTextString(m) }
func (*GetRecentBlocksRequest) ProtoMessage()    {}
func (*GetRecentBlocksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}
func (m *GetRecentBlocksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecentBlocksRequest.Unmarshal(m, b)
}
func (m *GetRecentBlocksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecentBlocksRequest.Marshal(b, m, deterministic)
}
func (m *GetRecentBlocksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecentBlocksRequest.Merge(m, src)
}
func (m *GetRecentBlocksRequest) XXX_Size() int {
	return xxx_messageInfo_GetRecentBlocksRequest.Size(m)
}
func (m *GetRecentBlocksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecentBlocksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecentBlocksRequest proto.InternalMessageInfo

func (m *GetRecentBlocksRequest) GetFromBlock() uint64 {
	if m != nil {
		return m.FromBlock
	}
	return 0
}

func (m *GetRecentBlocksRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetRecentBlocksReply struct {
	Blocks               []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRecentBlocksReply) Reset()         { *m = GetRecentBlocksReply{} }
func (m *GetRecentBlocksReply) String() string { return proto.CompactTextString(m) }
func (*GetRecentBlocksReply) ProtoMessage()    {}
func (*GetRecentBlocksReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}
func (m *GetRecentBlocksReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecentBlocksReply.Unmarshal(m, b)
}
func (m *GetRecentBlocksReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecentBlocksReply.Marshal(b, m, deterministic)
}
func (m *GetRecentBlocksReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecentBlocksReply.Merge(m, src)
}
func (m *GetRecentBlocksReply) XXX_Size() int {
	return xxx_messageInfo_GetRecentBlocksReply.Size(m)
}
func (m *GetRecentBlocksReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecentBlocksReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecentBlocksReply proto.InternalMessageInfo

func (m *GetRecentBlocksReply) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type GetRecentTxsRequest struct {
	FromTs               uint64   `protobuf:"varint,1,opt,name=from_ts,json=fromTs,proto3" json:"from_ts,omitempty"`
	Limit                uint64   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRecentTxsRequest) Reset()         { *m = GetRecentTxsRequest{} }
func (m *GetRecentTxsRequest) String() string { return proto.CompactTextString(m) }
func (*GetRecentTxsRequest) ProtoMessage()    {}
func (*GetRecentTxsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{2}
}
func (m *GetRecentTxsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecentTxsRequest.Unmarshal(m, b)
}
func (m *GetRecentTxsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecentTxsRequest.Marshal(b, m, deterministic)
}
func (m *GetRecentTxsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecentTxsRequest.Merge(m, src)
}
func (m *GetRecentTxsRequest) XXX_Size() int {
	return xxx_messageInfo_GetRecentTxsRequest.Size(m)
}
func (m *GetRecentTxsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecentTxsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecentTxsRequest proto.InternalMessageInfo

func (m *GetRecentTxsRequest) GetFromTs() uint64 {
	if m != nil {
		return m.FromTs
	}
	return 0
}

func (m *GetRecentTxsRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetRecentTxsReply struct {
	Txs                  []*Transaction `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetRecentTxsReply) Reset()         { *m = GetRecentTxsReply{} }
func (m *GetRecentTxsReply) String() string { return proto.CompactTextString(m) }
func (*GetRecentTxsReply) ProtoMessage()    {}
func (*GetRecentTxsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{3}
}
func (m *GetRecentTxsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecentTxsReply.Unmarshal(m, b)
}
func (m *GetRecentTxsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecentTxsReply.Marshal(b, m, deterministic)
}
func (m *GetRecentTxsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecentTxsReply.Merge(m, src)
}
func (m *GetRecentTxsReply) XXX_Size() int {
	return xxx_messageInfo_GetRecentTxsReply.Size(m)
}
func (m *GetRecentTxsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecentTxsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecentTxsReply proto.InternalMessageInfo

func (m *GetRecentTxsReply) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type GetAddressRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAddressRequest) Reset()         { *m = GetAddressRequest{} }
func (m *GetAddressRequest) String() string { return proto.CompactTextString(m) }
func (*GetAddressRequest) ProtoMessage()    {}
func (*GetAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{4}
}
func (m *GetAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAddressRequest.Unmarshal(m, b)
}
func (m *GetAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAddressRequest.Marshal(b, m, deterministic)
}
func (m *GetAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAddressRequest.Merge(m, src)
}
func (m *GetAddressRequest) XXX_Size() int {
	return xxx_messageInfo_GetAddressRequest.Size(m)
}
func (m *GetAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAddressRequest proto.InternalMessageInfo

func (m *GetAddressRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetAddressReply struct {
	Address              *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAddressReply) Reset()         { *m = GetAddressReply{} }
func (m *GetAddressReply) String() string { return proto.CompactTextString(m) }
func (*GetAddressReply) ProtoMessage()    {}
func (*GetAddressReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{5}
}
func (m *GetAddressReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAddressReply.Unmarshal(m, b)
}
func (m *GetAddressReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAddressReply.Marshal(b, m, deterministic)
}
func (m *GetAddressReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAddressReply.Merge(m, src)
}
func (m *GetAddressReply) XXX_Size() int {
	return xxx_messageInfo_GetAddressReply.Size(m)
}
func (m *GetAddressReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAddressReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetAddressReply proto.InternalMessageInfo

func (m *GetAddressReply) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type GetBlockByHashOrNumberRequest struct {
	Number               uint64   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockByHashOrNumberRequest) Reset()         { *m = GetBlockByHashOrNumberRequest{} }
func (m *GetBlockByHashOrNumberRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockByHashOrNumberRequest) ProtoMessage()    {}
func (*GetBlockByHashOrNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{6}
}
func (m *GetBlockByHashOrNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockByHashOrNumberRequest.Unmarshal(m, b)
}
func (m *GetBlockByHashOrNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockByHashOrNumberRequest.Marshal(b, m, deterministic)
}
func (m *GetBlockByHashOrNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockByHashOrNumberRequest.Merge(m, src)
}
func (m *GetBlockByHashOrNumberRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockByHashOrNumberRequest.Size(m)
}
func (m *GetBlockByHashOrNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockByHashOrNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockByHashOrNumberRequest proto.InternalMessageInfo

func (m *GetBlockByHashOrNumberRequest) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *GetBlockByHashOrNumberRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetBlockByHashOrNumberReply struct {
	Block                *BlockDetails `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetBlockByHashOrNumberReply) Reset()         { *m = GetBlockByHashOrNumberReply{} }
func (m *GetBlockByHashOrNumberReply) String() string { return proto.CompactTextString(m) }
func (*GetBlockByHashOrNumberReply) ProtoMessage()    {}
func (*GetBlockByHashOrNumberReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{7}
}
func (m *GetBlockByHashOrNumberReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockByHashOrNumberReply.Unmarshal(m, b)
}
func (m *GetBlockByHashOrNumberReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockByHashOrNumberReply.Marshal(b, m, deterministic)
}
func (m *GetBlockByHashOrNumberReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockByHashOrNumberReply.Merge(m, src)
}
func (m *GetBlockByHashOrNumberReply) XXX_Size() int {
	return xxx_messageInfo_GetBlockByHashOrNumberReply.Size(m)
}
func (m *GetBlockByHashOrNumberReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockByHashOrNumberReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockByHashOrNumberReply proto.InternalMessageInfo

func (m *GetBlockByHashOrNumberReply) GetBlock() *BlockDetails {
	if m != nil {
		return m.Block
	}
	return nil
}

type GetTransactionByHashRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionByHashRequest) Reset()         { *m = GetTransactionByHashRequest{} }
func (m *GetTransactionByHashRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionByHashRequest) ProtoMessage()    {}
func (*GetTransactionByHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{8}
}
func (m *GetTransactionByHashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionByHashRequest.Unmarshal(m, b)
}
func (m *GetTransactionByHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionByHashRequest.Marshal(b, m, deterministic)
}
func (m *GetTransactionByHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionByHashRequest.Merge(m, src)
}
func (m *GetTransactionByHashRequest) XXX_Size() int {
	return xxx_messageInfo_GetTransactionByHashRequest.Size(m)
}
func (m *GetTransactionByHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionByHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionByHashRequest proto.InternalMessageInfo

func (m *GetTransactionByHashRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetTransactionByHashReply struct {
	Transaction          *TransactionDetails `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetTransactionByHashReply) Reset()         { *m = GetTransactionByHashReply{} }
func (m *GetTransactionByHashReply) String() string { return proto.CompactTextString(m) }
func (*GetTransactionByHashReply) ProtoMessage()    {}
func (*GetTransactionByHashReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{9}
}
func (m *GetTransactionByHashReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionByHashReply.Unmarshal(m, b)
}
func (m *GetTransactionByHashReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionByHashReply.Marshal(b, m, deterministic)
}
func (m *GetTransactionByHashReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionByHashReply.Merge(m, src)
}
func (m *GetTransactionByHashReply) XXX_Size() int {
	return xxx_messageInfo_GetTransactionByHashReply.Size(m)
}
func (m *GetTransactionByHashReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionByHashReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionByHashReply proto.InternalMessageInfo

func (m *GetTransactionByHashReply) GetTransaction() *TransactionDetails {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRecentBlocksRequest)(nil), "com.ankr.GetRecentBlocksRequest")
	proto.RegisterType((*GetRecentBlocksReply)(nil), "com.ankr.GetRecentBlocksReply")
	proto.RegisterType((*GetRecentTxsRequest)(nil), "com.ankr.GetRecentTxsRequest")
	proto.RegisterType((*GetRecentTxsReply)(nil), "com.ankr.GetRecentTxsReply")
	proto.RegisterType((*GetAddressRequest)(nil), "com.ankr.GetAddressRequest")
	proto.RegisterType((*GetAddressReply)(nil), "com.ankr.GetAddressReply")
	proto.RegisterType((*GetBlockByHashOrNumberRequest)(nil), "com.ankr.GetBlockByHashOrNumberRequest")
	proto.RegisterType((*GetBlockByHashOrNumberReply)(nil), "com.ankr.GetBlockByHashOrNumberReply")
	proto.RegisterType((*GetTransactionByHashRequest)(nil), "com.ankr.GetTransactionByHashRequest")
	proto.RegisterType((*GetTransactionByHashReply)(nil), "com.ankr.GetTransactionByHashReply")
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor_f1a937782ebbded5) }

var fileDescriptor_f1a937782ebbded5 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5f, 0x6f, 0x12, 0x4f,
	0x14, 0x0d, 0x94, 0xd2, 0x5f, 0xef, 0xfe, 0x68, 0x65, 0x8a, 0x14, 0xb6, 0x45, 0xc9, 0x1a, 0x43,
	0x13, 0x0d, 0x9b, 0xe2, 0xa3, 0x06, 0x23, 0x69, 0x82, 0x49, 0xe3, 0x9f, 0x6c, 0x78, 0xd2, 0x87,
	0x66, 0x80, 0x11, 0x08, 0xcb, 0xee, 0xba, 0x33, 0x28, 0xa4, 0xa9, 0x0f, 0xbe, 0xf9, 0xec, 0x47,
	0xd3, 0x8f, 0xe0, 0x07, 0x31, 0x73, 0x67, 0xff, 0x21, 0xbb, 0xed, 0x0b, 0x61, 0xee, 0x3d, 0xf7,
	0xdc, 0xb3, 0x67, 0xce, 0x2e, 0x94, 0x26, 0x54, 0xb0, 0xaf, 0x74, 0xdd, 0xf6, 0x7c, 0x57, 0xb8,
	0xe4, 0xbf, 0x91, 0xbb, 0x68, 0x53, 0x67, 0xee, 0xeb, 0xda, 0xd0, 0x76, 0x47, 0x73, 0x55, 0xd6,
	0xcb, 0xc2, 0xa7, 0x0e, 0xa7, 0x23, 0x31, 0x73, 0x9d, 0xa0, 0x54, 0xa2, 0xe3, 0xb1, 0xcf, 0x38,
	0x0f, 0x8e, 0xa7, 0x13, 0xd7, 0x9d, 0xd8, 0xcc, 0xa4, 0xde, 0xcc, 0xa4, 0x8e, 0xe3, 0x0a, 0x2a,
	0xb1, 0x41, 0xd7, 0x78, 0x03, 0xd5, 0x3e, 0x13, 0x16, 0x1b, 0x31, 0x47, 0xf4, 0x24, 0x2f, 0xb7,
	0xd8, 0xe7, 0x25, 0xe3, 0x82, 0x34, 0x00, 0x3e, 0xf9, 0xee, 0xe2, 0x0a, 0xb7, 0xd5, 0x72, 0xcd,
	0xdc, 0x59, 0xc1, 0xda, 0x97, 0x15, 0x84, 0x91, 0x0a, 0xec, 0xda, 0xb3, 0xc5, 0x4c, 0xd4, 0xf2,
	0xd8, 0x51, 0x07, 0xe3, 0x25, 0x54, 0xb6, 0xe8, 0x3c, 0x7b, 0x4d, 0x5a, 0x50, 0x44, 0x1e, 0x5e,
	0xcb, 0x35, 0x77, 0xce, 0xb4, 0xce, 0x61, 0x3b, 0x7c, 0x9c, 0x36, 0xc2, 0xac, 0xa0, 0x6d, 0x5c,
	0xc0, 0x51, 0x44, 0x30, 0x58, 0x45, 0x62, 0x8e, 0x61, 0x0f, 0xc5, 0x08, 0x1e, 0x28, 0x29, 0xca,
	0xe3, 0x80, 0x67, 0xc8, 0x78, 0x01, 0xe5, 0x4d, 0x16, 0xa5, 0x61, 0x47, 0xac, 0x42, 0x01, 0xf7,
	0x63, 0x01, 0x83, 0xd8, 0x41, 0x4b, 0x22, 0x8c, 0x16, 0x4e, 0xbf, 0x52, 0x2e, 0x86, 0x0a, 0x08,
	0x14, 0xa6, 0x94, 0x4f, 0x71, 0xfd, 0xbe, 0x85, 0xff, 0x8d, 0x2e, 0x1c, 0x26, 0x81, 0x72, 0xc9,
	0x13, 0xd8, 0x0b, 0xec, 0x47, 0xa4, 0xd6, 0x29, 0xc7, 0x8b, 0x42, 0x60, 0x88, 0x30, 0x2e, 0xa1,
	0xd1, 0x67, 0xca, 0xa7, 0xde, 0xfa, 0x35, 0xe5, 0xd3, 0x77, 0xfe, 0xdb, 0xe5, 0x62, 0xc8, 0xfc,
	0x70, 0x69, 0x15, 0x8a, 0x0e, 0x16, 0xc2, 0xa7, 0x56, 0xa7, 0x48, 0x4c, 0x3e, 0x21, 0xe6, 0x12,
	0x4e, 0xb2, 0xc8, 0xa4, 0xb0, 0xa7, 0xb0, 0x1b, 0xdf, 0xa4, 0xd6, 0xa9, 0xfe, 0x73, 0x01, 0x17,
	0x4c, 0xd0, 0x99, 0xcd, 0x2d, 0x05, 0x32, 0xce, 0x91, 0x2c, 0xe1, 0x8c, 0xa2, 0xbc, 0xcd, 0x8c,
	0x8f, 0x50, 0x4f, 0x1f, 0x91, 0xdb, 0xbb, 0xa0, 0x25, 0x82, 0x1a, 0x68, 0x38, 0x4d, 0xbd, 0x83,
	0x50, 0x49, 0x72, 0xa0, 0xf3, 0xbb, 0x00, 0x65, 0x95, 0xa7, 0x91, 0xbb, 0x14, 0x7d, 0xf5, 0x66,
	0x90, 0x39, 0xfa, 0x9f, 0x4c, 0x1b, 0x69, 0xc6, 0x9c, 0xe9, 0xb9, 0xd6, 0x1f, 0xdc, 0x82, 0xf0,
	0xec, 0xb5, 0x51, 0xfd, 0xfe, 0xeb, 0xcf, 0xcf, 0xfc, 0x3d, 0x72, 0x60, 0x7e, 0x39, 0xa7, 0xb6,
	0x37, 0xa5, 0x26, 0x5a, 0x42, 0xae, 0xe0, 0xff, 0x64, 0xa6, 0x48, 0x23, 0x85, 0x27, 0x4e, 0xac,
	0x7e, 0x92, 0xd5, 0x96, 0x3b, 0x8e, 0x70, 0x47, 0x89, 0x68, 0xd1, 0x0e, 0xb1, 0x22, 0x0c, 0x20,
	0x4e, 0x13, 0xd9, 0x9c, 0xdf, 0x0c, 0xa3, 0x5e, 0x4f, 0x6f, 0x4a, 0xea, 0x87, 0x48, 0x5d, 0x27,
	0xc7, 0x11, 0x75, 0x90, 0x36, 0xf3, 0x5a, 0x5e, 0xd3, 0x0d, 0xf9, 0x91, 0xc3, 0x57, 0x3e, 0x25,
	0x28, 0xa4, 0xb5, 0x41, 0x9b, 0x9d, 0x4b, 0xfd, 0xf1, 0xdd, 0xc0, 0x74, 0x2d, 0x68, 0xa5, 0x79,
	0xad, 0x62, 0x7c, 0x43, 0xbe, 0xe1, 0xe7, 0x62, 0x2b, 0x33, 0x64, 0x93, 0x3f, 0x2b, 0x86, 0xfa,
	0xa3, 0xbb, 0x60, 0x52, 0x84, 0x8e, 0x22, 0x2a, 0x84, 0x24, 0xbc, 0x0e, 0xbc, 0xe8, 0x75, 0xe1,
	0x20, 0x62, 0xc0, 0xef, 0x61, 0xaf, 0xba, 0x95, 0xb2, 0xf7, 0xb2, 0xfe, 0xa1, 0xd2, 0x36, 0xf9,
	0x94, 0xfa, 0x6c, 0x6c, 0x8a, 0xb5, 0xc7, 0xf8, 0x73, 0xfc, 0x1d, 0x16, 0x71, 0xe8, 0xd9, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x81, 0xc2, 0x73, 0xac, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlockscoutGatewayClient is the client API for BlockscoutGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockscoutGatewayClient interface {
	GetRecentBlocks(ctx context.Context, in *GetRecentBlocksRequest, opts ...grpc.CallOption) (*GetRecentBlocksReply, error)
	GetRecentTxs(ctx context.Context, in *GetRecentTxsRequest, opts ...grpc.CallOption) (*GetRecentTxsReply, error)
	GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressReply, error)
	GetBlockByHashOrNumber(ctx context.Context, in *GetBlockByHashOrNumberRequest, opts ...grpc.CallOption) (*GetBlockByHashOrNumberReply, error)
	GetTransactionByHash(ctx context.Context, in *GetTransactionByHashRequest, opts ...grpc.CallOption) (*GetTransactionByHashReply, error)
}

type blockscoutGatewayClient struct {
	cc *grpc.ClientConn
}

func NewBlockscoutGatewayClient(cc *grpc.ClientConn) BlockscoutGatewayClient {
	return &blockscoutGatewayClient{cc}
}

func (c *blockscoutGatewayClient) GetRecentBlocks(ctx context.Context, in *GetRecentBlocksRequest, opts ...grpc.CallOption) (*GetRecentBlocksReply, error) {
	out := new(GetRecentBlocksReply)
	err := c.cc.Invoke(ctx, "/com.ankr.BlockscoutGateway/GetRecentBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockscoutGatewayClient) GetRecentTxs(ctx context.Context, in *GetRecentTxsRequest, opts ...grpc.CallOption) (*GetRecentTxsReply, error) {
	out := new(GetRecentTxsReply)
	err := c.cc.Invoke(ctx, "/com.ankr.BlockscoutGateway/GetRecentTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockscoutGatewayClient) GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressReply, error) {
	out := new(GetAddressReply)
	err := c.cc.Invoke(ctx, "/com.ankr.BlockscoutGateway/GetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockscoutGatewayClient) GetBlockByHashOrNumber(ctx context.Context, in *GetBlockByHashOrNumberRequest, opts ...grpc.CallOption) (*GetBlockByHashOrNumberReply, error) {
	out := new(GetBlockByHashOrNumberReply)
	err := c.cc.Invoke(ctx, "/com.ankr.BlockscoutGateway/GetBlockByHashOrNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockscoutGatewayClient) GetTransactionByHash(ctx context.Context, in *GetTransactionByHashRequest, opts ...grpc.CallOption) (*GetTransactionByHashReply, error) {
	out := new(GetTransactionByHashReply)
	err := c.cc.Invoke(ctx, "/com.ankr.BlockscoutGateway/GetTransactionByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockscoutGatewayServer is the server API for BlockscoutGateway service.
type BlockscoutGatewayServer interface {
	GetRecentBlocks(context.Context, *GetRecentBlocksRequest) (*GetRecentBlocksReply, error)
	GetRecentTxs(context.Context, *GetRecentTxsRequest) (*GetRecentTxsReply, error)
	GetAddress(context.Context, *GetAddressRequest) (*GetAddressReply, error)
	GetBlockByHashOrNumber(context.Context, *GetBlockByHashOrNumberRequest) (*GetBlockByHashOrNumberReply, error)
	GetTransactionByHash(context.Context, *GetTransactionByHashRequest) (*GetTransactionByHashReply, error)
}

// UnimplementedBlockscoutGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedBlockscoutGatewayServer struct {
}

func (*UnimplementedBlockscoutGatewayServer) GetRecentBlocks(ctx context.Context, req *GetRecentBlocksRequest) (*GetRecentBlocksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecentBlocks not implemented")
}
func (*UnimplementedBlockscoutGatewayServer) GetRecentTxs(ctx context.Context, req *GetRecentTxsRequest) (*GetRecentTxsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecentTxs not implemented")
}
func (*UnimplementedBlockscoutGatewayServer) GetAddress(ctx context.Context, req *GetAddressRequest) (*GetAddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddress not implemented")
}
func (*UnimplementedBlockscoutGatewayServer) GetBlockByHashOrNumber(ctx context.Context, req *GetBlockByHashOrNumberRequest) (*GetBlockByHashOrNumberReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByHashOrNumber not implemented")
}
func (*UnimplementedBlockscoutGatewayServer) GetTransactionByHash(ctx context.Context, req *GetTransactionByHashRequest) (*GetTransactionByHashReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionByHash not implemented")
}

func RegisterBlockscoutGatewayServer(s *grpc.Server, srv BlockscoutGatewayServer) {
	s.RegisterService(&_BlockscoutGateway_serviceDesc, srv)
}

func _BlockscoutGateway_GetRecentBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecentBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockscoutGatewayServer).GetRecentBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ankr.BlockscoutGateway/GetRecentBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockscoutGatewayServer).GetRecentBlocks(ctx, req.(*GetRecentBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockscoutGateway_GetRecentTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecentTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockscoutGatewayServer).GetRecentTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ankr.BlockscoutGateway/GetRecentTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockscoutGatewayServer).GetRecentTxs(ctx, req.(*GetRecentTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockscoutGateway_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockscoutGatewayServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ankr.BlockscoutGateway/GetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockscoutGatewayServer).GetAddress(ctx, req.(*GetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockscoutGateway_GetBlockByHashOrNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockByHashOrNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockscoutGatewayServer).GetBlockByHashOrNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ankr.BlockscoutGateway/GetBlockByHashOrNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockscoutGatewayServer).GetBlockByHashOrNumber(ctx, req.(*GetBlockByHashOrNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockscoutGateway_GetTransactionByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockscoutGatewayServer).GetTransactionByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ankr.BlockscoutGateway/GetTransactionByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockscoutGatewayServer).GetTransactionByHash(ctx, req.(*GetTransactionByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockscoutGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.ankr.BlockscoutGateway",
	HandlerType: (*BlockscoutGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecentBlocks",
			Handler:    _BlockscoutGateway_GetRecentBlocks_Handler,
		},
		{
			MethodName: "GetRecentTxs",
			Handler:    _BlockscoutGateway_GetRecentTxs_Handler,
		},
		{
			MethodName: "GetAddress",
			Handler:    _BlockscoutGateway_GetAddress_Handler,
		},
		{
			MethodName: "GetBlockByHashOrNumber",
			Handler:    _BlockscoutGateway_GetBlockByHashOrNumber_Handler,
		},
		{
			MethodName: "GetTransactionByHash",
			Handler:    _BlockscoutGateway_GetTransactionByHash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}
