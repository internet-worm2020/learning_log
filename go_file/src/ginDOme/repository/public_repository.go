package repository

import (
	"fmt"

	"gorm.io/gorm"
)

func sqlError(err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return nil
	case nil:
		return nil
	default:
		return fmt.Errorf("查询失败: %v", err)
	}
}
