package logging

import (
	"fmt"
	setting "gindome/config"
	"github.com/sirupsen/logrus"
	"os"
)

var WebLog *logrus.Logger

/*
Init

@description: 初始化函数
*/
func Init() {
	// 1. 初始化Web日志
	initWebLog()
}

/*
@description: 初始化Web日志
*/
func initWebLog() {
	// 1. 初始化Web日志名称
	webLogName := setting.Conf.LogConfig.WebLogName
	// 2. 初始化Web日志
	WebLog = initLog(webLogName)
}

/*
@description: 初始化日志句柄

@param: logFileName string 日志文件名

@return: *logrus.Logger 日志句柄
*/
func initLog(logFileName string) *logrus.Logger {
	// 1. 创建一个新的日志句柄
	log := logrus.New()
	// 2. 设置日志格式
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	// 3. 获取日志文件路径
	logFilePath := setting.Conf.LogFilePath
	// 4. 拼接日志文件名
	logName := logFilePath + logFileName
	var f *os.File
	var err error
	// 5. 判断日志文件夹是否存在，不存在则创建
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		os.MkdirAll(logFilePath, os.ModePerm)
	}
	// 6. 判断日志文件是否存在，不存在则创建，否则就直接打开
	if _, err := os.Stat(logName); os.IsNotExist(err) {
		f, err = os.Create(logName)
	} else {
		f, err = os.OpenFile(logName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}
	if err != nil {
		fmt.Println("open log file failed")
	}
	// 7. 设置日志输出文件
	log.Out = f
	// 8. 设置日志级别
	log.Level = logrus.DebugLevel
	// 9. 返回日志句柄
	return log
}
