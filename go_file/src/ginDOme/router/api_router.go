package router

import (
	"gindome/controller"
	"github.com/gin-gonic/gin"
)

func SetupApiRouters(r *gin.Engine) {
	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/user/:id", controller.GetUserDetailHandler)
	r.GET("/user", controller.GetUserHandler)
	//r.POST("/user", controller.CreateUserHandler)
}
