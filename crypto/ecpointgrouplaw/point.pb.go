// Copyright © 2020 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: github.com/getamis/alice/crypto/ecpointgrouplaw/point.proto

package ecpointgrouplaw

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

type EcPointMessage_Curve int32

const (
	EcPointMessage_P224 EcPointMessage_Curve = 0
	EcPointMessage_P256 EcPointMessage_Curve = 1
	EcPointMessage_P384 EcPointMessage_Curve = 2
	// Above curves are not implemented
	EcPointMessage_S256        EcPointMessage_Curve = 3
	EcPointMessage_EDWARD25519 EcPointMessage_Curve = 4
)

// Enum value maps for EcPointMessage_Curve.
var (
	EcPointMessage_Curve_name = map[int32]string{
		0: "P224",
		1: "P256",
		2: "P384",
		3: "S256",
		4: "EDWARD25519",
	}
	EcPointMessage_Curve_value = map[string]int32{
		"P224":        0,
		"P256":        1,
		"P384":        2,
		"S256":        3,
		"EDWARD25519": 4,
	}
)

func (x EcPointMessage_Curve) Enum() *EcPointMessage_Curve {
	p := new(EcPointMessage_Curve)
	*p = x
	return p
}

func (x EcPointMessage_Curve) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EcPointMessage_Curve) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_enumTypes[0].Descriptor()
}

func (EcPointMessage_Curve) Type() protoreflect.EnumType {
	return &file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_enumTypes[0]
}

func (x EcPointMessage_Curve) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EcPointMessage_Curve.Descriptor instead.
func (EcPointMessage_Curve) EnumDescriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_rawDescGZIP(), []int{0, 0}
}

type EcPointMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Curve EcPointMessage_Curve `protobuf:"varint,1,opt,name=curve,proto3,enum=ecpointgrouplaw.EcPointMessage_Curve" json:"curve,omitempty"`
	X     []byte               `protobuf:"bytes,2,opt,name=x,proto3" json:"x,omitempty"`
	Y     []byte               `protobuf:"bytes,3,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *EcPointMessage) Reset() {
	*x = EcPointMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcPointMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcPointMessage) ProtoMessage() {}

func (x *EcPointMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcPointMessage.ProtoReflect.Descriptor instead.
func (*EcPointMessage) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_rawDescGZIP(), []int{0}
}

func (x *EcPointMessage) GetCurve() EcPointMessage_Curve {
	if x != nil {
		return x.Curve
	}
	return EcPointMessage_P224
}

func (x *EcPointMessage) GetX() []byte {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *EcPointMessage) GetY() []byte {
	if x != nil {
		return x.Y
	}
	return nil
}

var File_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto protoreflect.FileDescriptor

var file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74,
	0x61, 0x6d, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x65, 0x63, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x6c, 0x61,
	0x77, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65,
	0x63, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x6c, 0x61, 0x77, 0x22, 0xab,
	0x01, 0x0a, 0x0e, 0x45, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x3b, 0x0a, 0x05, 0x63, 0x75, 0x72, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x25, 0x2e, 0x65, 0x63, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x6c,
	0x61, 0x77, 0x2e, 0x45, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x43, 0x75, 0x72, 0x76, 0x65, 0x52, 0x05, 0x63, 0x75, 0x72, 0x76, 0x65, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x79, 0x22, 0x40, 0x0a, 0x05, 0x43, 0x75,
	0x72, 0x76, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x32, 0x32, 0x34, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x50, 0x32, 0x35, 0x36, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x33, 0x38, 0x34, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x32, 0x35, 0x36, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x45,
	0x44, 0x57, 0x41, 0x52, 0x44, 0x32, 0x35, 0x35, 0x31, 0x39, 0x10, 0x04, 0x42, 0x31, 0x5a, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x6d,
	0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f,
	0x65, 0x63, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x6c, 0x61, 0x77, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_rawDescOnce sync.Once
	file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_rawDescData = file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_rawDesc
)

func file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_rawDescGZIP() []byte {
	file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_rawDescOnce.Do(func() {
		file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_rawDescData)
	})
	return file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_rawDescData
}

var file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_goTypes = []interface{}{
	(EcPointMessage_Curve)(0), // 0: ecpointgrouplaw.EcPointMessage.Curve
	(*EcPointMessage)(nil),    // 1: ecpointgrouplaw.EcPointMessage
}
var file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_depIdxs = []int32{
	0, // 0: ecpointgrouplaw.EcPointMessage.curve:type_name -> ecpointgrouplaw.EcPointMessage.Curve
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_init() }
func file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_init() {
	if File_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcPointMessage); i {
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
			RawDescriptor: file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_goTypes,
		DependencyIndexes: file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_depIdxs,
		EnumInfos:         file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_enumTypes,
		MessageInfos:      file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_msgTypes,
	}.Build()
	File_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto = out.File
	file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_rawDesc = nil
	file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_goTypes = nil
	file_github_com_getamis_alice_crypto_ecpointgrouplaw_point_proto_depIdxs = nil
}
