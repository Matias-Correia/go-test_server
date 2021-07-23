// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: protologs/logs.proto

package helloworld

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Log definition
type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockID        string               `protobuf:"bytes,1,opt,name=blockID,proto3" json:"blockID,omitempty"`
	RequestDelay   *durationpb.Duration `protobuf:"bytes,2,opt,name=request_delay,json=requestDelay,proto3" json:"request_delay,omitempty"`
	BlockDelivered bool                 `protobuf:"varint,3,opt,name=blockDelivered,proto3" json:"blockDelivered,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protologs_logs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_protologs_logs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_protologs_logs_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetBlockID() string {
	if x != nil {
		return x.BlockID
	}
	return ""
}

func (x *Log) GetRequestDelay() *durationpb.Duration {
	if x != nil {
		return x.RequestDelay
	}
	return nil
}

func (x *Log) GetBlockDelivered() bool {
	if x != nil {
		return x.BlockDelivered
	}
	return false
}

var File_protologs_logs_proto protoreflect.FileDescriptor

var file_protologs_logs_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6c, 0x6f, 0x67,
	0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87,
	0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44,
	0x12, 0x3e, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79,
	0x12, 0x26, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x32, 0x43, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x54,
	0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x4c,
	0x6f, 0x67, 0x73, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6c, 0x6f, 0x67, 0x73, 0x2e,
	0x4c, 0x6f, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protologs_logs_proto_rawDescOnce sync.Once
	file_protologs_logs_proto_rawDescData = file_protologs_logs_proto_rawDesc
)

func file_protologs_logs_proto_rawDescGZIP() []byte {
	file_protologs_logs_proto_rawDescOnce.Do(func() {
		file_protologs_logs_proto_rawDescData = protoimpl.X.CompressGZIP(file_protologs_logs_proto_rawDescData)
	})
	return file_protologs_logs_proto_rawDescData
}

var file_protologs_logs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protologs_logs_proto_goTypes = []interface{}{
	(*Log)(nil),                 // 0: protologs.Log
	(*durationpb.Duration)(nil), // 1: google.protobuf.Duration
	(*emptypb.Empty)(nil),       // 2: google.protobuf.Empty
}
var file_protologs_logs_proto_depIdxs = []int32{
	1, // 0: protologs.Log.request_delay:type_name -> google.protobuf.Duration
	0, // 1: protologs.LogTestData.SendLogs:input_type -> protologs.Log
	2, // 2: protologs.LogTestData.SendLogs:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protologs_logs_proto_init() }
func file_protologs_logs_proto_init() {
	if File_protologs_logs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protologs_logs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_protologs_logs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protologs_logs_proto_goTypes,
		DependencyIndexes: file_protologs_logs_proto_depIdxs,
		MessageInfos:      file_protologs_logs_proto_msgTypes,
	}.Build()
	File_protologs_logs_proto = out.File
	file_protologs_logs_proto_rawDesc = nil
	file_protologs_logs_proto_goTypes = nil
	file_protologs_logs_proto_depIdxs = nil
}