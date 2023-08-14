package chatModel

import (
	"Gin/Models/DB"
	"fmt"
	"strconv"
)

// History 从数据库中读取历史对话记录
func History(chatID string) (MessageHistoryReturn, error) {
	db := DB.Connect()

	var messagesHistory []MessageHistory
	if err := db.Where("chat_id = ?", chatID).Find(&messagesHistory).Error; err != nil {
		return MessageHistoryReturn{}, err
	}

	if len(messagesHistory) == 0 {
		return MessageHistoryReturn{}, fmt.Errorf("no messages found for chat")
	}

	//fmt.Println(MessageHistoryReturn{})

	chatIDint, _ := strconv.Atoi(chatID)
	messageReturn := MessageHistoryReturn{
		ChatID:        chatIDint,
		UserID:        messagesHistory[0].UserID,
		MessageReturn: make([]MessageReturn, len(messagesHistory)),
	}

	for i, message := range messagesHistory {
		messageReturn.MessageReturn[i] = MessageReturn{
			CreatedTime: message.CreatedTime,
			Role:        message.Role,
			Content:     message.Content,
		}
	}

	return messageReturn, nil
}
