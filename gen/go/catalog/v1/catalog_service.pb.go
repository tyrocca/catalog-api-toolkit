// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: catalog/v1/catalog_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_catalog_v1_catalog_service_proto protoreflect.FileDescriptor

var file_catalog_v1_catalog_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xff, 0x09, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x22, 0x28, 0xda, 0x41, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x3a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x0d, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x12, 0x67, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x25, 0xda,
	0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76,
	0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65,
	0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x6e, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x69, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0xda, 0x41, 0x00, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x69, 0x65, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x45, 0xda,
	0x41, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x07, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x32, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65,
	0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x70, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x25, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16,
	0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x69, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x8b, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x22,
	0x43, 0xda, 0x41, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x73, 0x12, 0x72, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x22, 0x30, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x85, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0xda, 0x41,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x69, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73,
	0x12, 0x98, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x12, 0x20, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x22, 0x50, 0xda, 0x41, 0x13, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x3a, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x32, 0x29, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x3d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x2f, 0x2a, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x7b, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x20, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x30, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x23, 0x2a, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x72, 0x6f, 0x63, 0x63, 0x61, 0x2f, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x6b,
	0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_catalog_v1_catalog_service_proto_goTypes = []interface{}{
	(*CreateCompanyRequest)(nil),  // 0: catalog.v1.CreateCompanyRequest
	(*GetCompanyRequest)(nil),     // 1: catalog.v1.GetCompanyRequest
	(*ListCompaniesRequest)(nil),  // 2: catalog.v1.ListCompaniesRequest
	(*UpdateCompanyRequest)(nil),  // 3: catalog.v1.UpdateCompanyRequest
	(*DeleteCompanyRequest)(nil),  // 4: catalog.v1.DeleteCompanyRequest
	(*CreateCatalogRequest)(nil),  // 5: catalog.v1.CreateCatalogRequest
	(*GetCatalogRequest)(nil),     // 6: catalog.v1.GetCatalogRequest
	(*ListCatalogsRequest)(nil),   // 7: catalog.v1.ListCatalogsRequest
	(*UpdateCatalogRequest)(nil),  // 8: catalog.v1.UpdateCatalogRequest
	(*DeleteCatalogRequest)(nil),  // 9: catalog.v1.DeleteCatalogRequest
	(*Company)(nil),               // 10: catalog.v1.Company
	(*ListCompaniesResponse)(nil), // 11: catalog.v1.ListCompaniesResponse
	(*emptypb.Empty)(nil),         // 12: google.protobuf.Empty
	(*Catalog)(nil),               // 13: catalog.v1.Catalog
	(*ListCatalogsResponse)(nil),  // 14: catalog.v1.ListCatalogsResponse
}
var file_catalog_v1_catalog_service_proto_depIdxs = []int32{
	0,  // 0: catalog.v1.CatalogService.CreateCompany:input_type -> catalog.v1.CreateCompanyRequest
	1,  // 1: catalog.v1.CatalogService.GetCompany:input_type -> catalog.v1.GetCompanyRequest
	2,  // 2: catalog.v1.CatalogService.ListCompanies:input_type -> catalog.v1.ListCompaniesRequest
	3,  // 3: catalog.v1.CatalogService.UpdateCompany:input_type -> catalog.v1.UpdateCompanyRequest
	4,  // 4: catalog.v1.CatalogService.DeleteCompany:input_type -> catalog.v1.DeleteCompanyRequest
	5,  // 5: catalog.v1.CatalogService.CreateCatalog:input_type -> catalog.v1.CreateCatalogRequest
	6,  // 6: catalog.v1.CatalogService.GetCatalog:input_type -> catalog.v1.GetCatalogRequest
	7,  // 7: catalog.v1.CatalogService.ListCatalogs:input_type -> catalog.v1.ListCatalogsRequest
	8,  // 8: catalog.v1.CatalogService.UpdateCatalog:input_type -> catalog.v1.UpdateCatalogRequest
	9,  // 9: catalog.v1.CatalogService.DeleteCatalog:input_type -> catalog.v1.DeleteCatalogRequest
	10, // 10: catalog.v1.CatalogService.CreateCompany:output_type -> catalog.v1.Company
	10, // 11: catalog.v1.CatalogService.GetCompany:output_type -> catalog.v1.Company
	11, // 12: catalog.v1.CatalogService.ListCompanies:output_type -> catalog.v1.ListCompaniesResponse
	10, // 13: catalog.v1.CatalogService.UpdateCompany:output_type -> catalog.v1.Company
	12, // 14: catalog.v1.CatalogService.DeleteCompany:output_type -> google.protobuf.Empty
	13, // 15: catalog.v1.CatalogService.CreateCatalog:output_type -> catalog.v1.Catalog
	13, // 16: catalog.v1.CatalogService.GetCatalog:output_type -> catalog.v1.Catalog
	14, // 17: catalog.v1.CatalogService.ListCatalogs:output_type -> catalog.v1.ListCatalogsResponse
	13, // 18: catalog.v1.CatalogService.UpdateCatalog:output_type -> catalog.v1.Catalog
	12, // 19: catalog.v1.CatalogService.DeleteCatalog:output_type -> google.protobuf.Empty
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_catalog_v1_catalog_service_proto_init() }
func file_catalog_v1_catalog_service_proto_init() {
	if File_catalog_v1_catalog_service_proto != nil {
		return
	}
	file_catalog_v1_company_proto_init()
	file_catalog_v1_catalog_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_catalog_v1_catalog_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catalog_v1_catalog_service_proto_goTypes,
		DependencyIndexes: file_catalog_v1_catalog_service_proto_depIdxs,
	}.Build()
	File_catalog_v1_catalog_service_proto = out.File
	file_catalog_v1_catalog_service_proto_rawDesc = nil
	file_catalog_v1_catalog_service_proto_goTypes = nil
	file_catalog_v1_catalog_service_proto_depIdxs = nil
}
