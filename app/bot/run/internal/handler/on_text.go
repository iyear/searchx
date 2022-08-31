package handler

import (
	"github.com/iyear/searchx/app/bot/internal/config"
	"github.com/iyear/searchx/app/bot/run/internal/handler/group"
	"github.com/iyear/searchx/pkg/searchbot"
	tele "gopkg.in/telebot.v3"
)

func OnText(c tele.Context) error {
	if c.Chat().Type == tele.ChatSuperGroup {
		return group.OnText(c)
	}
	return searchbot.Search(config.C.Ctrl.Search.PageSize)(c)
}
