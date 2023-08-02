package service

import "github.com/6a6ydoping/ChitChat/pkg/ws"

type DispatcherService interface {
	CreateRoom(room *ws.Room)
}
