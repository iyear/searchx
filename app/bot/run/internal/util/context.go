package util

import (
	config2 "github.com/iyear/searchx/app/bot/run/internal/config"
	"github.com/iyear/searchx/app/bot/run/internal/key"
	"github.com/iyear/searchx/app/bot/run/internal/model"
	tele "gopkg.in/telebot.v3"
)

func GetScope(c tele.Context) *model.Scope {
	return c.Get(config2.ContextScope).(*model.Scope)
}

func GetUserLanguage(storage *model.Storage, tid int64) string {
	v, found := storage.Cache.Get(key.Language(tid))
	if found {
		return v.(string)
	}

	lang, err := storage.KV.Get(key.Language(tid))
	if err != nil {
		lang = config2.C.Ctrl.DefaultLanguage
	}

	storage.Cache.Set(key.Language(tid), lang)
	return lang
}
