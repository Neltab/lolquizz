package dto

import "lolquizz/internal/domain/room"

type Player struct {
	Name   string `json:"name"`
	Score  int    `json:"score"`
	IsHost bool   `json:"is_host"`
}

func FromPlayer(p *room.Player) *Player {
	return &Player{
		Name:   p.Name,
		Score:  p.Score,
		IsHost: p.IsHost,
	}
}
