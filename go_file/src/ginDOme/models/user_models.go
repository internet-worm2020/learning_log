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
	Account string `json:"account" gorm:"unique;not null" validate:"required"`
	// Password 登录密码
	Password string `json:"password" gorm:"not null" validate:"required"`
	// RePassword 校验登录密码
	RePassword string `json:"re_password" gorm:"-" validate:"eqfield=Password"`
	// State 用户状态， -1 - 异常；0 - 锁定；1 - 正常；
	State         int16 `json:"state" gorm:"default:1"`
	UserProfileID uint
	// UserProfile 用户信息
	UserProfile *UserProfile `json:"user_profile" gorm:"constraint:OnDelete:CASCADE;"`
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
	Name string `json:"name" gorm:"not null" validate:"required"`
	// 年龄
	Age uint `json:"age" gorm:"default:0" validate:"lte=150"`
	// 性别
	Sex uint8 `json:"sex" gorm:"default:0" validate:"lte=2"`
	// 手机号
	Number string `json:"number"`
	// 地址
	Address string `json:"address"`
	// 身份证号
	IdCard string `json:"id_card"`
	// 邮箱
	Email string `json:"email"`
	// User关联外键
	// UserID uint `json:"user_id,omitempty"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}

func (u *UserProfile)ToMap()map[string]interface{}{
	data := make(map[string]interface{})
	data["name"] = u.Name
	data["age"] = u.Age
	data["sex"]=u.Sex
	data["number"]=u.Number
	data["address"]=u.Address
	data["id_card"]=u.IdCard
	data["email"] = u.Email
	return data
}