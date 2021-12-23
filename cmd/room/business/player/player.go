package player

const (
	NilState     uint32 = 0
	HandsUpState uint32 = 1
	PlayingState uint32 = 2
)

type Player struct {
	UserID uint64
	State  uint32
}

func NewPlayer(userID uint64) *Player {
	p := new(Player)
	p.UserID = userID
	p.State = NilState
	return p
}
