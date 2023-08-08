package middleware

import (
	"gindome/controller"
	"gindome/pkg"
	"gindome/repository"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 基于JWT的认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跳过认证操作的URL切片，根据需求添加
		skipAuthURLs := []string{"index", "login", "register", "swagger", "a"}
		requestURL := c.Request.URL
		// 遍历 skipAuthURLs 切片，如果请求URL包含其中任意一个字符串，则跳过认证步骤
		for _, skipURL := range skipAuthURLs {
			if strings.Contains(requestURL.String(), skipURL) {
				c.Next()
				return
			}
		}
		// 校验是否携带token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, pkg.CodeNeedLogin)
			c.Abort()
			return
		}

		// 解析 Authorization 头并校验 token 是否正确
		token, pkgErr := parseToken(authHeader)
		if token == nil {
			controller.ResponseErrorWithMsg(c, pkgErr.BusinessCode, pkgErr.Message)
			c.Abort()
			return
		}
		// 根据token中的uId和account获取用户信息
		id, account, err := repository.GetUserConsistent(token.UId, token.Account)
		if err != nil {
			pkgErr := pkg.NewErrorAutoMsg(pkg.CodeSuccess).WithErr(err)
			controller.ResponseErrorWithMsg(c, pkgErr.BusinessCode, pkgErr.Message)
			c.Abort()
			return
		}
		// 校验token中的uId和account是否与获取到的用户信息一致
		if id != token.UId || account != token.Account {
			controller.ResponseError(c, pkg.CodeWrongCredentials)
			c.Abort()
			return
		}
		// 信息保存到上下文
		c.Set("uId", token.UId)
		c.Set("account", token.Account)
		c.Next()
	}
}

// 解析 token 并验证是否有效
func parseToken(authHeader string) (*pkg.Claims, *pkg.Error) {
	// 检查 token 格式是否正确
	tokenSlice := strings.SplitN(authHeader, " ", 2)
	if len(tokenSlice) != 2 || tokenSlice[0] != "Bearer" {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeWrongCredentials)
	}

	// 解析 token 并验证是否有效
	mc, err := pkg.ParseToken(tokenSlice[1])
	if mc == nil {
		return nil, err
	}

	return mc, pkg.NewErrorAutoMsg(pkg.CodeSuccess)
}
