// Copyright 2020-2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: buf/alpha/audit/v1alpha1/convert.proto

package auditv1alpha1

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

type BufAlphaRegistryV1Alpha1ConvertFormat int32

const (
	BufAlphaRegistryV1Alpha1ConvertFormat_BUF_ALPHA_REGISTRY_V1_ALPHA1_CONVERT_FORMAT_UNSPECIFIED BufAlphaRegistryV1Alpha1ConvertFormat = 0
	BufAlphaRegistryV1Alpha1ConvertFormat_BUF_ALPHA_REGISTRY_V1_ALPHA1_CONVERT_FORMAT_BIN         BufAlphaRegistryV1Alpha1ConvertFormat = 1
	BufAlphaRegistryV1Alpha1ConvertFormat_BUF_ALPHA_REGISTRY_V1_ALPHA1_CONVERT_FORMAT_JSON        BufAlphaRegistryV1Alpha1ConvertFormat = 2
)

// Enum value maps for BufAlphaRegistryV1Alpha1ConvertFormat.
var (
	BufAlphaRegistryV1Alpha1ConvertFormat_name = map[int32]string{
		0: "BUF_ALPHA_REGISTRY_V1_ALPHA1_CONVERT_FORMAT_UNSPECIFIED",
		1: "BUF_ALPHA_REGISTRY_V1_ALPHA1_CONVERT_FORMAT_BIN",
		2: "BUF_ALPHA_REGISTRY_V1_ALPHA1_CONVERT_FORMAT_JSON",
	}
	BufAlphaRegistryV1Alpha1ConvertFormat_value = map[string]int32{
		"BUF_ALPHA_REGISTRY_V1_ALPHA1_CONVERT_FORMAT_UNSPECIFIED": 0,
		"BUF_ALPHA_REGISTRY_V1_ALPHA1_CONVERT_FORMAT_BIN":         1,
		"BUF_ALPHA_REGISTRY_V1_ALPHA1_CONVERT_FORMAT_JSON":        2,
	}
)

func (x BufAlphaRegistryV1Alpha1ConvertFormat) Enum() *BufAlphaRegistryV1Alpha1ConvertFormat {
	p := new(BufAlphaRegistryV1Alpha1ConvertFormat)
	*p = x
	return p
}

func (x BufAlphaRegistryV1Alpha1ConvertFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BufAlphaRegistryV1Alpha1ConvertFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_alpha_audit_v1alpha1_convert_proto_enumTypes[0].Descriptor()
}

func (BufAlphaRegistryV1Alpha1ConvertFormat) Type() protoreflect.EnumType {
	return &file_buf_alpha_audit_v1alpha1_convert_proto_enumTypes[0]
}

func (x BufAlphaRegistryV1Alpha1ConvertFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BufAlphaRegistryV1Alpha1ConvertFormat.Descriptor instead.
func (BufAlphaRegistryV1Alpha1ConvertFormat) EnumDescriptor() ([]byte, []int) {
	return file_buf_alpha_audit_v1alpha1_convert_proto_rawDescGZIP(), []int{0}
}

var File_buf_alpha_audit_v1alpha1_convert_proto protoreflect.FileDescriptor

var file_buf_alpha_audit_v1alpha1_convert_proto_rawDesc = []byte{
	0x0a, 0x26, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2a, 0xcf, 0x01, 0x0a, 0x25, 0x42, 0x75, 0x66, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x3b, 0x0a, 0x37,
	0x42, 0x55, 0x46, 0x5f, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54,
	0x52, 0x59, 0x5f, 0x56, 0x31, 0x5f, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x31, 0x5f, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x33, 0x0a, 0x2f, 0x42, 0x55, 0x46,
	0x5f, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x59, 0x5f,
	0x56, 0x31, 0x5f, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x31, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52,
	0x54, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x42, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x34,
	0x0a, 0x30, 0x42, 0x55, 0x46, 0x5f, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x5f, 0x52, 0x45, 0x47, 0x49,
	0x53, 0x54, 0x52, 0x59, 0x5f, 0x56, 0x31, 0x5f, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x31, 0x5f, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x4a, 0x53,
	0x4f, 0x4e, 0x10, 0x02, 0x42, 0x84, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x41, 0x41,
	0xaa, 0x02, 0x18, 0x42, 0x75, 0x66, 0x2e, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x18, 0x42, 0x75,
	0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x24, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70,
	0x68, 0x61, 0x5c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b,
	0x42, 0x75, 0x66, 0x3a, 0x3a, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x3a, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_buf_alpha_audit_v1alpha1_convert_proto_rawDescOnce sync.Once
	file_buf_alpha_audit_v1alpha1_convert_proto_rawDescData = file_buf_alpha_audit_v1alpha1_convert_proto_rawDesc
)

func file_buf_alpha_audit_v1alpha1_convert_proto_rawDescGZIP() []byte {
	file_buf_alpha_audit_v1alpha1_convert_proto_rawDescOnce.Do(func() {
		file_buf_alpha_audit_v1alpha1_convert_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_alpha_audit_v1alpha1_convert_proto_rawDescData)
	})
	return file_buf_alpha_audit_v1alpha1_convert_proto_rawDescData
}

var file_buf_alpha_audit_v1alpha1_convert_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_alpha_audit_v1alpha1_convert_proto_goTypes = []interface{}{
	(BufAlphaRegistryV1Alpha1ConvertFormat)(0), // 0: buf.alpha.audit.v1alpha1.BufAlphaRegistryV1Alpha1ConvertFormat
}
var file_buf_alpha_audit_v1alpha1_convert_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_buf_alpha_audit_v1alpha1_convert_proto_init() }
func file_buf_alpha_audit_v1alpha1_convert_proto_init() {
	if File_buf_alpha_audit_v1alpha1_convert_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_alpha_audit_v1alpha1_convert_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_alpha_audit_v1alpha1_convert_proto_goTypes,
		DependencyIndexes: file_buf_alpha_audit_v1alpha1_convert_proto_depIdxs,
		EnumInfos:         file_buf_alpha_audit_v1alpha1_convert_proto_enumTypes,
	}.Build()
	File_buf_alpha_audit_v1alpha1_convert_proto = out.File
	file_buf_alpha_audit_v1alpha1_convert_proto_rawDesc = nil
	file_buf_alpha_audit_v1alpha1_convert_proto_goTypes = nil
	file_buf_alpha_audit_v1alpha1_convert_proto_depIdxs = nil
}
