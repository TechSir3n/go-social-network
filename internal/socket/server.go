package socket

import (
	"bytes"
	"github.com/gorilla/websocket"
	"net/http"
	"social_network/utils/logger"
	"time"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Hub  *Hub
	Send chan *Message
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func NewClient(id string, hub *Hub, wrt http.ResponseWriter, req *http.Request) *Client {
	conn, err := upgrader.Upgrade(wrt, req, nil)
	if err != nil {
		logger.Error(err.Error())
	}

	client := &Client{
		ID:   id,
		Hub:  hub,
		Conn: conn,
		Send: make(chan *Message, 256),
	}

	go client.Write()
	go client.Read()

	return client
}

func (c *Client) Read() {

	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))

	c.Conn.SetPongHandler(func(png string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Error("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Hub.Broadcast <- &Message{
			Message:  string(message),
			ClientID: c.ID,
			Type:     "text",
		}

	}

}

func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.Conn.WriteJSON(message)
			if err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}

}
