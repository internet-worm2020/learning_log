package repository

import (
	"gindome/db/mysqlDB"
	"gindome/models"
)

func GetUserById(userId uint64) (models.User, error) {
	var user models.User = models.User{}
	var err error = mysqlDB.GetDB().Preload("UserProfile").Where("id = ?", userId).Take(&user).Error
	return user, err
}

func GetUserList(page, size int) ([]models.User, error) {
	var user []models.User = []models.User{}
	var err error = mysqlDB.GetDB().Preload("UserProfile").Limit(size).Offset((page - 1) * size).Find(&user).Error
	return user, err
}

func CreateUser(u *models.User) error {
	var err error = mysqlDB.GetDB().Create(u).Error
	return err
}
