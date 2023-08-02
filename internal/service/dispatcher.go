package service

import (
	"fmt"
	"github.com/6a6ydoping/ChitChat/pkg/ws"
	"github.com/gorilla/websocket"
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

func (m *Manager) JoinRoom(conn *websocket.Conn, token, roomID string) error {
	claims, err := m.Token.ExtractClaimsFromString(token)
	if err != nil || claims["id"] == "0" || claims["username"] == "" {
		return fmt.Errorf("error while extracting claims: %s", err)
	}
	// Get ID and username from jwt
	cl := &ws.Client{
		Conn:     conn,
		Message:  make(chan *ws.Message, 10),
		ID:       claims["id"],
		RoomID:   roomID,
		Username: claims["username"],
	}

	msg := &ws.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: claims["username"],
	}

	m.Dispatcher.Register <- cl
	m.Dispatcher.Broadcast <- msg

	go cl.WriteMessage()
	cl.ReadMessage(m.Dispatcher)

	return nil
}
