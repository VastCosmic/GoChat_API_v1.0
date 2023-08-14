package user

import (
	"os"
	"time"
)

// User 定义用户模型
type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Pwd       string `json:"pwd"`
	Admin     uint   `json:"admin"`
	Name      string `json:"name"`
	Api       string `json:"api"`
	ApiStatus uint   `json:"api_status"`
}

// RegisterUserJson 定义注册用户模型
type RegisterUserJson struct {
	//ID        uint   `json:"id" gorm:"primaryKey"`
	Pwd string `json:"pwd"`
	//Admin     uint   `json:"admin"`
	Name string `json:"name"`
	//Api  string `json:"api"`
	//ApiStatus uint   `json:"api_status"`
}

// LoginUserJson 定义登录用户模型
type LoginUserJson struct {
	ID  uint   `json:"id" gorm:"primaryKey"`
	Pwd string `json:"pwd"`
	//Admin     uint   `json:"admin"`
	Name string `json:"name"`
	//Api  string `json:"api"`
	//ApiStatus uint   `json:"api_status"`
}

// UpdateUserInfoJson 定义更新用户模型
type UpdateUserInfoJson struct {
	ID  uint   `json:"id" gorm:"primaryKey"`
	Pwd string `json:"pwd"`
	//Admin     uint   `json:"admin"`
	Name string `json:"name"`
	Api  string `json:"api"`
	//ApiStatus uint   `json:"api_status"`
}

// CurrentUserJson 定义查看当前用户信息的模型
type CurrentUserJson struct {
	ID uint `json:"id" gorm:"primaryKey"`
	//Pwd string `json:"pwd"`
	Admin uint   `json:"admin"`
	Name  string `json:"name"`
	//Api  string `json:"api"`
	ApiStatus uint `json:"api_status"`
}

// UpdateAPIUserJson 定义更新API用户模型
type UpdateAPIUserJson struct {
	ID uint `json:"id" gorm:"primaryKey"`
	//Pwd string `json:"pwd"`
	//Admin uint   `json:"admin"`
	//Name  string `json:"name"`
	Api string `json:"api"`
	//ApiStatus uint `json:"api_status"`
}

// ChatUserJson 定义更新API用户模型
type ChatUserJson struct {
	ID uint `json:"id" gorm:"primaryKey"`
	//Pwd string `json:"pwd"`
	Admin     uint   `json:"admin"`
	Name      string `json:"name"`
	Api       string `json:"api"`
	ApiStatus uint   `json:"api_status"`
}

// JwtToken 定义jwt_tokens模型
type JwtToken struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserID    uint   `json:"user_id"`
	Header    string `json:"header"`
	Payload   string `json:"payload"`
	Signature string `json:"signature"`
	CreatedAt time.Time
	ExpiresAt time.Time
}

// Message 定义Message结构体
type Message struct {
	Status string      `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data,omitempty"`
}

// 定义一个签名密钥，从环境变量文件中加载
var secretKey = []byte(os.Getenv("SECRET_KEY"))
