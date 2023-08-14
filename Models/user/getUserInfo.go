package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// GetUserInfo 获取用户Info
func GetUserInfo(c *gin.Context) (ChatUserJson, error) {
	var chatUserJson ChatUserJson
	// 从请求头中获取 Authorization 字段的值，即 JWT 令牌
	tokenString := c.GetHeader("Authorization")
	// 如果请求头中没有 Authorization 字段，返回 400 状态码和错误信息
	if tokenString == "" {
		c.JSON(200, Message{
			Status: "error",
			Info:   "missing token header",
		})
		return chatUserJson, errors.New("missing token header")
	}

	// 从上下文中获取数据库连接对象
	db := c.MustGet("db").(*gorm.DB)

	// 从 JWT 令牌中提取负载部分的内容，即用户的 ID 和角色
	claims, err := extractClaims(tokenString)
	// 如果提取出错，返回 401 状态码和错误信息
	if err != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "invalid token",
		})
		return chatUserJson, errors.New("invalid token")
	}

	// 将用户 ID 从字符串转换为整数
	userID, err := strconv.Atoi(claims.Id)
	// 如果转换出错，返回 500 状态码和错误信息
	if err != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "failed to parse user id",
		})
		return chatUserJson, errors.New("failed to parse user id")
	}

	// 定义一个用户结构体变量
	var user User

	// 在数据库中查询用户，使用用户 ID 作为条件
	result := db.First(&user, userID)
	// 如果查询出错，返回 404 状态码和错误信息
	if result.Error != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "user not found",
		})
		return chatUserJson, errors.New("user not found")
	}

	// 脱敏后返回信息
	chatUserJson.ID = user.ID
	chatUserJson.Admin = user.Admin
	chatUserJson.Name = user.Name
	chatUserJson.Api = user.Api
	chatUserJson.ApiStatus = user.ApiStatus

	return chatUserJson, nil
}
