package cache

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"
	"homework/pkg/redis"
)

type RedisCache struct {
	redis  redis.Redis
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger) *RedisCache {
	rd := redis.New()
	return &RedisCache{
		redis:  rd,
		logger: logger,
	}
}

func (r *RedisCache) GoDeleteFromCache(ctx context.Context, keys ...string) {
	go func() {
		err := r.redis.Del(ctx, keys...)
		if err != nil {
			r.logger.Errorf("error in deleting from cache: %v", err)
		}
	}()
}

func (r *RedisCache) GoAddToCache(ctx context.Context, key string, value interface{}) {
	go func() {
		dataJSON, err := json.Marshal(value)
		if err != nil {
			r.logger.Errorf("key: %s, value: %v: error in adding into redis: %v", key, value, err)
			return
		}
		err = r.redis.Set(ctx, key, dataJSON)
		if err != nil {
			r.logger.Errorf("key: %s, value: %v: error in adding into redis: %v", key, value, err)
		}
	}()
}

func (r *RedisCache) GetFromCache(ctx context.Context, key string) (interface{}, error) {
	return r.redis.Get(ctx, key)
}

func (r *RedisCache) Close() error {
	return r.redis.Close()
}
