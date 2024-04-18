package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Connect(host, port string) *redis.Client {
	redisAddr := fmt.Sprintf("%s:%s", host, port)
	fmt.Println(redisAddr)
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
	return client
}
