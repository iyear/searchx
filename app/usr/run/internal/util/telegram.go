package util

import (
	"github.com/gotd/td/tg"
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
	switch p := peer.(type) {
	case *tg.PeerUser:
		u := e.Users[p.UserID]
		return utils.String.GetSenderName(u.FirstName, u.LastName)
	case *tg.PeerChat:
		return e.Chats[p.ChatID].Title
	case *tg.PeerChannel:
		return e.Channels[p.ChannelID].Title
	}
	return ""
}
