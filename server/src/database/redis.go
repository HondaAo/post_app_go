package database

import (
	"github.com/go-redis/redis/v8"
)

func SetupRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
}
