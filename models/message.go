package models

import "time"

type Message struct {
	Id        int       `json:"id,omitempty" gorm:"primaryKey,autoIncrement:false"`
	ChatId    int       `json:"chatId,omitempty"`
	Content   string    `json:"content,omitempty"`
	SentBy    int       `json:"sentBy,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}