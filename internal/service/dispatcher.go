package service

import (
	"fmt"
	"github.com/6a6ydoping/ChitChat/api"
	"github.com/6a6ydoping/ChitChat/pkg/ws"
	"github.com/gorilla/websocket"
	"math/rand"
	"strconv"
	"time"
)

func (m *Manager) CreateRoom(room *ws.Room) {
	// TODO: check db for id's
	rand.Seed(time.Now().UnixNano())
	roomID := strconv.Itoa(rand.Intn(1000))
	m.Dispatcher.Rooms[roomID] = &ws.Room{
		ID:      roomID,
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
	if err != nil || claims["user_id"] == "0" || claims["username"] == "" {
		return fmt.Errorf("error while extracting claims: %s", err)
	}
	fmt.Println("claims!")
	fmt.Println(claims["username"])

	r, ok := m.Dispatcher.Rooms[roomID]
	if !ok {
		fmt.Println("Room doesnt exist, creating new one")
		// If the room doesn't exist, create it
		r = &ws.Room{
			ID:      roomID,
			Clients: make(map[string]*ws.Client),
		}
		m.Dispatcher.Rooms[roomID] = r
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

func (m *Manager) GetClients(roomId string) []api.ClientRes {
	var clients []api.ClientRes

	if _, ok := m.Dispatcher.Rooms[roomId]; !ok {
		clients = make([]api.ClientRes, 0)
		return nil
	}
	fmt.Println(m.Dispatcher.Rooms[roomId].Clients)
	for _, c := range m.Dispatcher.Rooms[roomId].Clients {
		clients = append(clients, api.ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	return clients
}
