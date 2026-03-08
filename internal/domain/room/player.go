package room

import "lolquizz/internal/domain/shared"

type Player struct {
	Id     shared.PlayerId
	Name   string
	IsHost bool
	Score  int
}

func NewPlayer(id shared.PlayerId, name string, isHost bool) *Player {
	return &Player{
		Id:     id,
		Name:   name,
		IsHost: isHost,
		Score:  0,
	}
}
