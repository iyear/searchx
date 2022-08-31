package utils

import (
	"errors"
	"golang.org/x/net/proxy"
	"net/url"
)

func IF(f bool, a, b interface{}) interface{} {
	if f {
		return a
	}
	return b
}

func ProxyFromURL(u string) (proxy.ContextDialer, error) {
	if u == "" {
		return proxy.Direct, nil
	}

	parse, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	dialer, err := proxy.FromURL(parse, proxy.Direct)
	if err != nil {
		return nil, err
	}

	d, ok := dialer.(proxy.ContextDialer)
	if !ok {
		return nil, errors.New("dialer is not context dialer")
	}

	return d, nil
}
