package searchbot

import (
	"context"
	"github.com/iyear/searchx/pkg/consts"
	"github.com/iyear/searchx/pkg/hashids"
	"github.com/iyear/searchx/pkg/models"
	"github.com/iyear/searchx/pkg/storage/search"
	"github.com/iyear/searchx/pkg/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/thinkeridea/go-extend/exunicode/exutf8"
	tele "gopkg.in/telebot.v3"
	"html"
	"strconv"
	"strings"
	"time"
)

func Search(pageSize int) tele.HandlerFunc {
	return func(c tele.Context) error {
		var btns [][]tele.InlineButton
		pn, order, keyword, ps := 0, 0, "", pageSize

		sp := getScope(c)

		start := time.Now()

		keyword = strings.ReplaceAll(c.Message().Text, "|", "")
		if c.Callback() == nil { // 初始
			// 由于c.Data长度限制，关键词长度也限制
			if len(keyword) > 55 {
				return c.EditOrSend(sp.Text.KeywordsTooLong.T(nil))
			}
		} else {
			keyword, pn, order = searchGetData(c.Data())
		}

		nextBtn := sp.Button.Next
		nextBtn.Data = searchSetData(keyword, pn+1, order)

		orderBtn := sp.Button.SwitchOrder
		orderBtn.Text = SearchOrders[order].Text
		nextOrder := (order + 1) % len(SearchOrders)
		orderBtn.Data = searchSetData(keyword, pn, nextOrder)

		prevBtn := sp.Button.Prev
		prevBtn.Data = searchSetData(keyword, pn-1, order)

		// 每次多查一个判断 total%ps==0 的情况
		searchResults := sp.Storage.Search.Search(context.TODO(), keyword, search.Options{
			From:   pn * ps,
			Size:   ps + 1,
			SortBy: SearchOrders[order].SortBy,
		})
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

		btns = append(btns, []tele.InlineButton{orderBtn})

		// 如果还有下页,len>ps,则最后一个不要,即只取到ps个
		// 如果没有下页,len<=ps,则都要,即只取到len个
		num := utils.Math.MinInt(len(searchResults), ps)
		results := make([]*TSearchResult, 0, num)
		msg := models.SearchMsg{}
		for i := 0; i < num; i++ {
			result := searchResults[i]
			if err := mapstructure.WeakDecode(result.Fields, &msg); err != nil {
				return err
			}

			maxHighlight := 3
			count := 0
			contents := make([]string, 0)

			for _, loc := range result.Location["text"] {
				contents = append(contents, utils.String.Highlight(msg.Text, int(loc.Start), int(loc.End),
					HighlightSpace, HighlightSpace, "\a", "\b"))
				count++
				if count == maxHighlight {
					break
				}
			}
			if count == 0 {
				contents = append(contents, exutf8.RuneSubString(msg.Text, 0, 10))
			}

			sender := utils.String.RuneSubString(msg.SenderName, SenderNameMax)
			if sender == "" {
				sender = strconv.FormatInt(msg.Sender, 10)
			}

			// set link
			golink := utils.Telegram.GetMsgLink(msg.Chat, msg.ID)
			if msg.ChatType == consts.ChatPrivate {
				deep, err := hashids.Encode64(TypeGoPrivate, msg.Chat, int64(msg.ID))
				if err != nil {
					return err
				}
				golink = utils.Telegram.GetDeepLink(c.Bot().Me.Username, deep)
			}

			deep, err := hashids.Encode64(TypeView, msg.Chat, int64(msg.ID))
			if err != nil {
				return err
			}
			viewlink := utils.Telegram.GetDeepLink(c.Bot().Me.Username, deep)

			results = append(results, &TSearchResult{
				Seq:        pn*ps + i + 1,
				ViewLink:   viewlink,
				SenderName: html.EscapeString(strings.TrimSpace(sender)),
				SenderLink: "tg://user?id=" + strconv.FormatInt(msg.Sender, 10),
				ChatName:   html.EscapeString(utils.String.RuneSubString(msg.ChatName, ChatNameMax)),
				Date:       time.Unix(msg.Date, 0).Format("2006.01.02"),
				Content:    strings.ReplaceAll(html.EscapeString(strings.Join(contents, "...")), "\n", " "),
				GoLink:     golink,
			})
		}

		text := strings.NewReplacer("\a", "<b>", "\b", "</b>").Replace(sp.Text.Results.T(&TSearchResults{
			Results: results,
			Keyword: keyword,
			Took:    time.Since(start).Milliseconds(),
		}))

		return c.EditOrSend(text, &tele.SendOptions{
			ReplyMarkup:           &tele.ReplyMarkup{InlineKeyboard: btns},
			DisableWebPagePreview: true,
		})
	}

}

func SearchNext(pageSize int) tele.HandlerFunc {
	return func(c tele.Context) error {
		return Search(pageSize)(c)
	}
}

func SearchPrev(pageSize int) tele.HandlerFunc {
	return func(c tele.Context) error {
		return Search(pageSize)(c)
	}
}

func SearchSwitchOrder(pageSize int) tele.HandlerFunc {
	return func(c tele.Context) error {
		return Search(pageSize)(c)
	}
}
