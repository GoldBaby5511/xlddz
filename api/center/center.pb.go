// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.20.3
// source: center.proto

package center

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

type CMDCenter int32

const (
	CMDCenter_IDNone          CMDCenter = 0
	CMDCenter_IDAppRegReq     CMDCenter = 1 //服务注册
	CMDCenter_IDAppRegRsp     CMDCenter = 2 //服务注册
	CMDCenter_IDAppState      CMDCenter = 3 //服务状态
	CMDCenter_IDHeartBeatReq  CMDCenter = 4 //服务心跳
	CMDCenter_IDHeartBeatRsp  CMDCenter = 5 //服务心跳
	CMDCenter_IDAppControlReq CMDCenter = 6 //控制消息
	CMDCenter_IDAppControlRsp CMDCenter = 7 //控制消息
)

// Enum value maps for CMDCenter.
var (
	CMDCenter_name = map[int32]string{
		0: "IDNone",
		1: "IDAppRegReq",
		2: "IDAppRegRsp",
		3: "IDAppState",
		4: "IDHeartBeatReq",
		5: "IDHeartBeatRsp",
		6: "IDAppControlReq",
		7: "IDAppControlRsp",
	}
	CMDCenter_value = map[string]int32{
		"IDNone":          0,
		"IDAppRegReq":     1,
		"IDAppRegRsp":     2,
		"IDAppState":      3,
		"IDHeartBeatReq":  4,
		"IDHeartBeatRsp":  5,
		"IDAppControlReq": 6,
		"IDAppControlRsp": 7,
	}
)

func (x CMDCenter) Enum() *CMDCenter {
	p := new(CMDCenter)
	*p = x
	return p
}

func (x CMDCenter) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMDCenter) Descriptor() protoreflect.EnumDescriptor {
	return file_center_proto_enumTypes[0].Descriptor()
}

func (CMDCenter) Type() protoreflect.EnumType {
	return &file_center_proto_enumTypes[0]
}

func (x CMDCenter) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CMDCenter.Descriptor instead.
func (CMDCenter) EnumDescriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{0}
}

type CtlId int32

const (
	CtlId_None              CtlId = 0
	CtlId_Maintenance       CtlId = 1 //开始维护
	CtlId_MaintenanceFinish CtlId = 2 //维护完成
	CtlId_ShowServerList    CtlId = 3 //显示列表
	CtlId_StartService      CtlId = 4 //启动服务
	CtlId_StopService       CtlId = 5 //停止服务
	CtlId_UpdateService     CtlId = 6 //更新服务
)

// Enum value maps for CtlId.
var (
	CtlId_name = map[int32]string{
		0: "None",
		1: "Maintenance",
		2: "MaintenanceFinish",
		3: "ShowServerList",
		4: "StartService",
		5: "StopService",
		6: "UpdateService",
	}
	CtlId_value = map[string]int32{
		"None":              0,
		"Maintenance":       1,
		"MaintenanceFinish": 2,
		"ShowServerList":    3,
		"StartService":      4,
		"StopService":       5,
		"UpdateService":     6,
	}
)

func (x CtlId) Enum() *CtlId {
	p := new(CtlId)
	*p = x
	return p
}

func (x CtlId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CtlId) Descriptor() protoreflect.EnumDescriptor {
	return file_center_proto_enumTypes[1].Descriptor()
}

func (CtlId) Type() protoreflect.EnumType {
	return &file_center_proto_enumTypes[1]
}

func (x CtlId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CtlId.Descriptor instead.
func (CtlId) EnumDescriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{1}
}

// 服务注册
type RegisterAppReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthKey    string `protobuf:"bytes,1,opt,name=auth_key,json=authKey,proto3" json:"auth_key,omitempty"`
	AttData    string `protobuf:"bytes,2,opt,name=att_data,json=attData,proto3" json:"att_data,omitempty"`
	MyAddress  string `protobuf:"bytes,3,opt,name=my_address,json=myAddress,proto3" json:"my_address,omitempty"`
	AppType    uint32 `protobuf:"varint,4,opt,name=app_type,json=appType,proto3" json:"app_type,omitempty"`
	AppId      uint32 `protobuf:"varint,5,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ReregToken string `protobuf:"bytes,6,opt,name=rereg_token,json=reregToken,proto3" json:"rereg_token,omitempty"` //如果中间网络断开了,可以使用rereg_token强行再次注册
	AppName    string `protobuf:"bytes,7,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`          //app的名称(一般为进程名)
}

func (x *RegisterAppReq) Reset() {
	*x = RegisterAppReq{}
	mi := &file_center_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterAppReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAppReq) ProtoMessage() {}

func (x *RegisterAppReq) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAppReq.ProtoReflect.Descriptor instead.
func (*RegisterAppReq) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterAppReq) GetAuthKey() string {
	if x != nil {
		return x.AuthKey
	}
	return ""
}

func (x *RegisterAppReq) GetAttData() string {
	if x != nil {
		return x.AttData
	}
	return ""
}

func (x *RegisterAppReq) GetMyAddress() string {
	if x != nil {
		return x.MyAddress
	}
	return ""
}

func (x *RegisterAppReq) GetAppType() uint32 {
	if x != nil {
		return x.AppType
	}
	return 0
}

func (x *RegisterAppReq) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *RegisterAppReq) GetReregToken() string {
	if x != nil {
		return x.ReregToken
	}
	return ""
}

func (x *RegisterAppReq) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

// 服务注册
type RegisterAppRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegResult  uint32 `protobuf:"varint,1,opt,name=reg_result,json=regResult,proto3" json:"reg_result,omitempty"`   //0表示成功，其它为错误码(rereg_token为出错内容)
	ReregToken string `protobuf:"bytes,2,opt,name=rereg_token,json=reregToken,proto3" json:"rereg_token,omitempty"` //如果中间网络断开了,可以使用rereg_token强行再次注册
	CenterId   uint32 `protobuf:"varint,3,opt,name=center_id,json=centerId,proto3" json:"center_id,omitempty"`
	AppType    uint32 `protobuf:"varint,4,opt,name=app_type,json=appType,proto3" json:"app_type,omitempty"` //Router 或其他
	AppId      uint32 `protobuf:"varint,5,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	AppName    string `protobuf:"bytes,6,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`          //app的名称(一般为进程名)
	AppAddress string `protobuf:"bytes,7,opt,name=app_address,json=appAddress,proto3" json:"app_address,omitempty"` //监听地址
}

func (x *RegisterAppRsp) Reset() {
	*x = RegisterAppRsp{}
	mi := &file_center_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterAppRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAppRsp) ProtoMessage() {}

func (x *RegisterAppRsp) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAppRsp.ProtoReflect.Descriptor instead.
func (*RegisterAppRsp) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterAppRsp) GetRegResult() uint32 {
	if x != nil {
		return x.RegResult
	}
	return 0
}

func (x *RegisterAppRsp) GetReregToken() string {
	if x != nil {
		return x.ReregToken
	}
	return ""
}

func (x *RegisterAppRsp) GetCenterId() uint32 {
	if x != nil {
		return x.CenterId
	}
	return 0
}

func (x *RegisterAppRsp) GetAppType() uint32 {
	if x != nil {
		return x.AppType
	}
	return 0
}

func (x *RegisterAppRsp) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *RegisterAppRsp) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *RegisterAppRsp) GetAppAddress() string {
	if x != nil {
		return x.AppAddress
	}
	return ""
}

// 服务状态
type AppStateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppState int32  `protobuf:"varint,1,opt,name=app_state,json=appState,proto3" json:"app_state,omitempty"`
	CenterId uint32 `protobuf:"varint,2,opt,name=center_id,json=centerId,proto3" json:"center_id,omitempty"`
	AppType  uint32 `protobuf:"varint,4,opt,name=app_type,json=appType,proto3" json:"app_type,omitempty"`
	AppId    uint32 `protobuf:"varint,5,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *AppStateNotify) Reset() {
	*x = AppStateNotify{}
	mi := &file_center_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppStateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppStateNotify) ProtoMessage() {}

func (x *AppStateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppStateNotify.ProtoReflect.Descriptor instead.
func (*AppStateNotify) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{2}
}

func (x *AppStateNotify) GetAppState() int32 {
	if x != nil {
		return x.AppState
	}
	return 0
}

func (x *AppStateNotify) GetCenterId() uint32 {
	if x != nil {
		return x.CenterId
	}
	return 0
}

func (x *AppStateNotify) GetAppType() uint32 {
	if x != nil {
		return x.AppType
	}
	return 0
}

func (x *AppStateNotify) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

// 服务心跳
type HeartBeatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeatId           int64  `protobuf:"varint,1,opt,name=beat_id,json=beatId,proto3" json:"beat_id,omitempty"`
	PulseTime        int64  `protobuf:"varint,2,opt,name=pulse_time,json=pulseTime,proto3" json:"pulse_time,omitempty"`
	ServiceState     int32  `protobuf:"varint,3,opt,name=service_state,json=serviceState,proto3" json:"service_state,omitempty"`
	StateDescription string `protobuf:"bytes,4,opt,name=state_description,json=stateDescription,proto3" json:"state_description,omitempty"`
	HttpAddress      string `protobuf:"bytes,5,opt,name=http_address,json=httpAddress,proto3" json:"http_address,omitempty"`
	RpcAddress       string `protobuf:"bytes,6,opt,name=rpc_address,json=rpcAddress,proto3" json:"rpc_address,omitempty"`
}

func (x *HeartBeatReq) Reset() {
	*x = HeartBeatReq{}
	mi := &file_center_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartBeatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatReq) ProtoMessage() {}

func (x *HeartBeatReq) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatReq.ProtoReflect.Descriptor instead.
func (*HeartBeatReq) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{3}
}

func (x *HeartBeatReq) GetBeatId() int64 {
	if x != nil {
		return x.BeatId
	}
	return 0
}

func (x *HeartBeatReq) GetPulseTime() int64 {
	if x != nil {
		return x.PulseTime
	}
	return 0
}

func (x *HeartBeatReq) GetServiceState() int32 {
	if x != nil {
		return x.ServiceState
	}
	return 0
}

func (x *HeartBeatReq) GetStateDescription() string {
	if x != nil {
		return x.StateDescription
	}
	return ""
}

func (x *HeartBeatReq) GetHttpAddress() string {
	if x != nil {
		return x.HttpAddress
	}
	return ""
}

func (x *HeartBeatReq) GetRpcAddress() string {
	if x != nil {
		return x.RpcAddress
	}
	return ""
}

// 服务心跳
type HeartBeatRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PulseTime int64 `protobuf:"varint,1,opt,name=pulse_time,json=pulseTime,proto3" json:"pulse_time,omitempty"`
}

func (x *HeartBeatRsp) Reset() {
	*x = HeartBeatRsp{}
	mi := &file_center_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartBeatRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatRsp) ProtoMessage() {}

func (x *HeartBeatRsp) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatRsp.ProtoReflect.Descriptor instead.
func (*HeartBeatRsp) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{4}
}

func (x *HeartBeatRsp) GetPulseTime() int64 {
	if x != nil {
		return x.PulseTime
	}
	return 0
}

type ControlItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type    uint32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Id      uint32   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Command string   `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"` //命令
	Args    []string `protobuf:"bytes,5,rep,name=args,proto3" json:"args,omitempty"`       //参数
}

func (x *ControlItem) Reset() {
	*x = ControlItem{}
	mi := &file_center_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ControlItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlItem) ProtoMessage() {}

func (x *ControlItem) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlItem.ProtoReflect.Descriptor instead.
func (*ControlItem) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{5}
}

func (x *ControlItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ControlItem) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ControlItem) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ControlItem) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ControlItem) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type AppControlReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CtlId      int32          `protobuf:"varint,1,opt,name=ctl_id,json=ctlId,proto3" json:"ctl_id,omitempty"` // 命令编号
	AppType    uint32         `protobuf:"varint,2,opt,name=app_type,json=appType,proto3" json:"app_type,omitempty"`
	AppId      uint32         `protobuf:"varint,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	CtlServers []*ControlItem `protobuf:"bytes,4,rep,name=ctl_servers,json=ctlServers,proto3" json:"ctl_servers,omitempty"` //操作服务
	Args       []string       `protobuf:"bytes,5,rep,name=args,proto3" json:"args,omitempty"`                               //参数
}

func (x *AppControlReq) Reset() {
	*x = AppControlReq{}
	mi := &file_center_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppControlReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppControlReq) ProtoMessage() {}

func (x *AppControlReq) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppControlReq.ProtoReflect.Descriptor instead.
func (*AppControlReq) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{6}
}

func (x *AppControlReq) GetCtlId() int32 {
	if x != nil {
		return x.CtlId
	}
	return 0
}

func (x *AppControlReq) GetAppType() uint32 {
	if x != nil {
		return x.AppType
	}
	return 0
}

func (x *AppControlReq) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *AppControlReq) GetCtlServers() []*ControlItem {
	if x != nil {
		return x.CtlServers
	}
	return nil
}

func (x *AppControlReq) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type AppControlRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CtlId   int32  `protobuf:"varint,1,opt,name=ctl_id,json=ctlId,proto3" json:"ctl_id,omitempty"` // 命令编号
	AppType uint32 `protobuf:"varint,2,opt,name=app_type,json=appType,proto3" json:"app_type,omitempty"`
	AppId   uint32 `protobuf:"varint,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Code    int32  `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Info    string `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *AppControlRsp) Reset() {
	*x = AppControlRsp{}
	mi := &file_center_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppControlRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppControlRsp) ProtoMessage() {}

func (x *AppControlRsp) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppControlRsp.ProtoReflect.Descriptor instead.
func (*AppControlRsp) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{7}
}

func (x *AppControlRsp) GetCtlId() int32 {
	if x != nil {
		return x.CtlId
	}
	return 0
}

func (x *AppControlRsp) GetAppType() uint32 {
	if x != nil {
		return x.AppType
	}
	return 0
}

func (x *AppControlRsp) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *AppControlRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AppControlRsp) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_center_proto protoreflect.FileDescriptor

var file_center_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x62, 0x73, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x22, 0xd3, 0x01, 0x0a, 0x0e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x74, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x72, 0x65, 0x67, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x72, 0x65, 0x67, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0xdb, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x52,
	0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x65, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x72, 0x65, 0x67, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x72, 0x65, 0x67, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x70, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x7c, 0x0a,
	0x0e, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x61, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x70, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0xdc, 0x01, 0x0a, 0x0c,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x62, 0x65, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62,
	0x65, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x6c, 0x73, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x75, 0x6c, 0x73, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x74,
	0x74, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x70, 0x63,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2d, 0x0a, 0x0c, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75,
	0x6c, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x75, 0x6c, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x73, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72,
	0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0xa5,
	0x01, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71,
	0x12, 0x15, 0x0a, 0x06, 0x63, 0x74, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x74, 0x6c, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x74, 0x6c,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x62, 0x73, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x63, 0x74, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x73, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x74, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x74, 0x6c, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x2a, 0x9b, 0x01, 0x0a, 0x09, 0x43, 0x4d,
	0x44, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x44, 0x4e, 0x6f, 0x6e,
	0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x44, 0x41, 0x70, 0x70, 0x52, 0x65, 0x67, 0x52,
	0x65, 0x71, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x44, 0x41, 0x70, 0x70, 0x52, 0x65, 0x67,
	0x52, 0x73, 0x70, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x44, 0x41, 0x70, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x44, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x44, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x73, 0x70, 0x10, 0x05, 0x12, 0x13, 0x0a,
	0x0f, 0x49, 0x44, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71,
	0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x44, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x52, 0x73, 0x70, 0x10, 0x07, 0x2a, 0x83, 0x01, 0x0a, 0x05, 0x43, 0x74, 0x6c, 0x49,
	0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4d,
	0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11,
	0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x74, 0x6f,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0x06, 0x32, 0x4e, 0x0a,
	0x0a, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x40, 0x0a, 0x0a, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x18, 0x2e, 0x62, 0x73, 0x2e, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x62, 0x73, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x73, 0x70, 0x42, 0x09, 0x5a,
	0x07, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_center_proto_rawDescOnce sync.Once
	file_center_proto_rawDescData = file_center_proto_rawDesc
)

func file_center_proto_rawDescGZIP() []byte {
	file_center_proto_rawDescOnce.Do(func() {
		file_center_proto_rawDescData = protoimpl.X.CompressGZIP(file_center_proto_rawDescData)
	})
	return file_center_proto_rawDescData
}

var file_center_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_center_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_center_proto_goTypes = []any{
	(CMDCenter)(0),         // 0: bs.center.CMDCenter
	(CtlId)(0),             // 1: bs.center.CtlId
	(*RegisterAppReq)(nil), // 2: bs.center.RegisterAppReq
	(*RegisterAppRsp)(nil), // 3: bs.center.RegisterAppRsp
	(*AppStateNotify)(nil), // 4: bs.center.AppStateNotify
	(*HeartBeatReq)(nil),   // 5: bs.center.HeartBeatReq
	(*HeartBeatRsp)(nil),   // 6: bs.center.HeartBeatRsp
	(*ControlItem)(nil),    // 7: bs.center.controlItem
	(*AppControlReq)(nil),  // 8: bs.center.AppControlReq
	(*AppControlRsp)(nil),  // 9: bs.center.AppControlRsp
}
var file_center_proto_depIdxs = []int32{
	7, // 0: bs.center.AppControlReq.ctl_servers:type_name -> bs.center.controlItem
	8, // 1: bs.center.AppControl.ControlReq:input_type -> bs.center.AppControlReq
	9, // 2: bs.center.AppControl.ControlReq:output_type -> bs.center.AppControlRsp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_center_proto_init() }
func file_center_proto_init() {
	if File_center_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_center_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_center_proto_goTypes,
		DependencyIndexes: file_center_proto_depIdxs,
		EnumInfos:         file_center_proto_enumTypes,
		MessageInfos:      file_center_proto_msgTypes,
	}.Build()
	File_center_proto = out.File
	file_center_proto_rawDesc = nil
	file_center_proto_goTypes = nil
	file_center_proto_depIdxs = nil
}
