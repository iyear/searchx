package query

import (
	"context"
	"github.com/fatih/color"
	"github.com/iyear/searchx/app/bot/query"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var (
	_query string
	json   bool
	pn     int
	ps     int
)

var Cmd = &cobra.Command{
	Use:   "query",
	Short: "Query messages",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
		defer cancel()

		cfg, err := cmd.Flags().GetString("config")
		if err != nil {
			color.Red("value of config flag not found")
			return
		}

		if err := query.Query(ctx, cfg, _query, pn, ps, json); err != nil {
			color.Red("error happens: %v", err)
			return
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&_query, "query", "q", "", "query keyword or statement")
	Cmd.PersistentFlags().IntVar(&pn, "pn", 0, "page number, starting from 0")
	Cmd.PersistentFlags().IntVar(&ps, "ps", 10, "page size")
	Cmd.PersistentFlags().BoolVar(&json, "json", false, "json format output")
}
