package redis

import (
	"context"
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
	ttl         time.Duration
}

func New(addr, password string, ttl time.Duration) *AppRedis {
	client := Connect(addr, password)
	return &AppRedis{
		redisClient: client,
		ttl:         ttl,
	}
}

func (r *AppRedis) Get(ctx context.Context, key string) (string, error) {
	return r.redisClient.Get(ctx, key).Result()
}

func (r *AppRedis) Set(ctx context.Context, key string, value interface{}) error {
	return r.redisClient.Set(ctx, key, value, r.ttl).Err()
}

func (r *AppRedis) Del(ctx context.Context, keys ...string) error {
	return r.redisClient.Del(ctx, keys...).Err()
}

func (r *AppRedis) Close() error {
	return r.redisClient.Close()
}
