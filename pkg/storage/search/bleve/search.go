package bleve

import (
	"github.com/blevesearch/bleve/v2"
	"github.com/iyear/searchx/pkg/storage/search"
)

func (b *Bleve) Search(query string, options *search.Options) []*search.Result {
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

func (b *Bleve) searchReq(req *bleve.SearchRequest, options *search.Options) []*search.Result {
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

	result, err := b.index.Search(req)
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
