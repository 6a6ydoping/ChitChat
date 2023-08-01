package handler

import (
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
