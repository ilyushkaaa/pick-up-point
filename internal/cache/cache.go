package cache

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Cache interface {
	GoDeleteFromCache(ctx context.Context, keys ...string)
	GoAddToCache(ctx context.Context, key string, value interface{})
	GetFromCache(ctx context.Context, key string) (interface{}, error)
	Close() error
}

type Config struct {
	OrderByIDTTl      time.Duration
	PPByIDTTl         time.Duration
	OrdersByClientTTl time.Duration
	RedisTTl          time.Duration
	RedisAddr         string
	RedisPassword     string
}

func GetConfig() (*Config, error) {
	OrderByIDTTl, err := strconv.ParseUint(os.Getenv("CACHE_ORDER_BY_ID_TTL"), 10, 64)
	if err != nil {
		return nil, err
	}
	PPByIDTTl, err := strconv.ParseUint(os.Getenv("CACHE_PP_BY_ID_TTL"), 10, 64)
	if err != nil {
		return nil, err
	}
	OrdersByClientTTl, err := strconv.ParseUint(os.Getenv("CACHE_ORDERS_BY_CLIENT_TTL"), 10, 64)
	if err != nil {
		return nil, err
	}
	RedisTTl, err := strconv.ParseUint(os.Getenv("CACHE_REDIS_TTL"), 10, 64)
	if err != nil {
		return nil, err
	}
	return &Config{
		OrderByIDTTl:      time.Duration(OrderByIDTTl),
		PPByIDTTl:         time.Duration(PPByIDTTl),
		OrdersByClientTTl: time.Duration(OrdersByClientTTl),
		RedisTTl:          time.Duration(RedisTTl),
		RedisAddr:         fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		RedisPassword:     os.Getenv("REDIS_PASSWORD"),
	}, nil
}
