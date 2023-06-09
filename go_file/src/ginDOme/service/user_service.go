package service

import (
	"gindome/models"
	"gindome/repository"
)

func RegisterUserService(u *models.User) (error) {
	var user models.User = models.User{
		Account:  u.Account,
		Password: u.Password,
		UserProfile:models.UserProfile{
			Name:u.Account,
		},
	}
	var err error = repository.RegisterUser(&user)
	return err
}
