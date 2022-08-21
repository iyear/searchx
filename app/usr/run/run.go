package run

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/telegram/updates"
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/app/usr/run/internal/config"
	"github.com/iyear/searchx/app/usr/run/internal/i18n"
	"github.com/iyear/searchx/app/usr/run/internal/login"
	"github.com/iyear/searchx/app/usr/run/internal/middleware"
	"github.com/iyear/searchx/app/usr/run/internal/model"
	"github.com/iyear/searchx/app/usr/run/internal/sto"
	"github.com/iyear/searchx/global"
	"github.com/iyear/searchx/pkg/logger"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/utils"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
	"log"
	"net/http"
	"time"
)

func Run(ctx context.Context, cfg string, _login bool) error {
	color.Blue(global.Logo)
	color.Blue("Initializing...")

	if err := config.Init(cfg); err != nil {
		log.Fatalf("init config failed: %v", err)
	}
	color.Blue("Config loaded")

	slog := logger.New(config.C.Log.Enable, "log/usr/latest.log", config.C.Log.Level)

	if err := i18n.Init(config.C.Ctrl.I18N); err != nil {
		slog.Fatalw("init i18n templates failed", "err", err)
	}
	color.Blue("I18n templates loaded")

	search, kv, cache, err := storage.Init(config.C.Storage)
	if err != nil {
		slog.Fatalw("init storage failed", "err", err)
	}
	color.Blue("Storage initialized")

	dialer, err := utils.ProxyFromURL(config.C.Proxy)
	if err != nil {
		slog.Fatalw("init proxy failed", "err", err, "proxy", config.C.Proxy)
	}

	if _login {
		if err = login.Start(kv, dialer); err != nil {
			return fmt.Errorf("login failed: %v", err)
		}
		return nil
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
		slog.Fatalw("create bot failed", "err", err)
	}
	color.Blue("Auth successfully! Bot: %s", bot.Me.Username)

	template, ok := i18n.Templates[config.C.Ctrl.Language]
	if !ok {
		slog.Fatalw("language is not supported", "language", config.C.Ctrl.Language)
	}

	_storage := &storage.Storage{
		KV:     kv,
		Search: search,
		Cache:  cache,
	}

	botScope := &model.BotScope{
		Storage:  _storage,
		Template: template,
		Log:      slog.Named("bot"),
	}

	bot.Use(middleware.SetScope(botScope), middleware.AutoResponder())

	handleBot(bot, template.Button)

	go bot.Start()
	defer bot.Stop()

	// init usr
	dispatcher := tg.NewUpdateDispatcher()
	handleUsr(&dispatcher)

	gaps := updates.New(updates.Config{
		Handler:      dispatcher,
		Logger:       zap.NewNop(),
		Storage:      sto.NewState(kv),
		AccessHasher: sto.NewAccessHasher(kv),
		OnChannelTooLong: func(channelID int64) {
			slog.Errorw("channel is too long", "channelID", channelID)
		},
	})

	c := telegram.NewClient(config.C.Account.ID, config.C.Account.Hash, telegram.Options{
		Resolver: dcs.Plain(dcs.PlainOptions{
			Dial: dialer.DialContext,
		}),
		SessionStorage: sto.NewSession(kv, false),
		Logger:         zap.NewNop(),
		UpdateHandler:  gaps,
	})

	usrScope := &model.UsrScope{
		Storage: _storage,
		Log:     slog.Named("usr"),
		Client:  c,
	}
	return c.Run(context.WithValue(ctx, config.ContextScope, usrScope), func(ctx context.Context) error {
		status, err := c.Auth().Status(ctx)
		if err != nil {
			return err
		}
		if !status.Authorized {
			return fmt.Errorf("not authorized. please login first")
		}

		// Notify update manager about authentication.
		if err := gaps.Auth(ctx, c.API(), status.User.ID, true, false); err != nil {
			return err
		}
		defer func() { _ = gaps.Logout() }()

		color.Blue("Auth successfully! User: %s", status.User.Username)

		<-ctx.Done()
		return ctx.Err()
	})
}
