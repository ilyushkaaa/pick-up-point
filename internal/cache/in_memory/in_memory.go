package cache

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

type InMemoryCache struct {
	pickUpPointsByID map[string]interface{}
	ppMu             *sync.RWMutex

	ordersByID map[string]interface{}
	orderMu    *sync.RWMutex

	ordersByClientID map[string]interface{}
	ordersByClientMu *sync.RWMutex

	expireTime time.Duration

	deleteFunctions map[string]func(string)
	addFunctions    map[string]func(string, interface{})
	getFunctions    map[string]func(string) (interface{}, error)

	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger) *InMemoryCache {
	cache := &InMemoryCache{
		pickUpPointsByID: make(map[string]interface{}),
		ppMu:             &sync.RWMutex{},
		ordersByID:       make(map[string]interface{}),
		orderMu:          &sync.RWMutex{},
		ordersByClientID: make(map[string]interface{}),
		ordersByClientMu: &sync.RWMutex{},
		expireTime:       time.Minute * 10,
		logger:           logger,
	}
	deleteFunctions := map[string]func(string){
		"pp":    cache.deleteFromPickUpPointsByID,
		"order": cache.deleteFromOrdersByID,
		"user":  cache.deleteOrdersByClientID,
	}

	addFunctions := map[string]func(string, interface{}){
		"pp":    cache.addPickUpPointsByID,
		"order": cache.addOrdersByID,
		"user":  cache.addOrdersByClientID,
	}

	getFunctions := map[string]func(string) (interface{}, error){
		"pp":    cache.getPickUpPointsByID,
		"order": cache.getOrdersByID,
		"user":  cache.getOrdersByClientID,
	}

	cache.deleteFunctions = deleteFunctions
	cache.addFunctions = addFunctions
	cache.getFunctions = getFunctions

	return cache
}

func (r *InMemoryCache) GoDeleteFromCache(_ context.Context, keys ...string) {
	go func() {
		for _, key := range keys {
			keyParts := strings.Split(key, "_")
			funcToCall, ok := r.deleteFunctions[keyParts[0]]
			if !ok {
				r.logger.Errorf("unknown key %s for cache", key)
			}
			funcToCall(keyParts[1])
		}
	}()
}

func (r *InMemoryCache) GoAddToCache(_ context.Context, key string, value interface{}) {
	go func() {
		keyParts := strings.Split(key, "_")
		funcToCall, ok := r.addFunctions[keyParts[0]]
		if !ok {
			r.logger.Errorf("unknown key %s for cache", key)
		}
		funcToCall(keyParts[1], value)
		r.goDeleteByExpirationTime(key)
	}()
}

func (r *InMemoryCache) GetFromCache(_ context.Context, key string) (interface{}, error) {
	keyParts := strings.Split(key, "_")
	funcToCall, ok := r.getFunctions[keyParts[0]]
	if !ok {
		r.logger.Errorf("unknown key %s for cache", key)
	}
	return funcToCall(keyParts[1])
}

func (r *InMemoryCache) Close() error {
	return nil
}

func (r *InMemoryCache) goDeleteByExpirationTime(key string) {
	go func() {
		<-time.After(r.expireTime)
		r.GoDeleteFromCache(context.Background(), key)
	}()
}

func (r *InMemoryCache) deleteFromPickUpPointsByID(key string) {
	r.ppMu.Lock()
	defer r.ppMu.Unlock()
	delete(r.pickUpPointsByID, key)
}

func (r *InMemoryCache) deleteFromOrdersByID(key string) {
	r.orderMu.Lock()
	defer r.orderMu.Unlock()
	delete(r.ordersByID, key)
}

func (r *InMemoryCache) deleteOrdersByClientID(key string) {
	r.ordersByClientMu.Lock()
	defer r.ordersByClientMu.Unlock()
	delete(r.ordersByClientID, key)
}

func (r *InMemoryCache) addPickUpPointsByID(key string, value interface{}) {
	r.ppMu.Lock()
	defer r.ppMu.Unlock()
	r.pickUpPointsByID[key] = value
}

func (r *InMemoryCache) addOrdersByID(key string, value interface{}) {
	r.orderMu.Lock()
	defer r.orderMu.Unlock()
	r.ordersByID[key] = value
}

func (r *InMemoryCache) addOrdersByClientID(key string, value interface{}) {
	r.ordersByClientMu.Lock()
	defer r.ordersByClientMu.Unlock()
	r.ordersByClientID[key] = value
}

func (r *InMemoryCache) getPickUpPointsByID(key string) (interface{}, error) {
	r.ppMu.RLock()
	defer r.ppMu.RUnlock()
	res, ok := r.pickUpPointsByID[key]
	if !ok {
		return nil, errors.New("no pick-up points with such id in cache")
	}
	return res, nil
}

func (r *InMemoryCache) getOrdersByID(key string) (interface{}, error) {
	r.orderMu.RLock()
	defer r.orderMu.RUnlock()
	res, ok := r.ordersByID[key]
	if !ok {
		return nil, errors.New("no pick-up points with such id in cache")
	}
	return res, nil
}

func (r *InMemoryCache) getOrdersByClientID(key string) (interface{}, error) {
	r.ordersByClientMu.RLock()
	defer r.ordersByClientMu.RUnlock()
	res, ok := r.ordersByClientID[key]
	if !ok {
		return nil, errors.New("no pick-up points with such id in cache")
	}
	return res, nil
}
