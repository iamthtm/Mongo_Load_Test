package connect

import (
	//redis
	"github.com/go-redis/redis"
)

// function connect redis and return client
func ConnectRedis(host string, password string) *redis.Client {
	//connect redis
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0, // use default DB
	})

	return client
}
