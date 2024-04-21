package cache

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.uber.org/zap"
)

type data struct {
	expirationTime time.Time
	value          interface{}
}

type InMemoryCache struct {
	cache  map[string]data
	mu     *sync.RWMutex
	ttl    time.Duration
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger, ttl time.Duration) *InMemoryCache {
	return &InMemoryCache{
		cache:  make(map[string]data),
		mu:     &sync.RWMutex{},
		ttl:    ttl,
		logger: logger,
	}
}

func (r *InMemoryCache) GoDeleteFromCache(_ context.Context, keys ...string) {
	go func() {
		for _, key := range keys {
			r.mu.Lock()
			delete(r.cache, key)
			r.mu.Unlock()
		}
	}()
}

func (r *InMemoryCache) GoAddToCache(_ context.Context, key string, value interface{}) {
	go func() {
		r.mu.Lock()
		defer r.mu.Unlock()
		r.cache[key] = data{
			expirationTime: time.Now().Add(r.ttl),
			value:          value,
		}
	}()
}

func (r *InMemoryCache) GetFromCache(ctx context.Context, key string) (interface{}, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	res, ok := r.cache[key]
	if !ok {
		return nil, errors.New("no values with such key in cache")
	}
	if res.expirationTime.Before(time.Now()) {
		r.GoDeleteFromCache(ctx, key)
		return nil, errors.New("value expired")
	}
	return res, nil
}

func (r *InMemoryCache) Close() error {
	return nil
}
