// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: api/rewards/rewards.proto

package rewards

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RewardsServiceClient is the client API for RewardsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RewardsServiceClient interface {
	GetTotalPoints(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTotalPointsResponse, error)
	GetQualifiedDevices(ctx context.Context, in *GetQualifiedDevicesRequest, opts ...grpc.CallOption) (RewardsService_GetQualifiedDevicesClient, error)
}

type rewardsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRewardsServiceClient(cc grpc.ClientConnInterface) RewardsServiceClient {
	return &rewardsServiceClient{cc}
}

func (c *rewardsServiceClient) GetTotalPoints(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetTotalPointsResponse, error) {
	out := new(GetTotalPointsResponse)
	err := c.cc.Invoke(ctx, "/users.RewardsService/GetTotalPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rewardsServiceClient) GetQualifiedDevices(ctx context.Context, in *GetQualifiedDevicesRequest, opts ...grpc.CallOption) (RewardsService_GetQualifiedDevicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &RewardsService_ServiceDesc.Streams[0], "/users.RewardsService/GetQualifiedDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &rewardsServiceGetQualifiedDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RewardsService_GetQualifiedDevicesClient interface {
	Recv() (*GetQualifiedDevicesDevice, error)
	grpc.ClientStream
}

type rewardsServiceGetQualifiedDevicesClient struct {
	grpc.ClientStream
}

func (x *rewardsServiceGetQualifiedDevicesClient) Recv() (*GetQualifiedDevicesDevice, error) {
	m := new(GetQualifiedDevicesDevice)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RewardsServiceServer is the server API for RewardsService service.
// All implementations must embed UnimplementedRewardsServiceServer
// for forward compatibility
type RewardsServiceServer interface {
	GetTotalPoints(context.Context, *emptypb.Empty) (*GetTotalPointsResponse, error)
	GetQualifiedDevices(*GetQualifiedDevicesRequest, RewardsService_GetQualifiedDevicesServer) error
	mustEmbedUnimplementedRewardsServiceServer()
}

// UnimplementedRewardsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRewardsServiceServer struct {
}

func (UnimplementedRewardsServiceServer) GetTotalPoints(context.Context, *emptypb.Empty) (*GetTotalPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalPoints not implemented")
}
func (UnimplementedRewardsServiceServer) GetQualifiedDevices(*GetQualifiedDevicesRequest, RewardsService_GetQualifiedDevicesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetQualifiedDevices not implemented")
}
func (UnimplementedRewardsServiceServer) mustEmbedUnimplementedRewardsServiceServer() {}

// UnsafeRewardsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RewardsServiceServer will
// result in compilation errors.
type UnsafeRewardsServiceServer interface {
	mustEmbedUnimplementedRewardsServiceServer()
}

func RegisterRewardsServiceServer(s grpc.ServiceRegistrar, srv RewardsServiceServer) {
	s.RegisterService(&RewardsService_ServiceDesc, srv)
}

func _RewardsService_GetTotalPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardsServiceServer).GetTotalPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.RewardsService/GetTotalPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardsServiceServer).GetTotalPoints(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RewardsService_GetQualifiedDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetQualifiedDevicesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RewardsServiceServer).GetQualifiedDevices(m, &rewardsServiceGetQualifiedDevicesServer{stream})
}

type RewardsService_GetQualifiedDevicesServer interface {
	Send(*GetQualifiedDevicesDevice) error
	grpc.ServerStream
}

type rewardsServiceGetQualifiedDevicesServer struct {
	grpc.ServerStream
}

func (x *rewardsServiceGetQualifiedDevicesServer) Send(m *GetQualifiedDevicesDevice) error {
	return x.ServerStream.SendMsg(m)
}

// RewardsService_ServiceDesc is the grpc.ServiceDesc for RewardsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RewardsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.RewardsService",
	HandlerType: (*RewardsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTotalPoints",
			Handler:    _RewardsService_GetTotalPoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetQualifiedDevices",
			Handler:       _RewardsService_GetQualifiedDevices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/rewards/rewards.proto",
}
