// Code generated by protoc-gen-go. DO NOT EDIT.
// source: get_services.proto

package grpcdomain

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// GetServicesRequest contains the name of the service the client is searching for
type GetServicesRequest struct {
	// Name is the name of the service which is requested
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServicesRequest) Reset()         { *m = GetServicesRequest{} }
func (m *GetServicesRequest) String() string { return proto.CompactTextString(m) }
func (*GetServicesRequest) ProtoMessage()    {}
func (*GetServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec18c318270b8d2, []int{0}
}

func (m *GetServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServicesRequest.Unmarshal(m, b)
}
func (m *GetServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServicesRequest.Marshal(b, m, deterministic)
}
func (m *GetServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServicesRequest.Merge(m, src)
}
func (m *GetServicesRequest) XXX_Size() int {
	return xxx_messageInfo_GetServicesRequest.Size(m)
}
func (m *GetServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServicesRequest proto.InternalMessageInfo

func (m *GetServicesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// GetServicesReply contains the list of known addresses hosting the requested service
type GetServicesReply struct {
	// Addresses is the list of addresses of the requested service
	Addresses            []string `protobuf:"bytes,1,rep,name=Addresses,proto3" json:"Addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServicesReply) Reset()         { *m = GetServicesReply{} }
func (m *GetServicesReply) String() string { return proto.CompactTextString(m) }
func (*GetServicesReply) ProtoMessage()    {}
func (*GetServicesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec18c318270b8d2, []int{1}
}

func (m *GetServicesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServicesReply.Unmarshal(m, b)
}
func (m *GetServicesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServicesReply.Marshal(b, m, deterministic)
}
func (m *GetServicesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServicesReply.Merge(m, src)
}
func (m *GetServicesReply) XXX_Size() int {
	return xxx_messageInfo_GetServicesReply.Size(m)
}
func (m *GetServicesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServicesReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetServicesReply proto.InternalMessageInfo

func (m *GetServicesReply) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func init() {
	proto.RegisterType((*GetServicesRequest)(nil), "grpcdomain.GetServicesRequest")
	proto.RegisterType((*GetServicesReply)(nil), "grpcdomain.GetServicesReply")
}

func init() { proto.RegisterFile("get_services.proto", fileDescriptor_7ec18c318270b8d2) }

var fileDescriptor_7ec18c318270b8d2 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x4f, 0x2d, 0x89,
	0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x4a, 0x2f, 0x2a, 0x48, 0x4e, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x93, 0xe2, 0xcd, 0x4f, 0xca, 0x4a,
	0x4d, 0x2e, 0x81, 0x4a, 0x29, 0x69, 0x70, 0x09, 0xb9, 0xa7, 0x96, 0x04, 0x43, 0xd5, 0x07, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xf8, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x06, 0x5c, 0x02, 0x28, 0x2a, 0x0b, 0x72, 0x2a, 0x85,
	0x64, 0xb8, 0x38, 0x1d, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x53, 0x8b, 0x25, 0x18, 0x15, 0x98,
	0x35, 0x38, 0x83, 0x10, 0x02, 0x49, 0x6c, 0x60, 0x2b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xe9, 0xdc, 0x28, 0x6a, 0x93, 0x00, 0x00, 0x00,
}