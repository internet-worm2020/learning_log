package router

import (
	"gindome/controller"
	docs "gindome/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*
@description: 设置API路由

@param: r *gin.Engine gin引擎.
*/
func setupApiRouters(r *gin.Engine) {
	// 注册
	r.POST("/register", controller.RegisterHandler)
	// 登录
	r.POST("/login", controller.LoginHandler)
	// 根据id查询用户详情信息
	r.GET("/user/:id", controller.GetUserDetailHandler)
	// 查询用户列表信息
	r.GET("/user", controller.GetUserHandler)
	// 根据id修改用户详情信息
	r.PATCH("/user",controller.UpdateUserHandler)
	// 根据id删除用户
	r.DELETE("/user",controller.DeleteUserHandler)
	r.GET("/a", controller.A)
	r.GET("/ws",controller.Ws)
}

/*
@description: 注册Swagger

@param: r *gin.Engine gin引擎.
*/
func registerSwagger(r *gin.Engine) {
	// 1. 设置Swagger基本信息
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "管理后台接口"
	docs.SwaggerInfo.Description = "实现一个管理系统的后端API服务"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8084"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// 2. 注册Swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
