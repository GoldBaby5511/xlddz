// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.20.3
// source: property.proto

package property

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	types "mango/api/types"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CMDProperty int32

const (
	CMDProperty_IDQueryPropertyReq  CMDProperty = 1 //查询财富
	CMDProperty_IDQueryPropertyRsp  CMDProperty = 2 //查询财富
	CMDProperty_IDModifyPropertyReq CMDProperty = 3 //修改财富
	CMDProperty_IDModifyPropertyRsp CMDProperty = 4 //修改财富
)

// Enum value maps for CMDProperty.
var (
	CMDProperty_name = map[int32]string{
		1: "IDQueryPropertyReq",
		2: "IDQueryPropertyRsp",
		3: "IDModifyPropertyReq",
		4: "IDModifyPropertyRsp",
	}
	CMDProperty_value = map[string]int32{
		"IDQueryPropertyReq":  1,
		"IDQueryPropertyRsp":  2,
		"IDModifyPropertyReq": 3,
		"IDModifyPropertyRsp": 4,
	}
)

func (x CMDProperty) Enum() *CMDProperty {
	p := new(CMDProperty)
	*p = x
	return p
}

func (x CMDProperty) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMDProperty) Descriptor() protoreflect.EnumDescriptor {
	return file_property_proto_enumTypes[0].Descriptor()
}

func (CMDProperty) Type() protoreflect.EnumType {
	return &file_property_proto_enumTypes[0]
}

func (x CMDProperty) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CMDProperty) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CMDProperty(num)
	return nil
}

// Deprecated: Use CMDProperty.Descriptor instead.
func (CMDProperty) EnumDescriptor() ([]byte, []int) {
	return file_property_proto_rawDescGZIP(), []int{0}
}

type QueryPropertyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId *uint64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"` //用户ID
}

func (x *QueryPropertyReq) Reset() {
	*x = QueryPropertyReq{}
	mi := &file_property_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryPropertyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPropertyReq) ProtoMessage() {}

func (x *QueryPropertyReq) ProtoReflect() protoreflect.Message {
	mi := &file_property_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPropertyReq.ProtoReflect.Descriptor instead.
func (*QueryPropertyReq) Descriptor() ([]byte, []int) {
	return file_property_proto_rawDescGZIP(), []int{0}
}

func (x *QueryPropertyReq) GetUserId() uint64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

type QueryPropertyRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    *uint64           `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`         //用户ID
	UserProps []*types.PropItem `protobuf:"bytes,2,rep,name=user_props,json=userProps" json:"user_props,omitempty"` //用户道具
	ErrInfo   *types.ErrorInfo  `protobuf:"bytes,99,opt,name=err_info,json=errInfo" json:"err_info,omitempty"`
}

func (x *QueryPropertyRsp) Reset() {
	*x = QueryPropertyRsp{}
	mi := &file_property_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryPropertyRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPropertyRsp) ProtoMessage() {}

func (x *QueryPropertyRsp) ProtoReflect() protoreflect.Message {
	mi := &file_property_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPropertyRsp.ProtoReflect.Descriptor instead.
func (*QueryPropertyRsp) Descriptor() ([]byte, []int) {
	return file_property_proto_rawDescGZIP(), []int{1}
}

func (x *QueryPropertyRsp) GetUserId() uint64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *QueryPropertyRsp) GetUserProps() []*types.PropItem {
	if x != nil {
		return x.UserProps
	}
	return nil
}

func (x *QueryPropertyRsp) GetErrInfo() *types.ErrorInfo {
	if x != nil {
		return x.ErrInfo
	}
	return nil
}

type ModifyPropertyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    *uint64           `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`         //用户ID
	OpType    *int32            `protobuf:"varint,2,opt,name=op_type,json=opType" json:"op_type,omitempty"`         //操作类型
	UserProps []*types.PropItem `protobuf:"bytes,3,rep,name=user_props,json=userProps" json:"user_props,omitempty"` //用户道具
}

func (x *ModifyPropertyReq) Reset() {
	*x = ModifyPropertyReq{}
	mi := &file_property_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModifyPropertyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyPropertyReq) ProtoMessage() {}

func (x *ModifyPropertyReq) ProtoReflect() protoreflect.Message {
	mi := &file_property_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyPropertyReq.ProtoReflect.Descriptor instead.
func (*ModifyPropertyReq) Descriptor() ([]byte, []int) {
	return file_property_proto_rawDescGZIP(), []int{2}
}

func (x *ModifyPropertyReq) GetUserId() uint64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *ModifyPropertyReq) GetOpType() int32 {
	if x != nil && x.OpType != nil {
		return *x.OpType
	}
	return 0
}

func (x *ModifyPropertyReq) GetUserProps() []*types.PropItem {
	if x != nil {
		return x.UserProps
	}
	return nil
}

type ModifyPropertyRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    *uint64           `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`         //用户ID
	OpType    *int32            `protobuf:"varint,2,opt,name=op_type,json=opType" json:"op_type,omitempty"`         //操作类型
	UserProps []*types.PropItem `protobuf:"bytes,3,rep,name=user_props,json=userProps" json:"user_props,omitempty"` //用户道具
	ErrInfo   *types.ErrorInfo  `protobuf:"bytes,99,opt,name=err_info,json=errInfo" json:"err_info,omitempty"`
}

func (x *ModifyPropertyRsp) Reset() {
	*x = ModifyPropertyRsp{}
	mi := &file_property_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModifyPropertyRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyPropertyRsp) ProtoMessage() {}

func (x *ModifyPropertyRsp) ProtoReflect() protoreflect.Message {
	mi := &file_property_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyPropertyRsp.ProtoReflect.Descriptor instead.
func (*ModifyPropertyRsp) Descriptor() ([]byte, []int) {
	return file_property_proto_rawDescGZIP(), []int{3}
}

func (x *ModifyPropertyRsp) GetUserId() uint64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *ModifyPropertyRsp) GetOpType() int32 {
	if x != nil && x.OpType != nil {
		return *x.OpType
	}
	return 0
}

func (x *ModifyPropertyRsp) GetUserProps() []*types.PropItem {
	if x != nil {
		return x.UserProps
	}
	return nil
}

func (x *ModifyPropertyRsp) GetErrInfo() *types.ErrorInfo {
	if x != nil {
		return x.ErrInfo
	}
	return nil
}

var File_property_proto protoreflect.FileDescriptor

var file_property_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x1a, 0x0b, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x10, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x73, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x73, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x65, 0x72, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x78, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x31, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x70, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x62, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x2e, 0x0a,
	0x08, 0x65, 0x72, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x62, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x65, 0x72, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0x6f, 0x0a,
	0x0b, 0x43, 0x4d, 0x44, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x12,
	0x49, 0x44, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x44, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x73, 0x70, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13,
	0x49, 0x44, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x44, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x73, 0x70, 0x10, 0x04, 0x42, 0x0b,
	0x5a, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
}

var (
	file_property_proto_rawDescOnce sync.Once
	file_property_proto_rawDescData = file_property_proto_rawDesc
)

func file_property_proto_rawDescGZIP() []byte {
	file_property_proto_rawDescOnce.Do(func() {
		file_property_proto_rawDescData = protoimpl.X.CompressGZIP(file_property_proto_rawDescData)
	})
	return file_property_proto_rawDescData
}

var file_property_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_property_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_property_proto_goTypes = []any{
	(CMDProperty)(0),          // 0: bs.property.CMDProperty
	(*QueryPropertyReq)(nil),  // 1: bs.property.QueryPropertyReq
	(*QueryPropertyRsp)(nil),  // 2: bs.property.QueryPropertyRsp
	(*ModifyPropertyReq)(nil), // 3: bs.property.ModifyPropertyReq
	(*ModifyPropertyRsp)(nil), // 4: bs.property.ModifyPropertyRsp
	(*types.PropItem)(nil),    // 5: bs.types.PropItem
	(*types.ErrorInfo)(nil),   // 6: bs.types.ErrorInfo
}
var file_property_proto_depIdxs = []int32{
	5, // 0: bs.property.QueryPropertyRsp.user_props:type_name -> bs.types.PropItem
	6, // 1: bs.property.QueryPropertyRsp.err_info:type_name -> bs.types.ErrorInfo
	5, // 2: bs.property.ModifyPropertyReq.user_props:type_name -> bs.types.PropItem
	5, // 3: bs.property.ModifyPropertyRsp.user_props:type_name -> bs.types.PropItem
	6, // 4: bs.property.ModifyPropertyRsp.err_info:type_name -> bs.types.ErrorInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_property_proto_init() }
func file_property_proto_init() {
	if File_property_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_property_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_property_proto_goTypes,
		DependencyIndexes: file_property_proto_depIdxs,
		EnumInfos:         file_property_proto_enumTypes,
		MessageInfos:      file_property_proto_msgTypes,
	}.Build()
	File_property_proto = out.File
	file_property_proto_rawDesc = nil
	file_property_proto_goTypes = nil
	file_property_proto_depIdxs = nil
}
