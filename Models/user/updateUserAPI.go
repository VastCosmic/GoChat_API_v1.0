package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UpdateAPI 定义UpdateAPI函数，接收一个用户ID和一个新的API值作为参数
func UpdateAPI(c *gin.Context) {
	// 连接数据库，这里假设已经存在一个名为db的*gorm.DB对象
	db := c.MustGet("db").(*gorm.DB)

	// 取得用户ID以及新的API值
	var updateUserAPIJson UpdateAPIUserJson
	err := c.ShouldBindJSON(&updateUserAPIJson)
	if err != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "invalid request body",
		})
		return
	}

	// 查询用户表，找到对应的用户id记录
	var user User
	result := db.First(&user, updateUserAPIJson.ID)
	if result.Error != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "invalid userid",
		})
		return
	}

	// 更新用户的API值和API状态
	user.Api = updateUserAPIJson.Api
	// 检查API是否可用
	if ValidateUserAPI(user.Api) == false {
		user.ApiStatus = 0
		c.JSON(200, Message{
			Status: "error",
			Info:   "invalid api",
		})
		return
	} else {
		user.ApiStatus = 1
	}

	result = db.Updates(&user)
	if result.Error != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "invalid db error",
		})
		return
	}

	c.JSON(200, Message{
		Status: "ok",
		Info:   "update api success",
	})
	return
}
