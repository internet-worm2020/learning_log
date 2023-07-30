package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

// AppConfig
// 定义 AppConfig 结构体，用于存储应用程序的配置信息.
type AppConfig struct {
	Mode           string                         `mapstructure:"mode"` // 应用程序运行模式
	Port           int                            `mapstructure:"port"` // 应用程序监听端口
	*LogConfig     `mapstructure:"log"`           // 日志配置信息
	*MySQLConfig   `mapstructure:"mysql"`         // MySQL 数据库配置信息
	*RedisConfig   `mapstructure:"redis"`         // Redis 数据库配置信息
	*KeyCollection `mapstructure:"keyCollection"` // 密钥集合
}

// LogConfig
// 定义 LogConfig 结构体，用于存储日志配置信息.
type LogConfig struct {
	Level           string `mapstructure:"level"`              // 日志级别
	WebLogName      string `mapstructure:"web_log_name"`       // Web 日志名称
	WebLogErrorName string `mapstructure:"web_log_error_name"` // Web 错误日志名称
	LogFilePath     string `mapstructure:"log_file_path"`      // 日志文件路径
}

// MySQLConfig
// 定义 MySQLConfig 结构体，用于存储 MySQL 数据库配置信息.
type MySQLConfig struct {
	Host         string `mapstructure:"host"`           // MySQL 数据库主机地址
	User         string `mapstructure:"user"`           // MySQL 数据库用户名
	Password     string `mapstructure:"password"`       // MySQL 数据库密码
	DB           string `mapstructure:"dbname"`         // MySQL 数据库名称
	Port         int    `mapstructure:"port"`           // MySQL 数据库端口号
	Timeout      string `mapstructure:"timeout"`        // 连接超时时间
	ReadTimeout  string `mapstructure:"readTimeout"`    // 读取超时时间
	WriteTimeout string `mapstructure:"writeTimeout"`   // 写入超时时间
	Loc          string `mapstructure:"loc"`            // 时区
	Charset      string `mapstructure:"charset"`        // 字符集
	ParseTime    bool   `mapstructure:"parseTime"`      // 是否解析时间
	MaxOpenConns int    `mapstructure:"max_open_conns"` // 最大连接数
	MaxIdleConns int    `mapstructure:"max_idle_conns"` // 最大空闲连接数
}

// RedisConfig
// 定义 RedisConfig 结构体，用于存储 Redis 数据库配置信息.
type RedisConfig struct {
	Host         string  `mapstructure:"host"`           // Redis 数据库主机地址
	Password     string  `mapstructure:"password"`       // Redis 数据库密码
	Port         int     `mapstructure:"port"`           // Redis 数据库端口号
	DB           [16]int `mapstructure:"db"`             // Redis 数据库编号
	PoolSize     int     `mapstructure:"pool_size"`      // 连接池大小
	MinIdleConns int     `mapstructure:"min_idle_conns"` // 最小空闲连接数
}

// KeyCollection
// 定义 KeyCollection 结构体，用于存储 各式密钥 数据库配置信息.
type KeyCollection struct {
	JwtKey string `mapstructure:"jwtKey"` // jwt密钥
	Md5Key string `mapstructure:"md5Key"` // md5密钥
}

// 定义配置文件路径常量.
const (
	devFilePath     string = "./config/config.dev.yaml"     // 开发环境配置文件路径
	releaseFilePath string = "./config/config.release.yaml" // 生产环境配置文件路径
	localFilePath   string = "./config/config.local.yaml"   // 本地环境配置文件路径
)

/*
Init

@description: 初始化配置文件

@param: mode string 运行模式.
*/
func Init(mode string) {
	// 1. 根据运行模式选择配置文件路径
	var filePath string
	if mode == "dev" {
		filePath = devFilePath
	} else if mode == "release" {
		filePath = releaseFilePath
	} else { // local
		filePath = localFilePath
	}

	// 2. 设置配置文件路径
	viper.SetConfigFile(filePath)

	// 3. 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("viper.ReadInConfig failed, err:%v\n", err))
	}

	// 4. 解析配置文件
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	// 5. 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Sprintf("viper.Unmarshal failed, err:%v\n", err))
		}
	})
}

func GetConf() *AppConfig {
	return Conf
}
