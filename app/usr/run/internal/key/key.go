package key

import (
	"github.com/iyear/searchx/pkg/keygen"
	"strconv"
)

func Session(apiID int) string {
	return keygen.New("session", strconv.Itoa(apiID))
}
