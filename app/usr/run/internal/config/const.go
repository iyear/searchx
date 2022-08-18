package config

import (
	"github.com/gotd/td/telegram"
	"github.com/iyear/searchx/global"
	"github.com/iyear/searchx/pkg/storage/search"
	"runtime"
)

var (
	Device = telegram.DeviceConfig{
		DeviceModel:   "SearchX",
		SystemVersion: runtime.GOOS,
		AppVersion:    global.Version,
	}
)

const (
	ContextScope = "scope"
)

type SearchOrder struct {
	SortBy []search.OptionSortByItem
	Text   string
}

// SearchOrders todo(iyear): i18n and refactor
var SearchOrders = []SearchOrder{
	{
		Text: "🔀 Normal",
	},
	{
		SortBy: []search.OptionSortByItem{{
			Field:   "date",
			Reverse: false,
		}},
		Text: "🔀 Date",
	},
	{
		SortBy: []search.OptionSortByItem{{
			Field:   "date",
			Reverse: true,
		}},
		Text: "🔀 Date Reverse",
	},
}

const (
	HighlightSpace = 12
	SenderNameMax  = 6
	ChatNameMax    = 8
)
