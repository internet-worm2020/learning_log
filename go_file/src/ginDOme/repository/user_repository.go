package repository

import (
	"fmt"
	"gindome/db/mysqlDB"
	"gindome/models"
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
	db = db.Where("account=?", account).Or("password=?", password).Find(&user)
	// 如果查找出错，则返回错误信息
	if db.Error != nil {
		return nil, sqlError(db.Error)
	}
	return &user, nil
}

/*
GetIDByAccount

@description: 根据账户名获取用户ID

@param: account string 账户名

@return: uint 用户ID

@return: error 错误信息.
*/
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
	mysqlDB.GetDB().Where("id=?",userId).Preload("UserProfile").Find(&user)
	// 返回用户信息变量
	return &user.UserProfile, nil
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
	err := mysqlDB.GetDB().Limit(size).Offset((page - 1) * size).Find(&userList).Error
	// 3. 返回用户列表和错误信息
	return userList, err
}
/*
DeleteUser

@description: 删除用户信息

@param: uId uint 用户ID

@retunr: error 错误信息.
*/
func DeleteUser(uId uint)(error) {
	// 获取数据库连接
	db := mysqlDB.GetDB()
	// 定义一个用户变量
	var user models.User
	// 根据id查找用户
	db.Where("id=?", uId)
	// 定义一个数量
	var count int64
	// 查询到多少数据
	db.Model(&user).Count(&count)
	// 要删除的数据是否存在
	if count==0{
		return fmt.Errorf("要删除数据不存在")
	}
	// 要删除的数据是否唯一
	if count >1 {
		return fmt.Errorf("删除数据不唯一")
	}
	// 删除数据
	db.Delete(&user)
	// 如果查找出错，则返回错误信息
	if db.Error != nil {
		return sqlError(db.Error)
	}
	// 返回错误信息
	return nil
}

func UpdateUserProfile(uId uint)(int64,error){
	db := mysqlDB.GetDB()
	var user models.User
	// db.Where("id=?",uId).Preload("UserProfile").Find(&user)
	db.Model(&user).Where("id=?",uId)
	fmt.Println(user)
	return 1,nil
}