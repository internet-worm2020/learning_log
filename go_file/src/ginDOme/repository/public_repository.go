package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func sqlError(err error)(error){
	switch err {
    case gorm.ErrRecordNotFound:
        return errors.New("查询不到数据")
    case nil:
        return nil
    default:
        return fmt.Errorf("查询失败: %w", err)
    }
}