package business

import (
	"github.com/golang/protobuf/proto"
	"mango/api/gateway"
	tCMD "mango/api/table"
	"mango/cmd/table/business/player"
	"mango/cmd/table/business/table"
	"mango/cmd/table/business/table/ddz"
	"mango/pkg/conf/apollo"
	g "mango/pkg/gate"
	"mango/pkg/log"
	n "mango/pkg/network"
)

var (
	freeTables map[uint64]*table.Table = make(map[uint64]*table.Table)
	usedTables map[uint64]*table.Table = make(map[uint64]*table.Table)
)

func init() {
	g.MsgRegister(&tCMD.ApplyReq{}, n.CMDTable, uint16(tCMD.CMDTable_IDApplyReq), handleApplyReq)
	g.MsgRegister(&tCMD.ReleaseReq{}, n.CMDTable, uint16(tCMD.CMDTable_IDReleaseReq), handleReleaseReq)
	g.MsgRegister(&tCMD.SetPlayerToTableReq{}, n.CMDTable, uint16(tCMD.CMDTable_IDSetPlayerToTableReq), handleSetPlayerToTableReq)
	g.MsgRegister(&tCMD.MatchTableReq{}, n.CMDTable, uint16(tCMD.CMDTable_IDMatchTableReq), handleMatchTableReq)
	g.MsgRegister(&tCMD.GameMessage{}, n.CMDTable, uint16(tCMD.CMDTable_IDGameMessage), handleGameMessage)
	g.EventRegister(g.ConnectSuccess, connectSuccess)
	g.EventRegister(g.Disconnect, disconnect)
	g.EventRegister(g.ConfigChangeNotify, configChangeNotify)
}

func connectSuccess(args []interface{}) {
	log.Info("连接", "来了老弟,参数数量=%d", len(args))
}

func disconnect(args []interface{}) {
	log.Info("连接", "告辞中,参数数量=%d", len(args))
}

func configChangeNotify(args []interface{}) {
	tableCount := apollo.GetConfigAsInt64("桌子数量", 2000)
	if len(freeTables) == 0 && len(usedTables) == 0 && tableCount != 0 {
		for i := 0; i < int(tableCount); i++ {
			freeTables[uint64(i)] = table.NewTable(uint64(i), new(ddz.Sink))
		}
	}

	log.Info("配置", "真的收到了配置消息=%d,%d", len(args), tableCount)
}

func handleApplyReq(args []interface{}) {
	b := args[n.DataIndex].(n.BaseMessage)
	m := (b.MyMessage).(*tCMD.ApplyReq)
	srcApp := args[n.OtherIndex].(n.BaseAgentInfo)

	log.Debug("", "收到申请,ApplyCount=%d,len=%d", m.GetApplyCount(), len(freeTables))
	if len(freeTables) < int(m.GetApplyCount()) {
		return
	}

	var rsp tCMD.ApplyRsp
	rsp.ApplyCount = proto.Uint32(m.GetApplyCount())
	for k, v := range freeTables {
		rsp.TableIds = append(rsp.TableIds, v.GeTableID())
		v.SetHostID(srcApp.AppID)
		delete(freeTables, k)
		usedTables[k] = v
		if len(rsp.GetTableIds()) == int(m.GetApplyCount()) {
			break
		}
	}

	g.SendData2App(srcApp.AppType, srcApp.AppID, n.CMDTable, uint32(tCMD.CMDTable_IDApplyRsp), &rsp)
}

func handleReleaseReq(args []interface{}) {
	m := args[n.DataIndex].(n.BaseMessage).MyMessage.(*tCMD.ReleaseReq)
	srcApp := args[n.OtherIndex].(n.BaseAgentInfo)

	log.Debug("", "收到释放,ApplyCount=%d,len=%d,srcID=%d", m.GetReleaseCount(), len(freeTables), srcApp.AppID)
	for _, tID := range m.GetTableIds() {
		t := getTable(srcApp.AppID, tID)
		if t == nil {
			continue
		}
		t.Reset()
		delete(usedTables, tID)
		freeTables[tID] = t
	}
}

func handleSetPlayerToTableReq(args []interface{}) {
	m := args[n.DataIndex].(n.BaseMessage).MyMessage.(*tCMD.SetPlayerToTableReq)
	srcApp := args[n.OtherIndex].(n.BaseAgentInfo)
	if _, ok := usedTables[m.GetTableId()]; !ok {
		log.Warning("", "没找到桌子啊,tableId=%v", m.GetTableId())
		return
	}
	t := getTable(srcApp.AppID, m.GetTableId())
	if t == nil {
		log.Warning("", "这桌子不是你的啊,tableId=%v,host=%v,srcId=%v", m.GetTableId(), usedTables[m.GetTableId()].HostAppID, srcApp.AppID)
		return
	}

	pl := getPlayer(m.GetUserId())
	if pl != nil {
		return
	}

	pl = player.NewPlayer()
	pl.UserID = m.GetUserId()
	pl.TableID = t.GeTableID()
	pl.SrcAppID = srcApp.AppID
	pl.SeatID = m.GetSeatId()
	pl.GateConnID = m.GetGateconnid()
	pl.State = player.SitdownState
	t.SetPlayer(pl)

	log.Debug("", "收到入座,UserID=%v,TableId=%d,len=%d,srcID=%d", pl.UserID, m.GetTableId(), len(freeTables), srcApp.AppID)
}

func handleMatchTableReq(args []interface{}) {
	m := args[n.DataIndex].(n.BaseMessage).MyMessage.(*tCMD.MatchTableReq)
	srcApp := args[n.OtherIndex].(n.BaseAgentInfo)

	t := getTable(srcApp.AppID, m.GetTableId())
	if t == nil {
		return
	}

	log.Debug("", "收到配桌,TableId=%d,len=%d,srcID=%d", m.GetTableId(), len(freeTables), srcApp.AppID)
	t.Start()
}

func handleGameMessage(args []interface{}) {
	m := args[n.DataIndex].(n.BaseMessage).MyMessage.(*tCMD.GameMessage)
	srcData := args[n.OtherIndex].(*gateway.TransferDataReq)

	userID := srcData.GetUserId()
	pl := getPlayer(userID)
	if pl == nil {
		log.Warning("", "游戏消息,没找到用户啊,userID=%v", userID)
		return
	}

	t := getTable(pl.SrcAppID, pl.TableID)
	if t == nil {
		log.Warning("", "游戏消息,没找到桌子啊,userID=%v,SrcAppID=%v,TableID=%v", userID, pl.SrcAppID, pl.TableID)
		return
	}
	t.GameMessage(pl.SeatID, m.GetSubCmdid(), m.GetData())
}

func getTable(srcAppID uint32, tableID uint64) *table.Table {
	for _, t := range usedTables {
		if t.GeTableID() == tableID && t.GetHostID() == srcAppID {
			return t
		}
	}
	return nil
}

func getPlayer(userID uint64) *player.Player {
	for _, t := range usedTables {
		for _, pl := range t.Players {
			if pl.UserID == userID {
				return pl
			}
		}
	}
	return nil
}
