package socket

import (
	_ "github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"social_network/utils/logger"
	"bytes"
)

type Client struct {
	ID   string
	conn *websocket.Conn
	hub  *Hub
	send chan []byte
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func UpgradeWS(hub *Hub, wrt http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(wrt, req, nil)
	if err != nil {
		logger.Error(err.Error())
	}

	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}

	client.hub.register <- client

	go client.Write()
	go client.Read()
}

func (c *Client) Read() {

	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Error("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}

}

func (c *Client) Write() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for msg := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
}
