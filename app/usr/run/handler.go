package run

import (
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/app/usr/run/internal/handler/bot"
	"github.com/iyear/searchx/app/usr/run/internal/handler/usr"
	"github.com/iyear/searchx/app/usr/run/internal/i18n"
	"github.com/iyear/searchx/app/usr/run/internal/middleware"
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
	p.Use(middleware.Private())
	{
		p.Handle(tele.OnText, bot.Search)
		p.Handle("/start", bot.Start)
		p.Handle(&button.Search.Next, bot.SearchNext)
		p.Handle(&button.Search.Prev, bot.SearchPrev)
		p.Handle(&button.Search.SwitchOrder, bot.SearchSwitchOrder)
	}

}
