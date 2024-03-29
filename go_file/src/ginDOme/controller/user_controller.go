package controller

import (
	"fmt"
	"gindome/db/redis"
	"gindome/models"
	"gindome/pkg"
	"gindome/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// RegisterHandler 注册账户
//
//	@Summary		注册账户
//	@Description	注册账户
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			account		body	string	true	"账户"
//	@Param			password	body	string	true	"密码"
//	@Param			re_password	body	string	true	"确认密码"
//	@Router			/register [post]
func RegisterHandler(c *gin.Context) {
	// 注册用户
	u := &models.User{}

	// 检查参数是否合法
	if err := c.ShouldBind(u); err != nil {
		ResponseError(c, pkg.CodeInvalidParam)
		return
	}

	// 检查账户是否合法
	if err := registerUserValid(u.Account, u.Password, u.RePassword); err != nil {
		ResponseError(c, pkg.CodeInvalidParam)
		return
	}
	data, err := service.RegisterUserService(u)
	if data == nil {
		ResponseErrorWithMsg(c, err.BusinessCode, err.Message)
		return
	}
	ResponseSuccess(c, data)
}

// LoginHandler 登录账号
//
//	@Summary		登录账号
//	@Description	登录账号
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			account		body	string	true	"账户"
//	@Param			password	body	string	true	"密码"
//	@Router			/login [post]
func LoginHandler(c *gin.Context) {
	u := &models.User{}

	// 检查参数是否合法
	if err := c.ShouldBind(u); err != nil {
		fmt.Println(err.Error())
		ResponseError(c, pkg.CodeInvalidParam)
		return
	}

	// 检查账户是否合法
	if err := loginUserValid(u.Account, u.Password); err != nil {
		ResponseError(c, pkg.CodeInvalidParam)
		return
	}

	// 登录用户
	data, err := service.LoginUserService(u)
	if data == nil {
		ResponseErrorWithMsg(c, err.BusinessCode, err.Message)
		return
	}
	ResponseSuccess(c, data)
}

// GetUserDetailHandler 获取用户信息
//
//	@Summary		获取用户信息
//	@Description	获取用户信息
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"User ID"
//	@Router			/user/{id} [get]
func GetUserDetailHandler(c *gin.Context) {
	// 从URL参数中获取用户ID
	userIdStr := c.Param("id")
	// 将字符串类型的用户ID转换为uint64类型
	userIdInt, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		// 如果转换失败，返回参数错误
		ResponseError(c, pkg.CodeInvalidParam)
	}

	// 根据用户ID获取用户信息
	data, err := service.GetUserByIdService(userIdInt)
	if err != nil {
		// 如果获取用户信息失败，返回服务器繁忙错误
		ResponseError(c, pkg.CodeServerBusy)
		return
	}
	// 返回用户信息
	ResponseSuccess(c, data)
}

// GetUserHandler 处理获取用户列表的请求
//
//	@Summary		获取用户列表
//	@Description	获取用户列表
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int	true	"页码"
//	@Param			size	query		int	true	"每页数量"
//	@Success		200		{object}	models.User
//	@Failure		500		{object}	pkg.ResCode
//	@Router			/user [get]
func GetUserHandler(c *gin.Context) {
	// 获取分页信息
	page, size := GetPageInfo(c)
	// 调用 service 层获取用户列表
	data, err := service.GetUserListService(page, size)
	// 如果出错，返回服务器繁忙
	if err != nil {
		ResponseError(c, pkg.CodeServerBusy)
		return
	}
	// 返回成功响应
	ResponseSuccess(c, data)
}

/*
@param:
*/
func DeleteUserHandler(c *gin.Context) {
	// 获取签名的string
	token := pkg.Token{Token: strings.Split(c.GetHeader("Authorization"), " ")[1]}
	// 调用删除用户服务
	err := service.DeleteUserService(token)
	// 返回错误信息
	if err != nil {
		ResponseErrorWithMsg(c, err.BusinessCode, err.Message)
		return
	}
	// 返回操作成功
	ResponseOperateSuccess(c)
}

func UpdateUserHandler(c *gin.Context) {
	// 获取签名的string
	token := pkg.Token{Token: strings.Split(c.GetHeader("Authorization"), " ")[1]}

	// 注册用户
	userProfile:=&models.UserProfile{}

	// 检查参数是否合法
	if err := c.ShouldBind(userProfile); err != nil {
			ResponseError(c, pkg.CodeInvalidParam)
			return
		}
		fmt.Println(userProfile)
	// 检查账户详情是否合法
	if err := UserProfileValid(userProfile); err != nil {
			ResponseError(c, pkg.CodeInvalidParam)
			return
		}
	if err := service.UpdateUserProfileService(token,userProfile);err != nil {
		ResponseErrorWithMsg(c, err.BusinessCode, err.Message)
		return
	}
	ResponseOperateSuccess(c)
}
func A(c *gin.Context) {
	redis0, _ := redis.GetRedis(0)
	redis0.Set("username", "zhangsan", 0).Err()
	username, _ := redis0.Get("username").Result()
	fmt.Println(username) // zhangsan
	panic("An unexpected error happen!")
}
