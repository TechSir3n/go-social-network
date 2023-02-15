package socket



type Hub struct {
	register   chan *Client     // for new connection
	unregister chan *Client     // for disconnection
	clients    map[*Client]bool //  view all clientsи
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			for client, _ := range h.clients {
				client.conn.WriteJSON("New User Connected")
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

			for client, _ := range h.clients {
				client.conn.WriteJSON("User Disconnected")
			}
			
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}

		}

	}
}
