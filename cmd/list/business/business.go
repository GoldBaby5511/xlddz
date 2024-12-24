package business

import (
	"mango/api/gateway"
	"mango/api/list"
	"mango/api/types"
	g "mango/pkg/gate"
	"mango/pkg/log"
	n "mango/pkg/network"
	"mango/pkg/util"
)

var (
	roomList = make(map[uint64]*types.RoomInfo)
)

func init() {
	g.MsgRegister(&list.RoomRegisterReq{}, n.AppList, uint16(list.CMDList_IDRoomRegisterReq), handleRoomRegisterReq)
	g.MsgRegister(&list.RoomListReq{}, n.AppList, uint16(list.CMDList_IDRoomListReq), handleRoomListReq)
}

func handleRoomRegisterReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*list.RoomRegisterReq)
	srcApp := args[n.OtherIndex].(n.BaseAgentInfo)

	regKey := util.MakeUint64FromUint32(m.GetInfo().GetAppInfo().GetType(), m.GetInfo().GetAppInfo().GetId())
	roomList[regKey] = m.GetInfo()
	log.Debug("", "收到注册,srcAppid=%d,len=%d", srcApp.AppId, m.GetInfo().GetAppInfo().GetId())

	rsp := list.RoomRegisterRsp{
		ErrInfo: &types.ErrorInfo{
			Info: "成功",
			Code: int32(list.RoomRegisterRsp_SUCCESS),
		},
	}
	g.SendData2App(srcApp.AppType, srcApp.AppId, n.AppList, uint32(list.CMDList_IDRoomRegisterRsp), &rsp)
}

func handleRoomListReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*list.RoomListReq)
	srcData := args[n.OtherIndex].(*gateway.TransferDataReq)

	log.Debug("", "收到列表请求,listID=%d,count=%d", m.GetListId(), len(roomList))

	var rsp list.RoomListRsp
	for _, r := range roomList {
		room := new(types.RoomInfo)
		room = r
		rsp.Rooms = append(rsp.Rooms, room)
	}
	rspBm := n.BaseMessage{MyMessage: &rsp, TraceId: ""}
	rspBm.Cmd = n.TCPCommand{MainCmdID: uint16(n.AppList), SubCmdID: uint16(list.CMDList_IDRoomListRsp)}
	g.SendMessage2Client(rspBm, srcData.GetGateconnid(), 0)
}
