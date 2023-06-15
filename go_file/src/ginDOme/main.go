package main

import (
	"fmt"
	setting "gindome/config"
	"gindome/db/mysqlDB"
	"gindome/db/redis"
	"gindome/pkg/jobs"
	log "gindome/logging"
	"gindome/router"
)


func main() {
	defer mysqlDB.CloseDB()
	defer redis.CloseRedis()
	// 1.加载配置
	// setting.Init(os.Args[1])
	setting.Init("a")
	// 2.初始化日志
	log.Init()
	// 3.初始化mysql
	mysqlDB.Init(setting.Conf.MySQLConfig)
	// 4.初始化表结构
	mysqlDB.AutoMigrateDB()
	// 5.初始化redis
	redis.Init(setting.Conf.RedisConfig);
	// 6.初始化定时任务
	jobs.InitJobs()
	// 7.注册路由
	r := router.InitRouter()
	err := r.Run(fmt.Sprintf("127.0.0.1:%d", setting.Conf.Port))
	if err != nil {
		return
	}
}
