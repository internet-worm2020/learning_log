package router

import (
	"gindome/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	//r.Use(middleware.AuthMiddleware())
	SetupApiRouters(r)
	return r
}
