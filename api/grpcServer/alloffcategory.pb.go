// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/grpcServer/protos/alloffcategory.proto

package grpcServer

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

type AlloffCategoryMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Keyname    string `protobuf:"bytes,3,opt,name=keyname,proto3" json:"keyname,omitempty"`
	Level      int32  `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	ParentId   string `protobuf:"bytes,5,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	ImgUrl     string `protobuf:"bytes,6,opt,name=img_url,json=imgUrl,proto3" json:"img_url,omitempty"`
}

func (x *AlloffCategoryMessage) Reset() {
	*x = AlloffCategoryMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_alloffcategory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlloffCategoryMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlloffCategoryMessage) ProtoMessage() {}

func (x *AlloffCategoryMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_alloffcategory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlloffCategoryMessage.ProtoReflect.Descriptor instead.
func (*AlloffCategoryMessage) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_alloffcategory_proto_rawDescGZIP(), []int{0}
}

func (x *AlloffCategoryMessage) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *AlloffCategoryMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlloffCategoryMessage) GetKeyname() string {
	if x != nil {
		return x.Keyname
	}
	return ""
}

func (x *AlloffCategoryMessage) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *AlloffCategoryMessage) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *AlloffCategoryMessage) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

type ListAlloffCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
}

func (x *ListAlloffCategoryRequest) Reset() {
	*x = ListAlloffCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_alloffcategory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlloffCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlloffCategoryRequest) ProtoMessage() {}

func (x *ListAlloffCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_alloffcategory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlloffCategoryRequest.ProtoReflect.Descriptor instead.
func (*ListAlloffCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_alloffcategory_proto_rawDescGZIP(), []int{1}
}

func (x *ListAlloffCategoryRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

type ListAlloffCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*AlloffCategoryMessage `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *ListAlloffCategoryResponse) Reset() {
	*x = ListAlloffCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_alloffcategory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlloffCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlloffCategoryResponse) ProtoMessage() {}

func (x *ListAlloffCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_alloffcategory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlloffCategoryResponse.ProtoReflect.Descriptor instead.
func (*ListAlloffCategoryResponse) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_alloffcategory_proto_rawDescGZIP(), []int{2}
}

func (x *ListAlloffCategoryResponse) GetCategories() []*AlloffCategoryMessage {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_api_grpcServer_protos_alloffcategory_proto protoreflect.FileDescriptor

var file_api_grpcServer_protos_alloffcategory_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72,
	0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0xb2, 0x01, 0x0a, 0x15, 0x41, 0x6c, 0x6c,
	0x6f, 0x66, 0x66, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x22, 0x38, 0x0a,
	0x19, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x32, 0x75, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x6f,
	0x66, 0x66, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x63, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65,
	0x73, 0x73, 0x62, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpcServer_protos_alloffcategory_proto_rawDescOnce sync.Once
	file_api_grpcServer_protos_alloffcategory_proto_rawDescData = file_api_grpcServer_protos_alloffcategory_proto_rawDesc
)

func file_api_grpcServer_protos_alloffcategory_proto_rawDescGZIP() []byte {
	file_api_grpcServer_protos_alloffcategory_proto_rawDescOnce.Do(func() {
		file_api_grpcServer_protos_alloffcategory_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpcServer_protos_alloffcategory_proto_rawDescData)
	})
	return file_api_grpcServer_protos_alloffcategory_proto_rawDescData
}

var file_api_grpcServer_protos_alloffcategory_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_grpcServer_protos_alloffcategory_proto_goTypes = []interface{}{
	(*AlloffCategoryMessage)(nil),      // 0: grpcServer.AlloffCategoryMessage
	(*ListAlloffCategoryRequest)(nil),  // 1: grpcServer.ListAlloffCategoryRequest
	(*ListAlloffCategoryResponse)(nil), // 2: grpcServer.ListAlloffCategoryResponse
}
var file_api_grpcServer_protos_alloffcategory_proto_depIdxs = []int32{
	0, // 0: grpcServer.ListAlloffCategoryResponse.categories:type_name -> grpcServer.AlloffCategoryMessage
	1, // 1: grpcServer.AlloffCategory.ListAlloffCategory:input_type -> grpcServer.ListAlloffCategoryRequest
	2, // 2: grpcServer.AlloffCategory.ListAlloffCategory:output_type -> grpcServer.ListAlloffCategoryResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_grpcServer_protos_alloffcategory_proto_init() }
func file_api_grpcServer_protos_alloffcategory_proto_init() {
	if File_api_grpcServer_protos_alloffcategory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpcServer_protos_alloffcategory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlloffCategoryMessage); i {
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
		file_api_grpcServer_protos_alloffcategory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlloffCategoryRequest); i {
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
		file_api_grpcServer_protos_alloffcategory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlloffCategoryResponse); i {
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
			RawDescriptor: file_api_grpcServer_protos_alloffcategory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpcServer_protos_alloffcategory_proto_goTypes,
		DependencyIndexes: file_api_grpcServer_protos_alloffcategory_proto_depIdxs,
		MessageInfos:      file_api_grpcServer_protos_alloffcategory_proto_msgTypes,
	}.Build()
	File_api_grpcServer_protos_alloffcategory_proto = out.File
	file_api_grpcServer_protos_alloffcategory_proto_rawDesc = nil
	file_api_grpcServer_protos_alloffcategory_proto_goTypes = nil
	file_api_grpcServer_protos_alloffcategory_proto_depIdxs = nil
}
