// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/v1/discover.proto

// Package Discover is used to find other Peers in the sonr network.

package service

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	common "github.com/sonr-io/sonr/third_party/types/common"
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

// LobbyMessage is message passed from Peer in Lobby
type LobbyMessage struct {
	From      *common.Peer `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Body      []byte       `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Signature []byte       `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	CreatedAt int64        `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (m *LobbyMessage) Reset()         { *m = LobbyMessage{} }
func (m *LobbyMessage) String() string { return proto.CompactTextString(m) }
func (*LobbyMessage) ProtoMessage()    {}
func (*LobbyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_57395aadf09e8ea7, []int{0}
}
func (m *LobbyMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LobbyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LobbyMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LobbyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LobbyMessage.Merge(m, src)
}
func (m *LobbyMessage) XXX_Size() int {
	return m.Size()
}
func (m *LobbyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LobbyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LobbyMessage proto.InternalMessageInfo

func (m *LobbyMessage) GetFrom() *common.Peer {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *LobbyMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *LobbyMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *LobbyMessage) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type RefreshEvent struct {
	TopicName  string         `protobuf:"bytes,1,opt,name=topic_name,json=topicName,proto3" json:"topic_name,omitempty"`
	Peers      []*common.Peer `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	ReceivedAt int64          `protobuf:"varint,3,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
}

func (m *RefreshEvent) Reset()         { *m = RefreshEvent{} }
func (m *RefreshEvent) String() string { return proto.CompactTextString(m) }
func (*RefreshEvent) ProtoMessage()    {}
func (*RefreshEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_57395aadf09e8ea7, []int{1}
}
func (m *RefreshEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RefreshEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RefreshEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RefreshEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshEvent.Merge(m, src)
}
func (m *RefreshEvent) XXX_Size() int {
	return m.Size()
}
func (m *RefreshEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshEvent proto.InternalMessageInfo

func (m *RefreshEvent) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *RefreshEvent) GetPeers() []*common.Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *RefreshEvent) GetReceivedAt() int64 {
	if m != nil {
		return m.ReceivedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*LobbyMessage)(nil), "sonrio.motor.service.v1.LobbyMessage")
	proto.RegisterType((*RefreshEvent)(nil), "sonrio.motor.service.v1.RefreshEvent")
}

func init() { proto.RegisterFile("service/v1/discover.proto", fileDescriptor_57395aadf09e8ea7) }

var fileDescriptor_57395aadf09e8ea7 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbd, 0x4e, 0xeb, 0x40,
	0x10, 0x85, 0xb3, 0x71, 0xee, 0x95, 0xbc, 0x49, 0xb5, 0xba, 0x3f, 0x06, 0x81, 0xb1, 0x52, 0x45,
	0x08, 0xd6, 0x32, 0x3c, 0x41, 0x90, 0xe8, 0x00, 0x45, 0x2e, 0x69, 0x22, 0xff, 0x4c, 0x92, 0x2d,
	0xec, 0xb1, 0x66, 0x37, 0x96, 0x2c, 0xf1, 0x06, 0x34, 0x3c, 0x16, 0x65, 0x4a, 0x4a, 0x94, 0xbc,
	0x08, 0xf2, 0xda, 0x11, 0x1d, 0xd5, 0xae, 0xce, 0x7c, 0x67, 0xce, 0xe8, 0xf0, 0x13, 0x0d, 0x54,
	0xab, 0x0c, 0xc2, 0x3a, 0x0a, 0x73, 0xa5, 0x33, 0xac, 0x81, 0x64, 0x45, 0x68, 0x50, 0xfc, 0xd7,
	0x58, 0x92, 0x42, 0x59, 0xa0, 0x41, 0x92, 0x3d, 0x27, 0xeb, 0xe8, 0xf4, 0x4f, 0x86, 0x45, 0x81,
	0x65, 0x6b, 0x51, 0xe5, 0x0a, 0x3b, 0x7c, 0xfa, 0xca, 0xf8, 0xe4, 0x01, 0xd3, 0xb4, 0x79, 0x04,
	0xad, 0x93, 0x35, 0x88, 0x4b, 0x3e, 0x5a, 0x11, 0x16, 0x1e, 0x0b, 0xd8, 0x6c, 0x7c, 0xf3, 0x4f,
	0xf6, 0xeb, 0x3a, 0xb3, 0xac, 0x23, 0xb9, 0x00, 0xa0, 0xd8, 0x32, 0x42, 0xf0, 0x51, 0x8a, 0x79,
	0xe3, 0x0d, 0x03, 0x36, 0x9b, 0xc4, 0xf6, 0x2f, 0xce, 0xb8, 0xab, 0xd5, 0xba, 0x4c, 0xcc, 0x96,
	0xc0, 0x73, 0xec, 0xe0, 0x5b, 0x10, 0xe7, 0x9c, 0x67, 0x04, 0x89, 0x81, 0x7c, 0x99, 0x18, 0x6f,
	0x14, 0xb0, 0x99, 0x13, 0xbb, 0xbd, 0x32, 0x37, 0xd3, 0x17, 0x3e, 0x89, 0x61, 0x45, 0xa0, 0x37,
	0xf7, 0x35, 0x94, 0xa6, 0xc5, 0x0d, 0x56, 0x2a, 0x5b, 0x96, 0x49, 0x01, 0xf6, 0x24, 0x37, 0x76,
	0xad, 0xf2, 0x94, 0x14, 0x20, 0xae, 0xf8, 0xaf, 0x0a, 0x80, 0xb4, 0x37, 0x0c, 0x9c, 0x1f, 0x8e,
	0xed, 0x20, 0x71, 0xc1, 0xc7, 0x04, 0x19, 0xa8, 0xba, 0x0b, 0x77, 0x6c, 0x38, 0x3f, 0x4a, 0x73,
	0x73, 0x97, 0xbf, 0xef, 0x7d, 0xb6, 0xdb, 0xfb, 0xec, 0x73, 0xef, 0xb3, 0xb7, 0x83, 0x3f, 0xd8,
	0x1d, 0xfc, 0xc1, 0xc7, 0xc1, 0x1f, 0xf0, 0xbf, 0x0a, 0xed, 0x6e, 0x69, 0x9a, 0x0a, 0xf4, 0xb1,
	0xd5, 0x05, 0x7b, 0x8e, 0xd6, 0xca, 0x6c, 0xb6, 0x69, 0x9b, 0x18, 0xb6, 0xc0, 0xb5, 0x42, 0xfb,
	0x86, 0x66, 0xa3, 0x28, 0x5f, 0x56, 0x09, 0x99, 0x26, 0xb4, 0xa6, 0xb0, 0x37, 0xa5, 0xbf, 0x6d,
	0xf1, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xe3, 0x24, 0xd0, 0xc4, 0x01, 0x00, 0x00,
}

func (m *LobbyMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LobbyMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LobbyMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintDiscover(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintDiscover(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintDiscover(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x12
	}
	if m.From != nil {
		{
			size, err := m.From.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDiscover(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RefreshEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RefreshEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RefreshEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReceivedAt != 0 {
		i = encodeVarintDiscover(dAtA, i, uint64(m.ReceivedAt))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Peers) > 0 {
		for iNdEx := len(m.Peers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Peers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDiscover(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.TopicName) > 0 {
		i -= len(m.TopicName)
		copy(dAtA[i:], m.TopicName)
		i = encodeVarintDiscover(dAtA, i, uint64(len(m.TopicName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDiscover(dAtA []byte, offset int, v uint64) int {
	offset -= sovDiscover(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LobbyMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.From != nil {
		l = m.From.Size()
		n += 1 + l + sovDiscover(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovDiscover(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovDiscover(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovDiscover(uint64(m.CreatedAt))
	}
	return n
}

func (m *RefreshEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TopicName)
	if l > 0 {
		n += 1 + l + sovDiscover(uint64(l))
	}
	if len(m.Peers) > 0 {
		for _, e := range m.Peers {
			l = e.Size()
			n += 1 + l + sovDiscover(uint64(l))
		}
	}
	if m.ReceivedAt != 0 {
		n += 1 + sovDiscover(uint64(m.ReceivedAt))
	}
	return n
}

func sovDiscover(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDiscover(x uint64) (n int) {
	return sovDiscover(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LobbyMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscover
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
			return fmt.Errorf("proto: LobbyMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LobbyMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscover
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
				return ErrInvalidLengthDiscover
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDiscover
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.From == nil {
				m.From = &common.Peer{}
			}
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscover
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
				return ErrInvalidLengthDiscover
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscover
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscover
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
				return ErrInvalidLengthDiscover
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscover
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscover
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDiscover(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDiscover
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
func (m *RefreshEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscover
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
			return fmt.Errorf("proto: RefreshEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RefreshEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopicName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscover
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
				return ErrInvalidLengthDiscover
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscover
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TopicName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscover
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
				return ErrInvalidLengthDiscover
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDiscover
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Peers = append(m.Peers, &common.Peer{})
			if err := m.Peers[len(m.Peers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedAt", wireType)
			}
			m.ReceivedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscover
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReceivedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDiscover(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDiscover
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
func skipDiscover(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDiscover
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
					return 0, ErrIntOverflowDiscover
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
					return 0, ErrIntOverflowDiscover
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
				return 0, ErrInvalidLengthDiscover
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDiscover
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDiscover
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDiscover        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDiscover          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDiscover = fmt.Errorf("proto: unexpected end of group")
)
