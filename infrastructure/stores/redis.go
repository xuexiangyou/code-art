package stores

import (
	"github.com/go-redis/redis/v8"
	"github.com/xuexiangyou/code-art/config"
)

func ConnectRedis(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Pass, // no password set
		DB:       0,  // use default DB
	})
	return rdb
}