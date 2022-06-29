package group

import (
	"github.com/iyear/searchx/app/bot/internal/key"
	"github.com/iyear/searchx/app/bot/internal/model"
	"github.com/iyear/searchx/app/bot/internal/util"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"time"
)

func Index(c tele.Context) error {
	m := c.Message()

	sp := util.GetScope(c)

	return sp.Storage.Search.Index(key.SearchMsgID(m.Chat.ID, m.ID), &model.SearchMsg{
		ID:     strconv.Itoa(m.ID),
		Chat:   strconv.FormatInt(m.Chat.ID, 10),
		Text:   m.Text,
		Sender: strconv.FormatInt(m.Sender.ID, 10),
		Date:   strconv.FormatInt(time.Now().Unix(), 10),
	})
}

func OnUserJoined(c tele.Context) error {
	//u := c.ChatMember().NewChatMember.User
	return nil
}
