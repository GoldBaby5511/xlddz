package business

import (
	"github.com/golang/protobuf/proto"
	"mango/api/gateway"
	"mango/api/lobby"
	"mango/api/property"
	"mango/api/types"
	g "mango/pkg/gate"
	"mango/pkg/log"
	n "mango/pkg/network"
	"mango/pkg/util"
)

var (
	userList = make(map[uint64]*types.BaseUserInfo)
)

func init() {
	g.MsgRegister(&lobby.LoginReq{}, n.AppLobby, uint16(lobby.CMDLobby_IDLoginReq), handleLoginReq)
	g.MsgRegister(&lobby.LogoutReq{}, n.AppLobby, uint16(lobby.CMDLobby_IDLogoutReq), handleLogoutReq)
	g.MsgRegister(&lobby.QueryUserInfoReq{}, n.AppLobby, uint16(lobby.CMDLobby_IDQueryUserInfoReq), handleQueryUserInfoReq)
	g.MsgRegister(&property.QueryPropertyRsp{}, n.AppProperty, uint16(property.CMDProperty_IDQueryPropertyRsp), handleQueryPropertyRsp)
	g.CallBackRegister(g.CbAppControlNotify, appControlNotify)
}

func appControlNotify(args []interface{}) {

}

func handleLoginReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*lobby.LoginReq)
	srcData := b.AgentInfo
	gateConnId := util.MakeUint64FromUint32(srcData.AppType, srcData.AppId)

	log.Debug("登录", "收到登录,AppType=%v,AppID=%v,Account=%v,gateConnId=%d,子渠道=%d",
		b.AgentInfo.AppType, b.AgentInfo.AppId, m.GetAccount(), gateConnId, m.GetSiteId())

	var userId uint64 = 0
	for _, v := range userList {
		if v.GetAccount() == m.GetAccount() {
			userId = v.GetUserId()
			v.GateConnId = *proto.Uint64(gateConnId)
		}
	}
	if userId == 0 {
		userId = uint64(10000 + len(userList))
		u := &types.BaseUserInfo{
			Account:    m.GetAccount(),
			UserId:     userId,
			GameId:     userId,
			GateConnId: gateConnId,
		}
		userList[userId] = u
	}
	req := property.QueryPropertyReq{
		UserId: userId,
	}
	g.SendData2App(n.AppProperty, n.Send2AnyOne, n.AppProperty, uint32(property.CMDProperty_IDQueryPropertyReq), &req)
}

func handleLogoutReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*lobby.LogoutReq)
	log.Debug("注销", "注销请求,userId=%v", m.GetUserId())
}

func handleQueryUserInfoReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*lobby.QueryUserInfoReq)
	srcApp := b.AgentInfo

	log.Info("", "查询用户,uid=%d", m.GetUserId())

	rsp := lobby.QueryUserInfoRsp{
		UserInfo: userList[m.GetUserId()],
		ErrInfo: &types.ErrorInfo{
			Code: types.ErrorInfo_success,
		},
	}
	cmd := n.TCPCommand{MainCmdID: uint16(n.AppLobby), SubCmdID: uint16(lobby.CMDLobby_IDQueryUserInfoRsp)}
	bm := n.BaseMessage{MyMessage: &rsp, Cmd: cmd}
	g.SendData(srcApp, bm)
}
func handleQueryPropertyRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*property.QueryPropertyRsp)

	if _, ok := userList[m.GetUserId()]; !ok {
		return
	}
	userList[m.GetUserId()].Props = append(userList[m.GetUserId()].Props, m.GetUserProps()...)

	log.Debug("", "财富查询,userId=%v,len=%v,gateConnId=%d", m.GetUserId(), len(m.GetUserProps()), userList[m.GetUserId()].GetGateConnId())

	authRsp := gateway.AuthInfo{
		UserId:     m.GetUserId(),
		GateConnId: userList[m.GetUserId()].GetGateConnId(),
		Result:     uint32(lobby.LoginRsp_SUCCESS),
	}
	g.SendData2App(n.AppGate, util.GetLUint32FromUint64(userList[m.GetUserId()].GetGateConnId()), n.AppGate, uint32(gateway.CMDGateway_IDAuthInfo), &authRsp)

	rsp := lobby.LoginRsp{
		ErrInfo: &types.ErrorInfo{
			Info: "成功",
			Code: types.ErrorInfo_ResultCode(lobby.LoginRsp_SUCCESS),
		},
	}
	rsp.UserInfo = new(types.BaseUserInfo)
	rsp.UserInfo = userList[m.GetUserId()]
	rspBm := n.BaseMessage{MyMessage: &rsp, TraceId: ""}
	rspBm.Cmd = n.TCPCommand{MainCmdID: uint16(n.AppLobby), SubCmdID: uint16(lobby.CMDLobby_IDLoginRsp)}
	g.SendMessage2Client(rspBm, userList[m.GetUserId()].GetGateConnId())
}
