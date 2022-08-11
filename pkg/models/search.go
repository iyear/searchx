package models

type SearchMsg struct {
	ID         string `json:"id" mapstructure:"id"`                   // message id
	Chat       string `json:"chat" mapstructure:"chat"`               // chat id
	ChatName   string `json:"chat_name" mapstructure:"chat_name"`     // chat name
	Text       string `json:"text" mapstructure:"text"`               // text content
	Sender     string `json:"sender" mapstructure:"sender"`           // sender id
	SenderName string `json:"sender_name" mapstructure:"sender_name"` // sender name
	Date       string `json:"date" mapstructure:"date"`               // unix timestamp
}
