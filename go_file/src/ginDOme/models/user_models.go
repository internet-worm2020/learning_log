package models

import (
	"crypto/sha256"
	"encoding/hex"
	setting "gindome/config"

	"gorm.io/gorm"
)

// User 用户模型.
type User struct {
	gorm.Model `json:"-"`
	// Account 登录账号
	Account string `json:"account,omitempty" gorm:"unique;not null" validate:"required"`
	// Password 登录密码
	Password string `json:"password,omitempty" gorm:"not null" validate:"required"`
	// RePassword 校验登录密码
	RePassword string `json:"re_password,omitempty" gorm:"-" validate:"eqfield=Password"`
	// State 用户状态， -1 - 异常；0 - 锁定；1 - 正常；
	State         int16 `json:"state,omitempty" gorm:"default:1"`
	UserProfileID uint
	// UserProfile 用户信息
	UserProfile UserProfile `json:"user_profile,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
}

func (User) TableName() string {
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
	// UserID uint `json:"user_id,omitempty"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}
