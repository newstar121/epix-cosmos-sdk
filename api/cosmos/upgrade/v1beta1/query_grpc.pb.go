// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cosmos/upgrade/v1beta1/query.proto

package upgradev1beta1

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
	Query_CurrentPlan_FullMethodName            = "/cosmos.upgrade.v1beta1.Query/CurrentPlan"
	Query_AppliedPlan_FullMethodName            = "/cosmos.upgrade.v1beta1.Query/AppliedPlan"
	Query_UpgradedConsensusState_FullMethodName = "/cosmos.upgrade.v1beta1.Query/UpgradedConsensusState"
	Query_ModuleVersions_FullMethodName         = "/cosmos.upgrade.v1beta1.Query/ModuleVersions"
	Query_Authority_FullMethodName              = "/cosmos.upgrade.v1beta1.Query/Authority"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// CurrentPlan queries the current upgrade plan.
	CurrentPlan(ctx context.Context, in *QueryCurrentPlanRequest, opts ...grpc.CallOption) (*QueryCurrentPlanResponse, error)
	// AppliedPlan queries a previously applied upgrade plan by its name.
	AppliedPlan(ctx context.Context, in *QueryAppliedPlanRequest, opts ...grpc.CallOption) (*QueryAppliedPlanResponse, error)
	// Deprecated: Do not use.
	// UpgradedConsensusState queries the consensus state that will serve
	// as a trusted kernel for the next version of this chain. It will only be
	// stored at the last height of this chain.
	// UpgradedConsensusState RPC not supported with legacy querier
	// This rpc is deprecated now that IBC has its own replacement
	// (https://github.com/cosmos/ibc-go/blob/2c880a22e9f9cc75f62b527ca94aa75ce1106001/proto/ibc/core/client/v1/query.proto#L54)
	UpgradedConsensusState(ctx context.Context, in *QueryUpgradedConsensusStateRequest, opts ...grpc.CallOption) (*QueryUpgradedConsensusStateResponse, error)
	// ModuleVersions queries the list of module versions from state.
	//
	// Since: cosmos-sdk 0.43
	ModuleVersions(ctx context.Context, in *QueryModuleVersionsRequest, opts ...grpc.CallOption) (*QueryModuleVersionsResponse, error)
	// Returns the account with authority to conduct upgrades
	//
	// Since: cosmos-sdk 0.46
	Authority(ctx context.Context, in *QueryAuthorityRequest, opts ...grpc.CallOption) (*QueryAuthorityResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CurrentPlan(ctx context.Context, in *QueryCurrentPlanRequest, opts ...grpc.CallOption) (*QueryCurrentPlanResponse, error) {
	out := new(QueryCurrentPlanResponse)
	err := c.cc.Invoke(ctx, Query_CurrentPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AppliedPlan(ctx context.Context, in *QueryAppliedPlanRequest, opts ...grpc.CallOption) (*QueryAppliedPlanResponse, error) {
	out := new(QueryAppliedPlanResponse)
	err := c.cc.Invoke(ctx, Query_AppliedPlan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *queryClient) UpgradedConsensusState(ctx context.Context, in *QueryUpgradedConsensusStateRequest, opts ...grpc.CallOption) (*QueryUpgradedConsensusStateResponse, error) {
	out := new(QueryUpgradedConsensusStateResponse)
	err := c.cc.Invoke(ctx, Query_UpgradedConsensusState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ModuleVersions(ctx context.Context, in *QueryModuleVersionsRequest, opts ...grpc.CallOption) (*QueryModuleVersionsResponse, error) {
	out := new(QueryModuleVersionsResponse)
	err := c.cc.Invoke(ctx, Query_ModuleVersions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Authority(ctx context.Context, in *QueryAuthorityRequest, opts ...grpc.CallOption) (*QueryAuthorityResponse, error) {
	out := new(QueryAuthorityResponse)
	err := c.cc.Invoke(ctx, Query_Authority_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// CurrentPlan queries the current upgrade plan.
	CurrentPlan(context.Context, *QueryCurrentPlanRequest) (*QueryCurrentPlanResponse, error)
	// AppliedPlan queries a previously applied upgrade plan by its name.
	AppliedPlan(context.Context, *QueryAppliedPlanRequest) (*QueryAppliedPlanResponse, error)
	// Deprecated: Do not use.
	// UpgradedConsensusState queries the consensus state that will serve
	// as a trusted kernel for the next version of this chain. It will only be
	// stored at the last height of this chain.
	// UpgradedConsensusState RPC not supported with legacy querier
	// This rpc is deprecated now that IBC has its own replacement
	// (https://github.com/cosmos/ibc-go/blob/2c880a22e9f9cc75f62b527ca94aa75ce1106001/proto/ibc/core/client/v1/query.proto#L54)
	UpgradedConsensusState(context.Context, *QueryUpgradedConsensusStateRequest) (*QueryUpgradedConsensusStateResponse, error)
	// ModuleVersions queries the list of module versions from state.
	//
	// Since: cosmos-sdk 0.43
	ModuleVersions(context.Context, *QueryModuleVersionsRequest) (*QueryModuleVersionsResponse, error)
	// Returns the account with authority to conduct upgrades
	//
	// Since: cosmos-sdk 0.46
	Authority(context.Context, *QueryAuthorityRequest) (*QueryAuthorityResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) CurrentPlan(context.Context, *QueryCurrentPlanRequest) (*QueryCurrentPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentPlan not implemented")
}
func (UnimplementedQueryServer) AppliedPlan(context.Context, *QueryAppliedPlanRequest) (*QueryAppliedPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppliedPlan not implemented")
}
func (UnimplementedQueryServer) UpgradedConsensusState(context.Context, *QueryUpgradedConsensusStateRequest) (*QueryUpgradedConsensusStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradedConsensusState not implemented")
}
func (UnimplementedQueryServer) ModuleVersions(context.Context, *QueryModuleVersionsRequest) (*QueryModuleVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModuleVersions not implemented")
}
func (UnimplementedQueryServer) Authority(context.Context, *QueryAuthorityRequest) (*QueryAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authority not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_CurrentPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCurrentPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CurrentPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_CurrentPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CurrentPlan(ctx, req.(*QueryCurrentPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AppliedPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAppliedPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AppliedPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AppliedPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AppliedPlan(ctx, req.(*QueryAppliedPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UpgradedConsensusState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUpgradedConsensusStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UpgradedConsensusState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_UpgradedConsensusState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UpgradedConsensusState(ctx, req.(*QueryUpgradedConsensusStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ModuleVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryModuleVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ModuleVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ModuleVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ModuleVersions(ctx, req.(*QueryModuleVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Authority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Authority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Authority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Authority(ctx, req.(*QueryAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.upgrade.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CurrentPlan",
			Handler:    _Query_CurrentPlan_Handler,
		},
		{
			MethodName: "AppliedPlan",
			Handler:    _Query_AppliedPlan_Handler,
		},
		{
			MethodName: "UpgradedConsensusState",
			Handler:    _Query_UpgradedConsensusState_Handler,
		},
		{
			MethodName: "ModuleVersions",
			Handler:    _Query_ModuleVersions_Handler,
		},
		{
			MethodName: "Authority",
			Handler:    _Query_Authority_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/upgrade/v1beta1/query.proto",
}
