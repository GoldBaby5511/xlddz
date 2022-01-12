package table

import (
	"github.com/golang/protobuf/proto"
	tCMD "mango/api/table"
	"mango/cmd/table/business/player"
	g "mango/pkg/gate"
	"mango/pkg/log"
	n "mango/pkg/network"
)

const InvalidSeadID = 0xFFFF

type GameSink interface {
	StartGame(f Frame)
	GameMessage(seatID, subCmdID uint32, data []byte)
}

type Frame interface {
	SendTableData(seatID uint32, bm n.BaseMessage)
	WriteGameScore()
	GameOver()
}

type Table struct {
	id        uint64
	HostAppID uint32
	gameSink  GameSink
	Players   map[uint32]*player.Player
}

func NewTable(id uint64, sink GameSink) *Table {
	t := new(Table)
	t.id = id
	t.HostAppID = 0
	t.gameSink = sink
	t.Players = make(map[uint32]*player.Player)
	return t
}

func (t *Table) SendTableData(seatID uint32, bm n.BaseMessage) {
	if seatID == InvalidSeadID {
		for _, pl := range t.Players {
			g.SendMessage2Client(bm, pl.GateConnID, 0)
		}
	} else {
		pl, ok := t.Players[seatID]
		if !ok {
			log.Warning("", "没找到,seatID=%d,id=%v,hostId=%v", seatID, t.id, t.HostAppID)
			return
		}
		g.SendMessage2Client(bm, pl.GateConnID, 0)
	}
}

func (t *Table) WriteGameScore() {
	var writeScore tCMD.WriteGameScore
	g.SendData2App(n.AppRoom, t.HostAppID, n.CMDTable, uint32(tCMD.CMDTable_IDWriteGameScore), &writeScore)
}

func (t *Table) GameOver() {
	t.Players = make(map[uint32]*player.Player)
	var over tCMD.GameOver
	over.TableId = proto.Uint64(t.id)
	g.SendData2App(n.AppRoom, t.HostAppID, n.CMDTable, uint32(tCMD.CMDTable_IDGameOver), &over)
}

func (t *Table) Reset() {
	t.HostAppID = 0
	t.Players = make(map[uint32]*player.Player)
}

func (t *Table) SetHostID(hostID uint32) {
	t.HostAppID = hostID
}

func (t *Table) GetHostID() uint32 {
	return t.HostAppID
}

func (t *Table) GeTableID() uint64 {
	return t.id
}

func (t *Table) SetPlayer(pl *player.Player) {
	if _, ok := t.Players[pl.SeatID]; ok {
		log.Warning("", "有人了,id=%v,userId=%v,seatId=%v", t.id, pl.UserID, pl.SeatID)
		return
	}
	t.Players[pl.SeatID] = pl
}

func (t *Table) Start() {
	t.gameSink.StartGame(t)
}

func (t *Table) GameMessage(seatID, subCmdID uint32, data []byte) {
	t.gameSink.GameMessage(seatID, subCmdID, data)
}
