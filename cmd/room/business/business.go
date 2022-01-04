package business

import (
	"github.com/golang/protobuf/proto"
	"mango/api/gate"
	rCMD "mango/api/room"
	tCMD "mango/api/table"
	"mango/api/types"
	"mango/cmd/room/business/player"
	"mango/pkg/conf/apollo"
	g "mango/pkg/gate"
	"mango/pkg/log"
	n "mango/pkg/network"
	"mango/pkg/timer"
	"time"
)

var (
	tables  []uint64
	players map[uint64]*player.Player = make(map[uint64]*player.Player)
)

func init() {
	g.MsgRegister(&tCMD.ApplyRsp{}, n.CMDTable, uint16(tCMD.CMDID_Table_IDApplyRsp), handleApplyRsp)
	g.MsgRegister(&rCMD.JoinReq{}, n.CMDRoom, uint16(rCMD.CMDID_Room_IDJoinReq), handleJoinReq)
	g.MsgRegister(&rCMD.UserActionReq{}, n.CMDRoom, uint16(rCMD.CMDID_Room_IDUserActionReq), handleUserActionReq)
	g.MsgRegister(&rCMD.ExitReq{}, n.CMDRoom, uint16(rCMD.CMDID_Room_IDExitReq), handleExitReq)
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
	//log.Debug("", "tables=%v", tables)
	log.Debug("", "收到桌子,ApplyCount=%d,AttAppid=%d,len=%d", m.GetApplyCount(), srcApp.AppID, len(tables))
}

func handleJoinReq(args []interface{}) {
	//b := args[n.DataIndex].(n.BaseMessage)
	//m := (b.MyMessage).(*rCMD.JoinReq)
	srcData := args[n.OtherIndex].(*gate.TransferDataReq)

	msgRespond := func(errCode int32) {
		var rsp rCMD.JoinRsp
		rsp.ErrInfo = new(types.ErrorInfo)
		rsp.ErrInfo.Code = proto.Int32(errCode)
		rspBm := n.BaseMessage{MyMessage: &rsp, TraceId: ""}
		rspBm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDRoom), SubCmdID: uint16(rCMD.CMDID_Room_IDJoinRsp)}
		g.SendMessage2Client(rspBm, srcData.GetGateconnid(), 0)
	}

	userID := srcData.GetUserId()
	if _, ok := players[userID]; ok {
		msgRespond(1)
		return
	}

	players[userID] = player.NewPlayer(userID)

	msgRespond(0)
}

func handleUserActionReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*rCMD.UserActionReq)
	srcData := args[n.OtherIndex].(*gate.TransferDataReq)

	msgRespond := func(errCode int32) {
		var rsp rCMD.UserActionRsp
		rsp.Action = (*rCMD.ActionType)(proto.Int32(int32(m.GetAction())))
		rsp.ErrInfo = new(types.ErrorInfo)
		rsp.ErrInfo.Code = proto.Int32(errCode)
		rspBm := n.BaseMessage{MyMessage: &rsp, TraceId: b.TraceId}
		rspBm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDRoom), SubCmdID: uint16(rCMD.CMDID_Room_IDUserActionRsp)}
		g.SendMessage2Client(rspBm, srcData.GetGateconnid(), 0)
	}

	userID := srcData.GetUserId()
	if _, ok := players[userID]; ok {
		msgRespond(1)
		return
	}

	if m.GetAction() == rCMD.ActionType_Ready {
		players[userID].State = player.HandsUpState
	}

	msgRespond(0)
}

func handleExitReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	//m := (b.MyMessage).(*rCMD.ExitReq)
	srcData := args[n.OtherIndex].(*gate.TransferDataReq)

	msgRespond := func(errCode int32) {
		var rsp rCMD.ExitRsp
		rsp.ErrInfo = new(types.ErrorInfo)
		rsp.ErrInfo.Code = proto.Int32(errCode)
		rspBm := n.BaseMessage{MyMessage: &rsp, TraceId: b.TraceId}
		rspBm.Cmd = n.TCPCommand{MainCmdID: uint16(n.CMDRoom), SubCmdID: uint16(rCMD.CMDID_Room_IDUserActionRsp)}
		g.SendMessage2Client(rspBm, srcData.GetGateconnid(), 0)
	}

	userID := srcData.GetUserId()
	if _, ok := players[userID]; ok {
		msgRespond(1)
		return
	}

	msgRespond(0)
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

		tableID := tables[len(tables)-1]
		tableAppID := apollo.GetConfigAsInt64("桌子服务AppID", 2000)
		var req tCMD.MatchTableReq
		req.TableId = proto.Uint64(tableID)
		for i := 0; i < int(seatCount); i++ {
			req.Players = append(req.Players, tablePlayers[i].UserID)
			tablePlayers[i].State = player.PlayingState
			setPlayerToTable(tableID, tablePlayers[i].UserID, uint32(tableAppID))
		}

		g.SendData2App(n.AppTable, uint32(tableAppID), n.CMDTable, uint32(tCMD.CMDID_Table_IDMatchTableReq), &req)
	}
}

func setPlayerToTable(tableID, userID uint64, tableAppID uint32) {
	var req tCMD.SetPlayerToTableReq
	req.UserId = proto.Uint64(userID)
	req.TableId = proto.Uint64(tableID)
	g.SendData2App(n.AppTable, tableAppID, n.CMDTable, uint32(tCMD.CMDID_Table_IDSetPlayerToTableReq), &req)
}
