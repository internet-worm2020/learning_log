package repository

import (
	"gindome/db/mysqlDB"
	"gindome/models"
)

func RegisterUser(u *models.User) error {
	var err error = sqlError(mysqlDB.GetDB().Create(u).Error)
	return err
}

func GetAccount(account string) (int64, error) {
	var user models.User
	db := mysqlDB.GetDB().Where("account=?", account).Find(&user)
	var totalData int64 = db.RowsAffected
	var err error = sqlError(db.Error)
	return totalData, err
}

func GetUserById(userId uint64) (*models.UserProfile, error) {
	var user models.User
	var userProfile models.UserProfile
	var err error
	err = mysqlDB.GetDB().Where("id=?", userId).Take(&user).Error
	if err != nil {
		return nil, err
	}
	err = mysqlDB.GetDB().Model(&user).Association("UserProfile").Find(&userProfile)
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}

func GetUserList(page, size int) ([]models.User, error) {
	var user []models.User
	var err error = mysqlDB.GetDB().Preload("UserProfile").Limit(size).Offset((page - 1) * size).Find(&user).Error
	return user, err
}

func CreateUser(u *models.User) error {
	var err error = mysqlDB.GetDB().Create(u).Error
	return err
}
