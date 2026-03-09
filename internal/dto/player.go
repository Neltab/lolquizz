package dto

import "lolquizz/internal/domain/room"

type PlayerDTO struct {
	Name   string `json:"name"`
	Score  int    `json:"score"`
	IsHost bool   `json:"is_host"`
}

func FromPlayer(p *room.Player) *PlayerDTO {
	return &PlayerDTO{
		Name:   p.Name,
		Score:  p.Score,
		IsHost: p.IsHost,
	}
}
