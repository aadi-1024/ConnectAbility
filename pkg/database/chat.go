package database

import (
	"errors"
	"github.com/aadi-1024/ConnectAbility/models"
	"time"
)

func (d *Database) CreateChat(chat models.Chat) error {
	return d.conn.Create(&chat).Error
}

func (d *Database) SendMessage(chatId, uid int, content string) (int, error) {
	chat := models.Chat{}
	if err := d.conn.Table("chats").Find(&chat, chatId).Error; err != nil {
		return 0, err
	}

	if chat.Member1 != uid && chat.Member2 != uid {
		return 0, errors.New("user not in specified chatroom")
	}

	msg := models.Message{
		ChatId:    chatId,
		Content:   content,
		SentBy:    uid,
		Timestamp: time.Now(),
	}
	return msg.Id, d.conn.Table("messages").Create(&msg).Error
}

func (d *Database) GetMessages(chatId, uid, limit int) ([]*models.Message, error) {
	chat := models.Chat{}
	if err := d.conn.Table("chats").Find(&chat, chatId).Error; err != nil {
		return nil, err
	}

	if chat.Member1 != uid && chat.Member2 != uid {
		return nil, errors.New("user not in specified chatroom")
	}

	data := make([]*models.Message, limit)
	err := d.conn.Table("messages").Limit(limit).Order("timestamp DESC").Find(&data).Where("Chat_Id = ?", chatId).Error

	return data, err
}
