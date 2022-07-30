package cache

import (
	"fmt"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/storage/cache/gocache"
)

func New(name string, options map[string]interface{}) (storage.Cache, error) {
	switch name {
	case "gocache":
		return gocache.New(options)
	default:
		panic(fmt.Sprintf("%s is a kv database that has not yet been implemented", name))
	}
}
