package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"pmokeev/web-chat/internal/services"
)

type ChatController struct {
	chatService services.ChattingService
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin: func(request *http.Request) bool {
		return request.Method == http.MethodGet
	},
}

func NewChatController(chattingService services.ChattingService) *ChatController {
	return &ChatController{chatService: chattingService}
}

func (chatController *ChatController) ChatHandler(context *gin.Context) {
	connection, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		log.Fatal("Error on websocket connection: ", err.Error())
	}

	chatController.chatService.AddUser("username", connection)
}
