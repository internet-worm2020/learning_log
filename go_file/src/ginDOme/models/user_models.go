package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	setting "gindome/config"
	"gindome/db/mysqlDB"
	"gorm.io/gorm"
)

// User 用户模型.
type User struct {
	gorm.Model `json:"-"`
	// Account 登录账号
	Account string `json:"account" gorm:"unique;not null" validate:"required"`
	// Password 登录密码
	Password string `json:"password,omitempty" gorm:"not null" validate:"required"`
	// RePassword 校验登录密码
	RePassword string `json:"re_password,omitempty" gorm:"-" validate:"eqfield=Password"`
	// State 用户状态， -1 - 异常；0 - 锁定；1 - 正常；
	State         int16 `json:"state" gorm:"default:1"`
	UserProfileID uint  `json:"-"`
	// UserProfile 用户信息
	UserProfile *UserProfile `json:"user_profile" gorm:"constraint:OnDelete:CASCADE;"`
}

func (*User) TableName() string {
	return "user"
}

func (u *User) HashPassword() {
	var salt []byte = []byte(setting.GetConf().Md5Key)
	var passwordWithSalt []byte
	passwordWithSalt = append(passwordWithSalt, salt[:]...)            // 添加盐值到密码数组
	passwordWithSalt = append(passwordWithSalt, []byte(u.Password)...) // 添加用户密码到密码数组
	var hash [32]byte = sha256.Sum256(passwordWithSalt)
	u.Password = hex.EncodeToString(hash[:]) // 将密码及盐值的哈希值转换为十六进制字符串，并保存在 User 结构体中作为新的密码
}

func (u *User) GetUser() (*User, error) {
	if err := mysqlDB.GetDB().Select("id", "account", "state", "user_profile_id").Where("id=?", u.ID).Preload("UserProfile").Find(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
func (u *User) GetAllUser(page, size int) ([]*User, error) {
	var userList []*User = make([]*User, 0, size)
	err := mysqlDB.GetDB().Select("id", "account", "state", "user_profile_id").Preload("UserProfile").Limit(size).Offset((page - 1) * size).Find(&userList).Error
	fmt.Println(userList, err)
	if err != nil {
		return nil, err
	}
	return userList, nil
}
func (u *User) Create() (*User, error) {
	// 将用户信息添加到数据库
	if err := mysqlDB.GetDB().Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
