package controller

import (
	"fmt"
	"gindome/models"
	"gindome/pkg"
	"gindome/repository"
	"gindome/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator/v10"
	"strconv"
)

//func GetUserDetailHandler(c *gin.Context) {
//	var userIdStr string = c.Param("id")
//	var userIdInt uint64
//	var err error
//	userIdInt, err = strconv.ParseUint(userIdStr, 10, 64)
//	if err != nil {
//		pkg.ResponseError(c, pkg.CodeInvalidParam)
//	}
//	data, err := repository.GetUserById(userIdInt)
//	if err != nil {
//		pkg.ResponseError(c, pkg.CodeServerBusy)
//		return
//	}
//	pkg.ResponseSuccess(c, data)
//}
//
//func GetUserHandler(c *gin.Context) {
//	page, size := pkg.GetPageInfo(c)
//	data, err := repository.GetUserList(page, size)
//	if err != nil {
//		pkg.ResponseError(c, pkg.CodeServerBusy)
//		return
//	}
//	pkg.ResponseSuccess(c, data)
//}
//
//func CreateUserHandler(c *gin.Context) {
//	u := new(models.User)
//	if err := c.ShouldBindJSON(&u); err != nil {
//		pkg.ResponseError(c, pkg.CodeInvalidParam)
//		return
//	}
//	if err := repository.CreateUser(u); err != nil {
//		pkg.ResponseError(c, pkg.CodeServerBusy)
//		return
//	}
//	pkg.ResponseSuccess(c, nil)
//}

// RegisterHandler 注册账户
func RegisterHandler(c *gin.Context) {
	var u *models.User = new(models.User)

	if err := c.ShouldBind(&u); err != nil {
		pkg.ResponseError(c, pkg.CodeInvalidParam)
		return
	}

	if err := registerUserValid(u.Account, u.Password, u.RePassword); err != nil {
		pkg.ResponseError(c, pkg.CodeInvalidParam)
		return
	}

	data, err := service.RegisterUserService(u)
	if data == nil {
		pkg.ResponseErrorWithMsg(c, err.BusinessCode, err.Message)
		return
	}
	pkg.ResponseSuccess(c, data)

}

// LoginHandler 登录账号
func LoginHandler(c *gin.Context) {
	var u *models.User = new(models.User)
	if err := c.ShouldBind(&u); err != nil {
		fmt.Println(err.Error())
		pkg.ResponseError(c, pkg.CodeInvalidParam)
		return
	}

	if err := loginUserValid(u.Account, u.Password); err != nil {
		pkg.ResponseError(c, pkg.CodeInvalidParam)
		return
	}
	data, err := service.LoginUserService(u)
	if data == nil {
		pkg.ResponseErrorWithMsg(c, err.BusinessCode, err.Message)
		return
	}
	pkg.ResponseSuccess(c, data)

}

func GetUserDetailHandler(c *gin.Context) {
	var userIdStr string = c.Param("id")
	var userIdInt uint64
	var err error
	userIdInt, err = strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		pkg.ResponseError(c, pkg.CodeInvalidParam)
	}
	data, err := repository.GetUserById(userIdInt)
	if err != nil {
		pkg.ResponseError(c, pkg.CodeServerBusy)
		return
	}
	pkg.ResponseSuccess(c, data)
}
