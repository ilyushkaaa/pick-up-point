package cache

import (
	"context"
)

type Cache interface {
	GoDeleteFromCache(ctx context.Context, keys ...string)
	GoAddToCache(ctx context.Context, key string, value interface{})
	GetFromCache(ctx context.Context, key string) (interface{}, error)
	Close() error
}
