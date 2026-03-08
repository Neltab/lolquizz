package main

import (
	"log"
	"lolquizz/internal/application"
	"lolquizz/internal/config"
	"lolquizz/internal/infrastructure/memory"
	"lolquizz/internal/interfaces/ws"
	"net/http"

	httpPkg "lolquizz/internal/interfaces/http"

	"github.com/google/uuid"
)

func main() {
	cfg := config.Load()

	roomRepo := memory.NewRoomRepository()
	questionRepo := memory.NewQuestionRepository()

	hub := ws.NewHub()
	go hub.Run()

	idGen := func() string { return uuid.New().String() }

	roomService := application.NewRoomService(roomRepo, hub, idGen)
	gameService := application.NewGameService(roomRepo, hub, questionRepo, idGen)
	sessionService := application.NewSessionService(cfg.SessionTTL)

	authHandler := httpPkg.NewAuthHandler(sessionService)
	roomHandler := httpPkg.NewRoomHandler(roomService, sessionService)

	wsRouter := ws.NewRouter(hub, roomService, gameService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/auth/login", authHandler.CreateSession)
	mux.HandleFunc("POST /api/rooms", roomHandler.CreateRoom)
	mux.HandleFunc("GET /api/rooms/{code}", roomHandler.GetRoom)

	mux.HandleFunc("/ws", httpPkg.HandleWebsocket(hub, wsRouter, sessionService))

	mux.HandleFunc("/", httpPkg.SPAHandler("./web/dist"))

	var handler http.Handler = mux
	handler = httpPkg.LoggingMiddleware(handler)
	handler = httpPkg.CORSMiddleware(handler, cfg.AllowedOrigin)

	log.Printf("Server starting on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Fatal(err)
	}
}
