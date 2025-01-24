package business

import (
	"fmt"
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

	//存在判断
	var userId uint64 = 0
	for _, v := range userList {
		if v.GetAccount() == m.GetAccount() {
			userId = v.GetUserId()
			v.GateConnId = *proto.Uint64(gateConnId)
		}
	}
	if userId == 0 {
		var err error
		userId, err = dbUserLogin(gateConnId, m)
		if err != nil {
			log.Error("", "数据库登录失败,err=%v", err)
			return
		}
	}

	log.Debug("登录", "收到登录,AppType=%v,AppID=%v,Account=%v,gateConnId=%d,userId=%d",
		b.AgentInfo.AppType, b.AgentInfo.AppId, m.GetAccount(), gateConnId, userId)

	if userId == 0 {
		respondUserLogin(userId, gateConnId, int32(lobby.LoginRsp_NOTEXIST), "用户不存在")
		return
	}

	//查询财富
	req := property.QueryPropertyReq{
		UserId:     userId,
		GateConnId: gateConnId,
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
	userId := m.GetUserId()

	if _, ok := userList[userId]; !ok {
		log.Warning("", "查询用户人没有?,uId=%v", userId)

		g.SendData(srcApp, n.BaseMessage{MyMessage: &lobby.QueryUserInfoRsp{
			ErrInfo: &types.ErrorInfo{
				Code: types.ErrorInfo_failed,
				Info: "查询用户人没有?",
			},
		}, Cmd: n.TCPCommand{MainCmdID: uint16(n.AppLobby), SubCmdID: uint16(lobby.CMDLobby_IDQueryUserInfoRsp)}})
		return
	}
	log.Info("", "查询用户,uid=%d,cId=%v", m.GetUserId(), userList[userId].GetGateConnId())

	rsp := lobby.QueryUserInfoRsp{
		UserInfo: userList[m.GetUserId()],
		ErrInfo: &types.ErrorInfo{
			Code: types.ErrorInfo_success,
		},
	}
	g.SendData(srcApp, n.BaseMessage{MyMessage: &rsp, Cmd: n.TCPCommand{MainCmdID: uint16(n.AppLobby), SubCmdID: uint16(lobby.CMDLobby_IDQueryUserInfoRsp)}})
}
func handleQueryPropertyRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*property.QueryPropertyRsp)

	userId := m.GetUserId()
	connId := m.GetGateConnId()

	if _, ok := userList[userId]; !ok {
		log.Warning("", "财富回来人没了?,uId=%v,cId=%v,code=%v,Info=%v", userId, connId, m.GetErrInfo().GetCode(), m.GetErrInfo().GetInfo())
		respondUserLogin(userId, connId, int32(lobby.LoginRsp_SERVERERROR), fmt.Sprintf("财富回来人没了?uId=%v", m.GetUserId()))
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

	//发送回复
	respondUserLogin(userId, connId, int32(lobby.LoginRsp_SUCCESS), "登录成功")
}

func respondUserLogin(userId, connId uint64, errCode int32, errInfo string) {

	log.Debug("", "登录回复,uId=%v,cId=%v,code=%v,errInfo=%v", userId, connId, errCode, errInfo)

	rsp := lobby.LoginRsp{
		ErrInfo: &types.ErrorInfo{
			Info: errInfo,
			Code: types.ErrorInfo_ResultCode(errCode),
		},
	}
	rsp.UserInfo = new(types.BaseUserInfo)
	rsp.UserInfo = userList[userId]
	rspBm := n.BaseMessage{MyMessage: &rsp, Cmd: n.TCPCommand{MainCmdID: uint16(n.AppLobby), SubCmdID: uint16(lobby.CMDLobby_IDLoginRsp)}}
	g.SendMessage2Client(rspBm, userList[userId].GetGateConnId())
}
