package middleware

import (
	"github.com/iyear/searchx/app/usr/run/internal/config"
	"github.com/iyear/searchx/app/usr/run/internal/model"
	tele "gopkg.in/telebot.v3"
)

func SetScope(sp *model.Scope) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			c.Set(config.ContextScope, sp)
			return next(c)
		}
	}
}
