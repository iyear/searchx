package config

import (
	"github.com/gotd/td/telegram"
	"github.com/iyear/searchx/global"
	"runtime"
)

var (
	Device = telegram.DeviceConfig{
		DeviceModel:   "SearchX",
		SystemVersion: runtime.GOOS,
		AppVersion:    global.Version,
	}
)

const (
	ContextScope  = "scope"
	ContextClient = "client"
)
