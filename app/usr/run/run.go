package run

import (
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/usr/run/internal/config"
	"github.com/iyear/searchx/global"
	"log"
)

func Run(cfg string, login bool) {
	color.Blue(global.Logo)
	color.Blue("Initializing...")

	if err := config.Init(cfg); err != nil {
		log.Fatalf("init config failed: %v", err)
	}
	color.Blue("Config loaded")

	if login {
		startLogin()
		return
	}

	// slog := logger.New(config.C.Ctrl.Log.Enable, "log/usr/latest.log", config.C.Ctrl.Log.Level)

}
