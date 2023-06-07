package users

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/user/:id", getUserDetailHandler)
	e.GET("/user",getUserHandler)
	e.POST("/user",createUserHandler)
}
