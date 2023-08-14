package Service

import (
	"Gin/Models/user"
	"github.com/gin-gonic/gin"
)

// RegisterUser 用户注册
func RegisterUser(c *gin.Context) {
	user.RegisterUser(c)
}

// LoginUser 用户登录
func LoginUser(c *gin.Context) {
	user.LoginUser(c)
}

// GetCurrentUser 获取当前用户信息
func GetCurrentUser(c *gin.Context) {
	user.GetCurrentUser(c)
}

// LogoutUser 用户退出登录
func LogoutUser(c *gin.Context) {
	user.LogoutUser(c)
}

// UpdateUserAPI 更新用户API
func UpdateUserAPI(c *gin.Context) {
	user.UpdateAPI(c)
}

// UpdateUser 更新用户信息
func UpdateUser(c *gin.Context) {
	user.UpdateUserInfo(c)
}
