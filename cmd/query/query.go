package query

import (
	"github.com/iyear/searchx/app/query"
	"github.com/spf13/cobra"
)

var (
	_query        string
	json          bool
	pn            int
	ps            int
	searchDriver  string
	searchOptions map[string]string
)

var Cmd = &cobra.Command{
	Use:   "query",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		query.Query(searchDriver, searchOptions, _query, pn, ps, json)
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&searchDriver, "driver", "d", "", "search driver used")
	Cmd.PersistentFlags().StringToStringVarP(&searchOptions, "options", "o", make(map[string]string), "options for search driver")
	Cmd.PersistentFlags().StringVarP(&_query, "query", "q", "", "query keyword or statement")
	Cmd.PersistentFlags().IntVar(&pn, "pn", 0, "page number, starting from 0")
	Cmd.PersistentFlags().IntVar(&ps, "ps", 10, "page size")
	Cmd.PersistentFlags().BoolVar(&json, "json", false, "json format output")
}
