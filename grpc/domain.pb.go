// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domain.proto

package grpcdomain

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("domain.proto", fileDescriptor_73e6234e76dbdb84) }

var fileDescriptor_73e6234e76dbdb84 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4a, 0xc5, 0x30,
	0x10, 0x86, 0xc5, 0xc5, 0x5b, 0xc4, 0xba, 0x30, 0xba, 0x31, 0x68, 0x15, 0x0f, 0xd0, 0x85, 0x1e,
	0x41, 0x41, 0x04, 0xc5, 0x62, 0x71, 0x1d, 0x6a, 0x3b, 0xd4, 0x81, 0x9a, 0xc4, 0xcc, 0x28, 0xf4,
	0xa0, 0xde, 0x47, 0xda, 0xa4, 0x9a, 0x8a, 0x7d, 0xdb, 0x7c, 0xff, 0x7c, 0xf3, 0x67, 0x44, 0xd6,
	0xda, 0xb7, 0x1a, 0x4d, 0xe1, 0xbc, 0x65, 0x2b, 0x45, 0xe7, 0x5d, 0x13, 0x5e, 0x94, 0xec, 0x80,
	0x35, 0x81, 0xff, 0xc4, 0x06, 0x28, 0x70, 0x75, 0x4c, 0xaf, 0xb5, 0x07, 0x8d, 0x2d, 0x18, 0x46,
	0x1e, 0x74, 0x8f, 0xc4, 0x11, 0x1d, 0x5a, 0x07, 0x46, 0x3b, 0x4b, 0xc8, 0x68, 0xa3, 0x4f, 0x1d,
	0x35, 0xbd, 0x25, 0xf8, 0xf3, 0x7a, 0xf9, 0xb5, 0x2b, 0x36, 0x37, 0xd3, 0x12, 0xf9, 0x20, 0xf6,
	0x6e, 0x81, 0xab, 0xb8, 0x45, 0xe6, 0xc5, 0x6f, 0x81, 0x22, 0x01, 0x4f, 0xf0, 0xfe, 0x01, 0xc4,
	0xea, 0x64, 0x95, 0xbb, 0x7e, 0xb8, 0xd8, 0x91, 0xcf, 0xe2, 0xa0, 0x1a, 0x1b, 0xde, 0xc5, 0x82,
	0xf7, 0x48, 0x2c, 0xcf, 0xd2, 0xa1, 0x94, 0xcc, 0xd6, 0xd3, 0xf5, 0x40, 0xd0, 0x96, 0x22, 0x7b,
	0x74, 0x60, 0xca, 0xf8, 0x8d, 0xa5, 0x31, 0x25, 0xff, 0x1a, 0x97, 0x81, 0x60, 0xac, 0xc4, 0xfe,
	0xf5, 0x78, 0x9a, 0x1f, 0xe5, 0x79, 0x3a, 0xb1, 0x40, 0xb3, 0x33, 0xdf, 0x92, 0x98, 0xa4, 0x2f,
	0x9b, 0xe9, 0xbc, 0x57, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x47, 0x85, 0x81, 0x1e, 0xd4, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DomainClient is the client API for Domain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DomainClient interface {
	// GetServices returns the availible services and their locations
	GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesReply, error)
	// ShareIdentityList Requests the IdentityList which the domain is aware of
	ShareIdentityList(ctx context.Context, in *IdentityListRequest, opts ...grpc.CallOption) (*IdentityListReply, error)
	// OpenPosition declares the new service which is needed and requests Appointments
	OpenPosition(ctx context.Context, in *OpenPositionRequest, opts ...grpc.CallOption) (*OpenPositionReply, error)
	// ClosePosition ends an election and informs if the position was awarded
	ClosePosition(ctx context.Context, in *ClosePositionRequest, opts ...grpc.CallOption) (*ClosePositionReply, error)
}

type domainClient struct {
	cc *grpc.ClientConn
}

func NewDomainClient(cc *grpc.ClientConn) DomainClient {
	return &domainClient{cc}
}

func (c *domainClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesReply, error) {
	out := new(GetServicesReply)
	err := c.cc.Invoke(ctx, "/grpcdomain.Domain/GetServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainClient) ShareIdentityList(ctx context.Context, in *IdentityListRequest, opts ...grpc.CallOption) (*IdentityListReply, error) {
	out := new(IdentityListReply)
	err := c.cc.Invoke(ctx, "/grpcdomain.Domain/ShareIdentityList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainClient) OpenPosition(ctx context.Context, in *OpenPositionRequest, opts ...grpc.CallOption) (*OpenPositionReply, error) {
	out := new(OpenPositionReply)
	err := c.cc.Invoke(ctx, "/grpcdomain.Domain/OpenPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *domainClient) ClosePosition(ctx context.Context, in *ClosePositionRequest, opts ...grpc.CallOption) (*ClosePositionReply, error) {
	out := new(ClosePositionReply)
	err := c.cc.Invoke(ctx, "/grpcdomain.Domain/ClosePosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DomainServer is the server API for Domain service.
type DomainServer interface {
	// GetServices returns the availible services and their locations
	GetServices(context.Context, *GetServicesRequest) (*GetServicesReply, error)
	// ShareIdentityList Requests the IdentityList which the domain is aware of
	ShareIdentityList(context.Context, *IdentityListRequest) (*IdentityListReply, error)
	// OpenPosition declares the new service which is needed and requests Appointments
	OpenPosition(context.Context, *OpenPositionRequest) (*OpenPositionReply, error)
	// ClosePosition ends an election and informs if the position was awarded
	ClosePosition(context.Context, *ClosePositionRequest) (*ClosePositionReply, error)
}

// UnimplementedDomainServer can be embedded to have forward compatible implementations.
type UnimplementedDomainServer struct {
}

func (*UnimplementedDomainServer) GetServices(ctx context.Context, req *GetServicesRequest) (*GetServicesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (*UnimplementedDomainServer) ShareIdentityList(ctx context.Context, req *IdentityListRequest) (*IdentityListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareIdentityList not implemented")
}
func (*UnimplementedDomainServer) OpenPosition(ctx context.Context, req *OpenPositionRequest) (*OpenPositionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenPosition not implemented")
}
func (*UnimplementedDomainServer) ClosePosition(ctx context.Context, req *ClosePositionRequest) (*ClosePositionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClosePosition not implemented")
}

func RegisterDomainServer(s *grpc.Server, srv DomainServer) {
	s.RegisterService(&_Domain_serviceDesc, srv)
}

func _Domain_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcdomain.Domain/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServer).GetServices(ctx, req.(*GetServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domain_ShareIdentityList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServer).ShareIdentityList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcdomain.Domain/ShareIdentityList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServer).ShareIdentityList(ctx, req.(*IdentityListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domain_OpenPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServer).OpenPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcdomain.Domain/OpenPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServer).OpenPosition(ctx, req.(*OpenPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Domain_ClosePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClosePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DomainServer).ClosePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcdomain.Domain/ClosePosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DomainServer).ClosePosition(ctx, req.(*ClosePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Domain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcdomain.Domain",
	HandlerType: (*DomainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServices",
			Handler:    _Domain_GetServices_Handler,
		},
		{
			MethodName: "ShareIdentityList",
			Handler:    _Domain_ShareIdentityList_Handler,
		},
		{
			MethodName: "OpenPosition",
			Handler:    _Domain_OpenPosition_Handler,
		},
		{
			MethodName: "ClosePosition",
			Handler:    _Domain_ClosePosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "domain.proto",
}
