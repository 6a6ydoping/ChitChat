package handler

import (
	"github.com/6a6ydoping/ChitChat/internal/service"
	"github.com/gorilla/websocket"
	"net/http"
)

type WebsocketHandler interface {
	HandleWebSocket(conn *websocket.Conn)
	BroadcastMessage(msg string)
	Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error)
}

type Handler struct {
	userService       service.UserService
	dispatcherService service.DispatcherService
	WebsocketHandler  WebsocketHandler
}

func New(s service.UserService, ds service.DispatcherService) *Handler {
	return &Handler{
		userService:       s,
		dispatcherService: ds,
	}
}
