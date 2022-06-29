package storage

type Search interface {
	Index(id string, data interface{}) error
	Search(query string, from, size int) []*SearchResult
}

type SearchResult struct {
	Score    float64                `json:"score"`
	Fields   map[string]interface{} `json:"fields"`
	Location SearchLocationMap      `json:"location"`
}

type SearchField = string
type SearchTerm = string
type SearchLocationMap map[SearchField]map[SearchTerm]SearchLocation

type SearchLocation struct {
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
}
