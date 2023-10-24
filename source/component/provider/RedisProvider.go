package provider

import (
	"time"

	"github.com/go-redis/redis"
)

type redisProvider struct {
	redis *redis.Client
}

func NewrRedisProvider(redis *redis.Client) *redisProvider {
	//return repository
	return &redisProvider{
		redis: redis,
	}
}

// GetRedis is a function
func (r *redisProvider) GetRedis(name string) string {
	return "Success"
}

func (r *redisProvider) Hget() (string, error) {
	key := "LOADTEST.GOLANG.HASH"

	//get redis
	data, _ := r.redis.HGet(key, "1234").Result()
	//set expire 15s
	r.redis.Expire(key, time.Duration(15)*time.Second)
	return data, nil

}
