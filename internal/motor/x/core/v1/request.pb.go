// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: core/v1/request.proto

// Package Motor is used for defining a Motor node and its properties.

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InitializeRequest_IPAddress_Family int32

const (
	InitializeRequest_IPAddress_FAMILY_UNSPECIFIED InitializeRequest_IPAddress_Family = 0
	InitializeRequest_IPAddress_FAMILY_IPV4        InitializeRequest_IPAddress_Family = 1 // IPv4 Address
	InitializeRequest_IPAddress_FAMILY_IPV6        InitializeRequest_IPAddress_Family = 2 // IPv6 Address
)

// Enum value maps for InitializeRequest_IPAddress_Family.
var (
	InitializeRequest_IPAddress_Family_name = map[int32]string{
		0: "FAMILY_UNSPECIFIED",
		1: "FAMILY_IPV4",
		2: "FAMILY_IPV6",
	}
	InitializeRequest_IPAddress_Family_value = map[string]int32{
		"FAMILY_UNSPECIFIED": 0,
		"FAMILY_IPV4":        1,
		"FAMILY_IPV6":        2,
	}
)

func (x InitializeRequest_IPAddress_Family) Enum() *InitializeRequest_IPAddress_Family {
	p := new(InitializeRequest_IPAddress_Family)
	*p = x
	return p
}

func (x InitializeRequest_IPAddress_Family) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InitializeRequest_IPAddress_Family) Descriptor() protoreflect.EnumDescriptor {
	return file_core_v1_request_proto_enumTypes[0].Descriptor()
}

func (InitializeRequest_IPAddress_Family) Type() protoreflect.EnumType {
	return &file_core_v1_request_proto_enumTypes[0]
}

func (x InitializeRequest_IPAddress_Family) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InitializeRequest_IPAddress_Family.Descriptor instead.
func (InitializeRequest_IPAddress_Family) EnumDescriptor() ([]byte, []int) {
	return file_core_v1_request_proto_rawDescGZIP(), []int{0, 4, 0}
}

// -----------------------------------------------------------------------------
// Motor Node API
// -----------------------------------------------------------------------------
// (Client) InitializeRequest Message to Establish Sonr Host/API/Room
type InitializeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostOptions    *InitializeRequest_HostOptions    `protobuf:"bytes,3,opt,name=host_options,json=hostOptions,proto3" json:"host_options,omitempty"`                                                                  // Libp2p Host config
	ServiceOptions *InitializeRequest_ServiceOptions `protobuf:"bytes,4,opt,name=service_options,json=serviceOptions,proto3" json:"service_options,omitempty"`                                                         // Service Config
	DeviceOptions  *InitializeRequest_DeviceOptions  `protobuf:"bytes,5,opt,name=device_options,json=deviceOptions,proto3" json:"device_options,omitempty"`                                                            // File System Config
	Variables      map[string]string                 `protobuf:"bytes,6,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Domain TXT Records
}

func (x *InitializeRequest) Reset() {
	*x = InitializeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeRequest) ProtoMessage() {}

func (x *InitializeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeRequest.ProtoReflect.Descriptor instead.
func (*InitializeRequest) Descriptor() ([]byte, []int) {
	return file_core_v1_request_proto_rawDescGZIP(), []int{0}
}

func (x *InitializeRequest) GetHostOptions() *InitializeRequest_HostOptions {
	if x != nil {
		return x.HostOptions
	}
	return nil
}

func (x *InitializeRequest) GetServiceOptions() *InitializeRequest_ServiceOptions {
	if x != nil {
		return x.ServiceOptions
	}
	return nil
}

func (x *InitializeRequest) GetDeviceOptions() *InitializeRequest_DeviceOptions {
	if x != nil {
		return x.DeviceOptions
	}
	return nil
}

func (x *InitializeRequest) GetVariables() map[string]string {
	if x != nil {
		return x.Variables
	}
	return nil
}

// Optional Message to Initialize FileSystem
type InitializeRequest_DeviceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Device ID
	HomeDir    string `protobuf:"bytes,2,opt,name=home_dir,json=homeDir,proto3" json:"home_dir,omitempty"`
	SupportDir string `protobuf:"bytes,3,opt,name=support_dir,json=supportDir,proto3" json:"support_dir,omitempty"`
	TempDir    string `protobuf:"bytes,4,opt,name=temp_dir,json=tempDir,proto3" json:"temp_dir,omitempty"`
}

func (x *InitializeRequest_DeviceOptions) Reset() {
	*x = InitializeRequest_DeviceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeRequest_DeviceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeRequest_DeviceOptions) ProtoMessage() {}

func (x *InitializeRequest_DeviceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeRequest_DeviceOptions.ProtoReflect.Descriptor instead.
func (*InitializeRequest_DeviceOptions) Descriptor() ([]byte, []int) {
	return file_core_v1_request_proto_rawDescGZIP(), []int{0, 1}
}

func (x *InitializeRequest_DeviceOptions) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InitializeRequest_DeviceOptions) GetHomeDir() string {
	if x != nil {
		return x.HomeDir
	}
	return ""
}

func (x *InitializeRequest_DeviceOptions) GetSupportDir() string {
	if x != nil {
		return x.SupportDir
	}
	return ""
}

func (x *InitializeRequest_DeviceOptions) GetTempDir() string {
	if x != nil {
		return x.TempDir
	}
	return ""
}

// Libp2p Host Options
type InitializeRequest_HostOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuicTransport bool                           `protobuf:"varint,1,opt,name=quic_transport,json=quicTransport,proto3" json:"quic_transport,omitempty"` // Enable QUIC Transport
	HttpTransport bool                           `protobuf:"varint,2,opt,name=http_transport,json=httpTransport,proto3" json:"http_transport,omitempty"` // Enable HTTP Transport
	Ipv4Only      bool                           `protobuf:"varint,3,opt,name=ipv4_only,json=ipv4Only,proto3" json:"ipv4_only,omitempty"`                // Enable IPv4 Only
	ListenAddrs   []*InitializeRequest_IPAddress `protobuf:"bytes,4,rep,name=listen_addrs,json=listenAddrs,proto3" json:"listen_addrs,omitempty"`        // List of Listen Addresses (optional)
}

func (x *InitializeRequest_HostOptions) Reset() {
	*x = InitializeRequest_HostOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeRequest_HostOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeRequest_HostOptions) ProtoMessage() {}

func (x *InitializeRequest_HostOptions) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeRequest_HostOptions.ProtoReflect.Descriptor instead.
func (*InitializeRequest_HostOptions) Descriptor() ([]byte, []int) {
	return file_core_v1_request_proto_rawDescGZIP(), []int{0, 2}
}

func (x *InitializeRequest_HostOptions) GetQuicTransport() bool {
	if x != nil {
		return x.QuicTransport
	}
	return false
}

func (x *InitializeRequest_HostOptions) GetHttpTransport() bool {
	if x != nil {
		return x.HttpTransport
	}
	return false
}

func (x *InitializeRequest_HostOptions) GetIpv4Only() bool {
	if x != nil {
		return x.Ipv4Only
	}
	return false
}

func (x *InitializeRequest_HostOptions) GetListenAddrs() []*InitializeRequest_IPAddress {
	if x != nil {
		return x.ListenAddrs
	}
	return nil
}

// Service Configuration
type InitializeRequest_ServiceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Textile  bool  `protobuf:"varint,1,opt,name=textile,proto3" json:"textile,omitempty"`   // Enable Textile Client and Threads
	Mailbox  bool  `protobuf:"varint,2,opt,name=mailbox,proto3" json:"mailbox,omitempty"`   // Enable Mailbox
	Buckets  bool  `protobuf:"varint,3,opt,name=buckets,proto3" json:"buckets,omitempty"`   // Enable Buckets
	Interval int32 `protobuf:"varint,4,opt,name=interval,proto3" json:"interval,omitempty"` // Auto Update Interval (seconds) - Default 5s
}

func (x *InitializeRequest_ServiceOptions) Reset() {
	*x = InitializeRequest_ServiceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeRequest_ServiceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeRequest_ServiceOptions) ProtoMessage() {}

func (x *InitializeRequest_ServiceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeRequest_ServiceOptions.ProtoReflect.Descriptor instead.
func (*InitializeRequest_ServiceOptions) Descriptor() ([]byte, []int) {
	return file_core_v1_request_proto_rawDescGZIP(), []int{0, 3}
}

func (x *InitializeRequest_ServiceOptions) GetTextile() bool {
	if x != nil {
		return x.Textile
	}
	return false
}

func (x *InitializeRequest_ServiceOptions) GetMailbox() bool {
	if x != nil {
		return x.Mailbox
	}
	return false
}

func (x *InitializeRequest_ServiceOptions) GetBuckets() bool {
	if x != nil {
		return x.Buckets
	}
	return false
}

func (x *InitializeRequest_ServiceOptions) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

// IP Address Interface
type InitializeRequest_IPAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                                                   // Name of Interface
	Address  string                             `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`                                                             // IP Address of Interface
	Internal bool                               `protobuf:"varint,3,opt,name=internal,proto3" json:"internal,omitempty"`                                                          // Wether it is a Loopback Interface
	Family   InitializeRequest_IPAddress_Family `protobuf:"varint,4,opt,name=family,proto3,enum=sonrio.motor.core.v1.InitializeRequest_IPAddress_Family" json:"family,omitempty"` // Address Family
}

func (x *InitializeRequest_IPAddress) Reset() {
	*x = InitializeRequest_IPAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_v1_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeRequest_IPAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeRequest_IPAddress) ProtoMessage() {}

func (x *InitializeRequest_IPAddress) ProtoReflect() protoreflect.Message {
	mi := &file_core_v1_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeRequest_IPAddress.ProtoReflect.Descriptor instead.
func (*InitializeRequest_IPAddress) Descriptor() ([]byte, []int) {
	return file_core_v1_request_proto_rawDescGZIP(), []int{0, 4}
}

func (x *InitializeRequest_IPAddress) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InitializeRequest_IPAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InitializeRequest_IPAddress) GetInternal() bool {
	if x != nil {
		return x.Internal
	}
	return false
}

func (x *InitializeRequest_IPAddress) GetFamily() InitializeRequest_IPAddress_Family {
	if x != nil {
		return x.Family
	}
	return InitializeRequest_IPAddress_FAMILY_UNSPECIFIED
}

var File_core_v1_request_proto protoreflect.FileDescriptor

var file_core_v1_request_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e,
	0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xf1, 0x08,
	0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b,
	0x68, 0x6f, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5f, 0x0a, 0x0f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f,
	0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5c, 0x0a, 0x0e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f,
	0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x54, 0x0a, 0x09, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x1a, 0x3c, 0x0a, 0x0e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x76,
	0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x68, 0x6f, 0x6d, 0x65, 0x44, 0x69, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x65, 0x6d, 0x70, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x65, 0x6d, 0x70, 0x44, 0x69, 0x72, 0x1a, 0xce, 0x01, 0x0a, 0x0b, 0x48, 0x6f, 0x73, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x71, 0x75, 0x69, 0x63, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x71, 0x75, 0x69, 0x63, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x6f, 0x6e, 0x6c,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x70, 0x76, 0x34, 0x4f, 0x6e, 0x6c,
	0x79, 0x12, 0x54, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f,
	0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x73, 0x1a, 0x7a, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x78,
	0x74, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x65, 0x78, 0x74,
	0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x1a, 0xeb, 0x01, 0x0a, 0x09, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x50, 0x0a, 0x06, 0x66,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x46,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x22, 0x42, 0x0a,
	0x06, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x41, 0x4d, 0x49, 0x4c,
	0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x49, 0x50, 0x56, 0x34, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x49, 0x50, 0x56, 0x36, 0x10,
	0x02, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2f, 0x78, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_v1_request_proto_rawDescOnce sync.Once
	file_core_v1_request_proto_rawDescData = file_core_v1_request_proto_rawDesc
)

func file_core_v1_request_proto_rawDescGZIP() []byte {
	file_core_v1_request_proto_rawDescOnce.Do(func() {
		file_core_v1_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_v1_request_proto_rawDescData)
	})
	return file_core_v1_request_proto_rawDescData
}

var file_core_v1_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_v1_request_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_core_v1_request_proto_goTypes = []interface{}{
	(InitializeRequest_IPAddress_Family)(0),  // 0: sonrio.motor.core.v1.InitializeRequest.IPAddress.Family
	(*InitializeRequest)(nil),                // 1: sonrio.motor.core.v1.InitializeRequest
	nil,                                      // 2: sonrio.motor.core.v1.InitializeRequest.VariablesEntry
	(*InitializeRequest_DeviceOptions)(nil),  // 3: sonrio.motor.core.v1.InitializeRequest.DeviceOptions
	(*InitializeRequest_HostOptions)(nil),    // 4: sonrio.motor.core.v1.InitializeRequest.HostOptions
	(*InitializeRequest_ServiceOptions)(nil), // 5: sonrio.motor.core.v1.InitializeRequest.ServiceOptions
	(*InitializeRequest_IPAddress)(nil),      // 6: sonrio.motor.core.v1.InitializeRequest.IPAddress
}
var file_core_v1_request_proto_depIdxs = []int32{
	4, // 0: sonrio.motor.core.v1.InitializeRequest.host_options:type_name -> sonrio.motor.core.v1.InitializeRequest.HostOptions
	5, // 1: sonrio.motor.core.v1.InitializeRequest.service_options:type_name -> sonrio.motor.core.v1.InitializeRequest.ServiceOptions
	3, // 2: sonrio.motor.core.v1.InitializeRequest.device_options:type_name -> sonrio.motor.core.v1.InitializeRequest.DeviceOptions
	2, // 3: sonrio.motor.core.v1.InitializeRequest.variables:type_name -> sonrio.motor.core.v1.InitializeRequest.VariablesEntry
	6, // 4: sonrio.motor.core.v1.InitializeRequest.HostOptions.listen_addrs:type_name -> sonrio.motor.core.v1.InitializeRequest.IPAddress
	0, // 5: sonrio.motor.core.v1.InitializeRequest.IPAddress.family:type_name -> sonrio.motor.core.v1.InitializeRequest.IPAddress.Family
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_core_v1_request_proto_init() }
func file_core_v1_request_proto_init() {
	if File_core_v1_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_v1_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeRequest_DeviceOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeRequest_HostOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeRequest_ServiceOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_v1_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeRequest_IPAddress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_v1_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_v1_request_proto_goTypes,
		DependencyIndexes: file_core_v1_request_proto_depIdxs,
		EnumInfos:         file_core_v1_request_proto_enumTypes,
		MessageInfos:      file_core_v1_request_proto_msgTypes,
	}.Build()
	File_core_v1_request_proto = out.File
	file_core_v1_request_proto_rawDesc = nil
	file_core_v1_request_proto_goTypes = nil
	file_core_v1_request_proto_depIdxs = nil
}
