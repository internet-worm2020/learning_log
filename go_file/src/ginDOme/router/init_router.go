package router

import (
	"gindome/middleware"

	"github.com/gin-gonic/gin"
)

/*
InitRouter

@description: 初始化路由

@return: *gin.Engine gin引擎.
*/
func InitRouter() *gin.Engine {
	// 1. 创建gin引擎
	r := gin.New()
	// 2. 添加日志中间件
	// r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.LoggerMiddleware1())
	// 3. 添加认证中间件
	r.Use(middleware.AuthMiddleware())
	// 4. 注册Swagger
	registerSwagger(r)
	// 5. 设置API路由
	setupApiRouters(r)
	// 6. 返回gin引擎
	return r
}
