package bot

import (
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/bot/internal/config"
	"github.com/iyear/searchx/app/bot/internal/i18n"
	"github.com/iyear/searchx/app/bot/internal/middleware"
	"github.com/iyear/searchx/app/bot/internal/model"
	"github.com/iyear/searchx/global"
	"github.com/iyear/searchx/pkg/logger"
	"github.com/iyear/searchx/pkg/storage/cache"
	"github.com/iyear/searchx/pkg/storage/kv"
	"github.com/iyear/searchx/pkg/storage/search"
	tele "gopkg.in/telebot.v3"
	"time"
)

func Run(cfg string) {
	log := logger.Init()

	if err := config.Init(cfg); err != nil {
		log.Fatalw("init config failed", "err", err)
	}

	if err := i18n.Init(config.C.Ctrl.I18N); err != nil {
		log.Fatalw("init i18n templates failed", "err", err)
	}

	_kv, err := kv.New(config.C.Storage.KV.Driver, config.C.Storage.KV.Options)
	if err != nil {
		log.Fatalw("init kv database failed", "err", err, "options", config.C.Storage.KV.Options)
	}

	_search, err := search.New(config.C.Storage.Search.Driver, config.C.Storage.Search.Options)
	if err != nil {
		log.Fatalw("init search engine database failed", "err", err, "options", config.C.Storage.Search.Options)
	}

	_cache, err := cache.New(config.C.Storage.Cache.Driver, config.C.Storage.Cache.Options)
	if err != nil {
		log.Fatalw("init cache failed", "err", err, "options", config.C.Storage.Cache.Options)
	}

	settings := tele.Settings{
		Token:     config.C.Bot.Token,
		Poller:    &tele.LongPoller{Timeout: 5 * time.Second},
		Client:    getClient(),
		OnError:   middleware.OnError(),
		ParseMode: tele.ModeMarkdown,
	}

	bot, err := tele.NewBot(settings)
	if err != nil {
		log.Fatalw("create bot failed", "err", err)
	}

	color.Blue(global.Logo)
	color.Blue("Bot: %s", bot.Me.Username)

	scope := &model.Scope{
		Storage: &model.Storage{
			KV:     _kv,
			Search: _search,
			Cache:  _cache,
		},
		Log: log.Named("bot"),
	}

	bot.Use(middleware.SetScope(scope), middleware.AutoResponder())

	makeHandlers(bot, i18n.Templates[config.C.Ctrl.DefaultLanguage].Button)

	bot.Start()
}
