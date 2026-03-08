package session

import (
	"lolquizz/internal/domain/shared"
	"time"
)

type Session struct {
	PlayerId  shared.PlayerId
	ExpiresAt time.Time
}
