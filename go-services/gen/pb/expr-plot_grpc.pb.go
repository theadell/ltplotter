// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: expr-plot.proto

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

// ExpressionPlotServiceClient is the client API for ExpressionPlotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExpressionPlotServiceClient interface {
	GeneratePlot(ctx context.Context, in *ExprPlotRequest, opts ...grpc.CallOption) (*ExprPlotResponse, error)
}

type expressionPlotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExpressionPlotServiceClient(cc grpc.ClientConnInterface) ExpressionPlotServiceClient {
	return &expressionPlotServiceClient{cc}
}

func (c *expressionPlotServiceClient) GeneratePlot(ctx context.Context, in *ExprPlotRequest, opts ...grpc.CallOption) (*ExprPlotResponse, error) {
	out := new(ExprPlotResponse)
	err := c.cc.Invoke(ctx, "/plot.ExpressionPlotService/GeneratePlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpressionPlotServiceServer is the server API for ExpressionPlotService service.
// All implementations must embed UnimplementedExpressionPlotServiceServer
// for forward compatibility
type ExpressionPlotServiceServer interface {
	GeneratePlot(context.Context, *ExprPlotRequest) (*ExprPlotResponse, error)
	mustEmbedUnimplementedExpressionPlotServiceServer()
}

// UnimplementedExpressionPlotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExpressionPlotServiceServer struct {
}

func (UnimplementedExpressionPlotServiceServer) GeneratePlot(context.Context, *ExprPlotRequest) (*ExprPlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePlot not implemented")
}
func (UnimplementedExpressionPlotServiceServer) mustEmbedUnimplementedExpressionPlotServiceServer() {}

// UnsafeExpressionPlotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExpressionPlotServiceServer will
// result in compilation errors.
type UnsafeExpressionPlotServiceServer interface {
	mustEmbedUnimplementedExpressionPlotServiceServer()
}

func RegisterExpressionPlotServiceServer(s grpc.ServiceRegistrar, srv ExpressionPlotServiceServer) {
	s.RegisterService(&ExpressionPlotService_ServiceDesc, srv)
}

func _ExpressionPlotService_GeneratePlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExprPlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressionPlotServiceServer).GeneratePlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plot.ExpressionPlotService/GeneratePlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressionPlotServiceServer).GeneratePlot(ctx, req.(*ExprPlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExpressionPlotService_ServiceDesc is the grpc.ServiceDesc for ExpressionPlotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExpressionPlotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plot.ExpressionPlotService",
	HandlerType: (*ExpressionPlotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeneratePlot",
			Handler:    _ExpressionPlotService_GeneratePlot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "expr-plot.proto",
}