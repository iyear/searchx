package run

import (
	"context"
	"github.com/google/martian/log"
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/app/usr/run/internal/handler/usr"
)

func handleUsr(dispatcher *tg.UpdateDispatcher) {
	dispatcher.OnNewMessage(usr.OnNewMessage)
	dispatcher.OnEditMessage(usr.OnEditMessage)
	dispatcher.OnNewScheduledMessage(usr.OnNewScheduledMessage)
	dispatcher.OnNewChannelMessage(usr.OnNewChannelMessage)
	dispatcher.OnEditChannelMessage(usr.OnEditChannelMessage)
	dispatcher.OnChannelTooLong(func(ctx context.Context, e tg.Entities, update *tg.UpdateChannelTooLong) error {
		log.Infof("channel too long. id: %d, pts: %d", update.ChannelID, update.Pts)
		return nil
	})
	// TODO(iyear): handler delete event?
}
