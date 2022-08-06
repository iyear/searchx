package session

import (
	"context"
	"github.com/iyear/searchx/app/usr/run/internal/key"
	"github.com/iyear/searchx/pkg/storage"
)

type Session struct {
	id int
	kv storage.KV
}

func (s *Session) LoadSession(_ context.Context) ([]byte, error) {
	b, err := s.kv.Get(key.Session(s.id))
	if err != nil {
		return nil, err
	}
	return []byte(b), nil
}

func (s *Session) StoreSession(_ context.Context, data []byte) error {
	return s.kv.Set(key.Session(s.id), string(data))
}

func New(kv storage.KV, apiID int, force bool) (*Session, error) {
	ss := &Session{kv: kv, id: apiID}
	if _, err := kv.Get(key.Session(apiID)); err == nil && !force {
		return ss, nil
	}

	if err := kv.Set(key.Session(apiID), ""); err != nil {
		return nil, err
	}

	return ss, nil
}
