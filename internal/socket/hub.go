package socket

import (
	"math/rand"
	"time"
)

type Hub struct {
	ID         int32
	Register   chan *Client     // for new connection
	Unregister chan *Client     // for disconnection
	Clients    map[*Client]bool //  view all clientsи
	Broadcast  chan *Message
}

type Message struct {
	Message  string `json:"message"`
	Type     string `json:"type"`
	ClientID string `json:"client_id"`
}

func NewHub() *Hub {
	rand.Seed(time.Now().UnixNano())
	hub := &Hub{
		ID:         rand.Int31(),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan *Message),
	}
	return hub
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			for client, _ := range h.Clients {
				client.Conn.WriteJSON("New User Connected")
			}

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

			for client, _ := range h.Clients {
				client.Conn.WriteJSON("User Disconnected")
			}

		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}

		}

	}
}
