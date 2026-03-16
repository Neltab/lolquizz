package room

type RoomCreatedEvent struct {
	RoomId RoomId
	Room   *Room
}

func (e *RoomCreatedEvent) EventName() string { return "room_created" }

type RoomDeletedEvent struct {
	RoomId RoomId
}

func (e *RoomDeletedEvent) EventName() string { return "room_deleted" }

type PlayerJoinedEvent struct {
	RoomId RoomId
	Player *Player
	Room   *Room
}

func (e *PlayerJoinedEvent) EventName() string { return "player_joined" }

type PlayerLeftEvent struct {
	RoomId RoomId
	Player *Player
	Room   *Room
}

func (e *PlayerLeftEvent) EventName() string { return "player_left" }

type SettingsUpdatedEvent struct {
	RoomId   RoomId
	Settings *Settings
	Room     *Room
}

func (e *SettingsUpdatedEvent) EventName() string { return "settings_updated" }
