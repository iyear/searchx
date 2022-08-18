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
	Search struct {
		KeywordsTooLong i18n.Text `mapstructure:"keywords_too_long"`
		Results         i18n.Text `mapstructure:"results"`
		View            i18n.Text `mapstructure:"view"`
	} `mapstructure:"search"`
	Start i18n.Text `mapstructure:"start"`
}

type TemplateButton struct {
	Back   tele.InlineButton `mapstructure:"back"`
	Search struct {
		Next        tele.InlineButton `mapstructure:"next"`
		Prev        tele.InlineButton `mapstructure:"prev"`
		SwitchOrder tele.InlineButton `mapstructure:"switch_order"`
	} `mapstructure:"search"`
}
