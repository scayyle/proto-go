// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: item_catalog_service/item.proto

package item_catalog_service

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateTestItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTestItemsRequest) Reset() {
	*x = CreateTestItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_catalog_service_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestItemsRequest) ProtoMessage() {}

func (x *CreateTestItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_item_catalog_service_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestItemsRequest.ProtoReflect.Descriptor instead.
func (*CreateTestItemsRequest) Descriptor() ([]byte, []int) {
	return file_item_catalog_service_item_proto_rawDescGZIP(), []int{0}
}

type CreateTestItemsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTestItemsReply) Reset() {
	*x = CreateTestItemsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_catalog_service_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestItemsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestItemsReply) ProtoMessage() {}

func (x *CreateTestItemsReply) ProtoReflect() protoreflect.Message {
	mi := &file_item_catalog_service_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestItemsReply.ProtoReflect.Descriptor instead.
func (*CreateTestItemsReply) Descriptor() ([]byte, []int) {
	return file_item_catalog_service_item_proto_rawDescGZIP(), []int{1}
}

var File_item_catalog_service_item_proto protoreflect.FileDescriptor

var file_item_catalog_service_item_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x73, 0x0a, 0x04, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x6b, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x2c, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x39,
	0x5a, 0x16, 0x2e, 0x3b, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xaa, 0x02, 0x1e, 0x49, 0x74, 0x65, 0x6d, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x52,
	0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_item_catalog_service_item_proto_rawDescOnce sync.Once
	file_item_catalog_service_item_proto_rawDescData = file_item_catalog_service_item_proto_rawDesc
)

func file_item_catalog_service_item_proto_rawDescGZIP() []byte {
	file_item_catalog_service_item_proto_rawDescOnce.Do(func() {
		file_item_catalog_service_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_item_catalog_service_item_proto_rawDescData)
	})
	return file_item_catalog_service_item_proto_rawDescData
}

var file_item_catalog_service_item_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_item_catalog_service_item_proto_goTypes = []interface{}{
	(*CreateTestItemsRequest)(nil), // 0: item_catalog_service.CreateTestItemsRequest
	(*CreateTestItemsReply)(nil),   // 1: item_catalog_service.CreateTestItemsReply
}
var file_item_catalog_service_item_proto_depIdxs = []int32{
	0, // 0: item_catalog_service.Item.CreateTestItems:input_type -> item_catalog_service.CreateTestItemsRequest
	1, // 1: item_catalog_service.Item.CreateTestItems:output_type -> item_catalog_service.CreateTestItemsReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_item_catalog_service_item_proto_init() }
func file_item_catalog_service_item_proto_init() {
	if File_item_catalog_service_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_item_catalog_service_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestItemsRequest); i {
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
		file_item_catalog_service_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestItemsReply); i {
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
			RawDescriptor: file_item_catalog_service_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_item_catalog_service_item_proto_goTypes,
		DependencyIndexes: file_item_catalog_service_item_proto_depIdxs,
		MessageInfos:      file_item_catalog_service_item_proto_msgTypes,
	}.Build()
	File_item_catalog_service_item_proto = out.File
	file_item_catalog_service_item_proto_rawDesc = nil
	file_item_catalog_service_item_proto_goTypes = nil
	file_item_catalog_service_item_proto_depIdxs = nil
}
