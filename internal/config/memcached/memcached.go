package config

import (
	"github.com/bradfitz/gomemcache/memcache"
	"os"
)

func ConnectMemcached() *memcache.Client {
	mc := memcache.New(os.Getenv("MEMCACHED_DSN"))
	return mc
}

