package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
)

// RegisterUser 定义一个注册用户的控制器函数
func RegisterUser(c *gin.Context) {
	// 定义一个用户结构体变量
	var registerUserJson RegisterUserJson

	// 从请求体中解析并绑定用户数据到 registerUserJson 变量
	err := c.ShouldBindJSON(&registerUserJson)
	// 如果解析出错，返回 400 状态码和错误信息
	if err != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "invalid request body",
			Data:   nil,
		})
		return
	}

	// 从上下文中获取数据库连接对象
	db := c.MustGet("db").(*gorm.DB)

	var user User
	user.Name = registerUserJson.Name
	user.Pwd = registerUserJson.Pwd

	// 查询用户名是否已经被使用
	result := db.Where("name = ? ", registerUserJson.Name).First(&user)
	// 如果查询出错，检查是否是 ErrRecordNotFound 错误
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 用户名不存在,不做操作
		} else {
			c.JSON(200, Message{
				Status: "error",
				Info:   "failed to create user",
				Data:   nil,
			})
			return
		}
	} else {
		// 用户名存在
		c.JSON(200, Message{
			Status: "error",
			Info:   "failed to create user, username already exists",
			Data:   nil,
		})
		return
	}

	// 在数据库中创建用户记录
	result = db.Create(&user)
	// 如果创建出错，返回 500 状态码和错误信息
	if result.Error != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "failed to create user",
			Data:   nil,
		})
		return
	}

	// 查询用户是否存在，取出完整的用户信息
	result = db.Where("name = ? AND pwd = ?", registerUserJson.Name, registerUserJson.Pwd).First(&user)
	if result.Error != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "failed to create user, invalid username or password",
			Data:   nil,
		})
		return
	}

	// 为用户生成一个 JWT 令牌
	tokenString, err := generateToken(user.ID)
	// 如果生成出错，返回 500 状态码和错误信息
	if err != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "failed to generate token",
			Data:   nil,
		})
		return
	}

	// 将 JWT 令牌分割成三部分：头部、负载和签名
	tokenParts := strings.Split(tokenString, ".")
	// 定义一个 JWT 令牌结构体变量，并赋值各个字段
	jwtToken := JwtToken{
		UserID:    user.ID,
		Header:    tokenParts[0],
		Payload:   tokenParts[1],
		Signature: tokenParts[2],
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	// 在数据库中保存 JWT 令牌记录
	result = db.Create(&jwtToken)
	// 如果保存出错，返回 500 状态码和错误信息
	if result.Error != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "failed to save token",
			Data:   nil,
		})
		return
	}

	// 返回 200 状态码和成功信息，以及用户 ID 和 JWT 令牌
	c.JSON(200, Message{
		Status: "success",
		Info:   "user registered successfully",
		Data: map[string]interface{}{
			"user_id": user.ID,
			"token":   tokenString,
		},
	})
}
