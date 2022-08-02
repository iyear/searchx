package group

import (
	"github.com/iyear/searchx/app/bot/internal/util"
	"github.com/iyear/searchx/pkg/keygen"
	"github.com/iyear/searchx/pkg/models"
	"github.com/iyear/searchx/pkg/storage/search"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"strings"
	"time"
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
	return util.GetScope(c).Storage.Search.Index([]*search.Item{{
		ID: keygen.SearchMsgID(msg.Chat.ID, msg.ID),
		Data: &models.SearchMsg{
			ID:     strconv.Itoa(msg.ID),
			Chat:   msg.Chat.Recipient(),
			Text:   strings.ReplaceAll(text, "\n", " "),
			Sender: msg.Sender.Recipient(),
			Date:   strconv.FormatInt(time.Now().Unix(), 10),
		},
	}})
}

func OnUserJoined(c tele.Context) error {
	// u := c.ChatMember().NewChatMember.User
	return nil
}
