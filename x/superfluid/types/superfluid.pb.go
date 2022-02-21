// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/superfluid/superfluid.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type SuperfluidAssetType int32

const (
	SuperfluidAssetTypeNative  SuperfluidAssetType = 0
	SuperfluidAssetTypeLPShare SuperfluidAssetType = 1
)

var SuperfluidAssetType_name = map[int32]string{
	0: "SuperfluidAssetTypeNative",
	1: "SuperfluidAssetTypeLPShare",
}

var SuperfluidAssetType_value = map[string]int32{
	"SuperfluidAssetTypeNative":  0,
	"SuperfluidAssetTypeLPShare": 1,
}

func (x SuperfluidAssetType) String() string {
	return proto.EnumName(SuperfluidAssetType_name, int32(x))
}

func (SuperfluidAssetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_79d3c29d82dbb734, []int{0}
}

// SuperfluidAsset stores the pair of superfluid asset type and denom pair
type SuperfluidAsset struct {
	Denom     string              `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	AssetType SuperfluidAssetType `protobuf:"varint,2,opt,name=asset_type,json=assetType,proto3,enum=osmosis.superfluid.SuperfluidAssetType" json:"asset_type,omitempty"`
}

func (m *SuperfluidAsset) Reset()         { *m = SuperfluidAsset{} }
func (m *SuperfluidAsset) String() string { return proto.CompactTextString(m) }
func (*SuperfluidAsset) ProtoMessage()    {}
func (*SuperfluidAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_79d3c29d82dbb734, []int{0}
}
func (m *SuperfluidAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SuperfluidAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SuperfluidAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SuperfluidAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuperfluidAsset.Merge(m, src)
}
func (m *SuperfluidAsset) XXX_Size() int {
	return m.Size()
}
func (m *SuperfluidAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_SuperfluidAsset.DiscardUnknown(m)
}

var xxx_messageInfo_SuperfluidAsset proto.InternalMessageInfo

// SuperfluidIntermediaryAccount takes the role of intermediary between LP token
// and OSMO tokens for superfluid staking
type SuperfluidIntermediaryAccount struct {
	Denom   string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	ValAddr string `protobuf:"bytes,2,opt,name=val_addr,json=valAddr,proto3" json:"val_addr,omitempty"`
	GaugeId uint64 `protobuf:"varint,3,opt,name=gauge_id,json=gaugeId,proto3" json:"gauge_id,omitempty"`
}

func (m *SuperfluidIntermediaryAccount) Reset()         { *m = SuperfluidIntermediaryAccount{} }
func (m *SuperfluidIntermediaryAccount) String() string { return proto.CompactTextString(m) }
func (*SuperfluidIntermediaryAccount) ProtoMessage()    {}
func (*SuperfluidIntermediaryAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_79d3c29d82dbb734, []int{1}
}
func (m *SuperfluidIntermediaryAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SuperfluidIntermediaryAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SuperfluidIntermediaryAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SuperfluidIntermediaryAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuperfluidIntermediaryAccount.Merge(m, src)
}
func (m *SuperfluidIntermediaryAccount) XXX_Size() int {
	return m.Size()
}
func (m *SuperfluidIntermediaryAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_SuperfluidIntermediaryAccount.DiscardUnknown(m)
}

var xxx_messageInfo_SuperfluidIntermediaryAccount proto.InternalMessageInfo

func (m *SuperfluidIntermediaryAccount) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *SuperfluidIntermediaryAccount) GetValAddr() string {
	if m != nil {
		return m.ValAddr
	}
	return ""
}

func (m *SuperfluidIntermediaryAccount) GetGaugeId() uint64 {
	if m != nil {
		return m.GaugeId
	}
	return 0
}

// The Osmo-Equivalent-Multiplier Record for epoch N refers to the osmo worth we
// treat an LP share as having, for all of epoch N. Eventually this is intended
// to be set as the Time-weighted-average-osmo-backing for the entire duration
// of epoch N-1. (Thereby locking whats in use for epoch N as based on the prior
// epochs rewards) However for now, this is not the TWAP but instead the spot
// price at the boundary.  For different types of assets in the future, it could
// change.
type OsmoEquivalentMultiplierRecord struct {
	EpochNumber int64                                  `protobuf:"varint,1,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
	Denom       string                                 `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Multiplier  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=multiplier,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"multiplier" yaml:"multiplier"`
}

func (m *OsmoEquivalentMultiplierRecord) Reset()         { *m = OsmoEquivalentMultiplierRecord{} }
func (m *OsmoEquivalentMultiplierRecord) String() string { return proto.CompactTextString(m) }
func (*OsmoEquivalentMultiplierRecord) ProtoMessage()    {}
func (*OsmoEquivalentMultiplierRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_79d3c29d82dbb734, []int{2}
}
func (m *OsmoEquivalentMultiplierRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OsmoEquivalentMultiplierRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OsmoEquivalentMultiplierRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OsmoEquivalentMultiplierRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OsmoEquivalentMultiplierRecord.Merge(m, src)
}
func (m *OsmoEquivalentMultiplierRecord) XXX_Size() int {
	return m.Size()
}
func (m *OsmoEquivalentMultiplierRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_OsmoEquivalentMultiplierRecord.DiscardUnknown(m)
}

var xxx_messageInfo_OsmoEquivalentMultiplierRecord proto.InternalMessageInfo

func (m *OsmoEquivalentMultiplierRecord) GetEpochNumber() int64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

func (m *OsmoEquivalentMultiplierRecord) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

// SuperfluidDelegationRecord takes the role of intermediary between LP token
// and OSMO tokens for superfluid staking
type SuperfluidDelegationRecord struct {
	DelegatorAddress string     `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	ValidatorAddress string     `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	DelegationAmount types.Coin `protobuf:"bytes,3,opt,name=delegation_amount,json=delegationAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"delegation_amount"`
}

func (m *SuperfluidDelegationRecord) Reset()         { *m = SuperfluidDelegationRecord{} }
func (m *SuperfluidDelegationRecord) String() string { return proto.CompactTextString(m) }
func (*SuperfluidDelegationRecord) ProtoMessage()    {}
func (*SuperfluidDelegationRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_79d3c29d82dbb734, []int{3}
}
func (m *SuperfluidDelegationRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SuperfluidDelegationRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SuperfluidDelegationRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SuperfluidDelegationRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuperfluidDelegationRecord.Merge(m, src)
}
func (m *SuperfluidDelegationRecord) XXX_Size() int {
	return m.Size()
}
func (m *SuperfluidDelegationRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_SuperfluidDelegationRecord.DiscardUnknown(m)
}

var xxx_messageInfo_SuperfluidDelegationRecord proto.InternalMessageInfo

func (m *SuperfluidDelegationRecord) GetDelegatorAddress() string {
	if m != nil {
		return m.DelegatorAddress
	}
	return ""
}

func (m *SuperfluidDelegationRecord) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *SuperfluidDelegationRecord) GetDelegationAmount() types.Coin {
	if m != nil {
		return m.DelegationAmount
	}
	return types.Coin{}
}

func init() {
	proto.RegisterEnum("osmosis.superfluid.SuperfluidAssetType", SuperfluidAssetType_name, SuperfluidAssetType_value)
	proto.RegisterType((*SuperfluidAsset)(nil), "osmosis.superfluid.SuperfluidAsset")
	proto.RegisterType((*SuperfluidIntermediaryAccount)(nil), "osmosis.superfluid.SuperfluidIntermediaryAccount")
	proto.RegisterType((*OsmoEquivalentMultiplierRecord)(nil), "osmosis.superfluid.OsmoEquivalentMultiplierRecord")
	proto.RegisterType((*SuperfluidDelegationRecord)(nil), "osmosis.superfluid.SuperfluidDelegationRecord")
}

func init() {
	proto.RegisterFile("osmosis/superfluid/superfluid.proto", fileDescriptor_79d3c29d82dbb734)
}

var fileDescriptor_79d3c29d82dbb734 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x4f, 0xd4, 0x4e,
	0x14, 0x6f, 0x81, 0xff, 0x1f, 0x76, 0x30, 0xba, 0x54, 0x0e, 0xb0, 0x09, 0xed, 0x5a, 0x13, 0xd9,
	0x68, 0xe8, 0x04, 0x8c, 0x31, 0xe1, 0xb6, 0x0b, 0x9a, 0x90, 0x28, 0x98, 0xe2, 0x89, 0xcb, 0x66,
	0xda, 0x19, 0xba, 0x13, 0xda, 0x4e, 0x9d, 0x99, 0x56, 0xf7, 0x1b, 0x70, 0xf4, 0x23, 0x90, 0x78,
	0xf3, 0x43, 0x78, 0xe6, 0xc8, 0xd1, 0x78, 0x58, 0xcd, 0xee, 0xc5, 0x9b, 0x89, 0x9f, 0xc0, 0x74,
	0xda, 0xed, 0x56, 0xc4, 0xc4, 0x53, 0xe7, 0xbd, 0xdf, 0xeb, 0xef, 0xfd, 0x7e, 0x6f, 0xde, 0x80,
	0xfb, 0x4c, 0x44, 0x4c, 0x50, 0x01, 0x45, 0x9a, 0x10, 0x7e, 0x1a, 0xa6, 0x14, 0xd7, 0x8e, 0x4e,
	0xc2, 0x99, 0x64, 0x86, 0x51, 0x16, 0x39, 0x33, 0xa4, 0xb5, 0x1a, 0xb0, 0x80, 0x29, 0x18, 0xe6,
	0xa7, 0xa2, 0xb2, 0x65, 0x06, 0x8c, 0x05, 0x21, 0x81, 0x2a, 0xf2, 0xd2, 0x53, 0x88, 0x53, 0x8e,
	0x24, 0x65, 0x71, 0x89, 0x5b, 0xd7, 0x71, 0x49, 0x23, 0x22, 0x24, 0x8a, 0x92, 0x29, 0x81, 0xaf,
	0x7a, 0x41, 0x0f, 0x09, 0x02, 0xb3, 0x6d, 0x8f, 0x48, 0xb4, 0x0d, 0x7d, 0x46, 0x4b, 0x02, 0x7b,
	0x08, 0xee, 0x1c, 0x57, 0x22, 0xba, 0x42, 0x10, 0x69, 0xac, 0x82, 0xff, 0x30, 0x89, 0x59, 0xb4,
	0xa6, 0xb7, 0xf5, 0x4e, 0xc3, 0x2d, 0x02, 0xe3, 0x39, 0x00, 0x28, 0x87, 0xfb, 0x72, 0x98, 0x90,
	0xb5, 0xb9, 0xb6, 0xde, 0xb9, 0xbd, 0xb3, 0xe9, 0xfc, 0x69, 0xc4, 0xb9, 0x46, 0xf7, 0x7a, 0x98,
	0x10, 0xb7, 0x81, 0xa6, 0xc7, 0xdd, 0xa5, 0xf3, 0x0b, 0x4b, 0xfb, 0x7e, 0x61, 0xe9, 0xf6, 0x19,
	0xd8, 0x98, 0xd5, 0x1e, 0xc4, 0x92, 0xf0, 0x88, 0x60, 0x8a, 0xf8, 0xb0, 0xeb, 0xfb, 0x2c, 0x8d,
	0xff, 0x26, 0x64, 0x1d, 0x2c, 0x65, 0x28, 0xec, 0x23, 0x8c, 0xb9, 0x92, 0xd1, 0x70, 0x17, 0x33,
	0x14, 0x76, 0x31, 0xe6, 0x39, 0x14, 0xa0, 0x34, 0x20, 0x7d, 0x8a, 0xd7, 0xe6, 0xdb, 0x7a, 0x67,
	0xc1, 0x5d, 0x54, 0xf1, 0x01, 0xb6, 0x3f, 0xe9, 0xc0, 0x3c, 0x12, 0x11, 0x7b, 0xf6, 0x26, 0xa5,
	0x19, 0x0a, 0x49, 0x2c, 0x5f, 0xa6, 0xa1, 0xa4, 0x49, 0x48, 0x09, 0x77, 0x89, 0xcf, 0x38, 0x36,
	0xee, 0x81, 0x5b, 0x24, 0x61, 0xfe, 0xa0, 0x1f, 0xa7, 0x91, 0x47, 0xb8, 0xea, 0x3a, 0xef, 0x2e,
	0xab, 0xdc, 0xa1, 0x4a, 0xcd, 0x14, 0xcd, 0xd5, 0x15, 0xf9, 0x00, 0x44, 0x15, 0x99, 0x6a, 0xdc,
	0xe8, 0xed, 0x5d, 0x8e, 0x2c, 0xed, 0xcb, 0xc8, 0x7a, 0x10, 0x50, 0x39, 0x48, 0x3d, 0xc7, 0x67,
	0x11, 0x2c, 0xaf, 0xa2, 0xf8, 0x6c, 0x09, 0x7c, 0x06, 0xf3, 0x59, 0x0a, 0x67, 0x9f, 0xf8, 0x3f,
	0x47, 0xd6, 0xca, 0x10, 0x45, 0xe1, 0xae, 0x3d, 0x63, 0xb2, 0xdd, 0x1a, 0xad, 0xfd, 0x43, 0x07,
	0xad, 0xd9, 0xb8, 0xf6, 0x49, 0x48, 0x02, 0xb5, 0x08, 0xa5, 0xf8, 0x47, 0x60, 0x05, 0x17, 0x39,
	0xc6, 0xd5, 0x6c, 0x88, 0x10, 0xe5, 0xdc, 0x9a, 0x15, 0xd0, 0x2d, 0xf2, 0x79, 0x71, 0x86, 0x42,
	0x8a, 0x7f, 0x2b, 0x2e, 0x2c, 0x35, 0x2b, 0x60, 0x5a, 0xfc, 0xb6, 0x62, 0xa6, 0x2c, 0xee, 0xa3,
	0x28, 0xbf, 0x1a, 0x65, 0x72, 0x79, 0x67, 0xdd, 0x29, 0xbc, 0x38, 0xf9, 0x76, 0x39, 0xe5, 0x76,
	0x39, 0x7b, 0x8c, 0xc6, 0x3d, 0x98, 0xfb, 0xff, 0xf8, 0xd5, 0xda, 0xfc, 0x07, 0xff, 0xf9, 0x0f,
	0x95, 0x4a, 0xca, 0xe2, 0xae, 0xea, 0xf1, 0xf0, 0x04, 0xdc, 0xbd, 0x61, 0x97, 0x8c, 0x0d, 0xb0,
	0x7e, 0x43, 0xfa, 0x10, 0x49, 0x9a, 0x91, 0xa6, 0x66, 0x98, 0xf5, 0x31, 0x55, 0xf0, 0x8b, 0x57,
	0xc7, 0x03, 0xc4, 0x49, 0x53, 0x6f, 0x2d, 0x9c, 0x7f, 0x30, 0xb5, 0xde, 0xd1, 0xe5, 0xd8, 0xd4,
	0xaf, 0xc6, 0xa6, 0xfe, 0x6d, 0x6c, 0xea, 0xef, 0x27, 0xa6, 0x76, 0x35, 0x31, 0xb5, 0xcf, 0x13,
	0x53, 0x3b, 0x79, 0x52, 0x13, 0x5c, 0x6e, 0xf7, 0x56, 0x88, 0x3c, 0x31, 0x0d, 0x60, 0xf6, 0x14,
	0xbe, 0xab, 0xbf, 0x6e, 0xe5, 0xc1, 0xfb, 0x5f, 0x3d, 0xa7, 0xc7, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xc4, 0xb7, 0x3f, 0xbd, 0x00, 0x04, 0x00, 0x00,
}

func (this *SuperfluidAsset) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SuperfluidAsset)
	if !ok {
		that2, ok := that.(SuperfluidAsset)
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
	if this.Denom != that1.Denom {
		return false
	}
	if this.AssetType != that1.AssetType {
		return false
	}
	return true
}
func (m *SuperfluidAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SuperfluidAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SuperfluidAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AssetType != 0 {
		i = encodeVarintSuperfluid(dAtA, i, uint64(m.AssetType))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintSuperfluid(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SuperfluidIntermediaryAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SuperfluidIntermediaryAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SuperfluidIntermediaryAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GaugeId != 0 {
		i = encodeVarintSuperfluid(dAtA, i, uint64(m.GaugeId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ValAddr) > 0 {
		i -= len(m.ValAddr)
		copy(dAtA[i:], m.ValAddr)
		i = encodeVarintSuperfluid(dAtA, i, uint64(len(m.ValAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintSuperfluid(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OsmoEquivalentMultiplierRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OsmoEquivalentMultiplierRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OsmoEquivalentMultiplierRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Multiplier.Size()
		i -= size
		if _, err := m.Multiplier.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSuperfluid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintSuperfluid(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if m.EpochNumber != 0 {
		i = encodeVarintSuperfluid(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SuperfluidDelegationRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SuperfluidDelegationRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SuperfluidDelegationRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DelegationAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSuperfluid(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintSuperfluid(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintSuperfluid(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSuperfluid(dAtA []byte, offset int, v uint64) int {
	offset -= sovSuperfluid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SuperfluidAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovSuperfluid(uint64(l))
	}
	if m.AssetType != 0 {
		n += 1 + sovSuperfluid(uint64(m.AssetType))
	}
	return n
}

func (m *SuperfluidIntermediaryAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovSuperfluid(uint64(l))
	}
	l = len(m.ValAddr)
	if l > 0 {
		n += 1 + l + sovSuperfluid(uint64(l))
	}
	if m.GaugeId != 0 {
		n += 1 + sovSuperfluid(uint64(m.GaugeId))
	}
	return n
}

func (m *OsmoEquivalentMultiplierRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochNumber != 0 {
		n += 1 + sovSuperfluid(uint64(m.EpochNumber))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovSuperfluid(uint64(l))
	}
	l = m.Multiplier.Size()
	n += 1 + l + sovSuperfluid(uint64(l))
	return n
}

func (m *SuperfluidDelegationRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovSuperfluid(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovSuperfluid(uint64(l))
	}
	l = m.DelegationAmount.Size()
	n += 1 + l + sovSuperfluid(uint64(l))
	return n
}

func sovSuperfluid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSuperfluid(x uint64) (n int) {
	return sovSuperfluid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SuperfluidAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSuperfluid
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
			return fmt.Errorf("proto: SuperfluidAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuperfluidAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetType", wireType)
			}
			m.AssetType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetType |= SuperfluidAssetType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSuperfluid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSuperfluid
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
func (m *SuperfluidIntermediaryAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSuperfluid
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
			return fmt.Errorf("proto: SuperfluidIntermediaryAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuperfluidIntermediaryAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GaugeId", wireType)
			}
			m.GaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSuperfluid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSuperfluid
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
func (m *OsmoEquivalentMultiplierRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSuperfluid
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
			return fmt.Errorf("proto: OsmoEquivalentMultiplierRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OsmoEquivalentMultiplierRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Multiplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSuperfluid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSuperfluid
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
func (m *SuperfluidDelegationRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSuperfluid
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
			return fmt.Errorf("proto: SuperfluidDelegationRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SuperfluidDelegationRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSuperfluid
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
				return ErrInvalidLengthSuperfluid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSuperfluid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DelegationAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSuperfluid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSuperfluid
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
func skipSuperfluid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSuperfluid
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
					return 0, ErrIntOverflowSuperfluid
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
					return 0, ErrIntOverflowSuperfluid
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
				return 0, ErrInvalidLengthSuperfluid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSuperfluid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSuperfluid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSuperfluid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSuperfluid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSuperfluid = fmt.Errorf("proto: unexpected end of group")
)