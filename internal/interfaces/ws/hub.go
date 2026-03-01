package ws

import (
	"encoding/json"
	"lolquizz/internal/domain/event"
	"lolquizz/internal/domain/room"
	"sync"
)

type Hub struct {
	clients    map[room.PlayerId]*Client
	rooms      map[room.RoomId]map[room.PlayerId]bool
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[room.PlayerId]*Client),
		// Set of players in each room
		rooms:      make(map[room.RoomId]map[room.PlayerId]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.registerClient(client)
		case client := <-h.unregister:
			h.unregisterClient(client)
		}
	}
}

func (h *Hub) registerClient(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.clients[client.playerId] = client
}

func (h *Hub) unregisterClient(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.clients[client.playerId]; ok {
		delete(h.clients, client.playerId)
		close(client.send)
		if client.roomId != "" {
			delete(h.rooms[client.roomId], client.playerId)
		}
	}
}

func (h *Hub) AddToRoom(playerId room.PlayerId, roomId room.RoomId) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.rooms[roomId] == nil {
		h.rooms[roomId] = make(map[room.PlayerId]bool)
	}
	h.rooms[roomId][playerId] = true

	if client, ok := h.clients[playerId]; ok {
		client.roomId = roomId
	}
}

func (h *Hub) BroadcastToRoom(roomID room.RoomId, msg OutgoingMessage) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	data, _ := json.Marshal(msg)
	playerIds := h.rooms[roomID]
	for playerId := range playerIds {
		if client, ok := h.clients[playerId]; ok {
			select {
			case client.send <- data:
			default:
			}
		}
	}
}

func (h *Hub) SendToPlayer(playerId room.PlayerId, msg OutgoingMessage) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if client, ok := h.clients[playerId]; ok {
		client.SendJson(msg)
	}
}

func (h *Hub) PublishToROom(roomId room.RoomId, event event.Event) {
	h.BroadcastToRoom(roomId, OutgoingMessage{
		Type:    event.EventName(),
		Payload: event,
	})
}

func (h *Hub) PublishToPlayer(playerId room.PlayerId, event event.Event) {
	h.SendToPlayer(playerId, OutgoingMessage{
		Type:    event.EventName(),
		Payload: event,
	})
}
