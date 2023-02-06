package database

import (
	"context"
	"log"
	"social_network/internal/api/v1/models"
	config "social_network/internal/config/redis"
	"time"
)

// NOSql redis

var CLRedis = config.InitRedis()

type AdminRepository interface {
	GreateAdminPassword(ctx context.Context) error
	GetAdminPassword(ctx context.Context) (models.Admin, error)
}

type AdminService struct {
	AdminRepository AdminRepository
	Admin           models.Admin
}

func NewAdminService(admin AdminRepository, model models.Admin) *AdminService {
	return &AdminService{
		AdminRepository: admin,
		Admin:           model,
	}
}

// add special_key for admin into redis 
func GreateAdminPassword(ctx context.Context, pass string) error {
	set, err := CLRedis.SetNX(ctx,"SPECIAL_KEY", pass, 10*time.Minute).Result()
	if err != nil {
		log.Println(err, " :Failed to set special_key into redis, [:ADMIN]")
		return err
	}

	if set {
		log.Println("SUCESS INSERTED")
	}

	return nil
}

// get special_key of redis 
func GetAdminPassword(ctx context.Context) (models.Admin, error) {
	get, err := CLRedis.Get(ctx, "SPECIAL_KEY").Result()
	if err != nil {
		log.Println(err, " Failed to get special_key,[:ADMIN]")
		return models.Admin{}, err
	}
	return models.Admin{
		Special_Key: string(get),
	}, nil
}
