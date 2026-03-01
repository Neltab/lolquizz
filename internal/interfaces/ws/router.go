package ws

import "lolquizz/internal/application"

type Router struct {
	roomService *application.RoomService
	gameService *application.GameService
}

func NewRouter(roomService *application.RoomService, gameService *application.GameService) *Router {
	return &Router{
		roomService: roomService,
		gameService: gameService,
	}
}

func (r *Router) Handle(client *Client, msg IncomingMessage) {
	switch msg.Type {
	case MsgJoinRoom:
		r.handleJoinRoom(client, msg)
	case MsgLeaveRoom:
		r.handleLeaveRoom(client, msg)
	case MsgUpdateSettings:
		r.handleUpdateSettings(client, msg)
	case MsgStartGame:
		r.handleStartGame(client, msg)
	case MsgSubmitAnswer:
		r.handleSubmitAnswer(client, msg)
	case MsgJudgeAnswer:
		r.handleJudgeAnswer(client, msg)
	case MsgNextRound:
		r.handleNextRound(client, msg)
	default:
        c.SendError("unknown message type: " + msg.Type)
    }
}

func (r *Router) handleJoinRoom(ctx context.Context, client *Client, msg IncomingMessage) {
	var req struct {
		RoomCode string `json:"room_code"`
		playerName string `json:"player_name"`
	}
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		client.SendError("invalid payload")
		return
	}
	room, err := r.roomService.JoinRoom(ctx, )