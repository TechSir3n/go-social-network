package config

import (
	"github.com/redis/go-redis/v9"
	"os"
)


func InitRedis() *redis.Client {
	rd := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DSN"),
		Password: "",
		DB:       0,
	})

	return rd
}
