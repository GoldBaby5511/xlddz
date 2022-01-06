// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: list.proto

package list

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

type CMDID_List int32

const (
	CMDID_List_IDRoomRegisterReq CMDID_List = 1 //房间注册
	CMDID_List_IDRoomRegisterRsp CMDID_List = 2 //房间注册
	CMDID_List_IDUserActionReq   CMDID_List = 3 //用户动作
	CMDID_List_IDUserActionRsp   CMDID_List = 4 //用户动作
	CMDID_List_IDExitReq         CMDID_List = 5 //离开房间
	CMDID_List_IDExitRsp         CMDID_List = 6 //离开房间
	CMDID_List_IDLast            CMDID_List = 10
)

// Enum value maps for CMDID_List.
var (
	CMDID_List_name = map[int32]string{
		1:  "IDRoomRegisterReq",
		2:  "IDRoomRegisterRsp",
		3:  "IDUserActionReq",
		4:  "IDUserActionRsp",
		5:  "IDExitReq",
		6:  "IDExitRsp",
		10: "IDLast",
	}
	CMDID_List_value = map[string]int32{
		"IDRoomRegisterReq": 1,
		"IDRoomRegisterRsp": 2,
		"IDUserActionReq":   3,
		"IDUserActionRsp":   4,
		"IDExitReq":         5,
		"IDExitRsp":         6,
		"IDLast":            10,
	}
)

func (x CMDID_List) Enum() *CMDID_List {
	p := new(CMDID_List)
	*p = x
	return p
}

func (x CMDID_List) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMDID_List) Descriptor() protoreflect.EnumDescriptor {
	return file_list_proto_enumTypes[0].Descriptor()
}

func (CMDID_List) Type() protoreflect.EnumType {
	return &file_list_proto_enumTypes[0]
}

func (x CMDID_List) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CMDID_List) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CMDID_List(num)
	return nil
}

// Deprecated: Use CMDID_List.Descriptor instead.
func (CMDID_List) EnumDescriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{0}
}

type RoomRegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomInfo *types.RoomInfo `protobuf:"bytes,1,opt,name=room_info,json=roomInfo" json:"room_info,omitempty"`
}

func (x *RoomRegisterReq) Reset() {
	*x = RoomRegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomRegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomRegisterReq) ProtoMessage() {}

func (x *RoomRegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomRegisterReq.ProtoReflect.Descriptor instead.
func (*RoomRegisterReq) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{0}
}

func (x *RoomRegisterReq) GetRoomInfo() *types.RoomInfo {
	if x != nil {
		return x.RoomInfo
	}
	return nil
}

type RoomRegisterRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrInfo *types.ErrorInfo `protobuf:"bytes,99,opt,name=err_info,json=errInfo" json:"err_info,omitempty"`
}

func (x *RoomRegisterRsp) Reset() {
	*x = RoomRegisterRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomRegisterRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomRegisterRsp) ProtoMessage() {}

func (x *RoomRegisterRsp) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomRegisterRsp.ProtoReflect.Descriptor instead.
func (*RoomRegisterRsp) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{1}
}

func (x *RoomRegisterRsp) GetErrInfo() *types.ErrorInfo {
	if x != nil {
		return x.ErrInfo
	}
	return nil
}

var File_list_proto protoreflect.FileDescriptor

var file_list_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x73,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0f, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x73, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x41, 0x0a, 0x0f, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x72, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x73,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x07, 0x65, 0x72, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0x8e, 0x01, 0x0a, 0x0a, 0x43, 0x4d,
	0x44, 0x49, 0x44, 0x5f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x44, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x10, 0x01, 0x12,
	0x15, 0x0a, 0x11, 0x49, 0x44, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x73, 0x70, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x44, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x49,
	0x44, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x10, 0x04,
	0x12, 0x0d, 0x0a, 0x09, 0x49, 0x44, 0x45, 0x78, 0x69, 0x74, 0x52, 0x65, 0x71, 0x10, 0x05, 0x12,
	0x0d, 0x0a, 0x09, 0x49, 0x44, 0x45, 0x78, 0x69, 0x74, 0x52, 0x73, 0x70, 0x10, 0x06, 0x12, 0x0a,
	0x0a, 0x06, 0x49, 0x44, 0x4c, 0x61, 0x73, 0x74, 0x10, 0x0a, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x6c,
	0x69, 0x73, 0x74,
}

var (
	file_list_proto_rawDescOnce sync.Once
	file_list_proto_rawDescData = file_list_proto_rawDesc
)

func file_list_proto_rawDescGZIP() []byte {
	file_list_proto_rawDescOnce.Do(func() {
		file_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_list_proto_rawDescData)
	})
	return file_list_proto_rawDescData
}

var file_list_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_list_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_list_proto_goTypes = []interface{}{
	(CMDID_List)(0),         // 0: bs.list.CMDID_List
	(*RoomRegisterReq)(nil), // 1: bs.list.RoomRegisterReq
	(*RoomRegisterRsp)(nil), // 2: bs.list.RoomRegisterRsp
	(*types.RoomInfo)(nil),  // 3: bs.types.RoomInfo
	(*types.ErrorInfo)(nil), // 4: bs.types.ErrorInfo
}
var file_list_proto_depIdxs = []int32{
	3, // 0: bs.list.RoomRegisterReq.room_info:type_name -> bs.types.RoomInfo
	4, // 1: bs.list.RoomRegisterRsp.err_info:type_name -> bs.types.ErrorInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_list_proto_init() }
func file_list_proto_init() {
	if File_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomRegisterReq); i {
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
		file_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomRegisterRsp); i {
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
			RawDescriptor: file_list_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_list_proto_goTypes,
		DependencyIndexes: file_list_proto_depIdxs,
		EnumInfos:         file_list_proto_enumTypes,
		MessageInfos:      file_list_proto_msgTypes,
	}.Build()
	File_list_proto = out.File
	file_list_proto_rawDesc = nil
	file_list_proto_goTypes = nil
	file_list_proto_depIdxs = nil
}
