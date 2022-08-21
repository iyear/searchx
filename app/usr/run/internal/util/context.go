package util

import (
	"context"
	"github.com/iyear/searchx/app/usr/run/internal/conf"
	"github.com/iyear/searchx/app/usr/run/internal/model"
	tele "gopkg.in/telebot.v3"
)

func GetBotScope(c tele.Context) *model.BotScope {
	return c.Get(conf.ContextScope).(*model.BotScope)
}

func GetUsrScope(c context.Context) *model.UsrScope {
	return c.Value(conf.ContextScope).(*model.UsrScope)
}
