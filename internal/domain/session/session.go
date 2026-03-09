package session

import (
	"time"
)

type Session struct {
	PlayerId  PlayerId
	ExpiresAt time.Time
}
