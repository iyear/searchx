package models

type SearchMsg struct {
	ID     string `json:"id"`     // message id
	Chat   string `json:"chat"`   // chat id
	Text   string `json:"text"`   // text content
	Sender string `json:"sender"` // sender id
	Date   string `json:"date"`   // unix timestamp
}
