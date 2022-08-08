package sto

import (
	"encoding/json"
	"errors"
	"github.com/gotd/td/telegram/updates"
	"github.com/iyear/searchx/app/usr/run/internal/key"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/storage/kv"
)

type State struct {
	kv storage.KV
}

func NewState(kv storage.KV) *State {
	return &State{kv: kv}
}

func (s *State) GetState(userID int64) (updates.State, bool, error) {
	state := updates.State{}

	data, err := s.kv.Get(key.State(userID))
	if err != nil {
		if errors.Is(err, kv.ErrNotFound) {
			return state, false, nil
		}
		return state, false, err
	}

	if err = json.Unmarshal([]byte(data), &state); err != nil {
		return state, false, err
	}

	return state, true, nil
}

func (s *State) SetState(userID int64, state updates.State) error {
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}

	if err = s.kv.Set(key.State(userID), string(data)); err != nil {
		return err
	}

	return s.kv.Set(key.StateChannel(userID), "{}")
}

func (s *State) SetPts(userID int64, pts int) error {
	state, found, err := s.GetState(userID)
	if err != nil {
		return err
	}
	if !found {
		return kv.ErrNotFound
	}

	state.Pts = pts

	data, err := json.Marshal(state)
	if err != nil {
		return err
	}

	return s.kv.Set(key.State(userID), string(data))
}

func (s *State) SetQts(userID int64, qts int) error {
	state, found, err := s.GetState(userID)
	if err != nil {
		return err
	}
	if !found {
		return kv.ErrNotFound
	}

	state.Qts = qts

	data, err := json.Marshal(state)
	if err != nil {
		return err
	}

	return s.kv.Set(key.State(userID), string(data))
}

func (s *State) SetDate(userID int64, date int) error {
	state, found, err := s.GetState(userID)
	if err != nil {
		return err
	}
	if !found {
		return kv.ErrNotFound
	}

	state.Date = date

	data, err := json.Marshal(state)
	if err != nil {
		return err
	}

	return s.kv.Set(key.State(userID), string(data))
}

func (s *State) SetSeq(userID int64, seq int) error {
	state, found, err := s.GetState(userID)
	if err != nil {
		return err
	}
	if !found {
		return kv.ErrNotFound
	}

	state.Seq = seq

	data, err := json.Marshal(state)
	if err != nil {
		return err
	}

	return s.kv.Set(key.State(userID), string(data))
}

func (s *State) SetDateSeq(userID int64, date, seq int) error {
	state, found, err := s.GetState(userID)
	if err != nil {
		return err
	}
	if !found {
		return kv.ErrNotFound
	}

	state.Date = date
	state.Seq = seq

	data, err := json.Marshal(state)
	if err != nil {
		return err
	}

	return s.kv.Set(key.State(userID), string(data))
}

func (s *State) GetChannelPts(userID, channelID int64) (int, bool, error) {
	data, err := s.kv.Get(key.StateChannel(userID))
	if err != nil {
		if errors.Is(err, kv.ErrNotFound) {
			return 0, false, nil
		}
		return 0, false, err
	}

	c := make(map[int64]int)
	if err = json.Unmarshal([]byte(data), &c); err != nil {
		return 0, false, err
	}

	pts, ok := c[channelID]
	if !ok {
		return 0, false, nil
	}

	return pts, true, nil
}

func (s *State) SetChannelPts(userID, channelID int64, pts int) error {
	data, err := s.kv.Get(key.StateChannel(userID))
	if err != nil {
		return err
	}

	c := make(map[int64]int)
	if err = json.Unmarshal([]byte(data), &c); err != nil {
		return err
	}

	c[channelID] = pts
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return s.kv.Set(key.StateChannel(userID), string(b))
}

func (s *State) ForEachChannels(userID int64, f func(channelID int64, pts int) error) error {
	data, err := s.kv.Get(key.StateChannel(userID))
	if err != nil {
		return err
	}

	c := make(map[int64]int)
	if err = json.Unmarshal([]byte(data), &c); err != nil {
		return err
	}

	for channelID, pts := range c {
		if err = f(channelID, pts); err != nil {
			return err
		}
	}

	return nil
}
