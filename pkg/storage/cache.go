package storage

import (
	"context"
	"fmt"
	"github.com/iyear/searchx/pkg/storage/cache/gocache"
)

type Cache interface {
	Get(ctx context.Context, key string) (interface{}, bool)
	Set(ctx context.Context, key string, val interface{})
}

func NewCache(name string, options map[string]interface{}) (Cache, error) {
	switch name {
	case "gocache":
		return gocache.New(options)
	default:
		panic(fmt.Sprintf("%s is a kv database that has not yet been implemented", name))
	}
}
