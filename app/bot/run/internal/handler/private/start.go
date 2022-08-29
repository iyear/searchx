package private

import (
	"github.com/iyear/searchx/app/bot/internal/config"
	"github.com/iyear/searchx/app/bot/run/internal/model"
	"github.com/iyear/searchx/app/bot/run/internal/util"
	"github.com/iyear/searchx/global"
	"github.com/iyear/searchx/pkg/searchbot"
	tele "gopkg.in/telebot.v3"
)

func Start(c tele.Context) error {
	sp := util.GetScope(c)

	// 返回首页则重置
	if c.Message().Payload == "" {
		chat := c.Chat()

		return c.EditOrSend(sp.Template.Text.Start.T(&model.TStart{
			ID:       chat.ID,
			Username: chat.Username,
			Notice:   config.C.Ctrl.Notice,
			// Chats:    []string{"chat1", "chat2", "chat3"},
			Version: global.Version,
		}), &tele.SendOptions{
			DisableWebPagePreview: true,
			ReplyMarkup: &tele.ReplyMarkup{
				InlineKeyboard: [][]tele.InlineButton{{sp.Template.Button.Start.Settings}},
			},
		})
	}

	return searchbot.OnText()(c)
}
