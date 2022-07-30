package bleve

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/iyear/searchx/pkg/storage"
)

func (b *Bleve) Search(query string, options *storage.SearchOptions) []*storage.SearchResult {
	// try query string query
	// if not have results, try wildcard query
	search := bleve.NewSearchRequestOptions(bleve.NewQueryStringQuery(query), options.Size, options.From, false)
	results := b.searchReq(search, options)
	if len(results) > 0 {
		return results
	}

	search = bleve.NewSearchRequestOptions(bleve.NewWildcardQuery("*"+query+"*"), options.Size, options.Size, false)
	return b.searchReq(search, options)
}

func (b *Bleve) searchReq(req *bleve.SearchRequest, options *storage.SearchOptions) []*storage.SearchResult {
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

	results := make([]*storage.SearchResult, 0)

	result, err := b.index.Search(req)
	if err != nil {
		return results
	}

	for _, hit := range result.Hits {
		locmap := make(storage.SearchLocationMap)
		// copy location map
		for field, locs := range hit.Locations {
			termloc := make(map[storage.SearchTerm]storage.SearchLocation)
			for term, loc := range locs {
				for _, l := range loc {
					termloc[term] = storage.SearchLocation{Start: l.Start, End: l.End}
				}
			}
			locmap[field] = termloc
		}
		results = append(results, &storage.SearchResult{
			Score:    hit.Score,
			Fields:   hit.Fields,
			Location: locmap,
		})
	}
	return results
}
