package usr

import (
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/app/usr/run/internal/model"
	"github.com/iyear/searchx/app/usr/run/internal/util"
	"github.com/iyear/searchx/pkg/keygen"
	"github.com/iyear/searchx/pkg/models"
	"github.com/iyear/searchx/pkg/storage/search"
	"strconv"
	"strings"
)

func index(sp *model.UsrScope, chatID int64, chatName string, msgID int, senderID int64, senderName string, text string, date int) error {
	// clean text
	if strings.TrimFunc(text, func(r rune) bool {
		return r == ' ' || r == '\n' || r == '\t' || r == '\r' || r == '\f' || r == '\v'
	}) == "" {
		return nil
	}

	sp.Log.Debugw("new message", "chatID", chatID, "chatName", chatName, "msgID", msgID, "senderID", senderID, "senderName", senderName, "text", text, "date", date)

	return sp.Storage.Search.Index([]*search.Item{{
		ID: keygen.SearchMsgID(chatID, msgID),
		Data: &models.SearchMsg{
			ID:         strconv.Itoa(msgID),
			Chat:       strconv.FormatInt(chatID, 10),
			ChatName:   chatName,
			Text:       strings.ReplaceAll(text, "\n", " "),
			Sender:     strconv.FormatInt(senderID, 10),
			SenderName: senderName,
			Date:       strconv.Itoa(date),
		},
	}})
}

func indexMessage(sp *model.UsrScope, e tg.Entities, msg tg.MessageClass) error {
	m, ok := msg.(*tg.Message)
	if !ok {
		return nil
	}

	// get chat info
	chatID := util.GetPeerID(m.PeerID)
	chatName := util.GetPeerName(m.PeerID, e)

	// get from info
	from, ok := m.GetFromID() // judge is from channel
	senderID, senderName := chatID, chatName
	if ok {
		senderID, senderName = util.GetPeerID(from), util.GetPeerName(from, e)
	}

	// get date
	date, ok := m.GetEditDate()
	if !ok {
		date = m.Date
	}

	return index(sp, chatID, chatName, m.ID, senderID, senderName, messageText(m), date)
}
