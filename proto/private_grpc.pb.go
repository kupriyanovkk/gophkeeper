// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: proto/private.proto

package proto

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
	Private_CreatePrivateData_FullMethodName    = "/proto.Private/CreatePrivateData"
	Private_GetPrivateData_FullMethodName       = "/proto.Private/GetPrivateData"
	Private_UpdatePrivateData_FullMethodName    = "/proto.Private/UpdatePrivateData"
	Private_DeletePrivateData_FullMethodName    = "/proto.Private/DeletePrivateData"
	Private_GetPrivateDataByType_FullMethodName = "/proto.Private/GetPrivateDataByType"
)

// PrivateClient is the client API for Private service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivateClient interface {
	CreatePrivateData(ctx context.Context, in *CreatePrivateDataRequest, opts ...grpc.CallOption) (*CreatePrivateDataResponse, error)
	GetPrivateData(ctx context.Context, in *GetPrivateDataRequest, opts ...grpc.CallOption) (*GetPrivateDataResponse, error)
	UpdatePrivateData(ctx context.Context, in *UpdatePrivateDataRequest, opts ...grpc.CallOption) (*UpdatePrivateDataResponse, error)
	DeletePrivateData(ctx context.Context, in *DeletePrivateDataRequest, opts ...grpc.CallOption) (*DeletePrivateDataResponse, error)
	GetPrivateDataByType(ctx context.Context, in *GetPrivateDataByTypeRequest, opts ...grpc.CallOption) (*GetPrivateDataByTypeResponse, error)
}

type privateClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivateClient(cc grpc.ClientConnInterface) PrivateClient {
	return &privateClient{cc}
}

func (c *privateClient) CreatePrivateData(ctx context.Context, in *CreatePrivateDataRequest, opts ...grpc.CallOption) (*CreatePrivateDataResponse, error) {
	out := new(CreatePrivateDataResponse)
	err := c.cc.Invoke(ctx, Private_CreatePrivateData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateClient) GetPrivateData(ctx context.Context, in *GetPrivateDataRequest, opts ...grpc.CallOption) (*GetPrivateDataResponse, error) {
	out := new(GetPrivateDataResponse)
	err := c.cc.Invoke(ctx, Private_GetPrivateData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateClient) UpdatePrivateData(ctx context.Context, in *UpdatePrivateDataRequest, opts ...grpc.CallOption) (*UpdatePrivateDataResponse, error) {
	out := new(UpdatePrivateDataResponse)
	err := c.cc.Invoke(ctx, Private_UpdatePrivateData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateClient) DeletePrivateData(ctx context.Context, in *DeletePrivateDataRequest, opts ...grpc.CallOption) (*DeletePrivateDataResponse, error) {
	out := new(DeletePrivateDataResponse)
	err := c.cc.Invoke(ctx, Private_DeletePrivateData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateClient) GetPrivateDataByType(ctx context.Context, in *GetPrivateDataByTypeRequest, opts ...grpc.CallOption) (*GetPrivateDataByTypeResponse, error) {
	out := new(GetPrivateDataByTypeResponse)
	err := c.cc.Invoke(ctx, Private_GetPrivateDataByType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivateServer is the server API for Private service.
// All implementations must embed UnimplementedPrivateServer
// for forward compatibility
type PrivateServer interface {
	CreatePrivateData(context.Context, *CreatePrivateDataRequest) (*CreatePrivateDataResponse, error)
	GetPrivateData(context.Context, *GetPrivateDataRequest) (*GetPrivateDataResponse, error)
	UpdatePrivateData(context.Context, *UpdatePrivateDataRequest) (*UpdatePrivateDataResponse, error)
	DeletePrivateData(context.Context, *DeletePrivateDataRequest) (*DeletePrivateDataResponse, error)
	GetPrivateDataByType(context.Context, *GetPrivateDataByTypeRequest) (*GetPrivateDataByTypeResponse, error)
	mustEmbedUnimplementedPrivateServer()
}

// UnimplementedPrivateServer must be embedded to have forward compatible implementations.
type UnimplementedPrivateServer struct {
}

func (UnimplementedPrivateServer) CreatePrivateData(context.Context, *CreatePrivateDataRequest) (*CreatePrivateDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrivateData not implemented")
}
func (UnimplementedPrivateServer) GetPrivateData(context.Context, *GetPrivateDataRequest) (*GetPrivateDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateData not implemented")
}
func (UnimplementedPrivateServer) UpdatePrivateData(context.Context, *UpdatePrivateDataRequest) (*UpdatePrivateDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrivateData not implemented")
}
func (UnimplementedPrivateServer) DeletePrivateData(context.Context, *DeletePrivateDataRequest) (*DeletePrivateDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePrivateData not implemented")
}
func (UnimplementedPrivateServer) GetPrivateDataByType(context.Context, *GetPrivateDataByTypeRequest) (*GetPrivateDataByTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateDataByType not implemented")
}
func (UnimplementedPrivateServer) mustEmbedUnimplementedPrivateServer() {}

// UnsafePrivateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivateServer will
// result in compilation errors.
type UnsafePrivateServer interface {
	mustEmbedUnimplementedPrivateServer()
}

func RegisterPrivateServer(s grpc.ServiceRegistrar, srv PrivateServer) {
	s.RegisterService(&Private_ServiceDesc, srv)
}

func _Private_CreatePrivateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrivateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServer).CreatePrivateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Private_CreatePrivateData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServer).CreatePrivateData(ctx, req.(*CreatePrivateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Private_GetPrivateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrivateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServer).GetPrivateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Private_GetPrivateData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServer).GetPrivateData(ctx, req.(*GetPrivateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Private_UpdatePrivateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePrivateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServer).UpdatePrivateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Private_UpdatePrivateData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServer).UpdatePrivateData(ctx, req.(*UpdatePrivateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Private_DeletePrivateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePrivateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServer).DeletePrivateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Private_DeletePrivateData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServer).DeletePrivateData(ctx, req.(*DeletePrivateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Private_GetPrivateDataByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrivateDataByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServer).GetPrivateDataByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Private_GetPrivateDataByType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServer).GetPrivateDataByType(ctx, req.(*GetPrivateDataByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Private_ServiceDesc is the grpc.ServiceDesc for Private service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Private_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Private",
	HandlerType: (*PrivateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePrivateData",
			Handler:    _Private_CreatePrivateData_Handler,
		},
		{
			MethodName: "GetPrivateData",
			Handler:    _Private_GetPrivateData_Handler,
		},
		{
			MethodName: "UpdatePrivateData",
			Handler:    _Private_UpdatePrivateData_Handler,
		},
		{
			MethodName: "DeletePrivateData",
			Handler:    _Private_DeletePrivateData_Handler,
		},
		{
			MethodName: "GetPrivateDataByType",
			Handler:    _Private_GetPrivateDataByType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/private.proto",
}
