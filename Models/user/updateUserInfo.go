package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UpdateUserInfo 定义UpdateUserInfo函数,用于更改用户名、密码,修改密码的操作会使所有token失效
func UpdateUserInfo(c *gin.Context) {
	// 连接数据库，这里假设已经存在一个名为db的*gorm.DB对象
	db := c.MustGet("db").(*gorm.DB)

	// 取得用户ID以及新的用户信息
	var updateUserInfoJson UpdateUserInfoJson
	err := c.ShouldBindJSON(&updateUserInfoJson)
	if err != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "invalid request body",
		})
		return
	}

	// 查询用户表，找到对应的用户id记录
	var user User
	result := db.First(&user, updateUserInfoJson.ID)
	if result.Error != nil {
		c.JSON(200, Message{
			Status: "error",
			Info:   "invalid userid",
		})
		return
	}

	// 更新用户的用户名
	if updateUserInfoJson.Name != "" && user.Name != updateUserInfoJson.Name {
		user.Name = updateUserInfoJson.Name
	}

	// 更新用户的密码
	if updateUserInfoJson.Pwd != "" && user.Pwd != updateUserInfoJson.Pwd {
		user.Pwd = updateUserInfoJson.Pwd
		// 使所有token失效
		// 查询jwt表，找到对应的用户id记录
		var jwtToken JwtToken
		// 在数据库中删除user_id为1的所有JwtToken记录
		result = db.Where("user_id = ?", user.ID).Delete(&jwtToken)
		if result.Error != nil {
			fmt.Println("删除失败:", result.Error)
			c.JSON(200, Message{
				Status: "error",
				Info:   "update pwd failed",
			})
			return
		} else {
			// 获取受影响的行数
			fmt.Println("删除了", result.RowsAffected, "个token记录")
		}
	}

	// 检查API是否更改，未更改则不做检测
	if user.Api != "" && user.Pwd != updateUserInfoJson.Api {
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
		// 更新用户的API值和API状态
		user.Api = updateUserInfoJson.Api
	}

	// 保存更改
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
		Info:   "update user info success",
	})
	return
}
