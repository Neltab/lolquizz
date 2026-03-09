package room

type PlayerJoinedEvent struct {
	RoomId RoomId
	Player *Player
}

func (e PlayerJoinedEvent) EventName() string { return "player_joined" }

type PlayerLeftEvent struct {
	RoomId  RoomId
	NewHost PlayerId
}

func (e PlayerLeftEvent) EventName() string { return "player_left" }

type SettingsUpdatedEvent struct {
	RoomId   RoomId
	Settings Settings
}

func (e SettingsUpdatedEvent) EventName() string { return "settings_updated" }
