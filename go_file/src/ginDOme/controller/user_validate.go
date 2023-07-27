package controller

import (
	"gindome/models"

	"github.com/go-playground/validator/v10"
)

/*
* @description: 校验注册账号数据

* @param: account string 账号

* @param: password string 密码

* @param: rePassword string 重复密码

 * @return: error 错误信息.
 */
func registerUserValid(account, password, rePassword string) error {
	// 1. 创建一个validator实例
	var valid *validator.Validate = validator.New()

	// 2. 创建一个用户实例
	var user *models.User = &models.User{
		Account:    account,
		Password:   password,
		RePassword: rePassword,
	}

	// 3. 校验用户实例
	err := valid.Struct(user)

	// 4. 返回错误信息
	return err
}

/*
* @description: 校验登录账号数据

* @param: account string 账号

* @param: password string 密码

 * @return: error 错误信息.
 */
func loginUserValid(account, password string) error {
	// 1. 创建一个validator实例
	var valid *validator.Validate = validator.New()

	// 2. 创建一个用户实例
	var user *models.User = &models.User{
		Account:    account,
		Password:   password,
		RePassword: password,
	}

	// 3. 校验用户实例
	err := valid.Struct(user)

	// 4. 返回错误信息
	return err
}

func UserProfileValid(u *models.UserProfile) error {
	// 1. 创建一个validator实例
	var valid *validator.Validate = validator.New()

	// 2. 创建一个用户实例
	var user *models.UserProfile = &models.UserProfile{
		Name: u.Name,
		Age: u.Age,
		Sex: u.Sex,
		Number: u.Number,
		Address: u.Address,
		IdCard: u.IdCard,
		Email: u.Email,
	}

	// 3. 校验用户实例
	err := valid.Struct(user)

	// 4. 返回错误信息
	return err
}