package login

import (
	"context"
	"errors"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"github.com/iyear/searchx/global"
	"github.com/tcnksm/go-input"
	"runtime"
	"strings"
)

var (
	Device = telegram.DeviceConfig{
		DeviceModel:   "SearchX",
		SystemVersion: runtime.GOOS,
		AppVersion:    global.Version,
	}
)

// noSignUp can be embedded to prevent signing up.
type noSignUp struct{}

func (c noSignUp) SignUp(_ context.Context) (auth.UserInfo, error) {
	return auth.UserInfo{}, errors.New("searchx don't support sign up Telegram account")
}

func (c noSignUp) AcceptTermsOfService(_ context.Context, tos tg.HelpTermsOfService) error {
	return &auth.SignUpRequired{TermsOfService: tos}
}

// termAuth implements authentication via terminal.
type termAuth struct {
	noSignUp
	phone string
}

func (a termAuth) Phone(_ context.Context) (string, error) {
	return a.phone, nil
}

func (a termAuth) Password(_ context.Context) (string, error) {
	pwd, err := input.DefaultUI().Ask("Enter 2FA Password:", &input.Options{
		Required: true,
		Loop:     true,
	})
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(pwd), nil
}

func (a termAuth) Code(_ context.Context, _ *tg.AuthSentCode) (string, error) {
	code, err := input.DefaultUI().Ask("Enter Code:", &input.Options{
		Required: true,
		Loop:     true,
	})
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(code), nil
}
