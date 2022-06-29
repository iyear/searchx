package search

import (
	"fmt"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/storage/search/internal/bleve"
)

func New(name string, options map[string]interface{}) (storage.Search, error) {
	switch name {
	case "bleve":
		return bleve.New(options)
	default:
		panic(fmt.Sprintf("%s is a search engine that has not yet been implemented", name))
	}
}
