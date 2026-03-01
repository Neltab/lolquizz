package room

import "context"

type Repository interface {
	FindByID(ctx context.Context, id RoomId) (*Room, error)
	FindByCode(ctx context.Context, code string) (*Room, error)
	Save(ctx context.Context, room *Room) error
	Delete(ctx context.Context, id RoomId) error
}
