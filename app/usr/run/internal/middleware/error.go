package middleware

import (
	"github.com/iyear/searchx/app/usr/run/internal/util"
	tele "gopkg.in/telebot.v3"
)

func OnError() func(err error, ctx tele.Context) {
	return func(err error, ctx tele.Context) {
		if err != nil {
			log := util.GetBotScope(ctx).Log
			if len(ctx.Recipient().Recipient()) > 0 {
				log.Errorw("error", "err", err, "recipient", ctx.Recipient().Recipient())
				return
			}
			log.Errorw("error", "err", err)
		}
	}
}
