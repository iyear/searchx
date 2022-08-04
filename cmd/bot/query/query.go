package query

import (
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/bot/query"
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
	Short: "Query messages",
	Run: func(cmd *cobra.Command, args []string) {
		if err := query.Query(searchDriver, searchOptions, _query, pn, ps, json); err != nil {
			color.Red("error happens: %v", err)
			return
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&searchDriver, "driver", "d", "", "used search engine driver")
	Cmd.PersistentFlags().StringToStringVarP(&searchOptions, "options", "o", make(map[string]string), "search engine options")
	Cmd.PersistentFlags().StringVarP(&_query, "query", "q", "", "query keyword or statement")
	Cmd.PersistentFlags().IntVar(&pn, "pn", 0, "page number, starting from 0")
	Cmd.PersistentFlags().IntVar(&ps, "ps", 10, "page size")
	Cmd.PersistentFlags().BoolVar(&json, "json", false, "json format output")
}
