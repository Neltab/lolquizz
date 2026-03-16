package room_test

import (
	"lolquizz/internal/domain/room"
)

func newTestRoom() (*room.Room, *room.Player) {
	host := room.NewPlayer("host", "Host", true)
	room := room.NewRoom("room-1", "code", host)
	return room, host
}
