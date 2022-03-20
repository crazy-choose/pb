// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stream

import (
	model "/out/model"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamServiceClient interface {
	Request(ctx context.Context, opts ...grpc.CallOption) (StreamService_RequestClient, error)
	Notify(ctx context.Context, opts ...grpc.CallOption) (StreamService_NotifyClient, error)
	SendTrap(ctx context.Context, opts ...grpc.CallOption) (StreamService_SendTrapClient, error)
}

type streamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServiceClient(cc grpc.ClientConnInterface) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) Request(ctx context.Context, opts ...grpc.CallOption) (StreamService_RequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[0], "/StreamService/Request", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceRequestClient{stream}
	return x, nil
}

type StreamService_RequestClient interface {
	Send(*model.ReqMsg) error
	CloseAndRecv() (*model.RspMsg, error)
	grpc.ClientStream
}

type streamServiceRequestClient struct {
	grpc.ClientStream
}

func (x *streamServiceRequestClient) Send(m *model.ReqMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceRequestClient) CloseAndRecv() (*model.RspMsg, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(model.RspMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) Notify(ctx context.Context, opts ...grpc.CallOption) (StreamService_NotifyClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[1], "/StreamService/Notify", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceNotifyClient{stream}
	return x, nil
}

type StreamService_NotifyClient interface {
	Send(*model.NotifyMsg) error
	Recv() (*model.NetResult, error)
	grpc.ClientStream
}

type streamServiceNotifyClient struct {
	grpc.ClientStream
}

func (x *streamServiceNotifyClient) Send(m *model.NotifyMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceNotifyClient) Recv() (*model.NetResult, error) {
	m := new(model.NetResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) SendTrap(ctx context.Context, opts ...grpc.CallOption) (StreamService_SendTrapClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[2], "/StreamService/SendTrap", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceSendTrapClient{stream}
	return x, nil
}

type StreamService_SendTrapClient interface {
	Send(*model.TrapMsg) error
	Recv() (*model.NetResult, error)
	grpc.ClientStream
}

type streamServiceSendTrapClient struct {
	grpc.ClientStream
}

func (x *streamServiceSendTrapClient) Send(m *model.TrapMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceSendTrapClient) Recv() (*model.NetResult, error) {
	m := new(model.NetResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
// All implementations must embed UnimplementedStreamServiceServer
// for forward compatibility
type StreamServiceServer interface {
	Request(StreamService_RequestServer) error
	Notify(StreamService_NotifyServer) error
	SendTrap(StreamService_SendTrapServer) error
	mustEmbedUnimplementedStreamServiceServer()
}

// UnimplementedStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (UnimplementedStreamServiceServer) Request(StreamService_RequestServer) error {
	return status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (UnimplementedStreamServiceServer) Notify(StreamService_NotifyServer) error {
	return status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedStreamServiceServer) SendTrap(StreamService_SendTrapServer) error {
	return status.Errorf(codes.Unimplemented, "method SendTrap not implemented")
}
func (UnimplementedStreamServiceServer) mustEmbedUnimplementedStreamServiceServer() {}

// UnsafeStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServiceServer will
// result in compilation errors.
type UnsafeStreamServiceServer interface {
	mustEmbedUnimplementedStreamServiceServer()
}

func RegisterStreamServiceServer(s grpc.ServiceRegistrar, srv StreamServiceServer) {
	s.RegisterService(&StreamService_ServiceDesc, srv)
}

func _StreamService_Request_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Request(&streamServiceRequestServer{stream})
}

type StreamService_RequestServer interface {
	SendAndClose(*model.RspMsg) error
	Recv() (*model.ReqMsg, error)
	grpc.ServerStream
}

type streamServiceRequestServer struct {
	grpc.ServerStream
}

func (x *streamServiceRequestServer) SendAndClose(m *model.RspMsg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRequestServer) Recv() (*model.ReqMsg, error) {
	m := new(model.ReqMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_Notify_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Notify(&streamServiceNotifyServer{stream})
}

type StreamService_NotifyServer interface {
	Send(*model.NetResult) error
	Recv() (*model.NotifyMsg, error)
	grpc.ServerStream
}

type streamServiceNotifyServer struct {
	grpc.ServerStream
}

func (x *streamServiceNotifyServer) Send(m *model.NetResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceNotifyServer) Recv() (*model.NotifyMsg, error) {
	m := new(model.NotifyMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_SendTrap_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).SendTrap(&streamServiceSendTrapServer{stream})
}

type StreamService_SendTrapServer interface {
	Send(*model.NetResult) error
	Recv() (*model.TrapMsg, error)
	grpc.ServerStream
}

type streamServiceSendTrapServer struct {
	grpc.ServerStream
}

func (x *streamServiceSendTrapServer) Send(m *model.NetResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceSendTrapServer) Recv() (*model.TrapMsg, error) {
	m := new(model.TrapMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamService_ServiceDesc is the grpc.ServiceDesc for StreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Request",
			Handler:       _StreamService_Request_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Notify",
			Handler:       _StreamService_Notify_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SendTrap",
			Handler:       _StreamService_SendTrap_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}