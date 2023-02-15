package config

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
	"social_network/utils/logger"
)

func InitRedis() *redis.Client {
	rd := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DSN"),
		Password: "",
		DB:       0,
	})

	_, err := rd.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatal(err.Error(), "Unable connect to redis")
	}

	return rd
}
