package channel

import (
	"context"
	"github.com/iyear/searchx/app/bot/run/internal/util"
	"github.com/iyear/searchx/pkg/consts"
	"github.com/iyear/searchx/pkg/keygen"
	"github.com/iyear/searchx/pkg/models"
	"github.com/iyear/searchx/pkg/storage/search"
	"github.com/iyear/searchx/pkg/utils"
	tele "gopkg.in/telebot.v3"
	"time"
)

func Index(c tele.Context) error {
	m := c.Message()

	msg := &models.SearchMsg{
		ID:         m.ID,
		Chat:       (-m.Chat.ID) - 1e12,
		ChatType:   consts.ChatChannel,
		ChatName:   m.Chat.Title,
		Text:       c.Text(),
		Sender:     m.SenderChat.ID,
		SenderName: utils.Telegram.GetName(m.SenderChat.FirstName, m.SenderChat.LastName, m.SenderChat.Username),
		Date:       time.Now().Unix(),
	}
	data, err := msg.Encode()
	if err != nil {
		return err
	}

	return util.GetScope(c).Storage.Search.Index(context.TODO(), []*search.Item{{
		ID:   keygen.SearchMsgID(m.Chat.ID, m.ID),
		Data: data,
	}})
}
