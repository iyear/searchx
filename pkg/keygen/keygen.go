package keygen

import (
	"bytes"
	"strconv"
	"strings"
	"sync"
)

var keyPool = sync.Pool{
	New: func() interface{} {
		b := &bytes.Buffer{}
		b.Grow(16)
		return b
	},
}

func New(indexes ...string) string {
	buf := keyPool.Get().(*bytes.Buffer)
	buf.WriteString(strings.Join(indexes, ":"))

	t := buf.String()
	buf.Reset()
	keyPool.Put(buf)
	return t
}

func SearchMsgID(chat int64, id int) string {
	return New(strconv.FormatInt(chat, 10), strconv.Itoa(id))
}
