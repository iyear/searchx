package login

import (
	"context"
	"github.com/fatih/color"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/dcs"
	"github.com/iyear/searchx/app/usr/internal/config"
	"github.com/iyear/searchx/app/usr/internal/sto"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/utils"
	"github.com/tcnksm/go-input"
	"go.uber.org/zap"
)

func Start(ctx context.Context, cfg string) error {
	if err := config.Init(cfg); err != nil {
		return err
	}
	color.Blue("Config loaded")

	_, kv, _, err := storage.Init(config.C.Storage)
	if err != nil {
		return err
	}
	color.Blue("Storage initialized")

	dialer, err := utils.ProxyFromURL(config.C.Proxy)
	if err != nil {
		return err
	}

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

	c := telegram.NewClient(config.C.App.ID, config.C.App.Hash, telegram.Options{
		Resolver: dcs.Plain(dcs.PlainOptions{
			Dial: dialer.DialContext,
		}),
		Device:         Device,
		SessionStorage: sto.NewSession(kv, true),
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
