package business

import (
	"github.com/golang/protobuf/proto"
	"mango/api/client"
	"mango/api/gate"
	"mango/api/types"
	g "mango/pkg/gate"
	"mango/pkg/log"
	n "mango/pkg/network"
)

var (
	pls map[string]*types.BaseUserInfo = make(map[string]*types.BaseUserInfo)
)

func init() {
	g.MsgRegister(&client.LoginReq{}, n.CMDClient, uint16(client.CMDID_Client_IDLoginReq), handleLoginReq)
	g.MsgRegister(&client.LogoutReq{}, n.CMDClient, uint16(client.CMDID_Client_IDLogoutReq), handleLogoutReq)
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
	srcData := args[n.OtherIndex].(*gate.TransferDataReq)
	gateConnId := srcData.GetGateconnid()

	log.Debug("登录", "收到登录,Account=%v,主渠道=%d,子渠道=%d", m.GetLoginAccount(), m.GetChannelId(), m.GetSiteId())

	var authRsp gate.AuthInfo
	authRsp.UserId = proto.Uint64(10001)
	authRsp.Gateconnid = proto.Uint64(gateConnId)
	authRsp.Result = proto.Uint32(uint32(client.LoginRsp_SUCCESS))
	g.SendData2App(n.AppGate, uint32(gateConnId>>32), n.CMDGate, uint32(gate.CMDID_Gate_IDAuthInfo), &authRsp)

	var rsp client.LoginRsp
	rsp.LoginInfo = proto.String("成功")
	rsp.LoginResult = (*client.LoginRsp_Result)(proto.Int32(int32(client.LoginRsp_SUCCESS)))
	rsp.BaseInfo = new(types.BaseUserInfo)
	if pl, ok := pls[m.GetLoginAccount()]; ok {
		rsp.BaseInfo = pl
	} else {
		rsp.BaseInfo.Account = proto.String(m.GetLoginAccount())
		rsp.BaseInfo.UserId = proto.Uint64(uint64(10000 + len(pls)))
		rsp.BaseInfo.GameId = proto.Uint64(uint64(10000 + len(pls)))
	}

	rspBm := n.BaseMessage{MyMessage: &rsp, TraceId: ""}
	rspBm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDClient), SubCmdID: uint16(client.CMDID_Client_IDLoginRsp)}
	g.SendMessage2Client(rspBm, gateConnId, 0)

	//sendLoginRsp(m, srcData, "成功", uint32(client.LoginRsp_SUCCESS))
}

func handleLogoutReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*client.LogoutReq)
	log.Debug("注销", "注销请求,userId=%v", m.GetUserId())
}

func sendLoginRsp(m *client.LoginReq, srcData *gate.TransferDataReq, info string, code uint32) {
	gateConnId := srcData.GetGateconnid()
	log.Info("登录", "发送登录响应,gateConnId=%v,info=%v,code=%v", gateConnId, info, code)

	var authRsp gate.AuthInfo
	authRsp.UserId = proto.Uint64(10001)
	authRsp.Gateconnid = proto.Uint64(gateConnId)
	authRsp.Result = proto.Uint32(code)
	g.SendData2App(n.AppGate, uint32(gateConnId>>32), n.CMDGate, uint32(gate.CMDID_Gate_IDAuthInfo), &authRsp)

	var rsp client.LoginRsp
	rsp.LoginInfo = proto.String(info)
	rsp.LoginResult = (*client.LoginRsp_Result)(proto.Int32(int32(code)))
	rsp.BaseInfo = new(types.BaseUserInfo)
	if pl, ok := pls[m.GetLoginAccount()]; ok {
		rsp.BaseInfo = pl
	} else {
		rsp.BaseInfo.Account = proto.String(m.GetLoginAccount())
		rsp.BaseInfo.UserId = proto.Uint64(uint64(10000 + len(pls)))
		rsp.BaseInfo.GameId = proto.Uint64(uint64(10000 + len(pls)))
	}

	rspBm := n.BaseMessage{MyMessage: &rsp, TraceId: ""}
	rspBm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDClient), SubCmdID: uint16(client.CMDID_Client_IDLoginRsp)}
	g.SendMessage2Client(rspBm, gateConnId, 0)
}
