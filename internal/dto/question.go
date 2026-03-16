package dto

import (
	"lolquizz/internal/domain/game"
	"time"
)

type Question struct {
	Text     string        `json:"text"`
	Duration time.Duration `json:"duration"`
}

func FromQuestion(q *game.Question) *Question {
	return &Question{
		Text:     q.Text,
		Duration: q.Duration,
	}
}
