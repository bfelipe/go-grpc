// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: streamer.proto

package pb

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

type Result int32

const (
	Result_CLEAR   Result = 0
	Result_ANOMALY Result = 1
	Result_UNKOWN  Result = 2
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0: "CLEAR",
		1: "ANOMALY",
		2: "UNKOWN",
	}
	Result_value = map[string]int32{
		"CLEAR":   0,
		"ANOMALY": 1,
		"UNKOWN":  2,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file_streamer_proto_enumTypes[0].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_streamer_proto_enumTypes[0]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_streamer_proto_rawDescGZIP(), []int{0}
}

type ScanningRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RobotId      string `protobuf:"bytes,1,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"`
	ScanningData int64  `protobuf:"varint,2,opt,name=scanning_data,json=scanningData,proto3" json:"scanning_data,omitempty"`
}

func (x *ScanningRequest) Reset() {
	*x = ScanningRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streamer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanningRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanningRequest) ProtoMessage() {}

func (x *ScanningRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streamer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanningRequest.ProtoReflect.Descriptor instead.
func (*ScanningRequest) Descriptor() ([]byte, []int) {
	return file_streamer_proto_rawDescGZIP(), []int{0}
}

func (x *ScanningRequest) GetRobotId() string {
	if x != nil {
		return x.RobotId
	}
	return ""
}

func (x *ScanningRequest) GetScanningData() int64 {
	if x != nil {
		return x.ScanningData
	}
	return 0
}

type ScanningResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result Result `protobuf:"varint,1,opt,name=result,proto3,enum=pb.Result" json:"result,omitempty"`
}

func (x *ScanningResponse) Reset() {
	*x = ScanningResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streamer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanningResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanningResponse) ProtoMessage() {}

func (x *ScanningResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streamer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanningResponse.ProtoReflect.Descriptor instead.
func (*ScanningResponse) Descriptor() ([]byte, []int) {
	return file_streamer_proto_rawDescGZIP(), []int{1}
}

func (x *ScanningResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_CLEAR
}

var File_streamer_proto protoreflect.FileDescriptor

var file_streamer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x51, 0x0a, 0x0f, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x22, 0x36, 0x0a, 0x10, 0x53, 0x63, 0x61, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a,
	0x2c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x45,
	0x41, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x4f, 0x4d, 0x41, 0x4c, 0x59, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x32, 0x3e, 0x0a,
	0x07, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x04, 0x53, 0x63, 0x61, 0x6e,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x14, 0x5a,
	0x12, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_streamer_proto_rawDescOnce sync.Once
	file_streamer_proto_rawDescData = file_streamer_proto_rawDesc
)

func file_streamer_proto_rawDescGZIP() []byte {
	file_streamer_proto_rawDescOnce.Do(func() {
		file_streamer_proto_rawDescData = protoimpl.X.CompressGZIP(file_streamer_proto_rawDescData)
	})
	return file_streamer_proto_rawDescData
}

var file_streamer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_streamer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_streamer_proto_goTypes = []interface{}{
	(Result)(0),              // 0: pb.Result
	(*ScanningRequest)(nil),  // 1: pb.ScanningRequest
	(*ScanningResponse)(nil), // 2: pb.ScanningResponse
}
var file_streamer_proto_depIdxs = []int32{
	0, // 0: pb.ScanningResponse.result:type_name -> pb.Result
	1, // 1: pb.Scanner.Scan:input_type -> pb.ScanningRequest
	2, // 2: pb.Scanner.Scan:output_type -> pb.ScanningResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_streamer_proto_init() }
func file_streamer_proto_init() {
	if File_streamer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_streamer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanningRequest); i {
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
		file_streamer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanningResponse); i {
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
			RawDescriptor: file_streamer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_streamer_proto_goTypes,
		DependencyIndexes: file_streamer_proto_depIdxs,
		EnumInfos:         file_streamer_proto_enumTypes,
		MessageInfos:      file_streamer_proto_msgTypes,
	}.Build()
	File_streamer_proto = out.File
	file_streamer_proto_rawDesc = nil
	file_streamer_proto_goTypes = nil
	file_streamer_proto_depIdxs = nil
}
