// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: lookup/v1/lookup.proto

package v1

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

type ObjectId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ObjectId) Reset() {
	*x = ObjectId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lookup_v1_lookup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectId) ProtoMessage() {}

func (x *ObjectId) ProtoReflect() protoreflect.Message {
	mi := &file_lookup_v1_lookup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectId.ProtoReflect.Descriptor instead.
func (*ObjectId) Descriptor() ([]byte, []int) {
	return file_lookup_v1_lookup_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectId) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ObjectId) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_lookup_v1_lookup_proto protoreflect.FileDescriptor

var file_lookup_v1_lookup_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x2e, 0x76, 0x31, 0x22, 0x2e, 0x0a, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x69, 0x72, 0x79, 0x63, 0x68, 0x75, 0x6b, 0x79, 0x75, 0x72, 0x69, 0x69, 0x2f,
	0x77, 0x61, 0x73, 0x6b, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lookup_v1_lookup_proto_rawDescOnce sync.Once
	file_lookup_v1_lookup_proto_rawDescData = file_lookup_v1_lookup_proto_rawDesc
)

func file_lookup_v1_lookup_proto_rawDescGZIP() []byte {
	file_lookup_v1_lookup_proto_rawDescOnce.Do(func() {
		file_lookup_v1_lookup_proto_rawDescData = protoimpl.X.CompressGZIP(file_lookup_v1_lookup_proto_rawDescData)
	})
	return file_lookup_v1_lookup_proto_rawDescData
}

var file_lookup_v1_lookup_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lookup_v1_lookup_proto_goTypes = []interface{}{
	(*ObjectId)(nil), // 0: lookup.v1.ObjectId
}
var file_lookup_v1_lookup_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lookup_v1_lookup_proto_init() }
func file_lookup_v1_lookup_proto_init() {
	if File_lookup_v1_lookup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lookup_v1_lookup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectId); i {
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
			RawDescriptor: file_lookup_v1_lookup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lookup_v1_lookup_proto_goTypes,
		DependencyIndexes: file_lookup_v1_lookup_proto_depIdxs,
		MessageInfos:      file_lookup_v1_lookup_proto_msgTypes,
	}.Build()
	File_lookup_v1_lookup_proto = out.File
	file_lookup_v1_lookup_proto_rawDesc = nil
	file_lookup_v1_lookup_proto_goTypes = nil
	file_lookup_v1_lookup_proto_depIdxs = nil
}
