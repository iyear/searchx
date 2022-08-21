package configs

import (
	"github.com/creasty/defaults"
	"github.com/iyear/searchx/pkg/validator"
	"github.com/spf13/viper"
)

func Init(path string, v interface{}) error {
	c := viper.New()
	c.SetConfigFile(path)
	if err := c.ReadInConfig(); err != nil {
		return err
	}

	if err := c.Unmarshal(v); err != nil {
		return err
	}

	if err := defaults.Set(v); err != nil {
		return err
	}

	if err := validator.Struct(v); err != nil {
		return err
	}

	return nil
}
