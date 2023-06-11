package middleware

import (
	// "bamboo.com/pipeline/Go-assault-squad/controller"
	"fmt"
	"gindome/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("中间件开始执行了")
		requestURL := c.Request.URL
		urlSlice := strings.Split(requestURL.String(), "/")
		if urlSlice[1] == "index" || urlSlice[1] == "login" || urlSlice[1] == "register" {
			c.Next()
			return
		}
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			pkg.ResponseError(c, pkg.CodeTokenIsEmpty)
			c.Abort()
			return
		}

		//pkg.ParseToken(authHeader)
		//fmt.Printf("authHeader: %v\n", authHeader)
	}
}
