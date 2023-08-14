package chatModel

import (
	"Gin/Models/DB"
	"errors"
	"fmt"
	"strconv"
)

// SaveNewChat 将响应的聊天对话保存到数据库中,返回聊天ID
func SaveNewChat(userID uint, userMessage string, completion ChatCompletion) (uint, error) {
	// 首先存储用户消息，并取得对话ID
	var messageTable MessageTable
	messageTable.UserID = userID
	messageTable.Role = "user"
	messageTable.Content = userMessage
	messageTable.PromptTokens = 0
	messageTable.CompletionTokens = 0
	messageTable.TotalTokens = 0

	// 将对话表转储到数据库中,并从数据库中取得对话ID
	// 备注：新的对话ID使用新建的单条对话记录ID
	db := DB.Connect()
	result := db.Create(&messageTable)
	if result.Error != nil {
		return 0, result.Error
	}

	// 声明对话ID
	var chatID uint

	if result.RowsAffected == 1 { // 插入成功
		// 取得数据库中的对话ID, 将插入后的数据扫描到newMessageTable中
		var newMessageTable MessageTable
		result.Scan(&newMessageTable)
		// 将单条对话记录ID赋值给对话ID
		newMessageTable.ChatID = newMessageTable.RecordID
		chatID = newMessageTable.ChatID
		// 更新对话表中的对话ID
		db.Model(&newMessageTable).Where("record_id = ?", newMessageTable.RecordID).Update("chat_id", newMessageTable.ChatID)
		// 返回对话ID
		println("newMessageTable.ChatID:", newMessageTable.ChatID)
	}

	fmt.Println("messageTable:", messageTable)

	// 修改对话表的内容为响应的内容
	messageTable = MessageTable{} // 重置对话表
	messageTable.UserID = userID
	messageTable.ChatID = chatID
	messageTable.Role = completion.Choices[0].Message.Role
	messageTable.Content = completion.Choices[0].Message.Content
	messageTable.PromptTokens = completion.Usage.PromptTokens
	messageTable.CompletionTokens = completion.Usage.CompletionTokens
	messageTable.TotalTokens = completion.Usage.TotalTokens

	fmt.Println("messageTable:", messageTable)

	result = db.Create(&messageTable)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 1 { // 插入成功
		return chatID, nil
	}
	return 0, errors.New("插入失败")
}

// SaveOldChat 将响应的聊天对话保存到数据库中
func SaveOldChat(userID uint, chatID string, completion ChatCompletion) error {
	// 声明一个对话表，用于转储对话
	var messageTable MessageTable
	messageTable.UserID = userID

	uChatID, err := strconv.ParseUint(chatID, 10, 0) // uint64
	if err != nil {
		fmt.Println(err) // 转换失败，打印错误信息
		return err
	}
	messageTable.ChatID = uint(uChatID)

	messageTable.Role = completion.Choices[0].Message.Role
	messageTable.Content = completion.Choices[0].Message.Content
	messageTable.PromptTokens = completion.Usage.PromptTokens
	messageTable.CompletionTokens = completion.Usage.CompletionTokens
	messageTable.TotalTokens = completion.Usage.TotalTokens

	// 将对话表转储到数据库中,并从数据库中取得对话ID
	db := DB.Connect()
	result := db.Create(&messageTable)
	if result.Error != nil {
		println(result.Error)
		return result.Error
	}
	if result.RowsAffected == 1 { // 插入成功
		return nil
	}
	return errors.New("插入失败")
}

// SaveUserMessage 将用户的消息保存到数据库中
func SaveUserMessage(userID uint, chatID string, message string) error {
	// 声明一个对话表，用于转储对话
	var messageTable MessageTable
	messageTable.UserID = userID

	uChatID, err := strconv.ParseUint(chatID, 10, 0) // uint64
	if err != nil {
		fmt.Println(err) // 转换失败，打印错误信息
		return err
	}
	messageTable.ChatID = uint(uChatID)

	messageTable.Role = "user"
	messageTable.Content = message

	// 将对话表转储到数据库中,并从数据库中取得对话ID
	db := DB.Connect()
	result := db.Create(&messageTable)
	if result.Error != nil {
		println(result.Error)
		return result.Error
	}
	if result.RowsAffected == 1 { // 插入成功
		return nil
	}
	return errors.New("插入失败")
}
