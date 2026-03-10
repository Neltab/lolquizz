package room

type PlayerJoinedEvent struct {
	RoomId RoomId
	Player *Player
	Room   *Room
}

func (e PlayerJoinedEvent) EventName() string { return "player_joined" }

type PlayerLeftEvent struct {
	RoomId RoomId
	Player *Player
	Room   *Room
}

func (e PlayerLeftEvent) EventName() string { return "player_left" }

type SettingsUpdatedEvent struct {
	RoomId   RoomId
	Settings *Settings
	Room     *Room
}

func (e SettingsUpdatedEvent) EventName() string { return "settings_updated" }
