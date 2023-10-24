package service

import (
	"example/testmongdb/component/provider"

	"github.com/go-redis/redis"
)

type RedisService struct {
	Redis *redis.Client
}

// RedisService is a call provider and return string success or error
func (r *RedisService) GetRedis() (string, error) {
	redisProvider := provider.NewrRedisProvider(r.Redis)
	return redisProvider.Hget()
}
