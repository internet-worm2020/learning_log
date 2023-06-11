package service

import (
	"gindome/models"
	"gindome/pkg"
	"gindome/repository"
)

func RegisterUserService(u *models.User) (string, pkg.Error) {
	// 检查用户是否已经存在
	totalData, err := repository.GetAccount(u.Account)
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
