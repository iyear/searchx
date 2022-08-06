package login

import (
	"context"
	"github.com/fatih/color"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/dcs"
	"github.com/iyear/searchx/app/usr/run/internal/config"
	"github.com/iyear/searchx/app/usr/run/internal/session"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/utils"
	"github.com/tcnksm/go-input"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

func Start() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	color.Blue("Login...")

	phone, err := input.DefaultUI().Ask("Enter your phone number:", &input.Options{
		Default:  "+86 12345678900",
		Loop:     true,
		Required: true,
	})
	if err != nil {
		return err
	}
	color.Blue("Send code...")

	dialer, err := utils.ProxyFromURL(config.C.Proxy)
	if err != nil {
		return err
	}

	kv, err := storage.NewKV(config.C.Storage.KV.Driver, config.C.Storage.KV.Options)
	if err != nil {
		return err
	}

	sess, err := session.New(kv, config.C.Account.ID, true)
	if err != nil {
		return err
	}

	c := telegram.NewClient(config.C.Account.ID, config.C.Account.Hash, telegram.Options{
		Resolver: dcs.Plain(dcs.PlainOptions{
			Dial: dialer.DialContext,
		}),
		SessionStorage: sess,
		Logger:         zap.NewNop(),
	})

	return c.Run(ctx, func(ctx context.Context) error {
		flow := auth.NewFlow(termAuth{phone: phone}, auth.SendCodeOptions{})
		if err := c.Auth().IfNecessary(ctx, flow); err != nil {
			return err
		}

		user, err := c.Self(ctx)
		if err != nil {
			return err
		}

		color.Blue("Login successfully! ID: %d, Username: %s", user.ID, user.Username)

		return nil
	})
}
