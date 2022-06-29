package kv

import (
	"fmt"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/storage/kv/internal/bolt"
)

func New(name string, options map[string]interface{}) (storage.KV, error) {
	switch name {
	case "bolt":
		return bolt.New(options)
	default:
		panic(fmt.Sprintf("%s is a kv database that has not yet been implemented", name))
	}
}
