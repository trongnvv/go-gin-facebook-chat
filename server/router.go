package server

import (
	"os"
	"trongnv-chat/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ROUTING := os.Getenv("ROUTING")

	chatController := controllers.NewChatController()
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	// router.Use(helpers.IsAuthenticated())

	v1 := router.Group(ROUTING)

	v1.GET("/:pageID/messages", chatController.GetMessages)
	v1.GET("/:pageID/conversations", chatController.GetConversations)

	return router
}
