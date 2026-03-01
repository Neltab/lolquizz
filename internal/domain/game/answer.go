package game

import (
	"time"

	"lolquizz/internal/domain/room"
)

type Answer struct {
	PlayerId     room.PlayerId
	Answer       string
	timeAnswered time.Time
}

func (a *Answer) Points() int {
	return 10
}
