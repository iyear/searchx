package i18n

import (
	"github.com/iyear/searchx/pkg/i18n"
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
	Search struct {
		KeywordsTooLong i18n.Text `mapstructure:"keywords_too_long"`
		Results         i18n.Text `mapstructure:"results"`
	} `mapstructure:"search"`
	Start i18n.Text `mapstructure:"start"`
}

type TemplateButton struct {
	Back   tele.InlineButton `mapstructure:"back"`
	Search struct {
		Next tele.InlineButton `mapstructure:"next"`
		Prev tele.InlineButton `mapstructure:"prev"`
	} `mapstructure:"search"`
	Start struct {
		Settings tele.InlineButton `mapstructure:"settings"`
	} `mapstructure:"start"`
	Settings struct {
		Language      tele.InlineButton `mapstructure:"language"`
		LanguagePlain tele.InlineButton `mapstructure:"language_plain"` // set manually
	} `mapstructure:"settings"`
}
