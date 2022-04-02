package stores

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/xuexiangyou/code-art/config"
	"go.uber.org/fx"
)

func ConnectRedis(lc fx.Lifecycle, config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Pass, // no password set
		DB:       0,  // use default DB
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			rdb.Close()
			return nil
		},
	})

	return rdb
}