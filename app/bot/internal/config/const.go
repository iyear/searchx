package config

import "github.com/iyear/searchx/pkg/storage"

const (
	ContextScope    = "scope"
	ContextLanguage = "lang"
)

const (
	CmdStart = "/start"
)

type SearchOrder struct {
	SortBy []storage.SearchOptionSortByItem
	Text   string
}

// SearchOrders todo(iyear): i18n and refactor
var SearchOrders = []SearchOrder{
	{
		Text: "🔀 Normal",
	},
	{
		SortBy: []storage.SearchOptionSortByItem{{
			Field:   "date",
			Reverse: false,
		}},
		Text: "🔀 Date",
	},
	{
		SortBy: []storage.SearchOptionSortByItem{{
			Field:   "date",
			Reverse: true,
		}},
		Text: "🔀 Date Reverse",
	},
}
