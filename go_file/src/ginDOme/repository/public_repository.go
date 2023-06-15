package repository

import (
	"fmt"
	"gorm.io/gorm"
)

/*
@description: 处理 SQL 查询错误

@param: err error SQL 查询错误

@return: error 错误信息
*/
func sqlError(err error) error {
	// 1. 判断错误类型
	switch err {
	// 2. 如果是记录未找到错误，则返回 nil
	case gorm.ErrRecordNotFound:
		return nil
	// 3. 如果是没有错误，则返回 nil
	case nil:
		return nil
	// 4. 如果是其他错误，则返回详细错误信息
	default:
		return fmt.Errorf("查询失败: %v", err)
	}
}
