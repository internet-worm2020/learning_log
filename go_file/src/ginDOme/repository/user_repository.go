package repository

import (
	"gindome/db/mysqlDB"
	"gindome/models"
)

// RegisterUser 注册用户，添加数据到数据库
func RegisterUser(u *models.User) error {
    if err := mysqlDB.GetDB().Create(u).Error; err != nil {
        return sqlError(err)
    }
    return nil
}

// GetAccount 根据账户名查找用户
func GetAccount(account string) (*models.User, int64, error) {
    var user models.User
    db := mysqlDB.GetDB().Where("account=?", account).Find(&user)
    if db.Error != nil {
        return nil, 0, sqlError(db.Error)
    }
    return &user, db.RowsAffected, nil
}

// GetIDByAccount 根据账户名获取用户ID
func GetIDByAccount(account string) (uint, error) {
    var Id uint
    if err := mysqlDB.GetDB().Table("user").Select("id").Where("account=?", account).Take(&Id).Error; err != nil {
        return 0, err
    }
    return Id, nil
}

// GetUserById 根据用户ID获取用户信息
func GetUserById(userId uint64) (*models.UserProfile, error) {
    var user models.User
    var userProfile models.UserProfile
    if err := mysqlDB.GetDB().Where("id=?", userId).Take(&user).Error; err != nil {
        return nil, err
    }
    if err := mysqlDB.GetDB().Model(&user).Association("UserProfile").Find(&userProfile); err != nil {
        return nil, err
    }
    return &userProfile, nil
}

// GetUserConsistent 查询用户数据是否一致
func GetUserConsistent(uId uint, uAccount string) (uint, string, error) {
    var user struct {
        ID      uint   `gorm:"column:id"`
        Account string `gorm:"column:account"`
    }
    if err := mysqlDB.GetDB().Table("user").Select("id, account").Where("id = ? AND account = ?", uId, uAccount).Take(&user).Error; err != nil {
        return 0, "", err
    }
    return user.ID, user.Account, nil
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
