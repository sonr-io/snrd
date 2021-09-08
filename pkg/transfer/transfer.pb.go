// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: proto/protocols/transfer.proto

package transfer

import (
	common "github.com/sonr-io/core/internal/common"
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

// InviteReach is Distance for an Invite
type InviteReach int32

const (
	InviteReach_NONE   InviteReach = 0 // No Reach
	InviteReach_LOCAL  InviteReach = 1 // Local Reach
	InviteReach_REMOTE InviteReach = 2 // Remote Reach
	InviteReach_DIRECT InviteReach = 3 // Direct Reach
)

// Enum value maps for InviteReach.
var (
	InviteReach_name = map[int32]string{
		0: "NONE",
		1: "LOCAL",
		2: "REMOTE",
		3: "DIRECT",
	}
	InviteReach_value = map[string]int32{
		"NONE":   0,
		"LOCAL":  1,
		"REMOTE": 2,
		"DIRECT": 3,
	}
)

func (x InviteReach) Enum() *InviteReach {
	p := new(InviteReach)
	*p = x
	return p
}

func (x InviteReach) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InviteReach) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_protocols_transfer_proto_enumTypes[0].Descriptor()
}

func (InviteReach) Type() protoreflect.EnumType {
	return &file_proto_protocols_transfer_proto_enumTypes[0]
}

func (x InviteReach) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InviteReach.Descriptor instead.
func (InviteReach) EnumDescriptor() ([]byte, []int) {
	return file_proto_protocols_transfer_proto_rawDescGZIP(), []int{0}
}

// InviteEvent notifies Peer that an Invite has been received
type InviteEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InviteId string           `protobuf:"bytes,1,opt,name=inviteId,proto3" json:"inviteId,omitempty"`
	From     *common.Peer     `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Transfer *common.Transfer `protobuf:"bytes,3,opt,name=transfer,proto3" json:"transfer,omitempty"`                                     // Attached Data
	Reach    InviteReach      `protobuf:"varint,4,opt,name=reach,proto3,enum=sonr.protocols.transfer.InviteReach" json:"reach,omitempty"` // Reach of Invite
}

func (x *InviteEvent) Reset() {
	*x = InviteEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteEvent) ProtoMessage() {}

func (x *InviteEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteEvent.ProtoReflect.Descriptor instead.
func (*InviteEvent) Descriptor() ([]byte, []int) {
	return file_proto_protocols_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *InviteEvent) GetInviteId() string {
	if x != nil {
		return x.InviteId
	}
	return ""
}

func (x *InviteEvent) GetFrom() *common.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *InviteEvent) GetTransfer() *common.Transfer {
	if x != nil {
		return x.Transfer
	}
	return nil
}

func (x *InviteEvent) GetReach() InviteReach {
	if x != nil {
		return x.Reach
	}
	return InviteReach_NONE
}

// Invitation Message sent on RPC
type InviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InviteId string           `protobuf:"bytes,1,opt,name=inviteId,proto3" json:"inviteId,omitempty"`
	From     *common.Peer     `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`                                             // Users Peer Data
	To       *common.Peer     `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`                                                 // Receivers Peer Data
	Transfer *common.Transfer `protobuf:"bytes,4,opt,name=transfer,proto3" json:"transfer,omitempty"`                                     // Attached Data
	Reach    InviteReach      `protobuf:"varint,5,opt,name=reach,proto3,enum=sonr.protocols.transfer.InviteReach" json:"reach,omitempty"` // Reach of Invite
	Metadata *common.Metadata `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`                                     // Metadata
}

func (x *InviteRequest) Reset() {
	*x = InviteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteRequest) ProtoMessage() {}

func (x *InviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteRequest.ProtoReflect.Descriptor instead.
func (*InviteRequest) Descriptor() ([]byte, []int) {
	return file_proto_protocols_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *InviteRequest) GetInviteId() string {
	if x != nil {
		return x.InviteId
	}
	return ""
}

func (x *InviteRequest) GetFrom() *common.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *InviteRequest) GetTo() *common.Peer {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *InviteRequest) GetTransfer() *common.Transfer {
	if x != nil {
		return x.Transfer
	}
	return nil
}

func (x *InviteRequest) GetReach() InviteReach {
	if x != nil {
		return x.Reach
	}
	return InviteReach_NONE
}

func (x *InviteRequest) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Reply Message sent on RPC
type InviteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool             `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                                      // Success
	InviteId   string           `protobuf:"bytes,2,opt,name=inviteId,proto3" json:"inviteId,omitempty"`                                     // Invitation ID
	TransferId string           `protobuf:"bytes,3,opt,name=transferId,proto3" json:"transferId,omitempty"`                                 // ID of new Transfer
	Reach      InviteReach      `protobuf:"varint,4,opt,name=reach,proto3,enum=sonr.protocols.transfer.InviteReach" json:"reach,omitempty"` // Reach of Invite
	Metadata   *common.Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`                                     // Metadata
}

func (x *InviteResponse) Reset() {
	*x = InviteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteResponse) ProtoMessage() {}

func (x *InviteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteResponse.ProtoReflect.Descriptor instead.
func (*InviteResponse) Descriptor() ([]byte, []int) {
	return file_proto_protocols_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *InviteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *InviteResponse) GetInviteId() string {
	if x != nil {
		return x.InviteId
	}
	return ""
}

func (x *InviteResponse) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

func (x *InviteResponse) GetReach() InviteReach {
	if x != nil {
		return x.Reach
	}
	return InviteReach_NONE
}

func (x *InviteResponse) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_proto_protocols_transfer_proto protoreflect.FileDescriptor

var file_proto_protocols_transfer_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x0b,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2f, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x52, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x3a, 0x0a,
	0x05, 0x72, 0x65, 0x61, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x61,
	0x63, 0x68, 0x52, 0x05, 0x72, 0x65, 0x61, 0x63, 0x68, 0x22, 0x8f, 0x02, 0x0a, 0x0d, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1f, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2f, 0x0a,
	0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x52, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x3a,
	0x0a, 0x05, 0x72, 0x65, 0x61, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x61, 0x63, 0x68, 0x52, 0x05, 0x72, 0x65, 0x61, 0x63, 0x68, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd3, 0x01, 0x0a, 0x0e,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x63, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x61, 0x63, 0x68, 0x52, 0x05, 0x72, 0x65, 0x61, 0x63, 0x68,
	0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2a, 0x3a, 0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x61, 0x63, 0x68,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f,
	0x43, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x54, 0x45, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10, 0x03, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72,
	0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_protocols_transfer_proto_rawDescOnce sync.Once
	file_proto_protocols_transfer_proto_rawDescData = file_proto_protocols_transfer_proto_rawDesc
)

func file_proto_protocols_transfer_proto_rawDescGZIP() []byte {
	file_proto_protocols_transfer_proto_rawDescOnce.Do(func() {
		file_proto_protocols_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_protocols_transfer_proto_rawDescData)
	})
	return file_proto_protocols_transfer_proto_rawDescData
}

var file_proto_protocols_transfer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_protocols_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_protocols_transfer_proto_goTypes = []interface{}{
	(InviteReach)(0),        // 0: sonr.protocols.transfer.InviteReach
	(*InviteEvent)(nil),     // 1: sonr.protocols.transfer.InviteEvent
	(*InviteRequest)(nil),   // 2: sonr.protocols.transfer.InviteRequest
	(*InviteResponse)(nil),  // 3: sonr.protocols.transfer.InviteResponse
	(*common.Peer)(nil),     // 4: sonr.core.Peer
	(*common.Transfer)(nil), // 5: sonr.core.Transfer
	(*common.Metadata)(nil), // 6: sonr.core.Metadata
}
var file_proto_protocols_transfer_proto_depIdxs = []int32{
	4,  // 0: sonr.protocols.transfer.InviteEvent.from:type_name -> sonr.core.Peer
	5,  // 1: sonr.protocols.transfer.InviteEvent.transfer:type_name -> sonr.core.Transfer
	0,  // 2: sonr.protocols.transfer.InviteEvent.reach:type_name -> sonr.protocols.transfer.InviteReach
	4,  // 3: sonr.protocols.transfer.InviteRequest.from:type_name -> sonr.core.Peer
	4,  // 4: sonr.protocols.transfer.InviteRequest.to:type_name -> sonr.core.Peer
	5,  // 5: sonr.protocols.transfer.InviteRequest.transfer:type_name -> sonr.core.Transfer
	0,  // 6: sonr.protocols.transfer.InviteRequest.reach:type_name -> sonr.protocols.transfer.InviteReach
	6,  // 7: sonr.protocols.transfer.InviteRequest.metadata:type_name -> sonr.core.Metadata
	0,  // 8: sonr.protocols.transfer.InviteResponse.reach:type_name -> sonr.protocols.transfer.InviteReach
	6,  // 9: sonr.protocols.transfer.InviteResponse.metadata:type_name -> sonr.core.Metadata
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_protocols_transfer_proto_init() }
func file_proto_protocols_transfer_proto_init() {
	if File_proto_protocols_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_protocols_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteEvent); i {
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
		file_proto_protocols_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteRequest); i {
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
		file_proto_protocols_transfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteResponse); i {
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
			RawDescriptor: file_proto_protocols_transfer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_protocols_transfer_proto_goTypes,
		DependencyIndexes: file_proto_protocols_transfer_proto_depIdxs,
		EnumInfos:         file_proto_protocols_transfer_proto_enumTypes,
		MessageInfos:      file_proto_protocols_transfer_proto_msgTypes,
	}.Build()
	File_proto_protocols_transfer_proto = out.File
	file_proto_protocols_transfer_proto_rawDesc = nil
	file_proto_protocols_transfer_proto_goTypes = nil
	file_proto_protocols_transfer_proto_depIdxs = nil
}
