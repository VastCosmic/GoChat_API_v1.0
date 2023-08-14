package Service

import (
	"Gin/Models/chatModel"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Chat 调用对话模型
func Chat(c *gin.Context) {
	chatModel.Chat(c)
}

// GetChatHistory 获取聊天记录
func GetChatHistory(c *gin.Context) {
	chatID := c.PostForm("chat_id")
	userID := c.PostForm("user_id")

	// 获取聊天记录
	history, err := chatModel.History(chatID)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 检查用户ID是否匹配
	if userID != strconv.Itoa(history.UserID) {
		c.JSON(200, gin.H{
			"error": "Unauthorized, check your user_id",
		})
		return
	}

	// 将history发送给客户端
	c.JSON(200, gin.H{
		"data":   history,
		"status": "ok",
	})
}
