package run

import (
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/usr/run/internal/config"
	"github.com/iyear/searchx/app/usr/run/internal/login"
	"github.com/iyear/searchx/global"
	"log"
)

func Run(cfg string, _login bool) {
	color.Blue(global.Logo)
	color.Blue("Initializing...")

	if err := config.Init(cfg); err != nil {
		log.Fatalf("init config failed: %v", err)
	}
	color.Blue("Config loaded")

	if _login {
		if err := login.Start(); err != nil {
			color.Red("Login failed: %v", err)
		}
		return
	}

	// slog := logger.New(config.C.Log.Enable, "log/usr/latest.log", config.C.Log.Level)
}
