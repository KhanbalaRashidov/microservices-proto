// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: shipping.proto

package shipping

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
	Shipping_Create_FullMethodName = "/Shipping/Create"
)

// ShippingClient is the client API for Shipping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShippingClient interface {
	Create(ctx context.Context, in *CreateShippingRequest, opts ...grpc.CallOption) (*CreateShippingResponse, error)
}

type shippingClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingClient(cc grpc.ClientConnInterface) ShippingClient {
	return &shippingClient{cc}
}

func (c *shippingClient) Create(ctx context.Context, in *CreateShippingRequest, opts ...grpc.CallOption) (*CreateShippingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateShippingResponse)
	err := c.cc.Invoke(ctx, Shipping_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServer is the server API for Shipping service.
// All implementations must embed UnimplementedShippingServer
// for forward compatibility.
type ShippingServer interface {
	Create(context.Context, *CreateShippingRequest) (*CreateShippingResponse, error)
	mustEmbedUnimplementedShippingServer()
}

// UnimplementedShippingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShippingServer struct{}

func (UnimplementedShippingServer) Create(context.Context, *CreateShippingRequest) (*CreateShippingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedShippingServer) mustEmbedUnimplementedShippingServer() {}
func (UnimplementedShippingServer) testEmbeddedByValue()                  {}

// UnsafeShippingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShippingServer will
// result in compilation errors.
type UnsafeShippingServer interface {
	mustEmbedUnimplementedShippingServer()
}

func RegisterShippingServer(s grpc.ServiceRegistrar, srv ShippingServer) {
	// If the following call pancis, it indicates UnimplementedShippingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Shipping_ServiceDesc, srv)
}

func _Shipping_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShippingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shipping_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServer).Create(ctx, req.(*CreateShippingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shipping_ServiceDesc is the grpc.ServiceDesc for Shipping service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shipping_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Shipping",
	HandlerType: (*ShippingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Shipping_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shipping.proto",
}
