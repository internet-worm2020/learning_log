package repository

import (
	"gindome/db/mysqlDB"
	"gindome/models"
)

// RegisterUser 注册用户，添加数据到数据库
// @Summary 注册用户
// @Description 添加数据到数据库
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param u body models.User true "用户信息"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /user/register [post]
func RegisterUser(u *models.User) error {
	// 将用户信息添加到数据库
	if err := mysqlDB.GetDB().Create(u).Error; err != nil {
		return sqlError(err)
	}
	return nil
}

// GetAccount 根据账户名查找用户
// @Summary 根据账户名查找用户
// @Description 根据账户名查找用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param account query string true "账户名"
// @Success 200 {object} models.User "用户信息"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /user/account [get]
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
// @Summary 根据账户名获取用户ID
// @Description 根据账户名获取用户ID
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param account query string true "账户名"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /user/id [get]
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
// @Summary 根据用户ID获取用户信息
// @Description 根据用户ID获取用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param userId query uint64 true "用户ID"
// @Success 200 {object} models.UserProfile "用户信息"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /user/info [get]
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
// @Summary 查询用户数据是否一致
// @Description 查询用户数据是否一致
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param uId query uint true "用户ID"
// @Param uAccount query string true "账户名"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /user/consistent [get]
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
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param page query int true "页码"
// @Param size query int true "每页数量"
// @Success 200 {array} models.User "用户列表"
// @Failure 400 {string} string "bad request"
// @Failure 500 {string} string "internal server error"
// @Router /user/list [get]
func GetUserList(page, size int) ([]*models.User, error) {
	// 初始化一个空的用户列表
	users := make([]*models.User, 0, size)
	// 查询用户列表
	err := mysqlDB.GetDB().Preload("UserProfile").Limit(size).Offset((page - 1) * size).Find(&users).Error
	// 返回用户列表和错误信息
	return users, err
}
