package pkg

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2

type Claims struct {
	UId uint
	jwt.StandardClaims
}

func GetToken(uid uint) (string, error) {

	claims := Claims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "lx-jwt",                                   // 签发人
		},
	}
	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥对令牌进行签名
	secretKey := []byte("my-secret-key")
	tokenString, err := token.SignedString(secretKey)
	fmt.Println("JWT token:", tokenString)
	return tokenString, err
}
