package group

import tele "gopkg.in/telebot.v3"

func Ping(c tele.Context) error {
	return c.Send("🏓 PONG")
}
