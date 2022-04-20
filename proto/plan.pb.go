// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: substrait/plan.proto

package proto

import (
	extensions "github.com/substrait-io/substrait-go/proto/extensions"
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

// Either a relation or root relation
type PlanRel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RelType:
	//	*PlanRel_Rel
	//	*PlanRel_Root
	RelType isPlanRel_RelType `protobuf_oneof:"rel_type"`
}

func (x *PlanRel) Reset() {
	*x = PlanRel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanRel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanRel) ProtoMessage() {}

func (x *PlanRel) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanRel.ProtoReflect.Descriptor instead.
func (*PlanRel) Descriptor() ([]byte, []int) {
	return file_substrait_plan_proto_rawDescGZIP(), []int{0}
}

func (m *PlanRel) GetRelType() isPlanRel_RelType {
	if m != nil {
		return m.RelType
	}
	return nil
}

func (x *PlanRel) GetRel() *Rel {
	if x, ok := x.GetRelType().(*PlanRel_Rel); ok {
		return x.Rel
	}
	return nil
}

func (x *PlanRel) GetRoot() *RelRoot {
	if x, ok := x.GetRelType().(*PlanRel_Root); ok {
		return x.Root
	}
	return nil
}

type isPlanRel_RelType interface {
	isPlanRel_RelType()
}

type PlanRel_Rel struct {
	// Any relation
	Rel *Rel `protobuf:"bytes,1,opt,name=rel,proto3,oneof"`
}

type PlanRel_Root struct {
	// The root of a relation tree
	Root *RelRoot `protobuf:"bytes,2,opt,name=root,proto3,oneof"`
}

func (*PlanRel_Rel) isPlanRel_RelType() {}

func (*PlanRel_Root) isPlanRel_RelType() {}

// Describe a set of operations to complete.
// For compactness sake, identifiers are normalized at the plan level.
type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// a list of yaml specifications this plan may depend on
	ExtensionUris []*extensions.SimpleExtensionURI `protobuf:"bytes,1,rep,name=extension_uris,json=extensionUris,proto3" json:"extension_uris,omitempty"`
	// a list of extensions this plan may depend on
	Extensions []*extensions.SimpleExtensionDeclaration `protobuf:"bytes,2,rep,name=extensions,proto3" json:"extensions,omitempty"`
	// one or more relation trees that are associated with this plan.
	Relations []*PlanRel `protobuf:"bytes,3,rep,name=relations,proto3" json:"relations,omitempty"`
	// additional extensions associated with this plan.
	AdvancedExtensions *extensions.AdvancedExtension `protobuf:"bytes,4,opt,name=advanced_extensions,json=advancedExtensions,proto3" json:"advanced_extensions,omitempty"`
	// A list of com.google.Any entities that this plan may use. Can be used to
	// warn if some embedded message types are unknown. Note that this list may
	// include message types that are ignorable (optimizations) or that are
	// unused. In many cases, a consumer may be able to work with a plan even if
	// one or more message types defined here are unknown.
	ExpectedTypeUrls []string `protobuf:"bytes,5,rep,name=expected_type_urls,json=expectedTypeUrls,proto3" json:"expected_type_urls,omitempty"`
}

func (x *Plan) Reset() {
	*x = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_plan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_plan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_substrait_plan_proto_rawDescGZIP(), []int{1}
}

func (x *Plan) GetExtensionUris() []*extensions.SimpleExtensionURI {
	if x != nil {
		return x.ExtensionUris
	}
	return nil
}

func (x *Plan) GetExtensions() []*extensions.SimpleExtensionDeclaration {
	if x != nil {
		return x.Extensions
	}
	return nil
}

func (x *Plan) GetRelations() []*PlanRel {
	if x != nil {
		return x.Relations
	}
	return nil
}

func (x *Plan) GetAdvancedExtensions() *extensions.AdvancedExtension {
	if x != nil {
		return x.AdvancedExtensions
	}
	return nil
}

func (x *Plan) GetExpectedTypeUrls() []string {
	if x != nil {
		return x.ExpectedTypeUrls
	}
	return nil
}

var File_substrait_plan_proto protoreflect.FileDescriptor

var file_substrait_plan_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x1a, 0x17, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2f, 0x61, 0x6c, 0x67,
	0x65, 0x62, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x73, 0x75, 0x62, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x63, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x6c, 0x12, 0x22, 0x0a, 0x03,
	0x72, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x03, 0x72, 0x65, 0x6c,
	0x12, 0x28, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x6c, 0x52, 0x6f,
	0x6f, 0x74, 0x48, 0x00, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xe3, 0x02, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x6e, 0x12,
	0x4f, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x69,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x52,
	0x49, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x73,
	0x12, 0x50, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x63, 0x6c, 0x61,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x30, 0x0a, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x6c, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x58, 0x0a, 0x13, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x61, 0x64, 0x76, 0x61,
	0x6e, 0x63, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c,
	0x0a, 0x12, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x65, 0x78, 0x70, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x42, 0x57, 0x0a, 0x12,
	0x69, 0x6f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x12, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_substrait_plan_proto_rawDescOnce sync.Once
	file_substrait_plan_proto_rawDescData = file_substrait_plan_proto_rawDesc
)

func file_substrait_plan_proto_rawDescGZIP() []byte {
	file_substrait_plan_proto_rawDescOnce.Do(func() {
		file_substrait_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_substrait_plan_proto_rawDescData)
	})
	return file_substrait_plan_proto_rawDescData
}

var file_substrait_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_substrait_plan_proto_goTypes = []interface{}{
	(*PlanRel)(nil),                       // 0: substrait.PlanRel
	(*Plan)(nil),                          // 1: substrait.Plan
	(*Rel)(nil),                           // 2: substrait.Rel
	(*RelRoot)(nil),                       // 3: substrait.RelRoot
	(*extensions.SimpleExtensionURI)(nil), // 4: substrait.extensions.SimpleExtensionURI
	(*extensions.SimpleExtensionDeclaration)(nil), // 5: substrait.extensions.SimpleExtensionDeclaration
	(*extensions.AdvancedExtension)(nil),          // 6: substrait.extensions.AdvancedExtension
}
var file_substrait_plan_proto_depIdxs = []int32{
	2, // 0: substrait.PlanRel.rel:type_name -> substrait.Rel
	3, // 1: substrait.PlanRel.root:type_name -> substrait.RelRoot
	4, // 2: substrait.Plan.extension_uris:type_name -> substrait.extensions.SimpleExtensionURI
	5, // 3: substrait.Plan.extensions:type_name -> substrait.extensions.SimpleExtensionDeclaration
	0, // 4: substrait.Plan.relations:type_name -> substrait.PlanRel
	6, // 5: substrait.Plan.advanced_extensions:type_name -> substrait.extensions.AdvancedExtension
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_substrait_plan_proto_init() }
func file_substrait_plan_proto_init() {
	if File_substrait_plan_proto != nil {
		return
	}
	file_substrait_algebra_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_substrait_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanRel); i {
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
		file_substrait_plan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plan); i {
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
	file_substrait_plan_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PlanRel_Rel)(nil),
		(*PlanRel_Root)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_substrait_plan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_substrait_plan_proto_goTypes,
		DependencyIndexes: file_substrait_plan_proto_depIdxs,
		MessageInfos:      file_substrait_plan_proto_msgTypes,
	}.Build()
	File_substrait_plan_proto = out.File
	file_substrait_plan_proto_rawDesc = nil
	file_substrait_plan_proto_goTypes = nil
	file_substrait_plan_proto_depIdxs = nil
}
