package http

import (
	"encoding/json"
	"lolquizz/internal/application"
	"lolquizz/internal/domain/shared"
	"net/http"

	"github.com/google/uuid"
)

type AuthHandler struct {
	sessions *application.SessionService
}

func NewAuthHandler(sessions *application.SessionService) *AuthHandler {
	return &AuthHandler{
		sessions: sessions,
	}
}

func (h *AuthHandler) CreateSession(w http.ResponseWriter, r *http.Request) {
	playerId := shared.PlayerId(uuid.New().String())
	token, err := h.sessions.Create(playerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})
}
