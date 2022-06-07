package cache

import (
	"github.com/go-redis/redis/v8"
)

type Cache struct {
  Cache *redis.Client
}

func ProvideCache(cache *redis.Client) Cache {
	return Cache{Cache: cache}
}

