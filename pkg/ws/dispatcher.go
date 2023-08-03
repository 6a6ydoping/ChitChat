package ws

import "fmt"

type Dispatcher struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

type Room struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Clients map[string]*Client
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 10),
	}
}

func (d *Dispatcher) Run() {
	for {
		select {
		// Join Room
		case client := <-d.Register:
			// Check if such room exists
			if _, ok := d.Rooms[client.RoomID]; ok {
				r := d.Rooms[client.RoomID]
				// Check if user already in the room
				if _, ok := r.Clients[client.ID]; !ok {
					r.Clients[client.ID] = client
					for _, cl := range d.Rooms[client.RoomID].Clients {
						fmt.Println(cl.Username)
					}
				}
			}
		// Leave Room
		case client := <-d.Unregister:
			// Check for room existence
			if _, ok := d.Rooms[client.RoomID]; ok {
				r := d.Rooms[client.RoomID]
				// Check if user in room
				if _, ok := r.Clients[client.ID]; ok {
					// If room not empty
					if len(r.Clients) != 0 {
						d.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomID:   client.RoomID,
							Username: client.Username,
						}
					}
					// finally delete client
					delete(d.Rooms[client.RoomID].Clients, client.ID)
					close(client.Message)
				}
			}
		// Broadcast message
		case message := <-d.Broadcast:
			if _, ok := d.Rooms[message.RoomID]; ok {
				for _, cl := range d.Rooms[message.RoomID].Clients {
					cl.Message <- message
				}
			}
		}
	}
}
