// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: protobuf/bridge.proto

package pb

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

// BridgeServiceClient is the client API for BridgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BridgeServiceClient interface {
	CrossChainTransfer(ctx context.Context, in *CrossChainTransferRequest, opts ...grpc.CallOption) (*CrossChainTransferResponse, error)
	ChangeTransferStatus(ctx context.Context, in *CrossChainTransferStatusRequest, opts ...grpc.CallOption) (*CrossChainTransferStatusResponse, error)
	UpdateDepositFundingPoolBalance(ctx context.Context, in *UpdateDepositFundingPoolBalanceRequest, opts ...grpc.CallOption) (*UpdateDepositFundingPoolBalanceResponse, error)
	UpdateWithdrawFundingPoolBalance(ctx context.Context, in *UpdateWithdrawFundingPoolBalanceRequest, opts ...grpc.CallOption) (*UpdateWithdrawFundingPoolBalanceResponse, error)
	UnstakeBatch(ctx context.Context, in *UnstakeBatchRequest, opts ...grpc.CallOption) (*UnstakeBatchResponse, error)
	UnstakeSingle(ctx context.Context, in *UnstakeSingleRequest, opts ...grpc.CallOption) (*UnstakeSingleResponse, error)
}

type bridgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBridgeServiceClient(cc grpc.ClientConnInterface) BridgeServiceClient {
	return &bridgeServiceClient{cc}
}

func (c *bridgeServiceClient) CrossChainTransfer(ctx context.Context, in *CrossChainTransferRequest, opts ...grpc.CallOption) (*CrossChainTransferResponse, error) {
	out := new(CrossChainTransferResponse)
	err := c.cc.Invoke(ctx, "/selaginella.proto_rpc.BridgeService/crossChainTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeServiceClient) ChangeTransferStatus(ctx context.Context, in *CrossChainTransferStatusRequest, opts ...grpc.CallOption) (*CrossChainTransferStatusResponse, error) {
	out := new(CrossChainTransferStatusResponse)
	err := c.cc.Invoke(ctx, "/selaginella.proto_rpc.BridgeService/changeTransferStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeServiceClient) UpdateDepositFundingPoolBalance(ctx context.Context, in *UpdateDepositFundingPoolBalanceRequest, opts ...grpc.CallOption) (*UpdateDepositFundingPoolBalanceResponse, error) {
	out := new(UpdateDepositFundingPoolBalanceResponse)
	err := c.cc.Invoke(ctx, "/selaginella.proto_rpc.BridgeService/UpdateDepositFundingPoolBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeServiceClient) UpdateWithdrawFundingPoolBalance(ctx context.Context, in *UpdateWithdrawFundingPoolBalanceRequest, opts ...grpc.CallOption) (*UpdateWithdrawFundingPoolBalanceResponse, error) {
	out := new(UpdateWithdrawFundingPoolBalanceResponse)
	err := c.cc.Invoke(ctx, "/selaginella.proto_rpc.BridgeService/UpdateWithdrawFundingPoolBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeServiceClient) UnstakeBatch(ctx context.Context, in *UnstakeBatchRequest, opts ...grpc.CallOption) (*UnstakeBatchResponse, error) {
	out := new(UnstakeBatchResponse)
	err := c.cc.Invoke(ctx, "/selaginella.proto_rpc.BridgeService/UnstakeBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeServiceClient) UnstakeSingle(ctx context.Context, in *UnstakeSingleRequest, opts ...grpc.CallOption) (*UnstakeSingleResponse, error) {
	out := new(UnstakeSingleResponse)
	err := c.cc.Invoke(ctx, "/selaginella.proto_rpc.BridgeService/UnstakeSingle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BridgeServiceServer is the server API for BridgeService service.
// All implementations must embed UnimplementedBridgeServiceServer
// for forward compatibility
type BridgeServiceServer interface {
	CrossChainTransfer(context.Context, *CrossChainTransferRequest) (*CrossChainTransferResponse, error)
	ChangeTransferStatus(context.Context, *CrossChainTransferStatusRequest) (*CrossChainTransferStatusResponse, error)
	UpdateDepositFundingPoolBalance(context.Context, *UpdateDepositFundingPoolBalanceRequest) (*UpdateDepositFundingPoolBalanceResponse, error)
	UpdateWithdrawFundingPoolBalance(context.Context, *UpdateWithdrawFundingPoolBalanceRequest) (*UpdateWithdrawFundingPoolBalanceResponse, error)
	UnstakeBatch(context.Context, *UnstakeBatchRequest) (*UnstakeBatchResponse, error)
	UnstakeSingle(context.Context, *UnstakeSingleRequest) (*UnstakeSingleResponse, error)
}

// UnimplementedBridgeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBridgeServiceServer struct {
}

func (UnimplementedBridgeServiceServer) CrossChainTransfer(context.Context, *CrossChainTransferRequest) (*CrossChainTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CrossChainTransfer not implemented")
}
func (UnimplementedBridgeServiceServer) ChangeTransferStatus(context.Context, *CrossChainTransferStatusRequest) (*CrossChainTransferStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeTransferStatus not implemented")
}
func (UnimplementedBridgeServiceServer) UpdateDepositFundingPoolBalance(context.Context, *UpdateDepositFundingPoolBalanceRequest) (*UpdateDepositFundingPoolBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDepositFundingPoolBalance not implemented")
}
func (UnimplementedBridgeServiceServer) UpdateWithdrawFundingPoolBalance(context.Context, *UpdateWithdrawFundingPoolBalanceRequest) (*UpdateWithdrawFundingPoolBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWithdrawFundingPoolBalance not implemented")
}
func (UnimplementedBridgeServiceServer) UnstakeBatch(context.Context, *UnstakeBatchRequest) (*UnstakeBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnstakeBatch not implemented")
}
func (UnimplementedBridgeServiceServer) UnstakeSingle(context.Context, *UnstakeSingleRequest) (*UnstakeSingleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnstakeSingle not implemented")
}

// UnsafeBridgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BridgeServiceServer will
// result in compilation errors.
type UnsafeBridgeServiceServer interface {
	mustEmbedUnimplementedBridgeServiceServer()
}

func RegisterBridgeServiceServer(s grpc.ServiceRegistrar, srv BridgeServiceServer) {
	s.RegisterService(&BridgeService_ServiceDesc, srv)
}

func _BridgeService_CrossChainTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrossChainTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).CrossChainTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selaginella.proto_rpc.BridgeService/crossChainTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).CrossChainTransfer(ctx, req.(*CrossChainTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BridgeService_ChangeTransferStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrossChainTransferStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).ChangeTransferStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selaginella.proto_rpc.BridgeService/changeTransferStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).ChangeTransferStatus(ctx, req.(*CrossChainTransferStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BridgeService_UpdateDepositFundingPoolBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDepositFundingPoolBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).UpdateDepositFundingPoolBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selaginella.proto_rpc.BridgeService/UpdateDepositFundingPoolBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).UpdateDepositFundingPoolBalance(ctx, req.(*UpdateDepositFundingPoolBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BridgeService_UpdateWithdrawFundingPoolBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWithdrawFundingPoolBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).UpdateWithdrawFundingPoolBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selaginella.proto_rpc.BridgeService/UpdateWithdrawFundingPoolBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).UpdateWithdrawFundingPoolBalance(ctx, req.(*UpdateWithdrawFundingPoolBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BridgeService_UnstakeBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnstakeBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).UnstakeBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selaginella.proto_rpc.BridgeService/UnstakeBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).UnstakeBatch(ctx, req.(*UnstakeBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BridgeService_UnstakeSingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnstakeSingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).UnstakeSingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/selaginella.proto_rpc.BridgeService/UnstakeSingle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).UnstakeSingle(ctx, req.(*UnstakeSingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BridgeService_ServiceDesc is the grpc.ServiceDesc for BridgeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BridgeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "selaginella.proto_rpc.BridgeService",
	HandlerType: (*BridgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "crossChainTransfer",
			Handler:    _BridgeService_CrossChainTransfer_Handler,
		},
		{
			MethodName: "changeTransferStatus",
			Handler:    _BridgeService_ChangeTransferStatus_Handler,
		},
		{
			MethodName: "UpdateDepositFundingPoolBalance",
			Handler:    _BridgeService_UpdateDepositFundingPoolBalance_Handler,
		},
		{
			MethodName: "UpdateWithdrawFundingPoolBalance",
			Handler:    _BridgeService_UpdateWithdrawFundingPoolBalance_Handler,
		},
		{
			MethodName: "UnstakeBatch",
			Handler:    _BridgeService_UnstakeBatch_Handler,
		},
		{
			MethodName: "UnstakeSingle",
			Handler:    _BridgeService_UnstakeSingle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/bridge.proto",
}
