package run

import (
	"github.com/iyear/searchx/app/bot/run/internal/config"
	"github.com/iyear/searchx/app/bot/run/internal/handler"
	"github.com/iyear/searchx/app/bot/run/internal/handler/channel"
	"github.com/iyear/searchx/app/bot/run/internal/handler/group"
	"github.com/iyear/searchx/app/bot/run/internal/handler/private"
	"github.com/iyear/searchx/app/bot/run/internal/i18n"
	"github.com/iyear/searchx/app/bot/run/internal/middleware"
	tele "gopkg.in/telebot.v3"
)

func makeHandlers(bot *tele.Bot, button *i18n.TemplateButton) {
	g := bot.Group()
	g.Use(middleware.SuperGroup())
	{
		bot.Handle(config.CmdPing, group.Ping)

		bot.Handle(tele.OnPhoto, group.OnPhoto)
		bot.Handle(tele.OnVideo, group.OnVideo)
		bot.Handle(tele.OnVoice, group.OnVoice)
		bot.Handle(tele.OnDocument, group.OnDocument)
		bot.Handle(tele.OnAudio, group.OnAudio)
		bot.Handle(tele.OnAnimation, group.OnAnimation)
		bot.Handle(tele.OnEdited, group.OnText)
	}

	p := bot.Group()
	p.Use(middleware.Private())
	{
		p.Handle(config.CmdStart, private.Start)

		p.Handle(&button.Start.Settings, private.SettingsBtn)
		p.Handle(&button.Back, private.Start)
		p.Handle(&button.Search.Next, private.SearchNext)
		p.Handle(&button.Search.Prev, private.SearchPrev)
		p.Handle(&button.Settings.Language, private.SettingsLanguage)
		p.Handle(&button.Settings.LanguagePlain, private.SettingsSwitchLanguage)
		p.Handle(&button.Search.SwitchOrder, private.SearchSwitchOrder)
	}

	bot.Handle(tele.OnText, handler.OnText)

	// channel handlers
	bot.Handle(tele.OnChannelPost, channel.Index)
	bot.Handle(tele.OnEditedChannelPost, channel.Index)

	bot.Handle(tele.OnUserJoined, group.OnUserJoined)
	bot.Handle(tele.OnAddedToGroup, group.OnAdded)
}
