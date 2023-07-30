package models

import "gorm.io/gorm"

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

func (*UserProfile) TableName() string {
	return "user_profile"
}

func (u *UserProfile) ToMap() map[string]interface{} {
	data := make(map[string]interface{})
	data["name"] = u.Name
	data["age"] = u.Age
	data["sex"] = u.Sex
	data["number"] = u.Number
	data["address"] = u.Address
	data["id_card"] = u.IdCard
	data["email"] = u.Email
	return data
}
