package application

import (
	"context"
	"crypto/rand"
	"fmt"
	"lolquizz/internal/domain/room"
	"lolquizz/internal/domain/shared"
	"math/big"
	"sync"

	"github.com/google/uuid"
)

type RoomService struct {
	rooms  room.Repository
	events EventPublisher
	idGen  func() string
	mu     sync.Mutex
}

func NewRoomService(rooms room.Repository, events EventPublisher, idGen func() string) *RoomService {
	return &RoomService{
		rooms:  rooms,
		events: events,
		idGen:  idGen,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, hostName string) (*room.Room, error) {
	code, err := s.generateUniqueCode(ctx)
	if err != nil {
		return nil, fmt.Errorf("generate room code: %w", err)
	}

	hostId := shared.PlayerId(uuid.New().String())

	host := room.NewPlayer(hostId, hostName)
	r := room.NewRoom(shared.RoomId(code), code, host)

	if err := s.rooms.Save(ctx, r); err != nil {
		return nil, fmt.Errorf("save room: %w", err)
	}

	return r, nil
}

func (s *RoomService) JoinRoom(ctx context.Context, code string, playerId shared.PlayerId, playerName string) (*room.Room, error) {
	r, err := s.rooms.FindByCode(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("find room: %w", err)
	}

	player := room.NewPlayer(playerId, playerName)
	if err := r.Join(player); err != nil {
		return nil, fmt.Errorf("join room: %w", err)
	}

	if err := s.rooms.Save(ctx, r); err != nil {
		return nil, fmt.Errorf("save room: %w", err)
	}

	s.events.PublishToRoom(r.Id, room.PlayerJoinedEvent{
		RoomId: r.Id,
		Player: player,
	})

	return r, nil
}

func (s *RoomService) LeaveRoom(ctx context.Context, code string, playerId shared.PlayerId) error {
	r, err := s.rooms.FindByCode(ctx, code)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	if err := r.Leave(playerId); err != nil {
		return fmt.Errorf("leave room: %w", err)
	}

	if err := s.rooms.Save(ctx, r); err != nil {
		return fmt.Errorf("save room: %w", err)
	}

	s.events.PublishToRoom(r.Id, room.PlayerLeftEvent{
		RoomId:  r.Id,
		NewHost: r.HostId,
	})

	return nil
}

func (s *RoomService) UpdateSettings(ctx context.Context, roomId shared.RoomId, settings room.Settings) error {
	r, err := s.rooms.FindById(ctx, roomId)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	r.Settings = settings

	if err := s.rooms.Save(ctx, r); err != nil {
		return fmt.Errorf("save room: %w", err)
	}

	s.events.PublishToRoom(r.Id, room.SettingsUpdatedEvent{
		RoomId:   r.Id,
		Settings: settings,
	})

	return nil
}

func (s *RoomService) GetRoom(ctx context.Context, code string) (*room.Room, error) {
	return s.rooms.FindByCode(ctx, code)
}

func (s *RoomService) generateUniqueCode(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	const charset = "ABCDEFGHJKLMNPQRSTUVWXYZ" // no I or O to avoid confusion
	const codeLen = 6
	const maxAttempts = 20

	for attempt := 0; attempt < maxAttempts; attempt++ {
		code := make([]byte, codeLen)
		for i := range code {
			n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
			if err != nil {
				return "", fmt.Errorf("random generation: %w", err)
			}
			code[i] = charset[n.Int64()]
		}

		candidate := string(code)

		_, err := s.rooms.FindByCode(ctx, candidate)
		if err != nil {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("failed to generate unique room code after %d attempts", maxAttempts)
}
