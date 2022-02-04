package models

import (
	"math/rand"
	"time"
)

func GetRandom() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetRandomI64() int64 {
	return GetRandom().Int63()
}

type Message struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	Sender string `json:"sender"`
}

func NewMessage(body string, sender string) *Message {
	return &Message{
		ID:     GetRandomI64(), // do normal numbering of messages
		Body:   body,
		Sender: sender,
	}
}
