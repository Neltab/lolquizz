package ws

import (
	"encoding/json"
	"log"
	"time"

	"lolquizz/internal/domain/shared"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 4096
)

type Client struct {
	hub      *Hub
	conn     *websocket.Conn
	send     chan []byte
	playerId shared.PlayerId
	roomId   shared.RoomId
}

func NewClient(hub *Hub, conn *websocket.Conn, playerId shared.PlayerId) *Client {
	return &Client{
		hub:      hub,
		conn:     conn,
		send:     make(chan []byte, 256),
		playerId: playerId,
	}
}

func (c *Client) ReadPump(router *Router) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		var msg IncomingMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			c.SendError("invalid message format")
			continue
		}

		// Route to the appropriate handler
		router.Handle(c, msg)
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) SendJson(msg OutgoingMessage) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("marshal error: %v", err)
		return
	}
	select {
	case c.send <- data:
	default:
		close(c.send)
	}
}

func (c *Client) SendError(msg string) {
	c.SendJson(OutgoingMessage{
		Type:    MsgError,
		Payload: map[string]string{"message": msg},
	})
}
