package source

import (
	"context"
	"github.com/gotd/td/telegram/query"
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/pkg/utils"
)

func getBlockedDialogs(ctx context.Context, client *tg.Client) (map[int64]struct{}, error) {
	blocks, err := query.GetBlocked(client).BatchSize(100).Collect(ctx)
	if err != nil {
		return nil, err
	}

	blockids := make(map[int64]struct{})
	for _, b := range blocks {
		blockids[utils.Telegram.GetPeerID(b.Contact.PeerID)] = struct{}{}
	}
	return blockids, nil
}
