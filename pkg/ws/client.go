package ws

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ID       string `json:"id"`
	Conn     *websocket.Conn
	Message  chan *Message
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}
