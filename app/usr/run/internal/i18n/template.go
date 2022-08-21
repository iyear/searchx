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
	Search searchbot.SearchContextTextTemplate `mapstructure:"search"`
	Start  i18n.Text                           `mapstructure:"start"`
}

type TemplateButton struct {
	Back   tele.InlineButton                     `mapstructure:"back"`
	Search searchbot.SearchContextButtonTemplate `mapstructure:"search"`
}
