package searchbot

import (
	"fmt"
	"github.com/iyear/searchx/pkg/consts"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func GetScope(c tele.Context) *SearchScope {
	return c.Get(consts.ContextSearch).(*SearchScope)
}

func GetMsgLink(chat int64, msg int) string {
	return fmt.Sprintf("https://t.me/c/%d/%d", chat, msg)
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
