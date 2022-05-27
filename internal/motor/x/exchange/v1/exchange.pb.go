// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: exchange/v1/exchange.proto

// Package exchange defines interfaces and types for exchange between two nodes in the network.

package v1

import (
	v11 "github.com/sonr-io/sonr/internal/motor/x/discover/v1"
	v1 "github.com/sonr-io/sonr/internal/motor/x/transmit/v1"
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

type MailboxMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                 // ID is the Message ID
	Payload   *v1.Payload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`                       // Payload is the message data
	From      *v11.Peer   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`                             // Users Peer Data
	To        *v11.Peer   `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`                                 // Receivers Peer Data
	Metadata  *Metadata   `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`                     // Metadata
	CreatedAt int64       `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // Timestamp
}

func (x *MailboxMessage) Reset() {
	*x = MailboxMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_v1_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxMessage) ProtoMessage() {}

func (x *MailboxMessage) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_v1_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxMessage.ProtoReflect.Descriptor instead.
func (*MailboxMessage) Descriptor() ([]byte, []int) {
	return file_exchange_v1_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *MailboxMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MailboxMessage) GetPayload() *v1.Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *MailboxMessage) GetFrom() *v11.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *MailboxMessage) GetTo() *v11.Peer {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *MailboxMessage) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *MailboxMessage) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

// Invitation Message sent on RPC
type InviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  *v1.Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`   // Attached Data
	From     *v11.Peer   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`         // Users Peer Data
	To       *v11.Peer   `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`             // Receivers Peer Data
	Metadata *Metadata   `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"` // Metadata
}

func (x *InviteRequest) Reset() {
	*x = InviteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_v1_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteRequest) ProtoMessage() {}

func (x *InviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_v1_exchange_proto_msgTypes[1]
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
	return file_exchange_v1_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *InviteRequest) GetPayload() *v1.Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *InviteRequest) GetFrom() *v11.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *InviteRequest) GetTo() *v11.Peer {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *InviteRequest) GetMetadata() *Metadata {
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

	Decision bool      `protobuf:"varint,1,opt,name=decision,proto3" json:"decision,omitempty"` // Success
	From     *v11.Peer `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`          // Users Peer Data
	To       *v11.Peer `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`              // Receivers Peer Data
	Metadata *Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`  // Metadata
}

func (x *InviteResponse) Reset() {
	*x = InviteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_v1_exchange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteResponse) ProtoMessage() {}

func (x *InviteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_v1_exchange_proto_msgTypes[2]
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
	return file_exchange_v1_exchange_proto_rawDescGZIP(), []int{2}
}

func (x *InviteResponse) GetDecision() bool {
	if x != nil {
		return x.Decision
	}
	return false
}

func (x *InviteResponse) GetFrom() *v11.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *InviteResponse) GetTo() *v11.Peer {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *InviteResponse) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Received Mail Event
type OnMailboxMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`             // ID is the Message ID
	Buffer   []byte    `protobuf:"bytes,2,opt,name=buffer,proto3" json:"buffer,omitempty"`     // Buffer is the message data
	From     *v11.Peer `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`         // Users Peer Data
	To       *v11.Peer `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`             // Receivers Peer Data
	Metadata *Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"` // Metadata
}

func (x *OnMailboxMessageResponse) Reset() {
	*x = OnMailboxMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_v1_exchange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnMailboxMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnMailboxMessageResponse) ProtoMessage() {}

func (x *OnMailboxMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_v1_exchange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnMailboxMessageResponse.ProtoReflect.Descriptor instead.
func (*OnMailboxMessageResponse) Descriptor() ([]byte, []int) {
	return file_exchange_v1_exchange_proto_rawDescGZIP(), []int{3}
}

func (x *OnMailboxMessageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OnMailboxMessageResponse) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

func (x *OnMailboxMessageResponse) GetFrom() *v11.Peer {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *OnMailboxMessageResponse) GetTo() *v11.Peer {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *OnMailboxMessageResponse) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Shared Metadata for Messages on all Protocols
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix timestamp
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Node ID
	NodeId string `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Signature of the message
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	// Public Key of the message sender
	PublicKey []byte `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_v1_exchange_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_v1_exchange_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_exchange_v1_exchange_proto_rawDescGZIP(), []int{4}
}

func (x *Metadata) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Metadata) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Metadata) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Metadata) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

var File_exchange_v1_exchange_proto protoreflect.FileDescriptor

var file_exchange_v1_exchange_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x6f,
	0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0,
	0x02, 0x0a, 0x0e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x3b, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f,
	0x72, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x32,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x2e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f,
	0x74, 0x6f, 0x72, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xf0, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f,
	0x74, 0x6f, 0x72, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x32, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e,
	0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xd0, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74,
	0x6f, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xe6, 0x01, 0x0a, 0x18, 0x4f, 0x6e, 0x4d, 0x61,
	0x69, 0x6c, 0x62, 0x6f, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x6f, 0x6e,
	0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x2e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x69, 0x6f, 0x2e, 0x6d, 0x6f, 0x74, 0x6f,
	0x72, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x7e, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x74, 0x6f, 0x72, 0x2f, 0x78, 0x2f, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exchange_v1_exchange_proto_rawDescOnce sync.Once
	file_exchange_v1_exchange_proto_rawDescData = file_exchange_v1_exchange_proto_rawDesc
)

func file_exchange_v1_exchange_proto_rawDescGZIP() []byte {
	file_exchange_v1_exchange_proto_rawDescOnce.Do(func() {
		file_exchange_v1_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_exchange_v1_exchange_proto_rawDescData)
	})
	return file_exchange_v1_exchange_proto_rawDescData
}

var file_exchange_v1_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_exchange_v1_exchange_proto_goTypes = []interface{}{
	(*MailboxMessage)(nil),           // 0: sonrio.motor.exchange.v1.MailboxMessage
	(*InviteRequest)(nil),            // 1: sonrio.motor.exchange.v1.InviteRequest
	(*InviteResponse)(nil),           // 2: sonrio.motor.exchange.v1.InviteResponse
	(*OnMailboxMessageResponse)(nil), // 3: sonrio.motor.exchange.v1.OnMailboxMessageResponse
	(*Metadata)(nil),                 // 4: sonrio.motor.exchange.v1.Metadata
	(*v1.Payload)(nil),               // 5: sonrio.motor.transmit.v1.Payload
	(*v11.Peer)(nil),                 // 6: sonrio.motor.discover.v1.Peer
}
var file_exchange_v1_exchange_proto_depIdxs = []int32{
	5,  // 0: sonrio.motor.exchange.v1.MailboxMessage.payload:type_name -> sonrio.motor.transmit.v1.Payload
	6,  // 1: sonrio.motor.exchange.v1.MailboxMessage.from:type_name -> sonrio.motor.discover.v1.Peer
	6,  // 2: sonrio.motor.exchange.v1.MailboxMessage.to:type_name -> sonrio.motor.discover.v1.Peer
	4,  // 3: sonrio.motor.exchange.v1.MailboxMessage.metadata:type_name -> sonrio.motor.exchange.v1.Metadata
	5,  // 4: sonrio.motor.exchange.v1.InviteRequest.payload:type_name -> sonrio.motor.transmit.v1.Payload
	6,  // 5: sonrio.motor.exchange.v1.InviteRequest.from:type_name -> sonrio.motor.discover.v1.Peer
	6,  // 6: sonrio.motor.exchange.v1.InviteRequest.to:type_name -> sonrio.motor.discover.v1.Peer
	4,  // 7: sonrio.motor.exchange.v1.InviteRequest.metadata:type_name -> sonrio.motor.exchange.v1.Metadata
	6,  // 8: sonrio.motor.exchange.v1.InviteResponse.from:type_name -> sonrio.motor.discover.v1.Peer
	6,  // 9: sonrio.motor.exchange.v1.InviteResponse.to:type_name -> sonrio.motor.discover.v1.Peer
	4,  // 10: sonrio.motor.exchange.v1.InviteResponse.metadata:type_name -> sonrio.motor.exchange.v1.Metadata
	6,  // 11: sonrio.motor.exchange.v1.OnMailboxMessageResponse.from:type_name -> sonrio.motor.discover.v1.Peer
	6,  // 12: sonrio.motor.exchange.v1.OnMailboxMessageResponse.to:type_name -> sonrio.motor.discover.v1.Peer
	4,  // 13: sonrio.motor.exchange.v1.OnMailboxMessageResponse.metadata:type_name -> sonrio.motor.exchange.v1.Metadata
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_exchange_v1_exchange_proto_init() }
func file_exchange_v1_exchange_proto_init() {
	if File_exchange_v1_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exchange_v1_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxMessage); i {
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
		file_exchange_v1_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_exchange_v1_exchange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_exchange_v1_exchange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnMailboxMessageResponse); i {
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
		file_exchange_v1_exchange_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
			RawDescriptor: file_exchange_v1_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_exchange_v1_exchange_proto_goTypes,
		DependencyIndexes: file_exchange_v1_exchange_proto_depIdxs,
		MessageInfos:      file_exchange_v1_exchange_proto_msgTypes,
	}.Build()
	File_exchange_v1_exchange_proto = out.File
	file_exchange_v1_exchange_proto_rawDesc = nil
	file_exchange_v1_exchange_proto_goTypes = nil
	file_exchange_v1_exchange_proto_depIdxs = nil
}
