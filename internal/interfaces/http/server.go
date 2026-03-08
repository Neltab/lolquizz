package http

import (
	"log"
	"lolquizz/internal/application"
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

func HandleWebsocket(hub *ws.Hub, router *ws.Router, sessions *application.SessionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("WS request: path=%s query=%s", r.URL.Path, r.URL.RawQuery)
		token := r.URL.Query().Get("token")
		log.Printf("WS upgrade: token=%q", token)
		if token == "" {
			http.Error(w, "token is required", http.StatusBadRequest)
			return
		}

		session, err := sessions.Validate(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		log.Printf("WS upgrading for player %s", session.PlayerId)

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "failed to upgrade websocket connection", http.StatusInternalServerError)
			return
		}

		log.Printf("WS connected: player %s", session.PlayerId)

		client := ws.NewClient(hub, conn, session.PlayerId)
		hub.Register(client)

		go client.WritePump()
		go client.ReadPump(router)
	}
}
