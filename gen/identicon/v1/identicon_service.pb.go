// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: identicon/v1/identicon_service.proto

package identiconv1

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

// リクエストメッセージの定義
type GenerateBinaryIdenticonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// テキストデータ
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *GenerateBinaryIdenticonRequest) Reset() {
	*x = GenerateBinaryIdenticonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identicon_v1_identicon_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateBinaryIdenticonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateBinaryIdenticonRequest) ProtoMessage() {}

func (x *GenerateBinaryIdenticonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identicon_v1_identicon_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateBinaryIdenticonRequest.ProtoReflect.Descriptor instead.
func (*GenerateBinaryIdenticonRequest) Descriptor() ([]byte, []int) {
	return file_identicon_v1_identicon_service_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateBinaryIdenticonRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// レスポンスメッセージの定義
type GenerateBinaryIdenticonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 生成されたIdenticon画像を含むバイナリデータ
	ImageData []byte `protobuf:"bytes,1,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
}

func (x *GenerateBinaryIdenticonResponse) Reset() {
	*x = GenerateBinaryIdenticonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identicon_v1_identicon_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateBinaryIdenticonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateBinaryIdenticonResponse) ProtoMessage() {}

func (x *GenerateBinaryIdenticonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identicon_v1_identicon_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateBinaryIdenticonResponse.ProtoReflect.Descriptor instead.
func (*GenerateBinaryIdenticonResponse) Descriptor() ([]byte, []int) {
	return file_identicon_v1_identicon_service_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateBinaryIdenticonResponse) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

var File_identicon_v1_identicon_service_proto protoreflect.FileDescriptor

var file_identicon_v1_identicon_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x22, 0x34, 0x0a, 0x1e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x40, 0x0a, 0x1f, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x32, 0x8a, 0x01, 0x0a,
	0x10, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x76, 0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x73, 0x68, 0x69, 0x64, 0x61, 0x6d,
	0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x6f, 0x6e, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_identicon_v1_identicon_service_proto_rawDescOnce sync.Once
	file_identicon_v1_identicon_service_proto_rawDescData = file_identicon_v1_identicon_service_proto_rawDesc
)

func file_identicon_v1_identicon_service_proto_rawDescGZIP() []byte {
	file_identicon_v1_identicon_service_proto_rawDescOnce.Do(func() {
		file_identicon_v1_identicon_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_identicon_v1_identicon_service_proto_rawDescData)
	})
	return file_identicon_v1_identicon_service_proto_rawDescData
}

var file_identicon_v1_identicon_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_identicon_v1_identicon_service_proto_goTypes = []interface{}{
	(*GenerateBinaryIdenticonRequest)(nil),  // 0: identicon.v1.GenerateBinaryIdenticonRequest
	(*GenerateBinaryIdenticonResponse)(nil), // 1: identicon.v1.GenerateBinaryIdenticonResponse
}
var file_identicon_v1_identicon_service_proto_depIdxs = []int32{
	0, // 0: identicon.v1.IdenticonService.GenerateBinaryIdenticon:input_type -> identicon.v1.GenerateBinaryIdenticonRequest
	1, // 1: identicon.v1.IdenticonService.GenerateBinaryIdenticon:output_type -> identicon.v1.GenerateBinaryIdenticonResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_identicon_v1_identicon_service_proto_init() }
func file_identicon_v1_identicon_service_proto_init() {
	if File_identicon_v1_identicon_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identicon_v1_identicon_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateBinaryIdenticonRequest); i {
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
		file_identicon_v1_identicon_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateBinaryIdenticonResponse); i {
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
			RawDescriptor: file_identicon_v1_identicon_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_identicon_v1_identicon_service_proto_goTypes,
		DependencyIndexes: file_identicon_v1_identicon_service_proto_depIdxs,
		MessageInfos:      file_identicon_v1_identicon_service_proto_msgTypes,
	}.Build()
	File_identicon_v1_identicon_service_proto = out.File
	file_identicon_v1_identicon_service_proto_rawDesc = nil
	file_identicon_v1_identicon_service_proto_goTypes = nil
	file_identicon_v1_identicon_service_proto_depIdxs = nil
}
