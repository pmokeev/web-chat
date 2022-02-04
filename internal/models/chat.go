package models

type Chat struct {
	Users    map[string]*User
	Messages chan *Message
	Join     chan *User
	Leave    chan *User
}

func NewChat() *Chat {
	return &Chat{}
}
