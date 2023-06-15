package router

import (
	"gindome/controller"
	docs "gindome/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupApiRouters(r *gin.Engine) {
	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/user/:id", controller.GetUserDetailHandler)
	r.GET("/user", controller.GetUserHandler)
	//r.POST("/user", controller.CreateUserHandler)
}

func registerSwagger(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "管理后台接口"
	docs.SwaggerInfo.Description = "实现一个管理系统的后端API服务"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8084"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
