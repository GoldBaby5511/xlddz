package business

import (
	"github.com/golang/protobuf/proto"
	"mango/api/client"
	"mango/api/gate"
	"mango/api/list"
	"mango/api/property"
	rCMD "mango/api/room"
	tCMD "mango/api/table"
	"mango/api/types"
	"mango/cmd/room/business/player"
	"mango/pkg/conf"
	"mango/pkg/conf/apollo"
	g "mango/pkg/gate"
	"mango/pkg/log"
	n "mango/pkg/network"
	"mango/pkg/timer"
	"mango/pkg/util"
	"time"
)

var (
	tables  []uint64
	players map[uint64]*player.Player = make(map[uint64]*player.Player)
)

func init() {
	g.MsgRegister(&tCMD.ApplyRsp{}, n.CMDTable, uint16(tCMD.CMDID_Table_IDApplyRsp), handleApplyRsp)
	g.MsgRegister(&tCMD.WriteGameScore{}, n.CMDTable, uint16(tCMD.CMDID_Table_IDWriteGameScore), handleWriteGameScore)
	g.MsgRegister(&tCMD.GameOver{}, n.CMDTable, uint16(tCMD.CMDID_Table_IDGameOver), handleGameOver)
	g.MsgRegister(&list.RoomRegisterRsp{}, n.CMDList, uint16(list.CMDID_List_IDRoomRegisterRsp), handleRoomRegisterRsp)
	g.MsgRegister(&client.JoinRoomReq{}, n.CMDClient, uint16(client.CMDID_Client_IDJoinRoomReq), handleJoinReq)
	g.MsgRegister(&client.RoomActionReq{}, n.CMDClient, uint16(client.CMDID_Client_IDRoomActionReq), handleUserActionReq)
	g.MsgRegister(&client.ExitRoomReq{}, n.CMDClient, uint16(client.CMDID_Client_IDExitRoomReq), handleExitReq)
	g.MsgRegister(&property.ModifyPropertyRsp{}, n.CMDProperty, uint16(property.CMDID_Property_IDModifyPropertyRsp), handleModifyPropertyRsp)
	g.EventRegister(g.ConnectSuccess, connectSuccess)
	g.EventRegister(g.Disconnect, disconnect)
	g.EventRegister(g.ConfigChangeNotify, configChangeNotify)

	g.Skeleton.LoopFunc(1*time.Second, checkMatchTable, timer.LoopForever)
}

func connectSuccess(args []interface{}) {
	log.Info("连接", "来了老弟,参数数量=%d", len(args))
}

func disconnect(args []interface{}) {
	log.Info("连接", "告辞中,参数数量=%d", len(args))
}

func configChangeNotify(args []interface{}) {
	tableAppID := apollo.GetConfigAsInt64("桌子服务AppID", 0)
	if tableAppID != 0 {
		g.Skeleton.AfterFunc(3*time.Second, checkApplyTable)
	}

	log.Info("配置", "真的收到了配置消息=%d,%d", len(args), tableAppID)
}

func handleApplyRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*tCMD.ApplyRsp)
	srcApp := args[n.OtherIndex].(n.BaseAgentInfo)

	tables = append(tables, m.GetTableIds()...)
	log.Debug("", "收到桌子,ApplyCount=%d,AttAppid=%d,len=%d", m.GetApplyCount(), srcApp.AppID, len(tables))

	var req list.RoomRegisterReq
	req.RoomInfo = new(types.RoomInfo)
	req.RoomInfo.RoomInfo = new(types.BaseAppInfo)
	req.RoomInfo.RoomInfo.Name = proto.String(conf.AppInfo.AppName)
	req.RoomInfo.RoomInfo.Type = proto.Uint32(conf.AppInfo.AppType)
	req.RoomInfo.RoomInfo.Id = proto.Uint32(conf.AppInfo.AppID)
	g.SendData2App(n.AppList, n.Send2AnyOne, n.CMDList, uint32(list.CMDID_List_IDRoomRegisterReq), &req)
}

func handleWriteGameScore(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*tCMD.WriteGameScore)
	srcApp := args[n.OtherIndex].(n.BaseAgentInfo)

	log.Debug("", "收到写分,tableId=%d,AttAppid=%d,len=%d", m.GetTableId(), srcApp.AppID, len(tables))

	var req property.ModifyPropertyReq
	p := new(types.PropItem)
	p.PropId = (*types.PropType)(proto.Int32(int32(types.PropType_Score)))
	p.PropCount = proto.Int64(100)
	req.UserProps = append(req.UserProps, p)
	g.SendData2App(n.AppProperty, n.Send2AnyOne, n.CMDProperty, uint32(property.CMDID_Property_IDModifyPropertyReq), &req)
}

func handleGameOver(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*tCMD.GameOver)
	srcApp := args[n.OtherIndex].(n.BaseAgentInfo)

	log.Debug("", "收到结束,table=%d,AttAppid=%d,len=%d", m.GetTableId(), srcApp.AppID, len(tables))

	for _, pl := range players {
		if pl.TableId != m.GetTableId() {
			continue
		}
		pl.State = player.StandingInRoom
	}
}

func handleRoomRegisterRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	//m := (b.MyMessage).(*list.RoomRegisterRsp)
	//srcApp := args[n.OtherIndex].(n.BaseAgentInfo)

	//tables = append(tables, m.GetTableIds()...)
	//log.Debug("", "tables=%v", tables)
	log.Debug("", "注册返回,AttAppid=%d,len=%d", b.AgentInfo.AppID, len(tables))
}

func handleJoinReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	//m := (b.MyMessage).(*client.JoinRoomReq)
	srcData := args[n.OtherIndex].(*gate.TransferDataReq)

	log.Debug("", "进入房间,userId = %v,Gateconnid=%v,appID =%v",
		srcData.GetUserId(), srcData.GetGateconnid(), util.GetLUint32FromUint64(srcData.GetGateconnid()))

	msgRespond := func(errCode int32) {
		var rsp client.JoinRoomRsp
		rsp.AppId = proto.Uint32(conf.AppInfo.AppID)
		rsp.ErrInfo = new(types.ErrorInfo)
		rsp.ErrInfo.Code = proto.Int32(errCode)
		rspBm := n.BaseMessage{MyMessage: &rsp, TraceId: ""}
		rspBm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDClient), SubCmdID: uint16(client.CMDID_Client_IDJoinRoomRsp)}
		g.SendMessage2Client(rspBm, srcData.GetGateconnid(), 0)
	}

	userID := srcData.GetUserId()
	if _, ok := players[userID]; ok {
		msgRespond(1)
		return
	}

	pl := player.NewPlayer(userID)
	pl.GateConnId = util.MakeUint64FromUint32(b.AgentInfo.AppType, b.AgentInfo.AppID)
	players[userID] = pl

	msgRespond(0)
}

func handleUserActionReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*client.RoomActionReq)
	srcData := args[n.OtherIndex].(*gate.TransferDataReq)

	msgRespond := func(errCode int32) {
		var rsp client.RoomActionRsp
		rsp.Action = (*client.ActionType)(proto.Int32(int32(m.GetAction())))
		rsp.ErrInfo = new(types.ErrorInfo)
		rsp.ErrInfo.Code = proto.Int32(errCode)
		rspBm := n.BaseMessage{MyMessage: &rsp, TraceId: b.TraceId}
		rspBm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDClient), SubCmdID: uint16(client.CMDID_Client_IDRoomActionRsp)}
		g.SendMessage2Client(rspBm, srcData.GetGateconnid(), 0)
	}

	userID := srcData.GetUserId()
	if _, ok := players[userID]; !ok {
		msgRespond(1)
		return
	}

	if m.GetAction() == client.ActionType_Ready {
		players[userID].State = player.HandsUpState
	}

	msgRespond(0)
}

func handleExitReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	//m := (b.MyMessage).(*client.ExitRoomReq)
	srcData := args[n.OtherIndex].(*gate.TransferDataReq)

	msgRespond := func(errCode int32) {
		var rsp client.ExitRoomRsp
		rsp.ErrInfo = new(types.ErrorInfo)
		rsp.ErrInfo.Code = proto.Int32(errCode)
		rspBm := n.BaseMessage{MyMessage: &rsp, TraceId: b.TraceId}
		rspBm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDClient), SubCmdID: uint16(client.CMDID_Client_IDRoomActionRsp)}
		g.SendMessage2Client(rspBm, srcData.GetGateconnid(), 0)
	}

	userID := srcData.GetUserId()
	if _, ok := players[userID]; !ok {
		msgRespond(1)
		return
	}

	msgRespond(0)
}

func handleModifyPropertyRsp(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*property.ModifyPropertyRsp)

	log.Debug("", "修改财富返回,userId=%v, opType=%v", m.GetUserId(), m.GetOpType())
}

func checkApplyTable() {
	if len(tables) == 0 {
		var req tCMD.ApplyReq
		req.ApplyCount = proto.Uint32(100)
		tableAppID := apollo.GetConfigAsInt64("桌子服务AppID", 2000)
		g.SendData2App(n.AppTable, uint32(tableAppID), n.CMDTable, uint32(tCMD.CMDID_Table_IDApplyReq), &req)
		g.Skeleton.AfterFunc(3*time.Second, checkApplyTable)
	}
}

func checkMatchTable() {
	var matchPlayers []*player.Player
	for _, pl := range players {
		if pl.State == player.HandsUpState {
			matchPlayers = append(matchPlayers, pl)
		}
	}

	seatCount := apollo.GetConfigAsInt64("座位数量", 3)
	for len(matchPlayers) >= int(seatCount) {
		tablePlayers := matchPlayers[:seatCount]
		matchPlayers = matchPlayers[seatCount:]

		log.Debug("", "len=%v", len(matchPlayers))

		tableID := tables[len(tables)-1]
		tableAppID := apollo.GetConfigAsInt64("桌子服务AppID", 2000)
		var req tCMD.MatchTableReq
		req.TableId = proto.Uint64(tableID)
		for i := 0; i < int(seatCount); i++ {
			log.Debug("", "设置用户,userId=%v", tablePlayers[i].UserID)
			req.Players = append(req.Players, tablePlayers[i].UserID)
			tablePlayers[i].State = player.PlayingState
			tablePlayers[i].TableId = tableID
			tablePlayers[i].SeatId = uint32(i)
			setPlayerToTable(tablePlayers[i], uint32(tableAppID))
		}

		go func() {
			time.Sleep(500)
			g.SendData2App(n.AppTable, uint32(tableAppID), n.CMDTable, uint32(tCMD.CMDID_Table_IDMatchTableReq), &req)
		}()
	}
}

func setPlayerToTable(pl *player.Player, tableAppID uint32) {
	var req tCMD.SetPlayerToTableReq
	req.UserId = proto.Uint64(pl.UserID)
	req.TableId = proto.Uint64(pl.TableId)
	req.SeatId = proto.Uint32(pl.SeatId)
	req.Gateconnid = proto.Uint64(pl.GateConnId)
	g.SendData2App(n.AppTable, tableAppID, n.CMDTable, uint32(tCMD.CMDID_Table_IDSetPlayerToTableReq), &req)

	var state rCMD.UserStateChange
	state.UserId = proto.Uint64(pl.UserID)
	state.TableServiceId = proto.Uint32(tableAppID)
	state.TableId = proto.Uint64(pl.TableId)
	state.SeatId = proto.Uint32(pl.SeatId)
	rspBm := n.BaseMessage{MyMessage: &state, TraceId: ""}
	rspBm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDRoom), SubCmdID: uint16(rCMD.CMDID_Room_IDUserStateChange)}
	g.SendMessage2Client(rspBm, pl.GateConnId, 0)
}
