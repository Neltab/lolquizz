package dto

import "lolquizz/internal/domain/room"

type Room struct {
	Players  []*Player       `json:"players"`
	Code     string          `json:"code"`
	Status   room.RoomStatus `json:"status"`
	Settings *room.Settings  `json:"settings"`
}

func FromRoom(r *room.Room) *Room {
	dto := &Room{
		Players:  make([]*Player, 0, len(r.Players)),
		Code:     r.Code,
		Status:   r.Status,
		Settings: r.Settings,
	}

	for _, p := range r.Players {
		dto.Players = append(dto.Players, FromPlayer(p))
	}

	return dto
}
