package game

import (
	"lolquizz/internal/domain/shared"
	"time"
)

type Answer struct {
	PlayerId     shared.PlayerId
	Answer       string
	timeAnswered time.Time
}

func (a *Answer) Points() int {
	return 10
}
