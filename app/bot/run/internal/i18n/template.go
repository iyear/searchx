package i18n

import (
	"github.com/iyear/searchx/pkg/i18n"
	"github.com/iyear/searchx/pkg/searchbot"
	tele "gopkg.in/telebot.v3"
)

type Template struct {
	Text   *TemplateText   `mapstructure:"text"`
	Button *TemplateButton `mapstructure:"button"`
}

type TemplateText struct {
	Settings struct {
		Desc     i18n.Text `mapstructure:"desc"`
		Language i18n.Text `mapstructure:"language"`
	} `mapstructure:"settings"`
	Search       searchbot.SearchContextTextTemplate `mapstructure:"search"`
	AddedToGroup struct {
		Fail    i18n.Text `mapstructure:"fail"`
		Success i18n.Text `mapstructure:"success"`
	} `mapstructure:"added_to_group"`
	Start i18n.Text `mapstructure:"start"`
}

type TemplateButton struct {
	Back   tele.InlineButton                     `mapstructure:"back"`
	Search searchbot.SearchContextButtonTemplate `mapstructure:"search"`
	Start  struct {
		Settings tele.InlineButton `mapstructure:"settings"`
	} `mapstructure:"start"`
	Settings struct {
		Language      tele.InlineButton `mapstructure:"language"`
		LanguagePlain tele.InlineButton `mapstructure:"language_plain"` // set manually
	} `mapstructure:"settings"`
}
