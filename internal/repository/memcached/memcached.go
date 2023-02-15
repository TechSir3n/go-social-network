package memcached

import (
	"encoding/json"
	"social_network/internal/api/v1/models"
	"social_network/internal/config/memcached"
	"social_network/utils/logger"
	"github.com/bradfitz/gomemcache/memcache"
)

var mc = config.ConnectMemcached()

func SetMemcached(user models.User) error {
	val, _ := json.Marshal(user)
	err := mc.Set(&memcache.Item{Key: "credentials", Value: val, Expiration:int32(600)}) // 30 second time to live
	if err != nil {
		logger.Error(err.Error(), "Failed to set user's data into memcached")
		return err
	}

	return nil
}

func GetMemcached(key string) (models.User, error) {
	data, err := mc.Get(key)
	if err != nil {
		logger.Error(err.Error(), "Failed to get value of the memcached")
		return models.User{}, err
	}

	var user models.User
	err = json.Unmarshal(data.Value, &user)
	if err != nil {
		logger.Error(err.Error(), ": Unable unmarshal")
		return models.User{}, err
	}

	return user, nil
}
