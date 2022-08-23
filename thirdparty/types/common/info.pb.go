// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/v1/info.proto

// Package common defines commonly used types agnostic to the node role on the Sonr network.

package common

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

type EntityKind int32

const (
	EntityKind_ADDRESS EntityKind = 0
	EntityKind_DID     EntityKind = 1
	EntityKind_CID     EntityKind = 2
)

var EntityKind_name = map[int32]string{
	0: "ADDRESS",
	1: "DID",
	2: "CID",
}

var EntityKind_value = map[string]int32{
	"ADDRESS": 0,
	"DID":     1,
	"CID":     2,
}

func (x EntityKind) String() string {
	return proto.EnumName(EntityKind_name, int32(x))
}

func (EntityKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{0}
}

type BlockchainModule int32

const (
	// Query x/registry module
	BlockchainModule_REGISTRY BlockchainModule = 0
	// Query x/schema module
	BlockchainModule_SCHEMA BlockchainModule = 1
	// Query x/bucket module
	BlockchainModule_BUCKET BlockchainModule = 2
)

var BlockchainModule_name = map[int32]string{
	0: "REGISTRY",
	1: "SCHEMA",
	2: "BUCKET",
}

var BlockchainModule_value = map[string]int32{
	"REGISTRY": 0,
	"SCHEMA":   1,
	"BUCKET":   2,
}

func (x BlockchainModule) String() string {
	return proto.EnumName(BlockchainModule_name, int32(x))
}

func (BlockchainModule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{1}
}

// Peers Active Status
type Peer_Status int32

const (
	Peer_STATUS_UNSPECIFIED Peer_Status = 0
	Peer_STATUS_ONLINE      Peer_Status = 1
	Peer_STATUS_AWAY        Peer_Status = 2
	Peer_STATUS_BUSY        Peer_Status = 3
)

var Peer_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "STATUS_ONLINE",
	2: "STATUS_AWAY",
	3: "STATUS_BUSY",
}

var Peer_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"STATUS_ONLINE":      1,
	"STATUS_AWAY":        2,
	"STATUS_BUSY":        3,
}

func (x Peer_Status) String() string {
	return proto.EnumName(Peer_Status_name, int32(x))
}

func (Peer_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{0, 0}
}

// Basic Info Sent to Peers to Establish Connections
type Peer struct {
	PeerId string      `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Did    string      `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	Status Peer_Status `protobuf:"varint,3,opt,name=status,proto3,enum=sonrio.common.v1.Peer_Status" json:"status,omitempty"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6ffb5b3e6a498f4, []int{0}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return m.Size()
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *Peer) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *Peer) GetStatus() Peer_Status {
	if m != nil {
		return m.Status
	}
	return Peer_STATUS_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("sonrio.common.v1.EntityKind", EntityKind_name, EntityKind_value)
	proto.RegisterEnum("sonrio.common.v1.BlockchainModule", BlockchainModule_name, BlockchainModule_value)
	proto.RegisterEnum("sonrio.common.v1.Peer_Status", Peer_Status_name, Peer_Status_value)
	proto.RegisterType((*Peer)(nil), "sonrio.common.v1.Peer")
}

func init() { proto.RegisterFile("common/v1/info.proto", fileDescriptor_a6ffb5b3e6a498f4) }

var fileDescriptor_a6ffb5b3e6a498f4 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x8b, 0xda, 0x40,
	0x18, 0xc6, 0x33, 0x49, 0x89, 0xed, 0x6b, 0xff, 0x4c, 0x87, 0xd2, 0x7a, 0x69, 0x10, 0x4f, 0x62,
	0x69, 0x82, 0x2d, 0x85, 0x5e, 0xf3, 0x67, 0xda, 0x0e, 0x56, 0x2b, 0x99, 0x84, 0xe2, 0x5e, 0x44,
	0x4d, 0x76, 0x1d, 0x56, 0x33, 0x21, 0x19, 0x05, 0xbf, 0xc5, 0x7e, 0xa3, 0xbd, 0xee, 0xd1, 0xe3,
	0x1e, 0x17, 0xfd, 0x22, 0x4b, 0xa2, 0x87, 0x65, 0x4f, 0xf3, 0xcc, 0xef, 0x7d, 0x78, 0xe0, 0x7d,
	0x5e, 0xf8, 0xb0, 0x90, 0xeb, 0xb5, 0xcc, 0x9c, 0x6d, 0xdf, 0x11, 0xd9, 0xa5, 0xb4, 0xf3, 0x42,
	0x2a, 0x49, 0x70, 0x29, 0xb3, 0x42, 0x48, 0xfb, 0x34, 0xb4, 0xb7, 0xfd, 0xce, 0x2d, 0x82, 0x17,
	0xe3, 0x34, 0x2d, 0xc8, 0x27, 0x68, 0xe4, 0x69, 0x5a, 0x4c, 0x45, 0xd2, 0x42, 0x6d, 0xd4, 0x7d,
	0x15, 0x9a, 0xd5, 0x97, 0x25, 0x04, 0x83, 0x91, 0x88, 0xa4, 0xa5, 0xd7, 0xb0, 0x92, 0xe4, 0x07,
	0x98, 0xa5, 0x9a, 0xa9, 0x4d, 0xd9, 0x32, 0xda, 0xa8, 0xfb, 0xf6, 0xdb, 0x67, 0xfb, 0x79, 0xac,
	0x5d, 0x45, 0xda, 0xbc, 0x36, 0x85, 0x67, 0x73, 0x27, 0x06, 0xf3, 0x44, 0xc8, 0x47, 0x20, 0x3c,
	0x72, 0xa3, 0x98, 0x4f, 0xe3, 0x11, 0x1f, 0x53, 0x9f, 0xfd, 0x62, 0x34, 0xc0, 0x1a, 0x79, 0x0f,
	0x6f, 0xce, 0xfc, 0xdf, 0xe8, 0x2f, 0x1b, 0x51, 0x8c, 0xc8, 0x3b, 0x68, 0x9e, 0x91, 0xfb, 0xdf,
	0x9d, 0x60, 0xfd, 0x09, 0xf0, 0x62, 0x3e, 0xc1, 0x46, 0xef, 0x0b, 0x00, 0xcd, 0x94, 0x50, 0xbb,
	0x81, 0xc8, 0x12, 0xd2, 0x84, 0x86, 0x1b, 0x04, 0x21, 0xe5, 0x1c, 0x6b, 0xa4, 0x01, 0x46, 0xc0,
	0x02, 0x8c, 0x2a, 0xe1, 0xb3, 0x00, 0xeb, 0xbd, 0x9f, 0x80, 0xbd, 0x95, 0x5c, 0x5c, 0x2f, 0x96,
	0x33, 0x91, 0x0d, 0x65, 0xb2, 0x59, 0xa5, 0xe4, 0x35, 0xbc, 0x0c, 0xe9, 0x6f, 0xc6, 0xa3, 0x70,
	0x82, 0x35, 0x02, 0x60, 0x72, 0xff, 0x0f, 0x1d, 0xba, 0x18, 0x55, 0xda, 0x8b, 0xfd, 0x01, 0x8d,
	0xb0, 0xee, 0xb1, 0xbb, 0x83, 0x85, 0xf6, 0x07, 0x0b, 0x3d, 0x1c, 0x2c, 0x74, 0x73, 0xb4, 0xb4,
	0xfd, 0xd1, 0xd2, 0xee, 0x8f, 0x96, 0x76, 0xe1, 0x5c, 0x09, 0xb5, 0xdc, 0xcc, 0xab, 0xed, 0x9d,
	0xaa, 0x88, 0xaf, 0x42, 0xd6, 0xaf, 0xa3, 0x96, 0xa2, 0x48, 0xf2, 0x59, 0xa1, 0x76, 0x8e, 0xda,
	0xe5, 0x69, 0xe9, 0x9c, 0xfa, 0x99, 0x9b, 0xf5, 0x31, 0xbe, 0x3f, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x21, 0x14, 0x6d, 0x8c, 0xa4, 0x01, 0x00, 0x00,
}

func (m *Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Peer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintInfo(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PeerId) > 0 {
		i -= len(m.PeerId)
		copy(dAtA[i:], m.PeerId)
		i = encodeVarintInfo(dAtA, i, uint64(len(m.PeerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Peer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovInfo(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovInfo(uint64(m.Status))
	}
	return n
}

func sovInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInfo(x uint64) (n int) {
	return sovInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInfo
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
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
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
				return ErrInvalidLengthInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Peer_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInfo
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
func skipInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
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
					return 0, ErrIntOverflowInfo
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
				return 0, ErrInvalidLengthInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInfo = fmt.Errorf("proto: unexpected end of group")
)
