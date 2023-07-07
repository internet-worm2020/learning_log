package mysqlDB

import "gindome/models"

/*
AutoMigrateDB

@description: 自动迁移数据库.
*/
func AutoMigrateDB() {
	// 1. 获取数据库连接
	db := GetDB()
	// 2. 自动迁移用户和用户资料表
	db.AutoMigrate(
		&models.User{},
		&models.UserProfile{},
	)
}
