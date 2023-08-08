package middleware

import (
	"gindome/pkg/snowflake"

	"github.com/gin-gonic/gin"
)
func AddUniqueNumberMiddleware()func(c *gin.Context){
	return func(c *gin.Context){
	requestID:=snowflake.GenID()
	c.Set("request_id",requestID)
	}
}