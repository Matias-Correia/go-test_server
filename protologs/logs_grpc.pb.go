// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protologs

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

// LogTestDataClient is the client API for LogTestData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogTestDataClient interface {
	//Sends
	SendLogs(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Empty, error)
}

type logTestDataClient struct {
	cc grpc.ClientConnInterface
}

func NewLogTestDataClient(cc grpc.ClientConnInterface) LogTestDataClient {
	return &logTestDataClient{cc}
}

func (c *logTestDataClient) SendLogs(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protologs.LogTestData/SendLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogTestDataServer is the server API for LogTestData service.
// All implementations must embed UnimplementedLogTestDataServer
// for forward compatibility
type LogTestDataServer interface {
	//Sends
	SendLogs(context.Context, *Log) (*Empty, error)
	mustEmbedUnimplementedLogTestDataServer()
}

// UnimplementedLogTestDataServer must be embedded to have forward compatible implementations.
type UnimplementedLogTestDataServer struct {
}

func (UnimplementedLogTestDataServer) SendLogs(context.Context, *Log) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLogs not implemented")
}
func (UnimplementedLogTestDataServer) mustEmbedUnimplementedLogTestDataServer() {}

// UnsafeLogTestDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogTestDataServer will
// result in compilation errors.
type UnsafeLogTestDataServer interface {
	mustEmbedUnimplementedLogTestDataServer()
}

func RegisterLogTestDataServer(s grpc.ServiceRegistrar, srv LogTestDataServer) {
	s.RegisterService(&LogTestData_ServiceDesc, srv)
}

func _LogTestData_SendLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogTestDataServer).SendLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protologs.LogTestData/SendLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogTestDataServer).SendLogs(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

// LogTestData_ServiceDesc is the grpc.ServiceDesc for LogTestData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogTestData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protologs.LogTestData",
	HandlerType: (*LogTestDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendLogs",
			Handler:    _LogTestData_SendLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protologs/logs.proto",
}
