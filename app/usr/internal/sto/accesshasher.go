package sto

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/iyear/searchx/app/usr/internal/key"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/storage/kv"
)

type AccessHasher struct {
	kv storage.KV
}

func NewAccessHasher(kv storage.KV) *AccessHasher {
	return &AccessHasher{kv: kv}
}

func (h *AccessHasher) SetChannelAccessHash(userID, channelID, accessHash int64) error {
	data, err := h.kv.Get(context.TODO(), key.ChannelAccessHash(userID))
	if err != nil && !errors.Is(err, kv.ErrNotFound) {
		return err
	}

	if errors.Is(err, kv.ErrNotFound) {
		data = "{}"
	}

	m := make(map[int64]int64)
	if err = json.Unmarshal([]byte(data), &m); err != nil {
		return err
	}

	m[channelID] = accessHash

	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return h.kv.Set(context.TODO(), key.ChannelAccessHash(userID), string(b))
}

func (h *AccessHasher) GetChannelAccessHash(userID, channelID int64) (int64, bool, error) {
	data, err := h.kv.Get(context.TODO(), key.ChannelAccessHash(userID))
	if err != nil {
		if errors.Is(err, kv.ErrNotFound) {
			return 0, false, nil
		}
		return 0, false, err
	}

	m := make(map[int64]int64)
	if err = json.Unmarshal([]byte(data), &m); err != nil {
		return 0, false, err
	}

	hash, found := m[channelID]
	return hash, found, nil
}
