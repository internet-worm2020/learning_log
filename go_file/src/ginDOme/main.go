package main

import (
	"context"
	"fmt"
	setting "gindome/config"
	"gindome/db/mysqlDB"
	"gindome/db/redis"
	log "gindome/logging"
	"gindome/pkg/jobs"
	"gindome/pkg/snowflake"
	"gindome/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	logs "github.com/internet-worm2020/go-pkg/log"
)

func main() {
	defer mysqlDB.CloseDB()
	defer redis.CloseRedis()
	// 1.加载配置
	// setting.Init(os.Args[1])
	setting.Init("a")
	// 2.初始化日志
	// log.Init()
	logs.Init(log.Adcc())
	// 3.初始化mysql
	mysqlDB.Init(setting.Conf.MySQLConfig)
	// 4.初始化表结构
	mysqlDB.AutoMigrateDB()
	// 5.初始化redis
	redis.Init(setting.Conf.RedisConfig)
	// 6.初始化定时任务
	jobs.InitJobs()
	// 7. 初始化雪花算法
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// 7.注册路由
	r := router.InitRouter()
	srv:=&http.Server{
		Addr:fmt.Sprintf("127.0.0.1:%d", setting.Conf.Port),
		Handler: r,
	}
	go func(){
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown:", err)
	}
	fmt.Println("Server exiting")
	// err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", setting.Conf.Port), r)
	// if err != nil {
	// 	return
	// }
}
