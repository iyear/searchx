package searchbot

import (
	"github.com/iyear/searchx/pkg/consts"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func getScope(c tele.Context) *SearchScope {
	return c.Get(consts.ContextSearch).(*SearchScope)
}

func searchGetData(data string) (string, int, int) {
	v := strings.Split(data, "|")
	pn, _ := strconv.Atoi(v[1])
	order, _ := strconv.Atoi(v[2])
	return v[0], pn, order
}

func searchSetData(keywords string, pn int, order int) string {
	return keywords + "|" + strconv.Itoa(pn) + "|" + strconv.Itoa(order)
}

const MessageIDIncrement = 0x10000
const MessageIDOffset = 0xFFFFFFFF

func GetWebKMessageID(origin int) int64 {
	if origin >= MessageIDOffset {
		return int64(origin)
	}
	return int64(MessageIDOffset + (origin * MessageIDIncrement))
}
