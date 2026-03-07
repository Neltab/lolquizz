package ws

import (
	"context"
	"encoding/json"
	"lolquizz/internal/application"
)

type Router struct {
	roomService *application.RoomService
	gameService *application.GameService
	hub         *Hub
}

func NewRouter(hub *Hub, roomService *application.RoomService, gameService *application.GameService) *Router {
	return &Router{
		roomService: roomService,
		gameService: gameService,
		hub:         hub,
	}
}

func (r *Router) Handle(client *Client, msg IncomingMessage) {
	ctx := context.Background() //TODO: use context from client

	switch msg.Type {
	case MsgJoinRoom:
		r.handleJoinRoom(ctx, client, msg.Payload)
	// case MsgLeaveRoom:
	// 	r.handleLeaveRoom(client, msg)
	// case MsgUpdateSettings:
	// 	r.handleUpdateSettings(client, msg)
	// case MsgStartGame:
	// 	r.handleStartGame(client, msg)
	// case MsgSubmitAnswer:
	// 	r.handleSubmitAnswer(client, msg)
	// case MsgJudgeAnswer:
	// 	r.handleJudgeAnswer(client, msg)
	// case MsgNextRound:
	// 	r.handleNextRound(client, msg)
	default:
		client.SendError("unknown message type: " + msg.Type)
	}
}

func (r *Router) handleJoinRoom(ctx context.Context, client *Client, data json.RawMessage) {
	var req struct {
		RoomCode   string `json:"room_code"`
		PlayerName string `json:"player_name"`
	}
	if err := json.Unmarshal(data, &req); err != nil {
		client.SendError("invalid payload")
		return
	}
	room, err := r.roomService.JoinRoom(ctx, req.RoomCode, client.playerId, req.PlayerName)
	if err != nil {
		client.SendError(err.Error())
		return
	}

	r.hub.AddToRoom(client.playerId, room.Id)
}
