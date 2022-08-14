package private

import (
	"encoding/base64"
	"github.com/iyear/searchx/app/bot/run/internal/config"
	"github.com/iyear/searchx/app/bot/run/internal/model"
	"github.com/iyear/searchx/app/bot/run/internal/util"
	"github.com/iyear/searchx/global"
	"github.com/iyear/searchx/pkg/models"
	"github.com/iyear/searchx/pkg/utils"
	"github.com/mitchellh/mapstructure"
	tele "gopkg.in/telebot.v3"
	"html"
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

	return messageView(c)
}

func messageView(c tele.Context) error {
	sp := util.GetScope(c)

	id, err := base64.URLEncoding.DecodeString(c.Message().Payload)
	if err != nil {
		return err
	}

	result, err := sp.Storage.Search.Get(string(id))
	if err != nil {
		return err
	}

	msg := models.SearchMsg{}
	if err = mapstructure.Decode(result.Fields, &msg); err != nil {
		return err
	}

	return c.EditOrSend(sp.Template.Text.Search.View.T(&model.TSearchView{
		MsgID:      msg.ID,
		ChatID:     msg.Chat,
		ChatName:   msg.ChatName,
		SenderID:   msg.Sender,
		SenderName: msg.SenderName,
		Date:       utils.String.MustGetDate(msg.Date).Format("2006-01-02 15:04:05"),
		Content:    html.UnescapeString(msg.Text),
	}))
}
