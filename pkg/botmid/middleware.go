package botmid

import tele "gopkg.in/telebot.v3"

func AutoResponder() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if c.Callback() != nil {
				defer func(c tele.Context) {
					_ = c.Respond()
				}(c)
			}
			return next(c) // continue execution chain
		}
	}
}

func Private() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if c.Chat() != nil && c.Chat().Type != tele.ChatPrivate {
				return nil
			}
			return next(c)
		}
	}
}

func SuperGroup() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if c.Chat() != nil && c.Chat().Type != tele.ChatSuperGroup {
				return nil
			}
			return next(c)
		}
	}
}
