package models

import (
	"crypto/sha1"
	"encoding/hex"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// 登录账号
	Account string `json:"account,omitempty" gorm:"unique;not null" validate:"required"`
	// 登录密码
	Password string `json:"password,omitempty" gorm:"not null" validate:"required"`
	// 校验登录密码
	RePassword string `json:"re_password,omitempty" gorm:"-" validate:"eqfield=Password"`
	// 用户状态， -1 - 异常；0 - 锁定；1 - 正常；
	State       int16       `json:"state,omitempty" gorm:"default:1"`
	UserProfile UserProfile `json:"user_profile,omitempty"`
}

func (User) TableName() string {
	return "user"
}
func (u *User) HashPassword() {
	var hashPassword []byte = []byte(u.Password)
	var hash [20]byte = sha1.Sum(hashPassword)
	u.Password = hex.EncodeToString(hash[:])
}

type UserProfile struct {
	gorm.Model `json:"-"`
	// 账号名称
	Name string `json:"name,omitempty" gorm:"not null"`
	// 年龄
	Age uint `json:"age,omitempty"`
	// 性别
	Sex uint8 `json:"sex,omitempty" gorm:"default:0"`
	// 手机号
	Number string `json:"number,omitempty"`
	// 地址
	Address string `json:"address,omitempty"`
	// 身份证号
	IdCard string `json:"id_card,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// User关联外键
	UserID uint `json:"user_id,omitempty"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}
