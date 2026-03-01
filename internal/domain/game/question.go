package game

import (
	"time"

	"lolquizz/internal/domain/room"
)

type Difficulty int

const (
	DifficultyEasy Difficulty = iota
	DifficultyMedium
	DifficultyHard
)

type Question struct {
	Text       string
	Image      string
	Answers    map[room.PlayerId]*Answer
	Difficulty Difficulty
	Duration   time.Duration
}

func NewQuestion(text string, image string, difficulty Difficulty, duration time.Duration) *Question {
	return &Question{
		Text:       text,
		Image:      image,
		Answers:    make(map[room.PlayerId]*Answer),
		Difficulty: difficulty,
		Duration:   duration,
	}
}

func (q *Question) Answer(playerId room.PlayerId, answer string, answerTime time.Time) {
	q.Answers[playerId] = &Answer{
		PlayerId:     playerId,
		Answer:       answer,
		timeAnswered: answerTime,
	}
}

func (q *Question) GetAnswer(playerId room.PlayerId) *Answer {
	return q.Answers[playerId]
}
