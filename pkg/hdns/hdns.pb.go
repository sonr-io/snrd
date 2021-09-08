// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: proto/protocols/hdns.proto

package hdns

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

// CreateSNameRequest is Message for Signing Request (Hmac Sha256)
type CreateSNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value to be signed
	SName    string `protobuf:"bytes,1,opt,name=sName,proto3" json:"sName,omitempty"`       // SName combined with Device ID and Hashed
	Mnemonic string `protobuf:"bytes,2,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"` // Mnemonic Hashed with private key for fingerprint
}

func (x *CreateSNameRequest) Reset() {
	*x = CreateSNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_hdns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSNameRequest) ProtoMessage() {}

func (x *CreateSNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_hdns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSNameRequest.ProtoReflect.Descriptor instead.
func (*CreateSNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_protocols_hdns_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSNameRequest) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *CreateSNameRequest) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

// CreateSNameResponse is Message for Signing Response (Hmac Sha256)
type CreateSNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // If Values were Signed
	// Resulting Signed Values
	SignedPrefix      string `protobuf:"bytes,2,opt,name=signedPrefix,proto3" json:"signedPrefix,omitempty"`           // Message for List of Bytes
	SignedFingerprint string `protobuf:"bytes,3,opt,name=signedFingerprint,proto3" json:"signedFingerprint,omitempty"` // Fingerprint Value
	PublicKey         string `protobuf:"bytes,4,opt,name=publicKey,proto3" json:"publicKey,omitempty"`                 // Base64 Encoded Public Key
	GivenSName        string `protobuf:"bytes,5,opt,name=givenSName,proto3" json:"givenSName,omitempty"`               // Provided SName
	GivenMnemonic     string `protobuf:"bytes,6,opt,name=givenMnemonic,proto3" json:"givenMnemonic,omitempty"`         // Provided Mnemonic
}

func (x *CreateSNameResponse) Reset() {
	*x = CreateSNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_hdns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSNameResponse) ProtoMessage() {}

func (x *CreateSNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_hdns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSNameResponse.ProtoReflect.Descriptor instead.
func (*CreateSNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_protocols_hdns_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSNameResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateSNameResponse) GetSignedPrefix() string {
	if x != nil {
		return x.SignedPrefix
	}
	return ""
}

func (x *CreateSNameResponse) GetSignedFingerprint() string {
	if x != nil {
		return x.SignedFingerprint
	}
	return ""
}

func (x *CreateSNameResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *CreateSNameResponse) GetGivenSName() string {
	if x != nil {
		return x.GivenSName
	}
	return ""
}

func (x *CreateSNameResponse) GetGivenMnemonic() string {
	if x != nil {
		return x.GivenMnemonic
	}
	return ""
}

// LookupSNameRequest is Message for Verifying Request (Hmac Sha256)
type LookupSNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value to be verified
	SName string `protobuf:"bytes,1,opt,name=sName,proto3" json:"sName,omitempty"` // SName combined with Device ID and Hashed
}

func (x *LookupSNameRequest) Reset() {
	*x = LookupSNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_hdns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupSNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupSNameRequest) ProtoMessage() {}

func (x *LookupSNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_hdns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupSNameRequest.ProtoReflect.Descriptor instead.
func (*LookupSNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_protocols_hdns_proto_rawDescGZIP(), []int{2}
}

func (x *LookupSNameRequest) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

// LookupSNameResponse is Message for Verifying Response (Hmac Sha256)
type LookupSNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`    // If Values were Verified
	PublicKey string `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"` // Base64 Encoded Public Key
}

func (x *LookupSNameResponse) Reset() {
	*x = LookupSNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_hdns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupSNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupSNameResponse) ProtoMessage() {}

func (x *LookupSNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_hdns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupSNameResponse.ProtoReflect.Descriptor instead.
func (*LookupSNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_protocols_hdns_proto_rawDescGZIP(), []int{3}
}

func (x *LookupSNameResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LookupSNameResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

// VerifySNameRequest is Message for Verifying Request (Hmac Sha256)
type VerifySNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value to be verified
	SName    string `protobuf:"bytes,1,opt,name=sName,proto3" json:"sName,omitempty"`       // SName combined with Device ID and Hashed
	Mnemonic string `protobuf:"bytes,2,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"` // Mnemonic Hashed with public key for fingerprint
}

func (x *VerifySNameRequest) Reset() {
	*x = VerifySNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_hdns_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifySNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifySNameRequest) ProtoMessage() {}

func (x *VerifySNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_hdns_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifySNameRequest.ProtoReflect.Descriptor instead.
func (*VerifySNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_protocols_hdns_proto_rawDescGZIP(), []int{4}
}

func (x *VerifySNameRequest) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *VerifySNameRequest) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

// VerifySNameResponse is Message for Verifying Response (Hmac Sha256)
type VerifySNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // If Values were Verified
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`      // Error Message
	ShortID string `protobuf:"bytes,3,opt,name=shortID,proto3" json:"shortID,omitempty"`  // Short ID generated from HMac and Device
}

func (x *VerifySNameResponse) Reset() {
	*x = VerifySNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_hdns_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifySNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifySNameResponse) ProtoMessage() {}

func (x *VerifySNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_hdns_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifySNameResponse.ProtoReflect.Descriptor instead.
func (*VerifySNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_protocols_hdns_proto_rawDescGZIP(), []int{5}
}

func (x *VerifySNameResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *VerifySNameResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *VerifySNameResponse) GetShortID() string {
	if x != nil {
		return x.ShortID
	}
	return ""
}

var File_proto_protocols_hdns_proto protoreflect.FileDescriptor

var file_proto_protocols_hdns_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x2f, 0x68, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x68, 0x64, 0x6e,
	0x73, 0x22, 0x46, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x22, 0xe5, 0x01, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x2c, 0x0a, 0x11, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x67,
	0x69, 0x76, 0x65, 0x6e, 0x53, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x53, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x67,
	0x69, 0x76, 0x65, 0x6e, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69,
	0x63, 0x22, 0x2a, 0x0a, 0x12, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a,
	0x13, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x46, 0x0a, 0x12,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6e, 0x65, 0x6d,
	0x6f, 0x6e, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6e, 0x65, 0x6d,
	0x6f, 0x6e, 0x69, 0x63, 0x22, 0x5f, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x49, 0x44, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x68, 0x64, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_protocols_hdns_proto_rawDescOnce sync.Once
	file_proto_protocols_hdns_proto_rawDescData = file_proto_protocols_hdns_proto_rawDesc
)

func file_proto_protocols_hdns_proto_rawDescGZIP() []byte {
	file_proto_protocols_hdns_proto_rawDescOnce.Do(func() {
		file_proto_protocols_hdns_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_protocols_hdns_proto_rawDescData)
	})
	return file_proto_protocols_hdns_proto_rawDescData
}

var file_proto_protocols_hdns_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_protocols_hdns_proto_goTypes = []interface{}{
	(*CreateSNameRequest)(nil),  // 0: sonr.protocols.hdns.CreateSNameRequest
	(*CreateSNameResponse)(nil), // 1: sonr.protocols.hdns.CreateSNameResponse
	(*LookupSNameRequest)(nil),  // 2: sonr.protocols.hdns.LookupSNameRequest
	(*LookupSNameResponse)(nil), // 3: sonr.protocols.hdns.LookupSNameResponse
	(*VerifySNameRequest)(nil),  // 4: sonr.protocols.hdns.VerifySNameRequest
	(*VerifySNameResponse)(nil), // 5: sonr.protocols.hdns.VerifySNameResponse
}
var file_proto_protocols_hdns_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_protocols_hdns_proto_init() }
func file_proto_protocols_hdns_proto_init() {
	if File_proto_protocols_hdns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_protocols_hdns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSNameRequest); i {
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
		file_proto_protocols_hdns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSNameResponse); i {
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
		file_proto_protocols_hdns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupSNameRequest); i {
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
		file_proto_protocols_hdns_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupSNameResponse); i {
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
		file_proto_protocols_hdns_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifySNameRequest); i {
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
		file_proto_protocols_hdns_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifySNameResponse); i {
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
			RawDescriptor: file_proto_protocols_hdns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_protocols_hdns_proto_goTypes,
		DependencyIndexes: file_proto_protocols_hdns_proto_depIdxs,
		MessageInfos:      file_proto_protocols_hdns_proto_msgTypes,
	}.Build()
	File_proto_protocols_hdns_proto = out.File
	file_proto_protocols_hdns_proto_rawDesc = nil
	file_proto_protocols_hdns_proto_goTypes = nil
	file_proto_protocols_hdns_proto_depIdxs = nil
}
