package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
)

// LoginUser 定义一个登录用户的控制器函数
func LoginUser(c *gin.Context) {
	var loginUserJson LoginUserJson

	err := c.ShouldBindJSON(&loginUserJson)
	if err != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "invalid request body",
			Data:   nil,
		})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	var user User
	//// 在数据库中查询用户，使用 id 或者 name 或者 api 其中之一作为条件
	//result := db.Where("id = @id OR name = @name OR api = @api", map[string]interface{}{
	//	"id":   loginUserJson.ID,
	//	"name": loginUserJson.Name,
	//	"api":  loginUserJson.Api,
	//}).First(&user)
	//if result.Error != nil {
	//	c.JSON(200, Message{
	//		Status: "error",
	//		Info:   "invalid userid 、username、api or password",
	//		Data:   nil,
	//	})
	//	return
	//}

	// 在数据库中查询用户，使用 id 或者 name 其中之一作为条件
	result := db.Where("id = @id OR name = @name", map[string]interface{}{
		"id":   loginUserJson.ID,
		"name": loginUserJson.Name,
	}).First(&user)
	if result.Error != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "invalid userid 、username or password",
			Data:   nil,
		})
		return
	}
	// 检查密码是否正确
	if user.Pwd != loginUserJson.Pwd {
		c.JSON(200, Message{
			Status: "error",
			Info:   "invalid userid 、username or password",
			Data:   nil,
		})
		return
	}

	tokenString, err := generateToken(user.ID)
	if err != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "failed to generate token",
			Data:   nil,
		})
		return
	}

	tokenParts := strings.Split(tokenString, ".")
	jwtToken := JwtToken{
		UserID:    user.ID,
		Header:    tokenParts[0],
		Payload:   tokenParts[1],
		Signature: tokenParts[2],
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	result = db.Create(&jwtToken)
	if result.Error != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "failed to save token",
			Data:   nil,
		})
		return
	}

	c.JSON(200, Message{
		Status: "success",
		Info:   "user logged in successfully",
		Data: map[string]interface{}{
			"user_id": user.ID,
			"token":   tokenString,
		},
	})
}
