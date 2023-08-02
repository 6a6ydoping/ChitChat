package service

import (
	"github.com/6a6ydoping/ChitChat/pkg/ws"
)

func (m *Manager) CreateRoom(room *ws.Room) {
	m.Dispatcher.Rooms[room.ID] = &ws.Room{
		ID:      room.ID,
		Name:    room.Name,
		Clients: make(map[string]*ws.Client),
	}
}

func (m *Manager) GetRooms() *[]ws.Room {
	rooms := make([]ws.Room, 0)
	for _, r := range m.Dispatcher.Rooms {
		rooms = append(rooms, ws.Room{
			ID:   r.ID,
			Name: r.Name,
		})
	}
	return &rooms
}
