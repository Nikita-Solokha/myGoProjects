package driver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	DriverService_RegisterDriver_FullMethodName   = "/driver.DriverService/RegisterDriver"
	DriverService_UnregisterDriver_FullMethodName = "/driver.DriverService/UnregisterDriver"
)

type DriverServiceClient interface {
	RegisterDriver(ctx context.Context, in *RegisterDriverRequest, opts ...grpc.CallOption) (*RegisterDriverResponse, error)
	UnregisterDriver(ctx context.Context, in *RegisterDriverRequest, opts ...grpc.CallOption) (*RegisterDriverResponse, error)
}

type driverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverServiceClient(cc grpc.ClientConnInterface) DriverServiceClient {
	return &driverServiceClient{cc}
}

func (c *driverServiceClient) RegisterDriver(ctx context.Context, in *RegisterDriverRequest, opts ...grpc.CallOption) (*RegisterDriverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterDriverResponse)
	err := c.cc.Invoke(ctx, DriverService_RegisterDriver_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverServiceClient) UnregisterDriver(ctx context.Context, in *RegisterDriverRequest, opts ...grpc.CallOption) (*RegisterDriverResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterDriverResponse)
	err := c.cc.Invoke(ctx, DriverService_UnregisterDriver_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DriverServiceServer interface {
	RegisterDriver(context.Context, *RegisterDriverRequest) (*RegisterDriverResponse, error)
	UnregisterDriver(context.Context, *RegisterDriverRequest) (*RegisterDriverResponse, error)
	mustEmbedUnimplementedDriverServiceServer()
}

type UnimplementedDriverServiceServer struct{}

func (UnimplementedDriverServiceServer) RegisterDriver(context.Context, *RegisterDriverRequest) (*RegisterDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDriver not implemented")
}
func (UnimplementedDriverServiceServer) UnregisterDriver(context.Context, *RegisterDriverRequest) (*RegisterDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterDriver not implemented")
}
func (UnimplementedDriverServiceServer) mustEmbedUnimplementedDriverServiceServer() {}
func (UnimplementedDriverServiceServer) testEmbeddedByValue()                       {}

type UnsafeDriverServiceServer interface {
	mustEmbedUnimplementedDriverServiceServer()
}

func RegisterDriverServiceServer(s grpc.ServiceRegistrar, srv DriverServiceServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DriverService_ServiceDesc, srv)
}

func _DriverService_RegisterDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).RegisterDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverService_RegisterDriver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).RegisterDriver(ctx, req.(*RegisterDriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverService_UnregisterDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).UnregisterDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverService_UnregisterDriver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).UnregisterDriver(ctx, req.(*RegisterDriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var DriverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "driver.DriverService",
	HandlerType: (*DriverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDriver",
			Handler:    _DriverService_RegisterDriver_Handler,
		},
		{
			MethodName: "UnregisterDriver",
			Handler:    _DriverService_UnregisterDriver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driver.proto",
}
