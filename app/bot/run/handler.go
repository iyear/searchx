package run

import (
	"github.com/iyear/searchx/app/bot/run/internal/config"
	"github.com/iyear/searchx/app/bot/run/internal/handler"
	"github.com/iyear/searchx/app/bot/run/internal/handler/channel"
	group2 "github.com/iyear/searchx/app/bot/run/internal/handler/group"
	private2 "github.com/iyear/searchx/app/bot/run/internal/handler/private"
	"github.com/iyear/searchx/app/bot/run/internal/i18n"
	"github.com/iyear/searchx/app/bot/run/internal/middleware"
	tele "gopkg.in/telebot.v3"
)

func makeHandlers(bot *tele.Bot, button *i18n.TemplateButton) {
	g := bot.Group()
	g.Use(middleware.SuperGroup())
	{
		bot.Handle(config.CmdPing, group2.Ping)

		bot.Handle(tele.OnPhoto, group2.OnPhoto)
		bot.Handle(tele.OnVideo, group2.OnVideo)
		bot.Handle(tele.OnVoice, group2.OnVoice)
		bot.Handle(tele.OnDocument, group2.OnDocument)
		bot.Handle(tele.OnAudio, group2.OnAudio)
		bot.Handle(tele.OnAnimation, group2.OnAnimation)
		bot.Handle(tele.OnEdited, group2.OnText)
	}

	p := bot.Group()
	p.Use(middleware.Private())
	{
		p.Handle(config.CmdStart, private2.Start)

		p.Handle(&button.Start.Settings, private2.SettingsBtn)
		p.Handle(&button.Back, private2.Start)
		p.Handle(&button.Search.Next, private2.SearchNext)
		p.Handle(&button.Search.Prev, private2.SearchPrev)
		p.Handle(&button.Settings.Language, private2.SettingsLanguage)
		p.Handle(&button.Settings.LanguagePlain, private2.SettingsSwitchLanguage)
		p.Handle(&button.Search.SwitchOrder, private2.SearchSwitchOrder)
	}

	bot.Handle(tele.OnText, handler.OnText)

	// channel handlers
	bot.Handle(tele.OnChannelPost, channel.Index)
	bot.Handle(tele.OnEditedChannelPost, channel.Index)

	bot.Handle(tele.OnUserJoined, group2.OnUserJoined)
	bot.Handle(tele.OnAddedToGroup, group2.OnAdded)
}
