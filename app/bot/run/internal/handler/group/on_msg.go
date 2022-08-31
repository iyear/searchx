package group

import (
	"context"
	"github.com/iyear/searchx/app/bot/run/internal/util"
	"github.com/iyear/searchx/pkg/consts"
	"github.com/iyear/searchx/pkg/keygen"
	"github.com/iyear/searchx/pkg/models"
	"github.com/iyear/searchx/pkg/storage/search"
	"github.com/iyear/searchx/pkg/utils"
	tele "gopkg.in/telebot.v3"
)

func OnText(c tele.Context) error {
	return index(c, c.Message().Text)
}

func OnPhoto(c tele.Context) error {
	return index(c, c.Message().Caption)
}

func OnDocument(c tele.Context) error {
	return index(c, c.Message().Document.FileName+" "+c.Message().Caption)
}

func OnVoice(c tele.Context) error {
	return index(c, c.Message().Voice.Caption)
}

func OnVideo(c tele.Context) error {
	return index(c, c.Message().Video.FileName+" "+c.Message().Caption)
}

func OnAudio(c tele.Context) error {
	return index(c, c.Message().Audio.FileName+" "+c.Message().Caption)
}

func OnAnimation(c tele.Context) error {
	return index(c, c.Message().Animation.FileName+" "+c.Message().Caption)
}

func index(c tele.Context, text string) error {
	sp := util.GetScope(c)
	msg := c.Message()

	date := msg.LastEdit
	if date == 0 {
		date = msg.Unixtime
	}

	if msg.Chat == nil {
		sp.Log.Debugw("chat is nil", "msg", msg.ID, "sender", msg.Sender.ID)
		return nil
	}

	m := &models.SearchMsg{
		ID:         msg.ID,
		Chat:       (-msg.Chat.ID) - 1e12,
		ChatType:   consts.ChatGroup,
		ChatName:   msg.Chat.Title,
		Text:       text,
		Sender:     msg.Sender.ID,
		SenderName: utils.Telegram.GetName(msg.Sender.FirstName, msg.Sender.LastName, msg.Sender.Username),
		Date:       date,
	}

	data, err := m.Encode()
	if err != nil {
		return err
	}

	return util.GetScope(c).Storage.Search.Index(context.TODO(), []*search.Item{{
		ID:   keygen.SearchMsgID(msg.Chat.ID, msg.ID),
		Data: data,
	}})
}
