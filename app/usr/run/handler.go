package run

import (
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/app/usr/run/internal/handler/usr"
)

func handleUsr(dispatcher *tg.UpdateDispatcher) {
	dispatcher.OnNewMessage(usr.OnNewMessage)
}
