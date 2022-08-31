package models

import (
	"encoding/json"
)

type SearchMsg struct {
	ID         int    `json:"id,string" mapstructure:"id" index:"id"`                     // message id
	Chat       int64  `json:"chat,string" mapstructure:"chat" index:"chat"`               // chat id
	ChatType   string `json:"chat_type" mapstructure:"chat_type" index:"chat_type"`       // chat type
	ChatName   string `json:"chat_name" mapstructure:"chat_name" index:"chat_name"`       // chat name
	Text       string `json:"text" mapstructure:"text" index:"text"`                      // text content
	Sender     int64  `json:"sender,string" mapstructure:"sender" index:"sender"`         // sender id
	SenderName string `json:"sender_name" mapstructure:"sender_name" index:"sender_name"` // sender name
	Date       int64  `json:"date,string" mapstructure:"date" index:"date"`               // unix timestamp
}

func (m *SearchMsg) Encode() (map[string]string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	mm := make(map[string]string)
	if err = json.Unmarshal(b, &mm); err != nil {
		return nil, err
	}

	return mm, nil
}
