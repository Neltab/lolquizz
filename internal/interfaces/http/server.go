package http

import (
	"lolquizz/internal/domain/room"
	"lolquizz/internal/interfaces/ws"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebsocket(hub *ws.Hub, router *ws.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		playerId := room.PlayerId(r.URL.Query().Get("playerId"))
		if playerId == "" {
			http.Error(w, "playerId is required", http.StatusBadRequest)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "failed to upgrade websocket connection", http.StatusInternalServerError)
			return
		}

		client := ws.NewClient(hub, conn, playerId)
		hub.Register(client)

		go client.WritePump()
		go client.ReadPump(router)
	}
}
