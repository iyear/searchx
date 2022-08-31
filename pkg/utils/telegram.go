package utils

import (
	"fmt"
	"github.com/gotd/td/telegram/message/peer"
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/pkg/consts"
)

type telegram struct{}

var Telegram = telegram{}

func (t telegram) GetName(first, last, username string) string {
	if name := first + " " + last; name != " " {
		return name
	}
	return username
}

func (t telegram) GetDeepLink(bot string, code string) string {
	return "https://t.me/" + bot + "?start=" + code
}

func (t telegram) GetMsgLink(chat int64, msg int) string {
	return fmt.Sprintf("https://t.me/c/%d/%d", chat, msg)
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
		return t.GetName(n.FirstName, n.LastName, n.Username)
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

func (t telegram) GetInputPeerName(peer tg.InputPeerClass, e peer.Entities) string {
	id := t.GetInputPeerID(peer)

	if n, ok := e.Users()[id]; ok {
		return t.GetName(n.FirstName, n.LastName, n.Username)
	}

	if n, ok := e.Channels()[id]; ok {
		return n.Title
	}

	if n, ok := e.Chats()[id]; ok {
		return n.Title
	}

	return ""
}

func (t telegram) GetInputPeerID(peer tg.InputPeerClass) int64 {
	switch p := peer.(type) {
	case *tg.InputPeerUser:
		return p.UserID
	case *tg.InputPeerChat:
		return p.ChatID
	case *tg.InputPeerChannel:
		return p.ChannelID
	}

	return 0
}
