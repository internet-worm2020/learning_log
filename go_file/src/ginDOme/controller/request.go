package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPageInfo(c *gin.Context) (int, int) {
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")

	var (
		page int
		size int
		err  error
	)

	page, err = strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		size = 10
	}
	return page, size
}
