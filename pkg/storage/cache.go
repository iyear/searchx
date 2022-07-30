package storage

import (
	"fmt"
	"github.com/iyear/searchx/pkg/storage/cache/gocache"
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
}

func New(name string, options map[string]interface{}) (Cache, error) {
	switch name {
	case "gocache":
		return gocache.New(options)
	default:
		panic(fmt.Sprintf("%s is a kv database that has not yet been implemented", name))
	}
}
