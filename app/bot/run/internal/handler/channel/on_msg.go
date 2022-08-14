package channel

import (
	"github.com/iyear/searchx/app/bot/run/internal/util"
	"github.com/iyear/searchx/pkg/keygen"
	"github.com/iyear/searchx/pkg/models"
	"github.com/iyear/searchx/pkg/storage/search"
	"github.com/iyear/searchx/pkg/utils"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"time"
)

func Index(c tele.Context) error {
	m := c.Message()

	return util.GetScope(c).Storage.Search.Index([]*search.Item{{
		ID: keygen.SearchMsgID(m.Chat.ID, m.ID),
		Data: &models.SearchMsg{
			ID:         strconv.Itoa(m.ID),
			Chat:       strconv.FormatInt(m.Chat.ID, 10),
			ChatName:   m.Chat.Title,
			Text:       c.Text(),
			Sender:     m.SenderChat.Recipient(),
			SenderName: utils.Telegram.GetSenderName(m.SenderChat.FirstName, m.SenderChat.LastName),
			Date:       strconv.FormatInt(time.Now().Unix(), 10),
		}},
	})
}
