package user

import (
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

// 定义一个生成JWT的函数
func generateToken(userID uint) (string, error) {
	// 创建一个JWT标准声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        strconv.Itoa(int(userID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	// 对JWT进行签名，得到一个字符串形式的token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
