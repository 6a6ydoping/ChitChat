package service

import (
	"github.com/6a6ydoping/ChitChat/pkg/ws"
	"github.com/gorilla/websocket"
)

type DispatcherService interface {
	CreateRoom(*ws.Room)
	GetRooms() *[]ws.Room
	JoinRoom(conn *websocket.Conn, token, roomID string) error
}
