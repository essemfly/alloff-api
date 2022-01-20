// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/grpcServer/protos/productGroup.proto

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

// Product Group Service
type GetProductGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductGroupId string `protobuf:"bytes,1,opt,name=product_group_id,json=productGroupId,proto3" json:"product_group_id,omitempty"`
}

func (x *GetProductGroupRequest) Reset() {
	*x = GetProductGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductGroupRequest) ProtoMessage() {}

func (x *GetProductGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductGroupRequest.ProtoReflect.Descriptor instead.
func (*GetProductGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_productGroup_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductGroupRequest) GetProductGroupId() string {
	if x != nil {
		return x.ProductGroupId
	}
	return ""
}

type GetProductGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductGroup *ProductGroupMessage `protobuf:"bytes,1,opt,name=product_group,json=productGroup,proto3" json:"product_group,omitempty"`
}

func (x *GetProductGroupResponse) Reset() {
	*x = GetProductGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductGroupResponse) ProtoMessage() {}

func (x *GetProductGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductGroupResponse.ProtoReflect.Descriptor instead.
func (*GetProductGroupResponse) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_productGroup_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductGroupResponse) GetProductGroup() *ProductGroupMessage {
	if x != nil {
		return x.ProductGroup
	}
	return nil
}

type CreateProductGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	ShortTitle  string   `protobuf:"bytes,2,opt,name=short_title,json=shortTitle,proto3" json:"short_title,omitempty"`
	Instruction []string `protobuf:"bytes,3,rep,name=instruction,proto3" json:"instruction,omitempty"`
	ImageUrl    string   `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	StartTime   string   `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	FinishTime  string   `protobuf:"bytes,6,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
}

func (x *CreateProductGroupRequest) Reset() {
	*x = CreateProductGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductGroupRequest) ProtoMessage() {}

func (x *CreateProductGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateProductGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_productGroup_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProductGroupRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateProductGroupRequest) GetShortTitle() string {
	if x != nil {
		return x.ShortTitle
	}
	return ""
}

func (x *CreateProductGroupRequest) GetInstruction() []string {
	if x != nil {
		return x.Instruction
	}
	return nil
}

func (x *CreateProductGroupRequest) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CreateProductGroupRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *CreateProductGroupRequest) GetFinishTime() string {
	if x != nil {
		return x.FinishTime
	}
	return ""
}

type CreateProductGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pg *ProductGroupMessage `protobuf:"bytes,1,opt,name=pg,proto3" json:"pg,omitempty"`
}

func (x *CreateProductGroupResponse) Reset() {
	*x = CreateProductGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductGroupResponse) ProtoMessage() {}

func (x *CreateProductGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateProductGroupResponse) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_productGroup_proto_rawDescGZIP(), []int{3}
}

func (x *CreateProductGroupResponse) GetPg() *ProductGroupMessage {
	if x != nil {
		return x.Pg
	}
	return nil
}

type ListProductGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProductGroupsRequest) Reset() {
	*x = ListProductGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductGroupsRequest) ProtoMessage() {}

func (x *ListProductGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductGroupsRequest.ProtoReflect.Descriptor instead.
func (*ListProductGroupsRequest) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_productGroup_proto_rawDescGZIP(), []int{4}
}

type ListProductGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pgs []*ProductGroupMessage `protobuf:"bytes,1,rep,name=pgs,proto3" json:"pgs,omitempty"`
}

func (x *ListProductGroupsResponse) Reset() {
	*x = ListProductGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductGroupsResponse) ProtoMessage() {}

func (x *ListProductGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductGroupsResponse.ProtoReflect.Descriptor instead.
func (*ListProductGroupsResponse) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_productGroup_proto_rawDescGZIP(), []int{5}
}

func (x *ListProductGroupsResponse) GetPgs() []*ProductGroupMessage {
	if x != nil {
		return x.Pgs
	}
	return nil
}

type PushProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductGroupId string   `protobuf:"bytes,1,opt,name=product_group_id,json=productGroupId,proto3" json:"product_group_id,omitempty"`
	ProductId      []string `protobuf:"bytes,2,rep,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *PushProductsRequest) Reset() {
	*x = PushProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushProductsRequest) ProtoMessage() {}

func (x *PushProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushProductsRequest.ProtoReflect.Descriptor instead.
func (*PushProductsRequest) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_productGroup_proto_rawDescGZIP(), []int{6}
}

func (x *PushProductsRequest) GetProductGroupId() string {
	if x != nil {
		return x.ProductGroupId
	}
	return ""
}

func (x *PushProductsRequest) GetProductId() []string {
	if x != nil {
		return x.ProductId
	}
	return nil
}

type PushProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pg *ProductGroupMessage `protobuf:"bytes,1,opt,name=pg,proto3" json:"pg,omitempty"`
}

func (x *PushProductsResponse) Reset() {
	*x = PushProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushProductsResponse) ProtoMessage() {}

func (x *PushProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushProductsResponse.ProtoReflect.Descriptor instead.
func (*PushProductsResponse) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_productGroup_proto_rawDescGZIP(), []int{7}
}

func (x *PushProductsResponse) GetPg() *ProductGroupMessage {
	if x != nil {
		return x.Pg
	}
	return nil
}

type ProductInGroupMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product  *ProductMessage `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Priority int32           `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *ProductInGroupMessage) Reset() {
	*x = ProductInGroupMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductInGroupMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductInGroupMessage) ProtoMessage() {}

func (x *ProductInGroupMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductInGroupMessage.ProtoReflect.Descriptor instead.
func (*ProductInGroupMessage) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_productGroup_proto_rawDescGZIP(), []int{8}
}

func (x *ProductInGroupMessage) GetProduct() *ProductMessage {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *ProductInGroupMessage) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type ProductGroupMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string                   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	ShortTitle  string                   `protobuf:"bytes,2,opt,name=short_title,json=shortTitle,proto3" json:"short_title,omitempty"`
	Instruction []string                 `protobuf:"bytes,3,rep,name=instruction,proto3" json:"instruction,omitempty"`
	ImageUrl    string                   `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Products    []*ProductInGroupMessage `protobuf:"bytes,5,rep,name=products,proto3" json:"products,omitempty"`
	StartTime   string                   `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	FinishTime  string                   `protobuf:"bytes,7,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
}

func (x *ProductGroupMessage) Reset() {
	*x = ProductGroupMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductGroupMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductGroupMessage) ProtoMessage() {}

func (x *ProductGroupMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_productGroup_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductGroupMessage.ProtoReflect.Descriptor instead.
func (*ProductGroupMessage) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_productGroup_proto_rawDescGZIP(), []int{9}
}

func (x *ProductGroupMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductGroupMessage) GetShortTitle() string {
	if x != nil {
		return x.ShortTitle
	}
	return ""
}

func (x *ProductGroupMessage) GetInstruction() []string {
	if x != nil {
		return x.Instruction
	}
	return nil
}

func (x *ProductGroupMessage) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *ProductGroupMessage) GetProducts() []*ProductInGroupMessage {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ProductGroupMessage) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ProductGroupMessage) GetFinishTime() string {
	if x != nil {
		return x.FinishTime
	}
	return ""
}

var File_api_grpcServer_protos_productGroup_proto protoreflect.FileDescriptor

var file_api_grpcServer_protos_productGroup_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x1a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22,
	0x5f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x22, 0xd1, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x02, 0x70, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x02, 0x70, 0x67, 0x22, 0x1a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x4e, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x03,
	0x70, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x70, 0x67, 0x73, 0x22,
	0x5e, 0x0a, 0x13, 0x50, 0x75, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22,
	0x47, 0x0a, 0x14, 0x50, 0x75, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x02, 0x70, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x02, 0x70, 0x67, 0x22, 0x69, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x22, 0x8a, 0x02, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x3d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65,
	0x32, 0x84, 0x03, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x5a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x62, 0x75, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_grpcServer_protos_productGroup_proto_rawDescOnce sync.Once
	file_api_grpcServer_protos_productGroup_proto_rawDescData = file_api_grpcServer_protos_productGroup_proto_rawDesc
)

func file_api_grpcServer_protos_productGroup_proto_rawDescGZIP() []byte {
	file_api_grpcServer_protos_productGroup_proto_rawDescOnce.Do(func() {
		file_api_grpcServer_protos_productGroup_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpcServer_protos_productGroup_proto_rawDescData)
	})
	return file_api_grpcServer_protos_productGroup_proto_rawDescData
}

var file_api_grpcServer_protos_productGroup_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_grpcServer_protos_productGroup_proto_goTypes = []interface{}{
	(*GetProductGroupRequest)(nil),     // 0: grpcServer.GetProductGroupRequest
	(*GetProductGroupResponse)(nil),    // 1: grpcServer.GetProductGroupResponse
	(*CreateProductGroupRequest)(nil),  // 2: grpcServer.CreateProductGroupRequest
	(*CreateProductGroupResponse)(nil), // 3: grpcServer.CreateProductGroupResponse
	(*ListProductGroupsRequest)(nil),   // 4: grpcServer.ListProductGroupsRequest
	(*ListProductGroupsResponse)(nil),  // 5: grpcServer.ListProductGroupsResponse
	(*PushProductsRequest)(nil),        // 6: grpcServer.PushProductsRequest
	(*PushProductsResponse)(nil),       // 7: grpcServer.PushProductsResponse
	(*ProductInGroupMessage)(nil),      // 8: grpcServer.ProductInGroupMessage
	(*ProductGroupMessage)(nil),        // 9: grpcServer.ProductGroupMessage
	(*ProductMessage)(nil),             // 10: grpcServer.ProductMessage
}
var file_api_grpcServer_protos_productGroup_proto_depIdxs = []int32{
	9,  // 0: grpcServer.GetProductGroupResponse.product_group:type_name -> grpcServer.ProductGroupMessage
	9,  // 1: grpcServer.CreateProductGroupResponse.pg:type_name -> grpcServer.ProductGroupMessage
	9,  // 2: grpcServer.ListProductGroupsResponse.pgs:type_name -> grpcServer.ProductGroupMessage
	9,  // 3: grpcServer.PushProductsResponse.pg:type_name -> grpcServer.ProductGroupMessage
	10, // 4: grpcServer.ProductInGroupMessage.product:type_name -> grpcServer.ProductMessage
	8,  // 5: grpcServer.ProductGroupMessage.products:type_name -> grpcServer.ProductInGroupMessage
	0,  // 6: grpcServer.ProductGroup.GetProductGroup:input_type -> grpcServer.GetProductGroupRequest
	2,  // 7: grpcServer.ProductGroup.CreateProductGroup:input_type -> grpcServer.CreateProductGroupRequest
	4,  // 8: grpcServer.ProductGroup.ListProductGroups:input_type -> grpcServer.ListProductGroupsRequest
	6,  // 9: grpcServer.ProductGroup.PushProducts:input_type -> grpcServer.PushProductsRequest
	1,  // 10: grpcServer.ProductGroup.GetProductGroup:output_type -> grpcServer.GetProductGroupResponse
	3,  // 11: grpcServer.ProductGroup.CreateProductGroup:output_type -> grpcServer.CreateProductGroupResponse
	5,  // 12: grpcServer.ProductGroup.ListProductGroups:output_type -> grpcServer.ListProductGroupsResponse
	7,  // 13: grpcServer.ProductGroup.PushProducts:output_type -> grpcServer.PushProductsResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_grpcServer_protos_productGroup_proto_init() }
func file_api_grpcServer_protos_productGroup_proto_init() {
	if File_api_grpcServer_protos_productGroup_proto != nil {
		return
	}
	file_api_grpcServer_protos_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_grpcServer_protos_productGroup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductGroupRequest); i {
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
		file_api_grpcServer_protos_productGroup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductGroupResponse); i {
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
		file_api_grpcServer_protos_productGroup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductGroupRequest); i {
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
		file_api_grpcServer_protos_productGroup_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductGroupResponse); i {
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
		file_api_grpcServer_protos_productGroup_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductGroupsRequest); i {
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
		file_api_grpcServer_protos_productGroup_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductGroupsResponse); i {
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
		file_api_grpcServer_protos_productGroup_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushProductsRequest); i {
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
		file_api_grpcServer_protos_productGroup_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushProductsResponse); i {
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
		file_api_grpcServer_protos_productGroup_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductInGroupMessage); i {
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
		file_api_grpcServer_protos_productGroup_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductGroupMessage); i {
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
			RawDescriptor: file_api_grpcServer_protos_productGroup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpcServer_protos_productGroup_proto_goTypes,
		DependencyIndexes: file_api_grpcServer_protos_productGroup_proto_depIdxs,
		MessageInfos:      file_api_grpcServer_protos_productGroup_proto_msgTypes,
	}.Build()
	File_api_grpcServer_protos_productGroup_proto = out.File
	file_api_grpcServer_protos_productGroup_proto_rawDesc = nil
	file_api_grpcServer_protos_productGroup_proto_goTypes = nil
	file_api_grpcServer_protos_productGroup_proto_depIdxs = nil
}