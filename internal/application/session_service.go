package application

import (
	"lolquizz/internal/domain/session"
	"lolquizz/internal/domain/shared"
	"math/rand"
	"sync"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type SessionService struct {
	sessions map[string]session.Session
	mu       sync.RWMutex
	ttl      time.Duration
}

func NewSessionService(ttl time.Duration) *SessionService {
	return &SessionService{
		sessions: make(map[string]session.Session),
		ttl:      ttl,
	}
}

func (s *SessionService) Create(playerId shared.PlayerId, nickname string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	token := StringWithCharset(32, charset)
	expiresAt := time.Now().Add(s.ttl)

	s.sessions[token] = session.Session{
		PlayerId:  playerId,
		Nickname:  nickname,
		ExpiresAt: expiresAt,
	}

	return token, nil
}

func (s *SessionService) Validate(token string) (*session.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	session, ok := s.sessions[token]
	if !ok {
		return nil, nil
	}

	if session.ExpiresAt.Before(time.Now()) {
		return nil, nil
	}

	return &session, nil
}

func StringWithCharset(length int, charset string) string {
	// TODO: make a util package
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
