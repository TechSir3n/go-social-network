package database

import (
	"context"
	"log"
	"social_network/internal/api/v1/models"
	config "social_network/internal/config/redis"
	"time"
)

// NOSql redis

var CLRedis = config.ConnectRedis()

type Redis struct {
	Admin interface {
		GreateAdminPassword(ctx context.Context, pass string) error
		GetAdminPassword(ctx context.Context) (models.Admin, error)
	}
}

// add special_key for admin into redis
func (s Redis) GreateAdminPassword(ctx context.Context, pass string) error {
	_, err := CLRedis.SetNX(ctx, "SPECIAL_KEY", pass, 10*time.Minute).Result()
	if err != nil {
		log.Println(err, " :Failed to set special_key into redis, [:ADMIN]")
		return err
	}

	return nil
}

// get special_key of redis
func (s Redis) GetAdminPassword(ctx context.Context) (models.Admin, error) {
	get, err := CLRedis.Get(ctx, "SPECIAL_KEY").Result()
	if err != nil {
		log.Println(err, " Failed to get special_key,[:ADMIN]")
		return models.Admin{}, err
	}
	return models.Admin{
		Special_Key: string(get),
	}, nil
}
