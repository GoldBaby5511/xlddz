package business

import (
	"github.com/golang/protobuf/proto"
	"mango/api/client"
	"mango/api/gate"
	"mango/api/property"
	"mango/api/types"
	g "mango/pkg/gate"
	"mango/pkg/log"
	n "mango/pkg/network"
	"mango/pkg/util"
)

var (
	pls map[uint64]*types.BaseUserInfo = make(map[uint64]*types.BaseUserInfo)
)

func init() {
	g.MsgRegister(&client.LoginReq{}, n.CMDClient, uint16(client.CMDID_Client_IDLoginReq), handleLoginReq)
	g.MsgRegister(&client.LogoutReq{}, n.CMDClient, uint16(client.CMDID_Client_IDLogoutReq), handleLogoutReq)
	g.MsgRegister(&property.QueryPropertyRsp{}, n.CMDProperty, uint16(property.CMDID_Property_IDQueryPropertyRsp), handleQueryPropertyRsp)
	g.EventRegister(g.ConnectSuccess, connectSuccess)
	g.EventRegister(g.Disconnect, disconnect)
}

func connectSuccess(args []interface{}) {
	log.Info("连接", "来了老弟,参数数量=%d", len(args))
}

func disconnect(args []interface{}) {
	log.Info("连接", "告辞中,参数数量=%d", len(args))
}

func handleLoginReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*client.LoginReq)
	srcData := b.AgentInfo
	gateConnId := util.MakeUint64FromUint32(srcData.AppType, srcData.AppID)

	log.Debug("登录", "收到登录,AppType=%v,AppID=%v,Account=%v,gateConnId=%d,子渠道=%d",
		b.AgentInfo.AppType, b.AgentInfo.AppID, m.GetLoginAccount(), gateConnId, m.GetSiteId())

	var userId uint64 = 0
	for _, v := range pls {
		if v.GetAccount() == m.GetLoginAccount() {
			userId = v.GetUserId()
			v.GateConnid = proto.Uint64(gateConnId)
		}
	}
	if userId == 0 {
		userId = uint64(10000 + len(pls))
		pls[userId] = new(types.BaseUserInfo)
		pls[userId].Account = proto.String(m.GetLoginAccount())
		pls[userId].UserId = proto.Uint64(userId)
		pls[userId].GameId = proto.Uint64(userId)
		pls[userId].GateConnid = proto.Uint64(gateConnId)
	}
	var req property.QueryPropertyReq
	req.UserId = proto.Uint64(userId)
	g.SendData2App(n.AppProperty, n.Send2AnyOne, n.CMDProperty, uint32(property.CMDID_Property_IDQueryPropertyReq), &req)
}

func handleLogoutReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*client.LogoutReq)
	log.Debug("注销", "注销请求,userId=%v", m.GetUserId())
}

func handleQueryPropertyRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*property.QueryPropertyRsp)

	if _, ok := pls[m.GetUserId()]; !ok {
		return
	}
	pls[m.GetUserId()].UserProps = append(pls[m.GetUserId()].UserProps, m.GetUserProps()...)

	log.Debug("", "财富查询,userId=%v,len=%v,GateConnid=%d", m.GetUserId(), len(m.GetUserProps()), pls[m.GetUserId()].GetGateConnid())

	var authRsp gate.AuthInfo
	authRsp.UserId = proto.Uint64(m.GetUserId())
	authRsp.Gateconnid = proto.Uint64(pls[m.GetUserId()].GetGateConnid())
	authRsp.Result = proto.Uint32(uint32(client.LoginRsp_SUCCESS))
	g.SendData2App(n.AppGate, util.GetLUint32FromUint64(pls[m.GetUserId()].GetGateConnid()), n.CMDGate, uint32(gate.CMDID_Gate_IDAuthInfo), &authRsp)

	var rsp client.LoginRsp
	rsp.LoginInfo = proto.String("成功")
	rsp.LoginResult = (*client.LoginRsp_Result)(proto.Int32(int32(client.LoginRsp_SUCCESS)))
	rsp.BaseInfo = new(types.BaseUserInfo)
	rsp.BaseInfo = pls[m.GetUserId()]
	rspBm := n.BaseMessage{MyMessage: &rsp, TraceId: ""}
	rspBm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDClient), SubCmdID: uint16(client.CMDID_Client_IDLoginRsp)}
	g.SendMessage2Client(rspBm, pls[m.GetUserId()].GetGateConnid(), 0)
}
