package game

import (
	"time"

	"lolquizz/internal/domain/room"
	"lolquizz/internal/domain/shared"
)

type QuestionStartedEvent struct {
	RoomId         shared.RoomId
	QuestionNumber int
	QuestionText   string
	Duration       time.Duration
}

func (e *QuestionStartedEvent) EventName() string {
	return "round_started"
}

type AnswerSubmittedEvent struct {
	RoomId   shared.RoomId
	PlayerId shared.PlayerId
}

func (e *AnswerSubmittedEvent) EventName() string {
	return "answer_submitted"
}

type AnswerJudgedEvent struct {
	RoomId   shared.RoomId
	PlayerId shared.PlayerId
	Correct  bool
}

func (e *AnswerJudgedEvent) EventName() string {
	return "answer_judged"
}

type TimerExpiredEvent struct {
	RoomId shared.RoomId
}

func (e *TimerExpiredEvent) EventName() string {
	return "timer_expired"
}

type GameFinishedEvent struct {
	RoomId shared.RoomId
}

func (e *GameFinishedEvent) EventName() string {
	return "game_finished"
}

type PlayerJoinedEvent struct {
	RoomID  shared.RoomId  `json:"room_id"`
	Player  *room.Player   `json:"player"`
	Players []*room.Player `json:"players"`
}

func (e PlayerJoinedEvent) EventName() string { return "player_joined" }
