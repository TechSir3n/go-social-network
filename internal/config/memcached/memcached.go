package config

import (
	"github.com/bradfitz/gomemcache/memcache"
	"os"
	"social_network/utils/logger"
)

func ConnectMemcached() *memcache.Client {
	mc := memcache.New(os.Getenv("MEMCACHED_DSN"))

	err := mc.Ping()
	if err != nil {
		logger.Fatal(err.Error(), "Unable connect to memcaced")
	}

	return mc
}
