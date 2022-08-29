package run

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/gotd/contrib/middleware/floodwait"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/telegram/updates"
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/app/usr/internal/config"
	"github.com/iyear/searchx/app/usr/internal/sto"
	"github.com/iyear/searchx/app/usr/run/internal/conf"
	"github.com/iyear/searchx/app/usr/run/internal/i18n"
	"github.com/iyear/searchx/app/usr/run/internal/middleware"
	"github.com/iyear/searchx/app/usr/run/internal/model"
	"github.com/iyear/searchx/global"
	"github.com/iyear/searchx/pkg/botmid"
	"github.com/iyear/searchx/pkg/logger"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/utils"
	tele "gopkg.in/telebot.v3"
	"net/http"
	"time"
)

func Run(ctx context.Context, cfg string) error {
	color.Blue(global.Logo)
	color.Blue("Initializing...")

	if err := config.Init(cfg); err != nil {
		return fmt.Errorf("init config failed: %v", err)
	}
	color.Blue("Config loaded")

	slog := logger.New(config.C.Log.Enable, "log/usr/latest.log", config.C.Log.Level)

	if err := i18n.Init(config.C.Ctrl.I18N); err != nil {
		return fmt.Errorf("init i18n failed: %v", err)
	}
	color.Blue("I18n templates loaded")

	search, kv, cache, err := storage.Init(config.C.Storage)
	if err != nil {
		return fmt.Errorf("init storage failed: %v", err)
	}
	color.Blue("Storage initialized")
	_storage := &storage.Storage{
		KV:     kv,
		Search: search,
		Cache:  cache,
	}

	dialer, err := utils.ProxyFromURL(config.C.Proxy)
	if err != nil {
		return fmt.Errorf("init proxy failed: %v", err)
	}

	// init bot
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
	color.Blue("Auth successfully! Bot: %s", bot.Me.Username)

	template, ok := i18n.Templates[config.C.Ctrl.Language]
	if !ok {
		return fmt.Errorf("language [%s] is not supported", config.C.Ctrl.Language)
	}

	botScope := &model.BotScope{
		Storage:  _storage,
		Template: template,
		Log:      slog.Named("bot"),
	}
	bot.Use(middleware.SetScope(botScope), botmid.AutoResponder())

	// init usr
	dispatcher := tg.NewUpdateDispatcher()
	handleUsr(&dispatcher)

	gaps := updates.New(updates.Config{
		Handler:      dispatcher,
		Logger:       slog.Named("updates").Desugar(),
		Storage:      sto.NewState(kv),
		AccessHasher: sto.NewAccessHasher(kv),
		OnChannelTooLong: func(channelID int64) {
			slog.Errorw("channel is too long", "channelID", channelID)
		},
	})

	c := telegram.NewClient(config.C.App.ID, config.C.App.Hash, telegram.Options{
		Resolver: dcs.Plain(dcs.PlainOptions{
			Dial: dialer.DialContext,
		}),
		SessionStorage: sto.NewSession(kv, false),
		Logger:         slog.Named("telegram").Desugar(),
		UpdateHandler:  gaps,
		RetryInterval:  time.Second,
		Middlewares: []telegram.Middleware{
			floodwait.NewSimpleWaiter(),
		},
		MaxRetries: conf.MaxRetries,
	})

	usrScope := &model.UsrScope{
		Storage: _storage,
		Log:     slog.Named("usr"),
		Client:  c,
	}
	return c.Run(context.WithValue(ctx, conf.ContextScope, usrScope), func(ctx context.Context) error {
		status, err := c.Auth().Status(ctx)
		if err != nil {
			return err
		}
		if !status.Authorized {
			return fmt.Errorf("not authorized. please login first")
		}

		// set handler and start bot
		// only `self` can use the bot
		bot.Use(botmid.WhiteList(status.User.ID))
		handleBot(bot, template.Button)
		go bot.Start()
		defer bot.Stop()

		// notify update manager about authentication.
		// isBot set to `true` to avoid too long updates diff and can't fetch old messages.
		// forgot set to `false` to avoid replace local pts by remote pts.
		if err := gaps.Auth(ctx, c.API(), status.User.ID, true, false); err != nil {
			return err
		}
		defer func() { _ = gaps.Logout() }()

		color.Blue("Auth successfully! User: %s", status.User.Username)

		<-ctx.Done()
		return ctx.Err()
	})
}
