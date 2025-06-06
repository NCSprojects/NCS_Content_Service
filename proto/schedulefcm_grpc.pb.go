// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: schedulefcm.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ScheduleService_GetScheduleIdsByStartTime_FullMethodName = "/schedulefcm.ScheduleService/GetScheduleIdsByStartTime"
	ScheduleService_GetStartTimeByScheduleId_FullMethodName  = "/schedulefcm.ScheduleService/GetStartTimeByScheduleId"
)

// ScheduleServiceClient is the client API for ScheduleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScheduleServiceClient interface {
	// 특정 시작 시간에 해당하는 스케줄 ID 목록 조회
	GetScheduleIdsByStartTime(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
	// 스케줄 ID로 스케줄 시작 시간 조회
	GetStartTimeByScheduleId(ctx context.Context, in *ScheduleIdRequest, opts ...grpc.CallOption) (*StartTimeResponse, error)
}

type scheduleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScheduleServiceClient(cc grpc.ClientConnInterface) ScheduleServiceClient {
	return &scheduleServiceClient{cc}
}

func (c *scheduleServiceClient) GetScheduleIdsByStartTime(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScheduleResponse)
	err := c.cc.Invoke(ctx, ScheduleService_GetScheduleIdsByStartTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) GetStartTimeByScheduleId(ctx context.Context, in *ScheduleIdRequest, opts ...grpc.CallOption) (*StartTimeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartTimeResponse)
	err := c.cc.Invoke(ctx, ScheduleService_GetStartTimeByScheduleId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduleServiceServer is the server API for ScheduleService service.
// All implementations must embed UnimplementedScheduleServiceServer
// for forward compatibility.
type ScheduleServiceServer interface {
	// 특정 시작 시간에 해당하는 스케줄 ID 목록 조회
	GetScheduleIdsByStartTime(context.Context, *ScheduleRequest) (*ScheduleResponse, error)
	// 스케줄 ID로 스케줄 시작 시간 조회
	GetStartTimeByScheduleId(context.Context, *ScheduleIdRequest) (*StartTimeResponse, error)
	mustEmbedUnimplementedScheduleServiceServer()
}

// UnimplementedScheduleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedScheduleServiceServer struct{}

func (UnimplementedScheduleServiceServer) GetScheduleIdsByStartTime(context.Context, *ScheduleRequest) (*ScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduleIdsByStartTime not implemented")
}
func (UnimplementedScheduleServiceServer) GetStartTimeByScheduleId(context.Context, *ScheduleIdRequest) (*StartTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStartTimeByScheduleId not implemented")
}
func (UnimplementedScheduleServiceServer) mustEmbedUnimplementedScheduleServiceServer() {}
func (UnimplementedScheduleServiceServer) testEmbeddedByValue()                         {}

// UnsafeScheduleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScheduleServiceServer will
// result in compilation errors.
type UnsafeScheduleServiceServer interface {
	mustEmbedUnimplementedScheduleServiceServer()
}

func RegisterScheduleServiceServer(s grpc.ServiceRegistrar, srv ScheduleServiceServer) {
	// If the following call pancis, it indicates UnimplementedScheduleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ScheduleService_ServiceDesc, srv)
}

func _ScheduleService_GetScheduleIdsByStartTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetScheduleIdsByStartTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_GetScheduleIdsByStartTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetScheduleIdsByStartTime(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_GetStartTimeByScheduleId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetStartTimeByScheduleId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScheduleService_GetStartTimeByScheduleId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetStartTimeByScheduleId(ctx, req.(*ScheduleIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScheduleService_ServiceDesc is the grpc.ServiceDesc for ScheduleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScheduleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schedulefcm.ScheduleService",
	HandlerType: (*ScheduleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScheduleIdsByStartTime",
			Handler:    _ScheduleService_GetScheduleIdsByStartTime_Handler,
		},
		{
			MethodName: "GetStartTimeByScheduleId",
			Handler:    _ScheduleService_GetStartTimeByScheduleId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schedulefcm.proto",
}
