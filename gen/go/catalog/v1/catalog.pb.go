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

type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uid         string                 `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	DisplayName string                 `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	CreateTime  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Company) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Company) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Company) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Company) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type CreateCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company *Company `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
}

func (x *CreateCompanyRequest) Reset() {
	*x = CreateCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCompanyRequest) ProtoMessage() {}

func (x *CreateCompanyRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateCompanyRequest.ProtoReflect.Descriptor instead.
func (*CreateCompanyRequest) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCompanyRequest) GetCompany() *Company {
	if x != nil {
		return x.Company
	}
	return nil
}

type GetCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetCompanyRequest) Reset() {
	*x = GetCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyRequest) ProtoMessage() {}

func (x *GetCompanyRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCompanyRequest.ProtoReflect.Descriptor instead.
func (*GetCompanyRequest) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *GetCompanyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListCompaniesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Filter    string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	OrderBy   string `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
}

func (x *ListCompaniesRequest) Reset() {
	*x = ListCompaniesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCompaniesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCompaniesRequest) ProtoMessage() {}

func (x *ListCompaniesRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListCompaniesRequest.ProtoReflect.Descriptor instead.
func (*ListCompaniesRequest) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *ListCompaniesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCompaniesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListCompaniesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListCompaniesRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

type ListCompaniesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Companies     []*Company `protobuf:"bytes,1,rep,name=companies,proto3" json:"companies,omitempty"`
	NextPageToken string     `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListCompaniesResponse) Reset() {
	*x = ListCompaniesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_v1_catalog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCompaniesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCompaniesResponse) ProtoMessage() {}

func (x *ListCompaniesResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListCompaniesResponse.ProtoReflect.Descriptor instead.
func (*ListCompaniesResponse) Descriptor() ([]byte, []int) {
	return file_catalog_v1_catalog_proto_rawDescGZIP(), []int{4}
}

func (x *ListCompaniesResponse) GetCompanies() []*Company {
	if x != nil {
		return x.Companies
	}
	return nil
}

func (x *ListCompaniesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_catalog_v1_catalog_proto protoreflect.FileDescriptor

var file_catalog_v1_catalog_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12,
	0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0,
	0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07,
	0x72, 0x05, 0x20, 0x01, 0x28, 0x80, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xe0, 0x41, 0x05, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xe0, 0x41, 0x05, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4a,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x22, 0x72, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xd9, 0x02, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22,
	0x28, 0xda, 0x41, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x3a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x12, 0x65, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x23, 0xda, 0x41, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x12, 0x6e, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65,
	0x73, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0xda, 0x41, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73,
	0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x79, 0x72, 0x6f, 0x63, 0x63, 0x61, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_catalog_v1_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_catalog_v1_catalog_proto_goTypes = []interface{}{
	(*Company)(nil),               // 0: catalog.v1.Company
	(*CreateCompanyRequest)(nil),  // 1: catalog.v1.CreateCompanyRequest
	(*GetCompanyRequest)(nil),     // 2: catalog.v1.GetCompanyRequest
	(*ListCompaniesRequest)(nil),  // 3: catalog.v1.ListCompaniesRequest
	(*ListCompaniesResponse)(nil), // 4: catalog.v1.ListCompaniesResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_catalog_v1_catalog_proto_depIdxs = []int32{
	5, // 0: catalog.v1.Company.create_time:type_name -> google.protobuf.Timestamp
	5, // 1: catalog.v1.Company.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: catalog.v1.CreateCompanyRequest.company:type_name -> catalog.v1.Company
	0, // 3: catalog.v1.ListCompaniesResponse.companies:type_name -> catalog.v1.Company
	1, // 4: catalog.v1.CatalogService.CreateCompany:input_type -> catalog.v1.CreateCompanyRequest
	2, // 5: catalog.v1.CatalogService.GetCompany:input_type -> catalog.v1.GetCompanyRequest
	3, // 6: catalog.v1.CatalogService.ListCompanies:input_type -> catalog.v1.ListCompaniesRequest
	0, // 7: catalog.v1.CatalogService.CreateCompany:output_type -> catalog.v1.Company
	0, // 8: catalog.v1.CatalogService.GetCompany:output_type -> catalog.v1.Company
	4, // 9: catalog.v1.CatalogService.ListCompanies:output_type -> catalog.v1.ListCompaniesResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_catalog_v1_catalog_proto_init() }
func file_catalog_v1_catalog_proto_init() {
	if File_catalog_v1_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catalog_v1_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
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
			switch v := v.(*CreateCompanyRequest); i {
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
			switch v := v.(*GetCompanyRequest); i {
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
			switch v := v.(*ListCompaniesRequest); i {
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
			switch v := v.(*ListCompaniesResponse); i {
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
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
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
