package storage

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
}
