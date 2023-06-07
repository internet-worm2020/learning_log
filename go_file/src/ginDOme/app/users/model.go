package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	// 账号名称
	Username string `json:"username" gorm:"not null" binding:"required"`
	// 登录账号
	Account string `json:"account" gorm:"not null" binding:"required"`
	// 登录密码
	Password string `json:"password" gorm:"not null" binding:"required"`
	// 校验登录密码
	RePassword string `json:"re_password,omitempty" gorm:"-" binding:"required,eqfield=Password"`
	// 用户状态， -1 - 异常；0 - 锁定；1 - 正常；
	State       int16        `json:"state" gorm:"default:1"`
	UserProfile *UserProfile `json:"user_profile,omitempty"`
}

func (User) TableName() string {
	return "user"
}



type UserProfile struct {
	gorm.Model `json:"-"`
	Age        uint `json:"age"`
	Sex        uint `json:"sex" gorm:"default:0"`
	UserID     uint `json:"-"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}
