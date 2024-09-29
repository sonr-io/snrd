// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: did/v1/tx.proto

package didv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_ExecuteTx_FullMethodName          = "/did.v1.Msg/ExecuteTx"
	Msg_RegisterController_FullMethodName = "/did.v1.Msg/RegisterController"
	Msg_UpdateParams_FullMethodName       = "/did.v1.Msg/UpdateParams"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// ExecuteTx executes a transaction on the Sonr Blockchain. It leverages
	// Macaroon for verification.
	ExecuteTx(ctx context.Context, in *MsgExecuteTx, opts ...grpc.CallOption) (*MsgExecuteTxResponse, error)
	// RegisterController initializes a controller with the given authentication
	// set, address, cid, publicKey, and user-defined alias.
	RegisterController(ctx context.Context, in *MsgRegisterController, opts ...grpc.CallOption) (*MsgRegisterControllerResponse, error)
	// UpdateParams defines a governance operation for updating the parameters.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ExecuteTx(ctx context.Context, in *MsgExecuteTx, opts ...grpc.CallOption) (*MsgExecuteTxResponse, error) {
	out := new(MsgExecuteTxResponse)
	err := c.cc.Invoke(ctx, Msg_ExecuteTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegisterController(ctx context.Context, in *MsgRegisterController, opts ...grpc.CallOption) (*MsgRegisterControllerResponse, error) {
	out := new(MsgRegisterControllerResponse)
	err := c.cc.Invoke(ctx, Msg_RegisterController_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// ExecuteTx executes a transaction on the Sonr Blockchain. It leverages
	// Macaroon for verification.
	ExecuteTx(context.Context, *MsgExecuteTx) (*MsgExecuteTxResponse, error)
	// RegisterController initializes a controller with the given authentication
	// set, address, cid, publicKey, and user-defined alias.
	RegisterController(context.Context, *MsgRegisterController) (*MsgRegisterControllerResponse, error)
	// UpdateParams defines a governance operation for updating the parameters.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) ExecuteTx(context.Context, *MsgExecuteTx) (*MsgExecuteTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteTx not implemented")
}
func (UnimplementedMsgServer) RegisterController(context.Context, *MsgRegisterController) (*MsgRegisterControllerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterController not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_ExecuteTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecuteTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExecuteTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ExecuteTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExecuteTx(ctx, req.(*MsgExecuteTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegisterController_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterController)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterController(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RegisterController_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterController(ctx, req.(*MsgRegisterController))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "did.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteTx",
			Handler:    _Msg_ExecuteTx_Handler,
		},
		{
			MethodName: "RegisterController",
			Handler:    _Msg_RegisterController_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "did/v1/tx.proto",
}
