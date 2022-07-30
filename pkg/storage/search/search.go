package search

type Item struct {
	ID   string      `json:"id"`
	Data interface{} `json:"data"`
}

type Options struct {
	From   int                `json:"from"`
	Size   int                `json:"size"`
	SortBy []OptionSortByItem `json:"sort_by"`
}

type OptionSortByItem struct {
	Field   string `json:"field"`
	Reverse bool   `json:"reverse"`
}

type Result struct {
	Score    float64                `json:"score"`
	Fields   map[string]interface{} `json:"fields"`
	Location LocationMap            `json:"location"`
}

type Field = string
type Term = string
type LocationMap map[Field]map[Term]Location

type Location struct {
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
}
