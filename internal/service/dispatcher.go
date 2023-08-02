package service

import (
	"fmt"
	"github.com/6a6ydoping/ChitChat/pkg/ws"
)

func (m *Manager) CreateRoom(room *ws.Room) {
	fmt.Println(m.Dispatcher.Rooms)
	fmt.Println("GERE")
	m.Dispatcher.Rooms[room.ID] = &ws.Room{
		ID:      room.ID,
		Name:    room.Name,
		Clients: make(map[string]*ws.Client),
	}
}
