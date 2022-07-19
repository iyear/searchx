package bot

import (
	"github.com/iyear/searchx/app/bot/internal/config"
	"github.com/iyear/searchx/app/bot/internal/handler"
	"github.com/iyear/searchx/app/bot/internal/handler/channel"
	"github.com/iyear/searchx/app/bot/internal/handler/group"
	"github.com/iyear/searchx/app/bot/internal/handler/private"
	"github.com/iyear/searchx/app/bot/internal/i18n"
	"github.com/iyear/searchx/app/bot/internal/middleware"
	tele "gopkg.in/telebot.v3"
)

func makeHandlers(bot *tele.Bot, button *i18n.TemplateButton) {
	g := bot.Group()
	g.Use(middleware.SuperGroup())
	{
		bot.Handle(tele.OnPhoto, group.Index)
		bot.Handle(tele.OnVideo, group.Index)
		bot.Handle(tele.OnDocument, group.Index)
		bot.Handle(tele.OnAudio, group.Index)
		bot.Handle(tele.OnAnimation, group.Index)
		bot.Handle(tele.OnEdited, group.Index)
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
