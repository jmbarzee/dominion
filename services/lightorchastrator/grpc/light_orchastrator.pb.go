// Code generated by protoc-gen-go. DO NOT EDIT.
// source: light_orchastrator.proto

package grpcdomain

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// SubscribeLightsRequest contains the identifying information from the subscriber
type SubscribeLightsRequest struct {
	// ServiceName identifies the service which is subscribing to the service
	ServiceName          string   `protobuf:"bytes,1,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeLightsRequest) Reset()         { *m = SubscribeLightsRequest{} }
func (m *SubscribeLightsRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeLightsRequest) ProtoMessage()    {}
func (*SubscribeLightsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0a167111198e4c0, []int{0}
}

func (m *SubscribeLightsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeLightsRequest.Unmarshal(m, b)
}
func (m *SubscribeLightsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeLightsRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeLightsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeLightsRequest.Merge(m, src)
}
func (m *SubscribeLightsRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeLightsRequest.Size(m)
}
func (m *SubscribeLightsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeLightsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeLightsRequest proto.InternalMessageInfo

func (m *SubscribeLightsRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

// SubscribeLightsReply contains the time and series of RGBA colors which should be displayed
type SubscribeLightsReply struct {
	// DisplayTime is the time which the lights should be applied (UnixNano)
	DisplayTime int64 `protobuf:"varint,1,opt,name=DisplayTime,proto3" json:"DisplayTime,omitempty"`
	// Colors are the series of colors which should be displayed
	Colors               []int32  `protobuf:"varint,2,rep,packed,name=Colors,proto3" json:"Colors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeLightsReply) Reset()         { *m = SubscribeLightsReply{} }
func (m *SubscribeLightsReply) String() string { return proto.CompactTextString(m) }
func (*SubscribeLightsReply) ProtoMessage()    {}
func (*SubscribeLightsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0a167111198e4c0, []int{1}
}

func (m *SubscribeLightsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeLightsReply.Unmarshal(m, b)
}
func (m *SubscribeLightsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeLightsReply.Marshal(b, m, deterministic)
}
func (m *SubscribeLightsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeLightsReply.Merge(m, src)
}
func (m *SubscribeLightsReply) XXX_Size() int {
	return xxx_messageInfo_SubscribeLightsReply.Size(m)
}
func (m *SubscribeLightsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeLightsReply.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeLightsReply proto.InternalMessageInfo

func (m *SubscribeLightsReply) GetDisplayTime() int64 {
	if m != nil {
		return m.DisplayTime
	}
	return 0
}

func (m *SubscribeLightsReply) GetColors() []int32 {
	if m != nil {
		return m.Colors
	}
	return nil
}

func init() {
	proto.RegisterType((*SubscribeLightsRequest)(nil), "grpcdomain.SubscribeLightsRequest")
	proto.RegisterType((*SubscribeLightsReply)(nil), "grpcdomain.SubscribeLightsReply")
}

func init() { proto.RegisterFile("light_orchastrator.proto", fileDescriptor_d0a167111198e4c0) }

var fileDescriptor_d0a167111198e4c0 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xc9, 0x4c, 0xcf,
	0x28, 0x89, 0xcf, 0x2f, 0x4a, 0xce, 0x48, 0x2c, 0x2e, 0x29, 0x4a, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x2f, 0x2a, 0x48, 0x4e, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc,
	0x53, 0xb2, 0xe2, 0x12, 0x0b, 0x2e, 0x4d, 0x2a, 0x4e, 0x2e, 0xca, 0x4c, 0x4a, 0xf5, 0x01, 0x69,
	0x28, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52, 0xe0, 0xe2, 0x0e, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xf5, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x16,
	0x52, 0x0a, 0xe0, 0x12, 0xc1, 0xd0, 0x5b, 0x90, 0x53, 0x09, 0xd2, 0xe9, 0x92, 0x59, 0x5c, 0x90,
	0x93, 0x58, 0x19, 0x92, 0x09, 0xd5, 0xc9, 0x1c, 0x84, 0x2c, 0x24, 0x24, 0xc6, 0xc5, 0xe6, 0x9c,
	0x9f, 0x93, 0x5f, 0x54, 0x2c, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x1a, 0x04, 0xe5, 0x19, 0x15, 0x72,
	0x09, 0x81, 0x0d, 0xf2, 0x07, 0x39, 0xba, 0x08, 0xea, 0x6a, 0xa1, 0x68, 0x2e, 0x7e, 0x34, 0x7b,
	0x84, 0x94, 0xf4, 0x10, 0x7e, 0xd0, 0xc3, 0xee, 0x01, 0x29, 0x05, 0xbc, 0x6a, 0x0a, 0x72, 0x2a,
	0x95, 0x18, 0x0c, 0x18, 0x93, 0xd8, 0xc0, 0x61, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x7f,
	0x7c, 0x6d, 0xad, 0x2f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LightOrcharstratorClient is the client API for LightOrcharstrator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LightOrcharstratorClient interface {
	// DisplayLights initiates
	SubscribeLights(ctx context.Context, in *SubscribeLightsRequest, opts ...grpc.CallOption) (LightOrcharstrator_SubscribeLightsClient, error)
}

type lightOrcharstratorClient struct {
	cc *grpc.ClientConn
}

func NewLightOrcharstratorClient(cc *grpc.ClientConn) LightOrcharstratorClient {
	return &lightOrcharstratorClient{cc}
}

func (c *lightOrcharstratorClient) SubscribeLights(ctx context.Context, in *SubscribeLightsRequest, opts ...grpc.CallOption) (LightOrcharstrator_SubscribeLightsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LightOrcharstrator_serviceDesc.Streams[0], "/grpcdomain.LightOrcharstrator/SubscribeLights", opts...)
	if err != nil {
		return nil, err
	}
	x := &lightOrcharstratorSubscribeLightsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LightOrcharstrator_SubscribeLightsClient interface {
	Recv() (*SubscribeLightsReply, error)
	grpc.ClientStream
}

type lightOrcharstratorSubscribeLightsClient struct {
	grpc.ClientStream
}

func (x *lightOrcharstratorSubscribeLightsClient) Recv() (*SubscribeLightsReply, error) {
	m := new(SubscribeLightsReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LightOrcharstratorServer is the server API for LightOrcharstrator service.
type LightOrcharstratorServer interface {
	// DisplayLights initiates
	SubscribeLights(*SubscribeLightsRequest, LightOrcharstrator_SubscribeLightsServer) error
}

// UnimplementedLightOrcharstratorServer can be embedded to have forward compatible implementations.
type UnimplementedLightOrcharstratorServer struct {
}

func (*UnimplementedLightOrcharstratorServer) SubscribeLights(req *SubscribeLightsRequest, srv LightOrcharstrator_SubscribeLightsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeLights not implemented")
}

func RegisterLightOrcharstratorServer(s *grpc.Server, srv LightOrcharstratorServer) {
	s.RegisterService(&_LightOrcharstrator_serviceDesc, srv)
}

func _LightOrcharstrator_SubscribeLights_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeLightsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LightOrcharstratorServer).SubscribeLights(m, &lightOrcharstratorSubscribeLightsServer{stream})
}

type LightOrcharstrator_SubscribeLightsServer interface {
	Send(*SubscribeLightsReply) error
	grpc.ServerStream
}

type lightOrcharstratorSubscribeLightsServer struct {
	grpc.ServerStream
}

func (x *lightOrcharstratorSubscribeLightsServer) Send(m *SubscribeLightsReply) error {
	return x.ServerStream.SendMsg(m)
}

var _LightOrcharstrator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcdomain.LightOrcharstrator",
	HandlerType: (*LightOrcharstratorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeLights",
			Handler:       _LightOrcharstrator_SubscribeLights_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "light_orchastrator.proto",
}
