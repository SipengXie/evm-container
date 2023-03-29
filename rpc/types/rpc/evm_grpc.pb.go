// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.1
// source: evm.proto

package rpc

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
	Rpc_NewEnv_FullMethodName          = "/rpc.Rpc/NewEnv"
	Rpc_SetBlockContext_FullMethodName = "/rpc.Rpc/SetBlockContext"
	Rpc_Reset_FullMethodName           = "/rpc.Rpc/Reset"
	Rpc_Cancel_FullMethodName          = "/rpc.Rpc/Cancel"
	Rpc_Cancelled_FullMethodName       = "/rpc.Rpc/Cancelled"
	Rpc_Call_FullMethodName            = "/rpc.Rpc/Call"
	Rpc_DelegateCall_FullMethodName    = "/rpc.Rpc/DelegateCall"
	Rpc_StataicCall_FullMethodName     = "/rpc.Rpc/StataicCall"
	Rpc_Create_FullMethodName          = "/rpc.Rpc/Create"
	Rpc_Create2_FullMethodName         = "/rpc.Rpc/Create2"
	Rpc_ChainConfig_FullMethodName     = "/rpc.Rpc/ChainConfig"
)

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	NewEnv(ctx context.Context, in *NewEnvRequest, opts ...grpc.CallOption) (*NewEnvResponse, error)
	SetBlockContext(ctx context.Context, in *SetBlockContextRequest, opts ...grpc.CallOption) (*SetBlockContextResponse, error)
	Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error)
	Cancel(ctx context.Context, in *CancelRequset, opts ...grpc.CallOption) (*CancelResponse, error)
	Cancelled(ctx context.Context, in *CancelledRequest, opts ...grpc.CallOption) (*CancelledResponse, error)
	Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
	DelegateCall(ctx context.Context, in *DelegateCallRequest, opts ...grpc.CallOption) (*DelegateCallResponse, error)
	StataicCall(ctx context.Context, in *StataicCallRequest, opts ...grpc.CallOption) (*StataicCallResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Create2(ctx context.Context, in *Create2Request, opts ...grpc.CallOption) (*Create2Response, error)
	ChainConfig(ctx context.Context, in *ChainConfigRequest, opts ...grpc.CallOption) (*ChainConfigResponse, error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) NewEnv(ctx context.Context, in *NewEnvRequest, opts ...grpc.CallOption) (*NewEnvResponse, error) {
	out := new(NewEnvResponse)
	err := c.cc.Invoke(ctx, Rpc_NewEnv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) SetBlockContext(ctx context.Context, in *SetBlockContextRequest, opts ...grpc.CallOption) (*SetBlockContextResponse, error) {
	out := new(SetBlockContextResponse)
	err := c.cc.Invoke(ctx, Rpc_SetBlockContext_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error) {
	out := new(ResetResponse)
	err := c.cc.Invoke(ctx, Rpc_Reset_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Cancel(ctx context.Context, in *CancelRequset, opts ...grpc.CallOption) (*CancelResponse, error) {
	out := new(CancelResponse)
	err := c.cc.Invoke(ctx, Rpc_Cancel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Cancelled(ctx context.Context, in *CancelledRequest, opts ...grpc.CallOption) (*CancelledResponse, error) {
	out := new(CancelledResponse)
	err := c.cc.Invoke(ctx, Rpc_Cancelled_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, Rpc_Call_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) DelegateCall(ctx context.Context, in *DelegateCallRequest, opts ...grpc.CallOption) (*DelegateCallResponse, error) {
	out := new(DelegateCallResponse)
	err := c.cc.Invoke(ctx, Rpc_DelegateCall_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) StataicCall(ctx context.Context, in *StataicCallRequest, opts ...grpc.CallOption) (*StataicCallResponse, error) {
	out := new(StataicCallResponse)
	err := c.cc.Invoke(ctx, Rpc_StataicCall_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Rpc_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Create2(ctx context.Context, in *Create2Request, opts ...grpc.CallOption) (*Create2Response, error) {
	out := new(Create2Response)
	err := c.cc.Invoke(ctx, Rpc_Create2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) ChainConfig(ctx context.Context, in *ChainConfigRequest, opts ...grpc.CallOption) (*ChainConfigResponse, error) {
	out := new(ChainConfigResponse)
	err := c.cc.Invoke(ctx, Rpc_ChainConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServer is the server API for Rpc service.
// All implementations must embed UnimplementedRpcServer
// for forward compatibility
type RpcServer interface {
	NewEnv(context.Context, *NewEnvRequest) (*NewEnvResponse, error)
	SetBlockContext(context.Context, *SetBlockContextRequest) (*SetBlockContextResponse, error)
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)
	Cancel(context.Context, *CancelRequset) (*CancelResponse, error)
	Cancelled(context.Context, *CancelledRequest) (*CancelledResponse, error)
	Call(context.Context, *CallRequest) (*CallResponse, error)
	DelegateCall(context.Context, *DelegateCallRequest) (*DelegateCallResponse, error)
	StataicCall(context.Context, *StataicCallRequest) (*StataicCallResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Create2(context.Context, *Create2Request) (*Create2Response, error)
	ChainConfig(context.Context, *ChainConfigRequest) (*ChainConfigResponse, error)
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (UnimplementedRpcServer) NewEnv(context.Context, *NewEnvRequest) (*NewEnvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewEnv not implemented")
}
func (UnimplementedRpcServer) SetBlockContext(context.Context, *SetBlockContextRequest) (*SetBlockContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBlockContext not implemented")
}
func (UnimplementedRpcServer) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}
func (UnimplementedRpcServer) Cancel(context.Context, *CancelRequset) (*CancelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedRpcServer) Cancelled(context.Context, *CancelledRequest) (*CancelledResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancelled not implemented")
}
func (UnimplementedRpcServer) Call(context.Context, *CallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedRpcServer) DelegateCall(context.Context, *DelegateCallRequest) (*DelegateCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegateCall not implemented")
}
func (UnimplementedRpcServer) StataicCall(context.Context, *StataicCallRequest) (*StataicCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StataicCall not implemented")
}
func (UnimplementedRpcServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRpcServer) Create2(context.Context, *Create2Request) (*Create2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create2 not implemented")
}
func (UnimplementedRpcServer) ChainConfig(context.Context, *ChainConfigRequest) (*ChainConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainConfig not implemented")
}
func (UnimplementedRpcServer) mustEmbedUnimplementedRpcServer() {}

// UnsafeRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcServer will
// result in compilation errors.
type UnsafeRpcServer interface {
	mustEmbedUnimplementedRpcServer()
}

func RegisterRpcServer(s grpc.ServiceRegistrar, srv RpcServer) {
	s.RegisterService(&Rpc_ServiceDesc, srv)
}

func _Rpc_NewEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).NewEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_NewEnv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).NewEnv(ctx, req.(*NewEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_SetBlockContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBlockContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).SetBlockContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_SetBlockContext_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).SetBlockContext(ctx, req.(*SetBlockContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_Reset_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Reset(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelRequset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_Cancel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Cancel(ctx, req.(*CancelRequset))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Cancelled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Cancelled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_Cancelled_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Cancelled(ctx, req.(*CancelledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_Call_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Call(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_DelegateCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelegateCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).DelegateCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_DelegateCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).DelegateCall(ctx, req.(*DelegateCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_StataicCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StataicCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).StataicCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_StataicCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).StataicCall(ctx, req.(*StataicCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Create2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Create2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Create2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_Create2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Create2(ctx, req.(*Create2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_ChainConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).ChainConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_ChainConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).ChainConfig(ctx, req.(*ChainConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewEnv",
			Handler:    _Rpc_NewEnv_Handler,
		},
		{
			MethodName: "SetBlockContext",
			Handler:    _Rpc_SetBlockContext_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _Rpc_Reset_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Rpc_Cancel_Handler,
		},
		{
			MethodName: "Cancelled",
			Handler:    _Rpc_Cancelled_Handler,
		},
		{
			MethodName: "Call",
			Handler:    _Rpc_Call_Handler,
		},
		{
			MethodName: "DelegateCall",
			Handler:    _Rpc_DelegateCall_Handler,
		},
		{
			MethodName: "StataicCall",
			Handler:    _Rpc_StataicCall_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Rpc_Create_Handler,
		},
		{
			MethodName: "Create2",
			Handler:    _Rpc_Create2_Handler,
		},
		{
			MethodName: "ChainConfig",
			Handler:    _Rpc_ChainConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evm.proto",
}
