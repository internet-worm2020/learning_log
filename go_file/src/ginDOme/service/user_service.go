package service

import (
	"gindome/models"
	"gindome/pkg"
	"gindome/repository"
)

func RegisterUserService(u *models.User) (string, pkg.Error) {
	// 检查用户是否已经存在
	_, totalData, err := repository.GetAccount(u.Account)
	if err != nil {
		return "", pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}
	if totalData == 1 {
		return "", pkg.NewErrorAutoMsg(pkg.CodeUserExist)
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
		return "", pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	// 获取新用户的ID
	UId, err := repository.GetIDByAccount(u.Account)
	if err != nil {
		return "", pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	// 生成认证令牌
	token, err := pkg.GetToken(UId)
	if err != nil {
		return "", pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	// 返回结果和错误信息
	return token, pkg.NewErrorAutoMsg(pkg.CodeSuccess)
}

func LoginUserService(u *models.User) (string, pkg.Error) {
	// 哈希加密用户密码
	u.HashPassword()
	// 检查用户是否已经存在
	user, totalData, err := repository.GetAccount(u.Account)
	if err != nil {
		return "", pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}
	if totalData != 1 {
		return "", pkg.NewErrorAutoMsg(pkg.CodeUserNotExist)
	}
	// 比较用户输入的账号和密码是否与数据库中的记录匹配
	if user.Account != u.Account || user.Password != u.Password {
		return "", pkg.NewErrorAutoMsg(pkg.CodeInvalidPassword)
	}
	// 生成认证令牌
	token, err := pkg.GetToken(user.ID)
	if err != nil {
		return "", pkg.NewErrorAutoMsg(pkg.CodeServerBusy).WithErr(err)
	}

	return token, pkg.NewErrorAutoMsg(pkg.CodeSuccess)
}
