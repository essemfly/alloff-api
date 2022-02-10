// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/grpcServer/protos/brand.proto

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

// Brand Service
type ListBrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBrandRequest) Reset() {
	*x = ListBrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_brand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBrandRequest) ProtoMessage() {}

func (x *ListBrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_brand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBrandRequest.ProtoReflect.Descriptor instead.
func (*ListBrandRequest) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_brand_proto_rawDescGZIP(), []int{0}
}

type ListBrandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brands []*BrandMessage `protobuf:"bytes,1,rep,name=brands,proto3" json:"brands,omitempty"`
}

func (x *ListBrandResponse) Reset() {
	*x = ListBrandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_brand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBrandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBrandResponse) ProtoMessage() {}

func (x *ListBrandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_brand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBrandResponse.ProtoReflect.Descriptor instead.
func (*ListBrandResponse) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_brand_proto_rawDescGZIP(), []int{1}
}

func (x *ListBrandResponse) GetBrands() []*BrandMessage {
	if x != nil {
		return x.Brands
	}
	return nil
}

type EditBrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyname       string              `protobuf:"bytes,1,opt,name=keyname,proto3" json:"keyname,omitempty"`
	Korname       *string             `protobuf:"bytes,2,opt,name=korname,proto3,oneof" json:"korname,omitempty"`
	Engname       *string             `protobuf:"bytes,3,opt,name=engname,proto3,oneof" json:"engname,omitempty"`
	LogoImageUrl  *string             `protobuf:"bytes,4,opt,name=logo_image_url,json=logoImageUrl,proto3,oneof" json:"logo_image_url,omitempty"`
	Description   *string             `protobuf:"bytes,5,opt,name=description,proto3,oneof" json:"description,omitempty"`
	IsPopular     *bool               `protobuf:"varint,6,opt,name=is_popular,json=isPopular,proto3,oneof" json:"is_popular,omitempty"`
	IsOpen        *bool               `protobuf:"varint,7,opt,name=is_open,json=isOpen,proto3,oneof" json:"is_open,omitempty"`
	InMaintenance *bool               `protobuf:"varint,8,opt,name=in_maintenance,json=inMaintenance,proto3,oneof" json:"in_maintenance,omitempty"`
	SizeGuide     []*SizeGuideMessage `protobuf:"bytes,9,rep,name=size_guide,json=sizeGuide,proto3" json:"size_guide,omitempty"`
	BackImageUrl  *string             `protobuf:"bytes,10,opt,name=back_image_url,json=backImageUrl,proto3,oneof" json:"back_image_url,omitempty"`
}

func (x *EditBrandRequest) Reset() {
	*x = EditBrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_brand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditBrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditBrandRequest) ProtoMessage() {}

func (x *EditBrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_brand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditBrandRequest.ProtoReflect.Descriptor instead.
func (*EditBrandRequest) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_brand_proto_rawDescGZIP(), []int{2}
}

func (x *EditBrandRequest) GetKeyname() string {
	if x != nil {
		return x.Keyname
	}
	return ""
}

func (x *EditBrandRequest) GetKorname() string {
	if x != nil && x.Korname != nil {
		return *x.Korname
	}
	return ""
}

func (x *EditBrandRequest) GetEngname() string {
	if x != nil && x.Engname != nil {
		return *x.Engname
	}
	return ""
}

func (x *EditBrandRequest) GetLogoImageUrl() string {
	if x != nil && x.LogoImageUrl != nil {
		return *x.LogoImageUrl
	}
	return ""
}

func (x *EditBrandRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *EditBrandRequest) GetIsPopular() bool {
	if x != nil && x.IsPopular != nil {
		return *x.IsPopular
	}
	return false
}

func (x *EditBrandRequest) GetIsOpen() bool {
	if x != nil && x.IsOpen != nil {
		return *x.IsOpen
	}
	return false
}

func (x *EditBrandRequest) GetInMaintenance() bool {
	if x != nil && x.InMaintenance != nil {
		return *x.InMaintenance
	}
	return false
}

func (x *EditBrandRequest) GetSizeGuide() []*SizeGuideMessage {
	if x != nil {
		return x.SizeGuide
	}
	return nil
}

func (x *EditBrandRequest) GetBackImageUrl() string {
	if x != nil && x.BackImageUrl != nil {
		return *x.BackImageUrl
	}
	return ""
}

type EditBrandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brand *BrandMessage `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *EditBrandResponse) Reset() {
	*x = EditBrandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_brand_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditBrandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditBrandResponse) ProtoMessage() {}

func (x *EditBrandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_brand_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditBrandResponse.ProtoReflect.Descriptor instead.
func (*EditBrandResponse) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_brand_proto_rawDescGZIP(), []int{3}
}

func (x *EditBrandResponse) GetBrand() *BrandMessage {
	if x != nil {
		return x.Brand
	}
	return nil
}

type CreateBrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyname       string              `protobuf:"bytes,1,opt,name=keyname,proto3" json:"keyname,omitempty"`
	Korname       string              `protobuf:"bytes,2,opt,name=korname,proto3" json:"korname,omitempty"`
	Engname       string              `protobuf:"bytes,3,opt,name=engname,proto3" json:"engname,omitempty"`
	LogoImageUrl  string              `protobuf:"bytes,4,opt,name=logo_image_url,json=logoImageUrl,proto3" json:"logo_image_url,omitempty"`
	Description   string              `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	IsPopular     bool                `protobuf:"varint,6,opt,name=is_popular,json=isPopular,proto3" json:"is_popular,omitempty"`
	IsOpen        bool                `protobuf:"varint,7,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	InMaintenance bool                `protobuf:"varint,8,opt,name=in_maintenance,json=inMaintenance,proto3" json:"in_maintenance,omitempty"`
	SizeGuide     []*SizeGuideMessage `protobuf:"bytes,9,rep,name=size_guide,json=sizeGuide,proto3" json:"size_guide,omitempty"`
	BackImageUrl  string              `protobuf:"bytes,10,opt,name=back_image_url,json=backImageUrl,proto3" json:"back_image_url,omitempty"`
}

func (x *CreateBrandRequest) Reset() {
	*x = CreateBrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_brand_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBrandRequest) ProtoMessage() {}

func (x *CreateBrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_brand_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBrandRequest.ProtoReflect.Descriptor instead.
func (*CreateBrandRequest) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_brand_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBrandRequest) GetKeyname() string {
	if x != nil {
		return x.Keyname
	}
	return ""
}

func (x *CreateBrandRequest) GetKorname() string {
	if x != nil {
		return x.Korname
	}
	return ""
}

func (x *CreateBrandRequest) GetEngname() string {
	if x != nil {
		return x.Engname
	}
	return ""
}

func (x *CreateBrandRequest) GetLogoImageUrl() string {
	if x != nil {
		return x.LogoImageUrl
	}
	return ""
}

func (x *CreateBrandRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateBrandRequest) GetIsPopular() bool {
	if x != nil {
		return x.IsPopular
	}
	return false
}

func (x *CreateBrandRequest) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *CreateBrandRequest) GetInMaintenance() bool {
	if x != nil {
		return x.InMaintenance
	}
	return false
}

func (x *CreateBrandRequest) GetSizeGuide() []*SizeGuideMessage {
	if x != nil {
		return x.SizeGuide
	}
	return nil
}

func (x *CreateBrandRequest) GetBackImageUrl() string {
	if x != nil {
		return x.BackImageUrl
	}
	return ""
}

type CreateBrandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brand *BrandMessage `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *CreateBrandResponse) Reset() {
	*x = CreateBrandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_brand_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBrandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBrandResponse) ProtoMessage() {}

func (x *CreateBrandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_brand_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBrandResponse.ProtoReflect.Descriptor instead.
func (*CreateBrandResponse) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_brand_proto_rawDescGZIP(), []int{5}
}

func (x *CreateBrandResponse) GetBrand() *BrandMessage {
	if x != nil {
		return x.Brand
	}
	return nil
}

type BrandMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrandId       string              `protobuf:"bytes,1,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	Korname       string              `protobuf:"bytes,2,opt,name=korname,proto3" json:"korname,omitempty"`
	Keyname       string              `protobuf:"bytes,3,opt,name=keyname,proto3" json:"keyname,omitempty"`
	Engname       string              `protobuf:"bytes,4,opt,name=engname,proto3" json:"engname,omitempty"`
	LogoImageUrl  string              `protobuf:"bytes,5,opt,name=logo_image_url,json=logoImageUrl,proto3" json:"logo_image_url,omitempty"`
	Description   string              `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	IsPopular     bool                `protobuf:"varint,7,opt,name=is_popular,json=isPopular,proto3" json:"is_popular,omitempty"`
	IsOpen        bool                `protobuf:"varint,8,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	InMaintenance bool                `protobuf:"varint,9,opt,name=in_maintenance,json=inMaintenance,proto3" json:"in_maintenance,omitempty"`
	SizeGuide     []*SizeGuideMessage `protobuf:"bytes,10,rep,name=size_guide,json=sizeGuide,proto3" json:"size_guide,omitempty"`
	BackImageUrl  string              `protobuf:"bytes,11,opt,name=back_image_url,json=backImageUrl,proto3" json:"back_image_url,omitempty"`
}

func (x *BrandMessage) Reset() {
	*x = BrandMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_brand_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandMessage) ProtoMessage() {}

func (x *BrandMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_brand_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandMessage.ProtoReflect.Descriptor instead.
func (*BrandMessage) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_brand_proto_rawDescGZIP(), []int{6}
}

func (x *BrandMessage) GetBrandId() string {
	if x != nil {
		return x.BrandId
	}
	return ""
}

func (x *BrandMessage) GetKorname() string {
	if x != nil {
		return x.Korname
	}
	return ""
}

func (x *BrandMessage) GetKeyname() string {
	if x != nil {
		return x.Keyname
	}
	return ""
}

func (x *BrandMessage) GetEngname() string {
	if x != nil {
		return x.Engname
	}
	return ""
}

func (x *BrandMessage) GetLogoImageUrl() string {
	if x != nil {
		return x.LogoImageUrl
	}
	return ""
}

func (x *BrandMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BrandMessage) GetIsPopular() bool {
	if x != nil {
		return x.IsPopular
	}
	return false
}

func (x *BrandMessage) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *BrandMessage) GetInMaintenance() bool {
	if x != nil {
		return x.InMaintenance
	}
	return false
}

func (x *BrandMessage) GetSizeGuide() []*SizeGuideMessage {
	if x != nil {
		return x.SizeGuide
	}
	return nil
}

func (x *BrandMessage) GetBackImageUrl() string {
	if x != nil {
		return x.BackImageUrl
	}
	return ""
}

type SizeGuideMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label    string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	ImageUrl string `protobuf:"bytes,2,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *SizeGuideMessage) Reset() {
	*x = SizeGuideMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpcServer_protos_brand_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SizeGuideMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SizeGuideMessage) ProtoMessage() {}

func (x *SizeGuideMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpcServer_protos_brand_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SizeGuideMessage.ProtoReflect.Descriptor instead.
func (*SizeGuideMessage) Descriptor() ([]byte, []int) {
	return file_api_grpcServer_protos_brand_proto_rawDescGZIP(), []int{7}
}

func (x *SizeGuideMessage) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *SizeGuideMessage) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_api_grpcServer_protos_brand_proto protoreflect.FileDescriptor

var file_api_grpcServer_protos_brand_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22,
	0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x8e, 0x04, 0x0a, 0x10, 0x45,
	0x64, 0x69, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6b, 0x65, 0x79, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x6b, 0x6f, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6b, 0x6f,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x65, 0x6e, 0x67, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x67,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x6f, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x88,
	0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52,
	0x09, 0x69, 0x73, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a,
	0x07, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05,
	0x52, 0x06, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x69,
	0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x0d, 0x69, 0x6e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f,
	0x67, 0x75, 0x69, 0x64, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x47, 0x75, 0x69,
	0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x47,
	0x75, 0x69, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0c,
	0x62, 0x61, 0x63, 0x6b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x6b, 0x6f, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x65, 0x6e, 0x67, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6c, 0x6f, 0x67, 0x6f,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x69,
	0x73, 0x5f, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x69, 0x73,
	0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x69,
	0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x43, 0x0a, 0x11, 0x45,
	0x64, 0x69, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2e, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x22, 0xec, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x6f, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6b, 0x6f, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x67, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e,
	0x67, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c,
	0x6f, 0x67, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69,
	0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69,
	0x6e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a,
	0x73, 0x69, 0x7a, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x69,
	0x7a, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09,
	0x73, 0x69, 0x7a, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x63,
	0x6b, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22,
	0x45, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x81, 0x03, 0x0a, 0x0c, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x6f, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x6f, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6b, 0x65, 0x79, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x67, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x67, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x6f, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70,
	0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6f, 0x70,
	0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x6e, 0x4d, 0x61, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f,
	0x67, 0x75, 0x69, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x47, 0x75, 0x69,
	0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x47,
	0x75, 0x69, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61,
	0x63, 0x6b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x45, 0x0a, 0x10, 0x53, 0x69,
	0x7a, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x6c, 0x32, 0xeb, 0x01, 0x0a, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x48, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x45, 0x64, 0x69, 0x74, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x45, 0x64, 0x69, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x64,
	0x69, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x1e,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65,
	0x73, 0x73, 0x62, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x66, 0x66, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpcServer_protos_brand_proto_rawDescOnce sync.Once
	file_api_grpcServer_protos_brand_proto_rawDescData = file_api_grpcServer_protos_brand_proto_rawDesc
)

func file_api_grpcServer_protos_brand_proto_rawDescGZIP() []byte {
	file_api_grpcServer_protos_brand_proto_rawDescOnce.Do(func() {
		file_api_grpcServer_protos_brand_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpcServer_protos_brand_proto_rawDescData)
	})
	return file_api_grpcServer_protos_brand_proto_rawDescData
}

var file_api_grpcServer_protos_brand_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_grpcServer_protos_brand_proto_goTypes = []interface{}{
	(*ListBrandRequest)(nil),    // 0: grpcServer.ListBrandRequest
	(*ListBrandResponse)(nil),   // 1: grpcServer.ListBrandResponse
	(*EditBrandRequest)(nil),    // 2: grpcServer.EditBrandRequest
	(*EditBrandResponse)(nil),   // 3: grpcServer.EditBrandResponse
	(*CreateBrandRequest)(nil),  // 4: grpcServer.CreateBrandRequest
	(*CreateBrandResponse)(nil), // 5: grpcServer.CreateBrandResponse
	(*BrandMessage)(nil),        // 6: grpcServer.BrandMessage
	(*SizeGuideMessage)(nil),    // 7: grpcServer.SizeGuideMessage
}
var file_api_grpcServer_protos_brand_proto_depIdxs = []int32{
	6, // 0: grpcServer.ListBrandResponse.brands:type_name -> grpcServer.BrandMessage
	7, // 1: grpcServer.EditBrandRequest.size_guide:type_name -> grpcServer.SizeGuideMessage
	6, // 2: grpcServer.EditBrandResponse.brand:type_name -> grpcServer.BrandMessage
	7, // 3: grpcServer.CreateBrandRequest.size_guide:type_name -> grpcServer.SizeGuideMessage
	6, // 4: grpcServer.CreateBrandResponse.brand:type_name -> grpcServer.BrandMessage
	7, // 5: grpcServer.BrandMessage.size_guide:type_name -> grpcServer.SizeGuideMessage
	0, // 6: grpcServer.Brand.ListBrand:input_type -> grpcServer.ListBrandRequest
	2, // 7: grpcServer.Brand.EditBrand:input_type -> grpcServer.EditBrandRequest
	4, // 8: grpcServer.Brand.CreateBrand:input_type -> grpcServer.CreateBrandRequest
	1, // 9: grpcServer.Brand.ListBrand:output_type -> grpcServer.ListBrandResponse
	3, // 10: grpcServer.Brand.EditBrand:output_type -> grpcServer.EditBrandResponse
	5, // 11: grpcServer.Brand.CreateBrand:output_type -> grpcServer.CreateBrandResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_grpcServer_protos_brand_proto_init() }
func file_api_grpcServer_protos_brand_proto_init() {
	if File_api_grpcServer_protos_brand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpcServer_protos_brand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBrandRequest); i {
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
		file_api_grpcServer_protos_brand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBrandResponse); i {
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
		file_api_grpcServer_protos_brand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditBrandRequest); i {
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
		file_api_grpcServer_protos_brand_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditBrandResponse); i {
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
		file_api_grpcServer_protos_brand_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBrandRequest); i {
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
		file_api_grpcServer_protos_brand_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBrandResponse); i {
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
		file_api_grpcServer_protos_brand_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrandMessage); i {
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
		file_api_grpcServer_protos_brand_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SizeGuideMessage); i {
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
	file_api_grpcServer_protos_brand_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_grpcServer_protos_brand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpcServer_protos_brand_proto_goTypes,
		DependencyIndexes: file_api_grpcServer_protos_brand_proto_depIdxs,
		MessageInfos:      file_api_grpcServer_protos_brand_proto_msgTypes,
	}.Build()
	File_api_grpcServer_protos_brand_proto = out.File
	file_api_grpcServer_protos_brand_proto_rawDesc = nil
	file_api_grpcServer_protos_brand_proto_goTypes = nil
	file_api_grpcServer_protos_brand_proto_depIdxs = nil
}
