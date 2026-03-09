package game

import (
	"time"
)

type Answer struct {
	PlayerId     PlayerId
	Answer       string
	timeAnswered time.Time
}

func (a *Answer) Points() int {
	return 10
}
