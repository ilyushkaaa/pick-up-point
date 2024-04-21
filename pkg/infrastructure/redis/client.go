package redis

import (
	"github.com/redis/go-redis/v9"
)

func Connect(addr, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	return client
}
