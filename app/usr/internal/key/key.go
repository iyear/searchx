package key

import (
	"github.com/iyear/searchx/pkg/keygen"
	"strconv"
)

func Session() string {
	return keygen.New("session")
}

func State(userID int64) string {
	return keygen.New("state", strconv.FormatInt(userID, 10))
}

func StateChannel(userID int64) string {
	return keygen.New("chan", strconv.FormatInt(userID, 10))
}

func ChannelAccessHash(userID int64) string {
	return keygen.New("chash", strconv.FormatInt(userID, 10))
}
