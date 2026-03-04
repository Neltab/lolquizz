package room

import "lolquizz/internal/domain/shared"

type PlayerJoinedEvent struct {
	RoomId shared.RoomId
	Player *Player
}

func (e PlayerJoinedEvent) EventName() string { return "player_joined" }

type PlayerLeftEvent struct {
	RoomId  shared.RoomId
	NewHost shared.PlayerId
}

func (e PlayerLeftEvent) EventName() string { return "player_left" }

type SettingsUpdatedEvent struct {
	RoomId   shared.RoomId
	Settings Settings
}

func (e SettingsUpdatedEvent) EventName() string { return "settings_updated" }
