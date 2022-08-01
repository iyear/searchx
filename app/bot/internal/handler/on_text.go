package handler

import (
	"github.com/iyear/searchx/app/bot/internal/handler/group"
	"github.com/iyear/searchx/app/bot/internal/handler/private"
	tele "gopkg.in/telebot.v3"
)

func OnText(c tele.Context) error {
	if c.Chat().Type == tele.ChatSuperGroup {
		return group.OnText(c)
	}
	return private.Search(c)
}
