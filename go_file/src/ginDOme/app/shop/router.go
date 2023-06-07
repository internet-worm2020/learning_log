package shop

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/shop/:name/*action", shopGetHandler)
	e.POST("/shop",shopPosthandler)
}
