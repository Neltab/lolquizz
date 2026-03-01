package room

type Player struct {
	Id    PlayerId
	Name  string
	Score int
}

func NewPlayer(id PlayerId, name string) *Player {
	return &Player{
		Id:    id,
		Name:  name,
		Score: 0,
	}
}
