// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: identity.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Service represents a service running under a domain
type ServiceIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identity is the identity of the service
	Identity *Identity `protobuf:"bytes,1,opt,name=Identity,proto3" json:"Identity,omitempty"`
	// Type is the type of the service
	Type string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *ServiceIdentity) Reset() {
	*x = ServiceIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceIdentity) ProtoMessage() {}

func (x *ServiceIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceIdentity.ProtoReflect.Descriptor instead.
func (*ServiceIdentity) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceIdentity) GetIdentity() *Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *ServiceIdentity) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// Identity holds identifying information for a service, domain, or dominion
type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is a unique identifier
	ID []byte `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Version is the version of code being run
	Version string `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"`
	// Address is the network address where the identity is reachable
	Address *Address `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{1}
}

func (x *Identity) GetID() []byte {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Identity) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Identity) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IP is the IP address which the domain was last known to communicate over
	IP []byte `protobuf:"bytes,5,opt,name=IP,proto3" json:"IP,omitempty"`
	// Port is the port which the domain was last known to belong communicate over
	Port int32 `protobuf:"varint,6,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_identity_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetIP() []byte {
	if x != nil {
		return x.IP
	}
	return nil
}

func (x *Address) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_identity_proto protoreflect.FileDescriptor

var file_identity_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x51, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x08, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x5d, 0x0a, 0x08, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x27, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2d, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6d, 0x62, 0x61, 0x72, 0x7a, 0x65, 0x65, 0x2f, 0x64,
	0x6f, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_identity_proto_rawDescOnce sync.Once
	file_identity_proto_rawDescData = file_identity_proto_rawDesc
)

func file_identity_proto_rawDescGZIP() []byte {
	file_identity_proto_rawDescOnce.Do(func() {
		file_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_identity_proto_rawDescData)
	})
	return file_identity_proto_rawDescData
}

var file_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_identity_proto_goTypes = []interface{}{
	(*ServiceIdentity)(nil), // 0: grpc.ServiceIdentity
	(*Identity)(nil),        // 1: grpc.Identity
	(*Address)(nil),         // 2: grpc.Address
}
var file_identity_proto_depIdxs = []int32{
	1, // 0: grpc.ServiceIdentity.Identity:type_name -> grpc.Identity
	2, // 1: grpc.Identity.Address:type_name -> grpc.Address
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_identity_proto_init() }
func file_identity_proto_init() {
	if File_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceIdentity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_identity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_identity_proto_goTypes,
		DependencyIndexes: file_identity_proto_depIdxs,
		MessageInfos:      file_identity_proto_msgTypes,
	}.Build()
	File_identity_proto = out.File
	file_identity_proto_rawDesc = nil
	file_identity_proto_goTypes = nil
	file_identity_proto_depIdxs = nil
}
