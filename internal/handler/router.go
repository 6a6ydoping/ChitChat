package handler

import (
	_ "github.com/6a6ydoping/ChitChat/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(customCors()))

	apiV1 := router.Group("/api/v1")
	apiV1.Use(h.authMiddleware())

	//User routes
	user := apiV1.Group("/user")
	user.POST("/register", h.createUser)
	user.POST("/login", h.loginUser)

	//Message routes
	message := apiV1.Group("/message")
	message.GET("", h.handleMessage)

	//Room routes
	room := apiV1.Group("/room")
	room.GET("", h.getRooms)
	room.POST("", h.createRoom)
	room.GET("/:roomID/info", h.getClients)

	//Non auth routes
	nonAuthGroup := apiV1.Group("/")
	nonAuthGroup.GET("/join/:roomID", h.joinRoom)

	//Swagger routes
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
