package ws

import (
	"encoding/json"
	"lolquizz/internal/domain/room"
	"lolquizz/internal/dto"
)

type IncomingMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type OutgoingMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

const (
	// Client -> Server
	MsgJoinRoom       = "join_room"
	MsgLeaveRoom      = "leave_room"
	MsgUpdateSettings = "update_settings"
	MsgStartGame      = "start_game"
	MsgSubmitAnswer   = "submit_answer"
	MsgJudgeAnswer    = "judge_answer"
	MsgNextRound      = "next_round"

	// Server -> Client
	MsgPlayerJoined    = "player_joined"
	MsgPlayerLeft      = "player_left"
	MsgSettingsUpdated = "settings_updated"
	MsgGameStarted     = "game_started"
	MsgTimerExpired    = "timer_expired"
	MsgAnswerJudged    = "answer_judged"
	MsgGameFinished    = "game_finished"
	MsgError           = "error"
)

type PlayerJoinedEvent struct {
	RoomId room.RoomId
	Player dto.PlayerDTO
}

func (e PlayerJoinedEvent) EventName() string { return "player_joined" }

type PlayerLeftEvent struct {
	RoomId  room.RoomId
	NewHost dto.PlayerDTO
}

func (e PlayerLeftEvent) EventName() string { return "player_left" }

type SettingsUpdatedEvent struct {
	RoomId   room.RoomId
	Settings room.Settings
}

func (e SettingsUpdatedEvent) EventName() string { return "settings_updated" }
