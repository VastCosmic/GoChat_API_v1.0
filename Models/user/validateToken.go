package user

import (
	"Gin/Models/DB"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

// 定义一个验证JWT的函数
func validateToken(tokenString string) (bool, error) {
	// 解析token字符串，得到一个token对象
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return false, err
	}

	// 检查token是否有效和过期
	if !token.Valid || token.Claims.(jwt.MapClaims)["exp"].(float64) < float64(time.Now().Unix()) {
		return false, nil
	}

	// 检查数据库中是否仍存在该token
	// 连接数据库
	db := DB.Connect()
	fmt.Println(db)

	// 将token字符串分割成三部分：头部、负载和签名
	tokenParts := strings.Split(tokenString, ".")
	// 定义一个JwtToken结构体变量
	var jwtToken JwtToken
	// 在数据库中查询JwtToken，使用头部、负载和签名作为条件
	result := db.Where("header = ? AND payload = ? AND signature = ?", tokenParts[0], tokenParts[1], tokenParts[2]).First(&jwtToken)

	// 如果查询出错或者没有找到记录，返回false和错误信息
	if result.Error != nil {
		return false, result.Error
	}

	// 如果查询成功，返回true和nil
	return true, nil
}
