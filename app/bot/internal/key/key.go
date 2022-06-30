package key

import (
	"github.com/iyear/searchx/pkg/keygen"
	"strconv"
)

func Language(tid int64) string {
	return keygen.New("lang", strconv.FormatInt(tid, 10))
}
