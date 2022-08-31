package usr

import (
	"context"
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/app/usr/run/internal/util"
)

func OnChannelTooLong(ctx context.Context, _ tg.Entities, update *tg.UpdateChannelTooLong) error {
	sp := util.GetUsrScope(ctx)
	sp.Log.Warnw("channel too long", "channel", update.ChannelID, "pts", update.Pts)
	return nil
}
