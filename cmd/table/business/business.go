package business

import (
	"github.com/golang/protobuf/proto"
	"xlddz/api/gate"
	tCMD "xlddz/api/table"
	"xlddz/cmd/table/business/player"
	"xlddz/cmd/table/business/table"
	"xlddz/cmd/table/business/table/ddz"
	"xlddz/pkg/conf"
	"xlddz/pkg/conf/apollo"
	g "xlddz/pkg/gate"
	"xlddz/pkg/log"
	n "xlddz/pkg/network"
)

var (
	freeTables map[uint64]*table.Table = make(map[uint64]*table.Table)
	usedTables map[uint64]*table.Table = make(map[uint64]*table.Table)
)

func init() {
	g.MsgRegister(&tCMD.ApplyReq{}, n.CMDTable, uint16(tCMD.CMDID_Table_IDApplyReq), handleApplyReq)
	g.MsgRegister(&tCMD.ReleaseReq{}, n.CMDTable, uint16(tCMD.CMDID_Table_IDReleaseReq), handleReleaseReq)
	g.MsgRegister(&tCMD.SetPlayerToTableReq{}, n.CMDTable, uint16(tCMD.CMDID_Table_IDSetPlayerToTableReq), handleSetPlayerToTableReq)
	g.MsgRegister(&tCMD.MatchTableReq{}, n.CMDTable, uint16(tCMD.CMDID_Table_IDMatchTableReq), handleMatchTableReq)
	g.MsgRegister(&tCMD.GameMessage{}, n.CMDTable, uint16(tCMD.CMDID_Table_IDGameMessage), handleGameMessage)
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

	g.SendData2App(srcApp.AppType, srcApp.AppID, n.CMDTable, uint32(tCMD.CMDID_Table_IDApplyRsp), &rsp)
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
	if m.GetTableId() != uint64(conf.AppInfo.AppID) {
		return
	}

	pl := getPlayer(m.GetUserId())
	if pl != nil {
		return
	}
	t := getTable(srcApp.AppID, m.GetTableId())
	if t == nil {
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

	log.Debug("", "收到释放,TableId=%d,len=%d,srcID=%d", m.GetTableId(), len(freeTables), srcApp.AppID)
}

func handleMatchTableReq(args []interface{}) {
	m := args[n.DataIndex].(n.BaseMessage).MyMessage.(*tCMD.MatchTableReq)
	srcApp := args[n.OtherIndex].(n.BaseAgentInfo)
	if m.GetTableId() != uint64(conf.AppInfo.AppID) {
		return
	}

	t := getTable(srcApp.AppID, m.GetTableId())
	if t == nil {
		return
	}

	t.Start()

	log.Debug("", "收到释放,TableId=%d,len=%d,srcID=%d", m.GetTableId(), len(freeTables), srcApp.AppID)
}

func handleGameMessage(args []interface{}) {
	m := args[n.DataIndex].(n.BaseMessage).MyMessage.(*tCMD.GameMessage)
	srcData := args[n.OtherIndex].(*gate.TransferDataReq)

	userID := srcData.GetUserId()
	pl := getPlayer(userID)
	if pl == nil {
		return
	}

	t := getTable(pl.SrcAppID, pl.TableID)
	if t == nil {
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
