package searchbot

import (
	"context"
	"encoding/base64"
	"github.com/iyear/searchx/pkg/models"
	"github.com/mitchellh/mapstructure"
	tele "gopkg.in/telebot.v3"
	"html"
	"time"
)

func View() tele.HandlerFunc {
	return func(c tele.Context) error {
		sp := getScope(c)

		id, err := base64.URLEncoding.DecodeString(c.Message().Payload)
		if err != nil {
			return err
		}

		result, err := sp.Storage.Search.Get(context.TODO(), string(id))
		if err != nil {
			return err
		}

		msg := models.SearchMsg{}
		if err = mapstructure.WeakDecode(result.Fields, &msg); err != nil {
			return err
		}

		return c.EditOrSend(sp.Text.View.T(&TSearchView{
			MsgID:      msg.ID,
			ChatID:     msg.Chat,
			ChatType:   msg.ChatType,
			ChatName:   msg.ChatName,
			SenderID:   msg.Sender,
			SenderName: msg.SenderName,
			Date:       time.Unix(msg.Date, 0).Format("2006-01-02 15:04:05"),
			Content:    html.UnescapeString(msg.Text),
		}))
	}
}
