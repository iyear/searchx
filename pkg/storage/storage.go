package storage

type Storage struct {
	KV     KV
	Search Search
	Cache  Cache
}

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

func Init(conf Config) (s Search, k KV, c Cache, err error) {
	if s, err = NewSearch(conf.Search.Driver, conf.Search.Options); err != nil {
		return
	}

	if k, err = NewKV(conf.KV.Driver, conf.KV.Options); err != nil {
		return
	}

	if c, err = NewCache(conf.Cache.Driver, conf.Cache.Options); err != nil {
		return
	}

	return
}
