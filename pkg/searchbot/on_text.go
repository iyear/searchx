package searchbot

import (
	"context"
	"fmt"
	"github.com/iyear/searchx/pkg/hashids"
	"github.com/iyear/searchx/pkg/keygen"
	"github.com/iyear/searchx/pkg/models"
	"github.com/mitchellh/mapstructure"
	tele "gopkg.in/telebot.v3"
	"html"
	"time"
)

func OnText() tele.HandlerFunc {
	return func(c tele.Context) error {
		if c.Message().Payload == "" {
			return nil
		}

		ids, err := hashids.Decode64(c.Message().Payload)
		if err != nil {
			return err
		}
		if len(ids) == 0 {
			return fmt.Errorf("no id type in payload")
		}

		switch ids[0] {
		case TypeView:
			return view(c, ids)
		case TypeGoPrivate:
			return goPrivate(c, ids)
		}
		return fmt.Errorf("payload type is not supported")
	}
}

func view(c tele.Context, ids []int64) error {
	sp := getScope(c)

	if len(ids) != 3 {
		return fmt.Errorf("invalid ids format")
	}

	result, err := sp.Storage.Search.Get(context.TODO(), keygen.SearchMsgID(ids[1], int(ids[2])))
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

func goPrivate(c tele.Context, ids []int64) error {
	sp := getScope(c)

	if len(ids) != 3 {
		return fmt.Errorf("invalid ids format")
	}

	return c.EditOrSend(sp.Text.GoPrivate.T(&TSearchGoPrivate{
		PeerID: ids[1],
		MsgID:  GetWebKMessageID(int(ids[2])),
	}), &tele.SendOptions{DisableWebPagePreview: true})
}
