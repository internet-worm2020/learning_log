package repository

import (
	"fmt"
	"gindome/db/mysqlDB"
	"gindome/models"
	"github.com/internet-worm2020/go-pkg/log"
)

/*
RegisterUser

@description: RegisterUser 注册用户，添加数据到数据库

@param: u models.User 用户信息

@return: error 错误信息.
*/
func RegisterUser(u *models.User) (*models.User, error) {
	// 获取数据库连接
	db := mysqlDB.GetDB()
	// 将用户信息添加到数据库
	if err := db.Create(u).Error; err != nil {
		log.Error(sqlError(err).Error())
		return nil, sqlError(err)
	}
	return u, nil
}

/*
GetAccount

@description: 根据账户名查找用户

@param: account string 账户名

@return: int64 影响的行数

@return: error 错误信息.
*/
func GetAccount(account string) (int64, error) {
	// 获取数据库连接
	db := mysqlDB.GetDB()
	// 定义一个用户变量
	var user models.User
	// 根据账户名查找用户
	db = db.Where("account=?", account).Find(&user)
	// 如果查找出错，则返回错误信息
	if db.Error != nil {
		log.Error(sqlError(db.Error).Error())
		return 0, sqlError(db.Error)
	}
	// 返回用户信息和影响的行数
	return db.RowsAffected, nil
}

/*
GetAccountPassword

@description: 根据账号和密码查询用户

@param: account string 账户名

@param: password string 密码

@return: *models.User 用户信息

@return: error 错误信息.
*/
func GetAccountPassword(account string, password string) (*models.User, error) {
	// 获取数据库连接
	db := mysqlDB.GetDB()
	// 定义一个用户变量
	var user models.User
	// 根据账户名查找用户
	db = db.Where("account=? AND password=?", account, password).Find(&user)
	// 如果查找出错，则返回错误信息
	if db.Error != nil {
		log.Error(sqlError(db.Error).Error())
		return nil, sqlError(db.Error)
	}
	return &user, nil
}

/*
GetUserById

@description: 根据用户ID获取用户信息

@param: userId uint64 用户ID

@return: *models.UserProfile 用户信息变量

@return: error 错误信息.
*/
func GetUserById(userId uint64) (*models.UserProfile, error) {
	// 定义一个用户变量和用户信息变量
	var user models.User
	// 根据用户信息获取用户信息变量

	if err := mysqlDB.GetDB().Where("id=?", userId).Preload("UserProfile").Find(&user).Error; err != nil {
		log.Error(err.Error())
		return nil, err
	}
	// 返回用户信息变量
	return user.UserProfile, nil
}

/*
GetUserConsistent

@description: 查询用户数据是否一致

@param: uId uint 用户ID

@param: uAccount string 用户账户名

@return: uint 用户ID

@return: string 用户账户名

@return: error 错误信息.
*/
func GetUserConsistent(uId uint, uAccount string) (uint, string, error) {
	// 定义一个用户变量
	var user struct {
		ID      uint   `gorm:"column:id"`      // 用户ID
		Account string `gorm:"column:account"` // 用户账户名
	}
	// 查询用户数据是否一致
	if err := mysqlDB.GetDB().Table("user").Select("id, account").Where("id = ? AND account = ?", uId, uAccount).Take(&user).Error; err != nil {
		log.Error(err.Error())
		return 0, "", err
	}
	// 返回用户ID和账户名
	return user.ID, user.Account, nil
}

/*
GetUserList

@description: 获取用户列表

@param: page int 分页页码

@param: size int 分页大小

@return: []*models.UserProfile 用户列表

@return: error 错误信息.
*/
func GetUserList(page, size int) ([]*models.UserProfile, error) {
	// 1. 创建一个空的用户列表
	userList := make([]*models.UserProfile, 0, size)
	// 2. 查询用户列表
	if err := mysqlDB.GetDB().Limit(size).Offset((page - 1) * size).Find(&userList).Error; err != nil {
		log.Error(err.Error())
		return nil, err
	}
	// 3. 返回用户列表和错误信息
	return userList, nil
}

/*
DeleteUser

@description: 删除用户信息

@param: uId uint 用户ID

@return: error 错误信息.
*/
func DeleteUser(uId uint) error {
	// 获取数据库连接
	db := mysqlDB.GetDB()
	// 定义一个用户变量
	var user models.User
	// 定义一个用户详情变量
	var userProfile models.UserProfile
	// 根据id查找用户
	userDB := db.Where("id=?", uId)
	userDB.Find(&user)
	// 根据关联id查找详情
	userProfilDB := db.Where("id=?", user.UserProfileID)
	// 定义一个数量
	var count int64
	// 查询到多少数据
	userDB.Count(&count)
	// 要删除的数据是否存在
	if count == 0 {
		log.Error(fmt.Errorf("要删除数据不存在").Error())
		return fmt.Errorf("要删除数据不存在")
	}
	// 删除数据
	userDB.Delete(&user)
	userProfilDB.Delete(&userProfile)
	// 如果查找出错，则返回错误信息
	if db.Error != nil {
		log.Error(sqlError(db.Error).Error())
		return sqlError(db.Error)
	}
	// 返回错误信息
	return nil
}

/*
UpdateUserProfile

@description: 修改用户详情

@param: uId uint 用户ID

@param: userProfile *models.UserProfile 提交的用户详情信息

@return: error 错误信息.
*/
func UpdateUserProfile(uId uint, userProfile *models.UserProfile) error {
	// 获取数据库链接
	db := mysqlDB.GetDB()
	// 定义user变量
	var user models.User
	// 定义错误
	var err error
	// 根据id查询用户信息
	err = db.Where("id=?", uId).Find(&user).Error
	if err != nil {
		log.Error(sqlError(db.Error).Error())
		return sqlError(err)
	}
	// 根据关联主键修改用户详情信息
	err = db.Model(&user.UserProfile).Where("id=?", user.UserProfileID).Updates(userProfile.ToMap()).Error
	if err != nil {
		log.Error(sqlError(db.Error).Error())
		return sqlError(err)
	}
	return nil
}
