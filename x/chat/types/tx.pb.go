// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
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

type MsgRegister struct {
	FromAddress          string     `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	NodeAddress          string     `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	MortgageAmount       types.Coin `protobuf:"bytes,3,opt,name=mortgage_amount,json=mortgageAmount,proto3" json:"mortgage_amount" yaml:"mortgage_amount"`
	MobilePrefix         string     `protobuf:"bytes,4,opt,name=mobile_prefix,json=mobilePrefix,proto3" json:"mobile_prefix,omitempty" yaml:"mobile_prefix"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MsgRegister) Reset()         { *m = MsgRegister{} }
func (m *MsgRegister) String() string { return proto.CompactTextString(m) }
func (*MsgRegister) ProtoMessage()    {}
func (*MsgRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{0}
}
func (m *MsgRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgRegister.Unmarshal(m, b)
}
func (m *MsgRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgRegister.Marshal(b, m, deterministic)
}
func (m *MsgRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegister.Merge(m, src)
}
func (m *MsgRegister) XXX_Size() int {
	return xxx_messageInfo_MsgRegister.Size(m)
}
func (m *MsgRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegister.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegister proto.InternalMessageInfo

func (m *MsgRegister) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgRegister) GetNodeAddress() string {
	if m != nil {
		return m.NodeAddress
	}
	return ""
}

func (m *MsgRegister) GetMortgageAmount() types.Coin {
	if m != nil {
		return m.MortgageAmount
	}
	return types.Coin{}
}

func (m *MsgRegister) GetMobilePrefix() string {
	if m != nil {
		return m.MobilePrefix
	}
	return ""
}

type MsgSetChatFee struct {
	FromAddress          string     `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	Fee                  types.Coin `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee" yaml:"fee"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MsgSetChatFee) Reset()         { *m = MsgSetChatFee{} }
func (m *MsgSetChatFee) String() string { return proto.CompactTextString(m) }
func (*MsgSetChatFee) ProtoMessage()    {}
func (*MsgSetChatFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{1}
}
func (m *MsgSetChatFee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSetChatFee.Unmarshal(m, b)
}
func (m *MsgSetChatFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSetChatFee.Marshal(b, m, deterministic)
}
func (m *MsgSetChatFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetChatFee.Merge(m, src)
}
func (m *MsgSetChatFee) XXX_Size() int {
	return xxx_messageInfo_MsgSetChatFee.Size(m)
}
func (m *MsgSetChatFee) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetChatFee.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetChatFee proto.InternalMessageInfo

func (m *MsgSetChatFee) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgSetChatFee) GetFee() types.Coin {
	if m != nil {
		return m.Fee
	}
	return types.Coin{}
}

type MsgSendGift struct {
	FromAddress          string     `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	NodeAddress          string     `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	ToAddress            string     `protobuf:"bytes,3,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty" yaml:"to_address"`
	GiftId               int64      `protobuf:"varint,4,opt,name=gift_id,json=giftId,proto3" json:"gift_id,omitempty" yaml:"gift_id"`
	GiftAmount           int64      `protobuf:"varint,5,opt,name=gift_amount,json=giftAmount,proto3" json:"gift_amount,omitempty" yaml:"gift_amount"`
	GiftValue            types.Coin `protobuf:"bytes,6,opt,name=gift_value,json=giftValue,proto3" json:"gift_value" yaml:"gift_value"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MsgSendGift) Reset()         { *m = MsgSendGift{} }
func (m *MsgSendGift) String() string { return proto.CompactTextString(m) }
func (*MsgSendGift) ProtoMessage()    {}
func (*MsgSendGift) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{2}
}
func (m *MsgSendGift) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSendGift.Unmarshal(m, b)
}
func (m *MsgSendGift) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSendGift.Marshal(b, m, deterministic)
}
func (m *MsgSendGift) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendGift.Merge(m, src)
}
func (m *MsgSendGift) XXX_Size() int {
	return xxx_messageInfo_MsgSendGift.Size(m)
}
func (m *MsgSendGift) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendGift.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendGift proto.InternalMessageInfo

func (m *MsgSendGift) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgSendGift) GetNodeAddress() string {
	if m != nil {
		return m.NodeAddress
	}
	return ""
}

func (m *MsgSendGift) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *MsgSendGift) GetGiftId() int64 {
	if m != nil {
		return m.GiftId
	}
	return 0
}

func (m *MsgSendGift) GetGiftAmount() int64 {
	if m != nil {
		return m.GiftAmount
	}
	return 0
}

func (m *MsgSendGift) GetGiftValue() types.Coin {
	if m != nil {
		return m.GiftValue
	}
	return types.Coin{}
}

type MsgAddressBookSave struct {
	FromAddress          string   `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	AddressBook          []string `protobuf:"bytes,2,rep,name=address_book,json=addressBook,proto3" json:"address_book,omitempty" yaml:"address_book"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgAddressBookSave) Reset()         { *m = MsgAddressBookSave{} }
func (m *MsgAddressBookSave) String() string { return proto.CompactTextString(m) }
func (*MsgAddressBookSave) ProtoMessage()    {}
func (*MsgAddressBookSave) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{3}
}
func (m *MsgAddressBookSave) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgAddressBookSave.Unmarshal(m, b)
}
func (m *MsgAddressBookSave) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgAddressBookSave.Marshal(b, m, deterministic)
}
func (m *MsgAddressBookSave) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddressBookSave.Merge(m, src)
}
func (m *MsgAddressBookSave) XXX_Size() int {
	return xxx_messageInfo_MsgAddressBookSave.Size(m)
}
func (m *MsgAddressBookSave) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddressBookSave.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddressBookSave proto.InternalMessageInfo

func (m *MsgAddressBookSave) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgAddressBookSave) GetAddressBook() []string {
	if m != nil {
		return m.AddressBook
	}
	return nil
}

type MsgMobileTransfer struct {
	FromAddress          string   `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	ToAddress            string   `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty" yaml:"to_address"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty" yaml:"mobile"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgMobileTransfer) Reset()         { *m = MsgMobileTransfer{} }
func (m *MsgMobileTransfer) String() string { return proto.CompactTextString(m) }
func (*MsgMobileTransfer) ProtoMessage()    {}
func (*MsgMobileTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{4}
}
func (m *MsgMobileTransfer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgMobileTransfer.Unmarshal(m, b)
}
func (m *MsgMobileTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgMobileTransfer.Marshal(b, m, deterministic)
}
func (m *MsgMobileTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMobileTransfer.Merge(m, src)
}
func (m *MsgMobileTransfer) XXX_Size() int {
	return xxx_messageInfo_MsgMobileTransfer.Size(m)
}
func (m *MsgMobileTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMobileTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMobileTransfer proto.InternalMessageInfo

func (m *MsgMobileTransfer) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgMobileTransfer) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *MsgMobileTransfer) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type MsgChangeGateway struct {
	FromAddress          string   `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	Gateway              string   `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty" yaml:"gateway"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgChangeGateway) Reset()         { *m = MsgChangeGateway{} }
func (m *MsgChangeGateway) String() string { return proto.CompactTextString(m) }
func (*MsgChangeGateway) ProtoMessage()    {}
func (*MsgChangeGateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{5}
}
func (m *MsgChangeGateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgChangeGateway.Unmarshal(m, b)
}
func (m *MsgChangeGateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgChangeGateway.Marshal(b, m, deterministic)
}
func (m *MsgChangeGateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgChangeGateway.Merge(m, src)
}
func (m *MsgChangeGateway) XXX_Size() int {
	return xxx_messageInfo_MsgChangeGateway.Size(m)
}
func (m *MsgChangeGateway) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgChangeGateway.DiscardUnknown(m)
}

var xxx_messageInfo_MsgChangeGateway proto.InternalMessageInfo

func (m *MsgChangeGateway) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgChangeGateway) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

type MsgBurnGetMobile struct {
	FromAddress          string   `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	MobilePrefix         string   `protobuf:"bytes,2,opt,name=mobile_prefix,json=mobilePrefix,proto3" json:"mobile_prefix,omitempty" yaml:"mobile_prefix"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgBurnGetMobile) Reset()         { *m = MsgBurnGetMobile{} }
func (m *MsgBurnGetMobile) String() string { return proto.CompactTextString(m) }
func (*MsgBurnGetMobile) ProtoMessage()    {}
func (*MsgBurnGetMobile) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{6}
}
func (m *MsgBurnGetMobile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgBurnGetMobile.Unmarshal(m, b)
}
func (m *MsgBurnGetMobile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgBurnGetMobile.Marshal(b, m, deterministic)
}
func (m *MsgBurnGetMobile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnGetMobile.Merge(m, src)
}
func (m *MsgBurnGetMobile) XXX_Size() int {
	return xxx_messageInfo_MsgBurnGetMobile.Size(m)
}
func (m *MsgBurnGetMobile) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnGetMobile.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnGetMobile proto.InternalMessageInfo

func (m *MsgBurnGetMobile) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgBurnGetMobile) GetMobilePrefix() string {
	if m != nil {
		return m.MobilePrefix
	}
	return ""
}

type MsgSetChatInfo struct {
	FromAddress          string     `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	NodeAddress          string     `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	AddressBook          []string   `protobuf:"bytes,3,rep,name=address_book,json=addressBook,proto3" json:"address_book,omitempty" yaml:"address_book"`
	ChatBlacklist        []string   `protobuf:"bytes,4,rep,name=chat_blacklist,json=chatBlacklist,proto3" json:"chat_blacklist,omitempty" yaml:"chat_blacklist"`
	ChatRestrictedMode   string     `protobuf:"bytes,5,opt,name=chat_restricted_mode,json=chatRestrictedMode,proto3" json:"chat_restricted_mode,omitempty" yaml:"chat_limit"`
	ChatFee              types.Coin `protobuf:"bytes,6,opt,name=chat_fee,json=chatFee,proto3" json:"chat_fee" yaml:"chat_fee"`
	ChatWhitelist        []string   `protobuf:"bytes,7,rep,name=chat_whitelist,json=chatWhitelist,proto3" json:"chat_whitelist,omitempty" yaml:"chat_whitelist"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MsgSetChatInfo) Reset()         { *m = MsgSetChatInfo{} }
func (m *MsgSetChatInfo) String() string { return proto.CompactTextString(m) }
func (*MsgSetChatInfo) ProtoMessage()    {}
func (*MsgSetChatInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{7}
}
func (m *MsgSetChatInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSetChatInfo.Unmarshal(m, b)
}
func (m *MsgSetChatInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSetChatInfo.Marshal(b, m, deterministic)
}
func (m *MsgSetChatInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetChatInfo.Merge(m, src)
}
func (m *MsgSetChatInfo) XXX_Size() int {
	return xxx_messageInfo_MsgSetChatInfo.Size(m)
}
func (m *MsgSetChatInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetChatInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetChatInfo proto.InternalMessageInfo

func (m *MsgSetChatInfo) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgSetChatInfo) GetNodeAddress() string {
	if m != nil {
		return m.NodeAddress
	}
	return ""
}

func (m *MsgSetChatInfo) GetAddressBook() []string {
	if m != nil {
		return m.AddressBook
	}
	return nil
}

func (m *MsgSetChatInfo) GetChatBlacklist() []string {
	if m != nil {
		return m.ChatBlacklist
	}
	return nil
}

func (m *MsgSetChatInfo) GetChatRestrictedMode() string {
	if m != nil {
		return m.ChatRestrictedMode
	}
	return ""
}

func (m *MsgSetChatInfo) GetChatFee() types.Coin {
	if m != nil {
		return m.ChatFee
	}
	return types.Coin{}
}

func (m *MsgSetChatInfo) GetChatWhitelist() []string {
	if m != nil {
		return m.ChatWhitelist
	}
	return nil
}

// MsgConvertCoinResponse returns no fieldsyou
type MsgEmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgEmptyResponse) Reset()         { *m = MsgEmptyResponse{} }
func (m *MsgEmptyResponse) String() string { return proto.CompactTextString(m) }
func (*MsgEmptyResponse) ProtoMessage()    {}
func (*MsgEmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{8}
}
func (m *MsgEmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgEmptyResponse.Unmarshal(m, b)
}
func (m *MsgEmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgEmptyResponse.Marshal(b, m, deterministic)
}
func (m *MsgEmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEmptyResponse.Merge(m, src)
}
func (m *MsgEmptyResponse) XXX_Size() int {
	return xxx_messageInfo_MsgEmptyResponse.Size(m)
}
func (m *MsgEmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegister)(nil), "freemasonry.chat.v1.MsgRegister")
	proto.RegisterType((*MsgSetChatFee)(nil), "freemasonry.chat.v1.MsgSetChatFee")
	proto.RegisterType((*MsgSendGift)(nil), "freemasonry.chat.v1.MsgSendGift")
	proto.RegisterType((*MsgAddressBookSave)(nil), "freemasonry.chat.v1.MsgAddressBookSave")
	proto.RegisterType((*MsgMobileTransfer)(nil), "freemasonry.chat.v1.MsgMobileTransfer")
	proto.RegisterType((*MsgChangeGateway)(nil), "freemasonry.chat.v1.MsgChangeGateway")
	proto.RegisterType((*MsgBurnGetMobile)(nil), "freemasonry.chat.v1.MsgBurnGetMobile")
	proto.RegisterType((*MsgSetChatInfo)(nil), "freemasonry.chat.v1.MsgSetChatInfo")
	proto.RegisterType((*MsgEmptyResponse)(nil), "freemasonry.chat.v1.MsgEmptyResponse")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor_0fd2153dc07d3b5c) }

var fileDescriptor_0fd2153dc07d3b5c = []byte{
	// 871 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x26, 0x49, 0xbb, 0x3f, 0x27, 0x4d, 0x96, 0x75, 0xb7, 0x6d, 0x76, 0x85, 0xea, 0x68, 0x10,
	0x65, 0x11, 0xc8, 0xd6, 0x16, 0x24, 0xa4, 0x4a, 0x08, 0x9a, 0x15, 0xac, 0x7a, 0x61, 0x09, 0x39,
	0x88, 0x0a, 0x24, 0x64, 0x4d, 0xec, 0x13, 0x67, 0xb4, 0xb1, 0x27, 0xf2, 0xcc, 0x6e, 0x37, 0x12,
	0x8f, 0x50, 0xde, 0x83, 0x4b, 0x1e, 0x83, 0xa7, 0xb0, 0xb8, 0xe1, 0x05, 0x72, 0xc1, 0x35, 0xf2,
	0xcc, 0x24, 0xb1, 0x43, 0xc3, 0xa6, 0xca, 0x45, 0xef, 0xe6, 0xcc, 0xf9, 0xbe, 0x6f, 0x7c, 0xe6,
	0xfc, 0x8c, 0x61, 0x4f, 0xde, 0x38, 0x93, 0x8c, 0x4b, 0x6e, 0xdd, 0x1f, 0x66, 0x88, 0x09, 0x15,
	0x3c, 0xcd, 0xa6, 0x4e, 0x38, 0xa2, 0xd2, 0xb9, 0x3e, 0x3b, 0xf9, 0x20, 0xe6, 0x3c, 0x1e, 0xa3,
	0x4b, 0x27, 0xcc, 0xa5, 0x69, 0xca, 0x25, 0x95, 0x8c, 0xa7, 0x42, 0x53, 0x4e, 0x8e, 0x62, 0x1e,
	0x73, 0xb5, 0x74, 0x8b, 0x95, 0xd9, 0x7d, 0x1c, 0x72, 0x91, 0x70, 0xe1, 0x0e, 0xa8, 0x40, 0xf7,
	0xfa, 0x6c, 0x80, 0x92, 0x9e, 0xb9, 0x21, 0x67, 0xa9, 0xf6, 0x93, 0xdf, 0xeb, 0xd0, 0xf4, 0x44,
	0xec, 0x63, 0xcc, 0x84, 0xc4, 0xcc, 0x7a, 0x06, 0xf7, 0x86, 0x19, 0x4f, 0x02, 0x1a, 0x45, 0x19,
	0x0a, 0xd1, 0xa9, 0x75, 0x6b, 0xa7, 0xfb, 0xbd, 0x47, 0xb3, 0xdc, 0xbe, 0x3f, 0xa5, 0xc9, 0xf8,
	0x19, 0x29, 0x7b, 0x89, 0xdf, 0x2c, 0xcc, 0xe7, 0xda, 0x2a, 0xb8, 0x29, 0x8f, 0x70, 0xc1, 0xad,
	0xaf, 0x72, 0xcb, 0x5e, 0xe2, 0x37, 0x0b, 0x73, 0xce, 0x1d, 0xc0, 0x41, 0xc2, 0x33, 0x19, 0xd3,
	0x18, 0x03, 0x9a, 0xf0, 0xab, 0x54, 0x76, 0x1a, 0xdd, 0xda, 0x69, 0xf3, 0xe9, 0xb1, 0xa3, 0x23,
	0x70, 0x8a, 0x08, 0x1c, 0x13, 0x81, 0x73, 0xce, 0x59, 0xda, 0x7b, 0xfc, 0x67, 0x6e, 0xbf, 0x37,
	0xcb, 0xed, 0x87, 0x5a, 0x7d, 0x85, 0x4f, 0xfc, 0xf6, 0x7c, 0xe7, 0xb9, 0xda, 0xb0, 0xbe, 0x82,
	0x56, 0xc2, 0x07, 0x6c, 0x8c, 0xc1, 0x24, 0xc3, 0x21, 0xbb, 0xe9, 0xdc, 0x51, 0x1f, 0xd8, 0x99,
	0xe5, 0xf6, 0xd1, 0x5c, 0xa2, 0xe4, 0x26, 0xfe, 0x3d, 0x6d, 0x7f, 0xaf, 0xcd, 0xd7, 0x35, 0x68,
	0x79, 0x22, 0xee, 0xa3, 0x3c, 0x1f, 0x51, 0xf9, 0x1d, 0xe2, 0x56, 0x97, 0xf5, 0x35, 0x34, 0x86,
	0x88, 0xea, 0x8e, 0xfe, 0x37, 0x48, 0xcb, 0x04, 0x09, 0x46, 0x11, 0x91, 0xf8, 0x05, 0x93, 0xfc,
	0xa3, 0x33, 0xd7, 0xc7, 0x34, 0xba, 0x60, 0x43, 0xf9, 0xce, 0x32, 0xf7, 0x05, 0x80, 0xe4, 0x0b,
	0x66, 0x43, 0x31, 0x1f, 0xcc, 0x72, 0xfb, 0x50, 0x33, 0x97, 0x3e, 0xe2, 0xef, 0x4b, 0x3e, 0x67,
	0x7d, 0x0a, 0xbb, 0x31, 0x1b, 0xca, 0x80, 0x45, 0x2a, 0x0b, 0x8d, 0x9e, 0x35, 0xcb, 0xed, 0xb6,
	0xa6, 0x18, 0x07, 0xf1, 0x77, 0x8a, 0xd5, 0x8b, 0xc8, 0xfa, 0x12, 0x9a, 0x6a, 0xcf, 0x14, 0xc6,
	0x5d, 0x45, 0x78, 0x38, 0xcb, 0x6d, 0xab, 0x44, 0x98, 0x67, 0x1d, 0x0a, 0xcb, 0x64, 0xbc, 0x0f,
	0xca, 0x0a, 0xae, 0xe9, 0xf8, 0x0a, 0x3b, 0x3b, 0xb7, 0xdd, 0xf5, 0xb1, 0xb9, 0xeb, 0xc3, 0x92,
	0xac, 0xa2, 0x12, 0x7f, 0xbf, 0x30, 0x7e, 0x54, 0xeb, 0xd7, 0x35, 0xb0, 0x3c, 0x11, 0x9b, 0x48,
	0x7a, 0x9c, 0x5f, 0xf6, 0xe9, 0x35, 0x6e, 0x7b, 0xff, 0xc6, 0x11, 0x0c, 0x38, 0xbf, 0xec, 0xd4,
	0xbb, 0x8d, 0x2a, 0xb7, 0xec, 0x25, 0x7e, 0x93, 0x2e, 0xcf, 0x26, 0x7f, 0xd4, 0xe0, 0xd0, 0x13,
	0xb1, 0xa7, 0x4a, 0xf5, 0x87, 0x8c, 0xa6, 0x62, 0xb8, 0x65, 0x1f, 0x57, 0x33, 0x5a, 0xdf, 0x30,
	0xa3, 0x9f, 0xc0, 0x8e, 0x6e, 0x17, 0x53, 0x03, 0x87, 0xb3, 0xdc, 0x6e, 0x95, 0xdb, 0x8a, 0xf8,
	0x06, 0x40, 0x7e, 0x85, 0xf7, 0x3d, 0x11, 0x9f, 0x8f, 0x68, 0x1a, 0xe3, 0x05, 0x95, 0xf8, 0x8a,
	0x4e, 0xb7, 0xfa, 0xe0, 0xcf, 0x60, 0x37, 0xd6, 0x32, 0xe6, 0x6b, 0xcb, 0xc5, 0xa4, 0x1d, 0xc4,
	0x9f, 0x43, 0xc8, 0x6f, 0x35, 0x75, 0x7c, 0xef, 0x2a, 0x4b, 0x2f, 0x50, 0xea, 0x7b, 0xdb, 0xea,
	0xf8, 0xff, 0xcc, 0x95, 0xfa, 0x5b, 0xcd, 0x95, 0xbf, 0x1a, 0xd0, 0x5e, 0xce, 0x95, 0x17, 0xe9,
	0x90, 0xbf, 0xb3, 0x5e, 0x5e, 0xad, 0xc3, 0xc6, 0xe6, 0x75, 0x68, 0x7d, 0x03, 0xed, 0xe2, 0xa1,
	0x0a, 0x06, 0x63, 0x1a, 0x5e, 0x8e, 0x99, 0x90, 0x9d, 0x3b, 0x8a, 0x7d, 0x3c, 0xcb, 0xed, 0x07,
	0x9a, 0x5d, 0xf5, 0x13, 0xbf, 0x55, 0x6c, 0xf4, 0xe6, 0xb6, 0x75, 0x01, 0x47, 0x0a, 0x91, 0xa1,
	0x90, 0x19, 0x0b, 0x25, 0x46, 0x41, 0xc2, 0x23, 0x54, 0xfd, 0x5e, 0xa9, 0x40, 0x85, 0x1a, 0xb3,
	0x84, 0x49, 0xe2, 0x5b, 0x85, 0xe1, 0x2f, 0x18, 0x1e, 0x8f, 0xd0, 0xf2, 0x60, 0x4f, 0x41, 0x8a,
	0x01, 0x7b, 0x6b, 0xd3, 0x3f, 0x32, 0x4d, 0x7f, 0x50, 0xd2, 0x56, 0x53, 0x76, 0x37, 0x34, 0x63,
	0x7e, 0x1e, 0xd9, 0xab, 0x11, 0x93, 0xa8, 0x22, 0xdb, 0x7d, 0x63, 0x64, 0x0b, 0xbf, 0x89, 0xec,
	0xe5, 0xc2, 0xb6, 0x54, 0xc5, 0x7d, 0x9b, 0x4c, 0xe4, 0xd4, 0x47, 0x31, 0xe1, 0xa9, 0xc0, 0xa7,
	0x7f, 0xdf, 0x85, 0x86, 0x27, 0x62, 0xab, 0x0f, 0x7b, 0x8b, 0xd7, 0xb7, 0xeb, 0xbc, 0xe1, 0xdd,
	0x77, 0x4a, 0xef, 0xf3, 0xc9, 0x47, 0xeb, 0x10, 0x15, 0x71, 0xeb, 0x25, 0x40, 0xe9, 0x9d, 0x22,
	0xeb, 0x48, 0x4b, 0xcc, 0xa6, 0xc2, 0x7d, 0xd8, 0x5b, 0xbc, 0x38, 0xdd, 0xf5, 0xb2, 0x1a, 0xb1,
	0xa9, 0x28, 0x85, 0x83, 0xd5, 0x69, 0xfa, 0xf1, 0x3a, 0xe6, 0x0a, 0x70, 0xd3, 0x23, 0x02, 0x68,
	0xaf, 0x4c, 0xc8, 0x27, 0xeb, 0x88, 0x55, 0xdc, 0xa6, 0x07, 0xfc, 0x02, 0xad, 0xea, 0x40, 0x5b,
	0xcb, 0xab, 0xc0, 0xde, 0x42, 0xbe, 0x3a, 0xb0, 0xd6, 0xf2, 0x2a, 0xb0, 0x4d, 0xe5, 0x7f, 0x82,
	0x66, 0x79, 0xfe, 0x7c, 0x78, 0x4b, 0xc1, 0x14, 0xa0, 0x0d, 0xa5, 0x7b, 0xa7, 0x3f, 0x3f, 0xa9,
	0xe0, 0x42, 0x77, 0x30, 0xe6, 0xe1, 0x65, 0x38, 0xa2, 0x2c, 0x75, 0x6f, 0xdc, 0x82, 0xe7, 0xca,
	0xe9, 0x04, 0xc5, 0x60, 0x47, 0xfd, 0x92, 0x7e, 0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c,
	0x70, 0xce, 0x90, 0x07, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Register(ctx context.Context, in *MsgRegister, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	SetChatFee(ctx context.Context, in *MsgSetChatFee, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	SendGift(ctx context.Context, in *MsgSendGift, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	AddressBookSave(ctx context.Context, in *MsgAddressBookSave, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	MobileTransfer(ctx context.Context, in *MsgMobileTransfer, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	ChangeGateway(ctx context.Context, in *MsgChangeGateway, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	BurnGetMobile(ctx context.Context, in *MsgBurnGetMobile, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	SetChatInfo(ctx context.Context, in *MsgSetChatInfo, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
}

type msgClient struct {
	cc *grpc.ClientConn
}

func NewMsgClient(cc *grpc.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Register(ctx context.Context, in *MsgRegister, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetChatFee(ctx context.Context, in *MsgSetChatFee, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/SetChatFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendGift(ctx context.Context, in *MsgSendGift, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/SendGift", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddressBookSave(ctx context.Context, in *MsgAddressBookSave, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/AddressBookSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MobileTransfer(ctx context.Context, in *MsgMobileTransfer, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/MobileTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ChangeGateway(ctx context.Context, in *MsgChangeGateway, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/ChangeGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BurnGetMobile(ctx context.Context, in *MsgBurnGetMobile, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/BurnGetMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetChatInfo(ctx context.Context, in *MsgSetChatInfo, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/SetChatInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Register(context.Context, *MsgRegister) (*MsgEmptyResponse, error)
	SetChatFee(context.Context, *MsgSetChatFee) (*MsgEmptyResponse, error)
	SendGift(context.Context, *MsgSendGift) (*MsgEmptyResponse, error)
	AddressBookSave(context.Context, *MsgAddressBookSave) (*MsgEmptyResponse, error)
	MobileTransfer(context.Context, *MsgMobileTransfer) (*MsgEmptyResponse, error)
	ChangeGateway(context.Context, *MsgChangeGateway) (*MsgEmptyResponse, error)
	BurnGetMobile(context.Context, *MsgBurnGetMobile) (*MsgEmptyResponse, error)
	SetChatInfo(context.Context, *MsgSetChatInfo) (*MsgEmptyResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Register(ctx context.Context, req *MsgRegister) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedMsgServer) SetChatFee(ctx context.Context, req *MsgSetChatFee) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetChatFee not implemented")
}
func (*UnimplementedMsgServer) SendGift(ctx context.Context, req *MsgSendGift) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGift not implemented")
}
func (*UnimplementedMsgServer) AddressBookSave(ctx context.Context, req *MsgAddressBookSave) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressBookSave not implemented")
}
func (*UnimplementedMsgServer) MobileTransfer(ctx context.Context, req *MsgMobileTransfer) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MobileTransfer not implemented")
}
func (*UnimplementedMsgServer) ChangeGateway(ctx context.Context, req *MsgChangeGateway) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeGateway not implemented")
}
func (*UnimplementedMsgServer) BurnGetMobile(ctx context.Context, req *MsgBurnGetMobile) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnGetMobile not implemented")
}
func (*UnimplementedMsgServer) SetChatInfo(ctx context.Context, req *MsgSetChatInfo) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetChatInfo not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Register(ctx, req.(*MsgRegister))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetChatFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetChatFee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetChatFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/SetChatFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetChatFee(ctx, req.(*MsgSetChatFee))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendGift_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendGift)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendGift(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/SendGift",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendGift(ctx, req.(*MsgSendGift))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddressBookSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddressBookSave)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddressBookSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/AddressBookSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddressBookSave(ctx, req.(*MsgAddressBookSave))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MobileTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMobileTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MobileTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/MobileTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MobileTransfer(ctx, req.(*MsgMobileTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ChangeGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgChangeGateway)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ChangeGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/ChangeGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ChangeGateway(ctx, req.(*MsgChangeGateway))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BurnGetMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurnGetMobile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BurnGetMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/BurnGetMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BurnGetMobile(ctx, req.(*MsgBurnGetMobile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetChatInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetChatInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetChatInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/SetChatInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetChatInfo(ctx, req.(*MsgSetChatInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "freemasonry.chat.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Msg_Register_Handler,
		},
		{
			MethodName: "SetChatFee",
			Handler:    _Msg_SetChatFee_Handler,
		},
		{
			MethodName: "SendGift",
			Handler:    _Msg_SendGift_Handler,
		},
		{
			MethodName: "AddressBookSave",
			Handler:    _Msg_AddressBookSave_Handler,
		},
		{
			MethodName: "MobileTransfer",
			Handler:    _Msg_MobileTransfer_Handler,
		},
		{
			MethodName: "ChangeGateway",
			Handler:    _Msg_ChangeGateway_Handler,
		},
		{
			MethodName: "BurnGetMobile",
			Handler:    _Msg_BurnGetMobile_Handler,
		},
		{
			MethodName: "SetChatInfo",
			Handler:    _Msg_SetChatInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tx.proto",
}