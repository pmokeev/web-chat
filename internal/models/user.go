package models

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type User struct {
	Username   string
	Connection *websocket.Conn
	Global     *Chat
}

func NewUser(username string, connection *websocket.Conn, global *Chat) *User {
	return &User{Username: username,
		Connection: connection,
		Global:     global}
}

func (user *User) Read() {
	for {
		if _, message, err := user.Connection.ReadMessage(); err != nil {
			log.Println("Error on read message:", err.Error())
			break
		} else {
			user.Global.Messages <- NewMessage(string(message), user.Username)
		}
	}

	user.Global.Leave <- user
}

func (user *User) Write(message *Message) {
	body, _ := json.Marshal(message)
	if err := user.Connection.WriteMessage(websocket.TextMessage, body); err != nil {
		log.Println("Error on write message:", err.Error())
	}
}
