// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.20.3
// source: lobby.proto

package lobby

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

type CMDLobby int32

const (
	CMDLobby_IDNone             CMDLobby = 0
	CMDLobby_IDLoginReq         CMDLobby = 1 //登录请求
	CMDLobby_IDLoginRsp         CMDLobby = 2 //登录回复
	CMDLobby_IDLogoutReq        CMDLobby = 3 //注销登录
	CMDLobby_IDLogoutRsp        CMDLobby = 4 //注销登录
	CMDLobby_IDLogicRegReq      CMDLobby = 5 //逻辑注册
	CMDLobby_IDLogicRegRsp      CMDLobby = 6 //逻辑注册
	CMDLobby_IDQueryUserInfoReq CMDLobby = 7 //查询用户
	CMDLobby_IDQueryUserInfoRsp CMDLobby = 8 //查询用户
)

// Enum value maps for CMDLobby.
var (
	CMDLobby_name = map[int32]string{
		0: "IDNone",
		1: "IDLoginReq",
		2: "IDLoginRsp",
		3: "IDLogoutReq",
		4: "IDLogoutRsp",
		5: "IDLogicRegReq",
		6: "IDLogicRegRsp",
		7: "IDQueryUserInfoReq",
		8: "IDQueryUserInfoRsp",
	}
	CMDLobby_value = map[string]int32{
		"IDNone":             0,
		"IDLoginReq":         1,
		"IDLoginRsp":         2,
		"IDLogoutReq":        3,
		"IDLogoutRsp":        4,
		"IDLogicRegReq":      5,
		"IDLogicRegRsp":      6,
		"IDQueryUserInfoReq": 7,
		"IDQueryUserInfoRsp": 8,
	}
)

func (x CMDLobby) Enum() *CMDLobby {
	p := new(CMDLobby)
	*p = x
	return p
}

func (x CMDLobby) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMDLobby) Descriptor() protoreflect.EnumDescriptor {
	return file_lobby_proto_enumTypes[0].Descriptor()
}

func (CMDLobby) Type() protoreflect.EnumType {
	return &file_lobby_proto_enumTypes[0]
}

func (x CMDLobby) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CMDLobby.Descriptor instead.
func (CMDLobby) EnumDescriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{0}
}

type LoginReq_LoginType int32

const (
	LoginReq_acc     LoginReq_LoginType = 0 //账号
	LoginReq_token   LoginReq_LoginType = 1 //
	LoginReq_unionId LoginReq_LoginType = 2 //唯一标识(游客登录等)
)

// Enum value maps for LoginReq_LoginType.
var (
	LoginReq_LoginType_name = map[int32]string{
		0: "acc",
		1: "token",
		2: "unionId",
	}
	LoginReq_LoginType_value = map[string]int32{
		"acc":     0,
		"token":   1,
		"unionId": 2,
	}
)

func (x LoginReq_LoginType) Enum() *LoginReq_LoginType {
	p := new(LoginReq_LoginType)
	*p = x
	return p
}

func (x LoginReq_LoginType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginReq_LoginType) Descriptor() protoreflect.EnumDescriptor {
	return file_lobby_proto_enumTypes[1].Descriptor()
}

func (LoginReq_LoginType) Type() protoreflect.EnumType {
	return &file_lobby_proto_enumTypes[1]
}

func (x LoginReq_LoginType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginReq_LoginType.Descriptor instead.
func (LoginReq_LoginType) EnumDescriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{0, 0}
}

type LoginRsp_Result int32

const (
	LoginRsp_SUCCESS        LoginRsp_Result = 0    //成功
	LoginRsp_NOTEXIST       LoginRsp_Result = 1    //账号不存在
	LoginRsp_FROZEN         LoginRsp_Result = 2    //账号被冻结
	LoginRsp_FALSEPW        LoginRsp_Result = 3    //密码错误
	LoginRsp_NETERROR       LoginRsp_Result = 4    //网络异常
	LoginRsp_APPISBUSY      LoginRsp_Result = 5    //服务器忙，人数爆满
	LoginRsp_GUESTFORBID    LoginRsp_Result = 6    //禁止游客登录
	LoginRsp_CONNECTERROR   LoginRsp_Result = 7    //连接异常
	LoginRsp_VERSIONOLD     LoginRsp_Result = 8    //版本过低
	LoginRsp_NOMOREGUEST    LoginRsp_Result = 9    //游客分配失败
	LoginRsp_FREQUENTLY     LoginRsp_Result = 10   //所在ip登录过多
	LoginRsp_APPINITING     LoginRsp_Result = 11   //系统初始化，请稍后再试
	LoginRsp_SERVERERROR    LoginRsp_Result = 255  //服务端出错
	LoginRsp_UNKOWN         LoginRsp_Result = 1000 //未知错误
	LoginRsp_TOKEN_FAILED   LoginRsp_Result = 1001 //Token出错
	LoginRsp_TOKEN_EXPIRED  LoginRsp_Result = 1002 //token过期了
	LoginRsp_TOKEN_NOTMATCH LoginRsp_Result = 1003 //token与appid不匹配
)

// Enum value maps for LoginRsp_Result.
var (
	LoginRsp_Result_name = map[int32]string{
		0:    "SUCCESS",
		1:    "NOTEXIST",
		2:    "FROZEN",
		3:    "FALSEPW",
		4:    "NETERROR",
		5:    "APPISBUSY",
		6:    "GUESTFORBID",
		7:    "CONNECTERROR",
		8:    "VERSIONOLD",
		9:    "NOMOREGUEST",
		10:   "FREQUENTLY",
		11:   "APPINITING",
		255:  "SERVERERROR",
		1000: "UNKOWN",
		1001: "TOKEN_FAILED",
		1002: "TOKEN_EXPIRED",
		1003: "TOKEN_NOTMATCH",
	}
	LoginRsp_Result_value = map[string]int32{
		"SUCCESS":        0,
		"NOTEXIST":       1,
		"FROZEN":         2,
		"FALSEPW":        3,
		"NETERROR":       4,
		"APPISBUSY":      5,
		"GUESTFORBID":    6,
		"CONNECTERROR":   7,
		"VERSIONOLD":     8,
		"NOMOREGUEST":    9,
		"FREQUENTLY":     10,
		"APPINITING":     11,
		"SERVERERROR":    255,
		"UNKOWN":         1000,
		"TOKEN_FAILED":   1001,
		"TOKEN_EXPIRED":  1002,
		"TOKEN_NOTMATCH": 1003,
	}
)

func (x LoginRsp_Result) Enum() *LoginRsp_Result {
	p := new(LoginRsp_Result)
	*p = x
	return p
}

func (x LoginRsp_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginRsp_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_lobby_proto_enumTypes[2].Descriptor()
}

func (LoginRsp_Result) Type() protoreflect.EnumType {
	return &file_lobby_proto_enumTypes[2]
}

func (x LoginRsp_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginRsp_Result.Descriptor instead.
func (LoginRsp_Result) EnumDescriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{1, 0}
}

type LogoutRsp_LogoutReason int32

const (
	LogoutRsp_Normal       LogoutRsp_LogoutReason = 0
	LogoutRsp_AnotherLogin LogoutRsp_LogoutReason = 1 //被顶号
)

// Enum value maps for LogoutRsp_LogoutReason.
var (
	LogoutRsp_LogoutReason_name = map[int32]string{
		0: "Normal",
		1: "AnotherLogin",
	}
	LogoutRsp_LogoutReason_value = map[string]int32{
		"Normal":       0,
		"AnotherLogin": 1,
	}
)

func (x LogoutRsp_LogoutReason) Enum() *LogoutRsp_LogoutReason {
	p := new(LogoutRsp_LogoutReason)
	*p = x
	return p
}

func (x LogoutRsp_LogoutReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogoutRsp_LogoutReason) Descriptor() protoreflect.EnumDescriptor {
	return file_lobby_proto_enumTypes[3].Descriptor()
}

func (LogoutRsp_LogoutReason) Type() protoreflect.EnumType {
	return &file_lobby_proto_enumTypes[3]
}

func (x LogoutRsp_LogoutReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogoutRsp_LogoutReason.Descriptor instead.
func (LogoutRsp_LogoutReason) EnumDescriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{3, 0}
}

// 登录请求
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameKind      uint32             `protobuf:"varint,1,opt,name=game_kind,json=gameKind,proto3" json:"game_kind,omitempty"`                                     //游戏种类
	LoginType     LoginReq_LoginType `protobuf:"varint,2,opt,name=login_type,json=loginType,proto3,enum=bs.lobby.LoginReq_LoginType" json:"login_type,omitempty"` //登录类型
	Account       string             `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`                                                        //用户账号(根据 LoginAction填不通内容)
	Password      string             `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`                                                      //用户密码
	Version       string             `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`                                                        //客户端版本号
	Ip            string             `protobuf:"bytes,6,opt,name=ip,proto3" json:"ip,omitempty"`                                                                  //客户端IP
	SystemVersion string             `protobuf:"bytes,7,opt,name=system_version,json=systemVersion,proto3" json:"system_version,omitempty"`                       //操作系统版本号
	ChannelId     uint32             `protobuf:"varint,8,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`                                  //主渠道id
	SiteId        uint32             `protobuf:"varint,9,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`                                           //子渠道id
	DeviceId      string             `protobuf:"bytes,10,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`                                     //设备唯一码
	UserType      int32              `protobuf:"varint,11,opt,name=user_type,json=userType,proto3" json:"user_type,omitempty"`                                    //用户类型(客户端禁止使用)
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	mi := &file_lobby_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetGameKind() uint32 {
	if x != nil {
		return x.GameKind
	}
	return 0
}

func (x *LoginReq) GetLoginType() LoginReq_LoginType {
	if x != nil {
		return x.LoginType
	}
	return LoginReq_acc
}

func (x *LoginReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginReq) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *LoginReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *LoginReq) GetSystemVersion() string {
	if x != nil {
		return x.SystemVersion
	}
	return ""
}

func (x *LoginReq) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *LoginReq) GetSiteId() uint32 {
	if x != nil {
		return x.SiteId
	}
	return 0
}

func (x *LoginReq) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *LoginReq) GetUserType() int32 {
	if x != nil {
		return x.UserType
	}
	return 0
}

// 登录回复
type LoginRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   LoginRsp_Result     `protobuf:"varint,1,opt,name=result,proto3,enum=bs.lobby.LoginRsp_Result" json:"result,omitempty"` //登录结果
	UserInfo *types.BaseUserInfo `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`            //基本信息
	ErrInfo  *types.ErrorInfo    `protobuf:"bytes,99,opt,name=err_info,json=errInfo,proto3" json:"err_info,omitempty"`
}

func (x *LoginRsp) Reset() {
	*x = LoginRsp{}
	mi := &file_lobby_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRsp) ProtoMessage() {}

func (x *LoginRsp) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRsp.ProtoReflect.Descriptor instead.
func (*LoginRsp) Descriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRsp) GetResult() LoginRsp_Result {
	if x != nil {
		return x.Result
	}
	return LoginRsp_SUCCESS
}

func (x *LoginRsp) GetUserInfo() *types.BaseUserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *LoginRsp) GetErrInfo() *types.ErrorInfo {
	if x != nil {
		return x.ErrInfo
	}
	return nil
}

// 注销登录
type LogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GateConnId uint64 `protobuf:"varint,2,opt,name=gate_conn_id,json=gateConnId,proto3" json:"gate_conn_id,omitempty"` //客户端可不填
}

func (x *LogoutReq) Reset() {
	*x = LogoutReq{}
	mi := &file_lobby_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReq) ProtoMessage() {}

func (x *LogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReq.ProtoReflect.Descriptor instead.
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{2}
}

func (x *LogoutReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LogoutReq) GetGateConnId() uint64 {
	if x != nil {
		return x.GateConnId
	}
	return 0
}

// 注销登录
type LogoutRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason  LogoutRsp_LogoutReason `protobuf:"varint,1,opt,name=reason,proto3,enum=bs.lobby.LogoutRsp_LogoutReason" json:"reason,omitempty"`
	ErrInfo *types.ErrorInfo       `protobuf:"bytes,99,opt,name=err_info,json=errInfo,proto3" json:"err_info,omitempty"`
}

func (x *LogoutRsp) Reset() {
	*x = LogoutRsp{}
	mi := &file_lobby_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRsp) ProtoMessage() {}

func (x *LogoutRsp) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRsp.ProtoReflect.Descriptor instead.
func (*LogoutRsp) Descriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{3}
}

func (x *LogoutRsp) GetReason() LogoutRsp_LogoutReason {
	if x != nil {
		return x.Reason
	}
	return LogoutRsp_Normal
}

func (x *LogoutRsp) GetErrInfo() *types.ErrorInfo {
	if x != nil {
		return x.ErrInfo
	}
	return nil
}

type LogicRegReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogicRegReq) Reset() {
	*x = LogicRegReq{}
	mi := &file_lobby_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogicRegReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogicRegReq) ProtoMessage() {}

func (x *LogicRegReq) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogicRegReq.ProtoReflect.Descriptor instead.
func (*LogicRegReq) Descriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{4}
}

type LogicRegRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrInfo *types.ErrorInfo `protobuf:"bytes,99,opt,name=err_info,json=errInfo,proto3" json:"err_info,omitempty"`
}

func (x *LogicRegRsp) Reset() {
	*x = LogicRegRsp{}
	mi := &file_lobby_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogicRegRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogicRegRsp) ProtoMessage() {}

func (x *LogicRegRsp) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogicRegRsp.ProtoReflect.Descriptor instead.
func (*LogicRegRsp) Descriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{5}
}

func (x *LogicRegRsp) GetErrInfo() *types.ErrorInfo {
	if x != nil {
		return x.ErrInfo
	}
	return nil
}

// 查询用户
type QueryUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                //用户ID
	QueryType  uint32 `protobuf:"varint,2,opt,name=query_type,json=queryType,proto3" json:"query_type,omitempty"`       //操作类型 0进房查询
	RoomConnId uint64 `protobuf:"varint,19,opt,name=room_conn_id,json=roomConnId,proto3" json:"room_conn_id,omitempty"` //所在房间
}

func (x *QueryUserInfoReq) Reset() {
	*x = QueryUserInfoReq{}
	mi := &file_lobby_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserInfoReq) ProtoMessage() {}

func (x *QueryUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserInfoReq.ProtoReflect.Descriptor instead.
func (*QueryUserInfoReq) Descriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{6}
}

func (x *QueryUserInfoReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *QueryUserInfoReq) GetQueryType() uint32 {
	if x != nil {
		return x.QueryType
	}
	return 0
}

func (x *QueryUserInfoReq) GetRoomConnId() uint64 {
	if x != nil {
		return x.RoomConnId
	}
	return 0
}

// 查询用户
type QueryUserInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo  *types.BaseUserInfo `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`     //用户信息
	QueryType uint32              `protobuf:"varint,2,opt,name=query_type,json=queryType,proto3" json:"query_type,omitempty"` //操作类型 0进房查询
	ErrInfo   *types.ErrorInfo    `protobuf:"bytes,99,opt,name=err_info,json=errInfo,proto3" json:"err_info,omitempty"`
}

func (x *QueryUserInfoRsp) Reset() {
	*x = QueryUserInfoRsp{}
	mi := &file_lobby_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryUserInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserInfoRsp) ProtoMessage() {}

func (x *QueryUserInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserInfoRsp.ProtoReflect.Descriptor instead.
func (*QueryUserInfoRsp) Descriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{7}
}

func (x *QueryUserInfoRsp) GetUserInfo() *types.BaseUserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *QueryUserInfoRsp) GetQueryType() uint32 {
	if x != nil {
		return x.QueryType
	}
	return 0
}

func (x *QueryUserInfoRsp) GetErrInfo() *types.ErrorInfo {
	if x != nil {
		return x.ErrInfo
	}
	return nil
}

var File_lobby_proto protoreflect.FileDescriptor

var file_lobby_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62,
	0x73, 0x2e, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x03, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x3b,
	0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x62, 0x73, 0x2e, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x61, 0x63, 0x63, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x10, 0x02, 0x22, 0xbd, 0x03, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x12,
	0x31, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x62, 0x73, 0x2e, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x73, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x33, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x73, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07,
	0x65, 0x72, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x98, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x54, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x52, 0x4f, 0x5a, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x4c,
	0x53, 0x45, 0x50, 0x57, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x45, 0x54, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x50, 0x50, 0x49, 0x53, 0x42, 0x55, 0x53,
	0x59, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x55, 0x45, 0x53, 0x54, 0x46, 0x4f, 0x52, 0x42,
	0x49, 0x44, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x4f, 0x4c, 0x44, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x4d, 0x4f, 0x52, 0x45,
	0x47, 0x55, 0x45, 0x53, 0x54, 0x10, 0x09, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x52, 0x45, 0x51, 0x55,
	0x45, 0x4e, 0x54, 0x4c, 0x59, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x50, 0x50, 0x49, 0x4e,
	0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0b, 0x53, 0x45, 0x52, 0x56, 0x45,
	0x52, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xff, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x55, 0x4e, 0x4b,
	0x4f, 0x57, 0x4e, 0x10, 0xe8, 0x07, 0x12, 0x11, 0x0a, 0x0c, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xe9, 0x07, 0x12, 0x12, 0x0a, 0x0d, 0x54, 0x4f, 0x4b,
	0x45, 0x4e, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0xea, 0x07, 0x12, 0x13, 0x0a,
	0x0e, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10,
	0xeb, 0x07, 0x22, 0x46, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x61, 0x74, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x67, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x09, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x62, 0x73, 0x2e, 0x6c, 0x6f,
	0x62, 0x62, 0x79, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x73, 0x70, 0x2e, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x63,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x65, 0x72, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x2c, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x01,
	0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x22,
	0x3d, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x52, 0x65, 0x67, 0x52, 0x73, 0x70, 0x12, 0x2e,
	0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x62, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x65, 0x72, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x6c,
	0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x64, 0x22, 0x96, 0x01, 0x0a,
	0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73,
	0x70, 0x12, 0x33, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x73, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x65, 0x72,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0xae, 0x01, 0x0a, 0x08, 0x43, 0x4d, 0x44, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x44, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x49, 0x44, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x49, 0x44, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x10, 0x02, 0x12, 0x0f,
	0x0a, 0x0b, 0x49, 0x44, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x10, 0x03, 0x12,
	0x0f, 0x0a, 0x0b, 0x49, 0x44, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x73, 0x70, 0x10, 0x04,
	0x12, 0x11, 0x0a, 0x0d, 0x49, 0x44, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x52, 0x65, 0x67, 0x52, 0x65,
	0x71, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x44, 0x4c, 0x6f, 0x67, 0x69, 0x63, 0x52, 0x65,
	0x67, 0x52, 0x73, 0x70, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x44, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x10, 0x07, 0x12, 0x16,
	0x0a, 0x12, 0x49, 0x44, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x73, 0x70, 0x10, 0x08, 0x42, 0x11, 0x5a, 0x0f, 0x6d, 0x61, 0x6e, 0x67, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_lobby_proto_rawDescOnce sync.Once
	file_lobby_proto_rawDescData = file_lobby_proto_rawDesc
)

func file_lobby_proto_rawDescGZIP() []byte {
	file_lobby_proto_rawDescOnce.Do(func() {
		file_lobby_proto_rawDescData = protoimpl.X.CompressGZIP(file_lobby_proto_rawDescData)
	})
	return file_lobby_proto_rawDescData
}

var file_lobby_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_lobby_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_lobby_proto_goTypes = []any{
	(CMDLobby)(0),               // 0: bs.lobby.CMDLobby
	(LoginReq_LoginType)(0),     // 1: bs.lobby.LoginReq.LoginType
	(LoginRsp_Result)(0),        // 2: bs.lobby.LoginRsp.Result
	(LogoutRsp_LogoutReason)(0), // 3: bs.lobby.LogoutRsp.LogoutReason
	(*LoginReq)(nil),            // 4: bs.lobby.LoginReq
	(*LoginRsp)(nil),            // 5: bs.lobby.LoginRsp
	(*LogoutReq)(nil),           // 6: bs.lobby.LogoutReq
	(*LogoutRsp)(nil),           // 7: bs.lobby.LogoutRsp
	(*LogicRegReq)(nil),         // 8: bs.lobby.LogicRegReq
	(*LogicRegRsp)(nil),         // 9: bs.lobby.LogicRegRsp
	(*QueryUserInfoReq)(nil),    // 10: bs.lobby.QueryUserInfoReq
	(*QueryUserInfoRsp)(nil),    // 11: bs.lobby.QueryUserInfoRsp
	(*types.BaseUserInfo)(nil),  // 12: bs.types.BaseUserInfo
	(*types.ErrorInfo)(nil),     // 13: bs.types.ErrorInfo
}
var file_lobby_proto_depIdxs = []int32{
	1,  // 0: bs.lobby.LoginReq.login_type:type_name -> bs.lobby.LoginReq.LoginType
	2,  // 1: bs.lobby.LoginRsp.result:type_name -> bs.lobby.LoginRsp.Result
	12, // 2: bs.lobby.LoginRsp.user_info:type_name -> bs.types.BaseUserInfo
	13, // 3: bs.lobby.LoginRsp.err_info:type_name -> bs.types.ErrorInfo
	3,  // 4: bs.lobby.LogoutRsp.reason:type_name -> bs.lobby.LogoutRsp.LogoutReason
	13, // 5: bs.lobby.LogoutRsp.err_info:type_name -> bs.types.ErrorInfo
	13, // 6: bs.lobby.LogicRegRsp.err_info:type_name -> bs.types.ErrorInfo
	12, // 7: bs.lobby.QueryUserInfoRsp.user_info:type_name -> bs.types.BaseUserInfo
	13, // 8: bs.lobby.QueryUserInfoRsp.err_info:type_name -> bs.types.ErrorInfo
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_lobby_proto_init() }
func file_lobby_proto_init() {
	if File_lobby_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lobby_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lobby_proto_goTypes,
		DependencyIndexes: file_lobby_proto_depIdxs,
		EnumInfos:         file_lobby_proto_enumTypes,
		MessageInfos:      file_lobby_proto_msgTypes,
	}.Build()
	File_lobby_proto = out.File
	file_lobby_proto_rawDesc = nil
	file_lobby_proto_goTypes = nil
	file_lobby_proto_depIdxs = nil
}
