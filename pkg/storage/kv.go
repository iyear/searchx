package storage

import (
	"fmt"
	"github.com/iyear/searchx/pkg/storage/kv/bolt"
)

type KV interface {
	// Get returns kv.ErrNotFound if the key does not exist
	Get(key string) (string, error)
	Set(key string, val string) error
}

func NewKV(name string, options map[string]interface{}) (KV, error) {
	switch name {
	case "bolt":
		return bolt.New(options)
	default:
		panic(fmt.Sprintf("%s is a kv database that has not yet been implemented", name))
	}
}
