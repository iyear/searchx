package key

import (
	"github.com/iyear/searchx/pkg/keygen"
	"strconv"
)

func Language(tid int64) string {
	return keygen.New("lang", strconv.FormatInt(tid, 10))
}

func SearchMsgID(group int64, id int) string {
	return keygen.New(strconv.FormatInt(group, 10), strconv.Itoa(id))
}
