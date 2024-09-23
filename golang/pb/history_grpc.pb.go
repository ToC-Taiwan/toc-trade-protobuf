// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: forwarder/history.proto

package pb

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
	HistoryDataInterface_GetStockHistoryTick_FullMethodName  = "/forwarder.HistoryDataInterface/GetStockHistoryTick"
	HistoryDataInterface_GetStockHistoryKbar_FullMethodName  = "/forwarder.HistoryDataInterface/GetStockHistoryKbar"
	HistoryDataInterface_GetStockHistoryClose_FullMethodName = "/forwarder.HistoryDataInterface/GetStockHistoryClose"
	HistoryDataInterface_GetFutureHistoryKbar_FullMethodName = "/forwarder.HistoryDataInterface/GetFutureHistoryKbar"
)

// HistoryDataInterfaceClient is the client API for HistoryDataInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HistoryDataInterfaceClient interface {
	GetStockHistoryTick(ctx context.Context, in *StockNumArrWithDate, opts ...grpc.CallOption) (*HistoryTickResponse, error)
	GetStockHistoryKbar(ctx context.Context, in *StockNumArrWithDate, opts ...grpc.CallOption) (*HistoryKbarResponse, error)
	GetStockHistoryClose(ctx context.Context, in *StockNumArrWithDate, opts ...grpc.CallOption) (*HistoryCloseResponse, error)
	GetFutureHistoryKbar(ctx context.Context, in *FutureCodeArrWithDate, opts ...grpc.CallOption) (*HistoryKbarResponse, error)
}

type historyDataInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewHistoryDataInterfaceClient(cc grpc.ClientConnInterface) HistoryDataInterfaceClient {
	return &historyDataInterfaceClient{cc}
}

func (c *historyDataInterfaceClient) GetStockHistoryTick(ctx context.Context, in *StockNumArrWithDate, opts ...grpc.CallOption) (*HistoryTickResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HistoryTickResponse)
	err := c.cc.Invoke(ctx, HistoryDataInterface_GetStockHistoryTick_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyDataInterfaceClient) GetStockHistoryKbar(ctx context.Context, in *StockNumArrWithDate, opts ...grpc.CallOption) (*HistoryKbarResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HistoryKbarResponse)
	err := c.cc.Invoke(ctx, HistoryDataInterface_GetStockHistoryKbar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyDataInterfaceClient) GetStockHistoryClose(ctx context.Context, in *StockNumArrWithDate, opts ...grpc.CallOption) (*HistoryCloseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HistoryCloseResponse)
	err := c.cc.Invoke(ctx, HistoryDataInterface_GetStockHistoryClose_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyDataInterfaceClient) GetFutureHistoryKbar(ctx context.Context, in *FutureCodeArrWithDate, opts ...grpc.CallOption) (*HistoryKbarResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HistoryKbarResponse)
	err := c.cc.Invoke(ctx, HistoryDataInterface_GetFutureHistoryKbar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HistoryDataInterfaceServer is the server API for HistoryDataInterface service.
// All implementations must embed UnimplementedHistoryDataInterfaceServer
// for forward compatibility.
type HistoryDataInterfaceServer interface {
	GetStockHistoryTick(context.Context, *StockNumArrWithDate) (*HistoryTickResponse, error)
	GetStockHistoryKbar(context.Context, *StockNumArrWithDate) (*HistoryKbarResponse, error)
	GetStockHistoryClose(context.Context, *StockNumArrWithDate) (*HistoryCloseResponse, error)
	GetFutureHistoryKbar(context.Context, *FutureCodeArrWithDate) (*HistoryKbarResponse, error)
	mustEmbedUnimplementedHistoryDataInterfaceServer()
}

// UnimplementedHistoryDataInterfaceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHistoryDataInterfaceServer struct{}

func (UnimplementedHistoryDataInterfaceServer) GetStockHistoryTick(context.Context, *StockNumArrWithDate) (*HistoryTickResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockHistoryTick not implemented")
}
func (UnimplementedHistoryDataInterfaceServer) GetStockHistoryKbar(context.Context, *StockNumArrWithDate) (*HistoryKbarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockHistoryKbar not implemented")
}
func (UnimplementedHistoryDataInterfaceServer) GetStockHistoryClose(context.Context, *StockNumArrWithDate) (*HistoryCloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockHistoryClose not implemented")
}
func (UnimplementedHistoryDataInterfaceServer) GetFutureHistoryKbar(context.Context, *FutureCodeArrWithDate) (*HistoryKbarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFutureHistoryKbar not implemented")
}
func (UnimplementedHistoryDataInterfaceServer) mustEmbedUnimplementedHistoryDataInterfaceServer() {}
func (UnimplementedHistoryDataInterfaceServer) testEmbeddedByValue()                              {}

// UnsafeHistoryDataInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HistoryDataInterfaceServer will
// result in compilation errors.
type UnsafeHistoryDataInterfaceServer interface {
	mustEmbedUnimplementedHistoryDataInterfaceServer()
}

func RegisterHistoryDataInterfaceServer(s grpc.ServiceRegistrar, srv HistoryDataInterfaceServer) {
	// If the following call pancis, it indicates UnimplementedHistoryDataInterfaceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HistoryDataInterface_ServiceDesc, srv)
}

func _HistoryDataInterface_GetStockHistoryTick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockNumArrWithDate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryDataInterfaceServer).GetStockHistoryTick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HistoryDataInterface_GetStockHistoryTick_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryDataInterfaceServer).GetStockHistoryTick(ctx, req.(*StockNumArrWithDate))
	}
	return interceptor(ctx, in, info, handler)
}

func _HistoryDataInterface_GetStockHistoryKbar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockNumArrWithDate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryDataInterfaceServer).GetStockHistoryKbar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HistoryDataInterface_GetStockHistoryKbar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryDataInterfaceServer).GetStockHistoryKbar(ctx, req.(*StockNumArrWithDate))
	}
	return interceptor(ctx, in, info, handler)
}

func _HistoryDataInterface_GetStockHistoryClose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockNumArrWithDate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryDataInterfaceServer).GetStockHistoryClose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HistoryDataInterface_GetStockHistoryClose_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryDataInterfaceServer).GetStockHistoryClose(ctx, req.(*StockNumArrWithDate))
	}
	return interceptor(ctx, in, info, handler)
}

func _HistoryDataInterface_GetFutureHistoryKbar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FutureCodeArrWithDate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryDataInterfaceServer).GetFutureHistoryKbar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HistoryDataInterface_GetFutureHistoryKbar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryDataInterfaceServer).GetFutureHistoryKbar(ctx, req.(*FutureCodeArrWithDate))
	}
	return interceptor(ctx, in, info, handler)
}

// HistoryDataInterface_ServiceDesc is the grpc.ServiceDesc for HistoryDataInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HistoryDataInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "forwarder.HistoryDataInterface",
	HandlerType: (*HistoryDataInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStockHistoryTick",
			Handler:    _HistoryDataInterface_GetStockHistoryTick_Handler,
		},
		{
			MethodName: "GetStockHistoryKbar",
			Handler:    _HistoryDataInterface_GetStockHistoryKbar_Handler,
		},
		{
			MethodName: "GetStockHistoryClose",
			Handler:    _HistoryDataInterface_GetStockHistoryClose_Handler,
		},
		{
			MethodName: "GetFutureHistoryKbar",
			Handler:    _HistoryDataInterface_GetFutureHistoryKbar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "forwarder/history.proto",
}