package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

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
	room.GET("", h.getRooms)
	room.POST("", h.createRoom)
	room.GET("/join/:roomID", h.joinRoom)

	return router
}
