package util

import (
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/pkg/consts"
	"github.com/iyear/searchx/pkg/utils"
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

func GetPeerType(peer tg.PeerClass, e tg.Entities) string {
	id := GetPeerID(peer)

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
