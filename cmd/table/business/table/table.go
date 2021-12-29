package table

import (
	"mango/cmd/table/business/player"
	g "mango/pkg/gate"
	n "mango/pkg/network"
)

const InvalidSeadID = 0xFFFF

type GameSink interface {
	StartGame(f Frame)
	GameMessage(seatID, subCmdID uint32, data []byte)
}

type Frame interface {
	SendTableData(seatID uint32, bm n.BaseMessage)
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
			return
		}
		g.SendMessage2Client(bm, pl.GateConnID, 0)
	}
}

func (t *Table) Reset() {
	t.HostAppID = 0
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
	t.Players[pl.SeatID] = pl
}

func (t *Table) Start() {
	t.gameSink.StartGame(t)
}

func (t *Table) GameMessage(seatID, subCmdID uint32, data []byte) {
	t.gameSink.GameMessage(seatID, subCmdID, data)
}
