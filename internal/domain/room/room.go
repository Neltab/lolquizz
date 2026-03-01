package room

import "time"

type RoomId string
type PlayerId string
type RoomStatus int

const (
	StatusSetuping RoomStatus = iota
	StatusPlaying
	StatusReviewing
	StatusFinished
)

type Room struct {
	Id        RoomId
	Code      string
	Status    RoomStatus
	HostId    PlayerId
	Players   map[PlayerId]*Player
	Settings  Settings
	CreatedAt time.Time
}

func NewRoom(id RoomId, code string, host *Player) *Room {
	r := &Room{
		Id:        id,
		Code:      code,
		Status:    StatusSetuping,
		HostId:    host.Id,
		Players:   make(map[PlayerId]*Player),
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

func (r *Room) Leave(playerId PlayerId) error {
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

func (r *Room) StartGame(playerId PlayerId) error {
	if !r.IsHost(playerId) {
		return ErrPlayerNotHost
	}

	if r.Status != StatusSetuping {
		return ErrGameAlreadyStarted
	}

	r.Status = StatusPlaying
	return nil
}

func (r *Room) IsHost(playerId PlayerId) bool {
	return playerId == r.HostId
}
