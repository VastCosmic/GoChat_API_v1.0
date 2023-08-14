package main

import (
	"Gin/Models/DB"
	"Gin/Models/user"
	"Gin/Services"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // 使用gin默认的路由器来处理HTTP请求

	// 连接数据库
	db := DB.Connect()
	fmt.Println(db)
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	//router.Use(cors.Default()) // 允许跨域请求

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Authorization", "content-type"},
	}))

	users := router.Group("/users") // 创建一个users路由组
	{
		users.POST("/register", Service.RegisterUser)
		users.POST("/login", Service.LoginUser)
		users.GET("/logout", user.JwtMiddleware, Service.LogoutUser)
		users.GET("/userinfo", user.JwtMiddleware, Service.GetCurrentUser)
		users.POST("/update", user.JwtMiddleware, Service.UpdateUser)
	}

	chat := router.Group("/chat") // 创建一个chat路由组
	{
		chat.POST("/history", user.JwtMiddleware, Service.GetChatHistory)   // 获取聊天记录
		chat.POST("/update_api", user.JwtMiddleware, Service.UpdateUserAPI) //处理并更新api
		chat.POST("/", user.JwtMiddleware, Service.Chat)                    // 处理通过POST方法向/chat路径发出的请求,处理聊天对话
	}

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("HTTP服务器启动失败:", err)
		return
	} // 启动HTTP服务器并监听8080端口
}
