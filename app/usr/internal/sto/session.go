package sto

import (
	"context"
	"errors"
	"github.com/iyear/searchx/app/usr/internal/key"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/storage/kv"
)

type Session struct {
	kv    storage.KV
	login bool
}

func NewSession(kv storage.KV, login bool) *Session {
	return &Session{kv: kv, login: login}
}

func (s *Session) LoadSession(_ context.Context) ([]byte, error) {
	if s.login {
		return nil, nil
	}

	b, err := s.kv.Get(context.TODO(), key.Session())
	if err != nil {
		if errors.Is(err, kv.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return []byte(b), nil
}

func (s *Session) StoreSession(_ context.Context, data []byte) error {
	return s.kv.Set(context.TODO(), key.Session(), string(data))
}
