// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: sonrhq/sonr/identity/v1/identifier.proto

package identityv1

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

// UserIdentifierType represents the type of a user identifier.
type UserIdentifierType int32

const (
	// The type of identifier is unknown.
	UserIdentifierType_USER_IDENTIFIER_TYPE_UNSPECIFIED UserIdentifierType = 0
	// The type of identifier is an email address.
	UserIdentifierType_USER_IDENTIFIER_TYPE_EMAIL UserIdentifierType = 1
	// The type of identifier is a phone number.
	UserIdentifierType_USER_IDENTIFIER_TYPE_PHONE UserIdentifierType = 2
	// The type of identifier is a FIDO credential ID.
	UserIdentifierType_USER_IDENTIFIER_TYPE_FIDO UserIdentifierType = 3
	// The type of a Passkey identifier.
	UserIdentifierType_USER_IDENTIFIER_TYPE_PASSKEY UserIdentifierType = 4
	// The type of a public key identifier.
	UserIdentifierType_USER_IDENTIFIER_TYPE_PUBLIC_KEY UserIdentifierType = 5
)

// Enum value maps for UserIdentifierType.
var (
	UserIdentifierType_name = map[int32]string{
		0: "USER_IDENTIFIER_TYPE_UNSPECIFIED",
		1: "USER_IDENTIFIER_TYPE_EMAIL",
		2: "USER_IDENTIFIER_TYPE_PHONE",
		3: "USER_IDENTIFIER_TYPE_FIDO",
		4: "USER_IDENTIFIER_TYPE_PASSKEY",
		5: "USER_IDENTIFIER_TYPE_PUBLIC_KEY",
	}
	UserIdentifierType_value = map[string]int32{
		"USER_IDENTIFIER_TYPE_UNSPECIFIED": 0,
		"USER_IDENTIFIER_TYPE_EMAIL":       1,
		"USER_IDENTIFIER_TYPE_PHONE":       2,
		"USER_IDENTIFIER_TYPE_FIDO":        3,
		"USER_IDENTIFIER_TYPE_PASSKEY":     4,
		"USER_IDENTIFIER_TYPE_PUBLIC_KEY":  5,
	}
)

func (x UserIdentifierType) Enum() *UserIdentifierType {
	p := new(UserIdentifierType)
	*p = x
	return p
}

func (x UserIdentifierType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserIdentifierType) Descriptor() protoreflect.EnumDescriptor {
	return file_sonrhq_sonr_identity_v1_identifier_proto_enumTypes[0].Descriptor()
}

func (UserIdentifierType) Type() protoreflect.EnumType {
	return &file_sonrhq_sonr_identity_v1_identifier_proto_enumTypes[0]
}

func (x UserIdentifierType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserIdentifierType.Descriptor instead.
func (UserIdentifierType) EnumDescriptor() ([]byte, []int) {
	return file_sonrhq_sonr_identity_v1_identifier_proto_rawDescGZIP(), []int{0}
}

// UserIdentifier represents the structure of a user identifier.
type UserIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of identifier.
	IdentifierType UserIdentifierType `protobuf:"varint,1,opt,name=identifier_type,json=identifierType,proto3,enum=sonrhq.sonr.identity.v1.UserIdentifierType" json:"identifier_type,omitempty"`
	// The value of the identifier.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UserIdentifier) Reset() {
	*x = UserIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonrhq_sonr_identity_v1_identifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdentifier) ProtoMessage() {}

func (x *UserIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_sonrhq_sonr_identity_v1_identifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdentifier.ProtoReflect.Descriptor instead.
func (*UserIdentifier) Descriptor() ([]byte, []int) {
	return file_sonrhq_sonr_identity_v1_identifier_proto_rawDescGZIP(), []int{0}
}

func (x *UserIdentifier) GetIdentifierType() UserIdentifierType {
	if x != nil {
		return x.IdentifierType
	}
	return UserIdentifierType_USER_IDENTIFIER_TYPE_UNSPECIFIED
}

func (x *UserIdentifier) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// WalletIdentifier represents the structure of a wallet identifier.
type WalletIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bip44 coin type.
	CoinType uint32 `protobuf:"varint,1,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
	// The value of the identifier is the address of the wallet.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *WalletIdentifier) Reset() {
	*x = WalletIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonrhq_sonr_identity_v1_identifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletIdentifier) ProtoMessage() {}

func (x *WalletIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_sonrhq_sonr_identity_v1_identifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletIdentifier.ProtoReflect.Descriptor instead.
func (*WalletIdentifier) Descriptor() ([]byte, []int) {
	return file_sonrhq_sonr_identity_v1_identifier_proto_rawDescGZIP(), []int{1}
}

func (x *WalletIdentifier) GetCoinType() uint32 {
	if x != nil {
		return x.CoinType
	}
	return 0
}

func (x *WalletIdentifier) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_sonrhq_sonr_identity_v1_identifier_proto protoreflect.FileDescriptor

var file_sonrhq_sonr_identity_v1_identifier_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x6f, 0x6e, 0x72,
	0x68, 0x71, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x22, 0x7c, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x0f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x45, 0x0a, 0x10, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0xe0, 0x01, 0x0a, 0x12, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x24, 0x0a, 0x20, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49,
	0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44,
	0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4d,
	0x41, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44,
	0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x48,
	0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44,
	0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49,
	0x44, 0x4f, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x45,
	0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x53,
	0x53, 0x4b, 0x45, 0x59, 0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49,
	0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x55, 0x42, 0x4c, 0x49, 0x43, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x05, 0x42, 0xf7, 0x01, 0x0a, 0x1b,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x68,
	0x71, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x78, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2f, 0x73, 0x6f, 0x6e,
	0x72, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x49, 0xaa, 0x02,
	0x17, 0x53, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2e, 0x53, 0x6f, 0x6e, 0x72, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x53, 0x6f, 0x6e, 0x72, 0x68,
	0x71, 0x5c, 0x53, 0x6f, 0x6e, 0x72, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x23, 0x53, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x5c, 0x53, 0x6f, 0x6e, 0x72,
	0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x53, 0x6f, 0x6e, 0x72, 0x68,
	0x71, 0x3a, 0x3a, 0x53, 0x6f, 0x6e, 0x72, 0x3a, 0x3a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sonrhq_sonr_identity_v1_identifier_proto_rawDescOnce sync.Once
	file_sonrhq_sonr_identity_v1_identifier_proto_rawDescData = file_sonrhq_sonr_identity_v1_identifier_proto_rawDesc
)

func file_sonrhq_sonr_identity_v1_identifier_proto_rawDescGZIP() []byte {
	file_sonrhq_sonr_identity_v1_identifier_proto_rawDescOnce.Do(func() {
		file_sonrhq_sonr_identity_v1_identifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_sonrhq_sonr_identity_v1_identifier_proto_rawDescData)
	})
	return file_sonrhq_sonr_identity_v1_identifier_proto_rawDescData
}

var file_sonrhq_sonr_identity_v1_identifier_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sonrhq_sonr_identity_v1_identifier_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sonrhq_sonr_identity_v1_identifier_proto_goTypes = []interface{}{
	(UserIdentifierType)(0),  // 0: sonrhq.sonr.identity.v1.UserIdentifierType
	(*UserIdentifier)(nil),   // 1: sonrhq.sonr.identity.v1.UserIdentifier
	(*WalletIdentifier)(nil), // 2: sonrhq.sonr.identity.v1.WalletIdentifier
}
var file_sonrhq_sonr_identity_v1_identifier_proto_depIdxs = []int32{
	0, // 0: sonrhq.sonr.identity.v1.UserIdentifier.identifier_type:type_name -> sonrhq.sonr.identity.v1.UserIdentifierType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sonrhq_sonr_identity_v1_identifier_proto_init() }
func file_sonrhq_sonr_identity_v1_identifier_proto_init() {
	if File_sonrhq_sonr_identity_v1_identifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sonrhq_sonr_identity_v1_identifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdentifier); i {
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
		file_sonrhq_sonr_identity_v1_identifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletIdentifier); i {
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
			RawDescriptor: file_sonrhq_sonr_identity_v1_identifier_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sonrhq_sonr_identity_v1_identifier_proto_goTypes,
		DependencyIndexes: file_sonrhq_sonr_identity_v1_identifier_proto_depIdxs,
		EnumInfos:         file_sonrhq_sonr_identity_v1_identifier_proto_enumTypes,
		MessageInfos:      file_sonrhq_sonr_identity_v1_identifier_proto_msgTypes,
	}.Build()
	File_sonrhq_sonr_identity_v1_identifier_proto = out.File
	file_sonrhq_sonr_identity_v1_identifier_proto_rawDesc = nil
	file_sonrhq_sonr_identity_v1_identifier_proto_goTypes = nil
	file_sonrhq_sonr_identity_v1_identifier_proto_depIdxs = nil
}
