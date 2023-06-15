package main

import (
	"fmt"
	setting "gindome/config"
	"gindome/db/mysqlDB"
	log "gindome/logging"
	"gindome/router"
)


// @title helloworld
// @version 1.0
// @description chenwenyu test helloworld
// @termsOfService http://swagger.io/terms/
// @contact.name cwy
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host cwy.helloworld.com
// @BasePath /base/path
func main() {

	// 1.加载配置
	// setting.Init(os.Args[1])
	setting.Init("a")
	// 2.初始化日志
	log.Init()
	// 3.初始化mysql
	mysqlDB.Init(setting.Conf.MySQLConfig)
	// 4.初始化表结构
	mysqlDB.AutoMigrateDB()
	// 7.注册路由
	r := router.InitRouter()
	err := r.Run(fmt.Sprintf("127.0.0.1:%d", setting.Conf.Port))
	if err != nil {
		return
	}
}
