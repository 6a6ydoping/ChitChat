package handler

import (
	"fmt"
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
	r := &ws.Room{
		ID:   req.ID,
		Name: req.Name,
	}
	fmt.Println(r)
	h.dispatcherService.CreateRoom(r)

	c.JSON(http.StatusOK, req)
}

//func (h *Handler) getRooms(c *gin.Context) {
//	rooms := make([]api.RoomRes, 0)
//
//	for _, r := range h.Dispatcher.Rooms {
//		rooms = append(rooms, api.RoomRes{
//			ID:   r.ID,
//			Name: r.Name,
//		})
//	}
//
//	c.JSON(http.StatusOK, rooms)
//}
//
//func (h *Handler) joinRoom(ctx *gin.Context) {
//	conn, err := h.WebsocketHandler.Upgrade(ctx.Writer, ctx.Request)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	roomID := ctx.Param("roomId")
//
//	cl := &ws.Client{
//		Conn:     conn,
//		Message:  make(chan *ws.Message, 10),
//		ID:       "",
//		RoomID:   roomID,
//		Username: "",
//	}
//
//	m := &ws.Message{
//		Content:  "A new user has joined the room",
//		RoomID:   roomID,
//		Username: "",
//	}
//
//	h.Dispatcher.Register <- cl
//	h.Dispatcher.Broadcast <- m
//
//	go cl.WriteMessage()
//	cl.ReadMessage(h.Dispatcher)
//}
