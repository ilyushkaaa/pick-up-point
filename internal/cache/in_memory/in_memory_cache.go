package cache

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Node struct {
	prev, next     *Node
	key            string
	value          interface{}
	expirationTime time.Time
}

func newNode(key string, val interface{}, expirationTime time.Time) *Node {
	return &Node{
		key:            key,
		value:          val,
		expirationTime: expirationTime,
	}
}

type InMemoryCache struct {
	cache      map[string]*Node
	head, tail *Node
	capacity   int
	mu         *sync.RWMutex
	ttl        time.Duration
	logger     *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger, ttl time.Duration, capacity int) *InMemoryCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return &InMemoryCache{
		cache:    make(map[string]*Node),
		mu:       &sync.RWMutex{},
		ttl:      ttl,
		logger:   logger,
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

func (r *InMemoryCache) GoDeleteFromCache(_ context.Context, keys ...string) {
	go func() {
		for _, key := range keys {
			r.mu.Lock()
			node, ok := r.cache[key]
			if ok {
				r.remove(node)
			}
			r.mu.Unlock()
		}
	}()
}

func (r *InMemoryCache) GoAddToCache(_ context.Context, key string, value interface{}) {
	go func() {
		r.mu.Lock()
		defer r.mu.Unlock()
		if _, ok := r.cache[key]; ok {
			r.remove(r.cache[key])
		}

		if len(r.cache) == r.capacity {
			r.remove(r.tail.prev)
		}

		r.insert(newNode(key, value, time.Now().Add(r.ttl)))
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
	r.remove(res)
	r.insert(res)
	return res, nil
}

func (r *InMemoryCache) Close() error {
	return nil
}

func (r *InMemoryCache) remove(node *Node) {
	delete(r.cache, node.key)
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (r *InMemoryCache) insert(node *Node) {
	r.cache[node.key] = node
	next := r.head.next
	r.head.next = node
	node.prev = r.head
	next.prev = node
	node.next = next
}
