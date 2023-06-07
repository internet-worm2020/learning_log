package app

import (
	"gindome/middleware"
	"gindome/app/shop"
	"gindome/app/blog"
	"gindome/app/users"
	"gindome/db/mysqlDB"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	shop.Routers(r)
	blog.Routers(r)
	users.Routers(r)
	return r
}
func AutoMigrateDB() {
	mysqlDB.GetDB().AutoMigrate(
		&users.User{},
		&users.UserProfile{},
	)
}