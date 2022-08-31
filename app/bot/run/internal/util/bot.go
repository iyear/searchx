package util

import (
	tele "gopkg.in/telebot.v3"
)

func appendBack(back tele.InlineButton, opts ...interface{}) []interface{} {
	if len(opts) == 0 {
		return []interface{}{&tele.ReplyMarkup{InlineKeyboard: [][]tele.InlineButton{{back}}}}
	}

	for i, opt := range opts {
		switch t := opt.(type) {
		case *tele.SendOptions:
			if t.ReplyMarkup == nil {
				t.ReplyMarkup = &tele.ReplyMarkup{InlineKeyboard: [][]tele.InlineButton{{back}}}
			} else if t.ReplyMarkup.InlineKeyboard == nil {
				t.ReplyMarkup.InlineKeyboard = [][]tele.InlineButton{{back}}
			} else {
				t.ReplyMarkup.InlineKeyboard = append(t.ReplyMarkup.InlineKeyboard, []tele.InlineButton{back})
			}
			opts[i] = t
		case *tele.ReplyMarkup:
			if t.InlineKeyboard == nil {
				t.InlineKeyboard = [][]tele.InlineButton{{back}}
			} else {
				t.InlineKeyboard = append(t.InlineKeyboard, []tele.InlineButton{back})
			}
			opts[i] = t
		}
	}
	return opts
}

func doWithBack(c tele.Context, fn func(what interface{}, opts ...interface{}) error, what interface{}, opts ...interface{}) error {
	return fn(what, appendBack(GetScope(c).Template.Button.Back, opts...)...)
}

func EditOrSendWithBack(c tele.Context, what interface{}, opts ...interface{}) error {
	return doWithBack(c, c.EditOrSend, what, opts...)
}
