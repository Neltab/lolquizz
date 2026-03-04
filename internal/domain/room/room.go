package room

import (
	"lolquizz/internal/domain/shared"
	"time"
)

type RoomStatus int

const (
	StatusSetuping RoomStatus = iota
	StatusPlaying
	StatusReviewing
	StatusFinished
)

type Room struct {
	Id        shared.RoomId               `json:"id"`
	Code      string                      `json:"code"`
	Status    RoomStatus                  `json:"status"`
	HostId    shared.PlayerId             `json:"host_id"`
	Players   map[shared.PlayerId]*Player `json:"players"`
	Settings  Settings                    `json:"settings"`
	CreatedAt time.Time                   `json:"created_at"`
}

func NewRoom(id shared.RoomId, code string, host *Player) *Room {
	r := &Room{
		Id:        id,
		Code:      code,
		Status:    StatusSetuping,
		HostId:    host.Id,
		Players:   make(map[shared.PlayerId]*Player),
		Settings:  Settings{MaxPlayers: 10},
		CreatedAt: time.Now(),
	}
	r.Players[host.Id] = host
	return r
}

func (r *Room) Join(player *Player) error {
	if r.Status != StatusSetuping {
		return ErrGameAlreadyStarted
	}

	if len(r.Players) >= r.Settings.MaxPlayers {
		return ErrRoomFull
	}

	r.Players[player.Id] = player
	return nil
}

func (r *Room) Leave(playerId shared.PlayerId) error {
	if _, ok := r.Players[playerId]; !ok {
		return ErrPlayerNotFound
	}

	delete(r.Players, playerId)

	if playerId == r.HostId && len(r.Players) > 0 {
		for id := range r.Players {
			r.HostId = id
			break
		}
	}
	return nil
}

func (r *Room) StartGame(playerId shared.PlayerId) error {
	if !r.IsHost(playerId) {
		return ErrPlayerNotHost
	}

	if r.Status != StatusSetuping {
		return ErrGameAlreadyStarted
	}

	r.Status = StatusPlaying
	return nil
}

func (r *Room) IsHost(playerId shared.PlayerId) bool {
	return playerId == r.HostId
}
