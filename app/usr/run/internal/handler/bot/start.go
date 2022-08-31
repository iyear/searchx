package bot

import (
	"github.com/iyear/searchx/app/usr/run/internal/model"
	"github.com/iyear/searchx/app/usr/run/internal/util"
	"github.com/iyear/searchx/global"
	"github.com/iyear/searchx/pkg/searchbot"
	tele "gopkg.in/telebot.v3"
)

func Start(c tele.Context) error {
	sp := util.GetBotScope(c)

	// 返回首页则重置
	if c.Message().Payload == "" {
		chat := c.Chat()

		return c.EditOrSend(sp.Template.Text.Start.T(&model.TStart{
			ID:       chat.ID,
			Username: chat.Username,
			Version:  global.Version,
		}), &tele.SendOptions{
			DisableWebPagePreview: true,
		})
	}

	return searchbot.OnText()(c)
}
