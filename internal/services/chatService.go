package services

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"pmokeev/web-chat/internal/models"
)

type ChatService struct {
	chat *models.Chat
}

func NewChatSerivce() *ChatService {
	chatService := &ChatService{chat: &models.Chat{
		Users:    make(map[string]*models.User),
		Messages: make(chan *models.Message),
		Join:     make(chan *models.User),
		Leave:    make(chan *models.User),
	}}
	go chatService.RunChat()
	return chatService
}

func (chatService *ChatService) RunChat() {
	for {
		select {
		case user := <-chatService.chat.Join:
			chatService.connectUser(user)
		case message := <-chatService.chat.Messages:
			chatService.broadcastMessage(message)
		case user := <-chatService.chat.Leave:
			chatService.disconnectUser(user)
		}
	}
}

func (chatService *ChatService) AddUser(username string, connection *websocket.Conn) {
	user := models.NewUser(username, connection, chatService.chat)
	chatService.chat.Join <- user
	user.Read()
}

func (chatService *ChatService) connectUser(user *models.User) {
	if _, ok := chatService.chat.Users[user.Username]; !ok {
		chatService.chat.Users[user.Username] = user
		body := fmt.Sprintf("%s join to chat", user.Username)
		chatService.broadcastMessage(models.NewMessage(body, "Server"))
	}
}

func (chatService *ChatService) broadcastMessage(message *models.Message) {
	log.Printf("Broadcast message: %v\n", message)
	for _, user := range chatService.chat.Users {
		user.Write(message)
	}
}

func (chatService *ChatService) disconnectUser(user *models.User) {
	if _, ok := chatService.chat.Users[user.Username]; ok {
		defer user.Connection.Close()
		delete(chatService.chat.Users, user.Username)

		body := fmt.Sprintf("%s left the chat", user.Username)
		chatService.broadcastMessage(models.NewMessage(body, "Server"))
	}
}
