package http

import (
	"encoding/json"
	"lolquizz/internal/application"
	"net/http"
)

type RoomHandler struct {
	roomService    *application.RoomService
	sessionService *application.SessionService
}

func NewRoomHandler(roomService *application.RoomService, sessionService *application.SessionService) *RoomHandler {
	return &RoomHandler{
		roomService:    roomService,
		sessionService: sessionService,
	}
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token    string `json:"token"`
		Nickname string `json:"nickname"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	session, err := h.sessionService.Validate(req.Token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	room, err := h.roomService.CreateRoom(r.Context(), session.PlayerId, req.Nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    room.Code,
		"is_host": true,
	})
}

func (h *RoomHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	room, err := h.roomService.GetRoom(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(room)
}
