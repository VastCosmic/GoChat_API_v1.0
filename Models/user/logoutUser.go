package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

// LogoutUser 定义一个注销用户的控制器函数
func LogoutUser(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(200, Message{
			Status: "error",
			Info:   "missing token header",
			Data:   nil,
		})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	result := db.Where("header = ? AND payload = ? AND signature = ?", strings.Split(tokenString, ".")[0], strings.Split(tokenString, ".")[1], strings.Split(tokenString, ".")[2]).Delete(&JwtToken{})
	if result.Error != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "failed to delete token",
			Data:   nil,
		})
		return
	}

	c.JSON(200, Message{
		Status: "success",
		Info:   "user logged out successfully",
		Data:   nil,
	})
}
