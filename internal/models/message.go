package models

type Message struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	Sender string `json:"sender"`
}

func NewMessage(body string, sender string) *Message {
	return &Message{
		ID:     1, // do normal numbering of messages
		Body:   body,
		Sender: sender,
	}
}
