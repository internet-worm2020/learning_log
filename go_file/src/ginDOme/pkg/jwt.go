package pkg

import (
	"errors"
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

func ParseToken(tokenString string) (*Claims, error) {
	secretKey := []byte("my-secret-key")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效令牌")
}
