package redis

import (
	"fmt"
	setting "gindome/config"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

/*
Init

@description: 创建redis客户端并测试连接是否成功

@param: cfg *setting.RedisConfig Redis配置信息
*/
func Init(cfg *setting.RedisConfig) {
	// 1. 创建redis客户端
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	// 2. 测试连接是否成功
	if _, err := redisClient.Ping().Result(); err != nil {
		// 3. 连接失败，抛出panic
		panic(fmt.Sprintf("连接redis失败, %s", err.Error()))
	}
}

/*
GetRedis

@description: 返回redis客户端

@return: *redis.Client Redis客户端实例
*/
func GetRedis() *redis.Client {
	// 1. 返回redis客户端
	return redisClient
}

/*
CloseRedis

@description: 关闭redis客户端

@return: error 错误信息
*/
func CloseRedis() error {
	// 1. 关闭redis客户端
	return redisClient.Close()
}

/*
NilRedis

@description: 返回redis.Nil

@return: error 错误信息
*/
func NilRedis() error {
	// 1. 返回redis.Nil
	return redis.Nil
}
