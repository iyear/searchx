package storage

import (
	"fmt"
	"github.com/iyear/searchx/pkg/storage/search"
	"github.com/iyear/searchx/pkg/storage/search/bleve"
)

type Search interface {
	Index(items []*search.Item) error
	Search(query string, options search.Options) []*search.Result
	Get(id string) (*search.Result, error)
}

func NewSearch(name string, options map[string]interface{}) (Search, error) {
	switch name {
	case "bleve":
		return bleve.New(options)
	default:
		panic(fmt.Sprintf("%s is a search engine that has not yet been implemented", name))
	}
}
