package middleware

import tele "gopkg.in/telebot.v3"

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
