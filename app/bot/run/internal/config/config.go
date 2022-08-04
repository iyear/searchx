package config

import (
	"github.com/creasty/defaults"
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
		Token  string `mapstructure:"token" validate:"required"`
		Socks5 struct {
			Enable   bool   `mapstructure:"enable"`
			Host     string `mapstructure:"host" validate:"hostname" default:"localhost"`
			Port     int    `mapstructure:"port" default:"1080"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
		} `mapstructure:"socks5"`
		Admin []int64 `mapstructure:"admin" validate:"required"`
	} `mapstructure:"bot"`
	Storage storage.Config `mapstructure:"storage"`
	Ctrl    struct {
		Notice          string `mapstructure:"notice" default:"NO NOTICE"`
		I18N            string `mapstructure:"i18n" validate:"dir" default:"config/i18n"`
		DefaultLanguage string `mapstructure:"default_language" validate:"iso6391" default:"zh-cn"`
		Search          struct {
			PageSize int `mapstructure:"page_size" validate:"gte=1,lte=20" default:"10"`
		} `mapstructure:"search"`
	} `mapstructure:"ctrl"`
}
