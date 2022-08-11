package run

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/telegram/updates"
	updhook "github.com/gotd/td/telegram/updates/hook"
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

	kv, err := storage.NewKV(config.C.Storage.KV.Driver, config.C.Storage.KV.Options)
	if err != nil {
		slog.Fatalw("init kv database failed", "err", err, "options", config.C.Storage.KV.Options)
	}
	color.Blue("KV database initialized")

	search, err := storage.NewSearch(config.C.Storage.Search.Driver, config.C.Storage.Search.Options)
	if err != nil {
		slog.Fatalw("init search engine database failed", "err", err, "options", config.C.Storage.Search.Options)
	}
	color.Blue("Search engine initialized")

	cache, err := storage.New(config.C.Storage.Cache.Driver, config.C.Storage.Cache.Options)
	if err != nil {
		slog.Fatalw("init cache failed", "err", err, "options", config.C.Storage.Cache.Options)
	}
	color.Blue("Cache initialized")

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
		ParseMode: tele.ModeMarkdown,
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
	_storage := &model.Storage{
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
	})

	c := telegram.NewClient(config.C.Account.ID, config.C.Account.Hash, telegram.Options{
		Resolver: dcs.Plain(dcs.PlainOptions{
			Dial: dialer.DialContext,
		}),
		SessionStorage: sto.NewSession(kv, false),
		Logger:         zap.NewNop(),
		UpdateHandler:  gaps,
		Middlewares: []telegram.Middleware{
			updhook.UpdateHook(gaps.Handle),
		},
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
