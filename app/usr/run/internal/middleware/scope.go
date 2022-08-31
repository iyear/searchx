package middleware

import (
	"github.com/iyear/searchx/app/usr/run/internal/conf"
	"github.com/iyear/searchx/app/usr/run/internal/model"
	"github.com/iyear/searchx/pkg/consts"
	"github.com/iyear/searchx/pkg/searchbot"
	tele "gopkg.in/telebot.v3"
)

func SetScope(sp *model.BotScope) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			c.Set(conf.ContextScope, sp)
			c.Set(consts.ContextSearch, &searchbot.SearchScope{
				Text:    &sp.Template.Text.Search,
				Button:  &sp.Template.Button.Search,
				Storage: sp.Storage,
			})
			return next(c)
		}
	}
}
