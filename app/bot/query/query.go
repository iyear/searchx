package query

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iyear/searchx/pkg/i18n"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/storage/search"
	"github.com/mitchellh/mapstructure"
)

type tmplData struct {
	Seq    int
	Score  float64
	Fields map[string]interface{}
}

//go:embed query.tmpl
var tmpl string

func Query(ctx context.Context, driver string, searchOptions map[string]string, query string, pn, ps int, jsonFormat bool) error {
	if driver == "" {
		return errors.New("search driver can not be empty")
	}

	options := make(map[string]interface{})
	if err := mapstructure.WeakDecode(searchOptions, &options); err != nil {
		return err
	}

	_search, err := storage.NewSearch(driver, options)
	if err != nil {
		return err
	}

	// todo(iyear): add sortBy options
	results := _search.Search(ctx, query, search.Options{
		From: pn * ps,
		Size: ps,
		SortBy: []search.OptionSortByItem{{
			Field:   "date",
			Reverse: true,
		}},
	})

	if jsonFormat {
		b, err := json.MarshalIndent(results, "", "\t")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	}

	data := make([]*tmplData, 0)
	for i, r := range results {
		data = append(data, &tmplData{
			Seq:    i + 1,
			Score:  r.Score,
			Fields: r.Fields,
		})
	}

	t, err := i18n.NewText(tmpl)
	if err != nil {
		return err
	}
	fmt.Println(t.T(data))
	return nil
}
