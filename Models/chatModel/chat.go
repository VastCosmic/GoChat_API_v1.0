package chatModel

import (
	"Gin/Models/user"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Chat 定义一个函数，用于发送HTTP POST请求到OpenAI的API，并返回聊天对话补全结果
func Chat(c *gin.Context) {
	// 获取对话状态
	mode := c.PostForm("mode")
	if mode == "new" {
		fmt.Println("mode: new")
	} else if mode == "old" {
		fmt.Println("mode: old")
	} else {
		c.JSON(200, gin.H{"error": "mode error"})
		return
	}

	// 获取用户的消息
	userMessage := c.PostForm("message")
	if userMessage == "" {
		c.JSON(200, gin.H{"error": "message is required"})
		return
	}

	// 获取对话id,如果是新对话，chat_id为空
	chatID := c.PostForm("chat_id")
	if chatID == "" && mode == "old" {
		c.JSON(200, gin.H{"error": "chat_id is required"})
		return
	}

	// 获取用户Info
	var chatUserJson user.ChatUserJson
	chatUserJson, err := user.GetUserInfo(c)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	// 获取用户的API Key
	apikey := chatUserJson.Api

	fmt.Println("ChatID:", chatID)
	//fmt.Println("API:", apikey)
	fmt.Println("Message:", userMessage)

	params := ChatParams{}
	// 如果是新对话，只使用用户的消息
	if mode == "new" {
		params.Model = "gpt-3.5-turbo"
		params.Messages = []Message{
			{
				Role:    "user",
				Content: userMessage,
			},
		}
		params.Temperature = 1.0
	} else if mode == "old" {
		//如果是旧对话，将历史对话记录和用户的消息拼接起来
		//从数据库中读取历史对话记录
		//如果有历史对话记录，将历史对话记录和用户的消息拼接起来
		//如果没有历史对话记录，只使用用户的消息

		history, err := History(chatID)
		if err != nil {
			c.JSON(200, gin.H{"error": err.Error()})
			return
		}

		// 将历史对话记录提取出来
		messageHistory := []Message{}
		for _, value := range history.MessageReturn {
			messageHistory = append(messageHistory, Message{
				Role:    value.Role,
				Content: value.Content,
			})
		}
		// 将用户的消息拼接到历史对话记录中
		messageHistory = append(messageHistory, Message{
			Role:    "user",
			Content: userMessage,
		})

		// 打印历史对话记录 test
		fmt.Println("History:", messageHistory)

		params.Model = "gpt-3.5-turbo"
		params.Messages = messageHistory
		params.Temperature = 1.0
	}

	// 调用PostOpenaiChat函数，向openai发送请求，获取聊天对话和错误
	completion, err := PostOpenaiChat(params, apikey)
	// 如果有错误，返回错误
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	// 如果没有错误，返回聊天对话id、消息内容
	if len(completion.Choices) > 0 {
		// 将响应数据存储到数据库, 并取得聊天对话id
		if mode == "new" {
			newChatID, err := SaveNewChat(chatUserJson.ID, userMessage, completion)
			println("newChatID:", newChatID)
			chatID = strconv.Itoa(int(newChatID))
			if err != nil {
				c.JSON(200, gin.H{"error": err.Error()})
				return
			}
		} else {
			// 将对话数据存储到数据库
			err := SaveUserMessage(chatUserJson.ID, chatID, userMessage)
			if err != nil {
				c.JSON(200, gin.H{"error": err.Error()})
				return
			}
			err = SaveOldChat(chatUserJson.ID, chatID, completion)
			if err != nil {
				c.JSON(200, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(200, gin.H{
			"chat_id": chatID,
			"message": completion.Choices[0].Message.Content,
			"status":  "ok",
		})
		fmt.Println("Assistant:", completion.Choices[0].Message.Content)
	} else {
		c.JSON(200, gin.H{"error": "Request too ! No api available or No completion choices"})
	}
}
