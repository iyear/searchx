package run

import (
	"fmt"
	"github.com/iyear/searchx/app/bot/run/internal/config"
	"go.uber.org/zap"
	"golang.org/x/net/proxy"
	"net/http"
)

func getClient() *http.Client {
	if !config.C.Bot.Socks5.Enable {
		return http.DefaultClient
	}

	host := config.C.Bot.Socks5.Host
	port := config.C.Bot.Socks5.Port
	user := config.C.Bot.Socks5.User
	password := config.C.Bot.Socks5.Password

	dialer, err := proxy.SOCKS5("tcp",
		fmt.Sprintf("%s:%d", host, port),
		&proxy.Auth{User: user, Password: password},
		proxy.Direct)

	if err != nil {
		zap.S().Fatalw("failed to get dialer",
			"error", err,
			"host", host,
			"port", port,
			"user", user,
			"password", password)
	}
	return &http.Client{Transport: &http.Transport{DialContext: dialer.(proxy.ContextDialer).DialContext}}
}
