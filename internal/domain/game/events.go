package game

import (
	"time"

	"lolquizz/internal/domain/room"
)

type QuestionStartedEvent struct {
	RoomId         RoomId
	QuestionNumber int
	QuestionText   string
	Duration       time.Duration
}

func (e *QuestionStartedEvent) EventName() string {
	return "round_started"
}

type AnswerSubmittedEvent struct {
	RoomId   RoomId
	PlayerId PlayerId
}

func (e *AnswerSubmittedEvent) EventName() string {
	return "answer_submitted"
}

type AnswerJudgedEvent struct {
	RoomId   RoomId
	PlayerId PlayerId
	Correct  bool
}

func (e *AnswerJudgedEvent) EventName() string {
	return "answer_judged"
}

type TimerExpiredEvent struct {
	RoomId RoomId
}

func (e *TimerExpiredEvent) EventName() string {
	return "timer_expired"
}

type GameFinishedEvent struct {
	RoomId RoomId
}

func (e *GameFinishedEvent) EventName() string {
	return "game_finished"
}

type PlayerJoinedEvent struct {
	RoomID  RoomId         `json:"room_id"`
	Player  *room.Player   `json:"player"`
	Players []*room.Player `json:"players"`
}

func (e PlayerJoinedEvent) EventName() string { return "player_joined" }
