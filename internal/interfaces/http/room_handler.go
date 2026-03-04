package http

import (
	"encoding/json"
	"lolquizz/internal/application"
	"lolquizz/internal/domain/shared"
	"net/http"
)

type RoomHandler struct {
	service *application.RoomService
}

func NewRoomHandler(service *application.RoomService) *RoomHandler {
	return &RoomHandler{
		service: service,
	}
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PlayerId shared.PlayerId `json:"player_id"`
		Nickname string          `json:"nickname"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	room, err := h.service.CreateRoom(r.Context(), req.PlayerId, req.Nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"room_id": room.Id,
		"code":    room.Code,
	})
}

func (h *RoomHandler) GetRoom(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	room, err := h.service.GetRoom(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(room)
}
