package ws

import (
	"encoding/json"
)

type IncomingMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type OutgoingMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
	State   interface{} `json:"state"`
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
