package config

import "github.com/spf13/viper"

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

	return nil
}

type config struct {
	Bot struct {
		Token  string `mapstructure:"token"`
		Socks5 struct {
			Enable   bool   `mapstructure:"enable"`
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
		} `mapstructure:"socks5"`
		Admin []int64 `mapstructure:"admin"`
	} `mapstructure:"bot"`
	Storage struct {
		KV struct {
			Driver  string                 `mapstructure:"driver"`
			Options map[string]interface{} `mapstructure:"options"`
		} `mapstructure:"kv"`
		Search struct {
			Driver  string                 `mapstructure:"driver"`
			Options map[string]interface{} `mapstructure:"options"`
		} `mapstructure:"search"`
		Cache struct {
			Driver  string                 `mapstructure:"driver"`
			Options map[string]interface{} `mapstructure:"options"`
		} `mapstructure:"cache"`
	}
	Ctrl struct {
		Notice          string `mapstructure:"notice"`
		I18N            string `mapstructure:"i18n"`
		DefaultLanguage string `mapstructure:"default_language"`
		Search          struct {
			PageSize int `mapstructure:"page_size"`
		} `mapstructure:"search"`
	} `mapstructure:"ctrl"`
}
