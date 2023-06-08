package mysqlDB
import "gindome/models"
func AutoMigrateDB() {
	GetDB().AutoMigrate(
		&models.User{},
		&models.UserProfile{},
	)
}