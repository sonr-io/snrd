// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package vaultv1

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_DWN          protoreflect.MessageDescriptor
	fd_DWN_id       protoreflect.FieldDescriptor
	fd_DWN_alias    protoreflect.FieldDescriptor
	fd_DWN_cid      protoreflect.FieldDescriptor
	fd_DWN_resolver protoreflect.FieldDescriptor
)

func init() {
	file_vault_v1_state_proto_init()
	md_DWN = File_vault_v1_state_proto.Messages().ByName("DWN")
	fd_DWN_id = md_DWN.Fields().ByName("id")
	fd_DWN_alias = md_DWN.Fields().ByName("alias")
	fd_DWN_cid = md_DWN.Fields().ByName("cid")
	fd_DWN_resolver = md_DWN.Fields().ByName("resolver")
}

var _ protoreflect.Message = (*fastReflection_DWN)(nil)

type fastReflection_DWN DWN

func (x *DWN) ProtoReflect() protoreflect.Message {
	return (*fastReflection_DWN)(x)
}

func (x *DWN) slowProtoReflect() protoreflect.Message {
	mi := &file_vault_v1_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_DWN_messageType fastReflection_DWN_messageType
var _ protoreflect.MessageType = fastReflection_DWN_messageType{}

type fastReflection_DWN_messageType struct{}

func (x fastReflection_DWN_messageType) Zero() protoreflect.Message {
	return (*fastReflection_DWN)(nil)
}
func (x fastReflection_DWN_messageType) New() protoreflect.Message {
	return new(fastReflection_DWN)
}
func (x fastReflection_DWN_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_DWN
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_DWN) Descriptor() protoreflect.MessageDescriptor {
	return md_DWN
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_DWN) Type() protoreflect.MessageType {
	return _fastReflection_DWN_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_DWN) New() protoreflect.Message {
	return new(fastReflection_DWN)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_DWN) Interface() protoreflect.ProtoMessage {
	return (*DWN)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_DWN) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_DWN_id, value) {
			return
		}
	}
	if x.Alias != "" {
		value := protoreflect.ValueOfString(x.Alias)
		if !f(fd_DWN_alias, value) {
			return
		}
	}
	if x.Cid != "" {
		value := protoreflect.ValueOfString(x.Cid)
		if !f(fd_DWN_cid, value) {
			return
		}
	}
	if x.Resolver != "" {
		value := protoreflect.ValueOfString(x.Resolver)
		if !f(fd_DWN_resolver, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_DWN) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "vault.v1.DWN.id":
		return x.Id != uint64(0)
	case "vault.v1.DWN.alias":
		return x.Alias != ""
	case "vault.v1.DWN.cid":
		return x.Cid != ""
	case "vault.v1.DWN.resolver":
		return x.Resolver != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: vault.v1.DWN"))
		}
		panic(fmt.Errorf("message vault.v1.DWN does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DWN) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "vault.v1.DWN.id":
		x.Id = uint64(0)
	case "vault.v1.DWN.alias":
		x.Alias = ""
	case "vault.v1.DWN.cid":
		x.Cid = ""
	case "vault.v1.DWN.resolver":
		x.Resolver = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: vault.v1.DWN"))
		}
		panic(fmt.Errorf("message vault.v1.DWN does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_DWN) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "vault.v1.DWN.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "vault.v1.DWN.alias":
		value := x.Alias
		return protoreflect.ValueOfString(value)
	case "vault.v1.DWN.cid":
		value := x.Cid
		return protoreflect.ValueOfString(value)
	case "vault.v1.DWN.resolver":
		value := x.Resolver
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: vault.v1.DWN"))
		}
		panic(fmt.Errorf("message vault.v1.DWN does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DWN) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "vault.v1.DWN.id":
		x.Id = value.Uint()
	case "vault.v1.DWN.alias":
		x.Alias = value.Interface().(string)
	case "vault.v1.DWN.cid":
		x.Cid = value.Interface().(string)
	case "vault.v1.DWN.resolver":
		x.Resolver = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: vault.v1.DWN"))
		}
		panic(fmt.Errorf("message vault.v1.DWN does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DWN) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "vault.v1.DWN.id":
		panic(fmt.Errorf("field id of message vault.v1.DWN is not mutable"))
	case "vault.v1.DWN.alias":
		panic(fmt.Errorf("field alias of message vault.v1.DWN is not mutable"))
	case "vault.v1.DWN.cid":
		panic(fmt.Errorf("field cid of message vault.v1.DWN is not mutable"))
	case "vault.v1.DWN.resolver":
		panic(fmt.Errorf("field resolver of message vault.v1.DWN is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: vault.v1.DWN"))
		}
		panic(fmt.Errorf("message vault.v1.DWN does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_DWN) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "vault.v1.DWN.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "vault.v1.DWN.alias":
		return protoreflect.ValueOfString("")
	case "vault.v1.DWN.cid":
		return protoreflect.ValueOfString("")
	case "vault.v1.DWN.resolver":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: vault.v1.DWN"))
		}
		panic(fmt.Errorf("message vault.v1.DWN does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_DWN) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in vault.v1.DWN", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_DWN) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_DWN) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_DWN) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_DWN) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*DWN)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.Alias)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Cid)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Resolver)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*DWN)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Resolver) > 0 {
			i -= len(x.Resolver)
			copy(dAtA[i:], x.Resolver)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Resolver)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Cid) > 0 {
			i -= len(x.Cid)
			copy(dAtA[i:], x.Cid)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Cid)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Alias) > 0 {
			i -= len(x.Alias)
			copy(dAtA[i:], x.Alias)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Alias)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*DWN)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DWN: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: DWN: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Alias", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Alias = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Cid = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Resolver", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Resolver = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: vault/v1/state.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DWN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Alias    string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	Cid      string `protobuf:"bytes,3,opt,name=cid,proto3" json:"cid,omitempty"`
	Resolver string `protobuf:"bytes,4,opt,name=resolver,proto3" json:"resolver,omitempty"`
}

func (x *DWN) Reset() {
	*x = DWN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_v1_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DWN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DWN) ProtoMessage() {}

// Deprecated: Use DWN.ProtoReflect.Descriptor instead.
func (*DWN) Descriptor() ([]byte, []int) {
	return file_vault_v1_state_proto_rawDescGZIP(), []int{0}
}

func (x *DWN) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DWN) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *DWN) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *DWN) GetResolver() string {
	if x != nil {
		return x.Resolver
	}
	return ""
}

var File_vault_v1_state_proto protoreflect.FileDescriptor

var file_vault_v1_state_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x03, 0x44, 0x57,
	0x4e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x3a, 0x28, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x22, 0x0a, 0x06, 0x0a,
	0x02, 0x69, 0x64, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x10, 0x01,
	0x18, 0x01, 0x12, 0x09, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x10, 0x02, 0x18, 0x01, 0x18, 0x01, 0x42,
	0x88, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31,
	0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x73, 0x6f, 0x6e,
	0x72, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58,
	0x58, 0xaa, 0x02, 0x08, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x56,
	0x61, 0x75, 0x6c, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x09, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_vault_v1_state_proto_rawDescOnce sync.Once
	file_vault_v1_state_proto_rawDescData = file_vault_v1_state_proto_rawDesc
)

func file_vault_v1_state_proto_rawDescGZIP() []byte {
	file_vault_v1_state_proto_rawDescOnce.Do(func() {
		file_vault_v1_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_vault_v1_state_proto_rawDescData)
	})
	return file_vault_v1_state_proto_rawDescData
}

var file_vault_v1_state_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_vault_v1_state_proto_goTypes = []interface{}{
	(*DWN)(nil), // 0: vault.v1.DWN
}
var file_vault_v1_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vault_v1_state_proto_init() }
func file_vault_v1_state_proto_init() {
	if File_vault_v1_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vault_v1_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DWN); i {
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
			RawDescriptor: file_vault_v1_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vault_v1_state_proto_goTypes,
		DependencyIndexes: file_vault_v1_state_proto_depIdxs,
		MessageInfos:      file_vault_v1_state_proto_msgTypes,
	}.Build()
	File_vault_v1_state_proto = out.File
	file_vault_v1_state_proto_rawDesc = nil
	file_vault_v1_state_proto_goTypes = nil
	file_vault_v1_state_proto_depIdxs = nil
}
