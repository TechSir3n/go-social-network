package generator

import (
	"context"
	"log"

 "social_network/internal/repository/database/redis"

	"github.com/sethvargo/go-password/password"
	"gopkg.in/robfig/cron.v2"
)


// generate a random special key of 5 characters and 5 numbers, 
// which will be known only to the administrator
func GeneratePassword() (string, error) {
	var redis database.Redis
	res, err := password.Generate(10, 5, 5, false, false)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	ctx := context.Background()
	err = redis.GreateAdminPassword(ctx, res)
	if err != nil {
		log.Println(err, ":Failed to create admin password")
		return "", err
	}

	return res, nil
}


// will call Generate password every 10 minutes to update admins's special key
func init() {
	s := cron.New()
	s.AddFunc("@every 10m", func() {
		GeneratePassword()
	})

	s.Start()
}
