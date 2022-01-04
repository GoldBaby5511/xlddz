// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: types.proto

package types

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

type BaseUserInfo_UserType int32

const (
	BaseUserInfo_UNKNOW BaseUserInfo_UserType = 0  //未知
	BaseUserInfo_Normal BaseUserInfo_UserType = 1  //正常类型
	BaseUserInfo_Robot  BaseUserInfo_UserType = 10 //机器人
)

// Enum value maps for BaseUserInfo_UserType.
var (
	BaseUserInfo_UserType_name = map[int32]string{
		0:  "UNKNOW",
		1:  "Normal",
		10: "Robot",
	}
	BaseUserInfo_UserType_value = map[string]int32{
		"UNKNOW": 0,
		"Normal": 1,
		"Robot":  10,
	}
)

func (x BaseUserInfo_UserType) Enum() *BaseUserInfo_UserType {
	p := new(BaseUserInfo_UserType)
	*p = x
	return p
}

func (x BaseUserInfo_UserType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaseUserInfo_UserType) Descriptor() protoreflect.EnumDescriptor {
	return file_types_proto_enumTypes[0].Descriptor()
}

func (BaseUserInfo_UserType) Type() protoreflect.EnumType {
	return &file_types_proto_enumTypes[0]
}

func (x BaseUserInfo_UserType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *BaseUserInfo_UserType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = BaseUserInfo_UserType(num)
	return nil
}

// Deprecated: Use BaseUserInfo_UserType.Descriptor instead.
func (BaseUserInfo_UserType) EnumDescriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2, 0}
}

//错误信息
type ErrorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code *int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Info *string `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (x *ErrorInfo) Reset() {
	*x = ErrorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorInfo) ProtoMessage() {}

func (x *ErrorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorInfo.ProtoReflect.Descriptor instead.
func (*ErrorInfo) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorInfo) GetCode() int32 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

func (x *ErrorInfo) GetInfo() string {
	if x != nil && x.Info != nil {
		return *x.Info
	}
	return ""
}

//道具信息
type PropItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropId    *int64 `protobuf:"varint,1,opt,name=prop_id,json=propId" json:"prop_id,omitempty"`
	PropCount *int64 `protobuf:"varint,2,opt,name=prop_count,json=propCount" json:"prop_count,omitempty"`
}

func (x *PropItem) Reset() {
	*x = PropItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropItem) ProtoMessage() {}

func (x *PropItem) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropItem.ProtoReflect.Descriptor instead.
func (*PropItem) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *PropItem) GetPropId() int64 {
	if x != nil && x.PropId != nil {
		return *x.PropId
	}
	return 0
}

func (x *PropItem) GetPropCount() int64 {
	if x != nil && x.PropCount != nil {
		return *x.PropCount
	}
	return 0
}

//基础信息
type BaseUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       *uint64                `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`                                           //用户ID
	GameId       *uint64                `protobuf:"varint,2,opt,name=game_id,json=gameId" json:"game_id,omitempty"`                                           //数字ID
	Gender       *uint32                `protobuf:"varint,3,opt,name=gender" json:"gender,omitempty"`                                                         //性别
	FaceId       *uint32                `protobuf:"varint,4,opt,name=face_id,json=faceId" json:"face_id,omitempty"`                                           //头像id
	CustomFace   *string                `protobuf:"bytes,5,opt,name=custom_face,json=customFace" json:"custom_face,omitempty"`                                //自定义的图像地址
	NickName     *string                `protobuf:"bytes,6,opt,name=nick_name,json=nickName" json:"nick_name,omitempty"`                                      //昵称
	UserType     *BaseUserInfo_UserType `protobuf:"varint,7,opt,name=user_type,json=userType,enum=bs.types.BaseUserInfo_UserType" json:"user_type,omitempty"` //用户类别
	UserProps    []*PropItem            `protobuf:"bytes,8,rep,name=user_props,json=userProps" json:"user_props,omitempty"`                                   //用户道具
	MarketId     *uint32                `protobuf:"varint,9,opt,name=market_id,json=marketId" json:"market_id,omitempty"`                                     //登录主渠道
	SiteId       *uint32                `protobuf:"varint,10,opt,name=site_id,json=siteId" json:"site_id,omitempty"`                                          //登录子渠道
	RegMarketId  *uint32                `protobuf:"varint,11,opt,name=reg_market_id,json=regMarketId" json:"reg_market_id,omitempty"`                         //注册主渠道
	RegSiteId    *uint32                `protobuf:"varint,12,opt,name=reg_site_id,json=regSiteId" json:"reg_site_id,omitempty"`                               //注册子渠道
	RegisterData *string                `protobuf:"bytes,13,opt,name=register_data,json=registerData" json:"register_data,omitempty"`                         //注册时间
}

func (x *BaseUserInfo) Reset() {
	*x = BaseUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseUserInfo) ProtoMessage() {}

func (x *BaseUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseUserInfo.ProtoReflect.Descriptor instead.
func (*BaseUserInfo) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2}
}

func (x *BaseUserInfo) GetUserId() uint64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *BaseUserInfo) GetGameId() uint64 {
	if x != nil && x.GameId != nil {
		return *x.GameId
	}
	return 0
}

func (x *BaseUserInfo) GetGender() uint32 {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return 0
}

func (x *BaseUserInfo) GetFaceId() uint32 {
	if x != nil && x.FaceId != nil {
		return *x.FaceId
	}
	return 0
}

func (x *BaseUserInfo) GetCustomFace() string {
	if x != nil && x.CustomFace != nil {
		return *x.CustomFace
	}
	return ""
}

func (x *BaseUserInfo) GetNickName() string {
	if x != nil && x.NickName != nil {
		return *x.NickName
	}
	return ""
}

func (x *BaseUserInfo) GetUserType() BaseUserInfo_UserType {
	if x != nil && x.UserType != nil {
		return *x.UserType
	}
	return BaseUserInfo_UNKNOW
}

func (x *BaseUserInfo) GetUserProps() []*PropItem {
	if x != nil {
		return x.UserProps
	}
	return nil
}

func (x *BaseUserInfo) GetMarketId() uint32 {
	if x != nil && x.MarketId != nil {
		return *x.MarketId
	}
	return 0
}

func (x *BaseUserInfo) GetSiteId() uint32 {
	if x != nil && x.SiteId != nil {
		return *x.SiteId
	}
	return 0
}

func (x *BaseUserInfo) GetRegMarketId() uint32 {
	if x != nil && x.RegMarketId != nil {
		return *x.RegMarketId
	}
	return 0
}

func (x *BaseUserInfo) GetRegSiteId() uint32 {
	if x != nil && x.RegSiteId != nil {
		return *x.RegSiteId
	}
	return 0
}

func (x *BaseUserInfo) GetRegisterData() string {
	if x != nil && x.RegisterData != nil {
		return *x.RegisterData
	}
	return ""
}

//房间信息
type UserRoomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseInfo  *BaseUserInfo `protobuf:"bytes,1,opt,name=base_info,json=baseInfo" json:"base_info,omitempty"`     //基础信息
	TableId   *uint64       `protobuf:"varint,2,opt,name=table_id,json=tableId" json:"table_id,omitempty"`       //所有桌子
	SeatIndex *uint32       `protobuf:"varint,3,opt,name=seat_index,json=seatIndex" json:"seat_index,omitempty"` //所在位置
	UserState *uint32       `protobuf:"varint,4,opt,name=user_state,json=userState" json:"user_state,omitempty"` //用户状态
	LostCount *uint32       `protobuf:"varint,5,opt,name=lost_count,json=lostCount" json:"lost_count,omitempty"` //玩家总输局
	DrawCount *uint32       `protobuf:"varint,6,opt,name=draw_count,json=drawCount" json:"draw_count,omitempty"` //玩家总平局
	WinCount  *uint32       `protobuf:"varint,7,opt,name=win_count,json=winCount" json:"win_count,omitempty"`    //玩家总胜局
}

func (x *UserRoomInfo) Reset() {
	*x = UserRoomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRoomInfo) ProtoMessage() {}

func (x *UserRoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRoomInfo.ProtoReflect.Descriptor instead.
func (*UserRoomInfo) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{3}
}

func (x *UserRoomInfo) GetBaseInfo() *BaseUserInfo {
	if x != nil {
		return x.BaseInfo
	}
	return nil
}

func (x *UserRoomInfo) GetTableId() uint64 {
	if x != nil && x.TableId != nil {
		return *x.TableId
	}
	return 0
}

func (x *UserRoomInfo) GetSeatIndex() uint32 {
	if x != nil && x.SeatIndex != nil {
		return *x.SeatIndex
	}
	return 0
}

func (x *UserRoomInfo) GetUserState() uint32 {
	if x != nil && x.UserState != nil {
		return *x.UserState
	}
	return 0
}

func (x *UserRoomInfo) GetLostCount() uint32 {
	if x != nil && x.LostCount != nil {
		return *x.LostCount
	}
	return 0
}

func (x *UserRoomInfo) GetDrawCount() uint32 {
	if x != nil && x.DrawCount != nil {
		return *x.DrawCount
	}
	return 0
}

func (x *UserRoomInfo) GetWinCount() uint32 {
	if x != nil && x.WinCount != nil {
		return *x.WinCount
	}
	return 0
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62,
	0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x33, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x42, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x70, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0xee, 0x03, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x67, 0x61, 0x6d,
	0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x66, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66,
	0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x46, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x62, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x31, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x50, 0x72, 0x6f, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x70, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x65, 0x67,
	0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x72, 0x65, 0x67, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0b, 0x72, 0x65, 0x67, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x72, 0x65, 0x67, 0x53, 0x69, 0x74, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x2d, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f,
	0x72, 0x6d, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x10,
	0x0a, 0x22, 0xf7, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x33, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x62,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x65, 0x61, 0x74, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x77, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x77, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x08, 0x5a, 0x06, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_types_proto_goTypes = []interface{}{
	(BaseUserInfo_UserType)(0), // 0: bs.types.BaseUserInfo.UserType
	(*ErrorInfo)(nil),          // 1: bs.types.ErrorInfo
	(*PropItem)(nil),           // 2: bs.types.PropItem
	(*BaseUserInfo)(nil),       // 3: bs.types.BaseUserInfo
	(*UserRoomInfo)(nil),       // 4: bs.types.UserRoomInfo
}
var file_types_proto_depIdxs = []int32{
	0, // 0: bs.types.BaseUserInfo.user_type:type_name -> bs.types.BaseUserInfo.UserType
	2, // 1: bs.types.BaseUserInfo.user_props:type_name -> bs.types.PropItem
	3, // 2: bs.types.UserRoomInfo.base_info:type_name -> bs.types.BaseUserInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorInfo); i {
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
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropItem); i {
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
		file_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseUserInfo); i {
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
		file_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRoomInfo); i {
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
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		EnumInfos:         file_types_proto_enumTypes,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
