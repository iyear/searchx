package query

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/iyear/searchx/app/bot/internal/config"
	"github.com/iyear/searchx/pkg/i18n"
	"github.com/iyear/searchx/pkg/storage"
	"github.com/iyear/searchx/pkg/storage/search"
)

type tmplData struct {
	Seq    int
	Score  float64
	Fields map[string]interface{}
}

//go:embed query.tmpl
var tmpl string

func Query(ctx context.Context, cfg string, query string, pn, ps int, jsonFormat bool) error {
	if err := config.Init(cfg); err != nil {
		return err
	}

	_search, err := storage.NewSearch(config.C.Storage.Search.Driver, config.C.Storage.Search.Options)
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
