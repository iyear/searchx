package util

import (
	"github.com/iyear/searchx/app/usr/run/internal/config"
	"github.com/iyear/searchx/app/usr/run/internal/model"
	tele "gopkg.in/telebot.v3"
)

func GetBotScope(c tele.Context) *model.BotScope {
	return c.Get(config.ContextScope).(*model.BotScope)
}
