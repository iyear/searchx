package query

import (
	"errors"
	"fmt"
	"github.com/iyear/searchx/pkg/storage/search"
	"github.com/mitchellh/mapstructure"
)

func Query(driver string, searchOptions map[string]string, query string, pn, ps int, jsonFormat bool) error {
	if driver == "" {
		return errors.New("search driver can not be empty")
	}

	options := make(map[string]interface{})
	if err := mapstructure.WeakDecode(searchOptions, &options); err != nil {
		return err
	}

	_search, err := search.New(driver, options)
	if err != nil {
		return err
	}

	results := _search.Search(query, pn*ps, ps)

	fmt.Printf("%+v", results)
	return nil
}
