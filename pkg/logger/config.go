package logger

type Config struct {
	Enable bool   `mapstructure:"enable" default:"true"`
	Level  string `mapstructure:"level" validate:"oneof=debug info warn error fatal" default:"info"`
}
