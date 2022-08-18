package bot

import (
	"encoding/base64"
	"github.com/iyear/searchx/app/usr/run/internal/model"
	"github.com/iyear/searchx/app/usr/run/internal/util"
	"github.com/iyear/searchx/global"
	"github.com/iyear/searchx/pkg/models"
	"github.com/mitchellh/mapstructure"
	tele "gopkg.in/telebot.v3"
	"html"
	"time"
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

	return messageView(c)
}

func messageView(c tele.Context) error {
	sp := util.GetBotScope(c)

	id, err := base64.URLEncoding.DecodeString(c.Message().Payload)
	if err != nil {
		return err
	}

	result, err := sp.Storage.Search.Get(string(id))
	if err != nil {
		return err
	}

	msg := models.SearchMsg{}
	if err = mapstructure.WeakDecode(result.Fields, &msg); err != nil {
		return err
	}

	return c.EditOrSend(sp.Template.Text.Search.View.T(&model.TSearchView{
		MsgID:      msg.ID,
		ChatID:     msg.Chat,
		ChatName:   msg.ChatName,
		SenderID:   msg.Sender,
		SenderName: msg.SenderName,
		Date:       time.Unix(msg.Date, 0).Format("2006-01-02 15:04:05"),
		Content:    html.UnescapeString(msg.Text),
	}))
}
