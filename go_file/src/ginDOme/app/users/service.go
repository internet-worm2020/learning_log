package users

import (
	"gindome/db/mysqlDB"
)

func getUserList(page, size int) ([]User, error) {
	var user []User = []User{}
	var err error = mysqlDB.GetDB().Preload("UserProfile").Limit(size).Offset((page - 1) * size).Find(&user).Error
	return user,err
}

func createUser(u *User)(error){
	var err error = mysqlDB.GetDB().Create(u).Error
	return err
}