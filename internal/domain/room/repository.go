package room

import (
	"context"
	"lolquizz/internal/domain/shared"
)

type Repository interface {
	FindById(ctx context.Context, id shared.RoomId) (*Room, error)
	FindByCode(ctx context.Context, code string) (*Room, error)
	Save(ctx context.Context, room *Room) error
	Delete(ctx context.Context, id shared.RoomId) error
}
