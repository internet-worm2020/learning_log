package service

import (
	"fmt"
	"gindome/models"
	"gindome/pkg"
	"gindome/repository"
)

/*
RegisterUserService

@description: 注册用户服务

@param: u *models.User 用户信息

@return: *pkg.Token 认证令牌

@return: pkg.Error 错误信息.
*/
func RegisterUserService(u *models.User) (*pkg.Token, pkg.Error) {
	// 检查用户是否已经存在
	totalData, err := repository.GetAccount(u.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}
	// 查询到的数据是否大于0
	if totalData > 0 {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeUserExist)
	}

	// 创建新用户
	user := models.User{
		Account:  u.Account,
		Password: u.Password,
		UserProfile: models.UserProfile{
			Name: u.Account,
		},
	}
	user.HashPassword()

	// 保存新用户到数据库中
	userData, err := repository.RegisterUser(&user)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	// 4. 生成认证令牌
	token, err := pkg.GetToken(userData.ID, userData.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	// 6. 返回结果和错误信息
	return token, pkg.NewErrorAutoMsg(pkg.CodeSuccess)
}

/*
LoginUserService

@description: 登录用户服务

@param: u *models.User 用户信息

@return: *pkg.Token 认证令牌

@return: pkg.Error 错误信息.
*/
func LoginUserService(u *models.User) (*pkg.Token, pkg.Error) {
	// 1. 哈希加密用户密码
	u.HashPassword()

	// 2. 检查用户是否已经存在
	totalData, err := repository.GetAccount(u.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}
	if totalData != 1 {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeUserNotExist)
	}
	// 3. 根据账号和密码查询用户
	user, err := repository.GetAccountPassword(u.Account, u.Password)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeInvalidPassword)
	}
	// 4. 比较用户输入的账号和密码是否与数据库中的记录匹配
	if user.Account != u.Account || user.Password != u.Password {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeInvalidPassword)
	}

	// 4. 生成认证令牌
	token, err := pkg.GetToken(user.ID, user.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	return token, pkg.NewErrorAutoMsg(pkg.CodeSuccess)
}

/*
GetUserByIdService

@description: 根据用户 ID 获取用户信息

@param: userId uint64 用户 ID

@return: *models.UserProfile 用户信息

@return: error 错误信息.
*/
func GetUserByIdService(userId uint64) (*models.UserProfile, error) {
	// 1. 调用 repository 层获取用户信息
	data, err := repository.GetUserById(userId)
	// 2. 如果出现错误，返回错误信息
	if err != nil {
		return nil, err
	}
	// 3. 返回用户信息和 nil
	return data, nil
}

/*
GetUserListService

@description: 获取用户列表服务

@param: page int 分页页码

@param: size int 分页大小

@return: []*models.UserProfile 用户列表

@return: error 错误信息.
*/
func GetUserListService(page, size int) ([]*models.UserProfile, error) {
	// 1. 调用 repository 层获取用户列表
	data, err := repository.GetUserList(page, size)
	// 2. 如果出错，返回错误信息
	if err != nil {
		return nil, err // 返回值为 nil 和错误信息
	}
	// 3. 返回用户列表和 nil
	return data, nil
}

/*
DeleteUserService

@description: 删除用户信息服务

@param: token pkg.Token token信息

@return: string 成功信息

@return: pkg.Error 错误信息
*/
func DeleteUserService(token pkg.Token) (string, pkg.Error) {
	// 定义签名信息
	var claims *pkg.Claims
	// 定义用户id
	var uId uint
	// 定义自定义错误
	var tokenErr pkg.Error
	// 解析token获取签名信息
	claims, tokenErr = pkg.ParseToken(token.Token)
	// 签名是否解析成功
	if claims == nil {
		return "", tokenErr
	}
	uId = claims.UId
	// 调用数据操作删除用户数据
	err := repository.DeleteUser(uId)
	if err != nil {
		return "", pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}
	return "ok", pkg.NewError(pkg.CodeSuccess, pkg.CodeSuccess.Msg())
}
func UpdateUserProfileService(token pkg.Token) {
	// 定义签名信息
	var claims *pkg.Claims
	// 定义用户id
	var uId uint
	// 定义自定义错误
	var tokenErr pkg.Error
	// 解析token获取签名信息
	claims, tokenErr = pkg.ParseToken(token.Token)
	// 签名是否解析成功
	// if claims == nil {
	// 	return "",tokenErr
	// }
	uId = claims.UId
	a, b := repository.UpdateUserProfile(uId)
	fmt.Println(a, b, tokenErr)

}
