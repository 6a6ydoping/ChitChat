package httpserver

import (
	"context"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	server          *http.Server
	conns           map[*websocket.Conn]bool
	connsLock       sync.RWMutex
	shutdownTimeout time.Duration
	notify          chan error
	upgrader        websocket.Upgrader
}

func New(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler: handler,
	}

	s := &Server{
		server:    httpServer,
		conns:     make(map[*websocket.Conn]bool),
		connsLock: sync.RWMutex{},
		upgrader: websocket.Upgrader{ //TODO: check request origins
			CheckOrigin: func(r *http.Request) bool {
				return true
			}},
		notify: make(chan error, 1),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) HandleWebSocket(conn *websocket.Conn) {
	defer func() {
		s.connsLock.Lock()
		delete(s.conns, conn)
		s.connsLock.Unlock()
		conn.Close()
	}()

	s.connsLock.Lock()
	s.conns[conn] = true
	s.connsLock.Unlock()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		if messageType == websocket.TextMessage {
			msg := string(p)
			s.BroadcastMessage(msg)
		}
	}
	conn.WriteMessage(websocket.TextMessage, []byte("Thank you for writing!!!!"))
}

func (s *Server) BroadcastMessage(msg string) {
	s.connsLock.RLock()
	defer s.connsLock.RUnlock()

	for conn := range s.conns {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("Error sending message:", err)
			conn.Close()
			delete(s.conns, conn)
		}
	}
}

func (s *Server) Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return nil, err
	}
	return conn, nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)

	defer cancel()

	return s.server.Shutdown(ctx)
}

func (s *Server) Notify() <-chan error {
	return s.notify
}
