package storage

type KV interface {
	Get(key string) (string, error)
	Set(key string, val string) error
}
