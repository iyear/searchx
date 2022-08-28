package run

import (
	"github.com/iyear/searchx/app/bot/internal/config"
	"github.com/iyear/searchx/app/bot/run/internal/conf"
	"github.com/iyear/searchx/app/bot/run/internal/handler"
	"github.com/iyear/searchx/app/bot/run/internal/handler/channel"
	"github.com/iyear/searchx/app/bot/run/internal/handler/group"
	"github.com/iyear/searchx/app/bot/run/internal/handler/private"
	"github.com/iyear/searchx/app/bot/run/internal/i18n"
	"github.com/iyear/searchx/pkg/botmid"
	"github.com/iyear/searchx/pkg/searchbot"
	tele "gopkg.in/telebot.v3"
)

func makeHandlers(bot *tele.Bot, button *i18n.TemplateButton) {
	g := bot.Group()
	g.Use(botmid.SuperGroup())
	{
		bot.Handle(conf.CmdPing, group.Ping)

		bot.Handle(tele.OnPhoto, group.OnPhoto)
		bot.Handle(tele.OnVideo, group.OnVideo)
		bot.Handle(tele.OnVoice, group.OnVoice)
		bot.Handle(tele.OnDocument, group.OnDocument)
		bot.Handle(tele.OnAudio, group.OnAudio)
		bot.Handle(tele.OnAnimation, group.OnAnimation)
		bot.Handle(tele.OnEdited, group.OnText)
	}

	p := bot.Group()
	p.Use(botmid.Private())
	{
		p.Handle(conf.CmdStart, private.Start)

		p.Handle(&button.Start.Settings, private.SettingsBtn)
		p.Handle(&button.Back, private.Start)
		p.Handle(&button.Search.Next, searchbot.SearchNext(config.C.Ctrl.Search.PageSize))
		p.Handle(&button.Search.Prev, searchbot.SearchPrev(config.C.Ctrl.Search.PageSize))
		p.Handle(&button.Settings.Language, private.SettingsLanguage)
		p.Handle(&button.Settings.LanguagePlain, private.SettingsSwitchLanguage)
		p.Handle(&button.Search.SwitchOrder, searchbot.SearchSwitchOrder(config.C.Ctrl.Search.PageSize))
	}

	bot.Handle(tele.OnText, handler.OnText)

	// channel handlers
	bot.Handle(tele.OnChannelPost, channel.Index)
	bot.Handle(tele.OnEditedChannelPost, channel.Index)

	bot.Handle(tele.OnAddedToGroup, group.OnAdded)
}
