package main

import (
	"strings"
	"time"
)

type Searcher interface {
	Contains(text string) bool
}

type TimeStamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTimeStamp() TimeStamp {
	return TimeStamp{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type Document struct {
	ID     int
	UserID int
	Name   string
	Value  string
	TimeStamp
}

func NewDocument(userID int, name, value string) *Document {
	return &Document{
		UserID:    userID,
		Name:      name,
		Value:     value,
		TimeStamp: NewTimeStamp(),
	}
}

func (document *Document) Contains(text string) bool {
	return strings.Contains(document.Name, text) || strings.Contains(document.Value, text)
}

type Message struct {
	ID         int
	RecieverID int
	SenderID   int
	Text       string
	TimeStamp
}

func NewMessage(receiverId, senderId int, text string) *Message {
	return &Message{
		RecieverID: receiverId,
		SenderID:   senderId,
		Text:       text,
		TimeStamp:  NewTimeStamp(),
	}
}

func (message *Message) Contains(text string) bool {
	return strings.Contains(message.Text, text)
}

func SeacrhText(text string, values []Searcher) []Searcher {
	result := make([]Searcher, 0)
	for _, v := range values {
		if v.Contains(text) {
			result = append(result, v)
		}
	}
	return result
}
