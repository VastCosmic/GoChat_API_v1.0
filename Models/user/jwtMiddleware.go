package user

import (
	"github.com/gin-gonic/gin"
)

// JwtMiddleware 定义一个JWT中间件函数
func JwtMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(200, Message{
			Status: "error",
			Info:   "missing token header",
			Data:   nil,
		})
		c.Abort()
		return
	}

	valid, err := validateToken(tokenString)
	if err != nil || !valid {
		c.JSON(200, Message{
			Status: "error",
			Info:   "invalid token",
			Data:   nil,
		})
		c.Abort()
		return
	}

	c.Next()
}
