package utils

import (
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/pkg/consts"
)

type telegram struct{}

var Telegram = telegram{}

func (t telegram) GetSenderName(first, last string) string {
	if last == "" {
		return first
	}
	if first == "" {
		return last
	}

	return first + " " + last
}

func (t telegram) GetDeepLink(bot string, code string) string {
	return "https://t.me/" + bot + "?start=" + code
}

func (t telegram) GetPeerID(peer tg.PeerClass) int64 {
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

func (t telegram) GetPeerName(peer tg.PeerClass, e tg.Entities) string {
	id := t.GetPeerID(peer)

	if n, ok := e.Users[id]; ok {
		return t.GetSenderName(n.FirstName, n.LastName)
	}

	if n, ok := e.Channels[id]; ok {
		return n.Title
	}

	if n, ok := e.Chats[id]; ok {
		return n.Title
	}

	return ""
}

func (t telegram) GetPeerType(peer tg.PeerClass, e tg.Entities) string {
	id := t.GetPeerID(peer)

	if _, ok := e.Users[id]; ok {
		return consts.ChatPrivate
	}

	if n, ok := e.Channels[id]; ok {
		if n.Megagroup || n.Gigagroup {
			return consts.ChatGroup
		}
		return consts.ChatChannel
	}

	if _, ok := e.Chats[id]; ok {
		return consts.ChatGroup
	}

	return consts.ChatUnknown
}
