package memory

import (
	"context"
	"fmt"
	"lolquizz/internal/domain/room"
	"sync"
)

type RoomRepository struct {
	rooms map[room.RoomId]*room.Room
	mu    sync.RWMutex
}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{
		rooms: make(map[room.RoomId]*room.Room),
	}
}

func (r *RoomRepository) FindById(ctx context.Context, id room.RoomId) (*room.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	room, ok := r.rooms[id]
	if !ok {
		return nil, fmt.Errorf("room %s not found", id)
	}

	return room, nil
}

func (r *RoomRepository) FindByCode(ctx context.Context, code string) (*room.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	room, ok := r.rooms[room.RoomId(code)]
	if !ok {
		return nil, fmt.Errorf("room %s not found", code)
	}

	return room, nil
}

func (r *RoomRepository) Save(ctx context.Context, room *room.Room) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.rooms[room.Id] = room
	return nil
}

func (r *RoomRepository) Delete(ctx context.Context, id room.RoomId) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.rooms, id)
	return nil
}
