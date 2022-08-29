package searchbot

import (
	"github.com/iyear/searchx/pkg/i18n"
	"github.com/iyear/searchx/pkg/storage"
	tele "gopkg.in/telebot.v3"
)

type SearchScope struct {
	Text    *SearchContextTextTemplate
	Button  *SearchContextButtonTemplate
	Storage *storage.Storage
}

type SearchContextTextTemplate struct {
	KeywordsTooLong i18n.Text `mapstructure:"keywords_too_long"`
	View            i18n.Text `mapstructure:"view"`
	Results         i18n.Text `mapstructure:"results"`
	GoPrivate       i18n.Text `mapstructure:"go_private"`
}

type SearchContextButtonTemplate struct {
	Next        tele.InlineButton `mapstructure:"next"`
	Prev        tele.InlineButton `mapstructure:"prev"`
	SwitchOrder tele.InlineButton `mapstructure:"switch_order"`
}
