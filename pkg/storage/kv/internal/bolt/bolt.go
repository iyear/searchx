package bolt

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go.etcd.io/bbolt"
	"os"
)

type Options struct {
	Path string `mapstructure:"path"`
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

	db, err := bbolt.Open(ops.Path, os.ModePerm, bbolt.DefaultOptions)
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

func (b *Bolt) Get(key string) (string, error) {
	var val []byte

	if err := b.db.View(func(tx *bbolt.Tx) error {
		val = tx.Bucket([]byte(bucket)).Get([]byte(key))
		return nil
	}); err != nil {
		return "", err
	}

	if val == nil {
		return "", fmt.Errorf("%s is not found", key)
	}

	return string(val), nil
}

func (b *Bolt) Set(key string, val string) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Put([]byte(key), []byte(val))
	})
}
