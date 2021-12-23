package player

const (
	NilState     uint32 = 0
	HandsUpState uint32 = 1
	PlayingState uint32 = 2
	SitdownState uint32 = 3
)

type Player struct {
	UserID     uint64
	State      uint32
	TableID    uint64
	SeatID     uint32
	GateConnID uint64
	SrcAppID   uint32
}

func NewPlayer() *Player {
	p := new(Player)
	p.UserID = 0
	p.State = NilState
	return p
}
