package db

import (
	"context"

	"github.com/ChaotenHG/auth-server/config"
	"github.com/redis/go-redis/v9"
)

var RedisContext context.Context = context.Background()
var Rdb *redis.Client

func LoadRedisClient(cfg *config.Config) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Database,
	})
}
