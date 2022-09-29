package bleve

import (
	"context"
	"fmt"
	"github.com/blevesearch/bleve/v2"
	"github.com/iyear/searchx/pkg/storage/search"
)

func (b *Bleve) Get(ctx context.Context, id string) (*search.Result, error) {
	s := bleve.NewSearchRequest(bleve.NewDocIDQuery([]string{id}))
	results := b.searchReq(ctx, s, search.Options{})

	if len(results) != 1 {
		return nil, fmt.Errorf("get doc failed, id: %s", id)
	}

	return results[0], nil
}

func (b *Bleve) Search(ctx context.Context, query string, options search.Options) []*search.Result {
	// try query string query
	// if not have results, try wildcard query
	s := bleve.NewSearchRequestOptions(bleve.NewQueryStringQuery(query), options.Size, options.From, false)
	results := b.searchReq(ctx, s, options)
	if len(results) > 0 {
		return results
	}

	s = bleve.NewSearchRequestOptions(bleve.NewWildcardQuery("*"+query+"*"), options.Size, options.From, false)
	return b.searchReq(ctx, s, options)
}

func (b *Bleve) searchReq(ctx context.Context, req *bleve.SearchRequest, options search.Options) []*search.Result {
	req.IncludeLocations = true
	req.Fields = []string{"*"}

	sortby := make([]string, 0, len(options.SortBy))
	for _, item := range options.SortBy {
		if item.Reverse {
			sortby = append(sortby, "-"+item.Field)
			continue
		}
		sortby = append(sortby, item.Field)
	}
	sortby = append(sortby, "-_score")
	req.SortBy(sortby)

	results := make([]*search.Result, 0)

	result, err := b.index.SearchInContext(ctx, req)
	if err != nil {
		return results
	}

	for _, hit := range result.Hits {
		locmap := make(search.LocationMap)
		// copy location map
		for field, locs := range hit.Locations {
			termloc := make(map[search.Term]search.Location)
			for term, loc := range locs {
				for _, l := range loc {
					termloc[term] = search.Location{Start: l.Start, End: l.End}
				}
			}
			locmap[field] = termloc
		}
		results = append(results, &search.Result{
			Score:    hit.Score,
			Fields:   hit.Fields,
			Location: locmap,
		})
	}
	return results
}
