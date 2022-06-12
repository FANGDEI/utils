package redis

import (
	"context"
	"utils/config"
	"utils/logs"

	"github.com/go-redis/redis/v8"
)

var M *Manager

func init() {
	var err error
	M, err = New()
	if err != nil {
		logs.M.FatallnError(err)
	}
}

type Manager struct {
	redisDb *redis.Client
}

func New() (*Manager, error) {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     config.M.Redis.Addr,     // redis地址
		Password: config.M.Redis.Password, // redis密码
		DB:       config.M.Redis.DBnum,    // redis连接数据库(0 ~ 15) 默认是0
	})

	// 通过 *redis.Client.Ping() 检查是否成功连接到redis服务器
	_, err := redisDb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	m := &Manager{
		redisDb: redisDb,
	}
	return m, nil
}
