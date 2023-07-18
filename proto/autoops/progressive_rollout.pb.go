// Copyright 2023 The Bucketeer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: proto/autoops/progressive_rollout.proto

package autoops

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProgressiveRollout_Type int32

const (
	ProgressiveRollout_MANUAL_SCHEDULE   ProgressiveRollout_Type = 0
	ProgressiveRollout_TEMPLATE_SCHEDULE ProgressiveRollout_Type = 1
)

// Enum value maps for ProgressiveRollout_Type.
var (
	ProgressiveRollout_Type_name = map[int32]string{
		0: "MANUAL_SCHEDULE",
		1: "TEMPLATE_SCHEDULE",
	}
	ProgressiveRollout_Type_value = map[string]int32{
		"MANUAL_SCHEDULE":   0,
		"TEMPLATE_SCHEDULE": 1,
	}
)

func (x ProgressiveRollout_Type) Enum() *ProgressiveRollout_Type {
	p := new(ProgressiveRollout_Type)
	*p = x
	return p
}

func (x ProgressiveRollout_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProgressiveRollout_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_autoops_progressive_rollout_proto_enumTypes[0].Descriptor()
}

func (ProgressiveRollout_Type) Type() protoreflect.EnumType {
	return &file_proto_autoops_progressive_rollout_proto_enumTypes[0]
}

func (x ProgressiveRollout_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProgressiveRollout_Type.Descriptor instead.
func (ProgressiveRollout_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_autoops_progressive_rollout_proto_rawDescGZIP(), []int{0, 0}
}

type ProgressiveRollout_Status int32

const (
	ProgressiveRollout_WAITING  ProgressiveRollout_Status = 0
	ProgressiveRollout_RUNNING  ProgressiveRollout_Status = 1
	ProgressiveRollout_FINISHED ProgressiveRollout_Status = 2
)

// Enum value maps for ProgressiveRollout_Status.
var (
	ProgressiveRollout_Status_name = map[int32]string{
		0: "WAITING",
		1: "RUNNING",
		2: "FINISHED",
	}
	ProgressiveRollout_Status_value = map[string]int32{
		"WAITING":  0,
		"RUNNING":  1,
		"FINISHED": 2,
	}
)

func (x ProgressiveRollout_Status) Enum() *ProgressiveRollout_Status {
	p := new(ProgressiveRollout_Status)
	*p = x
	return p
}

func (x ProgressiveRollout_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProgressiveRollout_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_autoops_progressive_rollout_proto_enumTypes[1].Descriptor()
}

func (ProgressiveRollout_Status) Type() protoreflect.EnumType {
	return &file_proto_autoops_progressive_rollout_proto_enumTypes[1]
}

func (x ProgressiveRollout_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProgressiveRollout_Status.Descriptor instead.
func (ProgressiveRollout_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_autoops_progressive_rollout_proto_rawDescGZIP(), []int{0, 1}
}

type ProgressiveRollout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FeatureId string                    `protobuf:"bytes,2,opt,name=feature_id,json=featureId,proto3" json:"feature_id,omitempty"`
	Clause    *anypb.Any                `protobuf:"bytes,3,opt,name=clause,proto3" json:"clause,omitempty"`
	Status    ProgressiveRollout_Status `protobuf:"varint,4,opt,name=status,proto3,enum=bucketeer.autoops.ProgressiveRollout_Status" json:"status,omitempty"`
	CreatedAt int64                     `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64                     `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Type      ProgressiveRollout_Type   `protobuf:"varint,7,opt,name=type,proto3,enum=bucketeer.autoops.ProgressiveRollout_Type" json:"type,omitempty"`
}

func (x *ProgressiveRollout) Reset() {
	*x = ProgressiveRollout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_autoops_progressive_rollout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressiveRollout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressiveRollout) ProtoMessage() {}

func (x *ProgressiveRollout) ProtoReflect() protoreflect.Message {
	mi := &file_proto_autoops_progressive_rollout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressiveRollout.ProtoReflect.Descriptor instead.
func (*ProgressiveRollout) Descriptor() ([]byte, []int) {
	return file_proto_autoops_progressive_rollout_proto_rawDescGZIP(), []int{0}
}

func (x *ProgressiveRollout) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProgressiveRollout) GetFeatureId() string {
	if x != nil {
		return x.FeatureId
	}
	return ""
}

func (x *ProgressiveRollout) GetClause() *anypb.Any {
	if x != nil {
		return x.Clause
	}
	return nil
}

func (x *ProgressiveRollout) GetStatus() ProgressiveRollout_Status {
	if x != nil {
		return x.Status
	}
	return ProgressiveRollout_WAITING
}

func (x *ProgressiveRollout) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ProgressiveRollout) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *ProgressiveRollout) GetType() ProgressiveRollout_Type {
	if x != nil {
		return x.Type
	}
	return ProgressiveRollout_MANUAL_SCHEDULE
}

var File_proto_autoops_progressive_rollout_proto protoreflect.FileDescriptor

var file_proto_autoops_progressive_rollout_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6f, 0x70, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x6c,
	0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x65, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6f, 0x70, 0x73, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x03, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x06, 0x63, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6f, 0x70, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x6f,
	0x75, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x3e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6f, 0x70,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6c,
	0x6c, 0x6f, 0x75, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x32, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x41, 0x4e, 0x55, 0x41,
	0x4c, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11,
	0x54, 0x45, 0x4d, 0x50, 0x4c, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c,
	0x45, 0x10, 0x01, 0x22, 0x30, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53,
	0x48, 0x45, 0x44, 0x10, 0x02, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2d, 0x69, 0x6f,
	0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6f, 0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_autoops_progressive_rollout_proto_rawDescOnce sync.Once
	file_proto_autoops_progressive_rollout_proto_rawDescData = file_proto_autoops_progressive_rollout_proto_rawDesc
)

func file_proto_autoops_progressive_rollout_proto_rawDescGZIP() []byte {
	file_proto_autoops_progressive_rollout_proto_rawDescOnce.Do(func() {
		file_proto_autoops_progressive_rollout_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_autoops_progressive_rollout_proto_rawDescData)
	})
	return file_proto_autoops_progressive_rollout_proto_rawDescData
}

var file_proto_autoops_progressive_rollout_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_autoops_progressive_rollout_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_autoops_progressive_rollout_proto_goTypes = []interface{}{
	(ProgressiveRollout_Type)(0),   // 0: bucketeer.autoops.ProgressiveRollout.Type
	(ProgressiveRollout_Status)(0), // 1: bucketeer.autoops.ProgressiveRollout.Status
	(*ProgressiveRollout)(nil),     // 2: bucketeer.autoops.ProgressiveRollout
	(*anypb.Any)(nil),              // 3: google.protobuf.Any
}
var file_proto_autoops_progressive_rollout_proto_depIdxs = []int32{
	3, // 0: bucketeer.autoops.ProgressiveRollout.clause:type_name -> google.protobuf.Any
	1, // 1: bucketeer.autoops.ProgressiveRollout.status:type_name -> bucketeer.autoops.ProgressiveRollout.Status
	0, // 2: bucketeer.autoops.ProgressiveRollout.type:type_name -> bucketeer.autoops.ProgressiveRollout.Type
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_autoops_progressive_rollout_proto_init() }
func file_proto_autoops_progressive_rollout_proto_init() {
	if File_proto_autoops_progressive_rollout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_autoops_progressive_rollout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressiveRollout); i {
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
			RawDescriptor: file_proto_autoops_progressive_rollout_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_autoops_progressive_rollout_proto_goTypes,
		DependencyIndexes: file_proto_autoops_progressive_rollout_proto_depIdxs,
		EnumInfos:         file_proto_autoops_progressive_rollout_proto_enumTypes,
		MessageInfos:      file_proto_autoops_progressive_rollout_proto_msgTypes,
	}.Build()
	File_proto_autoops_progressive_rollout_proto = out.File
	file_proto_autoops_progressive_rollout_proto_rawDesc = nil
	file_proto_autoops_progressive_rollout_proto_goTypes = nil
	file_proto_autoops_progressive_rollout_proto_depIdxs = nil
}
