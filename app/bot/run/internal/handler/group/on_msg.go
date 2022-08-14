package group

import (
	"github.com/iyear/searchx/app/bot/run/internal/util"
	"github.com/iyear/searchx/pkg/keygen"
	"github.com/iyear/searchx/pkg/models"
	"github.com/iyear/searchx/pkg/storage/search"
	"github.com/iyear/searchx/pkg/utils"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"strings"
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
	msg := c.Message()

	date := msg.LastEdit
	if date == 0 {
		date = msg.Unixtime
	}

	return util.GetScope(c).Storage.Search.Index([]*search.Item{{
		ID: keygen.SearchMsgID(msg.Chat.ID, msg.ID),
		Data: &models.SearchMsg{
			ID:         strconv.Itoa(msg.ID),
			Chat:       msg.Chat.Recipient(),
			ChatName:   msg.Chat.Title,
			Text:       strings.ReplaceAll(text, "\n", " "),
			Sender:     msg.Sender.Recipient(),
			SenderName: utils.Telegram.GetSenderName(msg.Sender.FirstName, msg.Sender.LastName),
			Date:       strconv.FormatInt(date, 10),
		},
	}})
}

func OnUserJoined(c tele.Context) error {
	// u := c.ChatMember().NewChatMember.User
	return nil
}
