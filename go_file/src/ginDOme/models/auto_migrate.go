package models

import "gindome/db/mysqlDB"

/*
AutoMigrateDB

@description: 自动迁移数据库.
*/
func AutoMigrateDB() {
	// 1. 获取数据库连接
	db := mysqlDB.GetDB()
	// 2. 自动迁移用户和用户资料表
	db.AutoMigrate(
		&User{},
		&UserProfile{},
	)
}
