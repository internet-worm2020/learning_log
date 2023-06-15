package controller

import (
	"fmt"
	"gindome/db/redis"
	"gindome/models"
	"gindome/pkg"
	"gindome/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RegisterHandler 注册账户
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
		pkg.ResponseError(c, pkg.CodeInvalidParam)
		return
	}

	// 检查账户是否合法
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
		pkg.ResponseError(c, pkg.CodeInvalidParam)
		return
	}

	// 检查账户是否合法
	if err := loginUserValid(u.Account, u.Password); err != nil {
		pkg.ResponseError(c, pkg.CodeInvalidParam)
		return
	}

	// 登录用户
	data, err := service.LoginUserService(u)
	if data == nil {
		pkg.ResponseErrorWithMsg(c, err.BusinessCode, err.Message)
		return
	}
	pkg.ResponseSuccess(c, data)

}

// GetUserDetailHandler 获取用户信息
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
		pkg.ResponseError(c, pkg.CodeInvalidParam)
	}

	// 根据用户ID获取用户信息
	data, err := service.GetUserByIdService(userIdInt)
	if err != nil {
		// 如果获取用户信息失败，返回服务器繁忙错误
		pkg.ResponseError(c, pkg.CodeServerBusy)
		return
	}
	// 返回用户信息
	pkg.ResponseSuccess(c, data)
}

// GetUserHandler 处理获取用户列表的请求
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
	page, size := pkg.GetPageInfo(c)
	// 调用 service 层获取用户列表
	data, err := service.GetUserListService(page, size)
	// 如果出错，返回服务器繁忙
	if err != nil {
		pkg.ResponseError(c, pkg.CodeServerBusy)
		return
	}
	// 返回成功响应
	pkg.ResponseSuccess(c, data)
}


func A(c *gin.Context){
	// redis.RedisClient.Set("username", "zhangsan", 0).Err()
	username, _ := redis.GetRedis().Get("username").Result()
	fmt.Println(username)  // zhangsan
}