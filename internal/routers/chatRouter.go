package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"pmokeev/web-chat/internal/controllers"
	"pmokeev/web-chat/internal/services"
	"time"
)

type ChatRouter struct {
	controller *controllers.Controller
}

func NewChatRouter(service *services.Service) *ChatRouter {
	return &ChatRouter{controller: controllers.NewController(service)}
}

func (chatRouter *ChatRouter) InitChatRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "GET"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	chat := router.Group("/api")
	{
		chat.GET("/chat", chatRouter.controller.ChatHandler)
	}

	return router
}
