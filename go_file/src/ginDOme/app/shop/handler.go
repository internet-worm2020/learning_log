package shop

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_"github.com/gin-gonic/gin/binding"
    _"gopkg.in/go-playground/validator.v8"
)

func shopGetHandler(c *gin.Context) {
	name:=c.Param("name")
	action := c.Param("action")
	fmt.Printf("action: %v\n", action)
	fmt.Printf("name: %v\n", name)
	c.JSON(http.StatusOK, gin.H{
		"message": "Shop Router",
	})
}

func shopPosthandler(c *gin.Context){
	b, _ := c.GetRawData()
	var m map[string]interface{}
		// 反序列化
		_ = json.Unmarshal(b, &m)
		c.JSON(http.StatusOK, m)
}
