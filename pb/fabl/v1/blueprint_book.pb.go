// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: fabl/v1/blueprint_book.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BlueprintBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     uint64                `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Item        string                `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Label       string                `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	ActiveIndex uint64                `protobuf:"varint,6,opt,name=active_index,json=activeIndex,proto3" json:"active_index,omitempty"`
	Blueprints  []*BlueprintBookEntry `protobuf:"bytes,7,rep,name=blueprints,proto3" json:"blueprints,omitempty"`
}

func (x *BlueprintBook) Reset() {
	*x = BlueprintBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fabl_v1_blueprint_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlueprintBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlueprintBook) ProtoMessage() {}

func (x *BlueprintBook) ProtoReflect() protoreflect.Message {
	mi := &file_fabl_v1_blueprint_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlueprintBook.ProtoReflect.Descriptor instead.
func (*BlueprintBook) Descriptor() ([]byte, []int) {
	return file_fabl_v1_blueprint_book_proto_rawDescGZIP(), []int{0}
}

func (x *BlueprintBook) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *BlueprintBook) GetItem() string {
	if x != nil {
		return x.Item
	}
	return ""
}

func (x *BlueprintBook) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *BlueprintBook) GetActiveIndex() uint64 {
	if x != nil {
		return x.ActiveIndex
	}
	return 0
}

func (x *BlueprintBook) GetBlueprints() []*BlueprintBookEntry {
	if x != nil {
		return x.Blueprints
	}
	return nil
}

var File_fabl_v1_blueprint_book_proto protoreflect.FileDescriptor

var file_fabl_v1_blueprint_book_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x66, 0x61, 0x62, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x66, 0x61, 0x62, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x66, 0x61, 0x62, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x0d,
	0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x3b, 0x0a, 0x0a, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x61, 0x62, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x62, 0x6c, 0x75, 0x65, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x73, 0x42, 0x1c, 0x5a, 0x1a, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x61, 0x62, 0x6c, 0x2e, 0x61, 0x70,
	0x70, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x61, 0x62, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fabl_v1_blueprint_book_proto_rawDescOnce sync.Once
	file_fabl_v1_blueprint_book_proto_rawDescData = file_fabl_v1_blueprint_book_proto_rawDesc
)

func file_fabl_v1_blueprint_book_proto_rawDescGZIP() []byte {
	file_fabl_v1_blueprint_book_proto_rawDescOnce.Do(func() {
		file_fabl_v1_blueprint_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_fabl_v1_blueprint_book_proto_rawDescData)
	})
	return file_fabl_v1_blueprint_book_proto_rawDescData
}

var file_fabl_v1_blueprint_book_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fabl_v1_blueprint_book_proto_goTypes = []interface{}{
	(*BlueprintBook)(nil),      // 0: fabl.v1.BlueprintBook
	(*BlueprintBookEntry)(nil), // 1: fabl.v1.BlueprintBookEntry
}
var file_fabl_v1_blueprint_book_proto_depIdxs = []int32{
	1, // 0: fabl.v1.BlueprintBook.blueprints:type_name -> fabl.v1.BlueprintBookEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_fabl_v1_blueprint_book_proto_init() }
func file_fabl_v1_blueprint_book_proto_init() {
	if File_fabl_v1_blueprint_book_proto != nil {
		return
	}
	file_fabl_v1_blueprint_book_entry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fabl_v1_blueprint_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlueprintBook); i {
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
			RawDescriptor: file_fabl_v1_blueprint_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fabl_v1_blueprint_book_proto_goTypes,
		DependencyIndexes: file_fabl_v1_blueprint_book_proto_depIdxs,
		MessageInfos:      file_fabl_v1_blueprint_book_proto_msgTypes,
	}.Build()
	File_fabl_v1_blueprint_book_proto = out.File
	file_fabl_v1_blueprint_book_proto_rawDesc = nil
	file_fabl_v1_blueprint_book_proto_goTypes = nil
	file_fabl_v1_blueprint_book_proto_depIdxs = nil
}
