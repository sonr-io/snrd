// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: did/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the module genesis state
type GenesisState struct {
	// Params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_fda181cae44f7c00, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// Document defines a DID document
type Document struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Controller           string   `protobuf:"bytes,2,opt,name=controller,proto3" json:"controller,omitempty"`
	Authentication       []string `protobuf:"bytes,3,rep,name=authentication,proto3" json:"authentication,omitempty"`
	AssertionMethod      []string `protobuf:"bytes,4,rep,name=assertion_method,json=assertionMethod,proto3" json:"assertion_method,omitempty"`
	CapabilityDelegation []string `protobuf:"bytes,5,rep,name=capability_delegation,json=capabilityDelegation,proto3" json:"capability_delegation,omitempty"`
	CapabilityInvocation []string `protobuf:"bytes,6,rep,name=capability_invocation,json=capabilityInvocation,proto3" json:"capability_invocation,omitempty"`
	Service              []string `protobuf:"bytes,7,rep,name=service,proto3" json:"service,omitempty"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_fda181cae44f7c00, []int{1}
}
func (m *Document) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Document.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return m.Size()
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Document) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *Document) GetAuthentication() []string {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *Document) GetAssertionMethod() []string {
	if m != nil {
		return m.AssertionMethod
	}
	return nil
}

func (m *Document) GetCapabilityDelegation() []string {
	if m != nil {
		return m.CapabilityDelegation
	}
	return nil
}

func (m *Document) GetCapabilityInvocation() []string {
	if m != nil {
		return m.CapabilityInvocation
	}
	return nil
}

func (m *Document) GetService() []string {
	if m != nil {
		return m.Service
	}
	return nil
}

// Params defines the set of module parameters.
type Params struct {
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_fda181cae44f7c00, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "did.v1.GenesisState")
	proto.RegisterType((*Document)(nil), "did.v1.Document")
	proto.RegisterType((*Params)(nil), "did.v1.Params")
}

func init() { proto.RegisterFile("did/v1/genesis.proto", fileDescriptor_fda181cae44f7c00) }

var fileDescriptor_fda181cae44f7c00 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0xdb, 0xc2, 0x2d, 0x97, 0xb9, 0x37, 0xdc, 0xeb, 0x04, 0x63, 0xc3, 0xa2, 0x60, 0x17,
	0x06, 0x8d, 0x69, 0x83, 0xec, 0x0c, 0x2b, 0x42, 0x62, 0x5c, 0x98, 0x18, 0xdc, 0xb9, 0x21, 0x43,
	0x67, 0x52, 0x26, 0x69, 0x67, 0x9a, 0xce, 0xd0, 0xc8, 0x2b, 0xb8, 0xd2, 0x9d, 0x4b, 0x1e, 0xc1,
	0xc7, 0x60, 0xc9, 0xd2, 0x95, 0x31, 0xb0, 0xd0, 0xc7, 0x30, 0x9d, 0x41, 0x34, 0x6c, 0x26, 0xe7,
	0x7c, 0xff, 0xff, 0x9f, 0x4c, 0xce, 0x01, 0x75, 0x4c, 0x71, 0x90, 0x77, 0x82, 0x88, 0x30, 0x22,
	0xa8, 0xf0, 0xd3, 0x8c, 0x4b, 0x0e, 0x6d, 0x4c, 0xb1, 0x9f, 0x77, 0x1a, 0x7b, 0x28, 0xa1, 0x8c,
	0x07, 0xea, 0xd5, 0x52, 0xa3, 0x1e, 0xf1, 0x88, 0xab, 0x32, 0x28, 0x2a, 0x4d, 0xbd, 0x1e, 0xf8,
	0x7b, 0xa1, 0x27, 0xdc, 0x48, 0x24, 0x09, 0x3c, 0x05, 0x76, 0x8a, 0x32, 0x94, 0x08, 0xc7, 0x6c,
	0x99, 0xed, 0x3f, 0x67, 0x35, 0x5f, 0x4f, 0xf4, 0xaf, 0x15, 0xed, 0x97, 0x17, 0xaf, 0x4d, 0x63,
	0xb8, 0xf1, 0x78, 0x8f, 0x16, 0xf8, 0x3d, 0xe0, 0xe1, 0x34, 0x21, 0x4c, 0xc2, 0x1a, 0xb0, 0x28,
	0x56, 0xb1, 0xea, 0xd0, 0xa2, 0x18, 0xba, 0x00, 0x84, 0x9c, 0xc9, 0x8c, 0xc7, 0x31, 0xc9, 0x1c,
	0x4b, 0xf1, 0x1f, 0x04, 0x1e, 0x81, 0x1a, 0x9a, 0xca, 0x09, 0x61, 0x92, 0x86, 0x48, 0x52, 0xce,
	0x9c, 0x52, 0xab, 0xd4, 0xae, 0x0e, 0x77, 0x28, 0x3c, 0x06, 0xff, 0x91, 0x10, 0x24, 0x2b, 0x9a,
	0x51, 0x42, 0xe4, 0x84, 0x63, 0xa7, 0xac, 0x9c, 0xff, 0xb6, 0xfc, 0x4a, 0x61, 0xd8, 0x05, 0xfb,
	0x21, 0x4a, 0xd1, 0x98, 0xc6, 0x54, 0xce, 0x46, 0x98, 0xc4, 0x24, 0xd2, 0x93, 0x7f, 0x29, 0x7f,
	0xfd, 0x5b, 0x1c, 0x6c, 0xb5, 0x9d, 0x10, 0x65, 0x39, 0xdf, 0x7c, 0xc7, 0xde, 0x0d, 0x5d, 0x6e,
	0x35, 0xe8, 0x80, 0x8a, 0x20, 0x59, 0x4e, 0x43, 0xe2, 0x54, 0x94, 0xed, 0xab, 0xf5, 0x0e, 0x81,
	0xad, 0x77, 0x75, 0x7e, 0xf0, 0x34, 0x6f, 0x1a, 0x1f, 0xf3, 0xa6, 0x79, 0xff, 0xfe, 0x7c, 0x02,
	0x8a, 0x7b, 0xe9, 0xb5, 0xf5, 0x7b, 0x8b, 0x95, 0x6b, 0x2e, 0x57, 0xae, 0xf9, 0xb6, 0x72, 0xcd,
	0x87, 0xb5, 0x6b, 0x2c, 0xd7, 0xae, 0xf1, 0xb2, 0x76, 0x8d, 0x5b, 0x2f, 0xa2, 0x72, 0x32, 0x1d,
	0xfb, 0x21, 0x4f, 0x02, 0xce, 0x04, 0x67, 0x59, 0xa0, 0x9e, 0xbb, 0xa0, 0x88, 0xcb, 0x59, 0x4a,
	0xc4, 0xd8, 0x56, 0x97, 0xeb, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x43, 0x9f, 0x66, 0x22, 0x02,
	0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	return true
}
func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Document) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Document) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Document) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Service) > 0 {
		for iNdEx := len(m.Service) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Service[iNdEx])
			copy(dAtA[i:], m.Service[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.Service[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.CapabilityInvocation) > 0 {
		for iNdEx := len(m.CapabilityInvocation) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CapabilityInvocation[iNdEx])
			copy(dAtA[i:], m.CapabilityInvocation[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.CapabilityInvocation[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.CapabilityDelegation) > 0 {
		for iNdEx := len(m.CapabilityDelegation) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CapabilityDelegation[iNdEx])
			copy(dAtA[i:], m.CapabilityDelegation[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.CapabilityDelegation[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.AssertionMethod) > 0 {
		for iNdEx := len(m.AssertionMethod) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AssertionMethod[iNdEx])
			copy(dAtA[i:], m.AssertionMethod[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.AssertionMethod[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Authentication) > 0 {
		for iNdEx := len(m.Authentication) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Authentication[iNdEx])
			copy(dAtA[i:], m.Authentication[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.Authentication[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Document) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Authentication) > 0 {
		for _, s := range m.Authentication {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AssertionMethod) > 0 {
		for _, s := range m.AssertionMethod {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CapabilityDelegation) > 0 {
		for _, s := range m.CapabilityDelegation {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CapabilityInvocation) > 0 {
		for _, s := range m.CapabilityInvocation {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Service) > 0 {
		for _, s := range m.Service {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Document) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Document: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Document: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authentication", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authentication = append(m.Authentication, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssertionMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssertionMethod = append(m.AssertionMethod, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapabilityDelegation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CapabilityDelegation = append(m.CapabilityDelegation, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapabilityInvocation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CapabilityInvocation = append(m.CapabilityInvocation, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = append(m.Service, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
