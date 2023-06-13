package pkg

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 2

type Claims struct {
	UId uint `json:"uid,omitempty"`
	Account string `json:"account,omitempty"`
	jwt.StandardClaims
}

type Token struct {
	Token string `json:"token,omitempty"`
}

// GetToken 生成 token
func GetToken(uid uint, account string) (*Token, error) {
    // 设置 token 的 claims
    claims := Claims{
        uid,
        account,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
            Issuer:    "lx-jwt",
        },
    }
    // 生成 token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    // 设置 secretKey
    secretKey := []byte("my-secret-key")
    // 签名 token
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return nil, err
    }
    return &Token{tokenString}, nil
}

// ParseToken 解析JWT令牌
func ParseToken(tokenString string) (*Claims, Error) {
	// 设置密钥
	secretKey := []byte("my-secret-key")
	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	// 如果解析失败，返回错误
	if err != nil {
		return nil, NewErrorAutoMsg(CodeInvalidToken).WithErr(err)
	}
	// 如果解析成功，返回Claims
	if claims, ok := token.Claims.(*Claims); ok {
		return claims, NewErrorAutoMsg(CodeSuccess)
	}
	// 如果解析成功但Claims类型不正确，返回错误
	return nil, NewErrorAutoMsg(CodeInvalidToken)
}
