package index

import (
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/pkg/models"
	"github.com/iyear/searchx/pkg/utils"
	"strings"
)

func Message(m *tg.Message, e tg.Entities) (*models.SearchMsg, bool) {
	// get chat info
	chatID := utils.Telegram.GetPeerID(m.PeerID)
	chatType := utils.Telegram.GetPeerType(m.PeerID, e)
	chatName := utils.Telegram.GetPeerName(m.PeerID, e)

	// get from info
	from, ok := m.GetFromID() // judge is from channel
	senderID, senderName := chatID, chatName
	if ok {
		senderID, senderName = utils.Telegram.GetPeerID(from), utils.Telegram.GetPeerName(from, e)
	}

	// get date
	date, ok := m.GetEditDate()
	if !ok {
		date = m.Date
	}

	text := MessageText(m)
	// clean text
	if strings.TrimFunc(text, func(r rune) bool {
		return r == ' ' || r == '\n' || r == '\t' || r == '\r' || r == '\f' || r == '\v'
	}) == "" {
		return nil, false
	}

	return &models.SearchMsg{
		ID:         m.ID,
		Chat:       chatID,
		ChatType:   chatType,
		ChatName:   chatName,
		Text:       text,
		Sender:     senderID,
		SenderName: senderName,
		Date:       int64(date),
	}, true
}

func MessageText(msg *tg.Message) string {
	// TODO(iyear): index service messages?

	media, ok := msg.GetMedia()
	if ok {
		return strings.Join(append(MediaText(media), msg.Message), " ")
	}

	return msg.Message
}
