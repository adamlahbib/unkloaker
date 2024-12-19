// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.6
// source: api/proto/api.proto

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
	Unkloaker_Deploy_FullMethodName   = "/proto.Unkloaker/Deploy"
	Unkloaker_Undeploy_FullMethodName = "/proto.Unkloaker/Undeploy"
)

// UnkloakerClient is the client API for Unkloaker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UnkloakerClient interface {
	Deploy(ctx context.Context, in *UnkloakerDeployRequest, opts ...grpc.CallOption) (*UnkloakerDeployResponse, error)
	Undeploy(ctx context.Context, in *UnkloakerUndeployRequest, opts ...grpc.CallOption) (*UnkloakerUndeployResponse, error)
}

type unkloakerClient struct {
	cc grpc.ClientConnInterface
}

func NewUnkloakerClient(cc grpc.ClientConnInterface) UnkloakerClient {
	return &unkloakerClient{cc}
}

func (c *unkloakerClient) Deploy(ctx context.Context, in *UnkloakerDeployRequest, opts ...grpc.CallOption) (*UnkloakerDeployResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnkloakerDeployResponse)
	err := c.cc.Invoke(ctx, Unkloaker_Deploy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unkloakerClient) Undeploy(ctx context.Context, in *UnkloakerUndeployRequest, opts ...grpc.CallOption) (*UnkloakerUndeployResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnkloakerUndeployResponse)
	err := c.cc.Invoke(ctx, Unkloaker_Undeploy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnkloakerServer is the server API for Unkloaker service.
// All implementations must embed UnimplementedUnkloakerServer
// for forward compatibility.
type UnkloakerServer interface {
	Deploy(context.Context, *UnkloakerDeployRequest) (*UnkloakerDeployResponse, error)
	Undeploy(context.Context, *UnkloakerUndeployRequest) (*UnkloakerUndeployResponse, error)
	mustEmbedUnimplementedUnkloakerServer()
}

// UnimplementedUnkloakerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUnkloakerServer struct{}

func (UnimplementedUnkloakerServer) Deploy(context.Context, *UnkloakerDeployRequest) (*UnkloakerDeployResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deploy not implemented")
}
func (UnimplementedUnkloakerServer) Undeploy(context.Context, *UnkloakerUndeployRequest) (*UnkloakerUndeployResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Undeploy not implemented")
}
func (UnimplementedUnkloakerServer) mustEmbedUnimplementedUnkloakerServer() {}
func (UnimplementedUnkloakerServer) testEmbeddedByValue()                   {}

// UnsafeUnkloakerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UnkloakerServer will
// result in compilation errors.
type UnsafeUnkloakerServer interface {
	mustEmbedUnimplementedUnkloakerServer()
}

func RegisterUnkloakerServer(s grpc.ServiceRegistrar, srv UnkloakerServer) {
	// If the following call pancis, it indicates UnimplementedUnkloakerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Unkloaker_ServiceDesc, srv)
}

func _Unkloaker_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnkloakerDeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnkloakerServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Unkloaker_Deploy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnkloakerServer).Deploy(ctx, req.(*UnkloakerDeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Unkloaker_Undeploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnkloakerUndeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnkloakerServer).Undeploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Unkloaker_Undeploy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnkloakerServer).Undeploy(ctx, req.(*UnkloakerUndeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Unkloaker_ServiceDesc is the grpc.ServiceDesc for Unkloaker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Unkloaker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Unkloaker",
	HandlerType: (*UnkloakerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deploy",
			Handler:    _Unkloaker_Deploy_Handler,
		},
		{
			MethodName: "Undeploy",
			Handler:    _Unkloaker_Undeploy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/api.proto",
}

const (
	Kloner_Clone_FullMethodName   = "/proto.Kloner/Clone"
	Kloner_Cleanup_FullMethodName = "/proto.Kloner/Cleanup"
)

// KlonerClient is the client API for Kloner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KlonerClient interface {
	Clone(ctx context.Context, in *KlonerCloneRequest, opts ...grpc.CallOption) (*KlonerCloneResponse, error)
	Cleanup(ctx context.Context, in *KlonerCleanupRequest, opts ...grpc.CallOption) (*KlonerCleanupResponse, error)
}

type klonerClient struct {
	cc grpc.ClientConnInterface
}

func NewKlonerClient(cc grpc.ClientConnInterface) KlonerClient {
	return &klonerClient{cc}
}

func (c *klonerClient) Clone(ctx context.Context, in *KlonerCloneRequest, opts ...grpc.CallOption) (*KlonerCloneResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KlonerCloneResponse)
	err := c.cc.Invoke(ctx, Kloner_Clone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klonerClient) Cleanup(ctx context.Context, in *KlonerCleanupRequest, opts ...grpc.CallOption) (*KlonerCleanupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KlonerCleanupResponse)
	err := c.cc.Invoke(ctx, Kloner_Cleanup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KlonerServer is the server API for Kloner service.
// All implementations must embed UnimplementedKlonerServer
// for forward compatibility.
type KlonerServer interface {
	Clone(context.Context, *KlonerCloneRequest) (*KlonerCloneResponse, error)
	Cleanup(context.Context, *KlonerCleanupRequest) (*KlonerCleanupResponse, error)
	mustEmbedUnimplementedKlonerServer()
}

// UnimplementedKlonerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKlonerServer struct{}

func (UnimplementedKlonerServer) Clone(context.Context, *KlonerCloneRequest) (*KlonerCloneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clone not implemented")
}
func (UnimplementedKlonerServer) Cleanup(context.Context, *KlonerCleanupRequest) (*KlonerCleanupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cleanup not implemented")
}
func (UnimplementedKlonerServer) mustEmbedUnimplementedKlonerServer() {}
func (UnimplementedKlonerServer) testEmbeddedByValue()                {}

// UnsafeKlonerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KlonerServer will
// result in compilation errors.
type UnsafeKlonerServer interface {
	mustEmbedUnimplementedKlonerServer()
}

func RegisterKlonerServer(s grpc.ServiceRegistrar, srv KlonerServer) {
	// If the following call pancis, it indicates UnimplementedKlonerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Kloner_ServiceDesc, srv)
}

func _Kloner_Clone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KlonerCloneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlonerServer).Clone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kloner_Clone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlonerServer).Clone(ctx, req.(*KlonerCloneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kloner_Cleanup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KlonerCleanupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlonerServer).Cleanup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kloner_Cleanup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlonerServer).Cleanup(ctx, req.(*KlonerCleanupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Kloner_ServiceDesc is the grpc.ServiceDesc for Kloner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Kloner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Kloner",
	HandlerType: (*KlonerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Clone",
			Handler:    _Kloner_Clone_Handler,
		},
		{
			MethodName: "Cleanup",
			Handler:    _Kloner_Cleanup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/api.proto",
}

const (
	Konstruktor_Build_FullMethodName = "/proto.Konstruktor/Build"
)

// KonstruktorClient is the client API for Konstruktor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KonstruktorClient interface {
	Build(ctx context.Context, in *KonstruktorBuildRequest, opts ...grpc.CallOption) (*KonstruktorBuildResponse, error)
}

type konstruktorClient struct {
	cc grpc.ClientConnInterface
}

func NewKonstruktorClient(cc grpc.ClientConnInterface) KonstruktorClient {
	return &konstruktorClient{cc}
}

func (c *konstruktorClient) Build(ctx context.Context, in *KonstruktorBuildRequest, opts ...grpc.CallOption) (*KonstruktorBuildResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KonstruktorBuildResponse)
	err := c.cc.Invoke(ctx, Konstruktor_Build_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KonstruktorServer is the server API for Konstruktor service.
// All implementations must embed UnimplementedKonstruktorServer
// for forward compatibility.
type KonstruktorServer interface {
	Build(context.Context, *KonstruktorBuildRequest) (*KonstruktorBuildResponse, error)
	mustEmbedUnimplementedKonstruktorServer()
}

// UnimplementedKonstruktorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKonstruktorServer struct{}

func (UnimplementedKonstruktorServer) Build(context.Context, *KonstruktorBuildRequest) (*KonstruktorBuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Build not implemented")
}
func (UnimplementedKonstruktorServer) mustEmbedUnimplementedKonstruktorServer() {}
func (UnimplementedKonstruktorServer) testEmbeddedByValue()                     {}

// UnsafeKonstruktorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KonstruktorServer will
// result in compilation errors.
type UnsafeKonstruktorServer interface {
	mustEmbedUnimplementedKonstruktorServer()
}

func RegisterKonstruktorServer(s grpc.ServiceRegistrar, srv KonstruktorServer) {
	// If the following call pancis, it indicates UnimplementedKonstruktorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Konstruktor_ServiceDesc, srv)
}

func _Konstruktor_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KonstruktorBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KonstruktorServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Konstruktor_Build_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KonstruktorServer).Build(ctx, req.(*KonstruktorBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Konstruktor_ServiceDesc is the grpc.ServiceDesc for Konstruktor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Konstruktor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Konstruktor",
	HandlerType: (*KonstruktorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Build",
			Handler:    _Konstruktor_Build_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/api.proto",
}
