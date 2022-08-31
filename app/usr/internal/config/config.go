package config

import (
	"github.com/iyear/searchx/pkg/configs"
	"github.com/iyear/searchx/pkg/logger"
	"github.com/iyear/searchx/pkg/storage"
)

var C config

func Init(path string) error {
	return configs.Init(path, &C)
}

type config struct {
	App struct {
		ID   int    `mapstructure:"id" validate:"required" default:"15055931"`
		Hash string `mapstructure:"hash" validate:"required" default:"021d433426cbb920eeb95164498fe3d3"`
	} `mapstructure:"app"`
	Bot struct {
		Token string `mapstructure:"token" validate:"required"`
	} `mapstructure:"bot"`
	Proxy   string         `mapstructure:"proxy" validate:"omitempty,url"`
	Storage storage.Config `mapstructure:"storage"`
	Log     logger.Config  `mapstructure:"log"`
	Ctrl    struct {
		I18N     string `mapstructure:"i18n" validate:"dir" default:"config/usr/i18n"`
		Language string `mapstructure:"language" validate:"iso6391" default:"zh-cn"`
		Search   struct {
			PageSize int `mapstructure:"page_size" validate:"gte=1,lte=20" default:"10"`
		} `mapstructure:"search"`
	} `mapstructure:"ctrl"`
}
