package controller

import (
	"gindome/models"
	"github.com/go-playground/validator/v10"
)

// 注册账号数据校验
func registerUserValid(account, password, rePassword string) error {
	var valid *validator.Validate = validator.New()

	var user *models.User = &models.User{
		Account:    account,
		Password:   password,
		RePassword: rePassword,
	}

	err := valid.Struct(user)

	return err
}

// 登录账号数据校验
func loginUserValid(account, password string) error {
	var valid *validator.Validate = validator.New()

	var user *models.User = &models.User{
		Account:    account,
		Password:   password,
		RePassword: password,
	}

	err := valid.Struct(user)

	return err
}
