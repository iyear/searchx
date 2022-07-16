package group

import (
	"github.com/iyear/searchx/app/bot/internal/util"
	"github.com/iyear/searchx/pkg/keygen"
	"github.com/iyear/searchx/pkg/models"
	"github.com/iyear/searchx/pkg/storage"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"time"
)

func Index(c tele.Context) error {
	m := c.Message()

	return util.GetScope(c).Storage.Search.Index([]*storage.SearchItem{{
		ID: keygen.SearchMsgID(m.Chat.ID, m.ID),
		Data: &models.SearchMsg{
			ID:     strconv.Itoa(m.ID),
			Chat:   m.Chat.Recipient(),
			Text:   c.Text(),
			Sender: m.Sender.Recipient(),
			Date:   strconv.FormatInt(time.Now().Unix(), 10),
		}},
	})
}

func OnUserJoined(c tele.Context) error {
	//u := c.ChatMember().NewChatMember.User
	return nil
}
