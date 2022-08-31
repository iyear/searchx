package util

import (
	"context"
	"github.com/iyear/searchx/app/bot/internal/config"
	"github.com/iyear/searchx/app/bot/run/internal/conf"
	"github.com/iyear/searchx/app/bot/run/internal/key"
	"github.com/iyear/searchx/app/bot/run/internal/model"
	"github.com/iyear/searchx/pkg/storage"
	tele "gopkg.in/telebot.v3"
)

func GetScope(c tele.Context) *model.Scope {
	return c.Get(conf.ContextScope).(*model.Scope)
}

func GetUserLanguage(storage *storage.Storage, tid int64) string {
	v, found := storage.Cache.Get(context.Background(), key.Language(tid))
	if found {
		return v.(string)
	}

	lang, err := storage.KV.Get(context.TODO(), key.Language(tid))
	if err != nil {
		lang = config.C.Ctrl.DefaultLanguage
	}

	storage.Cache.Set(context.Background(), key.Language(tid), lang)
	return lang
}
