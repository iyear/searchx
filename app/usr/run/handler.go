package run

import (
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/app/usr/internal/config"
	"github.com/iyear/searchx/app/usr/run/internal/handler/bot"
	"github.com/iyear/searchx/app/usr/run/internal/handler/usr"
	"github.com/iyear/searchx/app/usr/run/internal/i18n"
	"github.com/iyear/searchx/pkg/botmid"
	"github.com/iyear/searchx/pkg/searchbot"
	tele "gopkg.in/telebot.v3"
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

func handleBot(b *tele.Bot, button *i18n.TemplateButton) {
	p := b.Group()
	p.Use(botmid.Private())
	{
		p.Handle(tele.OnText, searchbot.Search(config.C.Ctrl.Search.PageSize))
		p.Handle("/start", bot.Start)
		p.Handle(&button.Search.Next, searchbot.Search(config.C.Ctrl.Search.PageSize))
		p.Handle(&button.Search.Prev, searchbot.Search(config.C.Ctrl.Search.PageSize))
		p.Handle(&button.Search.SwitchOrder, searchbot.Search(config.C.Ctrl.Search.PageSize))
	}

}
