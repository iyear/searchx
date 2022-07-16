package group

import (
	"github.com/iyear/searchx/app/bot/internal/util"
	tele "gopkg.in/telebot.v3"
)

func OnAdded(c tele.Context) error {
	sp := util.GetScope(c)

	// 不允许加入普通群组
	if c.Chat().Type == tele.ChatGroup {
		if err := c.Send(sp.Template.Text.AddedToGroup.Fail.T(nil), &tele.SendOptions{DisableWebPagePreview: true}); err != nil {
			return err
		}
		return c.Bot().Leave(c.Chat())
	}

	return c.Send(sp.Template.Text.AddedToGroup.Success.T(nil), &tele.SendOptions{DisableWebPagePreview: true})
}
