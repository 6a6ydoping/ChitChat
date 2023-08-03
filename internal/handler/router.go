package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))

	apiV1 := router.Group("/api/v1")

	user := apiV1.Group("/user")
	message := apiV1.Group("/message")
	room := apiV1.Group("/room")

	//User routes
	user.POST("/register", h.createUser)
	user.POST("/login", h.loginUser)

	apiV1.Use(h.authMiddleware())
	//Message routes
	message.GET("", h.handleMessage)

	//Room routes
	apiV1.GET("/join/:roomID", h.joinRoom)
	room.GET("", h.getRooms)
	room.POST("", h.createRoom)
	room.GET("/:roomID/info", h.getClients)

	return router
}
