package util

import (
	"context"
	"github.com/gotd/td/telegram"
	"github.com/iyear/searchx/app/usr/run/internal/config"
	"github.com/iyear/searchx/app/usr/run/internal/model"
	tele "gopkg.in/telebot.v3"
)

func GetScope(c tele.Context) *model.Scope {
	return c.Get(config.ContextScope).(*model.Scope)
}

func GetClient(c context.Context) *telegram.Client {
	return c.Value(config.ContextClient).(*telegram.Client)
}
