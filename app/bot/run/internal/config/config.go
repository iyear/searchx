package config

import (
	"github.com/creasty/defaults"
	"github.com/iyear/searchx/pkg/logger"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/validator"
	"github.com/spf13/viper"
)

var C config

func Init(path string) error {
	c := viper.New()
	c.SetConfigFile(path)
	if err := c.ReadInConfig(); err != nil {
		return err
	}

	if err := c.Unmarshal(&C); err != nil {
		return err
	}

	if err := defaults.Set(&C); err != nil {
		return err
	}

	if err := validator.Struct(&C); err != nil {
		return err
	}

	return nil
}

type config struct {
	Bot struct {
		Token string  `mapstructure:"token" validate:"required"`
		Admin []int64 `mapstructure:"admin" validate:"required"`
	} `mapstructure:"bot"`
	Proxy   string         `mapstructure:"proxy" validate:"omitempty,url"`
	Storage storage.Config `mapstructure:"storage"`
	Log     logger.Config  `mapstructure:"log"`
	Ctrl    struct {
		Notice          string `mapstructure:"notice" default:"NO NOTICE"`
		I18N            string `mapstructure:"i18n" validate:"dir" default:"config/bot/i18n"`
		DefaultLanguage string `mapstructure:"default_language" validate:"iso6391" default:"zh-cn"`
		Search          struct {
			PageSize int `mapstructure:"page_size" validate:"gte=1,lte=20" default:"10"`
		} `mapstructure:"search"`
	} `mapstructure:"ctrl"`
}
