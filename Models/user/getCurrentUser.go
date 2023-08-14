package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// GetCurrentUser 定义一个获取当前用户信息的控制器函数
func GetCurrentUser(c *gin.Context) {
	// 从请求头中获取 Authorization 字段的值，即 JWT 令牌
	tokenString := c.GetHeader("Authorization")
	// 如果请求头中没有 Authorization 字段，返回 400 状态码和错误信息
	if tokenString == "" {
		c.JSON(200, Message{
			Status: "error",
			Info:   "missing token header",
			Data:   nil,
		})
		return
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
			Data:   nil,
		})
		return
	}

	// 将用户 ID 从字符串转换为整数
	userID, err := strconv.Atoi(claims.Id)
	// 如果转换出错，返回 500 状态码和错误信息
	if err != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "failed to parse user id",
			Data:   nil,
		})
		return
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
			Data:   nil,
		})
		return
	}

	// 定义一个当前用户模型，仅返回不敏感信息
	var currentUserJson CurrentUserJson

	// 将用户结构体变量的字段赋值给当前用户模型的字段
	currentUserJson.ID = user.ID
	currentUserJson.Name = user.Name
	currentUserJson.Admin = user.Admin
	currentUserJson.ApiStatus = user.ApiStatus

	// 返回 200 状态码和成功信息，以及当前用户模型的数据
	c.JSON(200, Message{
		Status: "success",
		Info:   "user found",
		Data:   currentUserJson,
	})
}
