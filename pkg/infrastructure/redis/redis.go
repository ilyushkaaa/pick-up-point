package redis

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}) error
	Del(ctx context.Context, keys ...string) error
	Close() error
}

type AppRedis struct {
	redisClient *redis.Client
	expireTime  time.Duration
}

func New() *AppRedis {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	client := Connect(host, port)
	return &AppRedis{
		redisClient: client,
		expireTime:  time.Minute * 10,
	}
}

func (r *AppRedis) Get(ctx context.Context, key string) (string, error) {
	return r.redisClient.Get(ctx, key).Result()
}

func (r *AppRedis) Set(ctx context.Context, key string, value interface{}) error {
	return r.redisClient.Set(ctx, key, value, r.expireTime).Err()
}

func (r *AppRedis) Del(ctx context.Context, keys ...string) error {
	return r.redisClient.Del(ctx, keys...).Err()
}

func (r *AppRedis) Close() error {
	return r.redisClient.Close()
}
