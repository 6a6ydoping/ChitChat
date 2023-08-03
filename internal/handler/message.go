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
		ID:   "",
		Name: req.Name,
	}
	h.dispatcherService.CreateRoom(r)

	c.JSON(http.StatusOK, req)
}

func (h *Handler) getRooms(c *gin.Context) {
	rooms := h.dispatcherService.GetRooms()
	c.JSON(http.StatusOK, rooms)
}

func (h *Handler) getClients(c *gin.Context) {
	var clients []api.ClientRes
	roomId := c.Param("roomId")

	clients = h.dispatcherService.GetClients(roomId)

	c.JSON(http.StatusOK, clients)
}

func (h *Handler) joinRoom(ctx *gin.Context) {
	authToken := ctx.Query("token")
	if authToken == "" {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty token"})
		fmt.Println(1)
		return
	}
	conn, err := h.WebsocketHandler.Upgrade(ctx.Writer, ctx.Request)
	if err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(2)
		return
	}

	roomID := ctx.Param("roomID")
	err = h.dispatcherService.JoinRoom(conn, authToken, roomID)
	if err != nil {
		fmt.Println(err)
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(3)
		return
	}
}
