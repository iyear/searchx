package source

import (
	"context"
	"github.com/gotd/td/telegram/message/peer"
	"github.com/gotd/td/telegram/query"
	"github.com/gotd/td/tg"
)

func getBlockedDialogs(ctx context.Context, client *tg.Client) (map[int64]struct{}, error) {
	blocks, err := query.GetBlocked(client).BatchSize(100).Collect(ctx)
	if err != nil {
		return nil, err
	}

	blockids := make(map[int64]struct{})
	for _, b := range blocks {
		blockids[GetPeerID(b.Contact.PeerID)] = struct{}{}
	}
	return blockids, nil
}

func GetInputPeerID(peer tg.InputPeerClass) int64 {
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

func GetInputPeerName(peer tg.InputPeerClass, e peer.Entities) string {
	id := GetInputPeerID(peer)

	if n, ok := e.Users()[id]; ok {
		return n.FirstName + "" + n.LastName
	}

	if n, ok := e.Channels()[id]; ok {
		return n.Title
	}

	if n, ok := e.Chats()[id]; ok {
		return n.Title
	}

	return ""
}
