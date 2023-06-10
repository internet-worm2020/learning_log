package service

import (
	"fmt"
	"gindome/models"
	"gindome/repository"
)

func RegisterUserService(u *models.User) error {
	var totalData int64
	var err error
	totalData, err = repository.GetAccount(u.Account)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
		fmt.Println(2)
	}
	if totalData == 1 {
		return fmt.Errorf("账号已存在")
	}
	var user models.User = models.User{
		Account:  u.Account,
		Password: u.Password,
		UserProfile: models.UserProfile{
			Name: u.Account,
		},
	}
	err = repository.RegisterUser(&user)
	return err
}
