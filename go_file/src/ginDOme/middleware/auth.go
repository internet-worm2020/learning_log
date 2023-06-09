package middleware

import (
	// "bamboo.com/pipeline/Go-assault-squad/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// token验证成功，返回c.Next继续，否则返回c.Abort()直接返回
		// authHeader := c.Request.Header.Get("Authorization")
		// if authHeader == "" {
		// 	controller.ResponseError(c, controller.CodeNeedLogin)
		// 	c.Abort()
		// 	return
		// }
		// 将当前请求的userID信息保存到请求的上下文c上
		//c.Set(controller.CtxUserIDKey, mc.UserID)
		// c.Next() // 后续的处理请求的函数中 可以用过c.Get(CtxUserIDKey) 来获取当前请求的用户信息
		fmt.Println("中间件开始执行了")
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Printf("authHeader: %v\n", authHeader)
	}
}