package room

type Player struct {
	Id     PlayerId
	Name   string
	IsHost bool
	Score  int
}

func NewPlayer(id PlayerId, name string, isHost bool) *Player {
	return &Player{
		Id:     id,
		Name:   name,
		IsHost: isHost,
		Score:  0,
	}
}
