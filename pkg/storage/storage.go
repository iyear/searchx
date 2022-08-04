package storage

type Config struct {
	KV struct {
		Driver  string                 `mapstructure:"driver" validate:"oneof=bolt" default:"bolt"`
		Options map[string]interface{} `mapstructure:"options"`
	} `mapstructure:"kv"`
	Search struct {
		Driver  string                 `mapstructure:"driver" validate:"oneof=bleve" default:"bleve"`
		Options map[string]interface{} `mapstructure:"options"`
	} `mapstructure:"search"`
	Cache struct {
		Driver  string                 `mapstructure:"driver" validate:"oneof=gocache" default:"gocache"`
		Options map[string]interface{} `mapstructure:"options"`
	} `mapstructure:"cache"`
}
