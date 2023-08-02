package handler

import (
	"github.com/6a6ydoping/ChitChat/api"
	"github.com/6a6ydoping/ChitChat/pkg/ws"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) handleMessage(ctx *gin.Context) {
	conn, err := h.WebsocketHandler.Upgrade(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upgrade to WebSOCket"})
		return
	}
	defer conn.Close()

	h.WebsocketHandler.HandleWebSocket(conn)
}

func (h *Handler) createRoom(c *gin.Context) {
	var req api.CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.Dispatcher.Rooms[req.ID] = &ws.Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*ws.Client),
	}

	c.JSON(http.StatusOK, req)
}

func (h *Handler) getRooms(c *gin.Context) {
	rooms := make([]api.RoomRes, 0)

	for _, r := range h.Dispatcher.Rooms {
		rooms = append(rooms, api.RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}
