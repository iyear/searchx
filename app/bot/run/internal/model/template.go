package model

type TStart struct {
	ID       int64
	Username string
	Notice   string
	Chats    []string
	Version  string
}
