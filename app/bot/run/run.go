package run

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/bot/internal/config"
	"github.com/iyear/searchx/app/bot/run/internal/i18n"
	"github.com/iyear/searchx/app/bot/run/internal/middleware"
	"github.com/iyear/searchx/app/bot/run/internal/model"
	"github.com/iyear/searchx/global"
	"github.com/iyear/searchx/pkg/botmid"
	"github.com/iyear/searchx/pkg/logger"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/utils"
	tele "gopkg.in/telebot.v3"
	"net/http"
	"time"
)

func Run(cfg string) error {
	color.Blue(global.Logo)
	color.Blue("Initializing...")

	if err := config.Init(cfg); err != nil {
		return fmt.Errorf("init config failed: %v", err)
	}
	color.Blue("Config loaded")

	slog := logger.New(config.C.Log.Enable, "log/bot/latest.log", config.C.Log.Level)

	if err := i18n.Init(config.C.Ctrl.I18N); err != nil {
		return fmt.Errorf("init i18n failed: %v", err)
	}
	color.Blue("I18n templates loaded")

	search, kv, cache, err := storage.Init(config.C.Storage)
	if err != nil {
		return fmt.Errorf("init storage failed: %v", err)
	}
	color.Blue("Storage initialized")

	dialer, err := utils.ProxyFromURL(config.C.Proxy)
	if err != nil {
		return fmt.Errorf("init proxy failed: %v", err)
	}

	settings := tele.Settings{
		Token:     config.C.Bot.Token,
		Poller:    &tele.LongPoller{Timeout: 5 * time.Second},
		Client:    &http.Client{Transport: &http.Transport{DialContext: dialer.DialContext}},
		OnError:   middleware.OnError(),
		ParseMode: tele.ModeHTML,
	}

	bot, err := tele.NewBot(settings)
	if err != nil {
		return fmt.Errorf("init bot failed: %v", err)
	}

	color.Blue("Bot: %s", bot.Me.Username)

	scope := &model.Scope{
		Storage: &storage.Storage{
			KV:     kv,
			Search: search,
			Cache:  cache,
		},
		Log: slog.Named("bot"),
	}

	bot.Use(middleware.SetScope(scope), botmid.AutoResponder())

	makeHandlers(bot, i18n.Templates[config.C.Ctrl.DefaultLanguage].Button)

	bot.Start()
	return nil
}
