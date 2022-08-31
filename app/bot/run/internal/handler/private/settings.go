package private

import (
	"context"
	"fmt"
	iso6391 "github.com/iyear/iso-639-1"
	"github.com/iyear/searchx/app/bot/run/internal/i18n"
	"github.com/iyear/searchx/app/bot/run/internal/key"
	"github.com/iyear/searchx/app/bot/run/internal/middleware"
	"github.com/iyear/searchx/app/bot/run/internal/util"
	"github.com/iyear/searchx/pkg/utils"
	tele "gopkg.in/telebot.v3"
)

func SettingsBtn(c tele.Context) error {
	sp := util.GetScope(c)

	return util.EditOrSendWithBack(c, sp.Template.Text.Settings.Desc.T(nil),
		&tele.ReplyMarkup{InlineKeyboard: [][]tele.InlineButton{{sp.Template.Button.Settings.Language}}})
}

func SettingsLanguage(c tele.Context) error {
	sp := util.GetScope(c)

	langBtns := make([][]tele.InlineButton, 0)

	nowLang := util.GetUserLanguage(sp.Storage, c.Chat().ID)

	for _, code := range i18n.Languages {
		langBtn := sp.Template.Button.Settings.LanguagePlain
		lang := iso6391.FromCode(code)

		langBtn.Text = fmt.Sprintf("%s%s (%s)", utils.IF(nowLang == code, "✅ ", ""), lang.Name, lang.NativeName)
		langBtn.Data = code
		langBtns = append(langBtns, []tele.InlineButton{langBtn})
	}

	return util.EditOrSendWithBack(c, sp.Template.Text.Settings.Language.T(iso6391.FromCode(nowLang).NativeName),
		&tele.ReplyMarkup{InlineKeyboard: langBtns})
}

func SettingsSwitchLanguage(c tele.Context) error {
	sp := util.GetScope(c)

	// 相同则不做任何事
	if util.GetUserLanguage(sp.Storage, c.Chat().ID) == c.Data() {
		return nil
	}

	if err := sp.Storage.KV.Set(context.TODO(), key.Language(c.Chat().ID), c.Data()); err != nil {
		return err
	}

	sp.Storage.Cache.Set(context.Background(), key.Language(c.Chat().ID), c.Data())

	// 手动走一遍中间件刷新当前页面的语言
	return middleware.SetScope(sp)(SettingsLanguage)(c)
}
