package ws

import (
	"encoding/json"
	"log"
	"sync"
)

type playerId = string
type roomId = string

type Hub struct {
	clients    map[playerId]*Client
	rooms      map[roomId]map[playerId]bool
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[playerId]*Client),
		// Set of players in each room
		rooms:      make(map[roomId]map[playerId]bool),
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

func (h *Hub) Register(client *Client) {
	h.register <- client
}

func (h *Hub) Unregister(client *Client) {
	h.unregister <- client
}

func (h *Hub) AddToRoom(pId playerId, roomId roomId) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.rooms[roomId] == nil {
		h.rooms[roomId] = make(map[playerId]bool)
	}
	h.rooms[roomId][pId] = true

	if client, ok := h.clients[pId]; ok {
		client.roomId = roomId
	}
}

func (h *Hub) PublishToRoom(roomID roomId, msg OutgoingMessage) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	data, _ := json.Marshal(msg)
	playerIds := h.rooms[roomID]
	for playerId := range playerIds {
		if client, ok := h.clients[playerId]; ok {
			log.Printf("WS broadcast to player %s", playerId)
			select {
			case client.send <- data:
			default:
			}
		}
	}
}

func (h *Hub) PublishToPlayer(playerId playerId, msg OutgoingMessage) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if client, ok := h.clients[playerId]; ok {
		client.SendJson(msg)
	}
}
