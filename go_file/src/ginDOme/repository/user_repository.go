package repository

import (
	"gindome/db/mysqlDB"
	"gindome/models"
)

// RegisterUser 注册用户，添加数据到数据库
func RegisterUser(u *models.User) error {
	// 将用户信息添加到数据库
	if err := mysqlDB.GetDB().Create(u).Error; err != nil {
		return sqlError(err)
	}
	return nil
}

// GetAccount 根据账户名查找用户
func GetAccount(account string) (*models.User, int64, error) {
	// 定义一个用户变量
	var user models.User
	// 根据账户名查找用户
	db := mysqlDB.GetDB().Where("account=?", account).Find(&user)
	// 如果查找出错，则返回错误信息
	if db.Error != nil {
		return nil, 0, sqlError(db.Error)
	}
	// 返回用户信息和影响的行数
	return &user, db.RowsAffected, nil
}

// GetIDByAccount 根据账户名获取用户ID
func GetIDByAccount(account string) (uint, error) {
	// 定义一个ID变量
	var Id uint
	// 根据账户名获取用户ID
	if err := mysqlDB.GetDB().Table("user").Select("id").Where("account=?", account).Take(&Id).Error; err != nil {
		return 0, err
	}
	// 返回用户ID
	return Id, nil
}

// GetUserById 根据用户ID获取用户信息
func GetUserById(userId uint64) (*models.UserProfile, error) {
	// 定义一个用户变量和用户信息变量
	var user models.User
	var userProfile models.UserProfile
	// 根据用户ID获取用户信息
	if err := mysqlDB.GetDB().Where("id=?", userId).Take(&user).Error; err != nil {
		return nil, err
	}
	// 根据用户信息获取用户信息变量
	if err := mysqlDB.GetDB().Model(&user).Association("UserProfile").Find(&userProfile); err != nil {
		return nil, err
	}
	// 返回用户信息变量
	return &userProfile, nil
}

// GetUserConsistent 查询用户数据是否一致
func GetUserConsistent(uId uint, uAccount string) (uint, string, error) {
	// 定义一个用户变量
	var user struct {
		ID      uint   `gorm:"column:id"`
		Account string `gorm:"column:account"`
	}
	// 查询用户数据是否一致
	if err := mysqlDB.GetDB().Table("user").Select("id, account").Where("id = ? AND account = ?", uId, uAccount).Take(&user).Error; err != nil {
		return 0, "", err
	}
	// 返回用户ID和账户名
	return user.ID, user.Account, nil
}

// GetUserList 获取用户列表
func GetUserList(page, size int) ([]*models.User, error) {
	// 初始化一个空的用户列表
	users := make([]*models.User, 0, size)
	// 查询用户列表
	err := mysqlDB.GetDB().Preload("UserProfile").Limit(size).Offset((page - 1) * size).Find(&users).Error
	// 返回用户列表和错误信息
	return users, err
}
