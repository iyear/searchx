package private

import (
	"github.com/iyear/searchx/app/bot/internal/config"
	"github.com/iyear/searchx/app/bot/internal/model"
	"github.com/iyear/searchx/app/bot/internal/util"
	"github.com/iyear/searchx/pkg/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/thinkeridea/go-extend/exunicode/exutf8"
	tele "gopkg.in/telebot.v3"
	"strconv"
	"strings"
	"time"
)

func Search(c tele.Context) error {
	var pn int
	var keyword string
	var btns [][]tele.InlineButton
	ps := config.C.Ctrl.Search.PageSize

	sp := util.GetScope(c)

	start := time.Now()

	keyword = strings.ReplaceAll(c.Message().Text, "|", "")
	if c.Callback() == nil { // 初始
		pn = 0
		// 由于c.Data长度限制，关键词长度也限制
		if len(keyword) > 55 {
			return util.EditOrSendWithBack(c, sp.Template.Text.Search.KeywordsTooLong.T(nil))
		}
	} else {
		keyword, pn = searchGetData(c.Data())
	}

	nextBtn := sp.Template.Button.Search.Next
	nextBtn.Data = searchSetData(keyword, pn+1)

	prevBtn := sp.Template.Button.Search.Prev
	prevBtn.Data = searchSetData(keyword, pn-1)

	// 每次多查一个判断 total%ps==0 的情况
	searchResults := sp.Storage.Search.Search(keyword, pn*ps, ps+1)
	if pn == 0 {
		if len(searchResults) > ps {
			btns = append(btns, []tele.InlineButton{nextBtn})
		}
	} else if pn > 0 {
		if len(searchResults) > ps {
			btns = append(btns, []tele.InlineButton{prevBtn, nextBtn})
		} else {
			btns = append(btns, []tele.InlineButton{prevBtn})
		}
	}

	// 如果还有下页,len>ps,则最后一个不要,即只取到ps个
	// 如果没有下页,len<=ps,则都要,即只取到len个
	num := utils.MinInt(len(searchResults), ps)
	results := make([]*model.TSearchResult, 0, num)
	msg := model.SearchMsg{}
	for i := 0; i < num; i++ {
		result := searchResults[i]
		if err := mapstructure.Decode(result.Fields, &msg); err != nil {
			return err
		}

		maxHighlight := 3
		count := 0
		contents := []string{""} // 在两边也添加省略号
		for _, loc := range result.Location["text"] {
			contents = append(contents, util.Highlight(msg.Text, int(loc.Start), int(loc.End), 5, 5, "*"))
			count++
			if count == maxHighlight {
				break
			}
		}
		if count == 0 {
			contents = append(contents, exutf8.RuneSubString(msg.Text, 0, 10))
		}

		results = append(results, &model.TSearchResult{
			Seq:        pn*ps + i + 1,
			Sender:     msg.Sender,
			SenderLink: "tg://user?id=" + msg.Sender,
			Date:       util.MustGetDate(msg.Date).Format("2006.01.02"),
			Content:    strings.Join(append(contents, ""), "..."),
			Link:       util.GetMsgLink(msg.Chat, msg.ID),
		})
	}

	return util.EditOrSendWithBack(c, sp.Template.Text.Search.Results.T(&model.TSearchResults{
		Results: results,
		Keyword: keyword,
		Took:    time.Since(start).Milliseconds(),
	}), &tele.SendOptions{
		ReplyMarkup:           &tele.ReplyMarkup{InlineKeyboard: btns},
		DisableWebPagePreview: true,
	})

}

func SearchNext(c tele.Context) error {
	return Search(c)
}

func SearchPrev(c tele.Context) error {
	return Search(c)
}

func searchGetData(data string) (string, int) {
	v := strings.Split(data, "|")
	pn, _ := strconv.Atoi(v[1])
	return v[0], pn
}

func searchSetData(keywords string, pn int) string {
	return keywords + "|" + strconv.Itoa(pn)
}
