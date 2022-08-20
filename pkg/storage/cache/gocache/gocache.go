package gocache

import (
	"context"
	"github.com/patrickmn/go-cache"
	"time"
)

type Cache struct {
	cache *cache.Cache
}

func New(options map[string]interface{}) (*Cache, error) {
	_ = options
	return &Cache{cache: cache.New(10*time.Minute, 1*time.Minute)}, nil
}

func (c *Cache) Get(_ context.Context, key string) (interface{}, bool) {
	return c.cache.Get(key)
}

func (c *Cache) Set(_ context.Context, key string, val interface{}) {
	c.cache.Set(key, val, cache.DefaultExpiration)
}
