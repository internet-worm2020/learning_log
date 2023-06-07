package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level       string `mapstructure:"level"`
	WebLogName  string `mapstructure:"web_log_name"`
	LogFilePath string `mapstructure:"log_file_path"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	Timeout      string `mapstructure:"timeout"`
	ReadTimeout  string `mapstructure:"readTimeout"`
	WriteTimeout string `mapstructure:"writeTimeout"`
	Loc          string `mapstructure:"loc"`
	Charset      string `mapstructure:"charset"`
	ParseTime    bool   `mapstructure:"parseTime"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

var (
	devFilePath     string = "./config/config.dev.yaml"
	releaseFilePath string = "./config/config.dev.yaml"
	localFilePath   string = "./config/config.local.yaml"
)

func Init(mode string) {
	var filePath string
	if mode == "dev" {
		filePath = devFilePath
	} else if mode == "release" {
		filePath = releaseFilePath
	} else { // local
		filePath = localFilePath
	}
	fmt.Println()
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		panic(fmt.Sprintf("viper.ReadInConfig failed, err:%v\n", err))
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Sprintf("viper.Unmarshal failed, err:%v\n", err))
		}
	})
}
