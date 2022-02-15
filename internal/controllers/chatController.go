package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gopkg.in/dgrijalva/jwt-go.v3"
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
	cookie, err := context.Request.Cookie("jwt")
	if err != nil {
		if err == http.ErrNoCookie {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	JWTTokenString := cookie.Value
	username, err, isValid := chatController.chatService.GetUsername(JWTTokenString)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if !isValid {
		context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	connection, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		log.Fatal("Error on websocket connection: ", err.Error())
	}

	chatController.chatService.AddUser(username, connection)
}
