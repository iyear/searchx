package run

import (
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/app/usr/run/internal/handler/usr"
)

func handleUsr(dispatcher *tg.UpdateDispatcher) {
	dispatcher.OnNewMessage(usr.OnNewMessage)
	dispatcher.OnEditMessage(usr.OnEditMessage)
	dispatcher.OnNewScheduledMessage(usr.OnNewScheduledMessage)
	dispatcher.OnNewChannelMessage(usr.OnNewChannelMessage)
	dispatcher.OnEditChannelMessage(usr.OnEditChannelMessage)
	dispatcher.OnChannelTooLong(usr.OnChannelTooLong)
	// TODO(iyear): handler delete event?
}
