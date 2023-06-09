package controller

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"gindome/models"
	"gindome/repository"
	"gindome/service"
)

func GetUserDetailHandler(c *gin.Context) {
	var userIdStr string = c.Param("id")
	var userIdInt uint64
	var err error
	userIdInt, err = strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}
	data, err := repository.GetUserById(userIdInt)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetUserHandler(c *gin.Context) {
	page, size := GetPageInfo(c)
	data, err := repository.GetUserList(page, size)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func CreateUserHandler(c *gin.Context) {
	u := new(models.User)
	if err := c.ShouldBindJSON(&u); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := repository.CreateUser(u); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func RegisterHandler(c *gin.Context){
	var u *models.User=new(models.User)
	if err := c.ShouldBindJSON(&u); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if err := service.RegisterUserService(u); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}