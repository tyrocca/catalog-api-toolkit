// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: catalog/v1/catalog.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Catalog Stuff
type Catalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uid         string                 `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	DisplayName string                 `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	CreateTime  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Catalog) Reset() {
	*x = Catalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Catalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Catalog) ProtoMessage() {}

func (x *Catalog) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Catalog.ProtoReflect.Descriptor instead.
func (*Catalog) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *Catalog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Catalog) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Catalog) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Catalog) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Catalog) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type CreateCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent  string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Catalog *Catalog `protobuf:"bytes,2,opt,name=catalog,proto3" json:"catalog,omitempty"`
}

func (x *CreateCatalogRequest) Reset() {
	*x = CreateCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCatalogRequest) ProtoMessage() {}

func (x *CreateCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCatalogRequest.ProtoReflect.Descriptor instead.
func (*CreateCatalogRequest) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCatalogRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateCatalogRequest) GetCatalog() *Catalog {
	if x != nil {
		return x.Catalog
	}
	return nil
}

type GetCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetCatalogRequest) Reset() {
	*x = GetCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogRequest) ProtoMessage() {}

func (x *GetCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogRequest.ProtoReflect.Descriptor instead.
func (*GetCatalogRequest) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *GetCatalogRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListCatalogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent    string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Filter    string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	OrderBy   string `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
}

func (x *ListCatalogsRequest) Reset() {
	*x = ListCatalogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCatalogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatalogsRequest) ProtoMessage() {}

func (x *ListCatalogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatalogsRequest.ProtoReflect.Descriptor instead.
func (*ListCatalogsRequest) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *ListCatalogsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListCatalogsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCatalogsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListCatalogsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListCatalogsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

type ListCatalogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Catalogs      []*Catalog `protobuf:"bytes,1,rep,name=catalogs,proto3" json:"catalogs,omitempty"`
	NextPageToken string     `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListCatalogsResponse) Reset() {
	*x = ListCatalogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_catalog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCatalogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatalogsResponse) ProtoMessage() {}

func (x *ListCatalogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatalogsResponse.ProtoReflect.Descriptor instead.
func (*ListCatalogsResponse) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{4}
}

func (x *ListCatalogsResponse) GetCatalogs() []*Catalog {
	if x != nil {
		return x.Catalogs
	}
	return nil
}

func (x *ListCatalogsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type UpdateCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Catalog    *Catalog               `protobuf:"bytes,1,opt,name=catalog,proto3" json:"catalog,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateCatalogRequest) Reset() {
	*x = UpdateCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_catalog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCatalogRequest) ProtoMessage() {}

func (x *UpdateCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCatalogRequest.ProtoReflect.Descriptor instead.
func (*UpdateCatalogRequest) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCatalogRequest) GetCatalog() *Catalog {
	if x != nil {
		return x.Catalog
	}
	return nil
}

func (x *UpdateCatalogRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Force bool   `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *DeleteCatalogRequest) Reset() {
	*x = DeleteCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_catalog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCatalogRequest) ProtoMessage() {}

func (x *DeleteCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_v1_catalog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCatalogRequest.ProtoReflect.Descriptor instead.
func (*DeleteCatalogRequest) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCatalogRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteCatalogRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

var File_catalog_v1_catalog_proto protoreflect.FileDescriptor

var file_catalog_v1_catalog_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5,
	0x02, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x20, 0x01, 0x28, 0x80, 0x02, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xe0, 0x41, 0x05, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x43,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x06, 0xe0, 0x41, 0x05, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x3a, 0x51, 0xea, 0x41, 0x4e, 0x0a, 0x24, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x26,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x7d, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x7d, 0x22, 0x90, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x44, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2c, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26, 0x12, 0x24, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x22, 0x55, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41,
	0x02, 0xfa, 0x41, 0x26, 0x0a, 0x24, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2d, 0x61, 0x70,
	0x69, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xca, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x26,
	0x12, 0x24, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x74, 0x6f,
	0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x22, 0x6f, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x87,
	0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x45, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x42,
	0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x79,
	0x72, 0x6f, 0x63, 0x63, 0x61, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2d, 0x61, 0x70,
	0x69, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_catalog_v1_catalog_proto_rawDescOnce sync.Once
	file_catalog_v1_catalog_proto_rawDescData = file_catalog_v1_catalog_proto_rawDesc
)

func file_catalog_v1_catalog_proto_rawDescGZIP() []byte {
	file_catalog_v1_catalog_proto_rawDescOnce.Do(func() {
		file_catalog_v1_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_catalog_v1_catalog_proto_rawDescData)
	})
	return file_catalog_v1_catalog_proto_rawDescData
}

var file_catalog_v1_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_catalog_v1_catalog_proto_goTypes = []interface{}{
	(*Catalog)(nil),               // 0: catalog.v1.Catalog
	(*CreateCatalogRequest)(nil),  // 1: catalog.v1.CreateCatalogRequest
	(*GetCatalogRequest)(nil),     // 2: catalog.v1.GetCatalogRequest
	(*ListCatalogsRequest)(nil),   // 3: catalog.v1.ListCatalogsRequest
	(*ListCatalogsResponse)(nil),  // 4: catalog.v1.ListCatalogsResponse
	(*UpdateCatalogRequest)(nil),  // 5: catalog.v1.UpdateCatalogRequest
	(*DeleteCatalogRequest)(nil),  // 6: catalog.v1.DeleteCatalogRequest
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 8: google.protobuf.FieldMask
}
var file_catalog_v1_catalog_proto_depIdxs = []int32{
	7, // 0: catalog.v1.Catalog.create_time:type_name -> google.protobuf.Timestamp
	7, // 1: catalog.v1.Catalog.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: catalog.v1.CreateCatalogRequest.catalog:type_name -> catalog.v1.Catalog
	0, // 3: catalog.v1.ListCatalogsResponse.catalogs:type_name -> catalog.v1.Catalog
	0, // 4: catalog.v1.UpdateCatalogRequest.catalog:type_name -> catalog.v1.Catalog
	8, // 5: catalog.v1.UpdateCatalogRequest.update_mask:type_name -> google.protobuf.FieldMask
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_catalog_v1_catalog_proto_init() }
func file_catalog_v1_catalog_proto_init() {
	if File_catalog_v1_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catalog_v1_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Catalog); i {
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
		file_catalog_v1_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCatalogRequest); i {
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
		file_catalog_v1_catalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogRequest); i {
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
		file_catalog_v1_catalog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCatalogsRequest); i {
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
		file_catalog_v1_catalog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCatalogsResponse); i {
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
		file_catalog_v1_catalog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCatalogRequest); i {
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
		file_catalog_v1_catalog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCatalogRequest); i {
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
			RawDescriptor: file_catalog_v1_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_catalog_v1_catalog_proto_goTypes,
		DependencyIndexes: file_catalog_v1_catalog_proto_depIdxs,
		MessageInfos:      file_catalog_v1_catalog_proto_msgTypes,
	}.Build()
	File_catalog_v1_catalog_proto = out.File
	file_catalog_v1_catalog_proto_rawDesc = nil
	file_catalog_v1_catalog_proto_goTypes = nil
	file_catalog_v1_catalog_proto_depIdxs = nil
}
