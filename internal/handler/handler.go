package handler

import (
	"github.com/6a6ydoping/ChitChat/internal/service"
	"github.com/6a6ydoping/ChitChat/pkg/ws"
	"github.com/gorilla/websocket"
	"net/http"
)

type WebsocketHandler interface {
	HandleWebSocket(conn *websocket.Conn)
	BroadcastMessage(msg string)
	Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error)
}

type Handler struct {
	service          service.Service
	WebsocketHandler WebsocketHandler
	Dispatcher       *ws.Dispatcher
}

func New(s service.Service) *Handler {
	return &Handler{
		service:    s,
		Dispatcher: ws.NewDispatcher(),
	}
}
