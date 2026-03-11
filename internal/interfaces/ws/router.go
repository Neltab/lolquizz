package ws

import (
	"context"
	"encoding/json"
	"log"
	"lolquizz/internal/application"
	"lolquizz/internal/domain/event"
	"lolquizz/internal/domain/game"
	"lolquizz/internal/domain/room"
	"lolquizz/internal/dto"
)

type Router struct {
	roomService *application.RoomService
	gameService *application.GameService
	hub         *Hub
}

func NewRouter(hub *Hub, roomService *application.RoomService, gameService *application.GameService, eventBus event.Subscriber) *Router {
	r := &Router{
		roomService: roomService,
		gameService: gameService,
		hub:         hub,
	}

	eventBus.Subscribe(MsgPlayerJoined, func(e event.Event) {
		ev := e.(*room.PlayerJoinedEvent)
		r.hub.PublishToRoom(ev.RoomId, OutgoingMessage{
			Type:    MsgPlayerJoined,
			Payload: dto.FromPlayer(ev.Player),
			State:   dto.FromRoom(ev.Room),
		})
	})

	eventBus.Subscribe(MsgPlayerLeft, func(e event.Event) {
		ev := e.(*room.PlayerLeftEvent)
		r.hub.PublishToRoom(ev.RoomId, OutgoingMessage{
			Type:    MsgPlayerLeft,
			Payload: dto.FromPlayer(ev.Player),
			State:   dto.FromRoom(ev.Room),
		})
	})

	eventBus.Subscribe(MsgSettingsUpdated, func(e event.Event) {
		ev := e.(*room.SettingsUpdatedEvent)
		r.hub.PublishToRoom(ev.RoomId, OutgoingMessage{
			Type:    MsgSettingsUpdated,
			Payload: ev.Settings,
			State:   dto.FromRoom(ev.Room),
		})
	})

	eventBus.Subscribe("round_started", func(e event.Event) {
		ev := e.(*game.QuestionStartedEvent)
		r.hub.PublishToRoom(ev.RoomId, OutgoingMessage{
			Type:    "round_started",
			Payload: "round_started",
		})
	})

	return r
}

func (r *Router) Handle(client *Client, msg IncomingMessage) {
	ctx := context.Background() //TODO: use context from client

	switch msg.Type {
	case MsgJoinRoom:
		r.handleJoinRoom(ctx, client, msg.Payload)
	case MsgLeaveRoom:
		r.handleLeaveRoom(ctx, client, msg.Payload)
	case MsgUpdateSettings:
		r.handleUpdateSettings(ctx, client, msg.Payload)
	case MsgStartGame:
		r.handleStartGame(ctx, client, msg.Payload)
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

func (r *Router) handleJoinRoom(ctx context.Context, client *Client, payload json.RawMessage) {
	var req struct {
		RoomCode   string `json:"room_code"`
		PlayerName string `json:"player_name"`
	}
	if err := json.Unmarshal(payload, &req); err != nil {
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

func (r *Router) handleLeaveRoom(ctx context.Context, client *Client, payload json.RawMessage) {
	var req struct {
		RoomCode string `json:"room_code"`
	}
	if err := json.Unmarshal(payload, &req); err != nil {
		client.SendError("invalid payload")
		return
	}

	if err := r.roomService.LeaveRoom(ctx, req.RoomCode, client.playerId); err != nil {
		client.SendError(err.Error())
		return
	}
}

func (r *Router) handleUpdateSettings(ctx context.Context, client *Client, payload json.RawMessage) {
	var req struct {
		RoomCode string `json:"room_code"`
		Settings room.Settings
	}
	if err := json.Unmarshal(payload, &req); err != nil {
		client.SendError("invalid payload")
		return
	}

	if err := r.roomService.UpdateSettings(ctx, req.RoomCode, &req.Settings); err != nil {
		client.SendError(err.Error())
		return
	}
}

func (r *Router) handleStartGame(ctx context.Context, client *Client, payload json.RawMessage) {
	var req struct {
		RoomCode string `json:"room_code"`
	}
	if err := json.Unmarshal(payload, &req); err != nil {
		client.SendError("invalid payload")
		return
	}

	room, err := r.roomService.GetRoom(context.Background(), req.RoomCode)
	if err != nil {
		client.SendError(err.Error())
		return
	}

	log.Printf("starting game")

	if err := r.gameService.StartGame(context.Background(), room.Id, client.playerId); err != nil {
		client.SendError(err.Error())
		return
	}
}
