package business

import (
	"mango/api/client"
	"mango/api/gate"
	"mango/api/list"
	"mango/api/types"
	g "mango/pkg/gate"
	"mango/pkg/log"
	n "mango/pkg/network"
	"mango/pkg/util"
)

var (
	roomList map[uint64]*types.RoomInfo = make(map[uint64]*types.RoomInfo)
)

func init() {
	g.MsgRegister(&list.RoomRegisterReq{}, n.CMDList, uint16(list.CMDID_List_IDRoomRegisterReq), handleRoomRegisterReq)
	g.MsgRegister(&client.RoomListReq{}, n.CMDClient, uint16(client.CMDID_Client_IDRoomListReq), handleRoomListReq)
	g.EventRegister(g.ConnectSuccess, connectSuccess)
	g.EventRegister(g.Disconnect, disconnect)
}

func connectSuccess(args []interface{}) {
}

func disconnect(args []interface{}) {
}

func handleRoomRegisterReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*list.RoomRegisterReq)
	srcApp := args[n.OtherIndex].(n.BaseAgentInfo)

	regKey := util.MakeUint64FromUint32(m.GetRoomInfo().GetRoomInfo().GetType(), m.GetRoomInfo().GetRoomInfo().GetId())
	roomList[regKey] = m.GetRoomInfo()
	//tables = append(tables, m.GetTableIds()...)
	//log.Debug("", "tables=%v", tables)
	log.Debug("", "收到注册,AttAppid=%d,len=%d", srcApp.AppID, m.GetRoomInfo().GetRoomInfo().GetId())
}

func handleRoomListReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*client.RoomListReq)
	srcData := args[n.OtherIndex].(*gate.TransferDataReq)

	//tables = append(tables, m.GetTableIds()...)
	//log.Debug("", "GetListId=%v", m.GetListId())
	log.Debug("", "收到列表请求,userID=%d,listID=%d", srcData.GetUserId(), m.GetListId())

	var rsp client.RoomListRsp
	for _, r := range roomList {
		room := new(types.RoomInfo)
		room = r
		rsp.Rooms = append(rsp.Rooms, room)
	}
	rspBm := n.BaseMessage{MyMessage: &rsp, TraceId: ""}
	rspBm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDClient), SubCmdID: uint16(client.CMDID_Client_IDRoomListRsp)}
	g.SendMessage2Client(rspBm, srcData.GetGateconnid(), 0)
}
