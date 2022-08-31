package bolt

import (
	"context"
	"github.com/creasty/defaults"
	"github.com/iyear/searchx/pkg/storage/kv"
	"github.com/iyear/searchx/pkg/validator"
	"github.com/mitchellh/mapstructure"
	"go.etcd.io/bbolt"
	"os"
	"time"
)

type Options struct {
	Path string `mapstructure:"path" default:"data/data.kv"`
}

type Bolt struct {
	db *bbolt.DB
}

const bucket = "b"

func New(options map[string]interface{}) (*Bolt, error) {
	var ops Options

	if err := mapstructure.Decode(options, &ops); err != nil {
		return nil, err
	}

	if err := defaults.Set(&ops); err != nil {
		return nil, err
	}

	if err := validator.Struct(&ops); err != nil {
		return nil, err
	}

	db, err := bbolt.Open(ops.Path, os.ModePerm, &bbolt.Options{
		Timeout:      time.Second,
		NoGrowSync:   false,
		FreelistType: bbolt.FreelistArrayType,
	})
	if err != nil {
		return nil, err
	}

	if err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	}); err != nil {
		return nil, err
	}

	return &Bolt{db: db}, nil
}

func (b *Bolt) Get(_ context.Context, key string) (string, error) {
	var val []byte

	if err := b.db.View(func(tx *bbolt.Tx) error {
		val = tx.Bucket([]byte(bucket)).Get([]byte(key))
		return nil
	}); err != nil {
		return "", err
	}

	if val == nil {
		return "", kv.ErrNotFound
	}

	return string(val), nil
}

func (b *Bolt) Set(_ context.Context, key string, val string) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Put([]byte(key), []byte(val))
	})
}
