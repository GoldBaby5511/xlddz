package business

import (
	"mango/api/gateway"
	"mango/api/list"
	"mango/api/types"
	g "github.com/GoldBaby5511/go-mango-core/gate"
	"github.com/GoldBaby5511/go-mango-core/log"
	n "github.com/GoldBaby5511/go-mango-core/network"
	"github.com/GoldBaby5511/go-mango-core/util"
)

var (
	roomList map[uint64]*types.RoomInfo = make(map[uint64]*types.RoomInfo)
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
	log.Debug("", "收到注册,AttAppid=%d,len=%d", srcApp.AppID, m.GetInfo().GetAppInfo().GetId())
}

func handleRoomListReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*list.RoomListReq)
	srcData := args[n.OtherIndex].(*gateway.TransferDataReq)

	log.Debug("", "收到列表请求,listID=%d", m.GetListId())

	var rsp list.RoomListRsp
	for _, r := range roomList {
		room := new(types.RoomInfo)
		room = r
		rsp.Rooms = append(rsp.Rooms, room)
	}
	rspBm := n.BaseMessage{MyMessage: &rsp, TraceId: ""}
	rspBm.Cmd = n.TCPCommand{AppType: uint16(n.AppList), CmdId: uint16(list.CMDList_IDRoomListRsp)}
	g.SendMessage2Client(rspBm, srcData.GetGateconnid(), 0)
}
