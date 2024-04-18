package cache

import (
	"context"
	"time"

	modelOrder "homework/internal/order/model"
	modelPP "homework/internal/pick-up_point/model"
)

type InMemoryCache struct {
	pickUpPointsByID map[uint64]modelPP.PickUpPoint
	ordersByID       map[uint64]modelOrder.Order
	ordersByClientID map[uint64][]modelOrder.Order
	expireTime       time.Duration
}

func New() *InMemoryCache {
	return &InMemoryCache{
		pickUpPointsByID: make(map[uint64]modelPP.PickUpPoint),
		ordersByID:       make(map[uint64]modelOrder.Order),
		ordersByClientID: make(map[uint64][]modelOrder.Order),
		expireTime:       time.Minute * 10,
	}
}

func (r *InMemoryCache) GoDeleteFromCache(_ context.Context, keys ...string) {
	go func() {

	}()
}

func (r *InMemoryCache) GoAddToCache(_ context.Context, key string, value interface{}) {
	go func() {
		r.goDeleteByExpirationTime()

	}()
}

func (r *InMemoryCache) GetFromCache(_ context.Context, key string) (interface{}, error) {
	return nil, nil
}

func (r *InMemoryCache) Close() error {
	return nil
}

func (r *InMemoryCache) goDeleteByExpirationTime(key uint64) {
	go func() {
		<-time.After(r.expireTime)
		delete(r.pickUpPointsByID, key)
	}()
}
