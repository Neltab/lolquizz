package room

import "lolquizz/internal/domain/shared"

type Player struct {
	Id    shared.PlayerId
	Name  string
	Score int
}

func NewPlayer(id shared.PlayerId, name string) *Player {
	return &Player{
		Id:    id,
		Name:  name,
		Score: 0,
	}
}
