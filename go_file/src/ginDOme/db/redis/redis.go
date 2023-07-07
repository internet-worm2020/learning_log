package redis

import (
	"errors"
	"fmt"
	setting "gindome/config"

	"github.com/go-redis/redis"
)

var redisClients []*redis.Client

/*
Init

@description: 创建redis客户端并测试连接是否成功

@param: cfgs []*setting.RedisConfig Redis配置信息列表.
*/
func Init(cfg *setting.RedisConfig) {
	for _, s := range cfg.DB {
		// 1. 创建redis客户端
		client := redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Password:     cfg.Password, // no password set
			DB:           s,            // use default DB
			PoolSize:     cfg.PoolSize,
			MinIdleConns: cfg.MinIdleConns,
		})

		// 2. 测试连接是否成功
		if _, err := client.Ping().Result(); err != nil {
			// 3. 连接失败，抛出panic
			panic(fmt.Sprintf("连接redis失败, %s", err.Error()))
		}
		// 3. 将客户端添加到切片中
		redisClients = append(redisClients, client)
	}
}

/*
GetRedis

@description: 获取指定索引位置的Redis客户端实例

@param: index int 客户端索引位置

@return: *redis.Client Redis客户端实例

@return: error 错误信息.
*/
func GetRedis(index int) (*redis.Client, error) {
	// 1. 根据实际情况检查index参数的合法性
	if index < 0 || index > 15 {
		return nil, errors.New("未知数据库")
	}
	// 2. 返回对应索引的客户端实例
	return redisClients[index], nil
}

/*
CloseRedis

@description: 关闭所有redis客户端

@return: error 错误信息.
*/
func CloseRedis() error {
	// 遍历所有 Redis 客户端，逐一关闭
	for _, client := range redisClients {
		err := client.Close()
		if err != nil {
			return err
		}
	}

	// 返回空的错误信息（表示成功）
	return nil
}
