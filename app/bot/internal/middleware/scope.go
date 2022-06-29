package middleware

import (
	"github.com/iyear/searchx/app/bot/internal/config"
	"github.com/iyear/searchx/app/bot/internal/i18n"
	"github.com/iyear/searchx/app/bot/internal/model"
	"github.com/iyear/searchx/app/bot/internal/util"
	tele "gopkg.in/telebot.v3"
)

func SetScope(sp *model.Scope) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			lang := util.GetUserLanguage(sp.Storage, getID(c))
			tmpl, ok := i18n.Templates[lang]
			if !ok {
				return nil
			}

			c.Set(config.ContextScope, &model.Scope{
				Storage:  sp.Storage,
				Template: tmpl,
				Log:      sp.Log,
			})

			c.Set(config.ContextLanguage, lang)
			return next(c)
		}
	}
}

func getID(c tele.Context) int64 {
	if c.Chat() != nil {
		return c.Chat().ID
	}
	if c.Query() != nil {
		return c.Query().Sender.ID
	}
	if c.Message() != nil {
		return c.Message().Sender.ID
	}
	return 0
}
