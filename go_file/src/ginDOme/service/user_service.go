package service

import (
	"gindome/models"
	"gindome/pkg"
	"gindome/repository"
)

// RegisterUserService 注册用户服务
func RegisterUserService(u *models.User) (*pkg.Token, pkg.Error) {
	// 检查用户是否已经存在
	_, totalData, err := repository.GetAccount(u.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}
	if totalData == 1 {
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
	if err := repository.RegisterUser(&user); err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	// 获取新用户的ID
	uId, err := repository.GetIDByAccount(u.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	// 生成认证令牌
	token, err := pkg.GetToken(uId, u.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	// 返回结果和错误信息
	return token, pkg.NewErrorAutoMsg(pkg.CodeSuccess)
}

// LoginUserService 登录用户服务
func LoginUserService(u *models.User) (*pkg.Token, pkg.Error) {
	// 哈希加密用户密码
	u.HashPassword()
	// 检查用户是否已经存在
	user, totalData, err := repository.GetAccount(u.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}
	if totalData != 1 {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeUserNotExist)
	}
	// 比较用户输入的账号和密码是否与数据库中的记录匹配
	if user.Account != u.Account || user.Password != u.Password {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeInvalidPassword)
	}
	// 生成认证令牌
	token, err := pkg.GetToken(user.ID, user.Account)
	if err != nil {
		return nil, pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	return token, pkg.NewErrorAutoMsg(pkg.CodeSuccess)
}

// GetUserByIdService 获取用户信息
func GetUserByIdService(userId uint64) (*models.UserProfile, error) {
	// 调用 repository 层获取用户信息
	data, err := repository.GetUserById(userId)
	if err != nil {
		return nil, err // 返回错误信息
	}
	return data, nil // 返回用户信息和 nil
}

// GetUserList 获取用户列表
func GetUserListService(page, size int) ([]*models.User, error) {
	// 调用 repository 层获取用户列表
	data, err := repository.GetUserList(page, size)
	// 如果出错，返回错误信息
	if err != nil {
		return nil, err // 返回值为 nil 和错误信息
	}
	return data, nil // 返回用户列表和 nil
}
