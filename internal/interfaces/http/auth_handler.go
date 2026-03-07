package http

import (
	"encoding/json"
	"lolquizz/internal/application"
	"net/http"
)

type AuthHandler struct {
	sessions application.SessionService
}

func NewAuthHandler(sessions application.SessionService) *AuthHandler {
	return &AuthHandler{
		sessions: sessions,
	}
}

func (h *AuthHandler) CreateSession(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Nickname string `json:"nickname"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.sessions.Create(playerId, req.Nickname)
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
