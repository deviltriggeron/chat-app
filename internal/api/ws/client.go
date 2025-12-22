package ws

import (
	"chat-app/internal/usecase"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn    *websocket.Conn
	send    chan []byte
	hub     *Hub
	usecase usecase.Usecase
}

func NewClient(
	conn *websocket.Conn,
	hub *Hub,
	uc usecase.Usecase,
) *Client {
	return &Client{
		conn:    conn,
		send:    make(chan []byte),
		hub:     hub,
		usecase: uc,
	}
}
