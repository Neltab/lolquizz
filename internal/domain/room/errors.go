package room

import "errors"

var (
	ErrGameAlreadyStarted = errors.New("Game already started")
	ErrRoomFull           = errors.New("Room is full")
	ErrPlayerNotFound     = errors.New("Player not found")
	ErrPlayerNotHost      = errors.New("Player is not host")
)
