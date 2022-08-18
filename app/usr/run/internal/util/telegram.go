package util

import (
	"fmt"
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/pkg/utils"
	tele "gopkg.in/telebot.v3"
)

func GetPeerID(peer tg.PeerClass) int64 {
	switch p := peer.(type) {
	case *tg.PeerUser:
		return p.UserID
	case *tg.PeerChat:
		return p.ChatID
	case *tg.PeerChannel:
		return p.ChannelID
	}
	return 0
}

func GetPeerName(peer tg.PeerClass, e tg.Entities) string {
	id := GetPeerID(peer)

	if n, ok := e.Users[id]; ok {
		return utils.Telegram.GetSenderName(n.FirstName, n.LastName)
	}

	if n, ok := e.Channels[id]; ok {
		return n.Title
	}

	if n, ok := e.Chats[id]; ok {
		return n.Title
	}

	return ""
}

func appendBack(back tele.InlineButton, opts ...interface{}) []interface{} {
	if len(opts) == 0 {
		return []interface{}{&tele.ReplyMarkup{InlineKeyboard: [][]tele.InlineButton{{back}}}}
	}

	for i, opt := range opts {
		switch t := opt.(type) {
		case *tele.SendOptions:
			if t.ReplyMarkup == nil {
				t.ReplyMarkup = &tele.ReplyMarkup{InlineKeyboard: [][]tele.InlineButton{{back}}}
			} else if t.ReplyMarkup.InlineKeyboard == nil {
				t.ReplyMarkup.InlineKeyboard = [][]tele.InlineButton{{back}}
			} else {
				t.ReplyMarkup.InlineKeyboard = append(t.ReplyMarkup.InlineKeyboard, []tele.InlineButton{back})
			}
			opts[i] = t
		case *tele.ReplyMarkup:
			if t.InlineKeyboard == nil {
				t.InlineKeyboard = [][]tele.InlineButton{{back}}
			} else {
				t.InlineKeyboard = append(t.InlineKeyboard, []tele.InlineButton{back})
			}
			opts[i] = t
		}
	}
	return opts
}

func doWithBack(c tele.Context, fn func(what interface{}, opts ...interface{}) error, what interface{}, opts ...interface{}) error {
	return fn(what, appendBack(GetBotScope(c).Template.Button.Back, opts...)...)
}

func EditOrSendWithBack(c tele.Context, what interface{}, opts ...interface{}) error {
	return doWithBack(c, c.EditOrSend, what, opts...)
}

func GetMsgLink(chat int64, msg int) string {
	return fmt.Sprintf("https://t.me/c/%d/%d", chat, msg)
}
