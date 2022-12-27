// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: motor/bind/v1/request.proto

// Package Motor is used for defining a Motor node and its properties.

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	common "github.com/sonr-hq/sonr/pkg/common"
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

// -----------------------------------------------------------------------------
// Motor Node API
// -----------------------------------------------------------------------------
type NetworkMode int32

const (
	NetworkMode_NetworkMode_LOCAL NetworkMode = 0
	NetworkMode_NetworkMode_ALPHA NetworkMode = 1
	NetworkMode_NetworkMode_BETA  NetworkMode = 2
	NetworkMode_NetworkMode_TEST  NetworkMode = 3
)

var NetworkMode_name = map[int32]string{
	0: "NetworkMode_LOCAL",
	1: "NetworkMode_ALPHA",
	2: "NetworkMode_BETA",
	3: "NetworkMode_TEST",
}

var NetworkMode_value = map[string]int32{
	"NetworkMode_LOCAL": 0,
	"NetworkMode_ALPHA": 1,
	"NetworkMode_BETA":  2,
	"NetworkMode_TEST":  3,
}

func (x NetworkMode) String() string {
	return proto.EnumName(NetworkMode_name, int32(x))
}

func (NetworkMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_726257aebe87753d, []int{0}
}

// (Client) InitializeRequest Message to Establish Sonr Host/API/Room
type InitializeRequest struct {
	// Identifier of this Device
	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Device Home Directory
	HomeDir string `protobuf:"bytes,2,opt,name=home_dir,json=homeDir,proto3" json:"home_dir,omitempty"`
	// Device Library Support Directory
	SupportDir string `protobuf:"bytes,3,opt,name=support_dir,json=supportDir,proto3" json:"support_dir,omitempty"`
	// Device Temporary Storage Directory
	TempDir string `protobuf:"bytes,4,opt,name=temp_dir,json=tempDir,proto3" json:"temp_dir,omitempty"`
	// Logging level for the session, can be configured after setting
	// (info|debug|warn|error|fatal)
	LogLevel string `protobuf:"bytes,5,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	// Client Mode
	Network NetworkMode `protobuf:"varint,6,opt,name=network,proto3,enum=sonrhq.motor.bind.v1.NetworkMode" json:"network,omitempty"`
}

func (m *InitializeRequest) Reset()         { *m = InitializeRequest{} }
func (m *InitializeRequest) String() string { return proto.CompactTextString(m) }
func (*InitializeRequest) ProtoMessage()    {}
func (*InitializeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_726257aebe87753d, []int{0}
}
func (m *InitializeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InitializeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InitializeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InitializeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitializeRequest.Merge(m, src)
}
func (m *InitializeRequest) XXX_Size() int {
	return m.Size()
}
func (m *InitializeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitializeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitializeRequest proto.InternalMessageInfo

func (m *InitializeRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *InitializeRequest) GetHomeDir() string {
	if m != nil {
		return m.HomeDir
	}
	return ""
}

func (m *InitializeRequest) GetSupportDir() string {
	if m != nil {
		return m.SupportDir
	}
	return ""
}

func (m *InitializeRequest) GetTempDir() string {
	if m != nil {
		return m.TempDir
	}
	return ""
}

func (m *InitializeRequest) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *InitializeRequest) GetNetwork() NetworkMode {
	if m != nil {
		return m.Network
	}
	return NetworkMode_NetworkMode_LOCAL
}

// ConnectRequest initializes the libp2p host and connects to the Sonr network.
type ConnectRequest struct {
	// Peer ID of the node to connect to
	PeerId string `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	// Multiaddress of the node to connect to
	Multiaddr string `protobuf:"bytes,2,opt,name=multiaddr,proto3" json:"multiaddr,omitempty"`
	// Peer is the peer to connect to
	Node *common.NodeInfo `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_726257aebe87753d, []int{1}
}
func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return m.Size()
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *ConnectRequest) GetMultiaddr() string {
	if m != nil {
		return m.Multiaddr
	}
	return ""
}

func (m *ConnectRequest) GetNode() *common.NodeInfo {
	if m != nil {
		return m.Node
	}
	return nil
}

// CreateAccountWithKeyRequest allows the DSC to be specified manually when creating a request.
// Necessary for android (for now)
type RegisterRequest struct {
	Password  string            `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	AesDscKey []byte            `protobuf:"bytes,2,opt,name=aes_dsc_key,json=aesDscKey,proto3" json:"aes_dsc_key,omitempty"`
	AesPskKey []byte            `protobuf:"bytes,3,opt,name=aes_psk_key,json=aesPskKey,proto3" json:"aes_psk_key,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_726257aebe87753d, []int{2}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return m.Size()
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetAesDscKey() []byte {
	if m != nil {
		return m.AesDscKey
	}
	return nil
}

func (m *RegisterRequest) GetAesPskKey() []byte {
	if m != nil {
		return m.AesPskKey
	}
	return nil
}

func (m *RegisterRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterEnum("sonrhq.motor.bind.v1.NetworkMode", NetworkMode_name, NetworkMode_value)
	proto.RegisterType((*InitializeRequest)(nil), "sonrhq.motor.bind.v1.InitializeRequest")
	proto.RegisterType((*ConnectRequest)(nil), "sonrhq.motor.bind.v1.ConnectRequest")
	proto.RegisterType((*RegisterRequest)(nil), "sonrhq.motor.bind.v1.RegisterRequest")
	proto.RegisterMapType((map[string]string)(nil), "sonrhq.motor.bind.v1.RegisterRequest.MetadataEntry")
}

func init() { proto.RegisterFile("motor/bind/v1/request.proto", fileDescriptor_726257aebe87753d) }

var fileDescriptor_726257aebe87753d = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x4f, 0x6f, 0xda, 0x30,
	0x14, 0x27, 0xa5, 0x2d, 0x60, 0xb6, 0x2e, 0xb5, 0x98, 0x46, 0x61, 0xca, 0x18, 0x27, 0x34, 0x69,
	0x89, 0x4a, 0x0f, 0x9b, 0xd6, 0x13, 0x6d, 0x91, 0x86, 0x46, 0xd7, 0x2e, 0xe3, 0xb4, 0x4b, 0x14,
	0xf0, 0x1b, 0x58, 0x24, 0x71, 0xb0, 0x0d, 0x88, 0x7d, 0x8a, 0x7d, 0xac, 0x1d, 0x7b, 0xdc, 0x71,
	0x83, 0x8f, 0xb1, 0xcb, 0x64, 0x87, 0x50, 0x8a, 0x7a, 0x4a, 0x7e, 0x7f, 0x9e, 0xdf, 0x7b, 0x3f,
	0xcb, 0xa8, 0x1a, 0x32, 0xc9, 0xb8, 0xd3, 0xa7, 0x11, 0x71, 0x66, 0xa7, 0x0e, 0x87, 0xc9, 0x14,
	0x84, 0xb4, 0x63, 0xce, 0x24, 0xc3, 0x25, 0xc1, 0x22, 0x3e, 0x9a, 0xd8, 0xda, 0x63, 0x2b, 0x8f,
	0x3d, 0x3b, 0xad, 0x94, 0x06, 0x2c, 0x0c, 0x59, 0xa4, 0xec, 0x31, 0x00, 0x4f, 0xbc, 0xf5, 0xbf,
	0x06, 0x3a, 0xee, 0x44, 0x54, 0x52, 0x3f, 0xa0, 0x3f, 0xc0, 0x4d, 0xce, 0xc1, 0x55, 0x54, 0x20,
	0x30, 0xa3, 0x03, 0xf0, 0x28, 0x29, 0x1b, 0x35, 0xa3, 0x51, 0x70, 0xf3, 0x09, 0xd1, 0x21, 0xf8,
	0x04, 0xe5, 0x47, 0x2c, 0x04, 0x8f, 0x50, 0x5e, 0xde, 0xd3, 0x5a, 0x4e, 0xe1, 0x2b, 0xca, 0xf1,
	0x2b, 0x54, 0x14, 0xd3, 0x38, 0x66, 0x5c, 0x6a, 0x35, 0xab, 0x55, 0xb4, 0xa6, 0x94, 0xe1, 0x04,
	0xe5, 0x25, 0x84, 0xb1, 0x56, 0xf7, 0x93, 0x5a, 0x85, 0x95, 0x54, 0x45, 0x85, 0x80, 0x0d, 0xbd,
	0x00, 0x66, 0x10, 0x94, 0x0f, 0x92, 0x9e, 0x01, 0x1b, 0x76, 0x15, 0xc6, 0xe7, 0x28, 0x17, 0x81,
	0x9c, 0x33, 0x3e, 0x2e, 0x1f, 0xd6, 0x8c, 0xc6, 0x51, 0xf3, 0xb5, 0xfd, 0xd8, 0x92, 0xf6, 0xe7,
	0xc4, 0x74, 0xcd, 0x08, 0xb8, 0x69, 0x45, 0x7d, 0x8e, 0x8e, 0x2e, 0x59, 0x14, 0xc1, 0x40, 0xa6,
	0xfb, 0xbd, 0x40, 0x39, 0x95, 0xc1, 0xfd, 0x76, 0x87, 0x0a, 0x76, 0x08, 0x7e, 0x89, 0x0a, 0xe1,
	0x34, 0x90, 0xd4, 0x27, 0x24, 0x5d, 0xee, 0x9e, 0xc0, 0x36, 0xda, 0x8f, 0x18, 0x01, 0xbd, 0x57,
	0xb1, 0x59, 0x49, 0x47, 0x48, 0x82, 0xd5, 0xed, 0x19, 0x81, 0x4e, 0xf4, 0x9d, 0xb9, 0xda, 0x57,
	0xff, 0x67, 0xa0, 0x67, 0x2e, 0x0c, 0xa9, 0x90, 0xc0, 0xd3, 0xd6, 0x15, 0x94, 0x8f, 0x7d, 0x21,
	0xe6, 0x8c, 0x6f, 0x92, 0x4d, 0x31, 0xb6, 0x50, 0xd1, 0x07, 0xe1, 0x11, 0x31, 0xf0, 0xc6, 0xb0,
	0xd0, 0xfd, 0x9f, 0xb8, 0x05, 0x1f, 0xc4, 0x95, 0x18, 0x7c, 0x82, 0x45, 0xaa, 0xc7, 0x62, 0xac,
	0xf5, 0xec, 0x46, 0xbf, 0x15, 0x63, 0xa5, 0xdf, 0xa0, 0x7c, 0x08, 0xd2, 0x27, 0xbe, 0xf4, 0xcb,
	0xfb, 0xb5, 0x6c, 0xa3, 0xd8, 0x3c, 0x7b, 0x3c, 0xa6, 0x9d, 0xa1, 0xec, 0xeb, 0x75, 0x55, 0x3b,
	0x92, 0x7c, 0xe1, 0x6e, 0x0e, 0xa9, 0x9c, 0xa3, 0xa7, 0x0f, 0x24, 0x6c, 0xa2, 0xac, 0xea, 0x9c,
	0x0c, 0xae, 0x7e, 0x71, 0x09, 0x1d, 0xcc, 0xfc, 0x60, 0x0a, 0xeb, 0xb4, 0x12, 0xf0, 0x61, 0xef,
	0xbd, 0xf1, 0x66, 0x88, 0x8a, 0x5b, 0xd7, 0x81, 0x9f, 0xa3, 0xe3, 0x2d, 0xe8, 0x75, 0x6f, 0x2e,
	0x5b, 0x5d, 0x33, 0xb3, 0x4b, 0xb7, 0xba, 0xb7, 0x1f, 0x5b, 0xa6, 0x81, 0x4b, 0xc8, 0xdc, 0xa6,
	0x2f, 0xda, 0xbd, 0x96, 0xb9, 0xb7, 0xcb, 0xf6, 0xda, 0x5f, 0x7b, 0x66, 0xf6, 0xe2, 0xcb, 0xaf,
	0xa5, 0x65, 0xdc, 0x2d, 0x2d, 0xe3, 0xcf, 0xd2, 0x32, 0x7e, 0xae, 0xac, 0xcc, 0xdd, 0xca, 0xca,
	0xfc, 0x5e, 0x59, 0x99, 0x6f, 0xef, 0x86, 0x54, 0x8e, 0xa6, 0x7d, 0x75, 0x43, 0x8e, 0x0a, 0xe2,
	0xed, 0x68, 0xa2, 0xbf, 0x8e, 0x1c, 0x51, 0x4e, 0xbc, 0xd8, 0xe7, 0x72, 0xe1, 0xc8, 0x45, 0x0c,
	0xc2, 0x79, 0xf0, 0xa0, 0xfa, 0x87, 0xfa, 0x75, 0x9c, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x17,
	0x45, 0xb5, 0xfd, 0x68, 0x03, 0x00, 0x00,
}

func (m *InitializeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InitializeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InitializeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Network != 0 {
		i = encodeVarintRequest(dAtA, i, uint64(m.Network))
		i--
		dAtA[i] = 0x30
	}
	if len(m.LogLevel) > 0 {
		i -= len(m.LogLevel)
		copy(dAtA[i:], m.LogLevel)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.LogLevel)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TempDir) > 0 {
		i -= len(m.TempDir)
		copy(dAtA[i:], m.TempDir)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.TempDir)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SupportDir) > 0 {
		i -= len(m.SupportDir)
		copy(dAtA[i:], m.SupportDir)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.SupportDir)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.HomeDir) > 0 {
		i -= len(m.HomeDir)
		copy(dAtA[i:], m.HomeDir)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.HomeDir)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DeviceId) > 0 {
		i -= len(m.DeviceId)
		copy(dAtA[i:], m.DeviceId)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.DeviceId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConnectRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConnectRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Node != nil {
		{
			size, err := m.Node.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRequest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Multiaddr) > 0 {
		i -= len(m.Multiaddr)
		copy(dAtA[i:], m.Multiaddr)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.Multiaddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PeerId) > 0 {
		i -= len(m.PeerId)
		copy(dAtA[i:], m.PeerId)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.PeerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		for k := range m.Metadata {
			v := m.Metadata[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintRequest(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintRequest(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintRequest(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AesPskKey) > 0 {
		i -= len(m.AesPskKey)
		copy(dAtA[i:], m.AesPskKey)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.AesPskKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AesDscKey) > 0 {
		i -= len(m.AesDscKey)
		copy(dAtA[i:], m.AesDscKey)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.AesDscKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InitializeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DeviceId)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.HomeDir)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.SupportDir)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.TempDir)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.LogLevel)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	if m.Network != 0 {
		n += 1 + sovRequest(uint64(m.Network))
	}
	return n
}

func (m *ConnectRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.Multiaddr)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovRequest(uint64(l))
	}
	return n
}

func (m *RegisterRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.AesDscKey)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.AesPskKey)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovRequest(uint64(len(k))) + 1 + len(v) + sovRequest(uint64(len(v)))
			n += mapEntrySize + 1 + sovRequest(uint64(mapEntrySize))
		}
	}
	return n
}

func sovRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequest(x uint64) (n int) {
	return sovRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InitializeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: InitializeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InitializeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HomeDir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HomeDir = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportDir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupportDir = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TempDir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TempDir = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogLevel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LogLevel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			m.Network = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Network |= NetworkMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func (m *ConnectRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: ConnectRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiaddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Multiaddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &common.NodeInfo{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func (m *RegisterRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: RegisterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AesDscKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AesDscKey = append(m.AesDscKey[:0], dAtA[iNdEx:postIndex]...)
			if m.AesDscKey == nil {
				m.AesDscKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AesPskKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AesPskKey = append(m.AesPskKey[:0], dAtA[iNdEx:postIndex]...)
			if m.AesPskKey == nil {
				m.AesPskKey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRequest
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRequest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthRequest
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthRequest
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRequest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthRequest
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthRequest
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipRequest(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthRequest
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func skipRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
				return 0, ErrInvalidLengthRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequest = fmt.Errorf("proto: unexpected end of group")
)
